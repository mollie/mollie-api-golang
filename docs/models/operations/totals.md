# Totals

Totals are grouped according to the chosen grouping rule. The example response should give a good idea of what a
typical grouping looks like.

If grouping `status-balances` is chosen, the main grouping is as follows:

* `pendingBalance` containing an `open`, `pending`, `movedToAvailable`, and `close` sub-group
* `availableBalance` containing an `open`, `movedFromPending`, `immediatelyAvailable`, and `close` sub-group

If grouping `transaction-categories` is chosen, the main grouping is as follows:

* `open` and `close` groups, each containing a `pending` and `available` sub-group
* Transaction type groups such as `payments`, `refunds`, `chargebacks`, `capital`, `transfers`, `fee-prepayments`, `corrections`, `topups`
each containing a `pending`, `movedToAvailable`, and
`immediatelyAvailable` sub-group

Each sub-group typically has:

* An `amount` object containing the group's total amount
* A `count` integer if relevant (for example, counting the number of refunds)
* A `subtotals` array containing more sub-group objects if applicable


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `PendingBalance`                                                                                  | [*operations.PendingBalance](../../models/operations/pendingbalance.md)                           | :heavy_minus_sign:                                                                                | The pending balance. Only available if grouping is `status-balances`.                             |
| `AvailableBalance`                                                                                | [*operations.AvailableBalance](../../models/operations/availablebalance.md)                       | :heavy_minus_sign:                                                                                | The available balance. Only available if grouping is `status-balances`.                           |
| `Open`                                                                                            | [*operations.Open](../../models/operations/open.md)                                               | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Close`                                                                                           | [*operations.Close](../../models/operations/close.md)                                             | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Payments`                                                                                        | [*operations.GetBalanceReportPayments](../../models/operations/getbalancereportpayments.md)       | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Refunds`                                                                                         | [*operations.GetBalanceReportRefunds](../../models/operations/getbalancereportrefunds.md)         | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Chargebacks`                                                                                     | [*operations.GetBalanceReportChargebacks](../../models/operations/getbalancereportchargebacks.md) | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Capital`                                                                                         | [*operations.Capital](../../models/operations/capital.md)                                         | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Transfers`                                                                                       | [*operations.Transfers](../../models/operations/transfers.md)                                     | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `FeePrepayments`                                                                                  | [*operations.FeePrepayments](../../models/operations/feeprepayments.md)                           | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Corrections`                                                                                     | [*operations.Corrections](../../models/operations/corrections.md)                                 | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |
| `Topups`                                                                                          | [*operations.Topups](../../models/operations/topups.md)                                           | :heavy_minus_sign:                                                                                | Only available on `transaction-categories` grouping.                                              |