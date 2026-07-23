import { http, type Chain } from "viem";
import { createConfig } from "wagmi";
import { injected, safe, walletConnect } from "wagmi/connectors";
import { ENABLED_CHAINS, type AppChain } from "../config/chains";

function toViemChain(c: AppChain): Chain {
  return {
    id: c.id,
    name: c.name,
    nativeCurrency: c.nativeCurrency,
    rpcUrls: { default: { http: [c.rpcUrl] } },
    blockExplorers: { default: { name: c.name, url: c.explorerUrl } },
  };
}

const viemChains = ENABLED_CHAINS.map(toViemChain);

// ENABLED_CHAINS is empty when no enabled chain has synced contract addresses
// (e.g. a missing/stale generated/ directory). Fail loudly here instead of
// letting the non-empty-tuple cast below hide it as an opaque wagmi error.
if (viemChains.length === 0) {
  throw new Error(
    "No usable chains: ENABLED_CHAINS is empty. Run `yarn sync` to generate contract addresses, or enable a chain with a deployment in config/chains.ts.",
  );
}

// WalletConnect is optional: it needs a project id from WalletConnect/Reown
// Cloud. Without VITE_WALLETCONNECT_PROJECT_ID the connector is simply not
// registered, and injected + Safe still work.
const wcProjectId = import.meta.env.VITE_WALLETCONNECT_PROJECT_ID as string | undefined;

const connectors = [
  injected(),
  safe(),
  ...(wcProjectId ? [walletConnect({ projectId: wcProjectId })] : []),
];

export const wagmiConfig = createConfig({
  chains: viemChains as [Chain, ...Chain[]],
  connectors,
  transports: Object.fromEntries(ENABLED_CHAINS.map((c) => [c.id, http(c.rpcUrl)])),
});
