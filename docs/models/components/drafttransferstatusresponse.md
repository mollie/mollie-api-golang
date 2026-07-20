# DraftTransferStatusResponse

The status of the draft transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.DraftTransferStatusResponseAwaitingInitiation

// Open enum: custom values can be created with a direct type cast
custom := components.DraftTransferStatusResponse("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `DraftTransferStatusResponseAwaitingInitiation` | awaiting-initiation                             |
| `DraftTransferStatusResponseInitiated`          | initiated                                       |
| `DraftTransferStatusResponseDeclined`           | declined                                        |