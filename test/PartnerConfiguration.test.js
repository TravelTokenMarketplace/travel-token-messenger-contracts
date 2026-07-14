/**
 * @dev PartnerConfiguration tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

// Fixtures
const {
    setupSigners,
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
                .withArgs(serviceName, serviceHash);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(ttmAccount.connect(signers.otherAccount1).addService(serviceName, restrictedRate, capabilities))
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceName);

            // Should revert if the same service is added again
            await expect(ttmAccount.connect(signers.otherAccount1).addService(serviceName, restrictedRate, capabilities))
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
                .withArgs(serviceName, serviceHash);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(ttmAccount.connect(signers.otherAccount1).addService(serviceName, restrictedRate, capabilities))
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceName);

            // Try to remove with non-auth address
            await expect(ttmAccount.connect(signers.otherAccount3).removeService(serviceName))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount3.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Remove the service
            await expect(ttmAccount.connect(signers.otherAccount1).removeService(serviceName))
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceName);

            // Try to remove the service again, should fail
            await expect(
                ttmAccount.connect(signers.otherAccount1).removeService(serviceName),
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
                .withArgs(serviceName1, serviceHash1);
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName2))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceName2, serviceHash2);
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName3))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceName3, serviceHash3);

            // Get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            // Add all services
            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceName1, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceName1);
            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceName2, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceName2);
            await expect(
                ttmAccount.connect(signers.otherAccount1).addService(serviceName3, restrictedRate, capabilities),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceName3);

            // Verify services are added
            const [serviceNames] = await ttmAccount.getSupportedServices();
            expect(serviceNames).to.have.lengthOf(3);
            expect(serviceNames).to.include(serviceName1);
            expect(serviceNames).to.include(serviceName2);
            expect(serviceNames).to.include(serviceName3);

            // Try to remove all services with non-auth address
            await expect(ttmAccount.connect(signers.otherAccount3).removeAllServices())
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount3.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Remove all services
            await expect(ttmAccount.connect(signers.otherAccount1).removeAllServices())
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceName1)
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceName2)
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceName3);

            // Verify all services are removed
            const [remainingServiceNames] = await ttmAccount.getSupportedServices();
            expect(remainingServiceNames).to.have.lengthOf(0);

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
                .withArgs(serviceName, serviceHash);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE to otherAccount1
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            // Try to add a service with otherAccount2
            await expect(ttmAccount.connect(signers.otherAccount2).addService(serviceName, restrictedRate, capabilities))
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
                    .addService(services.serviceName1, restrictedRate1, capabilities1),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(services.serviceName1);

            expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addService(services.serviceName2, restrictedRate2, capabilities2),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(services.serviceName2);

            expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addService(services.serviceName3, restrictedRate3, capabilities3),
            )
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(services.serviceName3);

            // Get all services
            const servicesFromTTMAccount = await ttmAccount.getSupportedServices();
            expect(servicesFromTTMAccount).to.be.deep.equal([
                [services.serviceName1, services.serviceName2, services.serviceName3],
                [
                    [restrictedRate1, capabilities1],
                    [restrictedRate2, capabilities2],
                    [restrictedRate3, capabilities3],
                ],
            ]);

            // Get specific restricted rate for a service name
            expect(await ttmAccount["getServiceRestrictedRate(string)"](services.serviceName1)).to.be.equal(
                restrictedRate1,
            );
            expect(await ttmAccount["getServiceRestrictedRate(string)"](services.serviceName2)).to.be.equal(
                restrictedRate2,
            );
            expect(await ttmAccount["getServiceRestrictedRate(string)"](services.serviceName3)).to.be.equal(
                restrictedRate3,
            );

            // Get specific capabilities for a service name
            expect(await ttmAccount["getServiceCapabilities(string)"](services.serviceName1)).to.be.deep.equal(
                capabilities1,
            );
            expect(await ttmAccount["getServiceCapabilities(string)"](services.serviceName2)).to.be.deep.equal(
                capabilities2,
            );
            expect(await ttmAccount["getServiceCapabilities(string)"](services.serviceName3)).to.be.deep.equal(
                capabilities3,
            );

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
                    ["setServiceRestrictedRate(string,bool)"](services.serviceName1, newRestrictedRate1),
            )
                .to.emit(ttmAccount, "ServiceRestrictedRateUpdated")
                .withArgs(services.serviceName1, newRestrictedRate1);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    ["setServiceRestrictedRate(string,bool)"](services.serviceName2, newRestrictedRate2),
            )
                .to.emit(ttmAccount, "ServiceRestrictedRateUpdated")
                .withArgs(services.serviceName2, newRestrictedRate2);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    ["setServiceRestrictedRate(string,bool)"](services.serviceName3, newRestrictedRate3),
            ).to.emit(ttmAccount, "ServiceRestrictedRateUpdated");

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    ["setServiceRestrictedRate(string,bool)"](services.serviceName1, newRestrictedRate1),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Capabilities Setter
            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    ["setServiceCapabilities(string,string[])"](services.serviceName1, newCapabilities1),
            )
                .to.emit(ttmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceName1);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    ["setServiceCapabilities(string,string[])"](services.serviceName2, newCapabilities2),
            )
                .to.emit(ttmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceName2);

            await expect(
                await ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    ["setServiceCapabilities(string,string[])"](services.serviceName3, newCapabilities3),
            )
                .to.emit(ttmAccount, "ServiceCapabilitiesUpdated")
                .withArgs(services.serviceName3);

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    ["setServiceCapabilities(string,string[])"](services.serviceName1, newCapabilities1),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // Single Capability add/remove
            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addServiceCapability(services.serviceName3, "newCapabilities4"),
            )
                .to.emit(ttmAccount, "ServiceCapabilityAdded")
                .withArgs(services.serviceName3, "newCapabilities4");

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    .addServiceCapability(services.serviceName3, "newCapabilities4"),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            const newCapabilityList = newCapabilities3.concat(["newCapabilities4"]);

            expect(await ttmAccount["getServiceCapabilities(string)"](services.serviceName3)).to.be.deep.equal(
                newCapabilityList,
            );

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .removeServiceCapability(services.serviceName3, "newCapabilities4"),
            )
                .to.emit(ttmAccount, "ServiceCapabilityRemoved")
                .withArgs(services.serviceName3, "newCapabilities4");

            // Try with non-auth address
            await expect(
                ttmAccount
                    .connect(signers.otherAccount1)
                    .removeServiceCapability(services.serviceName3, "newCapabilities4"),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.SERVICE_ADMIN_ROLE());

            // TEST GETTERS with hashes

            // Get specific restricted rate for a service name
            expect(await ttmAccount["getServiceRestrictedRate(bytes32)"](services.serviceHash1)).to.be.equal(
                newRestrictedRate1,
            );
            expect(await ttmAccount["getServiceRestrictedRate(bytes32)"](services.serviceHash2)).to.be.equal(
                newRestrictedRate2,
            );
            expect(await ttmAccount["getServiceRestrictedRate(bytes32)"](services.serviceHash3)).to.be.equal(
                newRestrictedRate3,
            );

            // Get specific capabilities for a service name
            expect(await ttmAccount["getServiceCapabilities(bytes32)"](services.serviceHash1)).to.be.deep.equal(
                newCapabilities1,
            );
            expect(await ttmAccount["getServiceCapabilities(bytes32)"](services.serviceHash2)).to.be.deep.equal(
                newCapabilities2,
            );
            expect(await ttmAccount["getServiceCapabilities(bytes32)"](services.serviceHash3)).to.be.deep.equal(
                newCapabilities3,
            );

            // Test failures
            const nonExistingHash = ethers.keccak256(ethers.toUtf8Bytes("NON EXISTING HASH"));

            await expect(ttmAccount["getServiceCapabilities(bytes32)"](nonExistingHash))
                .to.be.revertedWithCustomError(ttmAccount, "ServiceDoesNotExist")
                .withArgs(nonExistingHash);
        });

        it("should revert if the service is not registered", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const serviceName = "ttm.service.accommodation.v0.AccommodationSearchService";
            const restrictedRate = false;
            const capabilities = [];

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addService(serviceName, restrictedRate, capabilities),
            ).to.be.revertedWithCustomError(ttmAccountManager, "ServiceNotRegistered");
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
                .withArgs(serviceName, serviceHash);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(ttmAccount.connect(signers.otherAccount1).addService(serviceName, restrictedRate, capabilities))
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceName);

            // Unregister the service
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceName, serviceHash);

            // Remove the service
            await expect(ttmAccount.connect(signers.otherAccount1).removeService(serviceName))
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(serviceName);
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
                .withArgs(serviceName, serviceHash);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            const restrictedRate = false;
            const capabilities = [];

            await expect(ttmAccount.connect(signers.otherAccount1).addService(serviceName, restrictedRate, capabilities))
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(serviceName);

            // Unregister the service on the manager
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceName, serviceHash);

            // Try to get all services
            const servicesFromTTMAccount = await ttmAccount.getSupportedServices();
            expect(servicesFromTTMAccount).to.be.deep.equal([[serviceName], [[restrictedRate, capabilities]]]);
        });

        it("should get all wanted services even if one is unregistered on the manager", async function () {
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
                .withArgs(serviceName, serviceHash);

            // get the SERVICE_ADMIN_ROLE
            const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

            // Grant SERVICE_ADMIN_ROLE
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(SERVICE_ADMIN_ROLE, signers.otherAccount1.address, signers.ttmAccountAdmin.address);

            // Add wanted service
            await expect(ttmAccount.connect(signers.otherAccount1).addWantedServices([serviceName]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(serviceName);

            // Unregister the service on the manager
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceName, serviceHash);

            // Try to get all wanted services
            const wantedServicesFromTTMAccount = await ttmAccount.getWantedServices();
            expect(wantedServicesFromTTMAccount).to.be.deep.equal([serviceName]);
        });
    });

    describe("Wanted Services", function () {
        it("should add a wanted service correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName1);

            // Try with non-auth address
            await expect(ttmAccount.connect(signers.otherAccount1).addWantedServices([services.serviceName1]))
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
                        services.serviceName1,
                        services.serviceName2,
                        services.serviceName3,
                        services.serviceName4,
                        services.serviceName5,
                        services.serviceName6,
                    ]),
            )
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName1)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName2)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName3)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName4)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName5)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName6);
        });

        it("should revert if a wanted service is already added", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName1);

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceName1]),
            ).to.be.revertedWithCustomError(ttmAccount, "WantedServiceAlreadyExists");
        });

        it("should revert if a wanted service is not registered", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addWantedServices(["ttm.service.test.v0.NonRegisteredService"]),
            ).to.be.revertedWithCustomError(ttmAccountManager, "ServiceNotRegistered");
        });

        it("should remove a wanted service correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName1);

            // Try with non-auth address
            await expect(
                ttmAccount.connect(signers.otherAccount1).removeWantedServices([services.serviceName1]),
            ).to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeWantedServices([services.serviceName1]))
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceName1);
        });

        it("should remove multiple wanted services correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addWantedServices([services.serviceName1, services.serviceName2]),
            )
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName1)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName2);

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .removeWantedServices([services.serviceName1, services.serviceName2]),
            )
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceName1)
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(services.serviceName2);
        });

        it("should revert removal if a wanted service does not exist", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).removeWantedServices([services.serviceName1]),
            ).to.be.revertedWithCustomError(ttmAccount, "WantedServiceDoesNotExist");
        });

        it("should add and get multiple wanted services correctly", async function () {
            const { ttmAccountManager, ttmAccount, services } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([services.serviceName1]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName1);

            await expect(
                ttmAccount
                    .connect(signers.ttmServiceAdmin)
                    .addWantedServices([services.serviceName2, services.serviceName3]),
            )
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName2)
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(services.serviceName3);

            // Get wanted services
            expect(await ttmAccount.getWantedServices()).to.be.deep.equal([
                services.serviceName1,
                services.serviceName2,
                services.serviceName3,
            ]);

            // Get wanted service by hash
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
            expect(await ttmAccount.isServiceSupported(services.serviceName1)).to.be.false;

            // Add the service
            await ttmAccount.connect(signers.ttmServiceAdmin).addService(services.serviceName1, false, []);

            // Now it should be supported
            expect(await ttmAccount.isServiceSupported(services.serviceName1)).to.be.true;

            // Non-registered or other services should still be unsupported
            expect(await ttmAccount.isServiceSupported(services.serviceName2)).to.be.false;
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
