# StatusReasonMerchantResponse

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.StatusReasonMerchantResponseMerchantIDNotFound

// Open enum: custom values can be created with a direct type cast
custom := components.StatusReasonMerchantResponse("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `StatusReasonMerchantResponseMerchantIDNotFound`                | merchant_id_not_found                                           |
| `StatusReasonMerchantResponseMerchantAccountClosed`             | merchant_account_closed                                         |
| `StatusReasonMerchantResponseTerminalIDNotFound`                | terminal_id_not_found                                           |
| `StatusReasonMerchantResponseTerminalClosed`                    | terminal_closed                                                 |
| `StatusReasonMerchantResponseInvalidCategoryCode`               | invalid_category_code                                           |
| `StatusReasonMerchantResponseInvalidCurrency`                   | invalid_currency                                                |
| `StatusReasonMerchantResponseMissingCvv2Cvc2`                   | missing_cvv2_cvc2                                               |
| `StatusReasonMerchantResponseCvv2NotAllowed`                    | cvv2_not_allowed                                                |
| `StatusReasonMerchantResponseMerchantNotRegisteredVbv`          | merchant_not_registered_vbv                                     |
| `StatusReasonMerchantResponseMerchantNotRegisteredForAmex`      | merchant_not_registered_for_amex                                |
| `StatusReasonMerchantResponseTransactionNotPermittedAtTerminal` | transaction_not_permitted_at_terminal                           |
| `StatusReasonMerchantResponseAgreementTerminalNotRelated`       | agreement_terminal_not_related                                  |
| `StatusReasonMerchantResponseInvalidProcessorID`                | invalid_processor_id                                            |
| `StatusReasonMerchantResponseInvalidMerchantData`               | invalid_merchant_data                                           |
| `StatusReasonMerchantResponseSubMerchantAccountClosed`          | sub_merchant_account_closed                                     |