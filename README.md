# Travel Token Messenger Contracts

[![CI](https://github.com/TravelTokenMarketplace/travel-token-messenger-contracts/actions/workflows/ci.yaml/badge.svg)](https://github.com/TravelTokenMarketplace/travel-token-messenger-contracts/actions/workflows/ci.yaml)

Smart contracts for the Travel Token Messenger ecosystem: per-partner
accounts (`TTMAccount`), a factory/registry (`TTMAccountManager`), and an
ERC-721 `BookingToken` representing bookings, targeting the Base network.

## Deployed Contracts

The contracts are deployed with a fresh state on Base Sepolia (testnet).
Addresses will be recorded here after each deployment.

| Base Sepolia (84532)          | Address                                      |
| ----------------------------- | -------------------------------------------- |
| `TTMAccountManager` (proxy)   | `0xFE6587D20F8F9a57823B6aB62a19Fc857202Ca03` |
| `TTMAccountManager` (impl)    | `0xEa8e753D5262fDc7D1D2bA2c2DfB89086B69F931` |
| `BookingToken` (proxy)        | `0xB9ac3D8898e8bE5481a0DDCb7692692979794efC` |
| `BookingToken` (impl)         | `0xe5F52581fA62c0A1d83A798f1621530d3620D0eA` |
| `TTMAccount` (implementation) | `0x792bDd0C6a2a58DfC055b03d664149977FffDDc1` |
| `BookingTokenOperator` (lib)  | `0xCbbdebde31ff19B1541Ab316f0FA456f3e79B3Dc` |

All six are verified on [Basescan](https://sepolia.basescan.org/).

## Getting Started

### Clone & install

```bash
git clone git@github.com:TravelTokenMarketplace/travel-token-messenger-contracts.git
cd travel-token-messenger-contracts
yarn install
```

### Configuration variables

RPC URLs and deployer keys come from Hardhat configuration variables:

```bash
yarn hardhat vars set BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY
yarn hardhat vars set BASE_SEPOLIA_URL        # optional, has a default
yarn hardhat vars set ETHERSCAN_API_KEY       # for contract verification
```

### Common commands

```bash
yarn compile   # compile contracts (runs contract-sizer + docgen)
yarn test      # hardhat test suite (REPORT_GAS=true)
yarn lint      # prettier + eslint + solhint
yarn hardhat export-abi   # regenerate abi/ (consumed by ui/)
```

### Deploy (Hardhat Ignition)

The Ignition module deploys the contracts but leaves the system **inert** —
`initialize` grants only DEFAULT_ADMIN/PAUSER/UPGRADER/VERSIONER, so no
services can be registered until roles are granted. Follow every step.

**Prerequisite — create custody keys.** Before deploying:

- Create the Safe on Base Sepolia via the Safe web app: `SafeL2` singleton
  (`0x29fcB43b46531BcA003ddC8FCB67FFE91900C762`) with the
  `CompatibilityFallbackHandler` (`0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99`),
  the owner set you control, and your chosen threshold. Record the Safe address.
- Provision a dedicated hot **pauser** EOA, separate from the Safe owner keys.

`base_sepolia_parameters.json` stays `{}` — every role deploys onto the deployer
key and is handed off in step 8. This keeps `managerVersioner` on the deployer
through the Ignition run (the module's `setAccountImplementation` /
`setBookingTokenAddress` calls need `VERSIONER_ROLE` as account 0).

```bash
# 1. Configuration
yarn hardhat vars set BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY
yarn hardhat vars set ETHERSCAN_API_KEY
yarn hardhat vars set BASE_SEPOLIA_URL       # optional; a public default exists

# 2. Deploy
yarn hardhat ignition deploy ignition/modules/messenger.js \
  --network base_sepolia --parameters ignition/base_sepolia_parameters.json

# 3. Grant the service-registry admin role (0 members after deploy)
yarn hardhat manager role:grant --role SERVICE_REGISTRY_ADMIN_ROLE \
  --address <deployer> --network base_sepolia

# 4. Register the 63 canonical services (~11.5M gas total)
yarn hardhat manager services:register \
  --json ./services/00_initial.json --network base_sepolia

# 5. Optional: grant MIN_EXPIRATION_ADMIN_ROLE to change the 60s default

# 6. Verify on Basescan
yarn hardhat ignition verify chain-84532
```

Marking the two `ERC1967Proxy` addresses as **proxies** on Basescan is a
separate manual step in the Basescan UI.

> **Why the parameters file stays `{}` (Approach H).** Every role deploys onto
> the deployer key and moves to the Safe in step 8 — a single, auditable
> handoff. Do **not** set roles in `base_sepolia_parameters.json`. Splitting
> custody between the parameters file and the handoff is fragile: Hardhat
> Ignition 0.15.8 cannot resolve a module parameter used as another parameter's
> default (`_resolveDefaultValue` only recurses into `AccountRuntimeValue`), so
> `managerPauser`, `managerUpgrader`, `managerVersioner`, `bookingAdmin`, and
> `bookingUpgrader` do **not** follow `managerAdmin` — each would silently stay
> on the deployer with nothing in the dry-run or deploy output to warn you.
>
> **`managerVersioner` must stay the deployer regardless.** The module itself
> calls `setAccountImplementation` and `setBookingTokenAddress`, both
> `onlyRole(VERSIONER_ROLE)`, as account 0 during the deploy. Pointing
> `managerVersioner` at a Safe in the parameters file makes those calls revert
> mid-module, leaving a partially-configured manager on chain. The role-handoff
> step (step 8) moves `VERSIONER_ROLE` to the Safe after the module has
> finished running.

**8. Hand off privileged roles to the Safe.**

```bash
yarn hardhat roles handoff --network base_sepolia \
  --safe <safe-address> --pauser <hot-pauser-address>
```

Before sending any transaction the task **preflights the three principals**: the
deployer, Safe and hot pauser must be three distinct, non-zero addresses; the
Safe must be a contract that answers `getOwners()`/`getThreshold()`; and the hot
pauser must be an EOA that holds no administrative role on either contract.

The task then prints the Safe's **singleton** (read from the proxy's storage slot
0), its owner set and its threshold, and stops short of granting anything until
you have read them. Answering the Safe interface proves an interface, not
provenance — **check the singleton against the `SafeL2` address in the
prerequisites above**; the task reports it rather than enforcing an allowlist, so
that our tooling carries no hard-coded per-chain Safe deployment addresses. Each of these is a one-way mistake —
a mistyped EOA passed as `--safe` would receive every administrative role and
pass the verify gate, and `--safe` equal to the deployer would strip the
manager's last admin during the renounce loop.

The task then grants the Safe `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`, `VERSIONER_ROLE`,
`PAUSER_ROLE` and `SERVICE_REGISTRY_ADMIN_ROLE` on the manager and
`DEFAULT_ADMIN_ROLE`/`UPGRADER_ROLE`/`PAUSER_ROLE` on BookingToken (plus
`MIN_EXPIRATION_ADMIN_ROLE` if the deployer was granted it in step 5); grants the
hot pauser `PAUSER_ROLE` on both; **verifies the Safe holds every role before it
renounces anything** (the manager is a singleton — an incomplete grant must not
strand it without an admin); then renounces the deployer's roles,
`DEFAULT_ADMIN_ROLE` last. It is idempotent and safe to re-run.

Pass `--keep-deployer-as-default-admin` to keep the deployer as a break-glass
recovery admin. **Testnet only** — the task refuses the flag on any network
outside `hardhat`, `localhost` and `base_sepolia`, so it cannot reach Base
mainnet, where it would be permanent.

**9. Commit `ignition/deployments/chain-84532/`.** The UI reads
`deployed_addresses.json` and filters enabled chains by whether contracts
exist, so until this is committed the UI has zero enabled chains.

**10. Fill in the address table** at the top of this README.

**11. Bump the contracts Go module** in the bot and matrix-app-service.

> **Known limitation (Base Sepolia).** `createTTMAccount` is currently
> permissionless — anyone can create an account. On Camino this was prevented
> by chain-level KYC, which does not exist on Base. Accepted for testnet
> (blast radius is spam, not funds); must be resolved before Base mainnet.
> See `docs/decisions/2026-07-21-contract-design-decisions.md`, Decision 1.

## Contracts

- **`TTMAccountManager`** — factory + registry + role management. Creates
  `ERC1967Proxy` instances of `TTMAccount` and tracks the canonical
  implementation, booking token address, and the service registry
  (services are referenced by `ttm.services.<package>.<version>.<Name>`).
- **`TTMAccount`** — a Travel Token Messenger account: manages funds,
  messenger bots, supported/wanted services, payment tokens, public keys,
  gas-money withdrawals, and minting/buying of Booking Tokens.
- **`BookingToken`** — ERC-721 representing bookings, with cancellation
  flows and an operator library.

Contracts are UUPS upgradeable (OpenZeppelin upgradeable contracts).
Generated API documentation lives in [`docs/`](docs/index.md).

## Management UI

A wallet-connected management UI for these contracts lives in
[`ui/`](ui/) — see its README/CLAUDE docs for commands.

## License

Licensed under the GNU Lesser General Public License v3.0 — see
[LICENSE.md](LICENSE.md).
