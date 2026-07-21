import { useMemo } from "react";
import { type Abi } from "viem";
import { useReadContract } from "wagmi";
import { useActiveContracts } from "./useActiveContracts";
import { useActiveChain } from "../wallet/activeChain";
import { buildServiceCatalog } from "../lib/serviceCatalog";

/**
 * Reads the registry's full service list once and derives the name↔hash map
 * locally. One eth_call replaces the per-hash resolution round-trips the
 * activity feed used to make.
 */
export function useServiceCatalog() {
  const { manager, managerAbi } = useActiveContracts();
  const { activeChainId } = useActiveChain();

  const { data, isLoading } = useReadContract({
    chainId: activeChainId,
    address: manager,
    abi: managerAbi as Abi,
    functionName: "getAllRegisteredServiceNames",
    query: { enabled: Boolean(manager) },
  });

  const catalog = useMemo(() => buildServiceCatalog((data as string[]) ?? []), [data]);

  return { catalog, isLoading };
}
