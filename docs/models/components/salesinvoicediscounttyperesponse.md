# SalesInvoiceDiscountTypeResponse

The type of discount.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceDiscountTypeResponseAmount

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoiceDiscountTypeResponse("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `SalesInvoiceDiscountTypeResponseAmount`     | amount                                       |
| `SalesInvoiceDiscountTypeResponsePercentage` | percentage                                   |