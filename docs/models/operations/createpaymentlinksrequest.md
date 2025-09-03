# CreatePaymentLinksRequest

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Self`                                                                                           | [operations.CreatePaymentSelfRequest](../../models/operations/createpaymentselfrequest.md)       | :heavy_check_mark:                                                                               | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field.       |
| `Payment`                                                                                        | [operations.CreatePaymentPaymentRequest](../../models/operations/createpaymentpaymentrequest.md) | :heavy_check_mark:                                                                               | The API resource URL of the [payment](get-payment) that belong to this route.                    |