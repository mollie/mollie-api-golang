# CancelPaymentLocale

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `CancelPaymentLocaleEnUs` | en_US                     |
| `CancelPaymentLocaleEnGb` | en_GB                     |
| `CancelPaymentLocaleNlNl` | nl_NL                     |
| `CancelPaymentLocaleNlBe` | nl_BE                     |
| `CancelPaymentLocaleDeDe` | de_DE                     |
| `CancelPaymentLocaleDeAt` | de_AT                     |
| `CancelPaymentLocaleDeCh` | de_CH                     |
| `CancelPaymentLocaleFrFr` | fr_FR                     |
| `CancelPaymentLocaleFrBe` | fr_BE                     |
| `CancelPaymentLocaleEsEs` | es_ES                     |
| `CancelPaymentLocaleCaEs` | ca_ES                     |
| `CancelPaymentLocalePtPt` | pt_PT                     |
| `CancelPaymentLocaleItIt` | it_IT                     |
| `CancelPaymentLocaleNbNo` | nb_NO                     |
| `CancelPaymentLocaleSvSe` | sv_SE                     |
| `CancelPaymentLocaleFiFi` | fi_FI                     |
| `CancelPaymentLocaleDaDk` | da_DK                     |
| `CancelPaymentLocaleIsIs` | is_IS                     |
| `CancelPaymentLocaleHuHu` | hu_HU                     |
| `CancelPaymentLocalePlPl` | pl_PL                     |
| `CancelPaymentLocaleLvLv` | lv_LV                     |
| `CancelPaymentLocaleLtLt` | lt_LT                     |