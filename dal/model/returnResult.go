package model

type ReturnResult[T interface{}] struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	value   T      `json:"value"`
	list    []T    `json:"list"`
}

func Ok(message string) ReturnResult[interface{}] {
	return ReturnResult[interface{}]{Code: 1000, Message: message}
}
func OkWithData[T interface{}](message string, value T) ReturnResult[T] {
	return ReturnResult[T]{Code: 1000, Message: message, value: value}
}
func OkWithList[T interface{}](message string, list []T) ReturnResult[T] {
	return ReturnResult[T]{Code: 1000, Message: message, list: list}
}
