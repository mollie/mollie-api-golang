# MandateMethod

Payment method of the mandate.

SEPA Direct Debit and PayPal mandates can be created directly.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.MandateMethodCreditcard
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `MandateMethodCreditcard`  | creditcard                 |
| `MandateMethodDirectdebit` | directdebit                |
| `MandateMethodPaypal`      | paypal                     |