# PaymentLineTypeResponse

The type of product purchased. For example, a physical or a digital product.

The `tip` payment line type is not available when creating a payment.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentLineTypeResponsePhysical

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentLineTypeResponse("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `PaymentLineTypeResponsePhysical`    | physical                             |
| `PaymentLineTypeResponseDigital`     | digital                              |
| `PaymentLineTypeResponseShippingFee` | shipping_fee                         |
| `PaymentLineTypeResponseDiscount`    | discount                             |
| `PaymentLineTypeResponseStoreCredit` | store_credit                         |
| `PaymentLineTypeResponseGiftCard`    | gift_card                            |
| `PaymentLineTypeResponseSurcharge`   | surcharge                            |
| `PaymentLineTypeResponseTip`         | tip                                  |