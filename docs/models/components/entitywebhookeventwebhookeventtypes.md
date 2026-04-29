# EntityWebhookEventWebhookEventTypes

The list of events to enable for this webhook. You may specify `'*'` to add all events, except those
that require explicit selection.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityWebhookEventWebhookEventTypesPaymentLinkPaid

// Open enum: custom values can be created with a direct type cast
custom := components.EntityWebhookEventWebhookEventTypes("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `EntityWebhookEventWebhookEventTypesPaymentLinkPaid`                      | payment-link.paid                                                         |
| `EntityWebhookEventWebhookEventTypesBalanceTransactionCreated`            | balance-transaction.created                                               |
| `EntityWebhookEventWebhookEventTypesSalesInvoiceCreated`                  | sales-invoice.created                                                     |
| `EntityWebhookEventWebhookEventTypesSalesInvoiceIssued`                   | sales-invoice.issued                                                      |
| `EntityWebhookEventWebhookEventTypesSalesInvoiceCanceled`                 | sales-invoice.canceled                                                    |
| `EntityWebhookEventWebhookEventTypesSalesInvoicePaid`                     | sales-invoice.paid                                                        |
| `EntityWebhookEventWebhookEventTypesBusinessAccountTransferRequested`     | business-account-transfer.requested                                       |
| `EntityWebhookEventWebhookEventTypesBusinessAccountTransferInitiated`     | business-account-transfer.initiated                                       |
| `EntityWebhookEventWebhookEventTypesBusinessAccountTransferPendingReview` | business-account-transfer.pending-review                                  |
| `EntityWebhookEventWebhookEventTypesBusinessAccountTransferProcessed`     | business-account-transfer.processed                                       |
| `EntityWebhookEventWebhookEventTypesBusinessAccountTransferFailed`        | business-account-transfer.failed                                          |
| `EntityWebhookEventWebhookEventTypesBusinessAccountTransferBlocked`       | business-account-transfer.blocked                                         |
| `EntityWebhookEventWebhookEventTypesBusinessAccountTransferReturned`      | business-account-transfer.returned                                        |
| `EntityWebhookEventWebhookEventTypesWildcard`                             | *                                                                         |