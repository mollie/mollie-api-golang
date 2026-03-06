# WebhookStatus

The subscription's current status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.WebhookStatusEnabled

// Open enum: custom values can be created with a direct type cast
custom := components.WebhookStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `WebhookStatusEnabled`  | enabled                 |
| `WebhookStatusBlocked`  | blocked                 |
| `WebhookStatusDisabled` | disabled                |
| `WebhookStatusDeleted`  | deleted                 |