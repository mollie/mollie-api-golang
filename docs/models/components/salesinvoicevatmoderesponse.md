# SalesInvoiceVatModeResponse

The VAT mode to use for VAT calculation. `exclusive` mode means we will apply the relevant VAT on top of the
price. `inclusive` means the prices you are providing to us already contain the VAT you want to apply.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceVatModeResponseExclusive

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoiceVatModeResponse("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `SalesInvoiceVatModeResponseExclusive` | exclusive                              |
| `SalesInvoiceVatModeResponseInclusive` | inclusive                              |