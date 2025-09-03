# MandateLinks

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Self`                                                                                     | [operations.MandateSelf](../../models/operations/mandateself.md)                           | :heavy_check_mark:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |
| `Customer`                                                                                 | [operations.ListMandatesCustomer](../../models/operations/listmandatescustomer.md)         | :heavy_check_mark:                                                                         | The API resource URL of the [customer](get-customer) that this mandate belongs to.         |
| `Documentation`                                                                            | [operations.MandateDocumentation](../../models/operations/mandatedocumentation.md)         | :heavy_check_mark:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |