import { describe, expect, it } from "vitest";
import {
  ACCOUNT_EVENTS,
  BOOKING_TOKEN_EVENTS,
  MANAGER_EVENTS,
  lookupEntry,
  renderSentence,
  toActivityEvent,
} from "./catalog";

const ACC = "0xaAaA000000000000000000000000000000000001";
const BUYER = "0xbBbB000000000000000000000000000000000002";

function render(source: Parameters<typeof lookupEntry>[0], name: string, args: Record<string, unknown>) {
  const entry = lookupEntry(source, name);
  if (!entry) throw new Error(`no catalog entry for ${source}:${name}`);
  return entry.render(args);
}

describe("catalog rendering", () => {
  it("renders manager account + service events", () => {
    expect(render("manager", "TTMAccountCreated", { account: ACC })).toBe("TTM Account 0xaAaA…0001 created");
    expect(render("manager", "ServiceRegistered", { serviceName: "ttm.x.v1.Foo" })).toBe(
      'Service "ttm.x.v1.Foo" registered',
    );
  });

  it("renders booking + cancellation events with the token id", () => {
    expect(render("bookingToken", "TokenBought", { tokenId: 42n, buyer: BUYER })).toBe(
      "Booking token #42 bought by 0xbBbB…0002",
    );
    expect(render("bookingToken", "TokenReservationExpired", { tokenId: 7n })).toBe(
      "Booking token #7 reservation expired",
    );
    expect(render("bookingToken", "CancellationFinalized", { tokenId: 9n })).toBe(
      "Cancellation finalized for booking token #9",
    );
  });

  it("renders account fund events with ether-formatted amounts", () => {
    expect(render("account", "Deposit", { sender: ACC, amount: 1_500000000000000000n })).toBe(
      "Deposit of 1.5 from 0xaAaA…0001",
    );
  });

  it("falls back to a short service hash when the name isn't resolved", () => {
    // ServiceAdded carries an indexed bytes32 `serviceHash` (Task 5), decoded as-is by viem.
    expect(
      render("account", "ServiceAdded", {
        serviceHash: "0xabcd000000000000000000000000000000000000000000000000000000001234",
      }),
    ).toBe("Supported service (0xabcd…1234) added");
  });

  it("falls back to a short service hash for wanted-service events too", () => {
    // WantedServiceAdded/Removed switched from an indexed `string serviceName` to an
    // indexed `bytes32 serviceHash` in Task 6, mirroring ServiceAdded/Removed (Task 5).
    expect(
      render("account", "WantedServiceAdded", {
        serviceHash: "0xabcd000000000000000000000000000000000000000000000000000000001234",
      }),
    ).toBe("Wanted service (0xabcd…1234) added");
  });

  it("uses the resolved service name when injected via serviceLabel", () => {
    expect(renderSentence("account", "ServiceAdded", { serviceHash: "0xabcd…", serviceLabel: "ttm.x.v1.Foo" })).toBe(
      'Supported service "ttm.x.v1.Foo" added',
    );
    expect(
      renderSentence("account", "ServiceCapabilityAdded", {
        serviceHash: "0xabcd…",
        serviceLabel: "ttm.x.v1.Foo",
        capability: "luggage",
      }),
    ).toBe('Service "ttm.x.v1.Foo" capability "luggage" added');
  });

  it("renders detail for gas-money and upgrade config events", () => {
    expect(render("account", "GasMoneyWithdrawalUpdated", { limit: 2_000000000000000000n, period: 86400n })).toBe(
      "Gas money limit updated to 2 per 86400s",
    );
    expect(render("account", "TTMAccountUpgraded", { oldImplementation: ACC, newImplementation: BUYER })).toBe(
      "Account upgraded to implementation 0xbBbB…0002",
    );
  });

  it("exposes non-overlapping event sets per source", () => {
    expect(MANAGER_EVENTS.length).toBe(3);
    expect(BOOKING_TOKEN_EVENTS.length).toBe(7);
    expect(ACCOUNT_EVENTS.length).toBeGreaterThan(15);
  });
});

describe("toActivityEvent", () => {
  const log = {
    eventName: "TokenBought",
    args: { tokenId: 5n, buyer: BUYER },
    blockNumber: 100n,
    logIndex: 3,
    transactionHash: "0xabc",
    address: "0xcontract",
  };

  it("maps a decoded log to a normalized event", () => {
    const ev = toActivityEvent(log as never, "bookingToken");
    expect(ev).toMatchObject({
      id: "0xabc#3",
      source: "bookingToken",
      category: "Bookings",
      blockNumber: 100n,
      logIndex: 3,
      sentence: "Booking token #5 bought by 0xbBbB…0002",
    });
  });

  it("returns undefined for an unknown event", () => {
    expect(toActivityEvent({ ...log, eventName: "Nope" } as never, "bookingToken")).toBeUndefined();
  });
});
