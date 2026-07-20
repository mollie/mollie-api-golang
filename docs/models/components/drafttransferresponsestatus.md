# DraftTransferResponseStatus

The status of the draft transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.DraftTransferResponseStatusAwaitingInitiation

// Open enum: custom values can be created with a direct type cast
custom := components.DraftTransferResponseStatus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `DraftTransferResponseStatusAwaitingInitiation` | awaiting-initiation                             |
| `DraftTransferResponseStatusInitiated`          | initiated                                       |
| `DraftTransferResponseStatusDeclined`           | declined                                        |