# AccountStatus

The status of the business account.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.AccountStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.AccountStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AccountStatusActive`  | active                 |
| `AccountStatusBlocked` | blocked                |
| `AccountStatusClosed`  | closed                 |