import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import { useChainMismatch } from "./useChainMismatch";

let mockIsConnected = true;
let mockWalletChainId: number | undefined = 84532;
let mockActiveChainId = 84532;

vi.mock("wagmi", () => ({
  useAccount: () => ({ isConnected: mockIsConnected, chainId: mockWalletChainId }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: mockActiveChainId, setActiveChainId: vi.fn() }),
}));

afterEach(() => {
  mockIsConnected = true;
  mockWalletChainId = 84532;
  mockActiveChainId = 84532;
});

describe("useChainMismatch", () => {
  it("reports no mismatch when the wallet is on the chain the app is showing", () => {
    const { result } = renderHook(() => useChainMismatch());
    expect(result.current.mismatched).toBe(false);
    expect(result.current.reason).toBeUndefined();
  });

  it("reports no mismatch when no wallet is connected", () => {
    mockIsConnected = false;
    mockWalletChainId = undefined;
    expect(renderHook(() => useChainMismatch()).result.current.mismatched).toBe(false);
  });

  it("names both networks when the wallet is on a different chain", () => {
    mockWalletChainId = 8453;
    const { result } = renderHook(() => useChainMismatch());
    expect(result.current.mismatched).toBe(true);
    expect(result.current.reason).toContain("Base");
    expect(result.current.reason).toContain("Base Sepolia");
  });

  it("treats a connected wallet with an unknown chain as a mismatch", () => {
    mockWalletChainId = undefined;
    expect(renderHook(() => useChainMismatch()).result.current.mismatched).toBe(true);
  });
});
