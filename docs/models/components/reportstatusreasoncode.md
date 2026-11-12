# ReportStatusReasonCode

A machine-readable code for the failure reason.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ReportStatusReasonCodeGenerationFailed

// Open enum: custom values can be created with a direct type cast
custom := components.ReportStatusReasonCode("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `ReportStatusReasonCodeGenerationFailed`       | generation-failed                              |
| `ReportStatusReasonCodeEmptyReport`            | empty-report                                   |
| `ReportStatusReasonCodeReportValidationFailed` | report-validation-failed                       |