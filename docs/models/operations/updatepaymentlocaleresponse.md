# UpdatePaymentLocaleResponse

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
| `UpdatePaymentLocaleResponseEnUs` | en_US                             |
| `UpdatePaymentLocaleResponseEnGb` | en_GB                             |
| `UpdatePaymentLocaleResponseNlNl` | nl_NL                             |
| `UpdatePaymentLocaleResponseNlBe` | nl_BE                             |
| `UpdatePaymentLocaleResponseDeDe` | de_DE                             |
| `UpdatePaymentLocaleResponseDeAt` | de_AT                             |
| `UpdatePaymentLocaleResponseDeCh` | de_CH                             |
| `UpdatePaymentLocaleResponseFrFr` | fr_FR                             |
| `UpdatePaymentLocaleResponseFrBe` | fr_BE                             |
| `UpdatePaymentLocaleResponseEsEs` | es_ES                             |
| `UpdatePaymentLocaleResponseCaEs` | ca_ES                             |
| `UpdatePaymentLocaleResponsePtPt` | pt_PT                             |
| `UpdatePaymentLocaleResponseItIt` | it_IT                             |
| `UpdatePaymentLocaleResponseNbNo` | nb_NO                             |
| `UpdatePaymentLocaleResponseSvSe` | sv_SE                             |
| `UpdatePaymentLocaleResponseFiFi` | fi_FI                             |
| `UpdatePaymentLocaleResponseDaDk` | da_DK                             |
| `UpdatePaymentLocaleResponseIsIs` | is_IS                             |
| `UpdatePaymentLocaleResponseHuHu` | hu_HU                             |
| `UpdatePaymentLocaleResponsePlPl` | pl_PL                             |
| `UpdatePaymentLocaleResponseLvLv` | lv_LV                             |
| `UpdatePaymentLocaleResponseLtLt` | lt_LT                             |