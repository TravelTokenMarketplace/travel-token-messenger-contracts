import { describe, expect, it } from "vitest";
import { encodeEventTopics } from "viem";
import { findCreatedAccount } from "./receipt";
import { MANAGER_ABI, TTMACCOUNT_ABI } from "../contracts";

describe("findCreatedAccount", () => {
  it("extracts the account, creator, and admin from logs", () => {
    const account = "0x1111111111111111111111111111111111111111";
    const creator = "0x2222222222222222222222222222222222222222";
    const admin = "0x3333333333333333333333333333333333333333";
    // TTMAccountCreated has only indexed args, so all values live in topics; data is empty.
    const topics = encodeEventTopics({
      abi: MANAGER_ABI,
      eventName: "TTMAccountCreated",
      args: { account, creator, admin },
    });
    const found = findCreatedAccount([{ data: "0x", topics }] as never, MANAGER_ABI as never);
    expect(found?.account.toLowerCase()).toBe(account);
    expect(found?.creator.toLowerCase()).toBe(creator);
    expect(found?.admin.toLowerCase()).toBe(admin);
  });

  it("returns undefined when no matching event", () => {
    expect(findCreatedAccount([{ data: "0x", topics: ["0xdead"] }] as never, MANAGER_ABI as never)).toBeUndefined();
  });

  it("returns undefined when decoded against the TTMAccount ABI, which never declares this event", () => {
    // TTMAccountCreated is declared only on the manager. Passing the account ABI (a
    // past bug in CreateAccount.tsx) makes decodeEventLog throw for every log, so the
    // production caller must pass MANAGER_ABI, not TTMACCOUNT_ABI, here.
    const account = "0x1111111111111111111111111111111111111111";
    const creator = "0x2222222222222222222222222222222222222222";
    const admin = "0x3333333333333333333333333333333333333333";
    const topics = encodeEventTopics({
      abi: MANAGER_ABI,
      eventName: "TTMAccountCreated",
      args: { account, creator, admin },
    });
    expect(findCreatedAccount([{ data: "0x", topics }] as never, TTMACCOUNT_ABI as never)).toBeUndefined();
  });
});
