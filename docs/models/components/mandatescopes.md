# MandateScopes

An array defining the eligible use cases for the mandate. For creditcard mandates, this field will always be 
present and can contain one or both of the following values:

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.MandateScopesCustomerPresent
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `MandateScopesCustomerPresent`    | customer-present                  |
| `MandateScopesCustomerNotPresent` | customer-not-present              |