# GetPaymentLocale

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `GetPaymentLocaleEnUs` | en_US                  |
| `GetPaymentLocaleEnGb` | en_GB                  |
| `GetPaymentLocaleNlNl` | nl_NL                  |
| `GetPaymentLocaleNlBe` | nl_BE                  |
| `GetPaymentLocaleDeDe` | de_DE                  |
| `GetPaymentLocaleDeAt` | de_AT                  |
| `GetPaymentLocaleDeCh` | de_CH                  |
| `GetPaymentLocaleFrFr` | fr_FR                  |
| `GetPaymentLocaleFrBe` | fr_BE                  |
| `GetPaymentLocaleEsEs` | es_ES                  |
| `GetPaymentLocaleCaEs` | ca_ES                  |
| `GetPaymentLocalePtPt` | pt_PT                  |
| `GetPaymentLocaleItIt` | it_IT                  |
| `GetPaymentLocaleNbNo` | nb_NO                  |
| `GetPaymentLocaleSvSe` | sv_SE                  |
| `GetPaymentLocaleFiFi` | fi_FI                  |
| `GetPaymentLocaleDaDk` | da_DK                  |
| `GetPaymentLocaleIsIs` | is_IS                  |
| `GetPaymentLocaleHuHu` | hu_HU                  |
| `GetPaymentLocalePlPl` | pl_PL                  |
| `GetPaymentLocaleLvLv` | lv_LV                  |
| `GetPaymentLocaleLtLt` | lt_LT                  |