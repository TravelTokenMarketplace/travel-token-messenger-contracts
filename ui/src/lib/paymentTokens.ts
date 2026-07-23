/**
 * Payment mode is encoded as a sentinel address in the on-chain allowlist:
 * address(0) = the network's native coin, address(1) = off-chain settlement,
 * anything else = a real ERC-20. This module is the single source of truth for
 * recognising and labelling the two sentinels in the UI. Wording mirrors the
 * CLI (`tasks/account.js`, `payment-token:list`).
 */
export const NATIVE_SENTINEL = "0x0000000000000000000000000000000000000000";
export const OFFCHAIN_SENTINEL = "0x0000000000000000000000000000000000000001";

export function isSentinel(address: string): boolean {
  const a = address.toLowerCase();
  return a === NATIVE_SENTINEL || a === OFFCHAIN_SENTINEL;
}

export function paymentTokenLabel(address: string): { symbol: string; name: string } | undefined {
  switch (address.toLowerCase()) {
    case NATIVE_SENTINEL:
      return { symbol: "Native currency", name: "On-chain, in the network's native coin" };
    case OFFCHAIN_SENTINEL:
      return { symbol: "Off-chain payment", name: "Settled off-chain, outside the contract" };
    default:
      return undefined;
  }
}
