# EntityBalanceTransactionType

The type of transaction, for example `payment` or `refund`. Values include the below examples, although this list
is not definitive.

* Regular payment processing: `payment` `capture` `unauthorized-direct-debit` `failed-payment`
* Refunds and chargebacks: `refund` `returned-refund` `chargeback` `chargeback-reversal`
* Settlements: `outgoing-transfer` `canceled-outgoing-transfer` `returned-transfer`
* Invoicing: `invoice-compensation` `balance-correction`
* Mollie Connect: `application-fee` `split-payment` `platform-payment-refund` `platform-payment-chargeback`


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `EntityBalanceTransactionTypeApplicationFee`                    | application-fee                                                 |
| `EntityBalanceTransactionTypeCapture`                           | capture                                                         |
| `EntityBalanceTransactionTypeChargeback`                        | chargeback                                                      |
| `EntityBalanceTransactionTypeChargebackReversal`                | chargeback-reversal                                             |
| `EntityBalanceTransactionTypeFailedPaymentFee`                  | failed-payment-fee                                              |
| `EntityBalanceTransactionTypeFailedPayment`                     | failed-payment                                                  |
| `EntityBalanceTransactionTypeInvoiceCompensation`               | invoice-compensation                                            |
| `EntityBalanceTransactionTypePayment`                           | payment                                                         |
| `EntityBalanceTransactionTypePaymentFee`                        | payment-fee                                                     |
| `EntityBalanceTransactionTypePaymentCommission`                 | payment-commission                                              |
| `EntityBalanceTransactionTypeRefund`                            | refund                                                          |
| `EntityBalanceTransactionTypeReturnedRefund`                    | returned-refund                                                 |
| `EntityBalanceTransactionTypeReturnedTransfer`                  | returned-transfer                                               |
| `EntityBalanceTransactionTypeSplitPayment`                      | split-payment                                                   |
| `EntityBalanceTransactionTypeOutgoingTransfer`                  | outgoing-transfer                                               |
| `EntityBalanceTransactionTypeCaptureCommission`                 | capture-commission                                              |
| `EntityBalanceTransactionTypeCanceledOutgoingTransfer`          | canceled-outgoing-transfer                                      |
| `EntityBalanceTransactionTypeIncomingTransfer`                  | incoming-transfer                                               |
| `EntityBalanceTransactionTypeAPIPaymentRollingReserveRelease`   | api-payment-rolling-reserve-release                             |
| `EntityBalanceTransactionTypeCaptureRollingReserveRelease`      | capture-rolling-reserve-release                                 |
| `EntityBalanceTransactionTypeReimbursementFee`                  | reimbursement-fee                                               |
| `EntityBalanceTransactionTypeBalanceCorrection`                 | balance-correction                                              |
| `EntityBalanceTransactionTypeUnauthorizedDirectDebit`           | unauthorized-direct-debit                                       |
| `EntityBalanceTransactionTypeBankChargedFailureFee`             | bank-charged-failure-fee                                        |
| `EntityBalanceTransactionTypePlatformPaymentRefund`             | platform-payment-refund                                         |
| `EntityBalanceTransactionTypeRefundCompensation`                | refund-compensation                                             |
| `EntityBalanceTransactionTypeReturnedRefundCompensation`        | returned-refund-compensation                                    |
| `EntityBalanceTransactionTypeReturnedPlatformPaymentRefund`     | returned-platform-payment-refund                                |
| `EntityBalanceTransactionTypePlatformPaymentChargeback`         | platform-payment-chargeback                                     |
| `EntityBalanceTransactionTypeChargebackCompensation`            | chargeback-compensation                                         |
| `EntityBalanceTransactionTypeReversedPlatformPaymentChargeback` | reversed-platform-payment-chargeback                            |
| `EntityBalanceTransactionTypeReversedChargebackCompensation`    | reversed-chargeback-compensation                                |
| `EntityBalanceTransactionTypeFailedSplitPaymentPlatform`        | failed-split-payment-platform                                   |
| `EntityBalanceTransactionTypeFailedSplitPaymentCompensation`    | failed-split-payment-compensation                               |
| `EntityBalanceTransactionTypeCashAdvanceLoan`                   | cash-advance-loan                                               |
| `EntityBalanceTransactionTypePlatformConnectedOrganizationsFee` | platform-connected-organizations-fee                            |
| `EntityBalanceTransactionTypeSplitTransaction`                  | split-transaction                                               |
| `EntityBalanceTransactionTypeManagedFee`                        | managed-fee                                                     |
| `EntityBalanceTransactionTypeReturnedManagedFee`                | returned-managed-fee                                            |
| `EntityBalanceTransactionTypeTopup`                             | topup                                                           |
| `EntityBalanceTransactionTypeBalanceReserve`                    | balance-reserve                                                 |
| `EntityBalanceTransactionTypeBalanceReserveReturn`              | balance-reserve-return                                          |
| `EntityBalanceTransactionTypeMovement`                          | movement                                                        |
| `EntityBalanceTransactionTypePostPaymentSplitPayment`           | post-payment-split-payment                                      |
| `EntityBalanceTransactionTypeCashCollateralIssuance`            | cash-collateral-issuance                                        |
| `EntityBalanceTransactionTypeCashCollateralRelease`             | cash-collateral-release                                         |