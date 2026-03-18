# ListCaptureResponseLinks

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Self`                                                                                     | [components.URLObj](../../models/components/urlobj.md)                                     | :heavy_check_mark:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |
| `Payment`                                                                                  | [components.URLObj](../../models/components/urlobj.md)                                     | :heavy_check_mark:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |
| `Settlement`                                                                               | [*components.URLNullable](../../models/components/urlnullable.md)                          | :heavy_minus_sign:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |
| `Shipment`                                                                                 | [*components.URLNullable](../../models/components/urlnullable.md)                          | :heavy_minus_sign:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |