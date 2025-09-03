# CreatePaymentLocaleResponse

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CreatePaymentLocaleResponseEnUs` | en_US                             |
| `CreatePaymentLocaleResponseEnGb` | en_GB                             |
| `CreatePaymentLocaleResponseNlNl` | nl_NL                             |
| `CreatePaymentLocaleResponseNlBe` | nl_BE                             |
| `CreatePaymentLocaleResponseDeDe` | de_DE                             |
| `CreatePaymentLocaleResponseDeAt` | de_AT                             |
| `CreatePaymentLocaleResponseDeCh` | de_CH                             |
| `CreatePaymentLocaleResponseFrFr` | fr_FR                             |
| `CreatePaymentLocaleResponseFrBe` | fr_BE                             |
| `CreatePaymentLocaleResponseEsEs` | es_ES                             |
| `CreatePaymentLocaleResponseCaEs` | ca_ES                             |
| `CreatePaymentLocaleResponsePtPt` | pt_PT                             |
| `CreatePaymentLocaleResponseItIt` | it_IT                             |
| `CreatePaymentLocaleResponseNbNo` | nb_NO                             |
| `CreatePaymentLocaleResponseSvSe` | sv_SE                             |
| `CreatePaymentLocaleResponseFiFi` | fi_FI                             |
| `CreatePaymentLocaleResponseDaDk` | da_DK                             |
| `CreatePaymentLocaleResponseIsIs` | is_IS                             |
| `CreatePaymentLocaleResponseHuHu` | hu_HU                             |
| `CreatePaymentLocaleResponsePlPl` | pl_PL                             |
| `CreatePaymentLocaleResponseLvLv` | lv_LV                             |
| `CreatePaymentLocaleResponseLtLt` | lt_LT                             |