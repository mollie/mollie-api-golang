# EntityRefundResponseStatus

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityRefundResponseStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.EntityRefundResponseStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `EntityRefundResponseStatusQueued`     | queued                                 |
| `EntityRefundResponseStatusPending`    | pending                                |
| `EntityRefundResponseStatusProcessing` | processing                             |
| `EntityRefundResponseStatusRefunded`   | refunded                               |
| `EntityRefundResponseStatusFailed`     | failed                                 |
| `EntityRefundResponseStatusCanceled`   | canceled                               |