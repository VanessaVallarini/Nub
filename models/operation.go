package models

type OperationInput struct {
	Operation string  `json:"operation" validate:"required"`
	UnitCost  float64 `json:"unit-cost" validate:"required"`
	Quantity  int64   `json:"quantity" validate:"required"`
}

type OperationOutput struct {
	Tax float64 `json:"tax"`
}

type OperationOutputJson struct {
	Tax string `json:"tax"`
}
