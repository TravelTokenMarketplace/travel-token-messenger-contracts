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
