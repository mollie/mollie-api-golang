# DraftTransferResponseStatus

The status of the draft transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.DraftTransferResponseStatusPendingReview

// Open enum: custom values can be created with a direct type cast
custom := components.DraftTransferResponseStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `DraftTransferResponseStatusPendingReview` | pending-review                             |
| `DraftTransferResponseStatusApproved`      | approved                                   |
| `DraftTransferResponseStatusDeclined`      | declined                                   |