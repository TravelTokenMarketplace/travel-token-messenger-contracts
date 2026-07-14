import { hasContracts } from "../contracts";

export interface AppChain {
  id: number;
  name: string;
  enabled: boolean;
  rpcUrl: string;
  explorerUrl: string;
  nativeCurrency: { name: string; symbol: string; decimals: number };
}

export const APP_CHAINS: AppChain[] = [
  {
    id: 8453,
    name: "Base",
    enabled: true,
    rpcUrl: "https://mainnet.base.org",
    explorerUrl: "https://basescan.org",
    nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  },
  {
    id: 84532,
    name: "Base Sepolia",
    enabled: true,
    rpcUrl: "https://base-sepolia.drpc.org",
    explorerUrl: "https://sepolia.basescan.org",
    nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  },
];

// A chain is usable only if marked enabled AND its contracts are deployed.
export const ENABLED_CHAINS: AppChain[] = APP_CHAINS.filter((c) => c.enabled && hasContracts(c.id));

export function explorerUrlFor(chainId: number): string | undefined {
  return APP_CHAINS.find((c) => c.id === chainId)?.explorerUrl;
}
