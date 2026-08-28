import { useAccount } from "wagmi";
import { APP_CHAINS } from "../config/chains";
import { useActiveChain } from "../wallet/activeChain";

export interface ChainMismatch {
  /** True when a wallet is connected but is not on the chain the app is showing. */
  mismatched: boolean;
  walletChainId?: number;
  activeChainId: number;
  /** Sentence naming both networks. Undefined unless `mismatched`. */
  reason?: string;
}

function chainName(id: number | undefined): string {
  if (id === undefined) return "an unrecognised network";
  return APP_CHAINS.find((c) => c.id === id)?.name ?? `chain ${id}`;
}

/**
 * Detects a wallet pointed at a different chain than the one the app is
 * displaying.
 *
 * The active chain only follows the wallet while the wallet sits on an enabled
 * chain, so moving the wallet elsewhere leaves the two silently disagreeing:
 * the UI keeps rendering the old chain's state and nothing about the page says
 * the next write would land somewhere else. Callers use this to refuse the
 * write rather than let it through.
 *
 * A connected wallet whose chain we cannot read counts as a mismatch — the safe
 * default is to block, not to submit and hope.
 */
export function useChainMismatch(): ChainMismatch {
  const { isConnected, chainId: walletChainId } = useAccount();
  const { activeChainId } = useActiveChain();

  const mismatched = Boolean(isConnected) && walletChainId !== activeChainId;

  return {
    mismatched,
    walletChainId,
    activeChainId,
    reason: mismatched
      ? `Your wallet is on ${chainName(walletChainId)} but this app is showing ${chainName(activeChainId)}. ` +
        `Switch your wallet to ${chainName(activeChainId)} to continue.`
      : undefined,
  };
}
