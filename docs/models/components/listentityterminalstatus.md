# ListEntityTerminalStatus

The status of the terminal.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListEntityTerminalStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntityTerminalStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ListEntityTerminalStatusPending`  | pending                            |
| `ListEntityTerminalStatusActive`   | active                             |
| `ListEntityTerminalStatusInactive` | inactive                           |