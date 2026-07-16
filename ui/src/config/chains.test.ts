import { describe, expect, it } from "vitest";
import { APP_CHAINS, ENABLED_CHAINS } from "./chains";

describe("chain config", () => {
  it("defines the Base networks", () => {
    expect(APP_CHAINS.map((c) => c.id).sort()).toEqual([8453, 84532]);
  });

  it("only enables chains that have deployed addresses", () => {
    for (const c of ENABLED_CHAINS) {
      expect(c.enabled).toBe(true);
    }
  });
});
