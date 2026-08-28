import { describe, expect, it, vi } from "vitest";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined, isConnected: false, chainId: 84532 }),
  useBalance: () => ({ data: undefined }),
  useReadContracts: () => ({ data: undefined, isLoading: false }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));
vi.mock("../../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({
    chainId: 84532,
    supported: true,
    manager: "0x2222222222222222222222222222222222222222",
    ttmAccount: "0x3333333333333333333333333333333333333333",
    ttmAccountImpl: "0x4444444444444444444444444444444444444444",
    managerAbi: [],
    ttmAccountAbi: [],
    bookingTokenAbi: [],
  }),
}));
vi.mock("../../hooks/useRoleMembers", () => ({
  useRoleMembers: () => ({ members: [], isLoading: false, refetch: vi.fn() }),
}));
vi.mock("../../hooks/useHasRole", () => ({
  useHasRole: () => ({ hasRole: false, isLoading: false }),
}));

import { BOT_ROLES } from "./BotsTab";

describe("BOT_ROLES", () => {
  it("lists exactly the two roles addMessengerBot grants (Decision 5)", () => {
    expect(BOT_ROLES).toEqual(["MESSENGER_BOT_ROLE", "BOOKING_OPERATOR_ROLE"]);
  });

  it("no longer treats GAS_WITHDRAWER_ROLE as a bot role", () => {
    expect(BOT_ROLES).not.toContain("GAS_WITHDRAWER_ROLE");
  });
});
