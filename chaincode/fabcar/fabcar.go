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
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ImageUrl  string `json:"imageUrl"`
	Hash      string `json:"hash"`
}

type Clinic struct {
	Name string `json:"name"`
}

type Research struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTrom"`
	User     []User `json:"users"`
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

	clinics := []Clinic{
		Clinic{Name: "НИИ онкологии им. Н.Н. Петрова"},
		Clinic{Name: "Университетская Клиника"},
		Clinic{Name: "Городской Клинический Онкологический Диспансер"},
		Clinic{Name: "Медлайн-Сервис на Октябрьском поле"},
		Clinic{Name: "Он клиник на Новом Арбате"},
		Clinic{Name: "Центр эндохирургии и литотрипсии (ЦЭЛТ)"},
		Clinic{Name: "Клиника Столица на Ленинском, 90"},
		Clinic{Name: "Клиника Столица на Арбате"},
		Clinic{Name: "Европейский медицинский центр на ул. Щепкина"},
		Clinic{Name: "ЭлЭн"},
		Clinic{Name: "Ортодонт комплекс"},
		Clinic{Name: "Simpladent на Дмитровской"},
		Clinic{Name: "Simpladent на Пролетарской"},
		Clinic{Name: "Перинатальный медицинский центр Мать и Дитя"},
		Clinic{Name: "Медлайн-Сервис на Полежаевской"},
		Clinic{Name: "Медлайн-Сервис на Сходненской"},
		Clinic{Name: "Медлайн-Сервис на Октябрьском поле"},
		Clinic{Name: "Медлайн-Сервис на ВДНХ"},
	}

	// for i := 0; i < 10; i++ {
	// 	append(clinics, Clinic{Name: "Simpladent на Пролетарской"})
	// }

	j := 0
	for j < len(clinics) {
		fmt.Println("j is ", j)
		clinicAsBytes, _ := json.Marshal(clinics[j])
		APIstub.PutState("CLINIC"+strconv.Itoa(j), clinicAsBytes)
		fmt.Println("Added", clinics[j])
		j = j + 1
	}

	researchs := []Research{
		Research{Name: "Исследование 1", Status: "Active", DateFrom: "", DateTo: "", User: []User{}},
	}

	for j := 0; j < 100; j++ {
		researchs = append(researchs, Research{Name: "Исследование 1", Status: "Active", DateFrom: "", DateTo: "", User: []User{}})
	}
	j = 0
	for j < len(researchs) {
		fmt.Println("j is ", j)
		asBytes, _ := json.Marshal(researchs[j])
		APIstub.PutState("RESEARCH"+strconv.Itoa(j), asBytes)
		fmt.Println("Added", researchs[j])
		j = j + 1
	}

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

	if len(args) <= 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	researchAsBytes, _ := APIstub.GetState(args[0])
	research := Research{}
	json.Unmarshal(researchAsBytes, &research)

	userAsBytes, _ := APIstub.GetState(args[1])
	user := User{}
	json.Unmarshal(userAsBytes, &user)

	research.User = append(research.User, user)
	researchAsBytes, _ = json.Marshal(research)

	APIstub.PutState(args[0], researchAsBytes)
	return shim.Success(researchAsBytes)
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
