# Travel Token Messenger Contracts

Solidity smart contracts for the Travel Token Messenger ecosystem, built with Hardhat. A wallet-connected management UI lives in `ui/` (see `ui/CLAUDE.md`).

## Layout

- `contracts/` — Solidity sources (Solidity 0.8.24, optimizer `runs: 1000`, `evmVersion: cancun`)
    - `manager/` — `TTMAccountManager` (factory + registry + roles), `ITTMAccountManager`
    - `account/` — `TTMAccount` (per-partner account: bots, tokens, services, pubkeys, withdrawals), `GasMoneyManager`, `ITTMAccount`
    - `booking-token/` — `BookingToken` (ERC-721) + cancellable/operator extensions
    - `partner/` — shared base contracts: `PartnerConfiguration`, `ServiceRegistry`
    - `test/` — mocks/helpers used by tests
- `tasks/` — Hardhat tasks for management: `manager.js`, `account.js`, and `roles.js` (all registered in `hardhat.config.js`). `manager.js`/`account.js` are mirrored by the UI; `roles.js` is a separate `roles` scope for privileged-role administration (deploy-time custody handoff), which the UI deliberately does not expose.
- `ignition/` — Hardhat Ignition deployment: `modules/messenger.js`, per-network `*_parameters.json`, and recorded deployments in `deployments/chain-<id>/`
- `abi/` — exported ABIs (see ABIs note below)
- `scripts/`, `services/`, `utils/`, `examples/`, `go/` — supporting code
- `docs/` — generated contract docs + design specs/plans under `docs/superpowers/`
- `ui/` — the React management UI (separate toolchain; has its own `CLAUDE.md`)

## Architecture notes

- Contracts are **UUPS upgradeable** (OpenZeppelin upgradeable + `@openzeppelin/hardhat-upgrades`). `TTMAccountManager.createTTMAccount` deploys an `ERC1967Proxy` per account and records it (with its creator) in the manager's account registry (`isTTMAccount`/`getTTMAccounts`) — the only writer.
- **Access control** is role-based (`AccessControlEnumerable`). `DEFAULT_ADMIN_ROLE` is `0x00…`; other roles are `keccak256(name)`. A "messenger bot" is an address granted two roles at once via `addMessengerBot`: `MESSENGER_BOT_ROLE` and `BOOKING_OPERATOR_ROLE`. `GAS_WITHDRAWER_ROLE` is not part of that grant — it must be granted separately by `DEFAULT_ADMIN_ROLE`.
- Services are referenced by name (`ttm.services.<package>.<version>.<Name>`) and must be registered in the manager's `ServiceRegistry` before an account can support/want them.

## Commands

- `yarn compile` — compile contracts (also runs `contract-sizer` and `docgen`)
- `yarn test` — run the Hardhat test suite (`REPORT_GAS=true`)
- `yarn lint` — Prettier + ESLint + Solhint; `yarn format` to auto-fix with Prettier
- `yarn docgen` — regenerate `docs/`
- `yarn hardhat export-abi` — regenerate `abi/` (see below)
- Deploy with Hardhat Ignition — the single deploy command is **not** sufficient by itself; it leaves the system inert (no roles granted beyond DEFAULT_ADMIN/PAUSER/UPGRADER/VERSIONER, zero services registered). Follow the full runbook in the README's [Deploy (Hardhat Ignition)](README.md#deploy-hardhat-ignition) section, including role grants, service registration, and the `managerAdmin`-does-not-cascade caveat.

## Networks

Defined in `hardhat.config.js`: `base` (8453), `base_sepolia` (84532), plus `localhost`. RPC URLs and deployer keys come from Hardhat **configuration variables** (`vars.get(...)`), e.g. `BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY`, `BASE_SEPOLIA_URL`, `ETHERSCAN_API_KEY` — set them with `yarn hardhat vars set <NAME>`.

## ABIs (important)

There are two ABI locations and they serve different purposes:

- `artifacts/` — Hardhat's compiler output. **Hardhat and tests use this.**
- `abi/` — a flattened export produced by `hardhat-abi-exporter` via `yarn hardhat export-abi`. **The UI consumes this** (`ui/scripts/sync-contracts.ts` reads it).

After changing contracts, run `yarn compile` and then `yarn hardhat export-abi` so `abi/` (and therefore the UI) stays in sync. `abi/` is committed; don't hand-edit it.
