# ListSalesInvoicesLineDiscount

The discount to be applied to the line item.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Type`                                                                                       | [operations.ListSalesInvoicesLineType](../../models/operations/listsalesinvoiceslinetype.md) | :heavy_check_mark:                                                                           | The type of discount.                                                                        | amount                                                                                       |
| `Value`                                                                                      | *string*                                                                                     | :heavy_check_mark:                                                                           | A string containing an exact monetary amount in the given currency, or the percentage.       | 10.00                                                                                        |