import { afterEach, describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { TxButton } from "./TxButton";

let mockMismatch: { mismatched: boolean; reason?: string } = { mismatched: false };

vi.mock("../hooks/useChainMismatch", () => ({
  useChainMismatch: () => ({ activeChainId: 84532, ...mockMismatch }),
}));

afterEach(() => {
  mockMismatch = { mismatched: false };
});

describe("TxButton", () => {
  it("calls write and shows confirmed on success", async () => {
    const write = vi.fn().mockResolvedValue("0xhash");
    const onConfirmed = vi.fn();
    render(<TxButton label="Do it" write={write} onConfirmed={onConfirmed} />);
    fireEvent.click(screen.getByRole("button", { name: /do it/i }));
    await waitFor(() => expect(write).toHaveBeenCalled());
    await waitFor(() => expect(onConfirmed).toHaveBeenCalled());
  });

  it("shows an error message on failure", async () => {
    const write = vi.fn().mockRejectedValue(new Error("user rejected"));
    render(<TxButton label="Do it" write={write} />);
    fireEvent.click(screen.getByRole("button", { name: /do it/i }));
    await waitFor(() => expect(screen.getByText(/user rejected/i)).toBeInTheDocument());
  });

  it("refuses to submit while the wallet is on a different chain than the app", () => {
    mockMismatch = { mismatched: true, reason: "Your wallet is on Base but this app is showing Base Sepolia." };
    const write = vi.fn();
    render(<TxButton label="Do it" write={write} />);
    const button = screen.getByRole("button", { name: /do it/i });
    expect(button).toBeDisabled();
    fireEvent.click(button);
    expect(write).not.toHaveBeenCalled();
  });

  it("explains the mismatch instead of the usual tooltip", () => {
    mockMismatch = { mismatched: true, reason: "Your wallet is on Base but this app is showing Base Sepolia." };
    render(<TxButton label="Do it" write={vi.fn()} tooltip="Adds a bot." />);
    const button = screen.getByRole("button", { name: /do it/i });
    expect(button).toHaveAttribute("data-chain-mismatch", "true");
    expect(screen.queryByText("Adds a bot.")).not.toBeInTheDocument();
  });
});
