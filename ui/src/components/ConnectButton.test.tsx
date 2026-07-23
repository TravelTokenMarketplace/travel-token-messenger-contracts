import { afterEach, describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen } from "@testing-library/react";

let mockConnectors: { id: string; name: string; uid: string }[];
const connect = vi.fn();

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined, isConnected: false, chainId: 84532 }),
  useConnect: () => ({ connect, connectors: mockConnectors }),
  useDisconnect: () => ({ disconnect: vi.fn() }),
}));

import { ConnectButton } from "./ConnectButton";

afterEach(() => {
  connect.mockClear();
});

describe("ConnectButton (disconnected)", () => {
  it("shows a picker for multiple connectors and connects the chosen one", () => {
    mockConnectors = [
      { id: "injected", name: "Injected", uid: "a" },
      { id: "walletConnect", name: "WalletConnect", uid: "b" },
    ];
    render(<ConnectButton />);
    fireEvent.click(screen.getByRole("button", { name: /connect wallet/i }));
    fireEvent.click(screen.getByText(/browser wallet/i));
    expect(connect).toHaveBeenCalledWith({ connector: mockConnectors[0] });
  });

  it("keeps a single connector as a one-click button", () => {
    mockConnectors = [{ id: "injected", name: "Injected", uid: "a" }];
    render(<ConnectButton />);
    fireEvent.click(screen.getByRole("button", { name: /connect wallet/i }));
    expect(connect).toHaveBeenCalledWith({ connector: mockConnectors[0] });
  });
});
