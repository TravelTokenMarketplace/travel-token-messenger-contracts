/**
 * @dev PartnerConfiguration tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

// Fixtures
const {
    setupSigners,
    serviceHash,
    deployTTMAccountManagerFixture,
    deployTTMAccountImplFixture,
    deployTTMAccountManagerWithTTMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployAndConfigureAllWithRegisteredServicesFixture,
} = require("./utils/fixtures");

describe("PartnerConfiguration", function () {
    describe("Services", function () {
        it("should add a supported service correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName = "ttm.service.accommodation.v1alpha.AccommodationSearchService";
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(serviceName));

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash, serviceName);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceHash);

            // Should revert if the same service is added again
            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash, restrictedRate, capabilities),
            )
                .to.be.revertedWithCustomError(ttmAccount, "ServiceAlreadyExists")
                .withArgs(serviceHash);
        });

        it("should remove a supported service correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName = "ttm.service.accommodation.v1alpha.AccommodationSearchService";
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(serviceName));

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash, serviceName);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceHash);

            // Try to remove with non-auth address
            await expect(ttmAccount.connect(signers.otherAccount3).removeService(serviceHash))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount3.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Remove the service
            await expect(ttmAccount.connect(signers.otherAccount1).removeService(serviceHash))
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceHash);

            // Try to remove the service again, should fail
            await expect(
                ttmAccount.connect(signers.otherAccount1).removeService(serviceHash),
            ).to.be.revertedWithCustomError(ttmAccount, "ServiceDoesNotExist");
        });

        it("should remove all supported services correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            // Register multiple services
            const serviceName1 = "ttm.service.accommodation.v1alpha.AccommodationSearchService";
            const serviceName2 = "ttm.service.transport.v1alpha.TransportSearchService";
            const serviceName3 = "ttm.service.activity.v1alpha.ActivitySearchService";

            const serviceHash1 = ethers.keccak256(ethers.toUtf8Bytes(serviceName1));
            const serviceHash2 = ethers.keccak256(ethers.toUtf8Bytes(serviceName2));
            const serviceHash3 = ethers.keccak256(ethers.toUtf8Bytes(serviceName3));

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName1))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash1, serviceName1);
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName2))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash2, serviceName2);
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName3))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash3, serviceName3);

            // Get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            // Add all services
            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash1, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceHash1);
            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash2, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceHash2);
            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash3, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceHash3);

            // Verify services are added
            const [serviceHashes] = await ttmAccount.getSupportedServices();
            expect(serviceHashes).to.have.lengthOf(3);
            expect(serviceHashes).to.include(serviceHash1);
            expect(serviceHashes).to.include(serviceHash2);
            expect(serviceHashes).to.include(serviceHash3);

            // Try to remove all services with non-auth address
            await expect(ttmAccount.connect(signers.otherAccount3).removeAllServices())
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount3.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Remove all services
            await expect(ttmAccount.connect(signers.otherAccount1).removeAllServices())
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceHash1)
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceHash2)
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceHash3);

            // Verify all services are removed
            const [remainingServiceHashes] = await ttmAccount.getSupportedServices();
            expect(remainingServiceHashes).to.have.lengthOf(0);

            // Try to remove all services again, should not fail (no-op)
            await expect(ttmAccount.connect(signers.otherAccount1).removeAllServices()).to.not.emit(
                ttmAccount,
                "ServiceRemoved",
            );
        });

        it("should revert if the caller does not have the SERVICE_ADMIN_ROLE", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName = "ttm.service.accommodation.v1alpha.AccommodationSearchService";
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(serviceName));

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash, serviceName);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE to otherAccount1
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            // Try to add a service with otherAccount2
            await expect(
                ttmAccount.connect(signers.otherAccount2).addService(serviceHash, restrictedRate, capabilities),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount2.address, SERVICE_ADMIN_ROLE);
        });

        it("should add and return all supported services correctly + setter/getters test", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const restrictedRate1 = false;
            const restrictedRate2 = true;
            const restrictedRate3 = false;

            const capabilities1 = ["test capability 1"];
            const capabilities2 = ["test capability 2"];
            const capabilities3 = ["test capability 3"];

            // Add services to TTM account
            expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addService(services.serviceHash1, restrictedRate1, capabilities1),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(services.serviceHash1);

            expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addService(services.serviceHash2, restrictedRate2, capabilities2),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(services.serviceHash2);

            expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addService(services.serviceHash3, restrictedRate3, capabilities3),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(services.serviceHash3);

            // Get all services
            const servicesFromTTMAccount = await ttmAccount.getSupportedServices();
            expect(servicesFromTTMAccount).to.be.deep.equal([
                [services.serviceHash1, services.serviceHash2, services.serviceHash3],
                [
                    [restrictedRate1, capabilities1],
                    [restrictedRate2, capabilities2],
                    [restrictedRate3, capabilities3],
                ],
            ]);

            // Get specific restricted rate for a service hash
            expect(await ttmAccount.getServiceRestrictedRate(services.serviceHash1)).to.be.equal(restrictedRate1);
            expect(await ttmAccount.getServiceRestrictedRate(services.serviceHash2)).to.be.equal(restrictedRate2);
            expect(await ttmAccount.getServiceRestrictedRate(services.serviceHash3)).to.be.equal(restrictedRate3);

            // Get specific capabilities for a service hash
            expect(await ttmAccount.getServiceCapabilities(services.serviceHash1)).to.be.deep.equal(capabilities1);
            expect(await ttmAccount.getServiceCapabilities(services.serviceHash2)).to.be.deep.equal(capabilities2);
            expect(await ttmAccount.getServiceCapabilities(services.serviceHash3)).to.be.deep.equal(capabilities3);

            // TEST SETTERS
            // with new values for each service field

            const newRestrictedRate1 = true;
            const newRestrictedRate2 = false;
            const newRestrictedRate3 = true;

            const newCapabilities1 = ["test capability 4"];
            const newCapabilities2 = ["test capability 5"];
            const newCapabilities3 = ["test capability 6"];

            // Restricted Rate Setter
            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .setServiceRestrictedRate(services.serviceHash1, newRestrictedRate1),
            )
                .to.emit(ttmAccount, "ServiceRestrictedRateUpdated")
                .withArgs(services.serviceHash1, newRestrictedRate1);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .setServiceRestrictedRate(services.serviceHash2, newRestrictedRate2),
            )
                .to.emit(ttmAccount, "ServiceRestrictedRateUpdated")
                .withArgs(services.serviceHash2, newRestrictedRate2);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .setServiceRestrictedRate(services.serviceHash3, newRestrictedRate3),
            ).to.emit(ttmAccount, "ServiceRestrictedRateUpdated");

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    .setServiceRestrictedRate(services.serviceHash1, newRestrictedRate1),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Capabilities Setter
            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .setServiceCapabilities(services.serviceHash1, newCapabilities1),
            )
                .to.emit(ttmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceHash1);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .setServiceCapabilities(services.serviceHash2, newCapabilities2),
            )
                .to.emit(ttmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceHash2);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .setServiceCapabilities(services.serviceHash3, newCapabilities3),
            )
                .to.emit(ttmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceHash3);

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    .setServiceCapabilities(services.serviceHash1, newCapabilities1),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Single Capability add/remove
            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addServiceCapability(services.serviceHash3, "newCapabilities4"),
            )
                .to.emit(ttmAccount, "ServiceCapabilityAdded")
                .withArgs(services.serviceHash3, "newCapabilities4");

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    .addServiceCapability(services.serviceHash3, "newCapabilities4"),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            const newCapabilityList = newCapabilities3.concat(["newCapabilities4"]);

            expect(await ttmAccount.getServiceCapabilities(services.serviceHash3)).to.be.deep.equal(newCapabilityList);

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .removeServiceCapability(services.serviceHash3, "newCapabilities4"),
            )
                .to.emit(ttmAccount, "ServiceCapabilityRemoved")
                .withArgs(services.serviceHash3, "newCapabilities4");

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    .removeServiceCapability(services.serviceHash3, "newCapabilities4"),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // TEST GETTERS with hashes

            // Get specific restricted rate for a service hash
            expect(await ttmAccount.getServiceRestrictedRate(services.serviceHash1)).to.be.equal(newRestrictedRate1);
            expect(await ttmAccount.getServiceRestrictedRate(services.serviceHash2)).to.be.equal(newRestrictedRate2);
            expect(await ttmAccount.getServiceRestrictedRate(services.serviceHash3)).to.be.equal(newRestrictedRate3);

            // Get specific capabilities for a service hash
            expect(await ttmAccount.getServiceCapabilities(services.serviceHash1)).to.be.deep.equal(newCapabilities1);
            expect(await ttmAccount.getServiceCapabilities(services.serviceHash2)).to.be.deep.equal(newCapabilities2);
            expect(await ttmAccount.getServiceCapabilities(services.serviceHash3)).to.be.deep.equal(newCapabilities3);

            // Test failures
            const nonExistingHash = ethers.keccak256(ethers.toUtf8Bytes("NON EXISTING HASH"));

            await expect(ttmAccount.getServiceCapabilities(nonExistingHash))
                .to.be.revertedWithCustomError(ttmAccount, "ServiceDoesNotExist")
                .withArgs(nonExistingHash);
        });

        it("should revert if the service is not registered", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const serviceName = "ttm.service.accommodation.v0.AccommodationSearchService";
            const hash = ethers.keccak256(ethers.toUtf8Bytes(serviceName));
            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, restrictedRate, capabilities),
            ).to.be.revertedWithCustomError(ttmAccountManager, "ServiceNotRegistered");
        });
    });

    describe("Capability removal", function () {
        it("should revert when removing a capability that does not exist", async function () {
            const { ttmAccount, services } = await loadFixture(deployAndConfigureAllWithRegisteredServicesFixture);

            const restrictedRate = false;
            const capabilities = ["existing-capability"];

            await ttmAccount
                .connect(signers.ttmServiceAdmin)
                .addService(services.serviceHash1, restrictedRate, capabilities);

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .removeServiceCapability(services.serviceHash1, "no-such-capability"),
            ).to.be.revertedWithCustomError(ttmAccount, "CapabilityDoesNotExist");
        });
    });

    describe("Unregistered Services", function () {
        it("should remove a service even if it's unregistered on the manager", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName = "ttm.service.accommodation.v1alpha.AccommodationSearchService";
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(serviceName));

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash, serviceName);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceHash);

            // Unregister the service
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceHash, serviceName);

            // Remove the service
            await expect(ttmAccount.connect(signers.otherAccount1).removeService(serviceHash))
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceHash);
        });

        it("should get all services even if one is unregistered on the manager", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName = "ttm.service.accommodation.v1alpha.AccommodationSearchService";
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(serviceName));

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceHash, serviceName);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceHash, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceHash);

            // Unregister the service on the manager
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceHash, serviceName);

            // Try to get all services
            const servicesFromTTMAccount = await ttmAccount.getSupportedServices();
            expect(servicesFromTTMAccount).to.be.deep.equal([[serviceHash], [[restrictedRate, capabilities]]]);

            // The registry deliberately keeps resolving the deprecated name by hash even
            // though the service is no longer registered (this is what the UI's bounded
            // fallback in useResolvedServiceNames/getServiceNameByHash relies on).
            expect(await ttmAccountManager.getServiceNameByHash(serviceHash)).to.equal(serviceName);
            await expect(ttmAccountManager.getRegisteredServiceNameByHash(serviceHash)).to.be.revertedWithCustomError(
                ttmAccountManager,
                "ServiceNotRegistered",
            );
        });

        it("should keep a wanted service hash even if it becomes unregistered on the manager", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName = "ttm.service.accommodation.v1alpha.AccommodationSearchService";
            // Independently computed rather than read back from the contract, so a wrong
            // storage key would still be caught.
            const hash = serviceHash(serviceName);

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(hash, serviceName);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            // Add wanted service by hash
            await expect(ttmAccount.connect(signers.otherAccount1).addWantedServices([hash]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(hash);

            // Unregister the service on the manager
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(hash, serviceName);

            // The account still stores the hash even though the manager no longer knows the name
            expect(await ttmAccount.getWantedServiceHashes()).to.be.deep.equal([hash]);
        });
    });

    describe("Wanted Services", function () {
        it("should add a wanted service correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceHash1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);

            // Try with non-auth address
            await expect(ttmAccount.connect(signers.otherAccount1).addWantedServices([services.serviceHash1]))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());
        });

        it("should add multiple (6) wanted services correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addWantedServices([
                        services.serviceHash1,
                        services.serviceHash2,
                        services.serviceHash3,
                        services.serviceHash4,
                        services.serviceHash5,
                        services.serviceHash6,
                    ]),
            )
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash2)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash3)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash4)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash5)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash6);

            expect(await ttmAccount.getWantedServiceHashes()).to.be.deep.equal([
                services.serviceHash1,
                services.serviceHash2,
                services.serviceHash3,
                services.serviceHash4,
                services.serviceHash5,
                services.serviceHash6,
            ]);
        });

        it("should revert if a wanted service is already added", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceHash1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceHash1]),
            ).to.be.revertedWithCustomError(ttmAccount, "WantedServiceAlreadyExists");
        });

        it("should revert if a wanted service is not registered", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // ttmServiceAdmin genuinely holds SERVICE_ADMIN_ROLE here (granted by the fixture),
            // so this must revert on the registry lookup, not on access control.
            const unregisteredHash = serviceHash("ttm.service.test.v0.NonRegisteredService");

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([unregisteredHash]),
            ).to.be.revertedWithCustomError(ttmAccountManager, "ServiceNotRegistered");
        });

        it("should remove a wanted service correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceHash1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);

            // Try with non-auth address
            await expect(
                ttmAccount.connect(signers.otherAccount1).removeWantedServices([services.serviceHash1]),
            ).to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeWantedServices([services.serviceHash1]))
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceHash1);

            expect(await ttmAccount.getWantedServiceHashes()).to.be.deep.equal([]);
        });

        it("should remove multiple wanted services correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addWantedServices([services.serviceHash1, services.serviceHash2]),
            )
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash2);

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .removeWantedServices([services.serviceHash1, services.serviceHash2]),
            )
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceHash1)
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceHash2);

            expect(await ttmAccount.getWantedServiceHashes()).to.be.deep.equal([]);
        });

        it("should revert removal if a wanted service does not exist", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).removeWantedServices([services.serviceHash1]),
            ).to.be.revertedWithCustomError(ttmAccount, "WantedServiceDoesNotExist");
        });

        it("should add and get multiple wanted services correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceHash1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash1);

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addWantedServices([services.serviceHash2, services.serviceHash3]),
            )
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash2)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceHash3);

            // Get wanted service hashes. services.serviceHash1-3 are independently computed
            // in fixtures.js via ethers.keccak256, not read back from the contract, so this
            // would catch a contract that stored under the wrong key even if self-consistent.
            expect(await ttmAccount.getWantedServiceHashes()).to.be.deep.equal([
                services.serviceHash1,
                services.serviceHash2,
                services.serviceHash3,
            ]);
        });

        it("should return correct status for isServiceSupported", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Prior to adding, it should be unsupported
            expect(await ttmAccount.isServiceSupported(services.serviceHash1)).to.be.false;

            // Add the service
            await ttmAccount.connect(signers.ttmServiceAdmin).addService(services.serviceHash1, false, []);

            // Now it should be supported
            expect(await ttmAccount.isServiceSupported(services.serviceHash1)).to.be.true;

            // Non-registered or other services should still be unsupported
            expect(await ttmAccount.isServiceSupported(services.serviceHash2)).to.be.false;
        });
    });

    describe("Payment", function () {
        it("should set and remove payment info correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Get off chain payment supported expecting false
            expect(await ttmAccount.offChainPaymentSupported()).to.be.equal(false);

            // Set off chain payment supported
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).setOffChainPaymentSupported(true))
                .to.emit(ttmAccount, "OffChainPaymentSupportUpdated")
                .withArgs(true);

            // Try with non-auth address
            await expect(
                ttmAccount.connect(signers.otherAccount1).setOffChainPaymentSupported(false),
            ).to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");

            // Get off chain payment supported expecting true
            expect(await ttmAccount.offChainPaymentSupported()).to.be.equal(true);

            // Set supported tokens
            const supportedToken1 = "0x0000000000000000000000000000000000000001";
            const supportedToken2 = "0x0000000000000000000000000000000000000002";

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addSupportedToken(supportedToken1))
                .to.emit(ttmAccount, "PaymentTokenAdded")
                .withArgs(supportedToken1);

            // Try with non-auth address
            await expect(
                ttmAccount.connect(signers.otherAccount1).addSupportedToken(supportedToken1),
            ).to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addSupportedToken(supportedToken2))
                .to.emit(ttmAccount, "PaymentTokenAdded")
                .withArgs(supportedToken2);

            // Get supported tokens
            const supportedTokens = await ttmAccount.getSupportedTokens();
            expect(supportedTokens).to.be.deep.equal([supportedToken1, supportedToken2]);

            // Revert if token is already supported
            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addSupportedToken(supportedToken1))
                .to.be.revertedWithCustomError(ttmAccount, "PaymentTokenAlreadyExists")
                .withArgs(supportedToken1);

            // Remove supported token
            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeSupportedToken(supportedToken1))
                .to.emit(ttmAccount, "PaymentTokenRemoved")
                .withArgs(supportedToken1);

            // Try with non-auth address
            await expect(
                ttmAccount.connect(signers.otherAccount1).removeSupportedToken(supportedToken1),
            ).to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");

            // Remove it again, should revert
            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeSupportedToken(supportedToken1))
                .to.be.revertedWithCustomError(ttmAccount, "PaymentTokenDoesNotExist")
                .withArgs(supportedToken1);

            // Get supported tokens, should only return supportedToken2
            const supportedTokensAfterRemoval = await ttmAccount.getSupportedTokens();
            expect(supportedTokensAfterRemoval).to.be.deep.equal([supportedToken2]);
        });

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
    });
    describe("PublicKeys", function () {
        it("should add a public keys correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Pubkey
            const pubkey =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr = ethers.computeAddress(pubkey);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addPublicKey(addr, pubkey))
                .to.emit(ttmAccount, "PublicKeyAdded")
                .withArgs(addr);

            // Try with non-auth address
            await expect(
                ttmAccount.connect(signers.otherAccount1).addPublicKey(addr, pubkey),
            ).to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");

            // Get public keys and check if they are correct, should include only addr and pubkey
            const publicKeys = await ttmAccount.getPublicKey(addr);
            expect(publicKeys).to.be.deep.equal(pubkey);
        });

        it("should remove a public key correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Pubkey
            const pubkey =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr = ethers.computeAddress(pubkey);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addPublicKey(addr, pubkey))
                .to.emit(ttmAccount, "PublicKeyAdded")
                .withArgs(addr);

            // Try with non-auth address
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            await expect(ttmAccount.connect(signers.otherAccount1).removePublicKey(addr))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, SERVICE_ADMIN_ROLE);

            // Remove public key
            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removePublicKey(addr))
                .to.emit(ttmAccount, "PublicKeyRemoved")
                .withArgs(addr);

            // Try to remove it again, should revert
            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removePublicKey(addr))
                .to.be.revertedWithCustomError(ttmAccount, "PublicKeyDoesNotExist")
                .withArgs(addr);

            // Get public keys, it should be a array of two empty arrays
            await expect(ttmAccount.getPublicKey(addr))
                .to.be.revertedWithCustomError(ttmAccount, "PublicKeyDoesNotExist")
                .withArgs(addr);
        });

        it("should get public keys and addresses correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Pubkey 1
            const pubkey1 =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr1 = ethers.computeAddress(pubkey1);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addPublicKey(addr1, pubkey1))
                .to.emit(ttmAccount, "PublicKeyAdded")
                .withArgs(addr1);

            // Pubkey 2
            const pubkey2 =
                "0x0407960fdb1ac968edc84eefe2aa4c5edc5b37ea0886eb4efecfd81c5993f9b00c77fc97dd94dc258fcf3c420f8a0601a8cb76030f2ffce68d104e7d83888083e5";
            // Address of the public key
            const addr2 = ethers.computeAddress(pubkey2);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addPublicKey(addr2, pubkey2))
                .to.emit(ttmAccount, "PublicKeyAdded")
                .withArgs(addr2);

            // Get public keys
            const publicKeys = await ttmAccount.getPublicKey(addr1);
            expect(publicKeys).to.be.deep.equal(pubkey1);
            const publicKeys2 = await ttmAccount.getPublicKey(addr2);
            expect(publicKeys2).to.be.deep.equal(pubkey2);

            // Get all public key addresses
            const allPublicKeys = await ttmAccount.getPublicKeysAddresses();
            expect(allPublicKeys).to.be.deep.equal([addr1, addr2]);
        });

        it("should revert when adding the same public key", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            // Pubkey
            const pubkey =
                "0x04fbe3e51d1e56c8ff935360cd32931f5a13ce4aac17f18ed8265c33f06468532fcb8b84eba84c0fae7ce88f64f97e7b6c7cf847b32b697b9e304de7ad2842e6ab";
            // Address of the public key
            const addr = ethers.computeAddress(pubkey);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addPublicKey(addr, pubkey))
                .to.emit(ttmAccount, "PublicKeyAdded")
                .withArgs(addr);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addPublicKey(addr, pubkey))
                .to.be.revertedWithCustomError(ttmAccount, "PublicKeyAlreadyExists")
                .withArgs(addr);
        });
    });
});
