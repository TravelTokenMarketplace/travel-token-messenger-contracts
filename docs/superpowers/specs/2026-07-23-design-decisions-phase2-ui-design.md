# Design Decisions Implementation — Phase 2 (UI) Design

**Date:** 2026-07-23
**Status:** Approved
**Parent spec:** `docs/superpowers/specs/2026-07-22-design-decisions-implementation-design.md`
**Follows:** phase 1 (contracts), merged as PR #6 (`dev` @ `0fc4e6f`)

## Why this document exists

The parent spec scoped phase 2 as a four-item bullet list under "Phase 2 — UI".
That list was written before phase 1 landed. Verified against `dev` @ `0fc4e6f`
on 2026-07-23, two of its items are already done and the connector item is
larger than the bullet implied. This spec is the reconciled, code-verified
phase-2 plan. It is UI-only: **no contract, ABI, or Go-binding change**, so the
contracts verification recipe is untouched and the gate is `cd ui && yarn sync
&& yarn test` plus `yarn lint`.

Nothing is deployed on any chain (`ignition/deployments/` absent,
`ui/src/contracts/generated/addresses.ts` exports an empty `ADDRESSES`), so this
work is exercised against local/testnet only.

## What phase 2 actually is, after verification

| Parent-spec item | Status on `dev` @ `0fc4e6f` | This spec |
| --- | --- | --- |
| WalletConnect + `safe()` connectors | Not done. `ui/src/wallet/wagmi.ts` wires only `injected()`. | **Section A** — and larger than a config edit (see below). |
| Grant/revoke `GAS_WITHDRAWER_ROLE` affordance | **Already done.** `RolesTab` → `RolesPanel` over all `ACCOUNT_ROLES` (includes `GAS_WITHDRAWER_ROLE`), enumerable, grant/revoke gated on `DEFAULT_ADMIN_ROLE`, reachable via `AccountWorkspace`. | Not redone. |
| `PaymentTokensTab` sentinel labels + drop off-chain boolean | Labels not done; the boolean control never existed. | **Section B** (labels only). |
| `tasks/account.js` CLI sentinel treatment | **Already done** in phase 1 (`payment-token:list` labels both sentinels at `tasks/account.js:471`; `payment:set-offchain` deleted). | Not redone. |
| — (new, not in parent spec) | `BotsTab` still lists `GAS_WITHDRAWER_ROLE` as a bot role — false after Decision 5. | **Section C**. |

## Decisions taken in brainstorming (2026-07-23)

1. **Connector UI: extend the bespoke `ConnectButton`, do not adopt RainbowKit.**
   The UI is developer-facing; the transit-board design and the single app-wide
   `Identicon` (used in `AddressDisplay`, `TokenDisplay`, `ConnectButton`) are
   deliberate. A custom picker keeps both and is less code. RainbowKit would be
   the right call for a consumer-facing, non-crypto-native audience (better
   onboarding and error handling), but that is not this UI. RainbowKit is left
   in `package.json` unused — a later dep-cleanup candidate, not removed now.
2. **WalletConnect project ID: optional env var, graceful when unset.** Read
   `import.meta.env.VITE_WALLETCONNECT_PROJECT_ID`; register the WalletConnect
   connector only when it is present. Safe + injected always work; WalletConnect
   activates once a real ID is set. No secret is committed.

---

## Section A — Wallet connectors (Decision 6)

### The real blocker

`ui/src/wallet/wagmi.ts` configures `connectors: [injected()]` and nothing else.
The parent spec framed the fix as "add two connectors". But the bespoke
`ui/src/components/ConnectButton.tsx` connects with
`connect({ connector: connectors[0] })` — it always grabs the *first* connector
and offers **no picker**. Adding connectors to the config therefore does nothing
user-visible on its own; the connect UI must let the user choose.

### Changes

**`ui/src/wallet/wagmi.ts`**
- Keep `createConfig` (not RainbowKit's `getDefaultConfig`).
- Build the connector list:
  ```ts
  const wcId = import.meta.env.VITE_WALLETCONNECT_PROJECT_ID as string | undefined;
  connectors: [
    injected(),
    safe(),
    ...(wcId ? [walletConnect({ projectId: wcId })] : []),
  ]
  ```
  `safe()` and `walletConnect` come from `wagmi/connectors`. wagmi's
  `walletConnect` connector renders its own QR modal (`showQrModal` defaults on).
  wagmi's `safe()` connector reports ready only inside the Safe-App iframe.
- Keep the existing `ENABLED_CHAINS`-empty guard that throws with the
  `yarn sync` hint — do not regress it.

**`ui/src/components/ConnectButton.tsx`**
- Disconnected state: replace the single-button `connect(connectors[0])` with a
  Headless-UI `Menu` (the app's existing menu primitive — the connected-state
  menu already uses it) listing the available connectors. Each item calls
  `connect({ connector })`.
- Friendly labels + icons mapped from connector `id`: `injected` → "Browser
  Wallet", `walletConnect` → "WalletConnect", `safe` → "Safe". A small
  `connectorMeta(id)` helper keeps the mapping in one place; unknown ids fall
  back to `connector.name`.
- Only surface connectors worth clicking: filter out the Safe connector when it
  is not ready (i.e. not in a Safe-App iframe) so users outside Safe do not get
  a dead entry. If exactly one connector is available, keep the current
  one-click behaviour (no menu) so the common single-wallet case is unchanged.
- Connected state (identicon, copy-address, view-on-explorer, disconnect) is
  **unchanged**.

**Config/docs**
- Add `VITE_WALLETCONNECT_PROJECT_ID` to `ui/.env.example` (create if absent)
  with a comment that it is optional and where to get one (WalletConnect/Reown
  Cloud).
- One line in the UI README/`ui/CLAUDE.md` noting the variable and that
  WalletConnect is inert without it.

### Testing (A)
- `ConnectButton` has no dedicated test today and is rendered only by
  `Layout.tsx`. Add a `ConnectButton.test.tsx` that, with a mocked wagmi
  `useConnect`, asserts: (1) multiple connectors render a picker, (2) a single
  connector renders the one-click button, (3) selecting an item calls `connect`
  with that connector. Keep the wagmi mock local; do **not** pull RainbowKit
  into the test tree.
- Confirm existing `Layout`-touching tests still pass.

---

## Section B — Payment-token sentinels (Decision 4)

### The problem

Since phase 1, `getSupportedTokens()` can contain `address(0)` (native) and
`address(1)` (off-chain). `ui/src/hooks/useTokenMetadata.ts` issues ERC-20
`symbol`/`name`/`decimals` reads for **every** address with `allowFailure:
true`; the sentinel reads fail, so `TokenDisplay` falls back to the raw address
and both render as bare `0x0000…0000` / `0x0000…0001` — meaningless to a partner.
The add flow (`PaymentTokensTab` → `ListManager`) is a single free-text address
input, so adding a sentinel today means typing a magic value.

### Changes

**`ui/src/lib/paymentTokens.ts` (new)** — single source of truth:
- `NATIVE_SENTINEL = "0x0000…0000"`, `OFFCHAIN_SENTINEL = "0x0000…0001"`.
- `paymentTokenLabel(address)` returning `{ symbol, name }` for the two
  sentinels and `undefined` otherwise. Wording mirrors the CLI
  (`tasks/account.js:471`): native → "Native currency", off-chain → "Off-chain
  payment".
- `isSentinel(address)` helper.

**`ui/src/hooks/useTokenMetadata.ts`**
- Exclude sentinels from the multicall read list (their reads always fail),
  then inject their labels into the returned `meta` map so `TokenDisplay`
  renders the name. Non-sentinel behaviour unchanged.

**`ui/src/components/ListManager.tsx`**
- Add an optional `presets?: { label: string; value: string; hint?: string }[]`
  prop. When provided and the caller has the role, render the presets as
  quick-add chips above the free-text input; a preset already present in
  `items` is hidden or disabled. No behaviour change when `presets` is omitted.

**`ui/src/pages/tabs/PaymentTokensTab.tsx`**
- Pass `presets` for the two sentinels (labels from `paymentTokens.ts`) so a
  partner adds Native / Off-chain by clicking, never by typing an address.

### Testing (B)
- Unit test `paymentTokenLabel`/`isSentinel`.
- `useTokenMetadata`: sentinels are not read and come back labelled.
- `PaymentTokensTab`/`ListManager`: presets render, a present sentinel is
  suppressed, clicking a preset calls `addSupportedToken` with the sentinel.

---

## Section C — `BotsTab` role list (Decision 5)

### The problem

`ui/src/pages/tabs/BotsTab.tsx:21` defines
`BOT_ROLES = ["MESSENGER_BOT_ROLE", "BOOKING_OPERATOR_ROLE", "GAS_WITHDRAWER_ROLE"]`
with a comment that "a fully-provisioned bot holds all three roles", and the
per-badge tooltip warns a missing role means the bot "may not function fully".
After Decision 5, `addMessengerBot` no longer grants `GAS_WITHDRAWER_ROLE`, so
that role reads as expected-and-missing (amber warning) on every correctly
provisioned bot.

### Changes

**`ui/src/pages/tabs/BotsTab.tsx`**
- `BOT_ROLES = ["MESSENGER_BOT_ROLE", "BOOKING_OPERATOR_ROLE"]`; fix the comment.
- Reword the missing-role tooltip so it no longer implies gas-withdrawer is
  required; gas-withdrawal is opt-in and managed in the Roles tab.
- Leave the `removeMessengerBot` "Remove" tooltip as-is: the on-chain call still
  defensively revokes `GAS_WITHDRAWER_ROLE` if present, so "revokes all bot
  roles" remains accurate.

### Testing (C)
- Update any `BotsTab` test asserting three badges to expect two.

---

## Out of scope / explicitly not redone

- Grant/revoke `GAS_WITHDRAWER_ROLE` affordance — already live via `RolesTab`.
- CLI sentinel labelling and `payment:set-offchain` removal — landed in phase 1.
- Guided Safe creation in `CreateAccount.tsx` — deferred in the parent spec;
  still deferred.
- Removing RainbowKit from `ui/package.json` — noted as a later cleanup.

## Files touched

New: `ui/src/lib/paymentTokens.ts`, `ui/src/components/ConnectButton.test.tsx`,
`ui/.env.example` (if absent), a `paymentTokens`/`useTokenMetadata` test file.
Edited: `ui/src/wallet/wagmi.ts`, `ui/src/components/ConnectButton.tsx`,
`ui/src/hooks/useTokenMetadata.ts`, `ui/src/components/ListManager.tsx`,
`ui/src/pages/tabs/PaymentTokensTab.tsx`, `ui/src/pages/tabs/BotsTab.tsx`, README
/ `ui/CLAUDE.md` line.

## Verification

```bash
cd ui
yarn sync
yarn test      # existing suite + new tests green
yarn lint      # 0 errors
yarn build     # type-check / bundle sane
```

No contract, ABI, docgen, or Go-binding regeneration — this phase changes none
of them.
