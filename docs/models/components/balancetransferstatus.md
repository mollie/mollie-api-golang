# BalanceTransferStatus

The status of the transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceTransferStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceTransferStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `BalanceTransferStatusCreated`   | created                          |
| `BalanceTransferStatusFailed`    | failed                           |
| `BalanceTransferStatusSucceeded` | succeeded                        |