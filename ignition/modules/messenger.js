const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const TravelTokenMessengerModule = buildModule("TravelTokenMessengerModule", (m) => {
    /***************************************************
     *                  SET ACCOUNTS                   *
     ***************************************************/

    // Use the first account as the admin. For local node this is the first account
    // from hardhat's example accounts. For Base Sepolia this is the vars from
    // hardhat's config vars BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY (defined in
    // hardhat.config.js).
    //const admin = m.getParameter("managerAdmin", m.getAccount(0));
    const admin = m.getAccount(0);

    const pauser = m.getParameter("managerPauser", admin);
    const upgrader = m.getParameter("managerUpgrader", admin);
    const versioner = m.getParameter("managerVersioner", admin);

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

    // Booking token admin and upgrader. Set default to the manager admin.
    // Configurable by a parameter json file.
    const bookingAdmin = m.getParameter("bookingAdmin", admin);
    const bookingUpgrader = m.getParameter("bookingUpgrader", admin);

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
