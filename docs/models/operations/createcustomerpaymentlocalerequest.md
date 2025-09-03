# CreateCustomerPaymentLocaleRequest

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `CreateCustomerPaymentLocaleRequestEnUs` | en_US                                    |
| `CreateCustomerPaymentLocaleRequestEnGb` | en_GB                                    |
| `CreateCustomerPaymentLocaleRequestNlNl` | nl_NL                                    |
| `CreateCustomerPaymentLocaleRequestNlBe` | nl_BE                                    |
| `CreateCustomerPaymentLocaleRequestDeDe` | de_DE                                    |
| `CreateCustomerPaymentLocaleRequestDeAt` | de_AT                                    |
| `CreateCustomerPaymentLocaleRequestDeCh` | de_CH                                    |
| `CreateCustomerPaymentLocaleRequestFrFr` | fr_FR                                    |
| `CreateCustomerPaymentLocaleRequestFrBe` | fr_BE                                    |
| `CreateCustomerPaymentLocaleRequestEsEs` | es_ES                                    |
| `CreateCustomerPaymentLocaleRequestCaEs` | ca_ES                                    |
| `CreateCustomerPaymentLocaleRequestPtPt` | pt_PT                                    |
| `CreateCustomerPaymentLocaleRequestItIt` | it_IT                                    |
| `CreateCustomerPaymentLocaleRequestNbNo` | nb_NO                                    |
| `CreateCustomerPaymentLocaleRequestSvSe` | sv_SE                                    |
| `CreateCustomerPaymentLocaleRequestFiFi` | fi_FI                                    |
| `CreateCustomerPaymentLocaleRequestDaDk` | da_DK                                    |
| `CreateCustomerPaymentLocaleRequestIsIs` | is_IS                                    |
| `CreateCustomerPaymentLocaleRequestHuHu` | hu_HU                                    |
| `CreateCustomerPaymentLocaleRequestPlPl` | pl_PL                                    |
| `CreateCustomerPaymentLocaleRequestLvLv` | lv_LV                                    |
| `CreateCustomerPaymentLocaleRequestLtLt` | lt_LT                                    |