# UnmatchedCreditTransferStatus

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.UnmatchedCreditTransferStatusReceived

// Open enum: custom values can be created with a direct type cast
custom := components.UnmatchedCreditTransferStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `UnmatchedCreditTransferStatusReceived` | received                                |
| `UnmatchedCreditTransferStatusMatched`  | matched                                 |
| `UnmatchedCreditTransferStatusReturned` | returned                                |
| `UnmatchedCreditTransferStatusExpired`  | expired                                 |