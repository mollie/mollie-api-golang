# ListBalanceTransactionsType

The type of transaction, for example `payment` or `refund`. Values include the below examples, although this list
is not definitive.

* Regular payment processing: `payment` `capture` `unauthorized-direct-debit` `failed-payment`
* Refunds and chargebacks: `refund` `returned-refund` `chargeback` `chargeback-reversal`
* Settlements: `outgoing-transfer` `canceled-outgoing-transfer` `returned-transfer`
* Invoicing: `invoice-compensation` `balance-correction`
* Mollie Connect: `application-fee` `split-payment` `platform-payment-refund` `platform-payment-chargeback`


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ListBalanceTransactionsTypeApplicationFee`                    | application-fee                                                |
| `ListBalanceTransactionsTypeCapture`                           | capture                                                        |
| `ListBalanceTransactionsTypeChargeback`                        | chargeback                                                     |
| `ListBalanceTransactionsTypeChargebackReversal`                | chargeback-reversal                                            |
| `ListBalanceTransactionsTypeFailedPaymentFee`                  | failed-payment-fee                                             |
| `ListBalanceTransactionsTypeFailedPayment`                     | failed-payment                                                 |
| `ListBalanceTransactionsTypeInvoiceCompensation`               | invoice-compensation                                           |
| `ListBalanceTransactionsTypePayment`                           | payment                                                        |
| `ListBalanceTransactionsTypePaymentFee`                        | payment-fee                                                    |
| `ListBalanceTransactionsTypePaymentCommission`                 | payment-commission                                             |
| `ListBalanceTransactionsTypeRefund`                            | refund                                                         |
| `ListBalanceTransactionsTypeReturnedRefund`                    | returned-refund                                                |
| `ListBalanceTransactionsTypeReturnedTransfer`                  | returned-transfer                                              |
| `ListBalanceTransactionsTypeSplitPayment`                      | split-payment                                                  |
| `ListBalanceTransactionsTypeOutgoingTransfer`                  | outgoing-transfer                                              |
| `ListBalanceTransactionsTypeCaptureCommission`                 | capture-commission                                             |
| `ListBalanceTransactionsTypeCanceledOutgoingTransfer`          | canceled-outgoing-transfer                                     |
| `ListBalanceTransactionsTypeIncomingTransfer`                  | incoming-transfer                                              |
| `ListBalanceTransactionsTypeAPIPaymentRollingReserveRelease`   | api-payment-rolling-reserve-release                            |
| `ListBalanceTransactionsTypeCaptureRollingReserveRelease`      | capture-rolling-reserve-release                                |
| `ListBalanceTransactionsTypeReimbursementFee`                  | reimbursement-fee                                              |
| `ListBalanceTransactionsTypeBalanceCorrection`                 | balance-correction                                             |
| `ListBalanceTransactionsTypeUnauthorizedDirectDebit`           | unauthorized-direct-debit                                      |
| `ListBalanceTransactionsTypeBankChargedFailureFee`             | bank-charged-failure-fee                                       |
| `ListBalanceTransactionsTypePlatformPaymentRefund`             | platform-payment-refund                                        |
| `ListBalanceTransactionsTypeRefundCompensation`                | refund-compensation                                            |
| `ListBalanceTransactionsTypeReturnedRefundCompensation`        | returned-refund-compensation                                   |
| `ListBalanceTransactionsTypeReturnedPlatformPaymentRefund`     | returned-platform-payment-refund                               |
| `ListBalanceTransactionsTypePlatformPaymentChargeback`         | platform-payment-chargeback                                    |
| `ListBalanceTransactionsTypeChargebackCompensation`            | chargeback-compensation                                        |
| `ListBalanceTransactionsTypeReversedPlatformPaymentChargeback` | reversed-platform-payment-chargeback                           |
| `ListBalanceTransactionsTypeReversedChargebackCompensation`    | reversed-chargeback-compensation                               |
| `ListBalanceTransactionsTypeFailedSplitPaymentPlatform`        | failed-split-payment-platform                                  |
| `ListBalanceTransactionsTypeFailedSplitPaymentCompensation`    | failed-split-payment-compensation                              |
| `ListBalanceTransactionsTypeCashAdvanceLoan`                   | cash-advance-loan                                              |
| `ListBalanceTransactionsTypePlatformConnectedOrganizationsFee` | platform-connected-organizations-fee                           |
| `ListBalanceTransactionsTypeSplitTransaction`                  | split-transaction                                              |
| `ListBalanceTransactionsTypeManagedFee`                        | managed-fee                                                    |
| `ListBalanceTransactionsTypeReturnedManagedFee`                | returned-managed-fee                                           |
| `ListBalanceTransactionsTypeTopup`                             | topup                                                          |
| `ListBalanceTransactionsTypeBalanceReserve`                    | balance-reserve                                                |
| `ListBalanceTransactionsTypeBalanceReserveReturn`              | balance-reserve-return                                         |
| `ListBalanceTransactionsTypeMovement`                          | movement                                                       |
| `ListBalanceTransactionsTypePostPaymentSplitPayment`           | post-payment-split-payment                                     |
| `ListBalanceTransactionsTypeCashCollateralIssuance`            | cash-collateral-issuance                                       |
| `ListBalanceTransactionsTypeCashCollateralRelease`             | cash-collateral-release                                        |