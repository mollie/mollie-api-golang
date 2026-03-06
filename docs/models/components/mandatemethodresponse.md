# MandateMethodResponse

Payment method of the mandate.

SEPA Direct Debit and PayPal mandates can be created directly.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.MandateMethodResponseCreditcard

// Open enum: custom values can be created with a direct type cast
custom := components.MandateMethodResponse("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `MandateMethodResponseCreditcard`  | creditcard                         |
| `MandateMethodResponseDirectdebit` | directdebit                        |
| `MandateMethodResponsePaypal`      | paypal                             |