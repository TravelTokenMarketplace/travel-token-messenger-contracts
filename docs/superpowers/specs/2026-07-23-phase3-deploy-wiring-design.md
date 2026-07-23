# Phase 3 — deployment wiring (Safe custody of privileged roles)

**Status:** design approved 2026-07-23. No deploy in this session — the
deliverables are the decided role topology, a `base_sepolia_parameters.json`
policy, a role-handoff Hardhat task with local tests, and README runbook
updates. Base Sepolia deployment follows in a later session; there is no
schedule pressure.

Phase 3 is the last of the six design-decisions phases. Phases 1 (contract
changes D2–D5, PR #6) and 2 (UI, PR #7) are merged; `dev` @ `46adcf3`. **Nothing
is deployed on any chain.** This phase writes **no Solidity** — it is deploy
wiring: a parameters policy, a task, its tests, and docs.

---

## The six decisions this closes

Decision 6 (admin-key custody = Safe multisig) is the subject. The concrete
choices, settled with the user on 2026-07-23:

| Question | Decision |
|---|---|
| Signers | Solo operator. The testnet Safe is a **placeholder** for a real mainnet signer set (TBD). |
| Threshold | **1-of-2** — two owner keys the user controls; either can act, losing one does not brick the Safe. |
| Pauser custody | A **dedicated hot EOA**, separate from both Safe owner keys. |
| Escalation | Pause fast on the hot key; recover via the Safe. Resolved as **Option B** — see below. |

### The pause/unpause constraint (why Option B)

`pause()` and `unpause()` are gated by the **same** `PAUSER_ROLE` on both
contracts (`TTMAccountManager.sol:213,220`; `BookingToken.sol:378,385`). There
is no admin-gated unpause. So "only the Safe may unpause" cannot be enforced
without a contract change. The three options were:

- **A** — hot key holds `PAUSER` only; "recover via Safe" means the remediation
  *upgrade* goes through the Safe, not the unpause. No code.
- **B (chosen)** — `PAUSER` granted to **both** the hot key and the Safe. The
  Safe *can* unpause (a literal recovery path); the hot key can pause fast. The
  hot key can also unpause — same role — so the "only Safe unpauses" discipline
  is convention, not enforced. No code.
- **C** — split `pause`/`unpause` into two roles so the separation is enforced
  on-chain. Rejected for this phase: it reopens contract work (new role, ABI,
  tests, Go bindings) and breaks the "phase 3 = no contract changes" framing.
  Recorded as a mainnet consideration.

---

## Role topology (steady state, after handoff)

- **Safe** — SafeL2 singleton + CompatibilityFallbackHandler, 2 owner keys the
  user controls, threshold 1, on Base Sepolia, created through the Safe web app.
- **Hot pauser** — a third EOA, distinct from both Safe owner keys.
- **Deployer** — `BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY`. Holds **nothing** after
  handoff, unless `--keep-deployer-as-default-admin` is passed (see below).

| Contract | Role | Holder(s) after handoff |
|---|---|---|
| `TTMAccountManager` | `DEFAULT_ADMIN_ROLE` | Safe |
| `TTMAccountManager` | `UPGRADER_ROLE` | Safe |
| `TTMAccountManager` | `VERSIONER_ROLE` | Safe |
| `TTMAccountManager` | `PAUSER_ROLE` | **Safe + hot pauser** |
| `TTMAccountManager` | `SERVICE_REGISTRY_ADMIN_ROLE` | Safe |
| `BookingToken` | `DEFAULT_ADMIN_ROLE` | Safe |
| `BookingToken` | `UPGRADER_ROLE` | Safe |
| `BookingToken` | `PAUSER_ROLE` | **Safe + hot pauser** |
| `BookingToken` | `MIN_EXPIRATION_ADMIN_ROLE` | Safe (only if the 60s default is ever changed) |

`SERVICE_REGISTRY_ADMIN_ROLE` and `MIN_EXPIRATION_ADMIN_ROLE` are **not** in the
current README step-8 handoff list. They are granted to the deployer during
setup, so leaving them there would contradict "deployer fully de-privileged."
This phase folds them into the handoff.

### Safe contract addresses (Base Sepolia, Safe v1.4.1)

Use `SafeL2`, not the non-L2 `Safe` singleton — Base is an L2, and `SafeL2`
emits per-transaction events that Safe's Transaction Service indexes without
node tracing. The non-L2 singleton works on-chain but can render wrong in Safe's
own UI. Set `CompatibilityFallbackHandler` at setup for EIP-1271.

| Contract | Address |
|---|---|
| `SafeL2` | `0x29fcB43b46531BcA003ddC8FCB67FFE91900C762` |
| `CompatibilityFallbackHandler` | `0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99` |
| `SafeProxyFactory` | `0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67` |

Addresses are deterministic across chains but deployment is per-chain — confirm
before any Base **mainnet** use.

---

## Approach — deploy as deployer, hand off in one task (Approach H)

Chosen over splitting role assignment between the parameters file and the
handoff (Approach P). Rationale:

- **One auditable source of truth.** The entire final topology lives in the
  handoff task, not split across a JSON file and a runbook step. This is exactly
  the footgun the README warns about — a role set in neither place "silently
  stays on the deployer," with nothing in the deploy output warning you.
- **Trap #7 is satisfied for free.** The Ignition module itself calls
  `setAccountImplementation` and `setBookingTokenAddress`, both
  `onlyRole(VERSIONER_ROLE)`, as account 0 during the run. With every role
  defaulting to the deployer, versioner is the deployer through the whole run —
  no special-casing, no mid-module revert.
- **The deployer needs `DEFAULT_ADMIN` through setup anyway** to grant
  `SERVICE_REGISTRY_ADMIN_ROLE` and register services. Keeping all roles on the
  deployer until a single explicit handoff is the natural sequence.

The extra grant/revoke transactions cost trivially more gas on testnet.

`base_sepolia_parameters.json` stays `{}` (all roles default to account 0).

---

## Deploy sequence

1. **Prerequisites (manual, no code):**
   - Create the Safe on Base Sepolia via the Safe web app — `SafeL2` singleton,
     `CompatibilityFallbackHandler`, 2 owner keys you control, threshold 1.
     Record the Safe address.
   - Provision the hot-pauser EOA (distinct from both Safe owner keys).
   - Fund the deployer EOA with testnet ETH.
   - `yarn hardhat vars set BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY` / `ETHERSCAN_API_KEY`.
2. **Deploy** — `yarn hardhat ignition deploy … --parameters
   ignition/base_sepolia_parameters.json` with the file left as `{}`. All roles
   land on the deployer; the module's versioner calls succeed.
3. **Setup as deployer** — grant `SERVICE_REGISTRY_ADMIN_ROLE` to the deployer,
   register the 63 canonical services. (The BookingToken `PAUSER` grant that was
   README step 6 moves into the handoff task.)
4. **Handoff** — `yarn hardhat roles:handoff --safe <addr> --pauser <addr>
   --network base_sepolia`. See the task spec below.
5. **Finalize** — commit `ignition/deployments/chain-84532/`, fill the README
   address table, bump the contracts Go module in the bot and matrix-app-service.

---

## The `roles:handoff` task — phase-3 code deliverable

A Hardhat task registered in `hardhat.config.js`, living under `tasks/`.

### Parameters

- `--safe <address>` (required) — the Safe that receives the admin roles.
- `--pauser <address>` (required) — the hot pauser EOA.
- `--keep-deployer-as-default-admin` (flag, default off) — skip **only** the
  final `DEFAULT_ADMIN_ROLE` renounce, on both contracts. A **testnet-only
  recovery hatch**: it leaves the deployer as a break-glass admin alongside the
  Safe if the Safe setup turns out wrong. Every other deployer role (upgrader,
  versioner, pauser, service-registry-admin) is still renounced regardless of
  this flag. Using it on mainnet would defeat moving admin to the Safe — the
  task must print a prominent warning when it is set.

The task reads the manager and BookingToken proxy addresses from
`ignition/deployments/chain-84532/deployed_addresses.json`.

### Sequence (strict order)

1. **Grant** the full steady-state topology:
   - Manager → Safe: `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`, `VERSIONER_ROLE`,
     `PAUSER_ROLE`, `SERVICE_REGISTRY_ADMIN_ROLE`.
   - Manager → hot pauser: `PAUSER_ROLE`.
   - BookingToken → Safe: `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`, `PAUSER_ROLE`.
   - BookingToken → hot pauser: `PAUSER_ROLE`.
2. **Verify gate** — read back and assert the Safe holds every role in the
   topology table on both contracts, and the hot pauser holds `PAUSER_ROLE` on
   both. **If any is missing, abort before renouncing anything.** This is the
   guard against bricking a singleton — the manager cannot be redeployed without
   losing all account registrations.
3. **Renounce** the deployer's roles via `renounceRole(role, deployer)`
   (self-renounce), `DEFAULT_ADMIN_ROLE` **last** on each contract:
   - Manager: `VERSIONER`, `UPGRADER`, `PAUSER`, `SERVICE_REGISTRY_ADMIN`, then
     `DEFAULT_ADMIN` (unless `--keep-deployer-as-default-admin`).
   - BookingToken: `UPGRADER`, `PAUSER`, then `DEFAULT_ADMIN` (unless the flag).
4. **Final verify** — assert the end state and print it:
   - Safe holds all admin roles on both contracts.
   - Hot pauser holds `PAUSER_ROLE` (and nothing else) on both.
   - Deployer holds **nothing**, or **only `DEFAULT_ADMIN_ROLE`** on both when
     the flag is set.

### Safety properties

- **Idempotent.** Re-runnable: skip a grant whose role the target already holds;
  skip a renounce the deployer has already done. Re-running after a partial
  failure converges to the target state.
- **Fail-safe ordering.** No renounce happens before the verify gate in step 2
  passes. `DEFAULT_ADMIN` renounced last so an aborted run never strands the
  contract without an admin.
- **`MIN_EXPIRATION_ADMIN_ROLE`** is granted to the Safe only if it was ever
  granted to the deployer (it is optional in the runbook). The task grants it to
  the Safe when the deployer holds it, then renounces the deployer's — otherwise
  it is a no-op.

---

## Tests

Against a **local** Hardhat deployment only — no chain interaction, honoring
"no deploy this session":

1. Deploy the Ignition module (or an equivalent fixture) to the local network so
   all roles start on the deployer.
2. Run the handoff logic pointing `--safe` at a stand-in address and `--pauser`
   at another. Assert the full steady-state topology, deployer holds nothing.
3. `--keep-deployer-as-default-admin`: assert the deployer retains
   `DEFAULT_ADMIN_ROLE` on both contracts and nothing else, Safe holds
   everything.
4. **Verify-gate abort:** simulate a missing grant (e.g. stub the grant of one
   Safe role) and assert the task reverts/aborts **before any renounce** — the
   deployer must still hold its roles afterward.
5. **Idempotency:** run the handoff twice; the second run is a no-op and the
   topology is unchanged.

---

## README runbook changes

Fold the above into the existing `Deploy (Hardhat Ignition)` runbook:

- **Prerequisite:** create the Safe and provision the hot pauser (new, before
  step 1).
- **Step 6** (manual BookingToken `PAUSER` grant) → removed; handled by the
  task.
- **Step 8** → replaced by `yarn hardhat roles:handoff …`, documenting the flag,
  the verify gate, and the full role list (including `SERVICE_REGISTRY_ADMIN_ROLE`
  and `PAUSER` to both Safe and hot key). Keep the "verify before renounce — the
  manager is a singleton" warning.

---

## Recorded for mainnet (out of scope now)

- **Real signer set + threshold.** Today's 1-of-2 is a placeholder; mainnet
  needs distinct signers and a real threshold.
- **Decision 1** (`createTTMAccount` is permissionless) must be resolved before
  Base mainnet — see `docs/decisions/2026-07-21-contract-design-decisions.md`
  and the "bounds a compromised bot key" finding in the phase-1 spec.
- **Option C** (split `pause`/`unpause`) if the halt/recovery separation should
  be enforced on-chain rather than by convention.
- `--keep-deployer-as-default-admin` is testnet-only; do not use it on mainnet.

---

## Traps this phase must respect

From `notes/contracts-next.md` (workspace) — verify each before relying on it:

- **Trap #7** — `managerVersioner` must stay the deployer through the Ignition
  run. Satisfied by Approach H (empty parameters).
- **Singleton brick risk** — the manager cannot be redeployed without losing
  account registrations. The verify gate + `DEFAULT_ADMIN`-last ordering exist
  for this.
- **Go bindings / docgen drift** — this phase touches no Solidity, so the
  bindings and `docs/` generated from contracts do not change. But the task and
  README still go through `yarn lint`, and anything under `docs/` must be
  prettier-formatted.
