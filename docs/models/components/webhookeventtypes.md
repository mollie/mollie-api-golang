# WebhookEventTypes

The list of events to enable for this webhook. You may specify `'*'` to add all events, except those
that require explicit selection.


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `WebhookEventTypesPaymentLinkPaid`           | payment-link.paid                            |
| `WebhookEventTypesBalanceTransactionCreated` | balance-transaction.created                  |
| `WebhookEventTypesSalesInvoiceCreated`       | sales-invoice.created                        |
| `WebhookEventTypesSalesInvoiceIssued`        | sales-invoice.issued                         |
| `WebhookEventTypesSalesInvoiceCanceled`      | sales-invoice.canceled                       |
| `WebhookEventTypesSalesInvoicePaid`          | sales-invoice.paid                           |
| `WebhookEventTypesWildcard`                  | *                                            |