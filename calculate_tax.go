// The calculate-tax package provides tools to calculate the tax to be paid on profits or losses from operations in the financial stock market.
package main

import (
	"bufio"
	"calculate-tax/models"
	"calculate-tax/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var stop bool

var currentAmountOfPurchases int64
var weightedAverageSale float64
var gain float64
var loss float64 // prejuízo

func main() {

	utils.Logger.Info("Starting Application..")

	var operationsOutput []models.OperationOutput
	scn := bufio.NewScanner(os.Stdin)

	for {
		utils.Logger.Info("Inform Operations:")

		for scn.Scan() {
			input := scn.Text()
			if len(input) == 0 {
				stop = true
				break
			}

			operations, err := scanOperations(input)
			if err != nil {
				stop = true
				break
			}

			operationsOutput = ProcessOperations(operations)

		}

		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		if stop {
			utils.Logger.Info("fim")
			result := Result(operationsOutput)
			fmt.Println(result)
			break
		}
	}
}

// scanOperations returns a list of operations
func scanOperations(input string) ([]models.OperationInput, error) {

	var operations []models.OperationInput

	err := json.Unmarshal([]byte(input), &operations)
	if err != nil {
		return nil, errors.New("json unmarshal")
	}

	validOperationType := validateOperationType(operations)
	if !validOperationType {
		return nil, errors.New("type operation is valid")
	}

	return operations, nil
}

// validateOperationType returns true for valid operation types or false for invalid operation types
func validateOperationType(operations []models.OperationInput) bool {

	for _, operation := range operations {
		_, typeOperationIsValid := models.ParseString(operation.Operation)
		if !typeOperationIsValid {
			utils.Logger.Error("type operation is valid: %t. Details: %s", typeOperationIsValid)
			return false
		}
	}

	return true
}

// ProcessOperations returns a list of fees to be paid on shares bought or sold
func ProcessOperations(operations []models.OperationInput) []models.OperationOutput {
	var operationsOutput []models.OperationOutput

	for _, operation := range operations {
		typeOperation, _ := models.ParseString(operation.Operation)
		if typeOperation == models.Buy {
			buy(operation)
			operationOutput := models.OperationOutput{
				Tax: 0.0,
			}
			operationsOutput = append(operationsOutput, operationOutput)
		} else {
			tax := sell(operation)
			oo := models.OperationOutput{
				Tax: tax,
			}
			operationsOutput = append(operationsOutput, oo)
		}
	}

	return operationsOutput
}

// buy add the number of shares purchased
func buy(operation models.OperationInput) {
	currentAmountOfPurchases += operation.Quantity
	calculateWeightedAverage(operation)
}

// sell returns a list of fees to be paid on shares sold
func sell(operation models.OperationInput) float64 {

	currentAmountOfPurchases -= operation.Quantity

	isLoss := isLoss(operation)

	l := 0.0
	g := 0.0
	if isLoss {
		l = float64(operation.Quantity) * weightedAverageSale
		loss += l
	} else {
		g = float64(operation.Quantity) * weightedAverageSale
		gain += g
	}

	if isLoss && l < gain {
		gain -= l
		loss -= l
	}

	if isLoss {
		return 0.0
	}

	if operation.UnitCost*float64(operation.Quantity) <= 20000 {
		return 0.0
	}

	return (g * 20) / 100
}

// calculateWeightedAverage calculates the average purchase price of shares
func calculateWeightedAverage(operation models.OperationInput) {
	if weightedAverageSale == 0 {
		weightedAverageSale = operation.UnitCost
		return
	}

	weightedAverageSale += ((float64(currentAmountOfPurchases) * weightedAverageSale) + (float64(operation.Quantity) * operation.UnitCost)) / (float64(currentAmountOfPurchases) + float64(operation.Quantity))
}

// isLoss check if there was damage
func isLoss(operation models.OperationInput) bool {
	typeOperation, _ := models.ParseString(operation.Operation)
	return typeOperation == models.Sell && operation.UnitCost < weightedAverageSale
}

// ProcessOutput format the output of the program
func Result(operationsOutput []models.OperationOutput) string {
	var operationsOutputJson []models.OperationOutputJson

	for _, operation := range operationsOutput {
		ret := models.OperationOutputJson{
			Tax: fmt.Sprintf("%.2f", operation.Tax),
		}
		operationsOutputJson = append(operationsOutputJson, ret)
	}

	j, err := json.Marshal(operationsOutputJson)
	if err != nil {
		utils.Logger.Error("Error: %s", err.Error())
	}

	return fmt.Sprintf("%s", string(j))
}
