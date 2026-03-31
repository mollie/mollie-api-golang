# TransferSchemeTypeResponse

The transfer scheme to be used for the transfer. The transfer scheme determines the processing time and method of the transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.TransferSchemeTypeResponseSepaCreditInst

// Open enum: custom values can be created with a direct type cast
custom := components.TransferSchemeTypeResponse("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `TransferSchemeTypeResponseSepaCreditInst` | sepa-credit-inst                           |
| `TransferSchemeTypeResponseSepaCredit`     | sepa-credit                                |