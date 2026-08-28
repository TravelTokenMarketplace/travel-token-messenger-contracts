import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import { useChainPinnedWrite } from "./useChainPinnedWrite";

const submit = vi.fn().mockResolvedValue("0xhash");
const submitSync = vi.fn();
let mockActiveChainId = 84532;

vi.mock("wagmi", () => ({
  useWriteContract: () => ({
    writeContract: submitSync,
    writeContractAsync: submit,
    isPending: false,
  }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: mockActiveChainId, setActiveChainId: vi.fn() }),
}));

afterEach(() => {
  submit.mockClear();
  submitSync.mockClear();
  mockActiveChainId = 84532;
});

const request = {
  address: "0x1111111111111111111111111111111111111111",
  abi: [],
  functionName: "createTTMAccount",
  args: [],
} as const;

describe("useChainPinnedWrite", () => {
  it("pins the app's active chain onto every write", async () => {
    const { result } = renderHook(() => useChainPinnedWrite());
    await result.current.writeContractAsync(request as never);
    expect(submit).toHaveBeenCalledWith(expect.objectContaining({ chainId: 84532 }));
  });

  it("pins the active chain onto the non-async writeContract too", () => {
    mockActiveChainId = 8453;
    const { result } = renderHook(() => useChainPinnedWrite());
    result.current.writeContract(request as never);
    expect(submitSync).toHaveBeenCalledWith(expect.objectContaining({ chainId: 8453 }));
  });

  it("does not let a caller-supplied chainId override the active chain", async () => {
    const { result } = renderHook(() => useChainPinnedWrite());
    await result.current.writeContractAsync({ ...request, chainId: 1 } as never);
    expect(submit).toHaveBeenCalledWith(expect.objectContaining({ chainId: 84532 }));
  });

  it("stays arity-transparent when the caller passes no options", async () => {
    const { result } = renderHook(() => useChainPinnedWrite());
    await result.current.writeContractAsync(request as never);
    expect(submit.mock.calls[0]).toHaveLength(1);
  });

  it("forwards the caller's request and options untouched", async () => {
    const options = { onSuccess: vi.fn() };
    const { result } = renderHook(() => useChainPinnedWrite());
    await result.current.writeContractAsync(request as never, options as never);
    expect(submit).toHaveBeenCalledWith(
      expect.objectContaining({ functionName: "createTTMAccount", abi: [] }),
      options,
    );
  });
});
