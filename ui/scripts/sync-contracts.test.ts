import { describe, expect, it } from "vitest";
import { resolveAddresses } from "./sync-contracts";

describe("resolveAddresses", () => {
  it("resolves canonical proxy addresses per chain", () => {
    const out = resolveAddresses({
      84532: {
        "TravelTokenMessengerModule#ManagerProxy": "0xMANAGER",
        "TravelTokenMessengerModule#BookingTokenProxy": "0xBT",
        "TravelTokenMessengerModule#TTMAccount": "0xIMPL",
        "TravelTokenMessengerModule#ManagerERC1967Proxy": "0xMANAGER",
      },
    });
    expect(out[84532]).toEqual({
      manager: "0xMANAGER",
      bookingToken: "0xBT",
      ttmAccountImpl: "0xIMPL",
    });
  });

  it("ignores non-canonical historical modules", () => {
    const out = resolveAddresses({
      501: {
        "TravelTokenMessengerModule#ManagerProxy": "0xGOOD",
        "TravelTokenMessengerModule#BookingTokenProxy": "0xGOODBT",
        "TravelTokenMessengerModule#TTMAccount": "0xGOODIMPL",
        "RefactorCancellationModule#ManagerProxy": "0xBAD",
        "ERC20ServiceFeeModule#TTMAccount": "0xBAD2",
      },
    });
    expect(out[501].manager).toBe("0xGOOD");
    expect(out[501].ttmAccountImpl).toBe("0xGOODIMPL");
  });

  it("omits chains missing required keys", () => {
    const out = resolveAddresses({ 8453: { "Other#Thing": "0x0" } });
    expect(out[8453]).toBeUndefined();
  });
});
