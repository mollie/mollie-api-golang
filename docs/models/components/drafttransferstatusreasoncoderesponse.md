# DraftTransferStatusReasonCodeResponse

A machine-readable code that indicates the reason for the draft transfer's current status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.DraftTransferStatusReasonCodeResponseDeletedByCreator

// Open enum: custom values can be created with a direct type cast
custom := components.DraftTransferStatusReasonCodeResponse("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `DraftTransferStatusReasonCodeResponseDeletedByCreator`    | deleted-by-creator                                         |
| `DraftTransferStatusReasonCodeResponseDeclinedByInitiator` | declined-by-initiator                                      |
| `DraftTransferStatusReasonCodeResponseAccountClosed`       | account-closed                                             |