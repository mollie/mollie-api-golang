# MethodStatus

The payment method's activation status for this profile.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.MethodStatusActivated

// Open enum: custom values can be created with a direct type cast
custom := components.MethodStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `MethodStatusActivated`       | activated                     |
| `MethodStatusPendingBoarding` | pending-boarding              |
| `MethodStatusPendingReview`   | pending-review                |
| `MethodStatusPendingExternal` | pending-external              |
| `MethodStatusRejected`        | rejected                      |