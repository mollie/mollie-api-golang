# ListRefundResponseStatus

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListRefundResponseStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.ListRefundResponseStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ListRefundResponseStatusQueued`     | queued                               |
| `ListRefundResponseStatusPending`    | pending                              |
| `ListRefundResponseStatusProcessing` | processing                           |
| `ListRefundResponseStatusRefunded`   | refunded                             |
| `ListRefundResponseStatusFailed`     | failed                               |
| `ListRefundResponseStatusCanceled`   | canceled                             |