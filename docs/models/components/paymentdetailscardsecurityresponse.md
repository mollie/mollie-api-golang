# PaymentDetailsCardSecurityResponse

The level of security applied during card processing.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsCardSecurityResponseNormal

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsCardSecurityResponse("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `PaymentDetailsCardSecurityResponseNormal`       | normal                                           |
| `PaymentDetailsCardSecurityResponseThreedsecure` | 3dsecure                                         |