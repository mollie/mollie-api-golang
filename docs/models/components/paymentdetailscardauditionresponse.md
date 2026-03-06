# PaymentDetailsCardAuditionResponse

The card's target audience, if known.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsCardAuditionResponseConsumer

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsCardAuditionResponse("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `PaymentDetailsCardAuditionResponseConsumer` | consumer                                     |
| `PaymentDetailsCardAuditionResponseBusiness` | business                                     |