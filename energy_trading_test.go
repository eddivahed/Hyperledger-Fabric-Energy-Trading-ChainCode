package main

import (
    "encoding/json"
    "testing"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-chaincode-go/shimtest"
    "github.com/stretchr/testify/assert"
    "strconv"
)

func TestCreateEnergyRequest(t *testing.T) {
    cc := new(EnergyTradingChaincode)
    stub := shimtest.NewMockStub("TestStub", cc)

    id := "request1"
    consumerId := "consumer1"
    energyAmount := 100
    timestamp := "2023-06-08T10:00:00Z"

    request := EnergyRequest{
        ID:           id,
        ConsumerID:   consumerId,
        EnergyAmount: energyAmount,
        Timestamp:    timestamp,
    }
    requestBytes, _ := json.Marshal(request)

    res := stub.MockInvoke("tx1", [][]byte{[]byte("CreateEnergyRequest"), []byte(id), []byte(consumerId), []byte(strconv.Itoa(energyAmount)), []byte(timestamp)})
    assert.Equal(t, int32(shim.OK), res.Status)

    state, err := stub.GetState(id)
    assert.NoError(t, err)
    assert.Equal(t, requestBytes, state)
}

func TestMatchRequestWithOffer(t *testing.T) {
    cc := new(EnergyTradingChaincode)
    stub := shimtest.NewMockStub("TestStub", cc)

    requestId := "request1"
    offerId := "offer1"
    energyAmount := 100

    request := EnergyRequest{
        ID:           requestId,
        ConsumerID:   "consumer1",
        EnergyAmount: energyAmount,
        Timestamp:    "2023-06-08T10:00:00Z",
    }
    requestBytes, _ := json.Marshal(request)

    offer := EnergyOffer{
        ID:           offerId,
        ProducerID:   "producer1",
        EnergyAmount: energyAmount,
        Timestamp:    "2023-06-08T10:00:00Z",
    }
    offerBytes, _ := json.Marshal(offer)

    stub.MockTransactionStart("tx1")
    stub.PutState(requestId, requestBytes)
    stub.PutState(offerId, offerBytes)
    stub.MockTransactionEnd("tx1")

    res := stub.MockInvoke("tx2", [][]byte{[]byte("MatchRequestWithOffer"), []byte(requestId), []byte(offerId)})
    assert.Equal(t, int32(shim.OK), res.Status)

    var transaction Transaction
    err := json.Unmarshal(res.Payload, &transaction)
    assert.NoError(t, err)
    assert.Equal(t, requestId, transaction.RequestID)
    assert.Equal(t, offerId, transaction.OfferID)
    assert.Equal(t, energyAmount, transaction.EnergyAmount)
}