import { useMemo } from "react";
import { type Abi } from "viem";
import { useReadContract, useReadContracts } from "wagmi";
import { useActiveContracts } from "./useActiveContracts";
import { useActiveChain } from "../wallet/activeChain";
import { buildServiceCatalog, serviceNameForHash, unresolvedHashes } from "../lib/serviceCatalog";

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

/**
 * Resolves a set of service hashes to names, merging two sources: the catalog
 * (one call, covers every currently-registered service — the common case),
 * plus a bounded per-hash fallback for whatever the catalog missed.
 *
 * The catalog alone would regress a service that was unregistered on the
 * manager after an account adopted it: `ServiceRegistry` deliberately keeps
 * that name mapping around (so such an account can still resolve it), but
 * `getAllRegisteredServiceNames()` only lists currently-registered services.
 * The fallback batch covers exactly that gap — in practice it's empty and
 * stays disabled, since deprecated services are rare.
 *
 * Used by the activity feed and by the Services tab (supported + wanted), so
 * every hash-to-name lookup in the UI goes through one place.
 */
export function useResolvedServiceNames(hashes: string[]) {
  const { catalog, isLoading: catalogLoading } = useServiceCatalog();
  const { manager, managerAbi, chainId } = useActiveContracts();
  const abi = managerAbi as Abi;

  const missing = useMemo(() => unresolvedHashes(catalog, hashes), [catalog, hashes]);

  const { data: fallbackReads, isLoading: fallbackLoading } = useReadContracts({
    allowFailure: true,
    contracts: missing.map(
      (hash) => ({ chainId, address: manager, abi, functionName: "getServiceNameByHash", args: [hash] }) as const,
    ),
    query: { enabled: Boolean(manager) && missing.length > 0 },
  });

  const resolve = useMemo(() => {
    const fallbackByHash = new Map<string, string>();
    missing.forEach((hash, i) => {
      const name = fallbackReads?.[i]?.result as string | undefined;
      if (name) fallbackByHash.set(hash.toLowerCase(), name);
    });
    return (hash: string): string | undefined =>
      serviceNameForHash(catalog, hash) ?? fallbackByHash.get(hash.toLowerCase());
  }, [catalog, missing, fallbackReads]);

  return { resolve, isLoading: catalogLoading || (missing.length > 0 && fallbackLoading) };
}
