# BalanceTransferCategory

The type of the transfer. Different fees may apply to different types of transfers.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceTransferCategoryInvoiceCollection
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `BalanceTransferCategoryInvoiceCollection`    | invoice_collection                            |
| `BalanceTransferCategoryPurchase`             | purchase                                      |
| `BalanceTransferCategoryChargeback`           | chargeback                                    |
| `BalanceTransferCategoryRefund`               | refund                                        |
| `BalanceTransferCategoryServicePenalty`       | service_penalty                               |
| `BalanceTransferCategoryDiscountCompensation` | discount_compensation                         |
| `BalanceTransferCategoryManualCorrection`     | manual_correction                             |
| `BalanceTransferCategoryOtherFee`             | other_fee                                     |