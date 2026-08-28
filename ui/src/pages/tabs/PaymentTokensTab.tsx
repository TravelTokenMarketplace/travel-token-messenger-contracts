import { type Abi, type Address } from "viem";
import { ListManager } from "../../components/ListManager";
import { TokenDisplay } from "../../components/TokenDisplay";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useChainPinnedWrite } from "../../hooks/useChainPinnedWrite";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";
import { useTokenMetadata } from "../../hooks/useTokenMetadata";
import { NATIVE_SENTINEL, OFFCHAIN_SENTINEL } from "../../lib/paymentTokens";

export function PaymentTokensTab({ account }: { account: Address }) {
  const { ttmAccountAbi } = useActiveContracts();
  const abi = ttmAccountAbi as Abi;
  const { writeContractAsync } = useChainPinnedWrite();
  const { items, isLoading, refetch } = useContractList(account, abi, "getSupportedTokens");
  const { hasRole, isLoading: roleLoading } = useHasRole(account, abi, "SERVICE_ADMIN_ROLE");
  const { meta } = useTokenMetadata(items as Address[]);

  return (
    <ListManager
      title="Payment Tokens"
      items={items}
      isLoading={isLoading || roleLoading}
      roleName="SERVICE_ADMIN_ROLE"
      hasRole={hasRole}
      addLabel="Add token"
      addPlaceholder="Token address 0x…"
      presets={[
        { label: "Native currency", value: NATIVE_SENTINEL },
        { label: "Off-chain payment", value: OFFCHAIN_SENTINEL },
      ]}
      onAdd={(v) =>
        writeContractAsync({ address: account, abi, functionName: "addSupportedToken", args: [v as Address] })
      }
      onRemove={(v) =>
        writeContractAsync({ address: account, abi, functionName: "removeSupportedToken", args: [v as Address] })
      }
      onChanged={refetch}
      renderItem={(v) => {
        const m = meta.get(v.toLowerCase());
        return <TokenDisplay address={v} symbol={m?.symbol} name={m?.name} />;
      }}
    />
  );
}
