# VerificationResultEnum

The result of the Verification of Payee check.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.VerificationResultEnumMatch

// Open enum: custom values can be created with a direct type cast
custom := components.VerificationResultEnum("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `VerificationResultEnumMatch`        | match                                |
| `VerificationResultEnumNoMatch`      | no-match                             |
| `VerificationResultEnumCloseMatch`   | close-match                          |
| `VerificationResultEnumNotAvailable` | not-available                        |