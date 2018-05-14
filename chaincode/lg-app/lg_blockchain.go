package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Request struct {
	Dealer_names      string `json:"dealer_names"`
	Beneficiary_names string `json:"beneficiary_names"`
	Guarantee_amount  string `json:"guarantee_amount"`
	Document_name     string `json:"document_name"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryLG" {
		return s.queryLG(APIstub, args)
	} else if function == "initLG" {
		return s.initLG(APIstub)
	} else if function == "recordLG" {
		return s.recordLG(APIstub, args)
	} else if function == "queryAllLG" {
		return s.queryAllLG(APIstub)
	}
	// else if function == "changeLGHolder" {
	// 	return s.changeLGHolder(stub, args)
	// }

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryLG(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	lgAsBytes, _ := APIstub.GetState(args[0])
	if lgAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(lgAsBytes)
}

func (s *SmartContract) initLG(APIstub shim.ChaincodeStubInterface) sc.Response {
	lg := []Request{
		Request{Dealer_names: "AAA", Beneficiary_names: "BBB", Guarantee_amount: "100000", Document_name: "No.one"},
		Request{Dealer_names: "CCC", Beneficiary_names: "DDD", Guarantee_amount: "543210", Document_name: "No.two"},
	}

	i := 0
	for i < len(lg) {
		fmt.Println("i is ", i)
		lgAsBytes, _ := json.Marshal(lg[i])
		APIstub.PutState(strconv.Itoa(i+1), lgAsBytes)
		fmt.Println("Added", lg[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) recordLG(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var lg = Request{Dealer_names: args[1], Beneficiary_names: args[2], Guarantee_amount: args[3], Document_name: args[4]}

	lgAsBytes, _ := json.Marshal(lg)
	err := APIstub.PutState(args[0], lgAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record tuna catch: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryAllLG(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllTuna:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
