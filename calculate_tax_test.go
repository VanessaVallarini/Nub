package main

import (
	"calculate-tax/models"
	"testing"
)

func TestCase1(t *testing.T) {

	var operations []models.OperationInput

	operation1 := models.OperationInput{
		Operation: "buy",

		UnitCost: 10.00,

		Quantity: 100,
	}

	operation2 := models.OperationInput{
		Operation: "sell",

		UnitCost: 15.00,

		Quantity: 50,
	}

	operation3 := models.OperationInput{
		Operation: "sell",

		UnitCost: 15.00,

		Quantity: 50,
	}

	operations = append(operations, operation1)
	operations = append(operations, operation2)
	operations = append(operations, operation3)

	output := ProcessOperations(operations)
	formattedOutput := Result(output)
	t.Log(formattedOutput)
	// Output: [{"tax":"0.00"},{"tax":"0.00"},{"tax":"0.00"}]
}

func TestCase2(t *testing.T) {

	var operations []models.OperationInput

	operation1 := models.OperationInput{
		Operation: "buy",

		UnitCost: 10.00,

		Quantity: 10000,
	}

	operation2 := models.OperationInput{
		Operation: "sell",

		UnitCost: 20.00,

		Quantity: 5000,
	}

	operation3 := models.OperationInput{
		Operation: "sell",

		UnitCost: 5.00,

		Quantity: 5000,
	}

	operations = append(operations, operation1)
	operations = append(operations, operation2)
	operations = append(operations, operation3)

	output := ProcessOperations(operations)
	formattedOutput := Result(output)
	t.Log(formattedOutput)
	// Output: [{"tax":"0.00"},{"tax":"10000.00"},{"tax":"0.00"}]
}
