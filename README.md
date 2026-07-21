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

```bash
yarn hardhat ignition deploy ignition/modules/messenger.js \
  --network base_sepolia --parameters ignition/base_sepolia_parameters.json
```

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
