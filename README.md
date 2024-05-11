# Hyperledger-Fabric-Energy-Trading-ChainCode

To install and commit the chaincode using the `./network.sh` script in the Fabric test network, you can follow these steps:

1. Navigate to the `test-network` directory in your Fabric samples:
   ```
   cd fabric-samples/test-network
   ```

2. Ensure that the test network is up and running. If not, start the network using the following command:
   ```
   ./network.sh up createChannel
   ```

3. Deploy the chaincode using the `deployCC` subcommand of `./network.sh`. Specify the chaincode name, path, and version as arguments. For example:
   ```
   ./network.sh deployCC -ccn energytrading -ccp ../energy_trading -ccv 1.0
   ```
   - `-ccn`: Chaincode name
   - `-ccp`: Chaincode path (relative to the `test-network` directory)
   - `-ccv`: Chaincode version

4. The `deployCC` subcommand will install the chaincode on the peers of both organizations (Org1 and Org2) and commit the chaincode definition to the channel.

5. Once the chaincode is deployed successfully, you can interact with it using the peer CLI or by developing a client application.

To interact with the chaincode using the peer CLI, you can follow these steps:

1. Open a new terminal and navigate to the `test-network` directory.

2. Set the environment variables for Org1:
   ```
   export CORE_PEER_LOCALMSPID="Org1MSP"
   export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
   export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
   export CORE_PEER_ADDRESS=localhost:7051
   ```

3. Invoke chaincode functions using the `peer chaincode invoke` command. For example, to create an energy request:
   ```
   peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n energytrading --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"CreateEnergyRequest","Args":["request1", "consumer1", "100", "2023-06-09T10:00:00Z"]}'
   ```

4. Query chaincode functions using the `peer chaincode query` command. For example, to get the transaction history for a user:
   ```
   peer chaincode query -C mychannel -n energytrading -c '{"function":"GetTransactionHistory","Args":["consumer1"]}'
   ```

Make sure to replace `energytrading` with the actual chaincode name you used during deployment, and adjust the function names and arguments according to your chaincode implementation.

Note: The commands assume you are using the default channel name `mychannel`. If you are using a different channel name, replace `mychannel` with the actual channel name.

Remember to set the appropriate environment variables and use the correct paths for the TLS root cert files based on your network configuration.

If you encounter any issues or errors during the deployment or invocation process, please let me know, and I'll be happy to help you further.

Certainly! Here's a step-by-step scenario to demonstrate the usage of the energy trading chaincode through the CLI:

1. Create an energy request:
   ```
   peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n energytrading --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"CreateEnergyRequest","Args":["request1", "consumer1", "100", "2023-06-09T10:00:00Z"]}'
   ```

2. Create an energy offer:
   ```
   peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n energytrading --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"CreateEnergyOffer","Args":["offer1", "producer1", "100", "2023-06-09T10:00:00Z"]}'
   ```

3. Match the energy request with the energy offer:
   ```
   peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n energytrading --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"MatchRequestWithOffer","Args":["request1", "offer1"]}'
   ```

4. Execute the transaction:
   ```
   peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n energytrading --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"ExecuteTransaction","Args":["<transaction-id>"]}'
   ```
   Replace `<transaction-id>` with the actual transaction ID returned from the `MatchRequestWithOffer` function.

5. Query the transaction history for a consumer:
   ```
   peer chaincode query -C mychannel -n energytrading -c '{"function":"GetTransactionHistory","Args":["consumer1"]}'
   ```

6. Query the transaction history for a producer:
   ```
   peer chaincode query -C mychannel -n energytrading -c '{"function":"GetTransactionHistory","Args":["producer1"]}'
   ```

These commands assume that you have the necessary environment variables set up and the peer CLI is configured correctly.

Remember to replace `mychannel` with the actual channel name where the chaincode is deployed, and adjust the peer addresses and TLS certificate paths according to your network configuration.

After executing these commands, you should see the corresponding results, such as the created energy requests, energy offers, and transaction history, depending on the function invoked.

If you encounter any issues or errors, please let me know, and I'll be happy to assist you further.
