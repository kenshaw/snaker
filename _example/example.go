// _example/example.go
package main

import (
	"fmt"

	"github.com/kenshaw/snaker"
)

func main() {
	fmt.Println("Change CamelCase -> snake_case:", snaker.CamelToSnake("AnIdentifier"))
	fmt.Println("Change CamelCase -> snake_case (2):", snaker.CamelToSnake("XMLHTTPACL"))
	fmt.Println("Change snake_case -> CamelCase:", snaker.SnakeToCamel("an_identifier"))
	fmt.Println("Force CamelCase:", snaker.ForceCamelIdentifier("APoorly_named_httpMethod"))
	fmt.Println("Force lower camelCase:", snaker.ForceLowerCamelIdentifier("APoorly_named_httpMethod"))
	fmt.Println("Force lower camelCase (2):", snaker.ForceLowerCamelIdentifier("XmlHttpACL"))
	fmt.Println("Change snake_case identifier -> CamelCase:", snaker.SnakeToCamelIdentifier("__2__xml___thing---"))
}
