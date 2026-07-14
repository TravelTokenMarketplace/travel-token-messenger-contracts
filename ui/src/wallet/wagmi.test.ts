import { describe, expect, it, vi } from "vitest";
import { wagmiConfig } from "./wagmi";
import { ENABLED_CHAINS } from "../config/chains";

// The real generated/addresses.ts (read by config/chains.ts via hasContracts)
// is populated by `yarn sync` from on-chain deployment journals and is empty
// until a deployment exists under the renamed Ignition module id, so tests
// must not depend on it being present.
vi.mock("../contracts", () => ({
  hasContracts: () => true,
  getContractsForChain: () => ({
    manager: "0x2222222222222222222222222222222222222222",
    bookingToken: "0x3333333333333333333333333333333333333333",
    ttmAccountImpl: "0x4444444444444444444444444444444444444444",
  }),
  MANAGER_ABI: [],
  TTMACCOUNT_ABI: [],
  BOOKINGTOKEN_ABI: [],
}));

describe("wagmiConfig", () => {
  it("registers a transport for every enabled chain", () => {
    for (const c of ENABLED_CHAINS) {
      expect(wagmiConfig.chains.some((wc) => wc.id === c.id)).toBe(true);
    }
  });
});
