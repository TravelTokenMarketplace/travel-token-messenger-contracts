# Rebrand: Camino Messenger → Travel Token Messenger

**Date:** 2026-07-14
**Status:** Approved
**Scope:** This repository (`travel-token-messenger-contracts`) only. Other
ecosystem repositories (protocol, bot, matrix app service) are rebranded
separately.

## Context

The project is rebranding from "Camino Messenger" to "Travel Token Messenger".
A fresh deployment on Base Sepolia will follow the rebrand, so nothing
on-chain constrains the rename: contract names, role hashes, event
signatures, and service-name hashes may all change freely.

Decisions made during brainstorming:

- **Full rename** — identifiers, files, roles, and prose all change (not a docs-only rebrand).
- **`TTM` prefix** replaces `CM` (Travel Token Messenger).
- **`ttm.` service namespace** replaces `cmp.` in all docs, examples, tasks, UI placeholders, and registered service names. Assumes the protocol repo adopts the same prefix.
- **Camino infrastructure removed entirely** — no legacy network configs kept.
- **Full git history** is pushed to the new repo; the old repo becomes an archive.
- **Execution: phased commits**, each compiling and passing tests.

## 1. Rename map

Identifier renames (case-aware; applied to contracts, tests, tasks, scripts,
ignition, UI, go bindings):

| Old | New |
|---|---|
| `CMAccount*` (incl. `CMAccountManager`, `CMAccountStorage`, `CMAccountInfo`, events `CMAccountCreated`/`CMAccountUpgraded`/…, errors `CMAccountInvalidImplementation`/…) | `TTMAccount*` |
| `ICMAccount`, `ICMAccountManager` | `ITTMAccount`, `ITTMAccountManager` |
| `CMACCOUNT_ROLE` | `TTMACCOUNT_ROLE` (new keccak hash — acceptable, fresh deployment) |
| `cmAccount*` variables (`cmAccountProxy`, `cmAccountAbi`, `cmAccountImpl`, …) | `ttmAccount*` |
| `CaminoMessengerModule` (ignition module) | `TravelTokenMessengerModule` |
| `cmp.services.` / `cmp.service` / `cmp.types` | `ttm.services.` / `ttm.service` / `ttm.types` |
| "CMP" / "Camino Messenger Protocol" (prose) | "Travel Token Messenger Protocol" |
| "Camino Messenger" / "CaminoMessenger" | "Travel Token Messenger" / "TravelTokenMessenger" |
| `camino-messenger-contracts` / `camino-messenger-ui` | `travel-token-messenger-contracts` / `travel-token-messenger-ui` |
| Test fixture names (`CaminoAdmin`, `cmServiceAdmin`, `cmBL`, `cmAccountAdmin`, …) | `TTMAdmin`, `ttmServiceAdmin`, and equivalents |

File renames via `git mv` (history follows):

- `contracts/account/CMAccount.sol` → `TTMAccount.sol`; `ICMAccount.sol` → `ITTMAccount.sol`
- `contracts/manager/CMAccountManager.sol` → `TTMAccountManager.sol`; `ICMAccountManager.sol` → `ITTMAccountManager.sol`
- `contracts/manager/test/CMAccountManagerTest.sol` → `TTMAccountManagerTest.sol`
- `test/CMAccount.test.js` → `TTMAccount.test.js`; `test/CMAccountManager.test.js` → `TTMAccountManager.test.js`
- `go/contracts/cmaccount/` → `ttmaccount/`; `go/contracts/cmaccountmanager/` → `ttmaccountmanager/` (regenerated)

Deletions:

- Hardhat networks `camino` (500) and `columbus` (501), caminoscan
  verification config (`etherscan.customChains` entries), and
  `CAMINO_*`/`COLUMBUS_*` configuration variables in `hardhat.config.js`.
- `ignition/camino_parameters.json`, `ignition/columbus_parameters.json`.
- **All** of `ignition/deployments/` — chain-500, chain-501, and the stale
  chain-84532 journal (Ignition keys futures by module/contract name; the
  Camino-named journal would conflict with the fresh deployment).
- README Camino mainnet/Columbus address tables, Camino badge, and Camino links.

Left untouched:

- Dated design docs under `docs/superpowers/specs/` (historical records).
- Git history (pushed as-is to the new repo).

## 2. Repo & remote setup (before any rebrand commits)

1. Create `TravelTokenMarketplace/travel-token-messenger-contracts` on GitHub
   (`gh repo create`, visibility matching the old repo).
2. `git remote rename origin old`;
   `git remote add origin git@github.com:TravelTokenMarketplace/travel-token-messenger-contracts.git`.
3. Push `dev` (and any other long-lived branches) plus tags to the new
   origin so the rebranding PR has a base; continue work on `rebranding`.

## 3. Phased commits on `rebranding` (each compiles and passes tests)

1. **Contracts** — Solidity identifier and file renames, NatSpec/comment
   updates. `yarn compile` passes. Run `yarn hardhat export-abi` in the same
   commit so `abi/` stays in sync (CLAUDE.md rule).
2. **Tests + tooling** — `test/`, `tasks/`, `scripts/`,
   `services/00_initial.json` (service names → `ttm.`), `utils/`,
   `examples/`, ignition module `ignition/modules/messenger.js`, and
   `ignition/base_sepolia_parameters.json`. `yarn test` green.
3. **UI** — re-run `ui/scripts/sync-contracts.ts`, rename package to
   `travel-token-messenger-ui`, update all UI text/config/placeholders.
   UI build and tests green.
4. **Docs & config** — README rewrite (Base-first, no Camino tables),
   `CLAUDE.md`, `DATA_PROTECTION.md`, root `package.json` name/repo URL,
   `hardhat.config.js` network cleanup, deletion of old ignition
   deployments/parameters, `yarn docgen` regeneration, go bindings
   regenerated via `scripts/generate_go_abi.sh`, `go.mod` module path
   updated to the new repo. Also `.github/` workflow references if any.

## 4. Verification & landing

- `yarn compile && yarn test && yarn lint`
- UI: `yarn build` and tests in `ui/`
- `go build ./...` in `go/contracts`
- Final sweep: `grep -ri camino` across the repo (excluding
  `node_modules/`, `artifacts/`, `cache/`, `.git/`) — expected hits only in
  `docs/superpowers/specs/` historical documents.
- PR `rebranding → dev` on the new repo.
- Fresh Base Sepolia deployment happens after merge as its own step:
  `ignition deploy` with `base_sepolia_parameters.json`, contract
  verification, service registration (`ttm.` names), README address table
  update.

## 5. Out of scope

- Rebranding the other ecosystem repositories (protocol, bot, matrix app
  service). They follow the same naming decisions (§ Context and §1) but are
  handled as separate projects.
- Any on-chain migration from existing deployments — the Base Sepolia
  deployment is fresh; old Camino/Columbus deployments are abandoned.
