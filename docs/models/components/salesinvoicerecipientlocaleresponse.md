# SalesInvoiceRecipientLocaleResponse

The locale for the recipient, to be used for translations in PDF generation and payment pages.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceRecipientLocaleResponseEnUs

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoiceRecipientLocaleResponse("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `SalesInvoiceRecipientLocaleResponseEnUs` | en_US                                     |
| `SalesInvoiceRecipientLocaleResponseEnGb` | en_GB                                     |
| `SalesInvoiceRecipientLocaleResponseNlNl` | nl_NL                                     |
| `SalesInvoiceRecipientLocaleResponseNlBe` | nl_BE                                     |
| `SalesInvoiceRecipientLocaleResponseDeDe` | de_DE                                     |
| `SalesInvoiceRecipientLocaleResponseDeAt` | de_AT                                     |
| `SalesInvoiceRecipientLocaleResponseDeCh` | de_CH                                     |
| `SalesInvoiceRecipientLocaleResponseFrFr` | fr_FR                                     |
| `SalesInvoiceRecipientLocaleResponseFrBe` | fr_BE                                     |