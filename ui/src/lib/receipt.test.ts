import { describe, expect, it } from "vitest";
import { encodeEventTopics, parseAbi } from "viem";
import { findCreatedAccount } from "./receipt";

const abi = parseAbi([
  "event TTMAccountCreated(address indexed account, address indexed creator, address indexed admin)",
]);

describe("findCreatedAccount", () => {
  it("extracts the account, creator, and admin from logs", () => {
    const account = "0x1111111111111111111111111111111111111111";
    const creator = "0x2222222222222222222222222222222222222222";
    const admin = "0x3333333333333333333333333333333333333333";
    // TTMAccountCreated has only indexed args, so all values live in topics; data is empty.
    const topics = encodeEventTopics({ abi, eventName: "TTMAccountCreated", args: { account, creator, admin } });
    const found = findCreatedAccount([{ data: "0x", topics }] as never, abi as never);
    expect(found?.account.toLowerCase()).toBe(account);
    expect(found?.creator.toLowerCase()).toBe(creator);
    expect(found?.admin.toLowerCase()).toBe(admin);
  });

  it("returns undefined when no matching event", () => {
    expect(findCreatedAccount([{ data: "0x", topics: ["0xdead"] }] as never, abi as never)).toBeUndefined();
  });
});
