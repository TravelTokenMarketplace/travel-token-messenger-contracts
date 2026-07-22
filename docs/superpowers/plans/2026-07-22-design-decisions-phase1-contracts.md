# Design Decisions — Phase 1 (Contracts) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Implement the contract-side changes for design decisions 2, 3, 4 and 5, leaving the suite green and the ABIs regenerated.

**Architecture:** Four contract changes against `contracts/booking-token/`, `contracts/account/` and `contracts/partner/`. Decisions 2 and 3 are coupled — D3's safety argument depends on D2's fix, so D2 lands first, and D2 additionally requires an `approveERC721` surface on `TTMAccount` without which its scenario is unreachable. Decision 4 is split across four tasks so the suite stays green throughout: add the query, declare tokens in fixtures, enforce at mint, then delete the superseded boolean.

**Tech Stack:** Solidity 0.8.24, Hardhat, `@openzeppelin/contracts-upgradeable` v5, Chai + `hardhat-network-helpers` for tests, Prettier/ESLint/Solhint for lint.

**Spec:** `docs/superpowers/specs/2026-07-22-design-decisions-implementation-design.md`

## Global Constraints

- Solidity `0.8.24`, optimizer `runs: 1000`, `evmVersion: cancun`. Do not change compiler settings.
- Contracts are **UUPS upgradeable**. Storage layout rules apply: append fields, never insert or reorder. Deleting a field is permitted **only** because nothing is deployed (`ignition/deployments/` does not exist; `ui/src/contracts/generated/addresses.ts` exports an empty `ADDRESSES`). Verify this still holds before Task 6.
- Baseline is **159 passing tests** (`yarn test`). Every task must end with at least 159 passing.
- Run `yarn lint` before each commit. Use `yarn format` to auto-fix Prettier.
- Do **not** hand-edit `abi/`. It is regenerated in Task 8.
- Phases 2 (UI) and 3 (deployment wiring) are out of scope for this plan and get their own plans.
- This work happens on branch `feat/design-decisions-implementation`, already created.

---

## File Structure

| File | Responsibility | Tasks |
| --- | --- | --- |
| `contracts/booking-token/BookingTokenCancellable.sol` | Proposal state machine. Gains an explicit `actor` parameter on reject/withdraw so authorization no longer assumes `msg.sender`. | 1 |
| `contracts/booking-token/BookingToken.sol` | Entry points, transfer path, mint. Transfer path keys off owner (T1); owner-must-be-TTM-Account gate (T2); payment enforcement at mint (T5). | 1, 2, 5 |
| `contracts/partner/PartnerConfiguration.sol` | Partner config. Gains `isSupportedToken`; loses `_supportsOffChainPayment`. | 3, 6 |
| `contracts/account/ITTMAccount.sol` | Currently a stub declaring only `initialize`. Gains `isSupportedToken` so `BookingToken` can call it. | 3 |
| `contracts/account/TTMAccount.sol` | Gains `approveERC721` (T1); loses `setOffChainPaymentSupported` (T6); bot role defaults and gas allowance (T7). | 1, 6, 7 |
| `test/utils/fixtures.js` | Shared fixtures. Adds `otherAccount4` (T1); must declare payment tokens before enforcement lands (T4). | 1, 4 |
| `test/BookingToken.test.js` | Transfer, cancellation and mint behaviour. | 1, 2, 5 |
| `test/PartnerConfiguration.test.js` | `isSupportedToken`; off-chain test replaced by a sentinel test. | 3, 6 |
| `test/TTMAccount.test.js` | `addMessengerBot` role grants. | 7 |
| `test/GasMoneyManager.test.js` | Gas defaults; needs explicit role grants once T7 lands. | 7 |

---

### Task 1: Decision 2 — authorize transfer-time proposal closing against the owner

**Files:**
- Modify: `contracts/account/TTMAccount.sol:462-467` (add `approveERC721`)
- Modify: `contracts/booking-token/BookingTokenCancellable.sol:388-450`
- Modify: `contracts/booking-token/BookingToken.sol:607-643`, `:830-848`
- Modify: `test/utils/fixtures.js:14-51` (add `otherAccount4`)
- Test: `test/BookingToken.test.js`

**Interfaces:**
- Consumes: nothing from earlier tasks.
- Produces: `approveERC721(IERC721 token, address to, uint256 tokenId)` on `TTMAccount`, gated on `WITHDRAWER_ROLE`; `signers.otherAccount4` in `test/utils/fixtures.js`; `_rejectCancellation(address actor, address owner, address supplier, uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion)` and `_withdrawCancellation(address actor, address owner, address supplier, uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalVersion)` — both `internal virtual`, both now taking `actor` as the first parameter. Task 2 does not touch these.

**Background:** `checkTransferable` auto-closes a `PENDING` proposal on transfer using protocol code `REJECTION_REASON_TRANSFER_ON_CHAIN = 99`. Two places assume `msg.sender` is a party to the booking: the `onlyOwnerOrSupplier` modifier on the two internal functions, and the branch choosing between them. When an approved operator transfers, both fail and the transfer reverts.

- [ ] **Step 1: Add a free signer**

Tasks 1, 2 and 5 all need an ordinary wallet that holds no role. `setupSigners`
(`test/utils/fixtures.js:14-51`) destructures only `otherAccount1`,
`otherAccount2` and `otherAccount3`, and `deployCancellationSupportFixture`
binds all three (`:286-288`). Add a fourth.

In `test/utils/fixtures.js`, add `otherAccount4` to both the destructuring array
(after `otherAccount3`, at `:31`) and the `signers` object (after
`otherAccount3`, at `:50`). `ethers.getSigners()` returns twenty accounts by
default, so there is capacity.

- [ ] **Step 2: Add `approveERC721` to `TTMAccount`**

Without this, the operator scenario is unreachable and this task cannot be
tested: Task 2 gates cancellation so a pending proposal implies the owner is a
TTM Account, and `TTMAccount` currently has no `approve`, no
`setApprovalForAll` and no generic call surface — so a TTM Account can never
authorize an operator. This is also what makes marketplace and custody
composability possible at all, which is why Option B was chosen.

In `contracts/account/TTMAccount.sol`, directly after `transferERC721`
(`:462-467`), add:

```solidity
    /**
     * @notice Approves an operator to transfer a specific ERC721 token held by
     * this account. Required for listing a booking token on a marketplace or
     * handing it to a custody provider.
     *
     * @param token The ERC721 contract
     * @param to The operator being approved
     * @param tokenId The token id
     */
    function approveERC721(IERC721 token, address to, uint256 tokenId) external onlyRole(WITHDRAWER_ROLE) {
        token.approve(to, tokenId);
    }
```

`WITHDRAWER_ROLE` mirrors `transferERC721`: approving an operator and
transferring outright are the same class of authority. `IERC721` is already
imported for `transferERC721`.

- [ ] **Step 3: Write the failing test**

Add inside the `describe("BookingToken")` block in `test/BookingToken.test.js`:

```javascript
describe("Transfer with pending cancellation", function () {
    it("should close a pending proposal when an approved operator transfers", async function () {
        const {
            supplierTTMAccount,
            distributorTTMAccount,
            bookingToken,
            tokenWithNativePayment,
            supplierBookingOperator,
        } = await loadFixture(deployCancellationSupportFixture);

        // Supplier opens a cancellation proposal on a BOUGHT token that the
        // distributor account owns.
        await supplierTTMAccount
            .connect(supplierBookingOperator)
            .initiateCancellation(tokenWithNativePayment, 0n, 1, 1);

        // The owner approves a third-party operator — a marketplace stand-in.
        const operator = signers.otherAccount4;
        await distributorTTMAccount
            .connect(signers.ttmAccountAdmin)
            .grantRole(await distributorTTMAccount.WITHDRAWER_ROLE(), signers.ttmAccountAdmin.address);
        await distributorTTMAccount
            .connect(signers.ttmAccountAdmin)
            .approveERC721(await bookingToken.getAddress(), operator.address, tokenWithNativePayment);

        // The operator transfers — it is neither the owner nor the supplier.
        await expect(
            bookingToken
                .connect(operator)
                .transferFrom(
                    await distributorTTMAccount.getAddress(),
                    operator.address,
                    tokenWithNativePayment,
                ),
        ).to.emit(bookingToken, "CancellationRejected");

        // Proposal is closed, and the token moved.
        const proposal = await bookingToken.getCancellationProposal(tokenWithNativePayment);
        expect(proposal.status).to.not.equal(1n); // 1 == PENDING
        expect(await bookingToken.ownerOf(tokenWithNativePayment)).to.equal(operator.address);
    });
});
```

- [ ] **Step 4: Run the test to verify it fails**

Run: `yarn test --grep "approved operator transfers"`
Expected: FAIL with `NotOwnerOrSupplier`. `msg.sender` is the operator, which is
neither owner nor supplier, so the `onlyOwnerOrSupplier` modifier on
`_rejectCancellation` reverts. If the failure is anything else, stop and
re-derive — a different error means the setup is wrong, not the premise.

- [ ] **Step 5: Add the `actor` parameter to `_withdrawCancellation`**

In `contracts/booking-token/BookingTokenCancellable.sol`, replace the `_withdrawCancellation` definition (currently at `:388-416`) with:

```solidity
    function _withdrawCancellation(
        address actor,
        address owner,
        address supplier,
        uint256 tokenId,
        uint16 withdrawalReason,
        uint16 withdrawalVersion
    ) internal virtual {
        // Authorize the acting party, which is not necessarily msg.sender: on
        // the transfer path the owner acts through an approved operator.
        if (actor != owner && actor != supplier) {
            revert NotOwnerOrSupplier();
        }

        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if not in PENDING state
        if (proposal.status != CancellationProposalStatus.PENDING) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Only current proposer can withdraw
        if (actor != proposal.currentProposer) {
            revert OnlyCurrentProposerCanWithdrawCancellation(tokenId);
        }

        // Set withdrawal reason
        proposal.withdrawalReason = withdrawalReason;
        proposal.withdrawalVersion = withdrawalVersion;

        // Set status to WITHDRAWN
        proposal.status = CancellationProposalStatus.WITHDRAWN;

        // Emit event
        emit CancellationWithdrawn(tokenId, withdrawalReason, withdrawalVersion);
    }
```

The check order is unchanged — authorization, then status, then proposer — so existing revert-reason assertions keep passing.

- [ ] **Step 6: Add the `actor` parameter to `_rejectCancellation`**

Replace the `_rejectCancellation` definition (currently at `:418-450`) with:

```solidity
    function _rejectCancellation(
        address actor,
        address owner,
        address supplier,
        uint256 tokenId,
        uint16 rejectionReason,
        uint16 rejectionReasonVersion
    ) internal virtual {
        // Authorize the acting party, which is not necessarily msg.sender: on
        // the transfer path the owner acts through an approved operator.
        if (actor != owner && actor != supplier) {
            revert NotOwnerOrSupplier();
        }

        Proposal storage proposal = _getBookingTokenCancellableStorage()._proposals[tokenId];

        // Revert if not in PENDING state
        if (proposal.status != CancellationProposalStatus.PENDING) {
            revert InvalidCancellationProposalStatus(tokenId, proposal.status);
        }

        // Proposer can not reject the cancellation
        if (actor == proposal.currentProposer) {
            revert ProposerCanNotRejectCancellation(tokenId);
        }

        // Set reason
        proposal.rejectionReason = rejectionReason;
        proposal.rejectionVersion = rejectionReasonVersion;

        // Set status to REJECTED
        proposal.status = CancellationProposalStatus.REJECTED;

        // Increment times rejected
        proposal.timesRejected++;

        // Emit event
        emit CancellationRejected(tokenId, rejectionReason, rejectionReasonVersion);
    }
```

Leave `_initiateCancellation`, `_acceptCancellation` and `_counterCancellation` untouched — they keep the `onlyOwnerOrSupplier` modifier, which is correct for direct calls.

- [ ] **Step 7: Update the two external entry points to pass `msg.sender`**

In `contracts/booking-token/BookingToken.sol`, in `withdrawCancellation` (at `:830-838`) change the internal call to:

```solidity
        _withdrawCancellation(msg.sender, owner, supplier, tokenId, withdrawalReason, withdrawalReasonVersion);
```

and in `rejectCancellation` (at `:840-848`):

```solidity
        _rejectCancellation(msg.sender, owner, supplier, tokenId, rejectionReason, rejectionReasonVersion);
```

Behaviour for direct calls is unchanged.

- [ ] **Step 8: Key the transfer path off the owner**

In `checkTransferable` (`contracts/booking-token/BookingToken.sol:622-642`), replace the `if (cancellationStatus == CancellationProposalStatus.PENDING)` block body with:

```solidity
        if (cancellationStatus == CancellationProposalStatus.PENDING) {
            address owner = _requireOwned(tokenId);
            address supplier = $._reservations[tokenId].supplier;

            // The acting party is the owner: a transfer is initiated by the
            // owner or by someone the owner approved. Keying off msg.sender
            // here would revert for marketplace and custody operators.
            if (owner == currentProposer) {
                // The owner is abandoning their own proposal.
                _withdrawCancellation(
                    owner,
                    owner,
                    supplier,
                    tokenId,
                    REJECTION_REASON_TRANSFER_ON_CHAIN,
                    REJECTION_REASON_VERSION
                );
            } else {
                // The counterparty's proposal is rejected.
                _rejectCancellation(
                    owner,
                    owner,
                    supplier,
                    tokenId,
                    REJECTION_REASON_TRANSFER_ON_CHAIN,
                    REJECTION_REASON_VERSION
                );
            }
        }
```

**Documented behaviour difference:** if the *supplier* is an approved operator and is also the current proposer, the old code withdrew and the new code rejects, because the transfer is now attributed to the owner. This is intentional per the spec and only affects that exotic combination.

- [ ] **Step 9: Run the new test and the full suite**

Run: `yarn test --grep "approved operator transfers"`
Expected: PASS

Run: `yarn test`
Expected: at least 160 passing, 0 failing.

- [ ] **Step 10: Lint and commit**

```bash
yarn lint
git add contracts/account/TTMAccount.sol contracts/booking-token/BookingTokenCancellable.sol contracts/booking-token/BookingToken.sol test/BookingToken.test.js test/utils/fixtures.js
git commit -m "fix(booking-token): authorize transfer-time proposal closing against the owner

checkTransferable auto-closes a pending cancellation proposal using protocol
code 99. It assumed msg.sender was a party to the booking, so transfers
initiated by an approved marketplace or custody operator reverted instead.

_rejectCancellation and _withdrawCancellation now take an explicit actor
parameter. External entry points pass msg.sender (unchanged behaviour); the
transfer path passes the owner.

Decision 2."
```

---

### Task 2: Decision 3 — cancellation requires a TTM Account owner

**Files:**
- Modify: `contracts/booking-token/BookingToken.sol:787-800`
- Test: `test/BookingToken.test.js`

**Interfaces:**
- Consumes: Task 1's fixed transfer path — the test below depends on a transfer closing a pending proposal.
- Produces: no new signatures. `_requireBoughtAndParties` keeps its signature and gains one check.

**Background:** all six cancellation entry points route through `_requireBoughtAndParties`. `ownerAccepted` is only ever set true when `msg.sender == owner`, and every entry point is `onlyTTMAccount(msg.sender)`. So when the owner is an ordinary wallet, the supplier can open a proposal that can never be accepted, and `finalizeCancellation` reverts with `OwnerNotAcceptedCancellation` forever.

- [ ] **Step 1: Write the failing test**

Add to `test/BookingToken.test.js`:

```javascript
describe("Cancellation requires a TTM Account owner", function () {
    it("should revert initiateCancellation when the owner is not a TTM Account", async function () {
        const {
            supplierTTMAccount,
            distributorTTMAccount,
            bookingToken,
            tokenWithNativePayment,
            supplierBookingOperator,
        } = await loadFixture(deployCancellationSupportFixture);

        // Move the BOUGHT token out to an ordinary wallet.
        const wallet = signers.otherAccount4;
        const WITHDRAWER_ROLE = await distributorTTMAccount.WITHDRAWER_ROLE();
        await distributorTTMAccount
            .connect(signers.ttmAccountAdmin)
            .grantRole(WITHDRAWER_ROLE, signers.ttmAccountAdmin.address);
        await distributorTTMAccount
            .connect(signers.ttmAccountAdmin)
            .transferERC721(await bookingToken.getAddress(), wallet.address, tokenWithNativePayment);

        expect(await bookingToken.ownerOf(tokenWithNativePayment)).to.equal(wallet.address);

        // The supplier can no longer open a proposal it could never finalize.
        await expect(
            supplierTTMAccount
                .connect(supplierBookingOperator)
                .initiateCancellation(tokenWithNativePayment, 0n, 1, 1),
        ).to.be.revertedWithCustomError(bookingToken, "NotTTMAccount");
    });
});
```

`NotTTMAccount(address account)` is the error `requireTTMAccount` reverts with (`contracts/booking-token/BookingToken.sol:724-728`).

- [ ] **Step 2: Run the test to verify it fails**

Run: `yarn test --grep "not a TTM Account"`
Expected: FAIL — the call currently succeeds, so the assertion reports that no revert happened.

- [ ] **Step 3: Add the gate**

In `contracts/booking-token/BookingToken.sol`, in `_requireBoughtAndParties` (`:787`), add after `owner = _requireOwned(tokenId);`:

```solidity
        // Cancellation is only possible while a TTM Account holds the token.
        // A wallet owner can never call acceptCancellation (every entry point
        // is onlyTTMAccount), so a proposal against one could never finalize.
        requireTTMAccount(owner);
```

This covers all six entry points. `checkTransferable` computes `owner` and `supplier` itself and does not use this helper, so the transfer-time auto-close is unaffected.

- [ ] **Step 4: Run the test and the full suite**

Run: `yarn test --grep "not a TTM Account"`
Expected: PASS

Run: `yarn test`
Expected: at least 161 passing, 0 failing.

- [ ] **Step 5: Commit**

```bash
yarn lint
git add contracts/booking-token/BookingToken.sol test/BookingToken.test.js
git commit -m "feat(booking-token): require a TTM Account owner for cancellation

A wallet owner can never call acceptCancellation, so a supplier could open a
proposal that finalizeCancellation would reject forever. Gate all six entry
points at their shared helper so the failure is immediate and decodable.

Decision 3."
```

---

### Task 3: Decision 4a — add `isSupportedToken`

**Files:**
- Modify: `contracts/partner/PartnerConfiguration.sol:384`
- Modify: `contracts/account/ITTMAccount.sol`
- Test: `test/PartnerConfiguration.test.js`

**Interfaces:**
- Consumes: nothing.
- Produces: `isSupportedToken(address _token) public view virtual returns (bool)` on `PartnerConfiguration`, declared as `external view returns (bool)` on `ITTMAccount`. Task 6 calls it from `BookingToken`.

**Background:** only `getSupportedTokens()` exists, returning the whole array. Looping that from `BookingToken` would make mint gas scale with a partner's allowlist size.

- [ ] **Step 1: Write the failing test**

Add to `test/PartnerConfiguration.test.js`, inside the existing `describe("Payment")` block (`:855`), after the `"should set and remove payment info correctly"` test:

```javascript
it("should report membership via isSupportedToken, including sentinels", async function () {
    const { ttmAccount } = await loadFixture(deployAndConfigureAllWithRegisteredServicesFixture);

    const NATIVE = ethers.ZeroAddress;
    const OFFCHAIN = ethers.getAddress("0x0000000000000000000000000000000000000001");
    const erc20 = signers.otherAccount1.address;

    // Nothing is declared to begin with.
    expect(await ttmAccount.isSupportedToken(NATIVE)).to.be.false;
    expect(await ttmAccount.isSupportedToken(OFFCHAIN)).to.be.false;
    expect(await ttmAccount.isSupportedToken(erc20)).to.be.false;

    // address(0) must be a legitimate set member, not treated as "unset".
    await ttmAccount.connect(signers.ttmServiceAdmin).addSupportedToken(NATIVE);
    await ttmAccount.connect(signers.ttmServiceAdmin).addSupportedToken(OFFCHAIN);
    await ttmAccount.connect(signers.ttmServiceAdmin).addSupportedToken(erc20);

    expect(await ttmAccount.isSupportedToken(NATIVE)).to.be.true;
    expect(await ttmAccount.isSupportedToken(OFFCHAIN)).to.be.true;
    expect(await ttmAccount.isSupportedToken(erc20)).to.be.true;
    expect(await ttmAccount.getSupportedTokens()).to.have.lengthOf(3);

    // Removal is reflected.
    await ttmAccount.connect(signers.ttmServiceAdmin).removeSupportedToken(NATIVE);
    expect(await ttmAccount.isSupportedToken(NATIVE)).to.be.false;
    expect(await ttmAccount.isSupportedToken(OFFCHAIN)).to.be.true;
});
```

`deployAndConfigureAllWithRegisteredServicesFixture` grants `SERVICE_ADMIN_ROLE` to `signers.ttmServiceAdmin` (`test/utils/fixtures.js:536-540`), which is the signer the surrounding `addSupportedToken` assertions already use.

- [ ] **Step 2: Run the test to verify it fails**

Run: `yarn test --grep "isSupportedToken"`
Expected: FAIL with `ttmAccount.isSupportedToken is not a function`.

- [ ] **Step 3: Implement the membership check**

In `contracts/partner/PartnerConfiguration.sol`, directly after `getSupportedTokens` (`:384-387`), add:

```solidity
    /**
     * @notice Returns whether a payment token is declared as supported.
     *
     * The two sentinel values are legitimate members of this set:
     * `address(0)` means native currency and `address(1)` means off-chain
     * payment, matching how `BookingToken` encodes payment mode.
     *
     * @param _token Payment token address, or a sentinel
     * @return supported Whether the token is declared
     */
    function isSupportedToken(address _token) public view virtual returns (bool supported) {
        PartnerConfigurationStorage storage $ = _getPartnerConfigurationStorage();
        return $._paymentInfo._supportedTokens.contains(_token);
    }
```

- [ ] **Step 4: Declare it on the interface**

In `contracts/account/ITTMAccount.sol`, add to the interface body:

```solidity
    function isSupportedToken(address _token) external view returns (bool);
```

- [ ] **Step 5: Run the test and the full suite**

Run: `yarn test --grep "isSupportedToken"`
Expected: PASS

Run: `yarn test`
Expected: at least 162 passing, 0 failing. Nothing is enforced yet, so no existing test changes behaviour.

- [ ] **Step 6: Commit**

```bash
yarn lint
git add contracts/partner/PartnerConfiguration.sol contracts/account/ITTMAccount.sol test/PartnerConfiguration.test.js
git commit -m "feat(partner): add isSupportedToken membership check

BookingToken needs an O(1) membership query; getSupportedTokens returns the
whole array, which would make mint gas scale with allowlist size. Declared on
ITTMAccount so BookingToken can call it. No enforcement yet.

Decision 4."
```

---

### Task 4: Decision 4b — declare payment tokens in fixtures

**Files:**
- Modify: `test/utils/fixtures.js`

**Interfaces:**
- Consumes: Task 3's `isSupportedToken` (not called here, but the declarations it queries are what this task creates).
- Produces: fixtures where every account that mints has declared the tokens it mints in. Task 6's enforcement depends on this.

**Background:** `addSupportedToken` appears zero times in `test/utils/fixtures.js`, and there are 30 mint call sites across the suite. Enforcement added before this task would break every mint path at once, making Task 6 impossible to review. This task is a no-op behaviourally — it must leave the suite green.

- [ ] **Step 1: Declare the tokens each minting account uses**

In `test/utils/fixtures.js`, in `deployBookingTokenWithNullUSDFixture` (`:256`) — or whichever fixture first creates `supplierTTMAccount` alongside `nullUSD` — add declarations after the account is created and before any mint:

```javascript
    // Declare the payment tokens this supplier accepts. Once the allowlist is
    // enforced at mint, a supplier that has declared nothing cannot trade.
    const SUPPORTED_NATIVE = ethers.ZeroAddress;
    const SUPPORTED_OFFCHAIN = ethers.getAddress("0x0000000000000000000000000000000000000001");

    await supplierTTMAccount.connect(signers.ttmAccountAdmin).addSupportedToken(SUPPORTED_NATIVE);
    await supplierTTMAccount.connect(signers.ttmAccountAdmin).addSupportedToken(SUPPORTED_OFFCHAIN);
    await supplierTTMAccount.connect(signers.ttmAccountAdmin).addSupportedToken(await nullUSD.getAddress());
```

`signers.ttmAccountAdmin` is the account's `defaultAdmin`, and `initialize` grants `SERVICE_ADMIN_ROLE` to `defaultAdmin` (`contracts/account/TTMAccount.sol:247`). No extra grant is needed.

- [ ] **Step 2: Cover the other minting accounts**

`deployCancellationSupportFixture` (`:278`) creates `otherTTMAccount` (`:425-460`) which is granted `BOOKING_OPERATOR_ROLE` and is used in failure-case tests. Any fixture-created account that calls `mintBookingToken` needs the same three declarations. Search for mint call sites:

Run: `grep -n "mintBookingToken" test/utils/fixtures.js`

For each distinct calling account, ensure its declarations exist before its first mint.

- [ ] **Step 3: Run the full suite**

Run: `yarn test`
Expected: still at least 162 passing, 0 failing. This task adds configuration only — if any test changes behaviour, something else is wrong; stop and investigate rather than adjusting the test.

- [ ] **Step 4: Commit**

```bash
yarn lint
git add test/utils/fixtures.js
git commit -m "test: declare payment tokens in fixtures

Preparation for allowlist enforcement. No fixture declared any payment token,
so enforcing at mint would break all 30 mint call sites at once. Behaviourally
a no-op today.

Decision 4."
```

---

### Task 5: Decision 4c — enforce the allowlist at mint

**Files:**
- Modify: `contracts/booking-token/BookingToken.sol` (imports, and `safeMintWithReservation` at `:390-448`)
- Test: `test/BookingToken.test.js`

**Interfaces:**
- Consumes: `ITTMAccount.isSupportedToken(address) → bool` from Task 3; fixture declarations from Task 4.
- Produces: `error PaymentTokenNotSupported(address paymentToken)` on `BookingToken`.

**Background:** the supplier is `msg.sender` at mint and sets its own price and payment token. Enforcement is therefore not anti-spoofing — it bounds what a compromised or misconfigured bot holding `BOOKING_OPERATOR_ROLE` can do, using configuration set at `SERVICE_ADMIN_ROLE`.

- [ ] **Step 1: Write the failing tests**

Add to `test/BookingToken.test.js`:

```javascript
describe("Payment token enforcement", function () {
    it("should revert minting with an undeclared payment token", async function () {
        const { supplierTTMAccount, distributorTTMAccount, bookingToken, supplierBookingOperator } =
            await loadFixture(deployCancellationSupportFixture);

        // otherAccount4 was added to setupSigners in Task 1 and holds no role.
        const undeclared = signers.otherAccount4.address;
        const expiration = (await ethers.provider.getBlock("latest")).timestamp + 120;

        await expect(
            supplierTTMAccount.connect(supplierBookingOperator).mintBookingToken(
                await distributorTTMAccount.getAddress(),
                "data:application/json;base64,e30K",
                expiration,
                ethers.parseEther("0.05"),
                undeclared, // never added via addSupportedToken
                0,
                true,
            ),
        ).to.be.revertedWithCustomError(bookingToken, "PaymentTokenNotSupported");
    });

    it("should revert minting when the supplier has declared nothing", async function () {
        const { distributorTTMAccount, bookingToken, otherTTMAccount, otherBookingOperator } =
            await loadFixture(deployCancellationSupportFixture);

        // Strip otherTTMAccount back to no payment configuration at all.
        // getSupportedTokens returns a copy, so removing while iterating is safe.
        // ttmAccountAdmin holds SERVICE_ADMIN_ROLE from initialize.
        for (const token of await otherTTMAccount.getSupportedTokens()) {
            await otherTTMAccount.connect(signers.ttmAccountAdmin).removeSupportedToken(token);
        }
        expect(await otherTTMAccount.getSupportedTokens()).to.have.lengthOf(0);

        const expiration = (await ethers.provider.getBlock("latest")).timestamp + 120;

        await expect(
            otherTTMAccount.connect(otherBookingOperator).mintBookingToken(
                await distributorTTMAccount.getAddress(),
                "data:application/json;base64,e30K",
                expiration,
                ethers.parseEther("0.05"),
                ethers.ZeroAddress,
                0,
                true,
            ),
        ).to.be.revertedWithCustomError(bookingToken, "PaymentTokenNotSupported");
    });
});
```

The existing fixture already exercises the accepting cases — native (token 0), ERC-20 (token 1) and off-chain (token 4) all mint successfully — so no extra positive tests are needed. If Task 4 was done correctly they keep passing.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `yarn test --grep "Payment token enforcement"`
Expected: FAIL — both mints currently succeed.

- [ ] **Step 3: Import the account interface**

In `contracts/booking-token/BookingToken.sol`, next to the existing `ITTMAccountManager` import (`:17`), add:

```solidity
import { ITTMAccount } from "../account/ITTMAccount.sol";
```

- [ ] **Step 4: Declare the error**

Add alongside the other errors in `contracts/booking-token/BookingToken.sol`:

```solidity
    /**
     * @notice The supplier has not declared this payment token as supported.
     *
     * @param paymentToken The rejected payment token, or sentinel
     */
    error PaymentTokenNotSupported(address paymentToken);
```

- [ ] **Step 5: Enforce at mint**

In `safeMintWithReservation`, immediately after `requireTTMAccount(reservedFor);` (`:400`), add:

```solidity
        // The supplier (msg.sender) must have declared this payment token.
        // This bounds what a compromised or misconfigured booking operator can
        // price a booking in. address(0) is native and address(1) is off-chain;
        // both are declared through the same allowlist.
        if (!ITTMAccount(msg.sender).isSupportedToken(address(paymentToken))) {
            revert PaymentTokenNotSupported(address(paymentToken));
        }
```

Do **not** add a matching check to `buyReservedToken`. A reservation is a standing offer; removing a token from configuration must not retroactively break an outstanding reservation.

- [ ] **Step 6: Run the tests and the full suite**

Run: `yarn test --grep "Payment token enforcement"`
Expected: PASS

Run: `yarn test`
Expected: at least 164 passing, 0 failing. If mint tests unrelated to this task now fail, Task 4 missed a minting account — fix the fixture, not the test.

- [ ] **Step 7: Commit**

```bash
yarn lint
git add contracts/booking-token/BookingToken.sol test/BookingToken.test.js
git commit -m "feat(booking-token): enforce the payment token allowlist at mint

Partner payment configuration was decoration: a booking could be priced in any
token. Enforcing it bounds what a compromised booking operator can do, since
minting runs on a hot bot key.

Enforced at mint only. Buy does not re-check, so removing a token cannot
retroactively break an outstanding reservation.

Decision 4."
```

---

### Task 6: Decision 4d — delete the off-chain payment boolean

**Files:**
- Modify: `contracts/partner/PartnerConfiguration.sol:36`, `:91`, `:394-406`
- Modify: `contracts/account/TTMAccount.sol:672-674`
- Test: `test/PartnerConfiguration.test.js:860-876`

**Interfaces:**
- Consumes: Task 3's `isSupportedToken`, Task 5's enforcement.
- Produces: removal of `offChainPaymentSupported()`, `setOffChainPaymentSupported(bool)` and the `OffChainPaymentSupportUpdated` event.

**Background:** off-chain support is now declared by adding `address(1)` to the allowlist, matching how `BookingToken` encodes payment mode. Keeping the boolean would mean two encodings for one concept. **Verify before starting** that `ignition/deployments/` still does not exist and `ui/src/contracts/generated/addresses.ts` still exports an empty `ADDRESSES` — deleting a storage field is only safe because nothing is deployed.

- [ ] **Step 1: Add the replacement test**

The off-chain assertions are not a standalone test — they are the first section
of `"should set and remove payment info correctly"`
(`test/PartnerConfiguration.test.js:856-876`), which continues into supported
tokens. Add this as a **new** test in the same `describe("Payment")` block
first, so coverage exists before anything is deleted:

```javascript
it("should declare off-chain payment support via the address(1) sentinel", async function () {
    const { ttmAccount } = await loadFixture(deployAndConfigureAllWithRegisteredServicesFixture);

    const OFFCHAIN = ethers.getAddress("0x0000000000000000000000000000000000000001");

    expect(await ttmAccount.isSupportedToken(OFFCHAIN)).to.be.false;

    await expect(ttmAccount.connect(signers.ttmServiceAdmin).addSupportedToken(OFFCHAIN))
        .to.emit(ttmAccount, "PaymentTokenAdded")
        .withArgs(OFFCHAIN);

    expect(await ttmAccount.isSupportedToken(OFFCHAIN)).to.be.true;

    await expect(
        ttmAccount.connect(signers.otherAccount1).addSupportedToken(OFFCHAIN),
    ).to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");
});
```

- [ ] **Step 2: Run it and confirm it passes**

Run: `yarn test --grep "address(1) sentinel"`
Expected: PASS. This is not a red-green step — Task 3 already built what it
exercises. Its purpose is to have replacement coverage in place **before** the
deletion, so removing the boolean cannot silently lose it.

- [ ] **Step 3: Remove the off-chain assertions from the existing test**

In `"should set and remove payment info correctly"`, delete the off-chain block
(`test/PartnerConfiguration.test.js:861-876`) — the two
`offChainPaymentSupported()` reads, the `setOffChainPaymentSupported(true)`
emit assertion, and the unauthorized-caller assertion against
`setOffChainPaymentSupported`. Leave the supported-tokens half of the test
intact; it starts at the `// Set supported tokens` comment.

- [ ] **Step 4: Delete the storage field and its accessors**

In `contracts/partner/PartnerConfiguration.sol`:

- Remove `bool _supportsOffChainPayment;` from the `PaymentInfo` struct (`:36`).
- Remove `event OffChainPaymentSupportUpdated(bool supportsOffChainPayment);` (`:91`).
- Remove `_setOffChainPaymentSupported` (`:394-398`) and `offChainPaymentSupported` (`:403-406`).

- [ ] **Step 5: Delete the account-level setter**

In `contracts/account/TTMAccount.sol`, remove `setOffChainPaymentSupported` (`:672-674`).

- [ ] **Step 6: Document the sentinels on the interface**

In `contracts/account/ITTMAccount.sol`, above the `isSupportedToken` declaration added in Task 3, add:

```solidity
    /**
     * @notice Whether a payment token is declared as supported by this account.
     *
     * Payment mode is encoded as an address, matching BookingToken:
     * `address(0)` is native currency, `address(1)` is off-chain payment, and
     * any other value is an ERC-20 address. All three are declared through the
     * same allowlist.
     */
```

- [ ] **Step 7: Compile, run the suite**

Run: `yarn compile`
Expected: success. If anything still references the removed symbols, the compiler names the file and line — fix those references.

Run: `yarn test`
Expected: at least 164 passing, 0 failing.

- [ ] **Step 8: Commit**

```bash
yarn lint
git add contracts/partner/PartnerConfiguration.sol contracts/account/TTMAccount.sol contracts/account/ITTMAccount.sol test/PartnerConfiguration.test.js
git commit -m "refactor(partner): declare off-chain payment through the allowlist

BookingToken already encodes payment mode as an address: address(0) native,
address(1) off-chain, anything else an ERC-20. Configuration now uses the same
encoding, so getSupportedTokens is the complete payment configuration in one
call.

Removes _supportsOffChainPayment and its accessors. Safe only because nothing
is deployed yet.

Decision 4."
```

---

### Task 7: Decision 5 — bot key policy

**Files:**
- Modify: `contracts/account/TTMAccount.sol:257`, `:725-741`
- Test: `test/TTMAccount.test.js`, `test/GasMoneyManager.test.js`

**Interfaces:**
- Consumes: nothing.
- Produces: no signature changes. `addMessengerBot(address bot, uint256 gasMoney)` keeps its signature and grants one fewer role.

**Background:** registering a bot grants three capabilities at once, including withdrawing up to 10 ETH per 24 hours. Bots run as hosted services with hot keys. Confirmed: `grep -rli gasmoney` over the bot repo returns nothing — our bot has never withdrawn gas money.

**Known breakage:** `test/GasMoneyManager.test.js` obtains `GAS_WITHDRAWER_ROLE` exclusively through `addMessengerBot` (at `:74`, `:134`, `:164`, `:226`) and asserts the 10 ETH default in roughly nine places. Both must be updated in this task.

- [ ] **Step 1: Write the failing tests**

Add to `test/TTMAccount.test.js`, in the messenger-bot `describe`:

```javascript
it("should not grant GAS_WITHDRAWER_ROLE when adding a bot", async function () {
    const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
    const bot = signers.otherAccount1;

    await ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(bot.address, 0n);

    expect(await ttmAccount.hasRole(await ttmAccount.MESSENGER_BOT_ROLE(), bot.address)).to.be.true;
    expect(await ttmAccount.hasRole(await ttmAccount.BOOKING_OPERATOR_ROLE(), bot.address)).to.be.true;
    expect(await ttmAccount.hasRole(await ttmAccount.GAS_WITHDRAWER_ROLE(), bot.address)).to.be.false;
});

it("should let the default admin grant GAS_WITHDRAWER_ROLE explicitly", async function () {
    const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
    const bot = signers.otherAccount1;
    const GAS_WITHDRAWER_ROLE = await ttmAccount.GAS_WITHDRAWER_ROLE();

    await ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(bot.address, 0n);
    await ttmAccount.connect(signers.ttmAccountAdmin).grantRole(GAS_WITHDRAWER_ROLE, bot.address);

    expect(await ttmAccount.hasRole(GAS_WITHDRAWER_ROLE, bot.address)).to.be.true;

    // Removal still fully de-authorizes a bot granted the role later.
    await ttmAccount.connect(signers.ttmAccountAdmin).removeMessengerBot(bot.address);
    expect(await ttmAccount.hasRole(GAS_WITHDRAWER_ROLE, bot.address)).to.be.false;
});

it("should default the gas allowance to 0.01 ETH per 24 hours", async function () {
    const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

    const [limit, period] = await ttmAccount.getGasMoneyWithdrawal();
    expect(limit).to.equal(ethers.parseEther("0.01"));
    expect(period).to.equal(24n * 60n * 60n);
});
```

`deployAndConfigureAllFixture` is the fixture the surrounding messenger-bot tests already use (`test/TTMAccount.test.js:310`).

- [ ] **Step 2: Run them to verify they fail**

Run: `yarn test --grep "GAS_WITHDRAWER_ROLE|0.01 ETH"`
Expected: FAIL — the role is currently granted by default and the limit is 10 ETH.

- [ ] **Step 3: Stop granting the role by default**

In `contracts/account/TTMAccount.sol`, in `addMessengerBot` (`:725`), delete this line:

```solidity
        _grantRole(GAS_WITHDRAWER_ROLE, bot);
```

Leave the revoke in `removeMessengerBot` (`:746`) in place — revoking an unheld role is a no-op, and keeping it means removal still fully de-authorizes a bot granted the role later.

Update the contract's doc comment at `:85-88` so it no longer says the role is granted by `addMessengerBot`. State instead that `DEFAULT_ADMIN_ROLE` grants it explicitly, and note the resulting separation: a `BOT_ADMIN_ROLE` holder can onboard bots but cannot give them access to funds.

- [ ] **Step 4: Lower the default allowance**

In `contracts/account/TTMAccount.sol:257`, change:

```solidity
        uint256 withdrawalLimit = 10 ether; // 10 ETH
```

to:

```solidity
        uint256 withdrawalLimit = 0.01 ether; // 0.01 ETH
```

Leave `withdrawalPeriod` at `24 hours`.

- [ ] **Step 5: Fix the gas money tests**

In `test/GasMoneyManager.test.js`, at each of `:74`, `:134`, `:164`, `:226`, the `addMessengerBot` call no longer confers withdrawal rights. Add an explicit grant immediately after each:

```javascript
            await ttmAccount
                .connect(signers.ttmAccountAdmin)
                .grantRole(await ttmAccount.GAS_WITHDRAWER_ROLE(), withdrawer.address);
```

Then update the limit expectations. Two distinct groups:

- Tests that assert the *default* (at `:22`, `:32`, `:270`) become
  `ethers.parseEther("0.01")`.
- Tests that *withdraw* amounts larger than the new default — including the one
  at `:74`, which withdraws 1 ETH, as well as `:131`, `:161`, `:223`, `:231`
  and `:253` — must call `setGasMoneyWithdrawal` to set the limit they need
  before withdrawing. Do not rewrite them to withdraw smaller amounts: they
  should test the mechanism, not the constant.

Note that `:74` needs **both** changes — the explicit role grant from above and
a raised limit — because 1 ETH exceeds the new 0.01 ETH default.

Run: `yarn test test/GasMoneyManager.test.js`
Expected: all passing. Work through failures one at a time; each is either a missing grant or a stale limit.

- [ ] **Step 6: Run the full suite**

Run: `yarn test`
Expected: at least 167 passing, 0 failing.

- [ ] **Step 7: Commit**

```bash
yarn lint
git add contracts/account/TTMAccount.sol test/TTMAccount.test.js test/GasMoneyManager.test.js
git commit -m "feat(account): drop gas withdrawal from bot defaults, lower allowance

Registering a bot granted three capabilities at once, including withdrawing up
to 10 ETH per 24 hours, to an address that runs as a hosted service with a hot
key. Gas withdrawal is now granted explicitly by DEFAULT_ADMIN, so a delegated
BOT_ADMIN can onboard bots but cannot give them access to funds.

Default allowance 10 ETH -> 0.01 ETH per 24 hours. Our bot never used this
path (no references to gas money in the bot repo).

Decision 5."
```

---

### Task 8: Regenerate ABIs and docs

**Files:**
- Modify: `abi/` (generated), `docs/` (generated)

**Interfaces:**
- Consumes: all preceding tasks.
- Produces: `abi/` in sync with the contracts, which phase 2 (UI) consumes via `ui/scripts/sync-contracts.ts`.

- [ ] **Step 1: Compile and export**

```bash
yarn compile
yarn hardhat export-abi
yarn docgen
```

- [ ] **Step 2: Confirm the ABI reflects the changes**

```bash
grep -c "offChainPaymentSupported" abi/*.json || echo "absent, as expected"
grep -c "isSupportedToken" abi/*.json
grep -c "PaymentTokenNotSupported" abi/*.json
```

Expected: `offChainPaymentSupported` absent; `isSupportedToken` and `PaymentTokenNotSupported` present.

- [ ] **Step 3: Full verification**

```bash
yarn test
yarn lint
```

Expected: at least 167 passing, 0 failing; lint clean.

- [ ] **Step 4: Commit**

```bash
git add abi/ docs/
git commit -m "chore: regenerate ABIs and docs after decision changes

abi/ is what the UI consumes via ui/scripts/sync-contracts.ts."
```

---

### Task 9: Update `BOT-MIGRATION.md`

**Files:**
- Modify: `../BOT-MIGRATION.md` (workspace parent folder — **not** inside any repo)

**Interfaces:**
- Consumes: all preceding tasks.
- Produces: nothing consumed by later tasks.

**Background:** `BOT-MIGRATION.md` lives in the workspace parent folder on purpose, alongside `REBRANDING.md`, `TODOS.md` and `CONTRACTS-NEXT.md`. **Do not commit it, or references to it, into any repo.**

- [ ] **Step 1: Add a section for these changes**

Append a section recording three items, written against the code as it now stands rather than against this plan:

- **Decision 4 changes onboarding order.** `safeMintWithReservation` reverts with `PaymentTokenNotSupported` until the supplier has declared the payment token. A partner who has declared nothing cannot trade at all. Payment mode is encoded as an address: `address(0)` native, `address(1)` off-chain, otherwise the ERC-20 address. `offChainPaymentSupported()` and `setOffChainPaymentSupported(bool)` no longer exist.
- **Decision 3 adds a new revert.** Cancellation entry points now revert when the token's owner is not a registered TTM Account. The bot should surface this rather than retry — it is terminal until the token returns to a TTM Account.
- **Decision 5 is confirmed no-impact.** `addMessengerBot` no longer grants `GAS_WITHDRAWER_ROLE`, and the default allowance is 0.01 ETH per 24 hours. The bot has no references to gas money, so nothing changes for it. Recorded so the question is not reopened.

Verify each signature and error name against the actual code before writing it down, in keeping with how the rest of that file was written.

- [ ] **Step 2: Confirm it is not staged into the repo**

```bash
cd /hgst/work/github.com/TravelTokenMarketplace/travel-token-messenger/travel-token-messenger-contracts
git status --short
```

Expected: `BOT-MIGRATION.md` does not appear. It lives outside the repo and is not tracked.

---

## Done criteria

- `yarn test` — at least 167 passing, 0 failing
- `yarn lint` — clean
- `abi/` regenerated and committed
- `BOT-MIGRATION.md` updated in the workspace parent folder, uncommitted
- Branch `feat/design-decisions-implementation` ready for a PR against `dev`

Phases 2 (UI) and 3 (deployment wiring) follow in their own plans. The Base Sepolia deployment happens after all three; there is no schedule pressure on it.
