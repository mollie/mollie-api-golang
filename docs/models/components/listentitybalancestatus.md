# ListEntityBalanceStatus

The status of the balance.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListEntityBalanceStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntityBalanceStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ListEntityBalanceStatusActive`   | active                            |
| `ListEntityBalanceStatusInactive` | inactive                          |