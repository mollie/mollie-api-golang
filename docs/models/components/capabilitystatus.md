# CapabilityStatus

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.CapabilityStatusUnrequested

// Open enum: custom values can be created with a direct type cast
custom := components.CapabilityStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CapabilityStatusUnrequested` | unrequested                   |
| `CapabilityStatusEnabled`     | enabled                       |
| `CapabilityStatusDisabled`    | disabled                      |
| `CapabilityStatusPending`     | pending                       |