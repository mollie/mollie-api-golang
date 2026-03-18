# CapabilityRequirementStatus

The status of the requirement depends on its due date.
If no due date is given, the status will be `requested`.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.CapabilityRequirementStatusCurrentlyDue

// Open enum: custom values can be created with a direct type cast
custom := components.CapabilityRequirementStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CapabilityRequirementStatusCurrentlyDue` | currently-due                             |
| `CapabilityRequirementStatusPastDue`      | past-due                                  |
| `CapabilityRequirementStatusRequested`    | requested                                 |