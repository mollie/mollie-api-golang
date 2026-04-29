# WebhookEventTypes

The list of events to enable for this webhook. You may specify `'*'` to add all events, except those
that require explicit selection.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.WebhookEventTypesPaymentLinkPaid
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `WebhookEventTypesPaymentLinkPaid`                      | payment-link.paid                                       |
| `WebhookEventTypesBalanceTransactionCreated`            | balance-transaction.created                             |
| `WebhookEventTypesSalesInvoiceCreated`                  | sales-invoice.created                                   |
| `WebhookEventTypesSalesInvoiceIssued`                   | sales-invoice.issued                                    |
| `WebhookEventTypesSalesInvoiceCanceled`                 | sales-invoice.canceled                                  |
| `WebhookEventTypesSalesInvoicePaid`                     | sales-invoice.paid                                      |
| `WebhookEventTypesBusinessAccountTransferRequested`     | business-account-transfer.requested                     |
| `WebhookEventTypesBusinessAccountTransferInitiated`     | business-account-transfer.initiated                     |
| `WebhookEventTypesBusinessAccountTransferPendingReview` | business-account-transfer.pending-review                |
| `WebhookEventTypesBusinessAccountTransferProcessed`     | business-account-transfer.processed                     |
| `WebhookEventTypesBusinessAccountTransferFailed`        | business-account-transfer.failed                        |
| `WebhookEventTypesBusinessAccountTransferBlocked`       | business-account-transfer.blocked                       |
| `WebhookEventTypesBusinessAccountTransferReturned`      | business-account-transfer.returned                      |
| `WebhookEventTypesWildcard`                             | *                                                       |