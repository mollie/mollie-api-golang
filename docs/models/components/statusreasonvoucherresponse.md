# StatusReasonVoucherResponse

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.StatusReasonVoucherResponseServiceFailed

// Open enum: custom values can be created with a direct type cast
custom := components.StatusReasonVoucherResponse("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `StatusReasonVoucherResponseServiceFailed`              | service_failed                                          |
| `StatusReasonVoucherResponseInvalidOperation`           | invalid_operation                                       |
| `StatusReasonVoucherResponseAuthorizationError`         | authorization_error                                     |
| `StatusReasonVoucherResponseLoginFailedWithoutReason`   | login_failed_without_reason                             |
| `StatusReasonVoucherResponseInvalidRetailer`            | invalid_retailer                                        |
| `StatusReasonVoucherResponseReferToCardIssuer`          | refer_to_card_issuer                                    |
| `StatusReasonVoucherResponseCardDoesNotExist`           | card_does_not_exist                                     |
| `StatusReasonVoucherResponseExpiredCard`                | expired_card                                            |
| `StatusReasonVoucherResponseCardIsBlocked`              | card_is_blocked                                         |
| `StatusReasonVoucherResponseInsufficientFunds`          | insufficient_funds                                      |
| `StatusReasonVoucherResponseInvalidCardID`              | invalid_card_id                                         |
| `StatusReasonVoucherResponseCardIsTransferred`          | card_is_transferred                                     |
| `StatusReasonVoucherResponseCardIsNotActive`            | card_is_not_active                                      |
| `StatusReasonVoucherResponseIncorrectPurchaseValue`     | incorrect_purchase_value                                |
| `StatusReasonVoucherResponseCardNotAvailable`           | card_not_available                                      |
| `StatusReasonVoucherResponseWrongCurrency`              | wrong_currency                                          |
| `StatusReasonVoucherResponseLoginFailedUnknownUser`     | login_failed_unknown_user                               |
| `StatusReasonVoucherResponseLoginFailedInvalidPassword` | login_failed_invalid_password                           |
| `StatusReasonVoucherResponseInvalidPin`                 | invalid_pin                                             |
| `StatusReasonVoucherResponseInvalidEanCode`             | invalid_ean_code                                        |