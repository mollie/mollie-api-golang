# RefundResponseStatus

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.RefundResponseStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.RefundResponseStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `RefundResponseStatusQueued`     | queued                           |
| `RefundResponseStatusPending`    | pending                          |
| `RefundResponseStatusProcessing` | processing                       |
| `RefundResponseStatusRefunded`   | refunded                         |
| `RefundResponseStatusFailed`     | failed                           |
| `RefundResponseStatusCanceled`   | canceled                         |