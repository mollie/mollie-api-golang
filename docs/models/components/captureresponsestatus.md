# CaptureResponseStatus

The capture's status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.CaptureResponseStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.CaptureResponseStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CaptureResponseStatusPending`   | pending                          |
| `CaptureResponseStatusSucceeded` | succeeded                        |
| `CaptureResponseStatusFailed`    | failed                           |