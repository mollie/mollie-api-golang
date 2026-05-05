# ListSettlementPaymentResponseSettlementAmount

The amount settled to your account for this payment, converted to the currency your account is settled in.

Amounts not settled by Mollie are not reflected here (e.g. PayPal or gift cards). If no amount is settled by
Mollie, this field is omitted from the response.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Currency`                                                          | `string`                                                            | :heavy_check_mark:                                                  | A three-character ISO 4217 currency code.                           | EUR                                                                 |
| `Value`                                                             | `string`                                                            | :heavy_check_mark:                                                  | A string containing an exact monetary amount in the given currency. | 10.00                                                               |