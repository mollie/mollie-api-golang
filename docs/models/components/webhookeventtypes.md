# WebhookEventTypes

The list of events to enable for this webhook. You may specify `'*'` to add all events, except those
that require explicit selection.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.WebhookEventTypesPaymentPaid
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `WebhookEventTypesPaymentPaid`                          | payment.paid                                            |
| `WebhookEventTypesPaymentAuthorized`                    | payment.authorized                                      |
| `WebhookEventTypesPaymentFailed`                        | payment.failed                                          |
| `WebhookEventTypesPaymentCanceled`                      | payment.canceled                                        |
| `WebhookEventTypesPaymentExpired`                       | payment.expired                                         |
| `WebhookEventTypesPaymentPending`                       | payment.pending                                         |
| `WebhookEventTypesRefundQueued`                         | refund.queued                                           |
| `WebhookEventTypesRefundPending`                        | refund.pending                                          |
| `WebhookEventTypesRefundProcessing`                     | refund.processing                                       |
| `WebhookEventTypesRefundRefunded`                       | refund.refunded                                         |
| `WebhookEventTypesRefundFailed`                         | refund.failed                                           |
| `WebhookEventTypesRefundCanceled`                       | refund.canceled                                         |
| `WebhookEventTypesPaymentLinkPaid`                      | payment-link.paid                                       |
| `WebhookEventTypesBalanceTransactionCreated`            | balance-transaction.created                             |
| `WebhookEventTypesPayoutInitiated`                      | payout.initiated                                        |
| `WebhookEventTypesPayoutProcessingAtBank`               | payout.processing-at-bank                               |
| `WebhookEventTypesPayoutCompleted`                      | payout.completed                                        |
| `WebhookEventTypesPayoutCanceled`                       | payout.canceled                                         |
| `WebhookEventTypesPayoutFailed`                         | payout.failed                                           |
| `WebhookEventTypesSalesInvoiceCreated`                  | sales-invoice.created                                   |
| `WebhookEventTypesSalesInvoiceIssued`                   | sales-invoice.issued                                    |
| `WebhookEventTypesSalesInvoiceCanceled`                 | sales-invoice.canceled                                  |
| `WebhookEventTypesSalesInvoicePaid`                     | sales-invoice.paid                                      |
| `WebhookEventTypesSalesInvoiceEInvoiceFailed`           | sales-invoice.e-invoice-failed                          |
| `WebhookEventTypesSalesInvoiceEInvoiceIssued`           | sales-invoice.e-invoice-issued                          |
| `WebhookEventTypesBusinessAccountTransferRequested`     | business-account-transfer.requested                     |
| `WebhookEventTypesBusinessAccountTransferInitiated`     | business-account-transfer.initiated                     |
| `WebhookEventTypesBusinessAccountTransferPendingReview` | business-account-transfer.pending-review                |
| `WebhookEventTypesBusinessAccountTransferProcessed`     | business-account-transfer.processed                     |
| `WebhookEventTypesBusinessAccountTransferFailed`        | business-account-transfer.failed                        |
| `WebhookEventTypesBusinessAccountTransferBlocked`       | business-account-transfer.blocked                       |
| `WebhookEventTypesBusinessAccountTransferReturned`      | business-account-transfer.returned                      |
| `WebhookEventTypesWildcard`                             | *                                                       |