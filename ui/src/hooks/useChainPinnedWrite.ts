import { useCallback } from "react";
import { useWriteContract } from "wagmi";
import { useActiveChain } from "../wallet/activeChain";

type WriteContract = ReturnType<typeof useWriteContract>;

/**
 * `useWriteContract`, with the app's active chain pinned onto every write.
 *
 * wagmi submits on whatever chain the connector happens to be on and asserts
 * nothing when `chainId` is absent. Reads are keyed by `activeChainId`, and the
 * active chain only follows the wallet while the wallet is on an enabled chain
 * — so a wallet left on another network leaves the two silently disagreeing:
 * the UI keeps rendering one chain's state while writes go to a different one,
 * at addresses that hold our contracts only on the chain being displayed. A
 * call to a codeless address still succeeds, so any value sent along is simply
 * lost while the transaction reports as confirmed.
 *
 * Pinning `chainId` makes wagmi reject the mismatch (or prompt a switch)
 * instead of submitting to the wrong chain. The pin is applied last, so a
 * caller cannot override it, and `options` is forwarded only when supplied so
 * the wrapper is indistinguishable from wagmi's own function at the call site.
 *
 * Every user-facing write must go through this hook rather than
 * `useWriteContract` directly — that is what stops a new call site from
 * reintroducing the problem. See `ui/CLAUDE.md`.
 */
export function useChainPinnedWrite(): WriteContract {
  const { writeContract, writeContractAsync, ...rest } = useWriteContract();
  const { activeChainId } = useActiveChain();

  // wagmi's write variables are a generic union that a plain object spread
  // cannot re-satisfy, so the pinned object is handed back through `never`.
  // The public signature below is unchanged, so call sites stay fully typed.
  const pinnedWriteContract = useCallback(
    ((variables, options) => {
      const pinned = { ...variables, chainId: activeChainId } as never;
      return options === undefined ? writeContract(pinned) : writeContract(pinned, options as never);
    }) as WriteContract["writeContract"],
    [writeContract, activeChainId],
  );

  const pinnedWriteContractAsync = useCallback(
    ((variables, options) => {
      const pinned = { ...variables, chainId: activeChainId } as never;
      return options === undefined ? writeContractAsync(pinned) : writeContractAsync(pinned, options as never);
    }) as WriteContract["writeContractAsync"],
    [writeContractAsync, activeChainId],
  );

  return { ...rest, writeContract: pinnedWriteContract, writeContractAsync: pinnedWriteContractAsync };
}
