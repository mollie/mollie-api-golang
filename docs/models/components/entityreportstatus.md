# EntityReportStatus

The report's current status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityReportStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.EntityReportStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `EntityReportStatusPending`   | pending                       |
| `EntityReportStatusCompleted` | completed                     |
| `EntityReportStatusFailed`    | failed                        |