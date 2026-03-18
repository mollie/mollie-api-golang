# PaymentDetailsReceiptCardReadMethodResponse

The method by which the card was read by the terminal.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsReceiptCardReadMethodResponseChip

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsReceiptCardReadMethodResponse("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `PaymentDetailsReceiptCardReadMethodResponseChip`                   | chip                                                                |
| `PaymentDetailsReceiptCardReadMethodResponseMagneticStripe`         | magnetic-stripe                                                     |
| `PaymentDetailsReceiptCardReadMethodResponseNearFieldCommunication` | near-field-communication                                            |
| `PaymentDetailsReceiptCardReadMethodResponseContactless`            | contactless                                                         |
| `PaymentDetailsReceiptCardReadMethodResponseMoto`                   | moto                                                                |