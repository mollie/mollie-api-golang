# ListCustomerPaymentsLocale

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
| `ListCustomerPaymentsLocaleEnUs` | en_US                            |
| `ListCustomerPaymentsLocaleEnGb` | en_GB                            |
| `ListCustomerPaymentsLocaleNlNl` | nl_NL                            |
| `ListCustomerPaymentsLocaleNlBe` | nl_BE                            |
| `ListCustomerPaymentsLocaleDeDe` | de_DE                            |
| `ListCustomerPaymentsLocaleDeAt` | de_AT                            |
| `ListCustomerPaymentsLocaleDeCh` | de_CH                            |
| `ListCustomerPaymentsLocaleFrFr` | fr_FR                            |
| `ListCustomerPaymentsLocaleFrBe` | fr_BE                            |
| `ListCustomerPaymentsLocaleEsEs` | es_ES                            |
| `ListCustomerPaymentsLocaleCaEs` | ca_ES                            |
| `ListCustomerPaymentsLocalePtPt` | pt_PT                            |
| `ListCustomerPaymentsLocaleItIt` | it_IT                            |
| `ListCustomerPaymentsLocaleNbNo` | nb_NO                            |
| `ListCustomerPaymentsLocaleSvSe` | sv_SE                            |
| `ListCustomerPaymentsLocaleFiFi` | fi_FI                            |
| `ListCustomerPaymentsLocaleDaDk` | da_DK                            |
| `ListCustomerPaymentsLocaleIsIs` | is_IS                            |
| `ListCustomerPaymentsLocaleHuHu` | hu_HU                            |
| `ListCustomerPaymentsLocalePlPl` | pl_PL                            |
| `ListCustomerPaymentsLocaleLvLv` | lv_LV                            |
| `ListCustomerPaymentsLocaleLtLt` | lt_LT                            |