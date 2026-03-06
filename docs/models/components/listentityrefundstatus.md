# ListEntityRefundStatus

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListEntityRefundStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntityRefundStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ListEntityRefundStatusQueued`     | queued                             |
| `ListEntityRefundStatusPending`    | pending                            |
| `ListEntityRefundStatusProcessing` | processing                         |
| `ListEntityRefundStatusRefunded`   | refunded                           |
| `ListEntityRefundStatusFailed`     | failed                             |
| `ListEntityRefundStatusCanceled`   | canceled                           |