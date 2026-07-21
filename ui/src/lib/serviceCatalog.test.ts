import { describe, expect, it } from "vitest";
import { keccak256, toBytes } from "viem";
import { buildServiceCatalog, hashServiceName, serviceNameForHash, unresolvedHashes } from "./serviceCatalog";

describe("serviceCatalog", () => {
  const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";

  // Independently computed with pycryptodome's Keccak-256 (NOT Node's hashlib
  // sha3_256, which is the different, NIST-padded SHA3 variant):
  //   python3 -c "from Crypto.Hash import keccak; h = keccak.new(digest_bits=256); \
  //     h.update('ttm.services.accommodation.v1alpha.AccommodationSearchService'.encode()); \
  //     print('0x' + h.hexdigest())"
  // -> 0x371ec7bcfbe014e2671831f3266bf9f6477ea12fe9bb7da1413e6939bd197931
  const knownHash = "0x371ec7bcfbe014e2671831f3266bf9f6477ea12fe9bb7da1413e6939bd197931";

  it("hashes a service name the same way the contracts do", () => {
    // Contracts use keccak256(abi.encodePacked(serviceName)), which for a lone
    // string argument is just the keccak of its UTF-8 bytes.
    expect(hashServiceName(name)).toBe(keccak256(toBytes(name)));
  });

  it("matches an independently computed hash literal", () => {
    // Guards against the implementation drifting to a different hashing
    // strategy (e.g. padding, a different digest) without the test above
    // (which re-derives its expectation via the same viem primitives) noticing.
    expect(hashServiceName(name)).toBe(knownHash);
  });

  it("builds a bidirectional map", () => {
    const catalog = buildServiceCatalog([name]);
    const hash = hashServiceName(name);

    expect(catalog.nameByHash.get(hash.toLowerCase())).toBe(name);
    expect(catalog.hashByName.get(name)).toBe(hash);
  });

  it("resolves a hash regardless of its casing", () => {
    const catalog = buildServiceCatalog([name]);
    const hash = hashServiceName(name);

    expect(serviceNameForHash(catalog, hash.toUpperCase().replace("0X", "0x"))).toBe(name);
  });

  it("returns undefined for an unknown hash", () => {
    const catalog = buildServiceCatalog([name]);

    expect(serviceNameForHash(catalog, hashServiceName("ttm.services.nope.v1.NopeService"))).toBeUndefined();
  });

  it("returns empty maps for an empty catalog", () => {
    const catalog = buildServiceCatalog([]);

    expect(catalog.nameByHash.size).toBe(0);
    expect(catalog.hashByName.size).toBe(0);
  });
});

describe("unresolvedHashes", () => {
  const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";
  const registeredHash = hashServiceName(name);
  // A hash that is *not* in the catalog: this is the deprecated-service case — a service
  // an account adopted while it was registered, later unregistered on the manager. The
  // catalog (seeded from currently-registered names only) can never contain it.
  const deprecatedHash = hashServiceName("ttm.services.nope.v1.NopeService");

  it("returns nothing when every hash is in the catalog", () => {
    const catalog = buildServiceCatalog([name]);

    expect(unresolvedHashes(catalog, [registeredHash])).toEqual([]);
  });

  it("returns exactly the hashes the catalog is missing", () => {
    const catalog = buildServiceCatalog([name]);

    expect(unresolvedHashes(catalog, [registeredHash, deprecatedHash])).toEqual([deprecatedHash]);
  });

  it("de-dupes case-insensitively, keeping the first-seen casing", () => {
    const catalog = buildServiceCatalog([]);
    const upper = deprecatedHash.toUpperCase().replace("0X", "0x");

    expect(unresolvedHashes(catalog, [deprecatedHash, upper])).toEqual([deprecatedHash]);
  });

  it("returns an empty list for an empty input", () => {
    const catalog = buildServiceCatalog([name]);

    expect(unresolvedHashes(catalog, [])).toEqual([]);
  });
});
