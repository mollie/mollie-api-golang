# ListRefundsLinks

Links to help navigate through the lists of items. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Self`                                                                                     | [operations.ListRefundsSelf](../../models/operations/listrefundsself.md)                   | :heavy_check_mark:                                                                         | The URL to the current set of items.                                                       |
| `Previous`                                                                                 | [operations.ListRefundsPrevious](../../models/operations/listrefundsprevious.md)           | :heavy_check_mark:                                                                         | The previous set of items, if available.                                                   |
| `Next`                                                                                     | [operations.ListRefundsNext](../../models/operations/listrefundsnext.md)                   | :heavy_check_mark:                                                                         | The next set of items, if available.                                                       |
| `Documentation`                                                                            | [operations.ListRefundsDocumentation](../../models/operations/listrefundsdocumentation.md) | :heavy_check_mark:                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field. |