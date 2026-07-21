# Design — Pre-Deploy Rework: Identity, Events, and a Hash-Native Account

**Date:** 2026-07-21
**Status:** Approved, ready for planning
**Predecessor:** `2026-07-21-predeploy-hardening-design.md` (shipped)
**Source material:** `docs/decisions/2026-07-21-technical-backlog.md` §§1–5, 7, 8

---

## Why now

Nothing is deployed. `ignition/deployments/` is empty, on every chain.

The technical backlog was written assuming a Base Sepolia deployment was
imminent, and it prices most of its items as "cheap now, expensive later". With
the deployment postponed, that window is still open, and it is the only reason
this work is scheduled ahead of the six business decisions in
`docs/decisions/2026-07-21-contract-design-decisions.md`.

Concretely, three costs the backlog anticipated are currently **zero**:

| Backlog concern                                                 | Cost today |
| --------------------------------------------------------------- | ---------- |
| §2 storage layout freeze — "an edit now, a migration after"     | Zero. No deployed state to migrate. |
| §1 log-format discontinuity — "the indexer must understand both formats indefinitely" | Zero. No indexed history exists. |
| §3 ABI break across bot, app-service, and UI simultaneously      | Near-zero. All three need a version bump from the hardening PR regardless. |

Every one of those becomes permanent on the day of the first deploy. This spec
therefore batches the backlog's "Spec 2" and "Spec 3" into a single change,
which the backlog's own sequencing kept separate only to limit consumer
churn — churn that does not exist yet.

## Scope

Everything inside `travel-token-messenger-contracts`: Solidity, tests, Go
bindings, `ui/`, and `tasks/`. The repo stays coherent and CI-green in one
merge.

`travel-token-messenger-bot` and `travel-token-matrix-app-service` are **out of
scope** and are handled by the migration deliverable in section G.

**Explicitly still deferred:** backlog §6 (bot role separation, blocked on
Decision 5) and all six items in the decision document.

## Corrections to the backlog document

The backlog is seed material, not gospel. Five of its claims did not survive
contact with the code, and the design departs from it accordingly. They are
recorded here because the backlog file itself stays as-written — it is a dated
artifact of the review that produced it.

1. **§2's "one wasted slot per account" is false.**
   `TTMAccountInfo { bool isTTMAccount; address creator; }`
   (`contracts/manager/TTMAccountManager.sol:97-99`) is 1 + 20 = 21 bytes and
   already packs into a single slot. Removing the bool saves no storage. §2's
   value is eliminating divergence, not gas.

2. **`TTMACCOUNT_ROLE` is not an authorization gate.** Across `contracts/`,
   `test/`, `ui/`, and `tasks/`, it is `_grantRole`'d exactly once at creation
   (`TTMAccountManager.sol:280`) and otherwise only ever *enumerated*
   (`ui/src/hooks/useMyAccounts.ts:42`, `tasks/manager.js:252`). Nothing calls
   `hasRole(TTMACCOUNT_ROLE, …)` to permit an action. This reframes §2 — see
   section A.

3. **§5 is already satisfied.** `ServiceRegistry._unregisterServiceName` emits
   `ServiceUnregistered(serviceName, serviceHash)` at
   `contracts/partner/ServiceRegistry.sol:117`. The backlog's claim that
   "nothing tells the ecosystem it was deprecated" was wrong, and its proposed
   `ServiceDeprecated` event is redundant. §5 reduces to indexing the hash,
   which folds into section B.

4. **§1.1 miscounts the affected events.** There are **eight** events declaring
   `string indexed serviceName`, at `contracts/account/TTMAccount.sol:164-174`,
   not seven at `:167-177`. The two `WantedService*` events are the ones the
   backlog missed, and they need the same treatment.

5. **§3's `TTMLens` contract is withdrawn.** See section C.

---

## A. Identity becomes a registry, not a role

**Backlog §2, plus the enumeration half of §8.**

### The reframing

The backlog treats this as "two sources of truth, pick a winner", and proposes
making `TTMACCOUNT_ROLE` authoritative while overriding `grantRole` and
`revokeRole` to revert for it.

Given correction 2 above, that is the wrong shape. Account identity is a
registry that was modelled as an AccessControl role, and *being a role* is
precisely what exposes it to `grantRole`. Making the role authoritative
preserves the mismodelling and then bolts on guards to contain it — guards that
must also cover `renounceRole`, and that a future refactor can quietly drop.

### The change

`TTMAccountManager` drops `TTMACCOUNT_ROLE` and the `TTMAccountInfo` struct.
Inside the existing ERC-7201 namespace:

```solidity
EnumerableSet.AddressSet _ttmAccounts;
mapping(address account => address creator) _ttmAccountCreator;
```

| Function                                | Behaviour |
| --------------------------------------- | --------- |
| `isTTMAccount(address)`                 | `_ttmAccounts.contains(a)` — signature unchanged |
| `getTTMAccountCreator(address)`         | unchanged signature, reads the new mapping |
| `getTTMAccountCount()`                  | new |
| `getTTMAccounts()`                      | new — replaces `getRoleMembers(TTMACCOUNT_ROLE)` |
| `getTTMAccounts(uint256 offset, uint256 limit)` | new — absorbs §8's pagination item |

`_createTTMAccount` becomes the sole writer, and no external mutator exists at
all. The divergence class disappears structurally rather than by guard, so no
`grantRole`/`revokeRole`/`renounceRole` overrides are needed.

`BookingToken.isTTMAccount` (`contracts/booking-token/BookingToken.sol:700`) is
unaffected — it already reads through `ITTMAccountManager`.

### Consumers in this repo

- `ui/src/hooks/useMyAccounts.ts` — two role calls become one `getTTMAccounts()`.
- `ui/src/components/AccountValidityNotice.tsx` — **deleted**. It exists solely
  to explain a divergence that can no longer occur.
- `tasks/manager.js:15,252` and `tasks/account.js:781` — updated to the new
  getters. The `role:all` comment at `tasks/manager.js:15` about
  `TTMACCOUNT_ROLE` slowing output becomes obsolete.

---

## B. Events

**Backlog §1 and §5.**

### Where names live

Section C makes `TTMAccount` hash-native, so it no longer knows service names.
It cannot emit them without reintroducing exactly the manager staticcall that
section C removes. The two backlog items therefore have to be resolved
together, and the resolution is: **names are emitted by the naming authority.**

`ServiceRegistry` is that authority, so it carries the readable name:

```solidity
event ServiceRegistered(bytes32 indexed serviceHash, string serviceName);
event ServiceUnregistered(bytes32 indexed serviceHash, string serviceName);
```

Both currently declare `(string serviceName, bytes32 serviceHash)` with neither
indexed (`ServiceRegistry.sol:44-45`). Indexing the hash makes them filterable;
keeping the name in the data section means **any consumer can build a complete
name↔hash map from logs alone, with no `eth_call` at all.**

`TTMAccount`'s eight service events (`TTMAccount.sol:164-174`) become hash-only:

```solidity
event ServiceAdded(bytes32 indexed serviceHash);
event ServiceRemoved(bytes32 indexed serviceHash);
event ServiceRestrictedRateUpdated(bytes32 indexed serviceHash, bool restrictedRate);
event ServiceCapabilitiesUpdated(bytes32 indexed serviceHash);
event ServiceCapabilityAdded(bytes32 indexed serviceHash, string capability);
event ServiceCapabilityRemoved(bytes32 indexed serviceHash, string capability);
event WantedServiceAdded(bytes32 indexed serviceHash);
event WantedServiceRemoved(bytes32 indexed serviceHash);
```

Capability payloads stay `string`: capabilities are free-form operator text,
not registry entries, so there is nothing to resolve a hash against (see
section D).

This fixes the underlying §1.1 defect. Today `string indexed serviceName`
stores only `keccak256(value)` in the topic and **nothing** in the data
section — the name is unrecoverable from the log. The new form is honest about
carrying a hash, and pairs it with a registry event that actually publishes the
mapping.

### Note on the UI

The backlog implies the UI's hash-resolution workaround disappears. It does
not, and this spec does not claim it will. Hash-in-topic remains; what changes
is that it stops being an accident the UI compensates for
(`ui/src/lib/activity/catalog.ts`'s `SERVICE_HASH_EVENTS` plus
`ui/src/hooks/useAccountActivity.ts:32-65`) and becomes the designed interface,
backed by a name-carrying registry event. The special-case machinery is
replaced by one shared resolver (section E), not removed.

### Other events

- `TTMAccountCreated(address indexed account, address indexed creator, address indexed admin)`
  — currently emits only the account (`TTMAccountManager.sol:142`), forcing
  indexers into a follow-up `getTTMAccountCreator` call per account.
- §1.3's two silent setters gain events: `BookingToken.setManagerAddress`
  (repoints the entire authorization oracle) and
  `setMinExpirationTimestampDiff` (changes a mint-time validation rule).

---

## C. Hash-native account

**Backlog §3, with the lens withdrawn.**

### Why no lens

The backlog proposed a stateless `TTMLens` view contract to do name resolution
and joins. Reading the actual consumers, it does not earn its keep.

The bot touches services in three places:

| Site | Call | What a lens would save |
| ---- | ---- | ---------------------- |
| `internal/messaging/service_registry.go:38` | `GetSupportedServices()`, once at startup | One `eth_call`, once per process |
| `pkg/ttm_accounts/ttm_accounts.go:189` | `IsServiceSupported(name)`, **per inbound message** | Nothing — a lens adds a hop on a hot path |
| `internal/eventlistener/subscriber/subscriber.go:127` | `WatchServiceAdded` | Nothing — this is §1's problem, not a lens's |

So the lens saves the bot a single startup call. Everything else the backlog
credited to §3 comes from going hash-native: the per-message staticcall in the
second row is deleted by hash-native alone, and the third row is fixed by
section B.

Against that: a fourth deployed contract, its address in every consumer's
config and in the UI's chain-enablement filter, an ownership and upgrade story,
and staleness risk if it caches a manager address that `setManagerAddress` can
repoint.

The resolution table is 63 entries. Any client can hold it in memory. If
read-batching later proves painful, the answer is Multicall3 — already deployed
on Base and Base Sepolia, with nothing for us to maintain. (Verify the
canonical address at the time of use.)

### The change

Every service function on `TTMAccount` takes `bytes32 serviceHash` instead of
`string memory serviceName`. This removes:

- the three private resolution helpers and their manager staticcalls
  (`TTMAccount.sol:527-551`);
- the per-item fan-out in `getSupportedServices()` (`:553-565`), which today
  costs two cross-contract staticcalls per service;
- `removeAllServices()`'s hash→name→hash round-trip (`:469-476`).

`getSupportedServices()` returns `(bytes32[] serviceHashes, Service[] services)`,
with a paginated variant alongside it. `addWantedServices` / `removeWantedServices`
take `bytes32[]`.

Expected saving: roughly 250 lines and ~2 KiB of bytecode.

### One staticcall stays, deliberately

`addService` must still verify the hash corresponds to a *registered* service,
so it keeps a single check against the manager. Today that validation is
implicit in `getRegisteredServiceHash` reverting on an unknown name; hash-native
makes it explicit.

Dropping it would let an account advertise a service that does not exist in the
registry. It is a write path, called rarely, and the invariant is worth one
staticcall. **Reads** lose the manager dependency entirely; writes keep it.

### Developer UX

`tasks/*.js` hash at the CLI boundary, so operators continue to type service
names. The string API is a presentation concern, and this is where it belongs.

---

## D. Cheap items

**Backlog §4, §7, §8.**

### Accepted

| Item | Change |
| ---- | ------ |
| §4 | Drop `onlyRole(BOOKING_OPERATOR_ROLE)` from the `TTMAccount.recordExpiration` wrapper (`:374-376`). The gate protects nothing — `BookingToken.recordExpiration` (`:632-656`) is already `public` and unrestricted, and correctly so: it only marks a reservation expired once `block.timestamp` has genuinely passed. Add NatSpec stating the permissionlessness is deliberate. |
| §8 | Override `supportsInterface` in `TTMAccount` (`:383`) to report `type(IERC721Receiver).interfaceId`. The contract implements `IERC721Receiver` (`:54`) but reports `false`, so counterparties doing capability detection conclude it cannot receive an NFT. |
| §8 | Collapse the six near-identical cancellation wrappers (`BookingToken.sol:721-850`) behind one `_requireBoughtAndParties(tokenId) returns (address owner, address supplier)`. ABI-identical, ~90 lines and ~1 KiB saved. |
| §8 | `nonReentrant` on `finalizeCancellation`, for consistency with `buyReservedToken`. |

### Rejected

**Packing the `Service` struct via `bytes32[] capabilities`**
(`PartnerConfiguration.sol:26-33`). Capabilities are arbitrary operator-typed
strings, rendered directly in `ui/src/pages/tabs/ServicesTab.tsx:156,213` and
editable there. Unlike service names they have no registry, so there is nothing
to reverse-resolve a hash against — hashing them would make them permanently
unreadable to save one storage slot. The O(n) `keccak256` comparison in
`_removeServiceCapability` (`:201`) stays; it iterates a handful of entries.

This is the one item where the hash-native direction of section C does **not**
generalise, and the distinction is exactly the presence of a registry.

### Tests (§7)

Five gaps, all backfilling behaviour the fee removal made load-bearing:

1. **`createTTMAccount` is permissionless.** Highest value in the list. It pins
   the accepted risk from Decision 1, so the day that decision is implemented a
   test fails loudly instead of the change passing silently. Camino enforced KYC
   at the chain level and that guarantee vanished in the move to Base without any
   code changing.
2. **No external path adds to the account set.** Replaces the backlog's
   "role/mapping divergence" test, which section A makes unrepresentable —
   assert instead that `isTTMAccount` is true only for factory-created accounts.
3. **`removeAllServices`** (`TTMAccount.sol:469-476`) — the multi-service
   iteration path is unexercised.
4. **`BookingTokenAddressUpdated` / `TTMAccountImplementationUpdated`** — emitted
   but never asserted.
5. **Cancellation refund to a contract that cannot receive ETH** — would have
   caught the stuck-refund bug behind Decision 3.

Tests 1 and 2 assert the *absence* of a restriction. Such tests do not catch
bugs you already know about; they pin a property so a future refactor or
environment change fails loudly rather than silently.

---

## E. UI

One shared service-catalog resolver module replaces the current special-case
handling. It seeds from a single `getAllRegisteredServiceNames()` call, derives
hashes locally (`serviceHash == keccak256(abi.encodePacked(name))`), and stays
current from `ServiceRegistered` / `ServiceUnregistered` logs.

Consumers of that module: `ServicesTab.tsx`, the activity feed
(`useAccountActivity.ts`, replacing `SERVICE_HASH_EVENTS`), and
`lib/receipt.ts` (which also picks up creator and admin from the widened
`TTMAccountCreated`).

Plus section A's changes to `useMyAccounts.ts` and the deletion of
`AccountValidityNotice.tsx`.

---

## F. Contract size

The 22.5 KiB gate is the binding constraint in this repo. Post-hardening:
`TTMAccountManager` 12.800, `TTMAccount` 21.371, `BookingToken` 21.552.

Section C takes roughly 2 KiB off `TTMAccount`; section D's wrapper dedup takes
roughly 1 KiB off `BookingToken`. Both currently sit within about 1 KiB of the
gate, so this work is also the cheapest available relief on the constraint most
likely to block the next feature.

The plan must record measured sizes before and after, not estimates.

---

## G. Consumer migration deliverable

Two files, written as the **final task** of implementation, from the merged
ABI:

| File | Audience |
| ---- | -------- |
| `BOT-MIGRATION.md` | a fresh `travel-token-messenger-bot` session |
| `APP-SERVICE-MIGRATION.md` | a fresh `travel-token-matrix-app-service` session |

Both live in the **local workspace parent folder**, alongside `REBRANDING.md`,
`TODOS.md`, and `CONTRACTS-NEXT.md`, and are **not committed to any repo** —
same rationale as those files: they reference the multi-repo layout.

They are written at the end, not from this design, because a migration guide
predicted from a spec drifts during implementation, and a fresh session has no
way to tell a stale instruction from a live one. Each must contain the exact
final signatures, the Go binding type names as generated, the module version to
pin, and a checklist of call sites in that repo.

Known call sites to cover for the bot: `pkg/ttm_accounts/ttm_accounts.go:189`
(`IsServiceSupported` — becomes hash-based, hash cached at startup),
`internal/messaging/service_registry.go:38` (`GetSupportedServices` — now
returns hashes, needs a local resolver seeded from the manager),
`internal/eventlistener/subscriber/subscriber.go:127` (`WatchServiceAdded` —
new event shape), and `tests/e2e/blockchain/client.go:163,191`.

---

## Testing strategy

- The existing 134 tests must stay green, adapted to the new signatures rather
  than deleted. A test removed during a signature change is a test that was
  quietly load-bearing — the fee removal already taught this lesson here, by
  deleting ~284 lines of manager tests and replacing none.
- Section D's five new tests.
- `yarn lint`, the `docs` clean-tree job, and the `abi/` drift guard all green.
- The UI's own test suite green, including `ui/src/lib/activity/catalog.test.ts`.

## Repo traps this work will hit

Carried forward from the hardening pass; each has already cost a CI cycle here.

1. **`yarn compile` does not reliably run docgen** — the compile cache
   short-circuits it despite `docgen.runOnCompile: true`, so `docs/` silently
   goes stale and CI's clean-tree check fails. Every symbol-adding or -removing
   change must run
   `yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi`.
   This spec adds and removes a great many symbols.
2. **Go bindings embed bytecode**, so even a comment change shifts them via
   Solidity's metadata hash. Regenerate with `scripts/generate_go_abi.sh` after
   any contract edit. It does its own `rm -rf node_modules && yarn install`, and
   `/` is ~97% full — route `TMPDIR`, `GOCACHE`, `GOMODCACHE`, `GOTMPDIR`, and
   `YARN_CACHE_FOLDER` to `/hgst`.
3. **Anything under `docs/` must be prettier-formatted**, including this file —
   `yarn docgen` ends in `prettier --write docs/`, so an unformatted file there
   dirties the tree and fails both the `docs` job and `lint`.
4. **`vars.get(NAME)` with no default throws at config load** (`HH1201`),
   breaking *every* Hardhat command rather than just the one that needs it.
   Always supply a fallback.
5. **Removing a contract requires updating the generator's `ARTIFACTS` list**,
   or the orphaned Go package is invisible to both regeneration and CI's
   `git status --porcelain` drift check. This bit us once with
   `go/contracts/servicefeetoken/`. No contract is removed by this spec, but
   section C removes many *symbols*.
6. **Verify comments against code before acting on them.** Stale comments have
   twice produced wrong conclusions in this repo.

## Out of scope

- The six business decisions in
  `docs/decisions/2026-07-21-contract-design-decisions.md`. Decisions 1 and 6
  remain mainnet blockers; this spec does not resolve them, and section D's
  first test deliberately pins the current permissionless behaviour rather than
  changing it.
- Backlog §6 (bot role separation), blocked on Decision 5.
- The Base Sepolia deployment itself, postponed by the user. This spec's entire
  justification is that it lands **before** that deployment.
- Any change to `travel-token-messenger-bot` or
  `travel-token-matrix-app-service` beyond the migration documents in section G.
