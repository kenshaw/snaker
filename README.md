# snaker [![Go Package][gopkg]][gopkg-link]

Package `snaker` provides methods to convert `CamelCase`, `snake_case`, and
`kebab-case` to and from each other. Correctly recognizes common (Go idiomatic)
initialisms (`HTTP`, `XML`, etc) with the ability to override/set recognized
initialisms.

[gopkg]: https://pkg.go.dev/badge/github.com/kenshaw/snaker.svg "Go Package"
[gopkg-link]: https://pkg.go.dev/github.com/kenshaw/snaker

## Example

```go
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
```

See the [package example][example].

[example]: https://pkg.go.dev/github.com/kenshaw/snaker/#Example
