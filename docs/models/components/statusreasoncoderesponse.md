# StatusReasonCodeResponse

A machine-readable code indicating the reason for the transfer's terminal status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.StatusReasonCodeResponseInsufficientFunds

// Open enum: custom values can be created with a direct type cast
custom := components.StatusReasonCodeResponse("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `StatusReasonCodeResponseInsufficientFunds` | insufficient-funds                          |
| `StatusReasonCodeResponseRejected`          | rejected                                    |
| `StatusReasonCodeResponseError`             | error                                       |