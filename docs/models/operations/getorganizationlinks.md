# GetOrganizationLinks

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Self`                                                                                              | [*operations.GetOrganizationSelf](../../models/operations/getorganizationself.md)                   | :heavy_minus_sign:                                                                                  | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field.          |
| `Dashboard`                                                                                         | [*operations.GetOrganizationDashboard](../../models/operations/getorganizationdashboard.md)         | :heavy_minus_sign:                                                                                  | Direct link to the organization's Mollie dashboard.                                                 |
| `Documentation`                                                                                     | [*operations.GetOrganizationDocumentation](../../models/operations/getorganizationdocumentation.md) | :heavy_minus_sign:                                                                                  | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field.          |