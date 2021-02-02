package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func InSlice(target interface{}, slice interface{}) bool {
	retorno := false

	for _, number := range slice.([]uint64) {

		if reflect.TypeOf(target) == reflect.TypeOf(number) && target == number {
			retorno = true
		}

	}

	return retorno
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
