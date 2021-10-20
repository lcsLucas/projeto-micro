package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
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

func MakeRateLimit(e endpoint.Endpoint, duration time.Duration, limit_request int) endpoint.Endpoint {
	return ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(duration), limit_request))(e)
}
