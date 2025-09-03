# PaymentLinkLinks

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Self`                                                                                                     | [operations.PaymentLinkSelf](../../models/operations/paymentlinkself.md)                                   | :heavy_check_mark:                                                                                         | In v2 endpoints, URLs are commonly represented as objects with an `href` and `type` field.                 |
| `PaymentLink`                                                                                              | [operations.ListPaymentLinksLinksPaymentLink](../../models/operations/listpaymentlinkslinkspaymentlink.md) | :heavy_check_mark:                                                                                         | The URL your customer should visit to make the payment. This is where you should redirect the customer to. |