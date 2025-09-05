# CreateWebhookWebhookEventTypesRequest

The list of events to enable for this webhook. You may specify `'*'` to add all events, except those
that require explicit selection. Separate multiple event types with a comma.


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `CreateWebhookWebhookEventTypesRequestPaymentLinkPaid`           | payment-link.paid                                                |
| `CreateWebhookWebhookEventTypesRequestBalanceTransactionCreated` | balance-transaction.created                                      |
| `CreateWebhookWebhookEventTypesRequestSalesInvoiceCreated`       | sales-invoice.created                                            |
| `CreateWebhookWebhookEventTypesRequestSalesInvoiceIssued`        | sales-invoice.issued                                             |
| `CreateWebhookWebhookEventTypesRequestSalesInvoiceCanceled`      | sales-invoice.canceled                                           |
| `CreateWebhookWebhookEventTypesRequestSalesInvoicePaid`          | sales-invoice.paid                                               |
| `CreateWebhookWebhookEventTypesRequestWildcard`                  | *                                                                |