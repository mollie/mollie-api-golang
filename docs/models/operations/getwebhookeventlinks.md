# GetWebhookEventLinks

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `Self`                                                                                             | [operations.GetWebhookEventSelf](../../models/operations/getwebhookeventself.md)                   | :heavy_check_mark:                                                                                 | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field.         |
| `Documentation`                                                                                    | [operations.GetWebhookEventDocumentation](../../models/operations/getwebhookeventdocumentation.md) | :heavy_check_mark:                                                                                 | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field.         |
| `Entity`                                                                                           | [*operations.LinksEntity](../../models/operations/linksentity.md)                                  | :heavy_minus_sign:                                                                                 | The API resource URL of the entity that this event belongs to.                                     |