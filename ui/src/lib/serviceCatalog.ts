import { keccak256, toBytes } from "viem";

/**
 * Service names and hashes are two views of one thing. Contracts store and emit
 * `bytes32` hashes; people read names. The registry is the only authority that
 * knows both, so we seed from it once and resolve locally thereafter — which is
 * why `TTMAccount` needs no name-resolution code and its reads cost no extra
 * cross-contract calls.
 */
export interface ServiceCatalog {
  /** Keyed by lowercase hash, because callers get hashes from several sources. */
  nameByHash: Map<string, string>;
  hashByName: Map<string, `0x${string}`>;
}

/**
 * Mirrors `keccak256(abi.encodePacked(serviceName))` in ServiceRegistry. For a
 * single string argument, `encodePacked` is just the raw UTF-8 bytes.
 */
export function hashServiceName(name: string): `0x${string}` {
  return keccak256(toBytes(name));
}

/** Builds the bidirectional map from the registry's list of registered names. */
export function buildServiceCatalog(names: string[]): ServiceCatalog {
  const nameByHash = new Map<string, string>();
  const hashByName = new Map<string, `0x${string}`>();

  for (const name of names) {
    const hash = hashServiceName(name);
    nameByHash.set(hash.toLowerCase(), name);
    hashByName.set(name, hash);
  }

  return { nameByHash, hashByName };
}

/** Resolves a hash to its registered name, or undefined if it is unknown. */
export function serviceNameForHash(catalog: ServiceCatalog, hash: string): string | undefined {
  return catalog.nameByHash.get(hash.toLowerCase());
}

/**
 * De-duplicated hashes from `hashes` that the catalog cannot resolve.
 *
 * The catalog is seeded from `getAllRegisteredServiceNames()`, which only lists
 * *currently registered* services. `ServiceRegistry` deliberately keeps a name
 * mapping around after a service is unregistered (so accounts that already
 * adopted it can still resolve the name), so a hash can be legitimate and
 * still miss the catalog. This is the input to the bounded per-hash fallback
 * batch: in the common case every hash resolves here and the result is empty,
 * so the fallback batch stays disabled.
 */
export function unresolvedHashes(catalog: ServiceCatalog, hashes: string[]): string[] {
  const seen = new Set<string>();
  const result: string[] = [];
  for (const hash of hashes) {
    const key = hash.toLowerCase();
    if (seen.has(key)) continue;
    seen.add(key);
    if (serviceNameForHash(catalog, hash) === undefined) result.push(hash);
  }
  return result;
}
