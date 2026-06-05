package parser

import "github.com/tidwall/gjson"

func Parser(data string) any {
	result := gjson.Get(data, "name.first")
	return result
}
