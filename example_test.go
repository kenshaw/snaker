package snaker_test

import (
	"fmt"

	"github.com/kenshaw/snaker"
)

func Example() {
	fmt.Println("Change CamelCase -> snake_case:", snaker.CamelToSnake("AnIdentifier"))
	fmt.Println("Change CamelCase -> snake_case (2):", snaker.CamelToSnake("XMLHTTPACL"))
	fmt.Println("Change snake_case -> CamelCase:", snaker.SnakeToCamel("an_identifier"))
	fmt.Println("Force CamelCase:", snaker.ForceCamelIdentifier("APoorly_named_httpMethod"))
	fmt.Println("Force lower camelCase:", snaker.ForceLowerCamelIdentifier("APoorly_named_httpMethod"))
	fmt.Println("Force lower camelCase (2):", snaker.ForceLowerCamelIdentifier("XmlHttpACL"))
	fmt.Println("Change snake_case identifier -> CamelCase:", snaker.SnakeToCamelIdentifier("__2__xml___thing---"))
	// Output:
	// Change CamelCase -> snake_case: an_identifier
	// Change CamelCase -> snake_case (2): xml_http_acl
	// Change snake_case -> CamelCase: AnIdentifier
	// Force CamelCase: APoorlyNamedHTTPMethod
	// Force lower camelCase: aPoorlyNamedHTTPMethod
	// Force lower camelCase (2): xmlHTTPACL
	// Change snake_case identifier -> CamelCase: XMLThing
}
