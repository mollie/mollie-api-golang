# PayoutStatusReasonCode

A machine-readable code describing the reason for the payout's current status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PayoutStatusReasonCodeRequested

// Open enum: custom values can be created with a direct type cast
custom := components.PayoutStatusReasonCode("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `PayoutStatusReasonCodeRequested`            | requested                                    |
| `PayoutStatusReasonCodeInitiated`            | initiated                                    |
| `PayoutStatusReasonCodeProcessingAtBank`     | processing_at_bank                           |
| `PayoutStatusReasonCodeCompleted`            | completed                                    |
| `PayoutStatusReasonCodeCanceled`             | canceled                                     |
| `PayoutStatusReasonCodeFailed`               | failed                                       |
| `PayoutStatusReasonCodeInsufficientFunds`    | insufficient_funds                           |
| `PayoutStatusReasonCodeReturned`             | returned                                     |
| `PayoutStatusReasonCodeInvalidRequest`       | invalid_request                              |
| `PayoutStatusReasonCodeOrganizationInactive` | organization_inactive                        |
| `PayoutStatusReasonCodePayoutsBlocked`       | payouts_blocked                              |
| `PayoutStatusReasonCodeBankProcessingFailed` | bank_processing_failed                       |
| `PayoutStatusReasonCodeBalanceNotFound`      | balance_not_found                            |
| `PayoutStatusReasonCodeExpired`              | expired                                      |