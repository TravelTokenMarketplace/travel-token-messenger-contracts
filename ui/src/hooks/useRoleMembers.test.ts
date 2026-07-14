import { describe, expect, it, vi } from "vitest";
import { toMemberList } from "./useRoleMembers";

// useRoleMembers.ts imports wallet/activeChain, which throws at module load
// if ENABLED_CHAINS is empty. The real generated/addresses.ts is populated by
// `yarn sync` from on-chain deployment journals and is empty until a
// deployment exists under the renamed Ignition module id, so this pure-helper
// test must not depend on it being present.
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

describe("toMemberList", () => {
  it("returns [] for undefined", () => {
    expect(toMemberList(undefined)).toEqual([]);
  });

  it("stringifies address array entries", () => {
    expect(toMemberList(["0xAAA", "0xBBB"])).toEqual(["0xAAA", "0xBBB"]);
  });
});
