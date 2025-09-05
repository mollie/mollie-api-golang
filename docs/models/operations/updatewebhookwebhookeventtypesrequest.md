# UpdateWebhookWebhookEventTypesRequest

The list of events to enable for this webhook. You may specify `'*'` to add all events, except those
that require explicit selection. Separate multiple event types with a comma.


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `UpdateWebhookWebhookEventTypesRequestPaymentLinkPaid`           | payment-link.paid                                                |
| `UpdateWebhookWebhookEventTypesRequestBalanceTransactionCreated` | balance-transaction.created                                      |
| `UpdateWebhookWebhookEventTypesRequestSalesInvoiceCreated`       | sales-invoice.created                                            |
| `UpdateWebhookWebhookEventTypesRequestSalesInvoiceIssued`        | sales-invoice.issued                                             |
| `UpdateWebhookWebhookEventTypesRequestSalesInvoiceCanceled`      | sales-invoice.canceled                                           |
| `UpdateWebhookWebhookEventTypesRequestSalesInvoicePaid`          | sales-invoice.paid                                               |
| `UpdateWebhookWebhookEventTypesRequestWildcard`                  | *                                                                |