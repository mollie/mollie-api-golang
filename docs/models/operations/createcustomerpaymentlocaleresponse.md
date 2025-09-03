# CreateCustomerPaymentLocaleResponse

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CreateCustomerPaymentLocaleResponseEnUs` | en_US                                     |
| `CreateCustomerPaymentLocaleResponseEnGb` | en_GB                                     |
| `CreateCustomerPaymentLocaleResponseNlNl` | nl_NL                                     |
| `CreateCustomerPaymentLocaleResponseNlBe` | nl_BE                                     |
| `CreateCustomerPaymentLocaleResponseDeDe` | de_DE                                     |
| `CreateCustomerPaymentLocaleResponseDeAt` | de_AT                                     |
| `CreateCustomerPaymentLocaleResponseDeCh` | de_CH                                     |
| `CreateCustomerPaymentLocaleResponseFrFr` | fr_FR                                     |
| `CreateCustomerPaymentLocaleResponseFrBe` | fr_BE                                     |
| `CreateCustomerPaymentLocaleResponseEsEs` | es_ES                                     |
| `CreateCustomerPaymentLocaleResponseCaEs` | ca_ES                                     |
| `CreateCustomerPaymentLocaleResponsePtPt` | pt_PT                                     |
| `CreateCustomerPaymentLocaleResponseItIt` | it_IT                                     |
| `CreateCustomerPaymentLocaleResponseNbNo` | nb_NO                                     |
| `CreateCustomerPaymentLocaleResponseSvSe` | sv_SE                                     |
| `CreateCustomerPaymentLocaleResponseFiFi` | fi_FI                                     |
| `CreateCustomerPaymentLocaleResponseDaDk` | da_DK                                     |
| `CreateCustomerPaymentLocaleResponseIsIs` | is_IS                                     |
| `CreateCustomerPaymentLocaleResponseHuHu` | hu_HU                                     |
| `CreateCustomerPaymentLocaleResponsePlPl` | pl_PL                                     |
| `CreateCustomerPaymentLocaleResponseLvLv` | lv_LV                                     |
| `CreateCustomerPaymentLocaleResponseLtLt` | lt_LT                                     |