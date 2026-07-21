import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BookingTokenTab } from "./BookingTokenTab";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: ({ functionName }: { functionName: string }) => {
    const map: Record<string, unknown> = {
      name: "BookingToken",
      symbol: "BToken",
      version: [1n, 0n, 0n],
      getManagerAddress: "0x2222222222222222222222222222222222222222",
      getMinExpirationTimestampDiff: 60n,
    };
    return { data: map[functionName], isLoading: false, refetch: vi.fn() };
  },
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));
// The real generated/addresses.ts is populated by `yarn sync` from on-chain
// deployment journals and is empty until a deployment exists under the
// renamed Ignition module id, so this test must not depend on it being
// present to exercise the "supported" path.
vi.mock("../../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({
    chainId: 84532,
    supported: true,
    manager: "0x2222222222222222222222222222222222222222",
    bookingToken: "0x3333333333333333333333333333333333333333",
    ttmAccountImpl: "0x4444444444444444444444444444444444444444",
    managerAbi: [],
    ttmAccountAbi: [],
    bookingTokenAbi: [],
  }),
}));

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("BookingTokenTab", () => {
  it("shows token info and settings", () => {
    wrap(<BookingTokenTab />);
    expect(screen.getByText("BookingToken")).toBeInTheDocument();
    expect(screen.getByText("BToken")).toBeInTheDocument();
    expect(screen.getAllByText(/min expiration/i).length).toBeGreaterThan(0);
  });
});
