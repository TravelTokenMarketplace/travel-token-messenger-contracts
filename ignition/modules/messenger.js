const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const TravelTokenMessengerModule = buildModule("TravelTokenMessengerModule", (m) => {
    /***************************************************
     *                  SET ACCOUNTS                   *
     ***************************************************/

    // Use the first account as the admin. For local node this is the first account
    // from hardhat's example accounts. For Base Sepolia this is the vars from
    // hardhat's config vars BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY (defined in
    // hardhat.config.js).
    const admin = m.getParameter("managerAdmin", m.getAccount(0));

    // NOTE: these can't default to `admin` itself, because hardhat-ignition
    // only resolves an AccountRuntimeValue recursively as another parameter's
    // default -- a ModuleParameterRuntimeValue (which `admin` now is) is left
    // unresolved and crashes ABI encoding. Defaulting to the same account(0)
    // preserves prior behaviour; override these explicitly if managerAdmin is
    // overridden and they should follow it.
    const pauser = m.getParameter("managerPauser", m.getAccount(0));
    const upgrader = m.getParameter("managerUpgrader", m.getAccount(0));
    const versioner = m.getParameter("managerVersioner", m.getAccount(0));

    // Deploy TTMAccountManager implementation contract
    const ttmAccountManager = m.contract("TTMAccountManager");

    // Encode the initialize function call for TTMAccountManager
    const initializeManagerData = m.encodeFunctionCall(ttmAccountManager, "initialize", [
        admin,
        pauser,
        upgrader,
        versioner,
    ]);

    // Deploy the proxy contract for TTMAccountManager with the initialize data
    const ManagerProxy = m.contract("ERC1967Proxy", [ttmAccountManager, initializeManagerData], {
        id: "ManagerERC1967Proxy",
    });

    // Create instance of the proxy contract with the TTMAccountManager ABI
    const managerProxy = m.contractAt("TTMAccountManager", ManagerProxy, { id: "ManagerProxy" });

    /***************************************************
     *             TTM ACCOUNT IMPLEMENTATION          *
     ***************************************************/

    // BookingTokenOperator library
    const bookingTokenOperator = m.library("BookingTokenOperator");

    // Deploy TTMAccount implementation with the BookingTokenOperator library
    const TTMAccountImpl = m.contract("TTMAccount", [], {
        libraries: { BookingTokenOperator: bookingTokenOperator },
    });

    // Set the TTMAccount implementation in the manager
    m.call(managerProxy, "setAccountImplementation", [TTMAccountImpl]);

    /***************************************************
     *                  BOOKING TOKEN                  *
     ***************************************************/

    // Booking token admin and upgrader. Set default to account(0) (the same
    // default `admin` uses) rather than nesting `admin` itself -- see note
    // above: hardhat-ignition doesn't resolve a ModuleParameterRuntimeValue
    // used as another parameter's default. Configurable by a parameter json file.
    const bookingAdmin = m.getParameter("bookingAdmin", m.getAccount(0));
    const bookingUpgrader = m.getParameter("bookingUpgrader", m.getAccount(0));

    // Deploy BookingToken implementation contract
    const bookingToken = m.contract("BookingToken");

    // Encode the initialize function call for BookingToken
    const initializeBookingTokenData = m.encodeFunctionCall(bookingToken, "initialize", [
        managerProxy.address,
        bookingAdmin,
        bookingUpgrader,
    ]);

    // Deploy the proxy contract for BookingToken with the initialize data
    const BookingTokenProxy = m.contract("ERC1967Proxy", [bookingToken, initializeBookingTokenData], {
        id: "BookingTokenERC1967Proxy",
    });

    // Create instance of the proxy contract with the BookingToken ABI
    const bookingTokenProxy = m.contractAt("BookingToken", BookingTokenProxy, { id: "BookingTokenProxy" });

    /***************************************************
     *                   POST CONFIG                   *
     ***************************************************/

    // Set the booking token address in the manager
    m.call(managerProxy, "setBookingTokenAddress", [bookingTokenProxy.address]);

    return { managerProxy, bookingTokenProxy, TTMAccountImpl };
});

module.exports = TravelTokenMessengerModule;
