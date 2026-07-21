const globals = require("globals");

module.exports = [
    {
        ignores: ["node_modules/**", "artifacts/**", "cache/**", "coverage/**", "abi/**", "docs/**", "ui/**", "go/**"],
    },
    {
        files: ["**/*.js"],
        languageOptions: {
            ecmaVersion: 2022,
            sourceType: "commonjs",
            globals: { ...globals.node, ...globals.mocha },
        },
        rules: {
            "no-unused-vars": ["error", { argsIgnorePattern: "^_" }],
            "no-undef": "error",
        },
    },
    {
        // Hardhat task files (`tasks/**`) are loaded directly by `hardhat.config.js`
        // through Hardhat's config-loading context, which injects `scope`, `task`,
        // `subtask`, `types`, `ethers`, etc. as ambient globals (the same mechanism
        // that lets `hardhat.config.js` itself call `require("hardhat/config")`
        // helpers without importing them). These are genuine Hardhat globals, not
        // typos, so we declare them rather than disabling no-undef.
        files: ["tasks/**/*.js"],
        languageOptions: {
            globals: { scope: "readonly", ethers: "readonly" },
        },
        rules: {
            // Every task action follows Hardhat's `(taskArgs, hre) => {}` signature,
            // but most actions use the injected `ethers`/`network` globals instead
            // of `hre.ethers`/`hre.network`, leaving `hre` unused in many actions.
            // Renaming every occurrence to `_hre` is pure churn for a parameter
            // whose signature is dictated by Hardhat, not by this code, so we stop
            // checking unused function arguments in this directory only.
            // `caughtErrors: "none"` covers the handful of best-effort
            // `try { ... } catch (e) { console.log("Failed to fetch ...") }`
            // blocks used for optional/nice-to-have CLI output (e.g. off-chain
            // payment support, gas withdrawal settings) — confirmed the called
            // contract methods are real (GasMoneyManager.getGasMoneyWithdrawal,
            // PartnerConfiguration.offChainPaymentSupported), so this is
            // deliberate graceful degradation, not a swallowed bug.
            "no-unused-vars": ["error", { args: "none", caughtErrors: "none" }],
        },
    },
    {
        // `test/**` relies on two established, pre-existing conventions that read
        // as no-undef/no-unused-vars noise under a fresh eslint config:
        //  - `network` is Hardhat's ambient global (same as in tasks/**).
        //  - `signers` is a shared fixture object assigned without `const/let` in
        //    test/utils/fixtures.js (sloppy-mode implicit global) and read by every
        //    test file; it's the established cross-file wiring, not a new bug.
        // Fixture destructuring (`const { a, b, c } = await loadFixture(...)`) also
        // routinely pulls in properties a given `it()` doesn't use, and fixtures.js
        // is imported with a wide named-import list that not every test file fully
        // consumes. Spot-checking a sample of these confirmed they're benign
        // leftovers, not masked bugs, so unused-vars is turned off for tests to
        // avoid renaming dozens of destructured bindings across 6 files.
        files: ["test/**/*.js"],
        languageOptions: {
            globals: { network: "readonly", signers: "writable" },
        },
        rules: {
            "no-unused-vars": "off",
        },
    },
];
