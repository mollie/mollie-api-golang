# ReportTypeResponse

The type of report to generate. Determines which other fields the request needs, and which fields show up
on the resulting report.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ReportTypeResponseTransactionMargin

// Open enum: custom values can be created with a direct type cast
custom := components.ReportTypeResponse("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ReportTypeResponseTransactionMargin` | transaction-margin                    |