import { useState } from "react";
import { ChevronRight, ShieldCheck, ShieldPlus, Trash2 } from "lucide-react";
import { type Abi, type Address, isAddress } from "viem";
import { useAccount } from "wagmi";
import { AddressDisplay } from "./AddressDisplay";
import { inputClass } from "./Input";
import { RoleGate } from "./RoleGate";
import { RowAction } from "./RowAction";
import { TxButton } from "./TxButton";
import { useChainPinnedWrite } from "../hooks/useChainPinnedWrite";
import { useHasRole } from "../hooks/useHasRole";
import { useRoleMembers } from "../hooks/useRoleMembers";
import { roleHash, type RoleName } from "../lib/roles";
import { shortRoleName } from "../lib/format";

function RoleHeader({
  label,
  role,
  open,
  badge,
}: {
  label: string;
  role: RoleName;
  open: boolean;
  badge: React.ReactNode;
}) {
  return (
    <>
      <ChevronRight className={`h-4 w-4 shrink-0 text-tarmac-400 transition-transform ${open ? "rotate-90" : ""}`} />
      <span className="flex-1 text-sm font-medium text-tarmac-800 dark:text-tarmac-100">{label}</span>
      <span className="font-mono text-[11px] text-tarmac-400">{role}</span>
      {badge}
    </>
  );
}

const REVOKE_CONFIRM_WORD = "REVOKE";

function sameAddress(a: string | undefined, b: string | undefined): boolean {
  return Boolean(a && b && a.toLowerCase() === b.toLowerCase());
}

/**
 * Revoke button that demands a typed confirmation for the revocations that
 * cannot be undone from the UI afterwards.
 *
 * Removing the last `DEFAULT_ADMIN_ROLE` holder strands the contract for good —
 * neither the account nor the manager overrides `revokeRole`/`renounceRole`, and
 * AccessControl itself is happy to remove the final admin. Revoking your own
 * role has the same shape from where the operator is sitting. Both used to be a
 * single mis-aimed click in a list of rows that all look alike.
 */
function RevokeAction({
  label,
  target,
  warning,
  disabled,
  write,
  onConfirmed,
}: {
  label: string;
  /** The address this button would revoke from; changing it clears the confirmation. */
  target: string;
  /** Why this particular revocation is dangerous; undefined for a routine one. */
  warning?: string;
  disabled?: boolean;
  write: () => Promise<`0x${string}`>;
  onConfirmed?: () => void;
}) {
  const [typed, setTyped] = useState("");
  // The non-enumerable row keeps one RevokeAction mounted while the operator
  // edits the address beside it, so the word typed for one address would still
  // count for whatever address replaced it — the ceremony would be performed
  // once and then stand for every revocation after it. Clearing during render
  // rather than in an effect means `confirmed` is never briefly true for an
  // address nobody confirmed.
  const [confirmedFor, setConfirmedFor] = useState(target);
  if (confirmedFor !== target) {
    setConfirmedFor(target);
    setTyped("");
  }
  const confirmed = !warning || typed.trim().toUpperCase() === REVOKE_CONFIRM_WORD;
  return (
    <div className="flex flex-col items-end gap-1">
      {warning && (
        <div className="flex items-center justify-end gap-2">
          <span className="max-w-xs text-right text-xs text-signal-fg dark:text-signal-dark">{warning}</span>
          <input
            className={`w-32 ${inputClass}`}
            placeholder={`Type ${REVOKE_CONFIRM_WORD}`}
            aria-label={`Type ${REVOKE_CONFIRM_WORD} to confirm`}
            value={typed}
            onChange={(e) => setTyped(e.target.value)}
          />
        </div>
      )}
      <TxButton
        label="Revoke"
        variant="danger"
        icon={<Trash2 className="h-4 w-4" />}
        disabled={disabled || !confirmed}
        tooltip={warning ?? `Revokes ${label} from this address — sends a transaction to your wallet.`}
        write={write}
        onConfirmed={onConfirmed}
      />
    </div>
  );
}

function GrantForm({
  account,
  abi,
  role,
  label,
  onDone,
}: {
  account: Address;
  abi: Abi;
  role: RoleName;
  label: string;
  onDone: () => void;
}) {
  const { writeContractAsync } = useChainPinnedWrite();
  const [grantee, setGrantee] = useState("");
  const trimmed = grantee.trim();
  const valid = isAddress(trimmed);
  return (
    <div className="flex items-end gap-2">
      <input
        className={`flex-1 ${inputClass}`}
        placeholder="Address 0x…"
        value={grantee}
        onChange={(e) => setGrantee(e.target.value)}
      />
      <TxButton
        label="Grant"
        icon={<ShieldPlus className="h-4 w-4" />}
        disabled={!valid}
        tooltip={`Grants ${label} to this address — sends a transaction to your wallet.`}
        write={() =>
          writeContractAsync({
            address: account,
            abi,
            functionName: "grantRole",
            args: [roleHash(role), trimmed as Address],
          })
        }
        onConfirmed={() => {
          setGrantee("");
          onDone();
        }}
      />
    </div>
  );
}

function EnumerableRoleRow({
  account,
  abi,
  role,
  hasAdmin,
  open,
  onToggle,
}: {
  account: Address;
  abi: Abi;
  role: RoleName;
  hasAdmin: boolean;
  open: boolean;
  onToggle: () => void;
}) {
  const { writeContractAsync } = useChainPinnedWrite();
  const { members, isLoading, refetch } = useRoleMembers(account, abi, role);
  const { address: connected } = useAccount();
  const label = shortRoleName(role);

  // The member list is already loaded here, so the last-admin check costs nothing extra.
  function revokeWarning(member: string): string | undefined {
    const warnings: string[] = [];
    if (role === "DEFAULT_ADMIN_ROLE" && members.length === 1) {
      warnings.push("This is the only admin — revoking it locks administration of this contract permanently.");
    }
    if (sameAddress(member, connected)) {
      warnings.push("This is your own address — you lose this role the moment it confirms.");
    }
    return warnings.length ? warnings.join(" ") : undefined;
  }

  return (
    <li className="rounded-md border border-tarmac-100 dark:border-tarmac-700/60">
      <button
        type="button"
        onClick={onToggle}
        aria-expanded={open}
        className="flex w-full items-center gap-2 px-3 py-2 text-left"
      >
        <RoleHeader
          label={label}
          role={role}
          open={open}
          badge={
            <span
              className={`rounded-full px-2 py-0.5 text-xs font-medium ${
                !isLoading && members.length > 0
                  ? "bg-brand-100 text-brand-700 dark:bg-brand-950 dark:text-brand-300"
                  : "bg-tarmac-100 text-tarmac-500 dark:bg-tarmac-700 dark:text-tarmac-400"
              }`}
              title={`${members.length} address${members.length === 1 ? "" : "es"} with this role`}
            >
              {isLoading ? "…" : members.length}
            </span>
          }
        />
      </button>
      {open && (
        <div className="space-y-3 border-t border-tarmac-100 px-3 py-3 dark:border-tarmac-700/60">
          <ul className="divide-y dark:divide-tarmac-700">
            {isLoading && <li className="py-2 text-sm text-tarmac-400">Loading…</li>}
            {!isLoading && members.length === 0 && <li className="py-2 text-sm text-tarmac-400">No members</li>}
            {members.map((m) => (
              <li key={m} className="group flex items-center justify-between gap-3 py-2">
                <AddressDisplay address={m} className="text-sm" />
                {hasAdmin && (
                  <RowAction>
                    <RevokeAction
                      label={label}
                      target={m}
                      warning={revokeWarning(m)}
                      write={() =>
                        writeContractAsync({
                          address: account,
                          abi,
                          functionName: "revokeRole",
                          args: [roleHash(role), m as Address],
                        })
                      }
                      onConfirmed={refetch}
                    />
                  </RowAction>
                )}
              </li>
            ))}
          </ul>
          <RoleGate hasRole={hasAdmin} roleName="DEFAULT_ADMIN_ROLE" action={`grant ${label}`}>
            <GrantForm account={account} abi={abi} role={role} label={label} onDone={refetch} />
          </RoleGate>
        </div>
      )}
    </li>
  );
}

function NonEnumerableRoleRow({
  account,
  abi,
  role,
  hasAdmin,
  open,
  onToggle,
}: {
  account: Address;
  abi: Abi;
  role: RoleName;
  hasAdmin: boolean;
  open: boolean;
  onToggle: () => void;
}) {
  const { writeContractAsync } = useChainPinnedWrite();
  const { hasRole: youHold } = useHasRole(account, abi, role);
  const { address: connected } = useAccount();
  const [revokee, setRevokee] = useState("");
  const label = shortRoleName(role);
  const trimmedRevokee = revokee.trim();

  // No member list to consult here, so the last-admin case cannot be ruled out —
  // say so rather than let an admin revocation through as though it were checked.
  function revokeWarning(): string | undefined {
    if (!isAddress(trimmedRevokee)) return undefined;
    const warnings: string[] = [];
    if (role === "DEFAULT_ADMIN_ROLE") {
      warnings.push(
        "This contract cannot list role members, so whether this is the last admin cannot be checked here. " +
          "Revoking the last admin locks administration permanently.",
      );
    }
    if (sameAddress(trimmedRevokee, connected)) {
      warnings.push("This is your own address — you lose this role the moment it confirms.");
    }
    return warnings.length ? warnings.join(" ") : undefined;
  }

  return (
    <li className="rounded-md border border-tarmac-100 dark:border-tarmac-700/60">
      <button
        type="button"
        onClick={onToggle}
        aria-expanded={open}
        className="flex w-full items-center gap-2 px-3 py-2 text-left"
      >
        <RoleHeader
          label={label}
          role={role}
          open={open}
          badge={
            youHold ? (
              <span className="inline-flex items-center gap-1 rounded-full bg-brand-100 px-2 py-0.5 text-xs font-medium text-brand-700 dark:bg-brand-950 dark:text-brand-300">
                <ShieldCheck className="h-3 w-3" /> You
              </span>
            ) : (
              <span className="w-px" />
            )
          }
        />
      </button>
      {open && (
        <div className="space-y-3 border-t border-tarmac-100 px-3 py-3 dark:border-tarmac-700/60">
          <p className="text-xs text-tarmac-400">
            This contract cannot list role members on-chain. Grant or revoke by address.
          </p>
          <RoleGate hasRole={hasAdmin} roleName="DEFAULT_ADMIN_ROLE" action={`manage ${label}`}>
            <div className="space-y-2">
              <GrantForm
                account={account}
                abi={abi}
                role={role}
                label={label}
                onDone={() => {
                  /* no member list to refetch in non-enumerable mode */
                }}
              />
              <div className="flex items-end gap-2">
                <input
                  className={`flex-1 ${inputClass}`}
                  placeholder="Address 0x… to revoke"
                  value={revokee}
                  onChange={(e) => setRevokee(e.target.value)}
                />
                <RevokeAction
                  label={label}
                  target={trimmedRevokee}
                  warning={revokeWarning()}
                  disabled={!isAddress(trimmedRevokee)}
                  write={() =>
                    writeContractAsync({
                      address: account,
                      abi,
                      functionName: "revokeRole",
                      args: [roleHash(role), trimmedRevokee as Address],
                    })
                  }
                  onConfirmed={() => setRevokee("")}
                />
              </div>
            </div>
          </RoleGate>
        </div>
      )}
    </li>
  );
}

export function RolesPanel({
  address,
  abi,
  roles,
  enumerable,
}: {
  address: Address;
  abi: Abi;
  roles: readonly RoleName[];
  enumerable: boolean;
}) {
  const { hasRole: hasAdmin } = useHasRole(address, abi, "DEFAULT_ADMIN_ROLE");
  const [openRole, setOpenRole] = useState<RoleName | null>(null);
  const Row = enumerable ? EnumerableRoleRow : NonEnumerableRoleRow;

  return (
    <ul className="space-y-2">
      {roles.map((r) => (
        <Row
          key={r}
          account={address}
          abi={abi}
          role={r}
          hasAdmin={hasAdmin}
          open={openRole === r}
          onToggle={() => setOpenRole((cur) => (cur === r ? null : r))}
        />
      ))}
    </ul>
  );
}
