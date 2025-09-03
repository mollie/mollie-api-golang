# CreatePaymentLocaleRequest

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CreatePaymentLocaleRequestEnUs` | en_US                            |
| `CreatePaymentLocaleRequestEnGb` | en_GB                            |
| `CreatePaymentLocaleRequestNlNl` | nl_NL                            |
| `CreatePaymentLocaleRequestNlBe` | nl_BE                            |
| `CreatePaymentLocaleRequestDeDe` | de_DE                            |
| `CreatePaymentLocaleRequestDeAt` | de_AT                            |
| `CreatePaymentLocaleRequestDeCh` | de_CH                            |
| `CreatePaymentLocaleRequestFrFr` | fr_FR                            |
| `CreatePaymentLocaleRequestFrBe` | fr_BE                            |
| `CreatePaymentLocaleRequestEsEs` | es_ES                            |
| `CreatePaymentLocaleRequestCaEs` | ca_ES                            |
| `CreatePaymentLocaleRequestPtPt` | pt_PT                            |
| `CreatePaymentLocaleRequestItIt` | it_IT                            |
| `CreatePaymentLocaleRequestNbNo` | nb_NO                            |
| `CreatePaymentLocaleRequestSvSe` | sv_SE                            |
| `CreatePaymentLocaleRequestFiFi` | fi_FI                            |
| `CreatePaymentLocaleRequestDaDk` | da_DK                            |
| `CreatePaymentLocaleRequestIsIs` | is_IS                            |
| `CreatePaymentLocaleRequestHuHu` | hu_HU                            |
| `CreatePaymentLocaleRequestPlPl` | pl_PL                            |
| `CreatePaymentLocaleRequestLvLv` | lv_LV                            |
| `CreatePaymentLocaleRequestLtLt` | lt_LT                            |