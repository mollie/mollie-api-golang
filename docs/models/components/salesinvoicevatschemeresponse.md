# SalesInvoiceVatSchemeResponse

The VAT scheme to create the invoice for. You must be enrolled with One Stop Shop enabled to use it.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceVatSchemeResponseStandard

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoiceVatSchemeResponse("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `SalesInvoiceVatSchemeResponseStandard`    | standard                                   |
| `SalesInvoiceVatSchemeResponseOneStopShop` | one-stop-shop                              |