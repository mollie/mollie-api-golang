# SubscriptionResponseStatus

The subscription's current status is directly related to the status of the underlying customer or mandate that is
enabling the subscription.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SubscriptionResponseStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionResponseStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `SubscriptionResponseStatusPending`   | pending                               |
| `SubscriptionResponseStatusActive`    | active                                |
| `SubscriptionResponseStatusCanceled`  | canceled                              |
| `SubscriptionResponseStatusSuspended` | suspended                             |
| `SubscriptionResponseStatusCompleted` | completed                             |