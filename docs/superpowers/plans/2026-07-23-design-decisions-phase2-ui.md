# Phase 2 (UI) — Contract Design Decisions — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Bring the contracts management UI in line with phase-1's on-chain changes: multi-connector wallet support (custom picker), payment-token sentinel labels, and a corrected Bots-tab role list.

**Architecture:** UI-only changes in `travel-token-messenger-contracts/ui`. No contract, ABI, docgen, or Go-binding change. Three independent slices — payment-token sentinels (Section B), Bots-tab role list (Section C), wallet connectors (Section A) — each landing behind the existing test suite plus new tests.

**Tech Stack:** React + TypeScript, Vite, wagmi v2 + viem, Headless UI, Tailwind (transit-board tokens `brand`/`tarmac`/`paper`), Vitest + Testing Library.

## Global Constraints

- **UI-only.** Do not touch `contracts/`, `abi/`, `docs/` (docgen), or `go/` bindings. The contracts verification recipe is not run for this work.
- **Gate per task:** run the task's own test file. **Gate for the branch:** `cd ui && yarn sync && yarn test && yarn build` (tsc typecheck + bundle), plus repo-root `yarn lint`.
- **Test runner:** `yarn test <path>` (script is `vitest run`; a path argument filters to that file).
- **Sentinel wording mirrors the CLI** (`tasks/account.js:471`): `address(0)` → "Native currency", `address(1)` → "Off-chain payment".
- **WalletConnect id is optional:** read `import.meta.env.VITE_WALLETCONNECT_PROJECT_ID`; register the WalletConnect connector only when it is set. Never commit a real id.
- **Do not regress** the `ENABLED_CHAINS`-empty guard in `wagmi.ts`, and do not introduce RainbowKit usage (it stays an unused dep).
- Work on branch `feat/design-decisions-phase2-ui` (already created off `dev`). All paths below are relative to `travel-token-messenger-contracts/`.

---

## File structure

New:
- `ui/src/lib/paymentTokens.ts` — sentinel constants + label/predicate helpers (single source of truth).
- `ui/src/lib/paymentTokens.test.ts`
- `ui/src/pages/tabs/PaymentTokensTab.test.tsx`
- `ui/src/components/ConnectButton.test.tsx`
- `ui/.env.example`

Modified:
- `ui/src/hooks/useTokenMetadata.ts` (+ `useTokenMetadata.test.tsx`)
- `ui/src/components/ListManager.tsx` (+ `ListManager.test.tsx`)
- `ui/src/pages/tabs/PaymentTokensTab.tsx`
- `ui/src/pages/tabs/BotsTab.tsx` (+ new `BotsTab.roles.test.ts`)
- `ui/src/wallet/wagmi.ts` (+ `wagmi.test.ts`)
- `ui/src/components/ConnectButton.tsx`
- `ui/CLAUDE.md` (one line documenting the env var)

---

## Section B — Payment-token sentinels

### Task 1: `paymentTokens` lib

**Files:**
- Create: `ui/src/lib/paymentTokens.ts`
- Test: `ui/src/lib/paymentTokens.test.ts`

**Interfaces:**
- Produces:
  - `NATIVE_SENTINEL: "0x0000000000000000000000000000000000000000"`
  - `OFFCHAIN_SENTINEL: "0x0000000000000000000000000000000000000001"`
  - `isSentinel(address: string): boolean`
  - `paymentTokenLabel(address: string): { symbol: string; name: string } | undefined`

- [ ] **Step 1: Write the failing test**

```ts
// ui/src/lib/paymentTokens.test.ts
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
```

- [ ] **Step 2: Run test to verify it fails**

Run: `yarn test src/lib/paymentTokens.test.ts`
Expected: FAIL — cannot resolve `./paymentTokens`.

- [ ] **Step 3: Write minimal implementation**

```ts
// ui/src/lib/paymentTokens.ts
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
```

- [ ] **Step 4: Run test to verify it passes**

Run: `yarn test src/lib/paymentTokens.test.ts`
Expected: PASS (3 tests).

- [ ] **Step 5: Commit**

```bash
git add ui/src/lib/paymentTokens.ts ui/src/lib/paymentTokens.test.ts
git commit -m "feat(ui): payment-token sentinel constants and labels"
```

---

### Task 2: `useTokenMetadata` skips and labels sentinels

**Files:**
- Modify: `ui/src/hooks/useTokenMetadata.ts`
- Test: `ui/src/hooks/useTokenMetadata.test.tsx`

**Interfaces:**
- Consumes: `NATIVE_SENTINEL`, `isSentinel`, `paymentTokenLabel` from Task 1.
- Produces: unchanged public signature — `useTokenMetadata(addresses): { meta: Map<string, TokenMeta>; isLoading }`. Sentinels now appear in `meta` with their label and are excluded from the on-chain multicall.

- [ ] **Step 1: Write the failing test** (append to existing file; the existing tests stay)

The existing test mocks `useReadContracts` as a value. Change the mock to a spy so the test can assert which contracts were read, then add two cases.

Replace the mock line at the top of `ui/src/hooks/useTokenMetadata.test.tsx`:

```ts
// was: vi.mock("wagmi", () => ({ useReadContracts: () => mockMulticall }));
// The spy MUST be `mock`-prefixed — vi.mock factories may only reference
// top-level variables whose names start with `mock`.
const mockReadSpy = vi.fn(() => mockMulticall);
vi.mock("wagmi", () => ({ useReadContracts: (arg: unknown) => mockReadSpy(arg) }));
```

Add the import and two tests:

```ts
import { NATIVE_SENTINEL, OFFCHAIN_SENTINEL } from "../lib/paymentTokens";

it("labels sentinels without reading them on-chain", () => {
  mockMulticall = { data: ok("EURe", "Monerium", 18), isLoading: false }; // 1 token worth of reads
  const { result } = renderHook(() =>
    useTokenMetadata([NATIVE_SENTINEL as Address, OFFCHAIN_SENTINEL as Address, A]),
  );
  // Only the ERC-20 (A) is read: 1 token * 3 calls.
  const passed = mockReadSpy.mock.calls.at(-1)?.[0] as { contracts: unknown[] };
  expect(passed.contracts).toHaveLength(3);
  expect(result.current.meta.get(NATIVE_SENTINEL)?.symbol).toBe("Native currency");
  expect(result.current.meta.get(OFFCHAIN_SENTINEL)?.symbol).toBe("Off-chain payment");
  expect(result.current.meta.get(A.toLowerCase())?.symbol).toBe("EURe");
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `yarn test src/hooks/useTokenMetadata.test.tsx`
Expected: FAIL — sentinels absent from `meta`, and `contracts` length is 9 (all three read).

- [ ] **Step 3: Write minimal implementation**

In `ui/src/hooks/useTokenMetadata.ts`, add the import and split the deduped `list` into sentinels and readable addresses. Replace the block from `const { data, isLoading } = useReadContracts(...)` through the `meta` construction with:

```ts
import { isSentinel, paymentTokenLabel } from "../lib/paymentTokens";
// ...existing dedupe producing `list`...

const readList = list.filter((a) => !isSentinel(a));

const { data, isLoading } = useReadContracts({
  allowFailure: true,
  contracts: readList.flatMap((address) => [
    { chainId, address, abi: ERC20_ABI, functionName: "symbol" } as const,
    { chainId, address, abi: ERC20_ABI, functionName: "name" } as const,
    { chainId, address, abi: ERC20_ABI, functionName: "decimals" } as const,
  ]),
  query: { enabled: readList.length > 0 },
});

const meta = new Map<string, TokenMeta>();

// Sentinels never hit the chain; they carry static labels.
for (const address of list) {
  const label = paymentTokenLabel(address);
  if (label) meta.set(address.toLowerCase(), { address, symbol: label.symbol, name: label.name });
}

if (data) {
  readList.forEach((address, i) => {
    const s = data[i * 3];
    const n = data[i * 3 + 1];
    const d = data[i * 3 + 2];
    meta.set(address.toLowerCase(), {
      address,
      symbol: s?.status === "success" ? (s.result as string) : undefined,
      name: n?.status === "success" ? (n.result as string) : undefined,
      decimals: d?.status === "success" ? Number(d.result) : undefined,
    });
  });
}

return { meta, isLoading };
```

- [ ] **Step 4: Run test to verify it passes**

Run: `yarn test src/hooks/useTokenMetadata.test.tsx`
Expected: PASS (existing 3 + new 1).

- [ ] **Step 5: Commit**

```bash
git add ui/src/hooks/useTokenMetadata.ts ui/src/hooks/useTokenMetadata.test.tsx
git commit -m "feat(ui): label payment sentinels and skip their ERC-20 reads"
```

---

### Task 3: `ListManager` gains a `presets` prop

**Files:**
- Modify: `ui/src/components/ListManager.tsx`
- Test: `ui/src/components/ListManager.test.tsx`

**Interfaces:**
- Produces: new optional prop `presets?: { label: string; value: string }[]`. When set and `hasRole`, render one quick-add button per preset above the input; a preset whose `value` is already in `items` (case-insensitive) is not rendered. Clicking a preset calls `onAdd(value)` then `onChanged`. Omitting `presets` leaves behaviour identical.

- [ ] **Step 1: Write the failing test** (append two tests to `ListManager.test.tsx`)

```ts
it("renders preset quick-add buttons and calls onAdd with the preset value", async () => {
  const onAdd = vi.fn().mockResolvedValue("0xhash");
  render(
    <ListManager
      {...base}
      hasRole
      items={[]}
      onAdd={onAdd}
      presets={[{ label: "Native currency", value: "0x0000000000000000000000000000000000000000" }]}
    />,
  );
  fireEvent.click(screen.getByRole("button", { name: /native currency/i }));
  await waitFor(() =>
    expect(onAdd).toHaveBeenCalledWith("0x0000000000000000000000000000000000000000"),
  );
});

it("hides a preset already present in items", () => {
  render(
    <ListManager
      {...base}
      hasRole
      items={["0x0000000000000000000000000000000000000000"]}
      onAdd={vi.fn()}
      presets={[{ label: "Native currency", value: "0x0000000000000000000000000000000000000000" }]}
    />,
  );
  expect(screen.queryByRole("button", { name: /native currency/i })).not.toBeInTheDocument();
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `yarn test src/components/ListManager.test.tsx`
Expected: FAIL — `presets` unknown, no such button.

- [ ] **Step 3: Write minimal implementation**

Add `presets` to `ListManagerProps`:

```ts
  /** Optional one-click add buttons (e.g. native / off-chain sentinels). */
  presets?: { label: string; value: string }[];
```

Inside the `<RoleGate>`, immediately above the `<div className="flex items-end gap-2">` input row, add:

```tsx
{props.presets && props.presets.length > 0 && (
  <div className="mb-2 flex flex-wrap gap-2">
    {props.presets
      .filter((p) => !items.some((it) => it.toLowerCase() === p.value.toLowerCase()))
      .map((p) => (
        <TxButton
          key={p.value}
          label={p.label}
          icon={<Plus className="h-4 w-4" />}
          write={() => props.onAdd(p.value)}
          onConfirmed={() => props.onChanged?.()}
        />
      ))}
  </div>
)}
```

(`items` is already destructured at the top of the component; `Plus` and `TxButton` are already imported.)

- [ ] **Step 4: Run test to verify it passes**

Run: `yarn test src/components/ListManager.test.tsx`
Expected: PASS (existing 2 + new 2).

- [ ] **Step 5: Commit**

```bash
git add ui/src/components/ListManager.tsx ui/src/components/ListManager.test.tsx
git commit -m "feat(ui): ListManager preset quick-add buttons"
```

---

### Task 4: `PaymentTokensTab` wires sentinel presets

**Files:**
- Modify: `ui/src/pages/tabs/PaymentTokensTab.tsx`
- Test: `ui/src/pages/tabs/PaymentTokensTab.test.tsx` (create)

**Interfaces:**
- Consumes: `NATIVE_SENTINEL`, `OFFCHAIN_SENTINEL` (Task 1); `presets` prop (Task 3); sentinel labels surfaced through `useTokenMetadata` (Task 2).

- [ ] **Step 1: Write the failing test**

```tsx
// ui/src/pages/tabs/PaymentTokensTab.test.tsx
import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { PaymentTokensTab } from "./PaymentTokensTab";
import { NATIVE_SENTINEL } from "../../lib/paymentTokens";

const writeContractAsync = vi.fn().mockResolvedValue("0xhash");

vi.mock("wagmi", () => ({ useWriteContract: () => ({ writeContractAsync }) }));
vi.mock("../../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({ ttmAccountAbi: [] }),
}));
vi.mock("../../hooks/useContractList", () => ({
  useContractList: () => ({ items: [], isLoading: false, refetch: vi.fn() }),
}));
vi.mock("../../hooks/useHasRole", () => ({
  useHasRole: () => ({ hasRole: true, isLoading: false }),
}));
vi.mock("../../hooks/useTokenMetadata", () => ({
  useTokenMetadata: () => ({ meta: new Map(), isLoading: false }),
}));

const account = "0x1111111111111111111111111111111111111111" as const;

describe("PaymentTokensTab", () => {
  it("offers native and off-chain as one-click presets", async () => {
    render(<PaymentTokensTab account={account} />);
    expect(screen.getByRole("button", { name: /native currency/i })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /off-chain payment/i })).toBeInTheDocument();
    fireEvent.click(screen.getByRole("button", { name: /native currency/i }));
    await waitFor(() =>
      expect(writeContractAsync).toHaveBeenCalledWith(
        expect.objectContaining({ functionName: "addSupportedToken", args: [NATIVE_SENTINEL] }),
      ),
    );
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `yarn test src/pages/tabs/PaymentTokensTab.test.tsx`
Expected: FAIL — no preset buttons rendered.

- [ ] **Step 3: Write minimal implementation**

In `ui/src/pages/tabs/PaymentTokensTab.tsx`, add the import and pass `presets` to `<ListManager>`:

```tsx
import { NATIVE_SENTINEL, OFFCHAIN_SENTINEL } from "../../lib/paymentTokens";
```

Add this prop to the `<ListManager ...>` element (alongside `addLabel`/`addPlaceholder`):

```tsx
      presets={[
        { label: "Native currency", value: NATIVE_SENTINEL },
        { label: "Off-chain payment", value: OFFCHAIN_SENTINEL },
      ]}
```

No other change — the sentinel display in the list already works via Task 2's labelled `meta`.

- [ ] **Step 4: Run test to verify it passes**

Run: `yarn test src/pages/tabs/PaymentTokensTab.test.tsx`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add ui/src/pages/tabs/PaymentTokensTab.tsx ui/src/pages/tabs/PaymentTokensTab.test.tsx
git commit -m "feat(ui): add native/off-chain quick-add to PaymentTokensTab"
```

---

## Section C — Bots-tab role list

### Task 5: Drop `GAS_WITHDRAWER_ROLE` from the bot role list

**Files:**
- Modify: `ui/src/pages/tabs/BotsTab.tsx`
- Test: `ui/src/pages/tabs/BotsTab.roles.test.ts` (create)

**Interfaces:**
- Produces: `BOT_ROLES` becomes an **exported** const `["MESSENGER_BOT_ROLE", "BOOKING_OPERATOR_ROLE"]` so it can be asserted directly (the component render path is heavy to mount; the data is the regression risk).

- [ ] **Step 1: Write the failing test**

```ts
// ui/src/pages/tabs/BotsTab.roles.test.ts
import { describe, expect, it } from "vitest";
import { BOT_ROLES } from "./BotsTab";

describe("BOT_ROLES", () => {
  it("lists exactly the two roles addMessengerBot grants (Decision 5)", () => {
    expect(BOT_ROLES).toEqual(["MESSENGER_BOT_ROLE", "BOOKING_OPERATOR_ROLE"]);
  });

  it("no longer treats GAS_WITHDRAWER_ROLE as a bot role", () => {
    expect(BOT_ROLES).not.toContain("GAS_WITHDRAWER_ROLE");
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `yarn test src/pages/tabs/BotsTab.roles.test.ts`
Expected: FAIL — `BOT_ROLES` not exported / still contains three roles.

- [ ] **Step 3: Write minimal implementation**

In `ui/src/pages/tabs/BotsTab.tsx`, change the declaration at line 20-21:

```ts
// addMessengerBot grants exactly these two roles (Decision 5). GAS_WITHDRAWER_ROLE
// is opt-in and managed on the Roles tab, so it is not listed here.
export const BOT_ROLES: RoleName[] = ["MESSENGER_BOT_ROLE", "BOOKING_OPERATOR_ROLE"];
```

No other change: with gas-withdrawer gone from the list, its badge and the "may not function fully" tooltip no longer apply to it, and `removeMessengerBot` still defensively revokes it on-chain, so the "Remove" tooltip stays accurate.

- [ ] **Step 4: Run test to verify it passes**

Run: `yarn test src/pages/tabs/BotsTab.roles.test.ts`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add ui/src/pages/tabs/BotsTab.tsx ui/src/pages/tabs/BotsTab.roles.test.ts
git commit -m "fix(ui): stop listing GAS_WITHDRAWER_ROLE as a bot role (Decision 5)"
```

---

## Section A — Wallet connectors

### Task 6: Register Safe + WalletConnect connectors

**Files:**
- Modify: `ui/src/wallet/wagmi.ts`
- Modify: `ui/src/wallet/wagmi.test.ts`
- Create: `ui/.env.example`
- Modify: `ui/CLAUDE.md`

**Interfaces:**
- Produces: `wagmiConfig.connectors` includes `injected` and `safe` always, and `walletConnect` only when `VITE_WALLETCONNECT_PROJECT_ID` is set.

- [ ] **Step 1: Write the failing test** (append to `wagmi.test.ts`)

```ts
describe("connectors", () => {
  it("registers injected and safe, and omits walletConnect without a project id", () => {
    const ids = wagmiConfig.connectors.map((c) => c.id);
    expect(ids).toContain("injected");
    expect(ids).toContain("safe");
    // VITE_WALLETCONNECT_PROJECT_ID is unset in the test env.
    expect(ids).not.toContain("walletConnect");
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `yarn test src/wallet/wagmi.test.ts`
Expected: FAIL — `safe` connector absent.

- [ ] **Step 3: Write minimal implementation**

In `ui/src/wallet/wagmi.ts`, change the connectors import and the `createConfig` call. Replace:

```ts
import { injected } from "wagmi/connectors";
```

with:

```ts
import { injected, safe, walletConnect } from "wagmi/connectors";
```

Above `createConfig`, build the connector list (keep the existing `ENABLED_CHAINS`-empty guard untouched):

```ts
// WalletConnect is optional: it needs a project id from WalletConnect/Reown
// Cloud. Without VITE_WALLETCONNECT_PROJECT_ID the connector is simply not
// registered, and injected + Safe still work.
const wcProjectId = import.meta.env.VITE_WALLETCONNECT_PROJECT_ID as string | undefined;

const connectors = [
  injected(),
  safe(),
  ...(wcProjectId ? [walletConnect({ projectId: wcProjectId })] : []),
];
```

Change the config to use it:

```ts
export const wagmiConfig = createConfig({
  chains: viemChains as [Chain, ...Chain[]],
  connectors,
  transports: Object.fromEntries(ENABLED_CHAINS.map((c) => [c.id, http(c.rpcUrl)])),
});
```

- [ ] **Step 4: Run test to verify it passes**

Run: `yarn test src/wallet/wagmi.test.ts`
Expected: PASS (existing + new).

- [ ] **Step 5: Add env example and docs**

Create `ui/.env.example`:

```bash
# Optional. WalletConnect / Reown Cloud project id (https://cloud.reown.com).
# When set, the WalletConnect connector is offered in the wallet picker;
# when unset, only Browser Wallet and Safe are available. Never commit a real id.
VITE_WALLETCONNECT_PROJECT_ID=
```

In `ui/CLAUDE.md`, add one line under the environment/config notes:

```markdown
- `VITE_WALLETCONNECT_PROJECT_ID` (optional): enables the WalletConnect wallet option; inert when unset. See `.env.example`.
```

- [ ] **Step 6: Commit**

```bash
git add ui/src/wallet/wagmi.ts ui/src/wallet/wagmi.test.ts ui/.env.example ui/CLAUDE.md
git commit -m "feat(ui): register Safe and optional WalletConnect connectors"
```

---

### Task 7: `ConnectButton` connector picker

**Files:**
- Modify: `ui/src/components/ConnectButton.tsx`
- Test: `ui/src/components/ConnectButton.test.tsx` (create)

**Interfaces:**
- Consumes: the connectors registered in Task 6.
- Behaviour: disconnected state offers every available connector (Safe hidden unless the app is inside an iframe — the Safe-App case) with friendly labels; a single available connector keeps the current one-click button. Connected state is unchanged.

- [ ] **Step 1: Write the failing test**

```tsx
// ui/src/components/ConnectButton.test.tsx
import { afterEach, describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen } from "@testing-library/react";

let mockConnectors: { id: string; name: string; uid: string }[];
const connect = vi.fn();

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined, isConnected: false, chainId: 84532 }),
  useConnect: () => ({ connect, connectors: mockConnectors }),
  useDisconnect: () => ({ disconnect: vi.fn() }),
}));

import { ConnectButton } from "./ConnectButton";

afterEach(() => {
  connect.mockClear();
});

describe("ConnectButton (disconnected)", () => {
  it("shows a picker for multiple connectors and connects the chosen one", () => {
    mockConnectors = [
      { id: "injected", name: "Injected", uid: "a" },
      { id: "walletConnect", name: "WalletConnect", uid: "b" },
    ];
    render(<ConnectButton />);
    fireEvent.click(screen.getByRole("button", { name: /connect wallet/i }));
    fireEvent.click(screen.getByText(/browser wallet/i));
    expect(connect).toHaveBeenCalledWith({ connector: mockConnectors[0] });
  });

  it("keeps a single connector as a one-click button", () => {
    mockConnectors = [{ id: "injected", name: "Injected", uid: "a" }];
    render(<ConnectButton />);
    fireEvent.click(screen.getByRole("button", { name: /connect wallet/i }));
    expect(connect).toHaveBeenCalledWith({ connector: mockConnectors[0] });
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `yarn test src/components/ConnectButton.test.tsx`
Expected: FAIL — no "Browser Wallet" item; current button connects `connectors[0]` directly with no menu.

- [ ] **Step 3: Write minimal implementation**

In `ui/src/components/ConnectButton.tsx`:

Extend the lucide import to add the picker icons:

```ts
import { Check, ChevronDown, Copy, ExternalLink, LogOut, QrCode, ShieldCheck, Wallet } from "lucide-react";
```

Add helpers above the component:

```tsx
// The Safe connector only works inside the Safe-App iframe; hide it elsewhere
// so users outside Safe do not get a dead menu entry.
function inSafeIframe(): boolean {
  return typeof window !== "undefined" && window.parent !== window.self;
}

function connectorMeta(id: string, name: string): { label: string; Icon: typeof Wallet } {
  switch (id) {
    case "injected":
      return { label: "Browser Wallet", Icon: Wallet };
    case "walletConnect":
      return { label: "WalletConnect", Icon: QrCode };
    case "safe":
      return { label: "Safe", Icon: ShieldCheck };
    default:
      return { label: name, Icon: Wallet };
  }
}
```

Replace the disconnected-state early return (the `if (!isConnected || !address) return (...)` block) with:

```tsx
  if (!isConnected || !address) {
    const available = connectors.filter((c) => c.id !== "safe" || inSafeIframe());

    if (available.length <= 1) {
      const only = available[0] ?? connectors[0];
      return (
        <button
          className="inline-flex items-center gap-1.5 rounded-md bg-brand-600 px-3 py-1.5 text-white transition-colors hover:bg-brand-700"
          onClick={() => only && connect({ connector: only })}
        >
          <Wallet className="h-4 w-4" /> Connect Wallet
        </button>
      );
    }

    return (
      <Menu as="div" className="relative">
        <MenuButton className="inline-flex items-center gap-1.5 rounded-md bg-brand-600 px-3 py-1.5 text-white transition-colors hover:bg-brand-700">
          <Wallet className="h-4 w-4" /> Connect Wallet <ChevronDown className="h-4 w-4" />
        </MenuButton>
        <MenuItems
          anchor="bottom end"
          className="z-30 mt-1 w-48 rounded-md border border-tarmac-200 bg-paper-raised py-1 text-sm shadow-lg focus:outline-none dark:border-tarmac-700 dark:bg-tarmac-800"
        >
          {available.map((c) => {
            const { label, Icon } = connectorMeta(c.id, c.name);
            return (
              <MenuItem key={c.uid}>
                <button
                  onClick={() => connect({ connector: c })}
                  className="flex w-full items-center gap-2 px-3 py-1.5 text-left text-tarmac-700 data-[focus]:bg-tarmac-100 dark:text-tarmac-200 dark:data-[focus]:bg-tarmac-700"
                >
                  <Icon className="h-4 w-4 text-tarmac-400" /> {label}
                </button>
              </MenuItem>
            );
          })}
        </MenuItems>
      </Menu>
    );
  }
```

(`Menu`, `MenuButton`, `MenuItem`, `MenuItems` are already imported for the connected-state menu.)

- [ ] **Step 4: Run test to verify it passes**

Run: `yarn test src/components/ConnectButton.test.tsx`
Expected: PASS (2 tests).

- [ ] **Step 5: Commit**

```bash
git add ui/src/components/ConnectButton.tsx ui/src/components/ConnectButton.test.tsx
git commit -m "feat(ui): connector picker in the custom ConnectButton"
```

---

## Final verification

- [ ] **Run the full UI gate**

```bash
cd ui
yarn sync
yarn test        # whole suite green, including the new tests
yarn build       # tsc -b typecheck + vite build succeed
```

- [ ] **Run the repo-root lint** (from `travel-token-messenger-contracts/`)

```bash
yarn lint        # 0 errors
```

- [ ] **Manual smoke (optional, if a dev server is run)**

Start `cd ui && yarn dev`. With no `VITE_WALLETCONNECT_PROJECT_ID`, the connect
button shows a picker with Browser Wallet only (Safe hidden outside an iframe);
the Payment Tokens tab offers "Native currency" / "Off-chain payment" quick-add
and renders existing sentinels with those labels; the Bots tab shows two role
badges per bot.

- [ ] **Open the PR against `dev`** once the branch is green.

---

## Self-review notes (for the plan author)

- **Spec coverage:** Section A → Tasks 6–7; Section B → Tasks 1–4; Section C → Task 5. "Already done" items (GAS_WITHDRAWER affordance, CLI sentinels) are intentionally untouched. `.env.example`/docs folded into Task 6.
- **Type consistency:** `NATIVE_SENTINEL`/`OFFCHAIN_SENTINEL`/`isSentinel`/`paymentTokenLabel` (Task 1) are used unchanged in Tasks 2 and 4; `presets` shape `{ label; value }` is identical in Tasks 3 and 4; `BOT_ROLES` exported in Task 5 matches its test import.
- **No placeholders:** every code step shows complete code; every run step shows the command and expected result.
