# ListPaymentsLocale

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ListPaymentsLocaleEnUs` | en_US                    |
| `ListPaymentsLocaleEnGb` | en_GB                    |
| `ListPaymentsLocaleNlNl` | nl_NL                    |
| `ListPaymentsLocaleNlBe` | nl_BE                    |
| `ListPaymentsLocaleDeDe` | de_DE                    |
| `ListPaymentsLocaleDeAt` | de_AT                    |
| `ListPaymentsLocaleDeCh` | de_CH                    |
| `ListPaymentsLocaleFrFr` | fr_FR                    |
| `ListPaymentsLocaleFrBe` | fr_BE                    |
| `ListPaymentsLocaleEsEs` | es_ES                    |
| `ListPaymentsLocaleCaEs` | ca_ES                    |
| `ListPaymentsLocalePtPt` | pt_PT                    |
| `ListPaymentsLocaleItIt` | it_IT                    |
| `ListPaymentsLocaleNbNo` | nb_NO                    |
| `ListPaymentsLocaleSvSe` | sv_SE                    |
| `ListPaymentsLocaleFiFi` | fi_FI                    |
| `ListPaymentsLocaleDaDk` | da_DK                    |
| `ListPaymentsLocaleIsIs` | is_IS                    |
| `ListPaymentsLocaleHuHu` | hu_HU                    |
| `ListPaymentsLocalePlPl` | pl_PL                    |
| `ListPaymentsLocaleLvLv` | lv_LV                    |
| `ListPaymentsLocaleLtLt` | lt_LT                    |