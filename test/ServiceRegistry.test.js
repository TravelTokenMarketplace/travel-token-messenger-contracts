/**
 * @dev ServiceRegistry tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

const {
    setupSigners,
    deployTTMAccountManagerFixture,
    deployTTMAccountImplFixture,
    deployTTMAccountManagerWithTTMAccountImplFixture,
    deployAndConfigureAllFixture,
} = require("./utils/fixtures");

describe("ServiceRegistry", function () {
    describe("Main", function () {
        it("should register a service correctly", async function () {
            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

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

            await expect(await ttmAccountManager.getRegisteredServiceHashByName(serviceName)).to.be.equal(serviceHash);
            await expect(await ttmAccountManager.getRegisteredServiceNameByHash(serviceHash)).to.be.equal(serviceName);

            // Try with non registered service hash
            const nonRegisteredServiceName = "nonRegisteredServiceHash";
            const nonRegisteredServiceHash = ethers.keccak256(ethers.toUtf8Bytes(nonRegisteredServiceName));

            await expect(
                ttmAccountManager.getRegisteredServiceHashByName(nonRegisteredServiceName),
            ).to.be.revertedWithCustomError(ttmAccountManager, "ServiceNotRegistered");
            await expect(
                ttmAccountManager.getRegisteredServiceNameByHash(nonRegisteredServiceHash),
            ).to.be.revertedWithCustomError(ttmAccountManager, "ServiceNotRegistered");
        });

        it("should unregister a service correctly", async function () {
            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

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

            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceName, serviceHash);

            // Try with non registered service name
            const nonRegisteredServiceName = "nonRegisteredServiceName";
            await expect(
                ttmAccountManager.connect(signers.otherAccount1).unregisterService(nonRegisteredServiceName),
            ).to.be.revertedWithCustomError(ttmAccountManager, "ServiceNotRegistered");
        });

        it("should revert if the service is already registered", async function () {
            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

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

            // Register the service
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName)).to.be.not
                .reverted;

            // Register the service again
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.be.revertedWithCustomError(ttmAccountManager, "ServiceAlreadyRegistered")
                .withArgs(serviceName);
        });

        it("should revert if the caller does not have the SERVICE_REGISTRY_ADMIN_ROLE", async function () {
            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            const serviceName = "ttm.service.accommodation.v1alpha.AccommodationSearchService";

            // registerService
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName))
                .to.be.revertedWithCustomError(ttmAccountManager, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, SERVICE_REGISTRY_ADMIN_ROLE);

            // unregisterService
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName))
                .to.be.revertedWithCustomError(ttmAccountManager, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, SERVICE_REGISTRY_ADMIN_ROLE);
        });

        it("should return all registered services correctly", async function () {
            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName1 = "ttm.service.accommodation.v1.AccommodationSearchService";
            const serviceHash1 = ethers.keccak256(ethers.toUtf8Bytes(serviceName1));

            const serviceName2 = "ttm.service.accommodation.v2.AccommodationSearchService";
            const serviceHash2 = ethers.keccak256(ethers.toUtf8Bytes(serviceName2));

            const serviceName3 = "ttm.service.accommodation.v3.AccommodationSearchService";
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

            // Check all registered service names
            const registeredServices = await ttmAccountManager.getAllRegisteredServiceNames();
            expect(registeredServices).to.be.deep.equal([serviceName1, serviceName2, serviceName3]);

            // Check all registered service hashes
            const registeredServiceHashes = await ttmAccountManager.getAllRegisteredServiceHashes();
            expect(registeredServiceHashes).to.be.deep.equal([serviceHash1, serviceHash2, serviceHash3]);
        });

        it("should return empty arrays when no services are registered", async function () {
            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            // Check that empty arrays are returned when no services are registered
            const registeredServices = await ttmAccountManager.getAllRegisteredServiceNames();
            expect(registeredServices).to.be.deep.equal([]);

            const registeredServiceHashes = await ttmAccountManager.getAllRegisteredServiceHashes();
            expect(registeredServiceHashes).to.be.deep.equal([]);
        });

        it("should handle service registration and unregistration with proper state updates", async function () {
            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();

            // Grant SERVICE_REGISTRY_ADMIN_ROLE
            await expect(
                ttmAccountManager
                    .connect(signers.managerAdmin)
                    .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address),
            )
                .to.emit(ttmAccountManager, "RoleGranted")
                .withArgs(SERVICE_REGISTRY_ADMIN_ROLE, signers.otherAccount1.address, signers.managerAdmin.address);

            const serviceName1 = "ttm.service.accommodation.v1.AccommodationSearchService";
            const serviceHash1 = ethers.keccak256(ethers.toUtf8Bytes(serviceName1));

            const serviceName2 = "ttm.service.accommodation.v2.AccommodationSearchService";
            const serviceHash2 = ethers.keccak256(ethers.toUtf8Bytes(serviceName2));

            // Register two services
            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName1))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceName1, serviceHash1);

            await expect(ttmAccountManager.connect(signers.otherAccount1).registerService(serviceName2))
                .to.emit(ttmAccountManager, "ServiceRegistered")
                .withArgs(serviceName2, serviceHash2);

            // Verify both services are registered
            let registeredServices = await ttmAccountManager.getAllRegisteredServiceNames();
            let registeredServiceHashes = await ttmAccountManager.getAllRegisteredServiceHashes();
            expect(registeredServices).to.have.length(2);
            expect(registeredServiceHashes).to.have.length(2);
            expect(registeredServices).to.include(serviceName1);
            expect(registeredServices).to.include(serviceName2);

            // Unregister first service
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName1))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceName1, serviceHash1);

            // Verify only second service remains
            registeredServices = await ttmAccountManager.getAllRegisteredServiceNames();
            registeredServiceHashes = await ttmAccountManager.getAllRegisteredServiceHashes();
            expect(registeredServices).to.have.length(1);
            expect(registeredServiceHashes).to.have.length(1);
            expect(registeredServices[0]).to.equal(serviceName2);
            expect(registeredServiceHashes[0]).to.equal(serviceHash2);

            // Verify first service is no longer accessible
            await expect(ttmAccountManager.getRegisteredServiceHashByName(serviceName1)).to.be.revertedWithCustomError(
                ttmAccountManager,
                "ServiceNotRegistered",
            );
            await expect(ttmAccountManager.getRegisteredServiceNameByHash(serviceHash1)).to.be.revertedWithCustomError(
                ttmAccountManager,
                "ServiceNotRegistered",
            );

            // Unregister second service
            await expect(ttmAccountManager.connect(signers.otherAccount1).unregisterService(serviceName2))
                .to.emit(ttmAccountManager, "ServiceUnregistered")
                .withArgs(serviceName2, serviceHash2);

            // Verify all services are unregistered
            registeredServices = await ttmAccountManager.getAllRegisteredServiceNames();
            registeredServiceHashes = await ttmAccountManager.getAllRegisteredServiceHashes();
            expect(registeredServices).to.be.deep.equal([]);
            expect(registeredServiceHashes).to.be.deep.equal([]);
        });
    });
});
