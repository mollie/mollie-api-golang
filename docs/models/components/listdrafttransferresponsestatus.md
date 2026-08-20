# ListDraftTransferResponseStatus

The status of the draft transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListDraftTransferResponseStatusPendingReview

// Open enum: custom values can be created with a direct type cast
custom := components.ListDraftTransferResponseStatus("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `ListDraftTransferResponseStatusPendingReview` | pending-review                                 |
| `ListDraftTransferResponseStatusApproved`      | approved                                       |
| `ListDraftTransferResponseStatusDeclined`      | declined                                       |