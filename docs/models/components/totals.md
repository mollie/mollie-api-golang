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

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `PendingBalance`                                                            | [*components.PendingBalance](../../models/components/pendingbalance.md)     | :heavy_minus_sign:                                                          | The pending balance. Only available if grouping is `status-balances`.       |
| `AvailableBalance`                                                          | [*components.AvailableBalance](../../models/components/availablebalance.md) | :heavy_minus_sign:                                                          | The available balance. Only available if grouping is `status-balances`.     |
| `Open`                                                                      | [*components.Open](../../models/components/open.md)                         | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Close`                                                                     | [*components.Close](../../models/components/close.md)                       | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Payments`                                                                  | [*components.Payments](../../models/components/payments.md)                 | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Refunds`                                                                   | [*components.Refunds](../../models/components/refunds.md)                   | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Chargebacks`                                                               | [*components.Chargebacks](../../models/components/chargebacks.md)           | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Capital`                                                                   | [*components.Capital](../../models/components/capital.md)                   | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Transfers`                                                                 | [*components.Transfers](../../models/components/transfers.md)               | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `FeePrepayments`                                                            | [*components.FeePrepayments](../../models/components/feeprepayments.md)     | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Corrections`                                                               | [*components.Corrections](../../models/components/corrections.md)           | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |
| `Topups`                                                                    | [*components.Topups](../../models/components/topups.md)                     | :heavy_minus_sign:                                                          | Only available on `transaction-categories` grouping.                        |