# TransferStatus

The status of the transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.TransferStatusRequested

// Open enum: custom values can be created with a direct type cast
custom := components.TransferStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `TransferStatusRequested`     | requested                     |
| `TransferStatusInitiated`     | initiated                     |
| `TransferStatusPendingReview` | pending-review                |
| `TransferStatusProcessed`     | processed                     |
| `TransferStatusFailed`        | failed                        |
| `TransferStatusBlocked`       | blocked                       |
| `TransferStatusReturned`      | returned                      |