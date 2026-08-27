# BalanceTransactionType

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceTransactionTypeAPIPaymentRollingReserveRelease

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceTransactionType("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `BalanceTransactionTypeAPIPaymentRollingReserveRelease`   | api-payment-rolling-reserve-release                       |
| `BalanceTransactionTypeApplicationFee`                    | application-fee                                           |
| `BalanceTransactionTypeBalanceChargeFee`                  | balance-charge-fee                                        |
| `BalanceTransactionTypeBalanceCorrection`                 | balance-correction                                        |
| `BalanceTransactionTypeBalanceReserve`                    | balance-reserve                                           |
| `BalanceTransactionTypeBalanceReserveReturn`              | balance-reserve-return                                    |
| `BalanceTransactionTypeBalanceTopup`                      | balance-topup                                             |
| `BalanceTransactionTypeCanceledTransfer`                  | canceled-transfer                                         |
| `BalanceTransactionTypeCapture`                           | capture                                                   |
| `BalanceTransactionTypeCashCollateralIssuance`            | cash-collateral-issuance                                  |
| `BalanceTransactionTypeCashCollateralRelease`             | cash-collateral-release                                   |
| `BalanceTransactionTypeChargeback`                        | chargeback                                                |
| `BalanceTransactionTypeChargebackCompensation`            | chargeback-compensation                                   |
| `BalanceTransactionTypeChargebackReversal`                | chargeback-reversal                                       |
| `BalanceTransactionTypeFailedPayment`                     | failed-payment                                            |
| `BalanceTransactionTypeFailedPlatformSplitPayment`        | failed-platform-split-payment                             |
| `BalanceTransactionTypeFailedSplitPaymentCompensation`    | failed-split-payment-compensation                         |
| `BalanceTransactionTypeFeePrepayment`                     | fee-prepayment                                            |
| `BalanceTransactionTypeHeldRollingReserve`                | held-rolling-reserve                                      |
| `BalanceTransactionTypeIncomingTransfer`                  | incoming-transfer                                         |
| `BalanceTransactionTypeInvoiceCompensation`               | invoice-compensation                                      |
| `BalanceTransactionTypeInvoiceRoundingCompensation`       | invoice-rounding-compensation                             |
| `BalanceTransactionTypeLoan`                              | loan                                                      |
| `BalanceTransactionTypeMovement`                          | movement                                                  |
| `BalanceTransactionTypeOutgoingCustomAmountTransfer`      | outgoing-custom-amount-transfer                           |
| `BalanceTransactionTypeOutgoingTransfer`                  | outgoing-transfer                                         |
| `BalanceTransactionTypePayment`                           | payment                                                   |
| `BalanceTransactionTypePaymentFee`                        | payment-fee                                               |
| `BalanceTransactionTypePendingRollingReserve`             | pending-rolling-reserve                                   |
| `BalanceTransactionTypePlatformPaymentChargeback`         | platform-payment-chargeback                               |
| `BalanceTransactionTypePlatformPaymentRefund`             | platform-payment-refund                                   |
| `BalanceTransactionTypePostPaymentSplitPayment`           | post-payment-split-payment                                |
| `BalanceTransactionTypeRefund`                            | refund                                                    |
| `BalanceTransactionTypeRefundCompensation`                | refund-compensation                                       |
| `BalanceTransactionTypeReleasedRollingReserve`            | released-rolling-reserve                                  |
| `BalanceTransactionTypeRepayment`                         | repayment                                                 |
| `BalanceTransactionTypeReturnedPlatformPaymentRefund`     | returned-platform-payment-refund                          |
| `BalanceTransactionTypeReturnedRefund`                    | returned-refund                                           |
| `BalanceTransactionTypeReturnedRefundCompensation`        | returned-refund-compensation                              |
| `BalanceTransactionTypeReturnedTransfer`                  | returned-transfer                                         |
| `BalanceTransactionTypeReversedChargebackCompensation`    | reversed-chargeback-compensation                          |
| `BalanceTransactionTypeReversedPlatformPaymentChargeback` | reversed-platform-payment-chargeback                      |
| `BalanceTransactionTypeRollingReserveHold`                | rolling-reserve-hold                                      |
| `BalanceTransactionTypeRollingReserveRelease`             | rolling-reserve-release                                   |
| `BalanceTransactionTypeSplitPayment`                      | split-payment                                             |
| `BalanceTransactionTypeSplitTransaction`                  | split-transaction                                         |
| `BalanceTransactionTypeToBeReleasedRollingReserve`        | to-be-released-rolling-reserve                            |
| `BalanceTransactionTypeTopup`                             | topup                                                     |