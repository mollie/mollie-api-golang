# ListMandateResponseScope

An array defining the eligible use cases for the mandate. For creditcard mandates, this field will always be 
present and can contain one or both of the following values:

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListMandateResponseScopeCustomerPresent

// Open enum: custom values can be created with a direct type cast
custom := components.ListMandateResponseScope("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `ListMandateResponseScopeCustomerPresent`    | customer-present                             |
| `ListMandateResponseScopeCustomerNotPresent` | customer-not-present                         |