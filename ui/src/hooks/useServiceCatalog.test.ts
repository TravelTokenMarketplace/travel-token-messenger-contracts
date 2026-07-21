import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import { keccak256, toBytes } from "viem";

let mockRegisteredNames: unknown;
let mockFallbackReads: { data: unknown; isLoading: boolean };
let lastFallbackContracts: unknown[] = [];
let lastFallbackEnabled = false;

vi.mock("wagmi", () => ({
  useReadContract: () => ({ data: mockRegisteredNames, isLoading: false }),
  useReadContracts: (args: { contracts: unknown[]; query?: { enabled?: boolean } }) => {
    lastFallbackContracts = args.contracts;
    lastFallbackEnabled = args.query?.enabled ?? true;
    return mockFallbackReads;
  },
}));
vi.mock("./useActiveContracts", () => ({
  useActiveContracts: () => ({
    manager: "0x2222222222222222222222222222222222222222",
    managerAbi: [],
    chainId: 84532,
  }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532 }),
}));

import { useServiceCatalog, useResolvedServiceNames } from "./useServiceCatalog";

const registeredName = "ttm.services.accommodation.v1alpha.AccommodationSearchService";
const registeredHash = keccak256(toBytes(registeredName));
// A hash for a service that is NOT in `getAllRegisteredServiceNames()` — this is the
// deprecated-service case: an account adopted it while it was registered, then it was
// unregistered on the manager. ServiceRegistry keeps the name resolvable on-chain via
// getServiceNameByHash; the catalog alone cannot see it.
const deprecatedName = "ttm.services.nope.v1.NopeService";
const deprecatedHash = keccak256(toBytes(deprecatedName));

afterEach(() => {
  mockRegisteredNames = [registeredName];
  mockFallbackReads = { data: undefined, isLoading: false };
  lastFallbackContracts = [];
  lastFallbackEnabled = false;
});

describe("useServiceCatalog", () => {
  it("builds the name/hash catalog from the registered service list", () => {
    mockRegisteredNames = [registeredName];
    const { result } = renderHook(() => useServiceCatalog());
    expect(result.current.catalog.nameByHash.get(registeredHash.toLowerCase())).toBe(registeredName);
  });
});

describe("useResolvedServiceNames", () => {
  it("resolves a currently-registered hash from the catalog alone, without a fallback call", () => {
    mockRegisteredNames = [registeredName];

    const { result } = renderHook(() => useResolvedServiceNames([registeredHash]));

    expect(result.current.resolve(registeredHash)).toBe(registeredName);
    // The common case: nothing missing, so the fallback batch is empty and disabled.
    expect(lastFallbackContracts).toEqual([]);
    expect(lastFallbackEnabled).toBe(false);
  });

  it("falls back to a per-hash lookup for a deprecated (unregistered) service and still resolves its name", () => {
    mockRegisteredNames = [registeredName];
    // The fallback batch is keyed by position in `missing`, which here is just [deprecatedHash].
    mockFallbackReads = { data: [{ status: "success", result: deprecatedName }], isLoading: false };

    const { result } = renderHook(() => useResolvedServiceNames([registeredHash, deprecatedHash]));

    // Still resolves the currently-registered one from the catalog...
    expect(result.current.resolve(registeredHash)).toBe(registeredName);
    // ...and resolves the deprecated one via the bounded fallback batch, which was
    // enabled with exactly the one hash the catalog couldn't cover.
    expect(result.current.resolve(deprecatedHash)).toBe(deprecatedName);
    expect(lastFallbackContracts).toHaveLength(1);
    expect(lastFallbackEnabled).toBe(true);
    // Pin the exact call: it must be the unconditional getServiceNameByHash, not the
    // sibling getRegisteredServiceNameByHash, which reverts with ServiceNotRegistered
    // for a service that was unregistered after an account adopted it — exactly the
    // deprecated-service case this fallback exists to handle. Swapping the two would
    // silently reinstate that regression while every other assertion here stayed green.
    expect(lastFallbackContracts[0]).toMatchObject({
      functionName: "getServiceNameByHash",
      args: [deprecatedHash],
    });
  });

  it("returns undefined for a hash neither the catalog nor the fallback batch resolves", () => {
    mockRegisteredNames = [registeredName];
    mockFallbackReads = { data: [{ status: "failure", error: new Error("not found") }], isLoading: false };

    const { result } = renderHook(() => useResolvedServiceNames([deprecatedHash]));

    expect(result.current.resolve(deprecatedHash)).toBeUndefined();
  });

  it("is loading while the fallback batch for a missing hash is still in flight", () => {
    mockRegisteredNames = [registeredName];
    mockFallbackReads = { data: undefined, isLoading: true };

    const { result } = renderHook(() => useResolvedServiceNames([deprecatedHash]));

    expect(result.current.isLoading).toBe(true);
  });

  it("is not loading on the fallback batch when there is nothing missing to resolve", () => {
    mockRegisteredNames = [registeredName];
    mockFallbackReads = { data: undefined, isLoading: true }; // would be loading if enabled

    const { result } = renderHook(() => useResolvedServiceNames([registeredHash]));

    expect(result.current.isLoading).toBe(false);
  });
});
