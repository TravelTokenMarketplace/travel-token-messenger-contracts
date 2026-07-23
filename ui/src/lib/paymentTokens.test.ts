import { describe, expect, it } from "vitest";
import { NATIVE_SENTINEL, OFFCHAIN_SENTINEL, isSentinel, paymentTokenLabel } from "./paymentTokens";

describe("paymentTokens", () => {
  it("recognises the two sentinels case-insensitively", () => {
    expect(isSentinel(NATIVE_SENTINEL)).toBe(true);
    expect(isSentinel(OFFCHAIN_SENTINEL.toUpperCase())).toBe(true);
    expect(isSentinel("0xAAaA000000000000000000000000000000000002")).toBe(false);
  });

  it("labels native and off-chain, mirroring the CLI wording", () => {
    expect(paymentTokenLabel(NATIVE_SENTINEL)?.symbol).toBe("Native currency");
    expect(paymentTokenLabel(OFFCHAIN_SENTINEL)?.symbol).toBe("Off-chain payment");
  });

  it("returns undefined for a real ERC-20 address", () => {
    expect(paymentTokenLabel("0xAAaA000000000000000000000000000000000002")).toBeUndefined();
  });
});
