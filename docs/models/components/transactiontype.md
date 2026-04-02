# TransactionType

Indicates what kind of transaction this is.

We may introduce new transaction types as we expand the product. We recommend building your integration
to handle unexpected values gracefully, so nothing breaks when that happens. 

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.TransactionTypeCardPayment

// Open enum: custom values can be created with a direct type cast
custom := components.TransactionType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `TransactionTypeCardPayment`       | card-payment                       |
| `TransactionTypeBankTransfer`      | bank-transfer                      |
| `TransactionTypePspTransfer`       | psp-transfer                       |
| `TransactionTypeInternalTransfer`  | internal-transfer                  |
| `TransactionTypeIdealPayment`      | ideal-payment                      |
| `TransactionTypeFee`               | fee                                |
| `TransactionTypeCorrection`        | correction                         |
| `TransactionTypeDirectDebit`       | direct-debit                       |
| `TransactionTypeDirectDebitRefund` | direct-debit-refund                |