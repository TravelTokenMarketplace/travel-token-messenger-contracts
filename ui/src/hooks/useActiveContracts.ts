import { type Address } from "viem";
import { BOOKINGTOKEN_ABI, TTMACCOUNT_ABI, MANAGER_ABI, getContractsForChain } from "../contracts";
import { useActiveChain } from "../wallet/activeChain";

export function useActiveContracts() {
  const { activeChainId } = useActiveChain();
  const resolved = getContractsForChain(activeChainId);
  return {
    chainId: activeChainId,
    supported: Boolean(resolved),
    manager: resolved?.manager as Address | undefined,
    bookingToken: resolved?.bookingToken as Address | undefined,
    ttmAccountImpl: resolved?.ttmAccountImpl as Address | undefined,
    managerAbi: MANAGER_ABI,
    ttmAccountAbi: TTMACCOUNT_ABI,
    bookingTokenAbi: BOOKINGTOKEN_ABI,
  };
}
