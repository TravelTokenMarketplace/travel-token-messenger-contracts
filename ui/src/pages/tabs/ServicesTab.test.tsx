import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ServicesTab } from "./ServicesTab";

const HASH = "0xaaaa000000000000000000000000000000000000000000000000000000000001" as const;
const NAME = "ttm.services.accommodation.v1.AccommodationSearchService";

// The wallet is on Base while the app is showing Base Sepolia. Every write in
// the row must refuse in that state, including the ones that are chips rather
// than TxButtons.
let mockWalletChainId: number | undefined = 8453;

vi.mock("wagmi", () => ({
  useAccount: () => ({
    address: "0x1111111111111111111111111111111111111111",
    isConnected: true,
    chainId: mockWalletChainId,
  }),
  useReadContract: ({ functionName }: { functionName: string }) => ({
    data:
      functionName === "getAllServiceHashes"
        ? [HASH]
        : functionName === "getAllRegisteredServiceNames"
          ? [NAME]
          : undefined,
    isLoading: false,
    refetch: vi.fn(),
  }),
  useReadContracts: () => ({
    // restricted rate, then capabilities, in the order the row batches them.
    data: [{ result: true }, { result: ["cancel"] }],
    isLoading: false,
    refetch: vi.fn(),
  }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../hooks/useHasRole", () => ({ useHasRole: () => ({ hasRole: true, isLoading: false }) }));
vi.mock("../../hooks/useContractList", () => ({
  useContractList: () => ({ items: [], isLoading: false, refetch: vi.fn() }),
}));
vi.mock("../../hooks/useServiceCatalog", () => ({
  useServiceCatalog: () => ({
    catalog: { hashByName: new Map([[NAME, HASH]]), nameByHash: new Map([[HASH, NAME]]) },
    isLoading: false,
  }),
  useResolvedServiceNames: () => ({ resolve: () => NAME, isLoading: false }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

function openRow() {
  render(
    <QueryClientProvider client={new QueryClient()}>
      <ServicesTab account="0x2222222222222222222222222222222222222222" />
    </QueryClientProvider>,
  );
  fireEvent.click(screen.getByRole("button", { name: /AccommodationSearchService/i }));
}

describe("ServicesTab wrong-network guard", () => {
  it("refuses the inline service writes while the wallet is on another chain", () => {
    mockWalletChainId = 8453;
    openRow();
    expect(screen.getByRole("button", { name: /restricted rate: on/i })).toBeDisabled();
    expect(screen.getByRole("button", { name: /remove capability cancel/i })).toBeDisabled();
    expect(screen.getByRole("button", { name: /^add$/i })).toBeDisabled();
    // The tracked write in the same row is covered by TxButton's own guard.
    expect(screen.getByRole("button", { name: /remove service/i })).toBeDisabled();
  });

  it("allows them once the wallet is on the chain the app is showing", () => {
    mockWalletChainId = 84532;
    openRow();
    expect(screen.getByRole("button", { name: /restricted rate: on/i })).toBeEnabled();
    expect(screen.getByRole("button", { name: /remove capability cancel/i })).toBeEnabled();
    expect(screen.getByRole("button", { name: /remove service/i })).toBeEnabled();
  });
});
