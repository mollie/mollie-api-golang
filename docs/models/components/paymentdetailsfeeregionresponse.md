# PaymentDetailsFeeRegionResponse

The applicable card fee region.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsFeeRegionResponseAmericanExpress

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsFeeRegionResponse("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `PaymentDetailsFeeRegionResponseAmericanExpress`  | american-express                                  |
| `PaymentDetailsFeeRegionResponseAmexIntraEea`     | amex-intra-eea                                    |
| `PaymentDetailsFeeRegionResponseCarteBancaire`    | carte-bancaire                                    |
| `PaymentDetailsFeeRegionResponseIntraEu`          | intra-eu                                          |
| `PaymentDetailsFeeRegionResponseIntraEuCorporate` | intra-eu-corporate                                |
| `PaymentDetailsFeeRegionResponseDomestic`         | domestic                                          |
| `PaymentDetailsFeeRegionResponseMaestro`          | maestro                                           |
| `PaymentDetailsFeeRegionResponseOther`            | other                                             |
| `PaymentDetailsFeeRegionResponseInter`            | inter                                             |
| `PaymentDetailsFeeRegionResponseIntraEea`         | intra_eea                                         |