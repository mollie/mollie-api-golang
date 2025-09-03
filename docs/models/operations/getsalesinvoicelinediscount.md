# GetSalesInvoiceLineDiscount

The discount to be applied to the line item.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Type`                                                                                   | [operations.GetSalesInvoiceLineType](../../models/operations/getsalesinvoicelinetype.md) | :heavy_check_mark:                                                                       | The type of discount.                                                                    | amount                                                                                   |
| `Value`                                                                                  | *string*                                                                                 | :heavy_check_mark:                                                                       | A string containing an exact monetary amount in the given currency, or the percentage.   | 10.00                                                                                    |