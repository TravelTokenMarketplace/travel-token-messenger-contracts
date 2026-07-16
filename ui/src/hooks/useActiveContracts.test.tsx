import { describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import { useActiveContracts } from "./useActiveContracts";

vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));

// The real generated/addresses.ts is populated by `yarn sync` reading on-chain
// deployment journals and is empty until a deployment exists under the
// renamed Ignition module id, so tests must not depend on it being present.
vi.mock("../contracts", () => ({
  hasContracts: (chainId: number) => chainId === 84532,
  getContractsForChain: (chainId: number) =>
    chainId === 84532
      ? {
          manager: "0x2222222222222222222222222222222222222222",
          bookingToken: "0x3333333333333333333333333333333333333333",
          ttmAccountImpl: "0x4444444444444444444444444444444444444444",
        }
      : undefined,
  MANAGER_ABI: [],
  TTMACCOUNT_ABI: [],
  BOOKINGTOKEN_ABI: [],
}));

describe("useActiveContracts", () => {
  it("reports supported=true for a chain with contracts", () => {
    const { result } = renderHook(() => useActiveContracts());
    expect(result.current.chainId).toBe(84532);
    expect(result.current.supported).toBe(true);
    expect(result.current.manager).toBeTruthy();
  });
});
