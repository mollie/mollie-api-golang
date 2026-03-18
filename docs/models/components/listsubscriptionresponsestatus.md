# ListSubscriptionResponseStatus

The subscription's current status is directly related to the status of the underlying customer or mandate that is
enabling the subscription.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListSubscriptionResponseStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.ListSubscriptionResponseStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ListSubscriptionResponseStatusPending`   | pending                                   |
| `ListSubscriptionResponseStatusActive`    | active                                    |
| `ListSubscriptionResponseStatusCanceled`  | canceled                                  |
| `ListSubscriptionResponseStatusSuspended` | suspended                                 |
| `ListSubscriptionResponseStatusCompleted` | completed                                 |