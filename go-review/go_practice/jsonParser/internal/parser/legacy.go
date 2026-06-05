package parser

import (
	"encoding/json"
	"fmt"
)

func Legacy(data string) interface{} {
	var parsedData interface{}
	err := json.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", parsedData)

	return parsedData
}
