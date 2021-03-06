/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Company struct {
	Name string `json:"name"`
}

type Type struct {
	Name string `json:"name"`
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ImageUrl  string `json:"imageUrl"`
	Hash      string `json:"hash"`
}

type Card struct {
	UserID    string `json:"userID"`
	CompanyID string `json:"companyID"`
	Name      string `json:"name"`
}

type CardItem struct {
	card          string `json:"card"`
	Key           string `json:"key"`
	Value         string `json:"value"`
	AditionalData string `json:"aditionalData"`
	Date          string `json:"date"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	//queryAllClinics

	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryPerson" {
		return s.queryCar(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createCar" {
		return s.createCar(APIstub, args)
	} else if function == "queryPersons" {
		return s.queryAllUsers(APIstub)
	} else if function == "changeCarOwner" {
		return s.changeCarOwner(APIstub, args)
	} else if function == "queryAllClinics" {
		return s.queryAllClinics(APIstub)
	} else if function == "queryAllResearches" {
		return s.queryAllResearches(APIstub)
	} else if function == "queryResearche" {
		return s.subscribe(APIstub, args)
	} else if function == "getAllSubscribers" {
		return s.getAllSubscribers(APIstub, args)
	} else if function == "queryCardItemByCARDID" {
		return s.queryCardItemByCardID(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name. 1")
}

func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

func (t *SmartContract) queryCardItemByCardID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//   0
	// "bob"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cardID := args[0] //strings.ToUpper(args[0])

	//	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"carditem\",\"card\":\"%s\"}}", cardID)
	//queryString := fmt.Sprintf("{\"selector\":%s}", cardID)

	queryResults, err := getStateByPartialCompositeKey(APIstub, "carditem~card", []string{cardID})

	//	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	users := []User{
		User{FirstName: "Pavel", LastName: "Pantyukhov", ImageUrl: "https://pp.userapi.com/c638918/v638918847/3d1d9/s_auB5cvB6M.jpg", Hash: "3u891738291hdiawhduiawdhiuawd"},
		User{FirstName: "Pavel", LastName: "Pantyukhov", ImageUrl: "https://pp.userapi.com/c638918/v638918847/3d1d9/s_auB5cvB6M.jpg", Hash: "3u891738291hdiawhduiawdhiuawd"},
		User{FirstName: "Maksim", LastName: "Kuznetsov", ImageUrl: "https://pp.userapi.com/c307310/v307310903/602d/Gyr1qLrB23Q.jpg", Hash: "2903821390218390jdioawjdiowajdoiaw"},
		User{FirstName: "Yakov", LastName: "Kanner", ImageUrl: "", Hash: "1892737128djwaiodjiawodjwoi"},
		User{FirstName: "Vitaliy", LastName: "Melnik", ImageUrl: "", Hash: "1231jdlawmdklawjdklawnmdlkwandjakwn"},
		User{FirstName: "Vladimir", LastName: "Ivanov", ImageUrl: "", Hash: "mcjkz7873827381hdaw"},
	}

	i := 0
	for i < len(users) {
		fmt.Println("i is ", i)
		userAsBytes, _ := json.Marshal(users[i])
		APIstub.PutState("USER"+strconv.Itoa(i), userAsBytes)
		fmt.Println("Added", users[i])
		i = i + 1
	}

	// Add clinic

	companies := []Company{
		Company{Name: "НИИ онкологии им. Н.Н. Петрова"},
		Company{Name: "Университетская Клиника"},
		Company{Name: "Городской Клинический Онкологический Диспансер"},
		Company{Name: "Медлайн-Сервис на Октябрьском поле"},
		Company{Name: "Он клиник на Новом Арбате"},
		Company{Name: "Центр эндохирургии и литотрипсии (ЦЭЛТ)"},
		Company{Name: "Клиника Столица на Ленинском, 90"},
		Company{Name: "Клиника Столица на Арбате"},
		Company{Name: "Европейский медицинский центр на ул. Щепкина"},
		Company{Name: "ЭлЭн"},
		Company{Name: "Ортодонт комплекс"},
		Company{Name: "Simpladent на Дмитровской"},
		Company{Name: "Simpladent на Пролетарской"},
		Company{Name: "Перинатальный медицинский центр Мать и Дитя"},
		Company{Name: "Медлайн-Сервис на Полежаевской"},
		Company{Name: "Медлайн-Сервис на Сходненской"},
		Company{Name: "Медлайн-Сервис на Октябрьском поле"},
		Company{Name: "Медлайн-Сервис на ВДНХ"},
	}

	j := 0
	for j < len(companies) {
		fmt.Println("j is ", j)
		asBytes, _ := json.Marshal(companies[j])
		APIstub.PutState("COMPANY"+strconv.Itoa(j), asBytes)
		fmt.Println("Added", companies[j])
		j = j + 1
	}

	for j := 0; j < 20; j++ {
		card := Card{UserID: "USER0", CompanyID: "COMPANY0", Name: "Карточка"}
		asBytes, _ := json.Marshal(card)

		fmt.Println("CARD", card)
		key := "CARD" + strconv.Itoa(j)
		APIstub.PutState(key, asBytes)

		for k := 0; k < 20; k++ {
			cardItem := CardItem{card: key, Key: "Принятие таблетки 1", Value: "1", AditionalData: "Заметка врача", Date: "2017.06.18"}

			asBytes, _ := json.Marshal(cardItem)
			fmt.Println("CARDITEM", cardItem)
			cardItemKey := "CARDITEM" + strconv.Itoa(j)
			APIstub.PutState(cardItemKey, asBytes)

			indexName := "carditem~card"
			colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{cardItem.card, cardItemKey})
			if err != nil {
				return shim.Error(err.Error())
			}
			//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
			//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
			value := []byte{0x00}
			APIstub.PutState(colorNameIndexKey, value)

		}
	}

	// researchs := []Research{
	// 	Research{Name: "Исследование 1", Status: "Active", DateFrom: "", DateTo: ""},
	// }

	// for j := 0; j < 100; j++ {
	// 	researchs = append(researchs, Research{Name: "Исследование 1", Status: "Active", DateFrom: "", DateTo: ""})
	// }
	// j = 0
	// for j < len(researchs) {
	// 	fmt.Println("j is ", j)
	// 	asBytes, _ := json.Marshal(researchs[j])
	// 	APIstub.PutState("RESEARCH"+strconv.Itoa(j), asBytes)
	// 	fmt.Println("Added", researchs[j])
	// 	j = j + 1
	// }

	return shim.Success(nil)
}

func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var user = User{FirstName: args[1], LastName: args[2], ImageUrl: args[3], Hash: args[4]}

	userAsBytes, _ := json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllClinics(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CLINIC0"
	endKey := "CLINIC20"

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
		// Add a comma before array members, suppress it for the first array member
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

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllResearches(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "RESEARCH0"
	endKey := "RESEARCH2000000"

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
		// Add a comma before array members, suppress it for the first array member
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

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) subscribe(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// if len(args) <= 1 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 1")
	// }

	// researchAsBytes, _ := APIstub.GetState(args[0])
	// research := Research{}
	// json.Unmarshal(researchAsBytes, &research)

	// userAsBytes, _ := APIstub.GetState(args[1])
	// user := User{}
	// json.Unmarshal(userAsBytes, &user)

	// asBytes, _ := json.Marshal(ResearchUser{UserID: args[1], ResearchID: args[0]})

	// key := "RESEARCHUSER" + strconv.Itoa(rand.Intn(10000000))
	// //TODO SOOOOOBAD
	// APIstub.PutState(key, asBytes)

	// itemAsBytes, _ := APIstub.GetState(key)
	// return shim.Success(itemAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getAllSubscribers(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	startKey := "RESEARCHUSER0"
	endKey := "RESEARCHUSER2000000000"

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
		// Add a comma before array members, suppress it for the first array member
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

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllUsers(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "USER0"
	endKey := "USER999"

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
		// Add a comma before array members, suppress it for the first array member
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

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//carAsBytes, _ := APIstub.GetState(args[0])
	//car := User{}
	//
	//json.Unmarshal(carAsBytes, &car)
	//car.Owner = args[1]
	//
	//carAsBytes, _ = json.Marshal(car)
	//APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
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

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getStateByPartialCompositeKey(APIstub shim.ChaincodeStubInterface, objectType string, keys []string) ([]byte, error) {

	fmt.Printf("- getQueryResultForobjectType objectType:\n%s\n", objectType)
	resultsIterator, err := APIstub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	var i int
	bArrayMemberAlreadyWritten := false

	for i = 0; resultsIterator.HasNext(); i++ {
		// Note that we don't get the value (2nd return variable), we'll just get the marble name from the composite key
		responseRange, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		// get the color and name from color~name composite key
		_, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		key := compositeKeyParts[1]

		fmt.Printf("KEY: %s", key)

		bytes, _ := APIstub.GetState(key)

		//	cardItem := CardItem{}
		//		json.Unmarshal(bytes, &cardItem)

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is

		//	b, _ := json.Marshal(cardItem)
		buffer.WriteString(string(bytes))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}
