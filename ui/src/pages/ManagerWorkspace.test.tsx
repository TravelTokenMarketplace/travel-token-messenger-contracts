import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ManagerWorkspace } from "./ManagerWorkspace";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useBalance: () => ({ data: undefined }),
  useReadContract: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useReadContracts: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn(), txs: [], dismiss: vi.fn() }) }));
// The real generated/addresses.ts is populated by `yarn sync` from on-chain
// deployment journals and is empty until a deployment exists under the
// renamed Ignition module id, so this test must not depend on it being
// present to exercise the "supported" path.
vi.mock("../hooks/useActiveContracts", () => ({
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

describe("ManagerWorkspace", () => {
  it("renders the manager tab bar", () => {
    render(
      <QueryClientProvider client={new QueryClient()}>
        <MemoryRouter initialEntries={["/manager"]}>
          <Routes>
            <Route path="manager" element={<ManagerWorkspace />} />
          </Routes>
        </MemoryRouter>
      </QueryClientProvider>,
    );
    expect(screen.getByRole("link", { name: /service registry/i })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: /booking token/i })).toBeInTheDocument();
  });
});
