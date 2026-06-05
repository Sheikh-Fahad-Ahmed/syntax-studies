package main

import (
	"fmt"

	"github.com/Sheikh-Fahad-Ahmed/jsonParser/internal/data"
	"github.com/Sheikh-Fahad-Ahmed/jsonParser/internal/parser"
)

func main() {

	jsonData := data.Data
	data := `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	fmt.Println("Legacy")
	res := parser.Legacy(data)
	fmt.Printf("inside main: %+v", res)

	fmt.Printf("\n\nParser")
	res = parser.Parser(jsonData)
	fmt.Printf("\nParser result: %+v", "orderid")
}
