import "@testing-library/jest-dom/vitest";

// jsdom has no ResizeObserver; @headlessui/react's Menu anchor positioning
// (floating-ui autoUpdate) needs one whenever a Menu is actually opened in a test.
if (typeof globalThis.ResizeObserver === "undefined") {
  globalThis.ResizeObserver = class ResizeObserver {
    observe() {}
    unobserve() {}
    disconnect() {}
  };
}
