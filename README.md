# Travel Token Messenger Contracts

[![CI](https://github.com/TravelTokenMarketplace/travel-token-messenger-contracts/actions/workflows/ci.yaml/badge.svg)](https://github.com/TravelTokenMarketplace/travel-token-messenger-contracts/actions/workflows/ci.yaml)

Smart contracts for the Travel Token Messenger ecosystem: per-partner
accounts (`TTMAccount`), a factory/registry (`TTMAccountManager`), and an
ERC-721 `BookingToken` representing bookings, targeting the Base network.

## Deployed Contracts

The contracts are deployed with a fresh state on Base Sepolia (testnet).
Addresses will be recorded here after each deployment.

| Base Sepolia (84532)       | Address |
| -------------------------- | ------- |
| _pending fresh deployment_ | —       |

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
# 6. Grant PAUSER_ROLE on BookingToken to the operations key

# 7. Verify on Basescan
yarn hardhat ignition verify chain-84532
```

Marking the two `ERC1967Proxy` addresses as **proxies** on Basescan is a
separate manual step in the Basescan UI.

> **Overriding `managerAdmin` does NOT cascade to the other roles.** Hardhat
> Ignition 0.15.8 cannot resolve a module parameter used as another parameter's
> default (`_resolveDefaultValue` only recurses into `AccountRuntimeValue`), so
> `managerPauser`, `managerUpgrader`, `managerVersioner`, `bookingAdmin`, and
> `bookingUpgrader` each default to **account 0 — the deployer key** —
> independently of `managerAdmin`.
>
> If you point `managerAdmin` at a Safe and forget the others, those roles
> silently stay on the deployer key. Nothing in the dry-run or deploy output
> warns you.
>
> **`managerVersioner` is the one exception — leave it as the deployer.** The
> module itself calls `setAccountImplementation` and `setBookingTokenAddress`,
> both `onlyRole(VERSIONER_ROLE)`, and both execute as account 0 during the
> deploy. Pointing `managerVersioner` at a Safe in the parameters file makes
> those calls revert mid-module, leaving a partially-configured manager on
> chain. Transfer `VERSIONER_ROLE` to the Safe in the role-handoff step (step 8) instead, after the module has finished running.
>
> **Set `managerPauser`, `managerUpgrader`, `bookingAdmin`, and
> `bookingUpgrader` explicitly in `base_sepolia_parameters.json`** if they
> should differ from the deployer, and verify role membership on-chain after
> deploying, before step 8.

**8. Hand off admin roles.** Transfer `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`,
and `VERSIONER_ROLE` to the Safe, verify the Safe can act, and only **then**
renounce the deployer's roles. The manager is a singleton — renouncing before
verifying is unrecoverable.

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
