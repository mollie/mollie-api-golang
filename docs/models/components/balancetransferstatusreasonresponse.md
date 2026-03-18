# BalanceTransferStatusReasonResponse

A machine-readable code that indicates the reason for the transfer's status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceTransferStatusReasonResponseRequestCreated

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceTransferStatusReasonResponse("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `BalanceTransferStatusReasonResponseRequestCreated`            | request_created                                                |
| `BalanceTransferStatusReasonResponseSuccess`                   | success                                                        |
| `BalanceTransferStatusReasonResponseSourceNotAllowed`          | source_not_allowed                                             |
| `BalanceTransferStatusReasonResponseDestinationNotAllowed`     | destination_not_allowed                                        |
| `BalanceTransferStatusReasonResponseInsufficientFunds`         | insufficient_funds                                             |
| `BalanceTransferStatusReasonResponseInvalidSourceBalance`      | invalid_source_balance                                         |
| `BalanceTransferStatusReasonResponseInvalidDestinationBalance` | invalid_destination_balance                                    |
| `BalanceTransferStatusReasonResponseTransferRequestExpired`    | transfer_request_expired                                       |
| `BalanceTransferStatusReasonResponseTransferLimitReached`      | transfer_limit_reached                                         |