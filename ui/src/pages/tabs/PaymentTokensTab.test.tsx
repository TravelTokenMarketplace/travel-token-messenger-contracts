import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { PaymentTokensTab } from "./PaymentTokensTab";
import { NATIVE_SENTINEL } from "../../lib/paymentTokens";

const writeContractAsync = vi.fn().mockResolvedValue("0xhash");

vi.mock("wagmi", () => ({ useWriteContract: () => ({ writeContractAsync }) }));
vi.mock("../../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({ ttmAccountAbi: [] }),
}));
vi.mock("../../hooks/useContractList", () => ({
  useContractList: () => ({ items: [], isLoading: false, refetch: vi.fn() }),
}));
vi.mock("../../hooks/useHasRole", () => ({
  useHasRole: () => ({ hasRole: true, isLoading: false }),
}));
vi.mock("../../hooks/useTokenMetadata", () => ({
  useTokenMetadata: () => ({ meta: new Map(), isLoading: false }),
}));

const account = "0x1111111111111111111111111111111111111111" as const;

describe("PaymentTokensTab", () => {
  it("offers native and off-chain as one-click presets", async () => {
    render(<PaymentTokensTab account={account} />);
    expect(screen.getByRole("button", { name: /native currency/i })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /off-chain payment/i })).toBeInTheDocument();
    fireEvent.click(screen.getByRole("button", { name: /native currency/i }));
    await waitFor(() =>
      expect(writeContractAsync).toHaveBeenCalledWith(
        expect.objectContaining({ functionName: "addSupportedToken", args: [NATIVE_SENTINEL] }),
      ),
    );
  });
});
