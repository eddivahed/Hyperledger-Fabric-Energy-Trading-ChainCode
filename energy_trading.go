package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type EnergyRequest struct {
	ID           string `json:"id"`
	ConsumerID   string `json:"consumerId"`
	EnergyAmount int    `json:"energyAmount"`
	Timestamp    string `json:"timestamp"`
}

type EnergyOffer struct {
	ID           string `json:"id"`
	ProducerID   string `json:"producerId"`
	EnergyAmount int    `json:"energyAmount"`
	Timestamp    string `json:"timestamp"`
}

type Transaction struct {
	ID           string `json:"id"`
	RequestID    string `json:"requestId"`
	OfferID      string `json:"offerId"`
	EnergyAmount int    `json:"energyAmount"`
	Timestamp    string `json:"timestamp"`
	Status       string `json:"status"`
}
type EnergyTradingChaincode struct {
	contractapi.Contract
}

func (c *EnergyTradingChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
func (c *EnergyTradingChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "CreateEnergyRequest":
		return c.createEnergyRequest(stub, args)
	case "CreateEnergyOffer":
		return c.createEnergyOffer(stub, args)
	case "MatchRequestWithOffer":
		return c.matchRequestWithOffer(stub, args)
	case "ExecuteTransaction":
		return c.executeTransaction(stub, args)
	case "GetTransactionHistory":
		return c.getTransactionHistory(stub, args)
	default:
		return shim.Error("Invalid function name")
	}
}
func (c *EnergyTradingChaincode) createEnergyRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	id := args[0]
	consumerId := args[1]
	energyAmount, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid energy amount")
	}
	timestamp := args[3]

	request := EnergyRequest{
		ID:           id,
		ConsumerID:   consumerId,
		EnergyAmount: energyAmount,
		Timestamp:    timestamp,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return shim.Error("Failed to marshal energy request")
	}

	err = stub.PutState(id, requestBytes)
	if err != nil {
		return shim.Error("Failed to put energy request in state")
	}

	return shim.Success(nil)
}
func (c *EnergyTradingChaincode) createEnergyOffer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	id := args[0]
	producerId := args[1]
	energyAmount, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid energy amount")
	}
	timestamp := args[3]

	offer := EnergyOffer{
		ID:           id,
		ProducerID:   producerId,
		EnergyAmount: energyAmount,
		Timestamp:    timestamp,
	}
	offerBytes, err := json.Marshal(offer)
	if err != nil {
		return shim.Error("Failed to marshal energy offer")
	}

	err = stub.PutState(id, offerBytes)
	if err != nil {
		return shim.Error("Failed to put energy offer in state")
	}

	return shim.Success(nil)
}
func (c *EnergyTradingChaincode) executeTransaction(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	transactionId := args[0]

	// Retrieve the transaction from the ledger
	transactionBytes, err := stub.GetState(transactionId)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to read transaction from world state: %v", err))
	}
	if transactionBytes == nil {
		return shim.Error(fmt.Sprintf("Transaction %s does not exist", transactionId))
	}
	var transaction Transaction
	err = json.Unmarshal(transactionBytes, &transaction)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal transaction: %v", err))
	}

	// Verify that the energy has been delivered and payment has been made
	// (Implement the verification logic based on your specific requirements)

	// Update the transaction status to "completed"
	transaction.Status = "completed"
	updatedTransactionBytes, err := json.Marshal(transaction)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal updated transaction: %v", err))
	}
	err = stub.PutState(transactionId, updatedTransactionBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to put updated transaction in state: %v", err))
	}

	return shim.Success(nil)
}
// func (c *EnergyTradingChaincode) getTransactionHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}

// 	userId := args[0]

// 	// Retrieve all transactions from the ledger
// 	resultsIterator, err := stub.GetStateByRange("", "")
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("Failed to get state by range: %v", err))
// 	}
// 	defer resultsIterator.Close()

// 	var transactions []Transaction
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return shim.Error(fmt.Sprintf("Failed to get next result: %v", err))
// 		}

// 		var transaction Transaction
// 		err = json.Unmarshal(queryResponse.Value, &transaction)
// 		if err != nil {
// 			return shim.Error(fmt.Sprintf("Failed to unmarshal transaction: %v", err))
// 		}

// 		// Retrieve the energy request and offer associated with the transaction
// 		requestBytes, err := stub.GetState(transaction.RequestID)
// 		if err != nil {
// 			return shim.Error(fmt.Sprintf("Failed to get energy request: %v", err))
// 		}
// 		var energyRequest EnergyRequest
// 		err = json.Unmarshal(requestBytes, &energyRequest)
// 		if err != nil {
// 			return shim.Error(fmt.Sprintf("Failed to unmarshal energy request: %v", err))
// 		}

// 		offerBytes, err := stub.GetState(transaction.OfferID)
// 		if err != nil {
// 			return shim.Error(fmt.Sprintf("Failed to get energy offer: %v", err))
// 		}
// 		var energyOffer EnergyOffer
// 		err = json.Unmarshal(offerBytes, &energyOffer)
// 		if err != nil {
// 			return shim.Error(fmt.Sprintf("Failed to unmarshal energy offer: %v", err))
// 		}

// 		// Filter transactions based on the user ID (consumer or producer)
// 		if energyRequest.ConsumerID == userId || energyOffer.ProducerID == userId {
// 			transactions = append(transactions, transaction)
// 		}
// 	}

// 	transactionsBytes, err := json.Marshal(transactions)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("Failed to marshal transactions: %v", err))
// 	}

// 	return shim.Success(transactionsBytes)
// }
func (c *EnergyTradingChaincode) getTransactionHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    userId := args[0]

    // Retrieve all transactions from the ledger
    resultsIterator, err := stub.GetStateByRange("", "")
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to get state by range: %v", err))
    }
    defer resultsIterator.Close()

    var transactions []Transaction
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(fmt.Sprintf("Failed to get next result: %v", err))
        }

        var transaction Transaction
        err = json.Unmarshal(queryResponse.Value, &transaction)
        if err != nil {
            // Skip transactions that cannot be unmarshaled
            fmt.Printf("Warning: Failed to unmarshal transaction: %v\n", err)
            continue
        }

        // Retrieve the energy request and offer associated with the transaction
        requestBytes, err := stub.GetState(transaction.RequestID)
        if err != nil {
            return shim.Error(fmt.Sprintf("Failed to get energy request: %v", err))
        }
        if requestBytes == nil {
            // Skip transactions with missing energy request
            fmt.Printf("Warning: Energy request not found for transaction %s\n", transaction.ID)
            continue
        }
        var energyRequest EnergyRequest
        err = json.Unmarshal(requestBytes, &energyRequest)
        if err != nil {
            // Skip energy requests that cannot be unmarshaled
            fmt.Printf("Warning: Failed to unmarshal energy request for transaction %s: %v\n", transaction.ID, err)
            continue
        }

        offerBytes, err := stub.GetState(transaction.OfferID)
        if err != nil {
            return shim.Error(fmt.Sprintf("Failed to get energy offer: %v", err))
        }
        if offerBytes == nil {
            // Skip transactions with missing energy offer
            fmt.Printf("Warning: Energy offer not found for transaction %s\n", transaction.ID)
            continue
        }
        var energyOffer EnergyOffer
        err = json.Unmarshal(offerBytes, &energyOffer)
        if err != nil {
            // Skip energy offers that cannot be unmarshaled
            fmt.Printf("Warning: Failed to unmarshal energy offer for transaction %s: %v\n", transaction.ID, err)
            continue
        }

        // Filter transactions based on the user ID (consumer or producer)
        if energyRequest.ConsumerID == userId || energyOffer.ProducerID == userId {
            transactions = append(transactions, transaction)
        }
    }

    transactionsBytes, err := json.Marshal(transactions)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to marshal transactions: %v", err))
    }

    return shim.Success(transactionsBytes)
}
func (c *EnergyTradingChaincode) matchRequestWithOffer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestId := args[0]
	offerId := args[1]

	// Retrieve the energy request from the ledger
	requestBytes, err := stub.GetState(requestId)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to read energy request from world state: %v", err))
	}
	if requestBytes == nil {
		return shim.Error(fmt.Sprintf("Energy request %s does not exist", requestId))
	}
	var energyRequest EnergyRequest
	err = json.Unmarshal(requestBytes, &energyRequest)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal energy request: %v", err))
	}

	// Retrieve the energy offer from the ledger
	offerBytes, err := stub.GetState(offerId)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to read energy offer from world state: %v", err))
	}
	if offerBytes == nil {
		return shim.Error(fmt.Sprintf("Energy offer %s does not exist", offerId))
	}
	var energyOffer EnergyOffer
	err = json.Unmarshal(offerBytes, &energyOffer)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal energy offer: %v", err))
	}

	// Check if the energy amounts match
	if energyRequest.EnergyAmount != energyOffer.EnergyAmount {
		return shim.Error("Energy amounts do not match between request and offer")
	}

	// Create a new transaction
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get transaction timestamp: %v", err))
	}
	transaction := Transaction{
		ID:           stub.GetTxID(),
		RequestID:    requestId,
		OfferID:      offerId,
		EnergyAmount: energyRequest.EnergyAmount,
		Timestamp:    timestamp.String(),
	}
	transactionBytes, err := json.Marshal(transaction)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal transaction: %v", err))
	}

	// Store the transaction in the ledger
	err = stub.PutState(transaction.ID, transactionBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to put transaction in state: %v", err))
	}

	// Update the status of the energy request and offer (e.g., mark them as "matched")
	// (Implement the logic to update the status based on your specific requirements)

	return shim.Success(transactionBytes)
}
func main() {
	err := shim.Start(new(EnergyTradingChaincode))
	if err != nil {
		fmt.Printf("Error starting energy trading chaincode: %s", err)
	}
}
