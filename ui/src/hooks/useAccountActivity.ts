import { useMemo } from "react";
import { type Address } from "viem";
import { type ActivitySourceInput, useActivity } from "./useActivity";
import { useBlockTimestamps } from "./useBlockTimestamps";
import { useActiveContracts } from "./useActiveContracts";
import { useResolvedServiceNames } from "./useServiceCatalog";
import { ACCOUNT_EVENTS, renderSentence } from "../lib/activity/catalog";

/**
 * All activity emitted by a single TTM Account proxy (bots, services, tokens,
 * pubkeys, funds, config). One address, so a single getLogs filter covers it.
 *
 * Every account-side service event (`ServiceAdded`, `WantedServiceAdded`, etc.)
 * carries an indexed `bytes32 serviceHash`. We resolve those hashes to names
 * via `useResolvedServiceNames` (the registry's service catalog, plus a
 * bounded fallback for a service unregistered after this account adopted it)
 * and re-render the affected sentences.
 */
export function useAccountActivity(account: Address) {
  const { chainId } = useActiveContracts();

  const sources = useMemo<ActivitySourceInput[]>(
    () => [{ source: "account", address: account, events: ACCOUNT_EVENTS }],
    [account],
  );

  const activity = useActivity({ sources, chainId });

  const serviceHashes = useMemo(
    () => activity.events.map((e) => e.args.serviceHash).filter((h): h is string => typeof h === "string"),
    [activity.events],
  );

  const { resolve } = useResolvedServiceNames(serviceHashes);

  const timestamps = useBlockTimestamps(
    chainId,
    activity.events.map((e) => e.blockNumber),
  );

  const events = useMemo(
    () =>
      activity.events.map((e) => {
        const hash = e.args.serviceHash;
        const serviceLabel = typeof hash === "string" ? resolve(hash) : undefined;
        const sentence = serviceLabel ? renderSentence(e.source, e.eventName, { ...e.args, serviceLabel }) : e.sentence;
        return { ...e, timestamp: timestamps.get(e.blockNumber), sentence };
      }),
    [activity.events, timestamps, resolve],
  );

  return { ...activity, events };
}
