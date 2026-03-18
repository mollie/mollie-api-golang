# ListCaptureResponseStatus

The capture's status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListCaptureResponseStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.ListCaptureResponseStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ListCaptureResponseStatusPending`   | pending                              |
| `ListCaptureResponseStatusSucceeded` | succeeded                            |
| `ListCaptureResponseStatusFailed`    | failed                               |