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

var qtdAtualDeCompras int64         //qtdAtualDeCompras
var novamediaponderadavenda float64 //nova-media-ponderada-venda
var lucro float64                   //lucro
var prejuizo float64                // prejuízo

func main() {

	fmt.Println("Starting Application...")

	var operationsOutput []models.OperationOutput
	scn := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Inform Operations:")

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

			for _, operation := range operations {
				typeOperation, _ := models.ParseString(operation.Operation)
				if typeOperation == models.Buy {
					buy(operation)
					o := models.OperationOutput{
						Tax: 0.0,
					}
					operationsOutput = append(operationsOutput, o)
				} else {
					t := sell(operation)
					o := models.OperationOutput{
						Tax: t,
					}
					operationsOutput = append(operationsOutput, o)
				}
			}

		}

		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		if stop {
			utils.Logger.Info("fim")
			fmt.Println(operationsOutput)
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

// validateOperationType Returns true for valid operation types or false for invalid operation types
func validateOperationType(operations []models.OperationInput) bool {

	for _, operation := range operations {
		_, typeOperationIsValid := models.ParseString(operation.Operation)
		if !typeOperationIsValid {
			utils.Logger.Error("type operation is valid: %t", typeOperationIsValid)
			return false
		}
	}

	return true
}

func buy(operation models.OperationInput) {
	qtdAtualDeCompras += operation.Quantity
	calcularMediaPonderada(operation)
}

func sell(operation models.OperationInput) float64 {

	qtdAtualDeCompras -= operation.Quantity

	isPrejuizo := isPrejuizo(operation)

	p := 0.0
	l := 0.0
	if isPrejuizo {
		p = float64(operation.Quantity) * novamediaponderadavenda
		prejuizo += p
	} else {
		l = float64(operation.Quantity) * novamediaponderadavenda
		lucro += l
	}

	if isPrejuizo && p < lucro {
		lucro -= p
		prejuizo -= p
	}

	if isPrejuizo {
		return 0.0
	}

	if operation.UnitCost*float64(operation.Quantity) <= 20000 {
		return 0.0
	}

	return (l * 20) / 100
}

func calcularMediaPonderada(operation models.OperationInput) {
	if novamediaponderadavenda == 0 {
		novamediaponderadavenda = operation.UnitCost
		return
	}

	novamediaponderadavenda += ((float64(qtdAtualDeCompras) * novamediaponderadavenda) + (float64(operation.Quantity) * operation.UnitCost)) / (float64(qtdAtualDeCompras) + float64(operation.Quantity))
}

func isPrejuizo(operation models.OperationInput) bool {
	typeOperation, _ := models.ParseString(operation.Operation)
	return typeOperation == models.Sell && operation.UnitCost < novamediaponderadavenda
}
