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

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `PaymentDetailsFeeRegionResponseAmericanExpress`                  | american-express                                                  |
| `PaymentDetailsFeeRegionResponseAmexIntraEea`                     | amex-intra-eea                                                    |
| `PaymentDetailsFeeRegionResponseCarteBancaire`                    | carte-bancaire                                                    |
| `PaymentDetailsFeeRegionResponseIntraEu`                          | intra-eu                                                          |
| `PaymentDetailsFeeRegionResponseIntraEuCorporate`                 | intra-eu-corporate                                                |
| `PaymentDetailsFeeRegionResponseDomestic`                         | domestic                                                          |
| `PaymentDetailsFeeRegionResponseMaestro`                          | maestro                                                           |
| `PaymentDetailsFeeRegionResponseMastercardCreditBusinessDomestic` | mastercard-credit-business-domestic                               |
| `PaymentDetailsFeeRegionResponseMastercardCreditConsumerDomestic` | mastercard-credit-consumer-domestic                               |
| `PaymentDetailsFeeRegionResponseMastercardCreditConsumerIntraEea` | mastercard-credit-consumer-intra-eea                              |
| `PaymentDetailsFeeRegionResponseMastercardDebitBusinessDomestic`  | mastercard-debit-business-domestic                                |
| `PaymentDetailsFeeRegionResponseMastercardDebitBusinessIntraEea`  | mastercard-debit-business-intra-eea                               |
| `PaymentDetailsFeeRegionResponseMastercardDebitConsumerDomestic`  | mastercard-debit-consumer-domestic                                |
| `PaymentDetailsFeeRegionResponseMastercardDebitConsumerIntraEea`  | mastercard-debit-consumer-intra-eea                               |
| `PaymentDetailsFeeRegionResponseOther`                            | other                                                             |
| `PaymentDetailsFeeRegionResponseInter`                            | inter                                                             |
| `PaymentDetailsFeeRegionResponseIntraEea`                         | intra_eea                                                         |
| `PaymentDetailsFeeRegionResponseVisaCreditBusinessDomestic`       | visa-credit-business-domestic                                     |
| `PaymentDetailsFeeRegionResponseVisaCreditConsumerDomestic`       | visa-credit-consumer-domestic                                     |
| `PaymentDetailsFeeRegionResponseVisaCreditConsumerIntraEea`       | visa-credit-consumer-intra-eea                                    |
| `PaymentDetailsFeeRegionResponseVisaDebitBusinessDomestic`        | visa-debit-business-domestic                                      |
| `PaymentDetailsFeeRegionResponseVisaDebitBusinessIntraEea`        | visa-debit-business-intra-eea                                     |
| `PaymentDetailsFeeRegionResponseVisaDebitConsumerDomestic`        | visa-debit-consumer-domestic                                      |