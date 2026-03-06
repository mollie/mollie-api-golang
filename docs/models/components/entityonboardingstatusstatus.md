# EntityOnboardingStatusStatus

The current status of the organization's onboarding process.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityOnboardingStatusStatusNeedsData

// Open enum: custom values can be created with a direct type cast
custom := components.EntityOnboardingStatusStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `EntityOnboardingStatusStatusNeedsData` | needs-data                              |
| `EntityOnboardingStatusStatusInReview`  | in-review                               |
| `EntityOnboardingStatusStatusCompleted` | completed                               |