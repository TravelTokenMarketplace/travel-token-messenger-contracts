import { afterEach, describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RolesPanel } from "./RolesPanel";

// Mutable mock state so individual tests can simulate a connected admin wallet.
// (vitest only allows factory-referenced vars prefixed with `mock`.)
let mockAddress: string | undefined;
let mockReadData: unknown;

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: mockAddress }),
  useReadContract: () => ({ data: mockReadData, isLoading: false, refetch: vi.fn() }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

const addr = "0x1111111111111111111111111111111111111111" as const;
const other = "0x2222222222222222222222222222222222222222" as const;

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

afterEach(() => {
  mockAddress = undefined;
  mockReadData = undefined;
});

describe("RolesPanel", () => {
  it("renders a row per role (enumerable)", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE", "PAUSER_ROLE"]} enumerable />);
    expect(screen.getByText("Admin")).toBeInTheDocument();
    expect(screen.getByText("Pauser")).toBeInTheDocument();
  });

  it("notes members are not listable in non-enumerable mode", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable={false} />);
    fireEvent.click(screen.getByRole("button", { name: /admin/i }));
    expect(screen.getByText(/cannot list/i)).toBeInTheDocument();
  });

  it("renders grant/revoke-by-address inputs for an admin in non-enumerable mode", () => {
    // Connected wallet that holds the admin role: useHasRole reads truthy.
    mockAddress = addr;
    mockReadData = true;
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable={false} />);
    fireEvent.click(screen.getByRole("button", { name: /admin/i }));
    expect(screen.getByPlaceholderText("Address 0x…")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Address 0x… to revoke")).toBeInTheDocument();
  });

  it("revokes without ceremony when the role has other members and is not your own", () => {
    mockAddress = addr;
    mockReadData = [other, "0x3333333333333333333333333333333333333333"];
    wrap(<RolesPanel address={addr} abi={[]} roles={["PAUSER_ROLE"]} enumerable />);
    fireEvent.click(screen.getByRole("button", { name: /pauser/i }));
    expect(screen.getAllByRole("button", { name: /revoke/i })[0]).toBeEnabled();
    expect(screen.queryByLabelText(/type revoke to confirm/i)).not.toBeInTheDocument();
  });

  it("makes revoking the only admin require a typed confirmation", () => {
    mockAddress = addr;
    mockReadData = [other];
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable />);
    fireEvent.click(screen.getByRole("button", { name: /admin/i }));
    expect(screen.getAllByText(/only admin/i).length).toBeGreaterThan(0);
    const revoke = screen.getAllByRole("button", { name: /revoke/i })[0];
    expect(revoke).toBeDisabled();
    fireEvent.change(screen.getByLabelText(/type revoke to confirm/i), { target: { value: "REVOKE" } });
    expect(screen.getAllByRole("button", { name: /revoke/i })[0]).toBeEnabled();
  });

  it("makes revoking your own role require a typed confirmation", () => {
    mockAddress = addr;
    mockReadData = [addr, other];
    wrap(<RolesPanel address={addr} abi={[]} roles={["PAUSER_ROLE"]} enumerable />);
    fireEvent.click(screen.getByRole("button", { name: /pauser/i }));
    expect(screen.getAllByText(/your own address/i).length).toBeGreaterThan(0);
    expect(screen.getAllByRole("button", { name: /revoke/i })[0]).toBeDisabled();
  });

  it("warns that the last admin cannot be checked when members are not enumerable", () => {
    mockAddress = addr;
    mockReadData = true;
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable={false} />);
    fireEvent.click(screen.getByRole("button", { name: /admin/i }));
    fireEvent.change(screen.getByPlaceholderText("Address 0x… to revoke"), { target: { value: other } });
    expect(screen.getAllByText(/cannot be checked/i).length).toBeGreaterThan(0);
    expect(screen.getByRole("button", { name: /revoke/i })).toBeDisabled();
  });
});
