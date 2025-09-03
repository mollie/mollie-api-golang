# GetInvoiceLinks

An object with several relevant URLs. Every URL object will contain an `href` and a `type` field.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Self`                                                                                    | [*operations.GetInvoiceSelf](../../models/operations/getinvoiceself.md)                   | :heavy_minus_sign:                                                                        | URL to the current invoice resource.                                                      |
| `Pdf`                                                                                     | [*operations.Pdf](../../models/operations/pdf.md)                                         | :heavy_minus_sign:                                                                        | URL to a downloadable PDF of the invoice.                                                 |
| `Documentation`                                                                           | [*operations.GetInvoiceDocumentation](../../models/operations/getinvoicedocumentation.md) | :heavy_minus_sign:                                                                        | URL to the API documentation.                                                             |