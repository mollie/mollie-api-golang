# BalanceTransferDestinationType

The default destination of automatic scheduled transfers. Currently only `bank-account` is supported.

* `bank-account` — Transfer the balance amount to an external bank account

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceTransferDestinationTypeBankAccount

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceTransferDestinationType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `BalanceTransferDestinationTypeBankAccount` | bank-account                                |