package main

import (
	"calculate-tax/models"
	"testing"
)

var (
	expectedCase1 = `[{"tax":"0.00"},{"tax":"0.00"},{"tax":"0.00"}]`
	expectedCase2 = `[{"tax":"0.00"},{"tax":"10000.00"},{"tax":"0.00"}]`
	expectedCase3 = `[{"tax":"0.00"},{"tax":"0.00"},{"tax":"1000.00"}]`
	expectedCase4 = `[{"tax":"0.00"},{"tax":"0.00"},{"tax":"0.00"}]`
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

	if formattedOutput != expectedCase1 {
		t.Error("Expected:", expectedCase1, "Got:", formattedOutput)
	}
	t.Log(formattedOutput)
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

	if formattedOutput != expectedCase2 {
		t.Error("Expected:", expectedCase2, "Got:", formattedOutput)
	}
	t.Log(formattedOutput)
}

func TestCase3(t *testing.T) {

	var operations []models.OperationInput

	operation1 := models.OperationInput{
		Operation: "buy",

		UnitCost: 10.00,

		Quantity: 10000,
	}

	operation2 := models.OperationInput{
		Operation: "sell",

		UnitCost: 5.00,

		Quantity: 5000,
	}

	operation3 := models.OperationInput{
		Operation: "sell",

		UnitCost: 20.00,

		Quantity: 3000,
	}

	operations = append(operations, operation1)
	operations = append(operations, operation2)
	operations = append(operations, operation3)

	output := ProcessOperations(operations)
	formattedOutput := Result(output)

	if formattedOutput != expectedCase3 {
		t.Error("Expected:", expectedCase3, "Got:", formattedOutput)
	}
	t.Log(formattedOutput)
}

func TestCase4(t *testing.T) {

	var operations []models.OperationInput

	operation1 := models.OperationInput{
		Operation: "buy",

		UnitCost: 10.00,

		Quantity: 10000,
	}

	operation2 := models.OperationInput{
		Operation: "buy",

		UnitCost: 25.00,

		Quantity: 5000,
	}

	operation3 := models.OperationInput{
		Operation: "sell",

		UnitCost: 15.00,

		Quantity: 10000,
	}

	operations = append(operations, operation1)
	operations = append(operations, operation2)
	operations = append(operations, operation3)

	output := ProcessOperations(operations)
	formattedOutput := Result(output)

	if formattedOutput != expectedCase4 {
		t.Error("Expected:", expectedCase4, "Got:", formattedOutput)
	}
	t.Log(formattedOutput)
}
