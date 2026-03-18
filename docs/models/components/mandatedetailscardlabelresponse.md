# MandateDetailsCardLabelResponse

The card's label. Available for card mandates, if the card label could be detected.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.MandateDetailsCardLabelResponseAmericanExpress

// Open enum: custom values can be created with a direct type cast
custom := components.MandateDetailsCardLabelResponse("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `MandateDetailsCardLabelResponseAmericanExpress` | American Express                                 |
| `MandateDetailsCardLabelResponseCartaSi`         | Carta Si                                         |
| `MandateDetailsCardLabelResponseCarteBleue`      | Carte Bleue                                      |
| `MandateDetailsCardLabelResponseDankort`         | Dankort                                          |
| `MandateDetailsCardLabelResponseDinersClub`      | Diners Club                                      |
| `MandateDetailsCardLabelResponseDiscover`        | Discover                                         |
| `MandateDetailsCardLabelResponseJcb`             | JCB                                              |
| `MandateDetailsCardLabelResponseLaser`           | Laser                                            |
| `MandateDetailsCardLabelResponseMaestro`         | Maestro                                          |
| `MandateDetailsCardLabelResponseMastercard`      | Mastercard                                       |
| `MandateDetailsCardLabelResponseUnionpay`        | Unionpay                                         |
| `MandateDetailsCardLabelResponseVisa`            | Visa                                             |