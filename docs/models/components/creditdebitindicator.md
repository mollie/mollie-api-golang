# CreditDebitIndicator

Indicates whether the transfer is a credit or debit transaction from the perspective of the account holder.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.CreditDebitIndicatorCredit

// Open enum: custom values can be created with a direct type cast
custom := components.CreditDebitIndicator("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CreditDebitIndicatorCredit` | credit                       |
| `CreditDebitIndicatorDebit`  | debit                        |