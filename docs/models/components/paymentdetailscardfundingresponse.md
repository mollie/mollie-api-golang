# PaymentDetailsCardFundingResponse

The card type.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsCardFundingResponseDebit

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsCardFundingResponse("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `PaymentDetailsCardFundingResponseDebit`         | debit                                            |
| `PaymentDetailsCardFundingResponseCredit`        | credit                                           |
| `PaymentDetailsCardFundingResponsePrepaid`       | prepaid                                          |
| `PaymentDetailsCardFundingResponseDeferredDebit` | deferred-debit                                   |