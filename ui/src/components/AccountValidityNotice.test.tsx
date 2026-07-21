import { afterEach, describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { AccountValidityNotice } from "./AccountValidityNotice";

let mockIsTTMAccount: boolean | undefined;
let mockIsLoading = false;

vi.mock("wagmi", () => ({
  useReadContract: () => ({ data: mockIsTTMAccount, isLoading: mockIsLoading }),
}));
vi.mock("../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({
    manager: "0x9999999999999999999999999999999999999999",
    managerAbi: [],
    chainId: 84532,
    supported: true,
  }),
}));

const validAccount = "0x1111111111111111111111111111111111111111" as const;

afterEach(() => {
  mockIsTTMAccount = undefined;
  mockIsLoading = false;
});

describe("AccountValidityNotice", () => {
  it("warns when the manager does not recognize the address as a TTM Account", () => {
    mockIsTTMAccount = false;
    render(<AccountValidityNotice account={validAccount} />);
    expect(screen.getByText(/not a ttm account/i)).toBeInTheDocument();
    expect(screen.getByText(/does not recognize this address as a ttm account/i)).toBeInTheDocument();
  });

  it("renders nothing while the manager read is loading", () => {
    mockIsTTMAccount = undefined;
    mockIsLoading = true;
    const { container } = render(<AccountValidityNotice account={validAccount} />);
    expect(container).toBeEmptyDOMElement();
  });

  it("renders nothing once the manager confirms the address is a TTM Account", () => {
    mockIsTTMAccount = true;
    const { container } = render(<AccountValidityNotice account={validAccount} />);
    expect(container).toBeEmptyDOMElement();
  });

  it("flags an invalid address without calling the manager", () => {
    render(<AccountValidityNotice account={"not-an-address" as unknown as typeof validAccount} />);
    expect(screen.getByText(/invalid address/i)).toBeInTheDocument();
  });
});
