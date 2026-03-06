# MandateResponseStatus

The status of the mandate. A status can be `pending` for mandates when the first payment is not yet finalized, or
when we did not received the IBAN yet from the first payment.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.MandateResponseStatusValid

// Open enum: custom values can be created with a direct type cast
custom := components.MandateResponseStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `MandateResponseStatusValid`   | valid                          |
| `MandateResponseStatusPending` | pending                        |
| `MandateResponseStatusInvalid` | invalid                        |