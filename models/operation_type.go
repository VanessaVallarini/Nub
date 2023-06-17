package models

import "strings"

type OperationType int32

const (
	Buy OperationType = iota
	Sell
)

var (
	operationTypeMap = map[string]OperationType{
		"buy":  Buy,
		"sell": Sell,
	}
)

func ParseString(str string) (OperationType, bool) {
	operationType, ok := operationTypeMap[strings.ToLower(str)]
	return operationType, ok
}
