# EntityBalanceStatus

The status of the balance.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityBalanceStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.EntityBalanceStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `EntityBalanceStatusActive`   | active                        |
| `EntityBalanceStatusInactive` | inactive                      |