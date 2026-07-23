import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import type { Address } from "viem";

let mockMulticall: { data: unknown; isLoading: boolean };

// The spy MUST be `mock`-prefixed — vi.mock factories may only reference
// top-level variables whose names start with `mock`.
const mockReadSpy = vi.fn((_arg: unknown) => mockMulticall);
vi.mock("wagmi", () => ({ useReadContracts: (arg: unknown) => mockReadSpy(arg) }));
vi.mock("./useActiveContracts", () => ({ useActiveContracts: () => ({ chainId: 84532 }) }));

import { useTokenMetadata } from "./useTokenMetadata";
import { NATIVE_SENTINEL, OFFCHAIN_SENTINEL } from "../lib/paymentTokens";

const A = "0xAAaA000000000000000000000000000000000001" as Address;
// 3 reads per token: symbol, name, decimals
const ok = (s: string, n: string, d: number) => [
  { status: "success", result: s },
  { status: "success", result: n },
  { status: "success", result: d },
];

afterEach(() => {
  mockMulticall = { data: undefined, isLoading: false };
});

describe("useTokenMetadata", () => {
  it("maps symbol/name/decimals keyed by lowercase address", () => {
    mockMulticall = { data: ok("EURe", "Monerium EUR emoney", 18), isLoading: false };
    const { result } = renderHook(() => useTokenMetadata([A]));
    expect(result.current.meta.get(A.toLowerCase())).toMatchObject({
      symbol: "EURe",
      name: "Monerium EUR emoney",
      decimals: 18,
    });
  });

  it("dedupes addresses case-insensitively", () => {
    mockMulticall = { data: ok("EURe", "Monerium", 18), isLoading: false };
    const { result } = renderHook(() => useTokenMetadata([A, A.toLowerCase() as Address]));
    expect(result.current.meta.size).toBe(1);
  });

  it("leaves a field undefined when its read fails", () => {
    mockMulticall = {
      data: [
        { status: "failure", error: new Error("x") }, // symbol
        { status: "success", result: "Token" }, // name
        { status: "success", result: 6 }, // decimals
      ],
      isLoading: false,
    };
    const { result } = renderHook(() => useTokenMetadata([A]));
    const m = result.current.meta.get(A.toLowerCase());
    expect(m?.symbol).toBeUndefined();
    expect(m?.name).toBe("Token");
    expect(m?.decimals).toBe(6);
  });

  it("labels sentinels without reading them on-chain", () => {
    mockMulticall = { data: ok("EURe", "Monerium", 18), isLoading: false }; // 1 token worth of reads
    const { result } = renderHook(() =>
      useTokenMetadata([NATIVE_SENTINEL as Address, OFFCHAIN_SENTINEL as Address, A]),
    );
    // Only the ERC-20 (A) is read: 1 token * 3 calls.
    const passed = mockReadSpy.mock.calls.at(-1)?.[0] as { contracts: unknown[] };
    expect(passed.contracts).toHaveLength(3);
    expect(result.current.meta.get(NATIVE_SENTINEL)?.symbol).toBe("Native currency");
    expect(result.current.meta.get(OFFCHAIN_SENTINEL)?.symbol).toBe("Off-chain payment");
    expect(result.current.meta.get(A.toLowerCase())?.symbol).toBe("EURe");
  });
});
