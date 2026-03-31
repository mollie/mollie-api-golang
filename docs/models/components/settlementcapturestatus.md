# SettlementCaptureStatus

The capture's status. Settlement captures always have a status of `succeeded`.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SettlementCaptureStatusSucceeded

// Open enum: custom values can be created with a direct type cast
custom := components.SettlementCaptureStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `SettlementCaptureStatusSucceeded` | succeeded                          |