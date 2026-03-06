# PaymentDetailsCardLabelResponse

The card's label, if known.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsCardLabelResponseAmericanExpress

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsCardLabelResponse("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `PaymentDetailsCardLabelResponseAmericanExpress` | American Express                                 |
| `PaymentDetailsCardLabelResponseCartaSi`         | Carta Si                                         |
| `PaymentDetailsCardLabelResponseCarteBleue`      | Carte Bleue                                      |
| `PaymentDetailsCardLabelResponseDankort`         | Dankort                                          |
| `PaymentDetailsCardLabelResponseDinersClub`      | Diners Club                                      |
| `PaymentDetailsCardLabelResponseDiscover`        | Discover                                         |
| `PaymentDetailsCardLabelResponseJcb`             | JCB                                              |
| `PaymentDetailsCardLabelResponseLaser`           | Laser                                            |
| `PaymentDetailsCardLabelResponseMaestro`         | Maestro                                          |
| `PaymentDetailsCardLabelResponseMastercard`      | Mastercard                                       |
| `PaymentDetailsCardLabelResponseUnionpay`        | Unionpay                                         |
| `PaymentDetailsCardLabelResponseVisa`            | Visa                                             |
| `PaymentDetailsCardLabelResponseVpay`            | Vpay                                             |