# GetMandateLinks

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Self`                                                                                     | [operations.GetMandateSelf](../../models/operations/getmandateself.md)                     | :heavy_check_mark:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |
| `Customer`                                                                                 | [operations.GetMandateCustomer](../../models/operations/getmandatecustomer.md)             | :heavy_check_mark:                                                                         | The API resource URL of the [customer](get-customer) that this mandate belongs to.         |
| `Documentation`                                                                            | [operations.GetMandateDocumentation](../../models/operations/getmandatedocumentation.md)   | :heavy_check_mark:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |