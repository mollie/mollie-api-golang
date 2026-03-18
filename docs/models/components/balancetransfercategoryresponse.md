# BalanceTransferCategoryResponse

The type of the transfer. Different fees may apply to different types of transfers.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceTransferCategoryResponseInvoiceCollection

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceTransferCategoryResponse("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `BalanceTransferCategoryResponseInvoiceCollection`    | invoice_collection                                    |
| `BalanceTransferCategoryResponsePurchase`             | purchase                                              |
| `BalanceTransferCategoryResponseChargeback`           | chargeback                                            |
| `BalanceTransferCategoryResponseRefund`               | refund                                                |
| `BalanceTransferCategoryResponseServicePenalty`       | service_penalty                                       |
| `BalanceTransferCategoryResponseDiscountCompensation` | discount_compensation                                 |
| `BalanceTransferCategoryResponseManualCorrection`     | manual_correction                                     |
| `BalanceTransferCategoryResponseOtherFee`             | other_fee                                             |