# EntityTerminalStatus

The status of the terminal.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityTerminalStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.EntityTerminalStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `EntityTerminalStatusPending`  | pending                        |
| `EntityTerminalStatusActive`   | active                         |
| `EntityTerminalStatusInactive` | inactive                       |