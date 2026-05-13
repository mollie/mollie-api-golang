# PayoutStatus

The status of the payout.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PayoutStatusRequested

// Open enum: custom values can be created with a direct type cast
custom := components.PayoutStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `PayoutStatusRequested`        | requested                      |
| `PayoutStatusInitiated`        | initiated                      |
| `PayoutStatusProcessingAtBank` | processing-at-bank             |
| `PayoutStatusCompleted`        | completed                      |
| `PayoutStatusFailed`           | failed                         |
| `PayoutStatusCanceled`         | canceled                       |