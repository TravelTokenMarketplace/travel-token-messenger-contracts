import { decodeEventLog, type Abi, type Address, type Log } from "viem";

/** The account, its creator, and its admin, as carried by `TTMAccountCreated`. */
export interface CreatedAccount {
  account: Address;
  creator: Address;
  admin: Address;
}

export function findCreatedAccount(logs: Pick<Log, "data" | "topics">[], abi: Abi): CreatedAccount | undefined {
  for (const log of logs) {
    try {
      const decoded = decodeEventLog({ abi, data: log.data, topics: log.topics });
      if (decoded.eventName === "TTMAccountCreated") {
        return decoded.args as unknown as CreatedAccount;
      }
    } catch {
      // not our event, skip
    }
  }
  return undefined;
}
