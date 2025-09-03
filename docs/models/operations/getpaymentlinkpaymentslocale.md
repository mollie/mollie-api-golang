# GetPaymentLinkPaymentsLocale

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `GetPaymentLinkPaymentsLocaleEnUs` | en_US                              |
| `GetPaymentLinkPaymentsLocaleEnGb` | en_GB                              |
| `GetPaymentLinkPaymentsLocaleNlNl` | nl_NL                              |
| `GetPaymentLinkPaymentsLocaleNlBe` | nl_BE                              |
| `GetPaymentLinkPaymentsLocaleDeDe` | de_DE                              |
| `GetPaymentLinkPaymentsLocaleDeAt` | de_AT                              |
| `GetPaymentLinkPaymentsLocaleDeCh` | de_CH                              |
| `GetPaymentLinkPaymentsLocaleFrFr` | fr_FR                              |
| `GetPaymentLinkPaymentsLocaleFrBe` | fr_BE                              |
| `GetPaymentLinkPaymentsLocaleEsEs` | es_ES                              |
| `GetPaymentLinkPaymentsLocaleCaEs` | ca_ES                              |
| `GetPaymentLinkPaymentsLocalePtPt` | pt_PT                              |
| `GetPaymentLinkPaymentsLocaleItIt` | it_IT                              |
| `GetPaymentLinkPaymentsLocaleNbNo` | nb_NO                              |
| `GetPaymentLinkPaymentsLocaleSvSe` | sv_SE                              |
| `GetPaymentLinkPaymentsLocaleFiFi` | fi_FI                              |
| `GetPaymentLinkPaymentsLocaleDaDk` | da_DK                              |
| `GetPaymentLinkPaymentsLocaleIsIs` | is_IS                              |
| `GetPaymentLinkPaymentsLocaleHuHu` | hu_HU                              |
| `GetPaymentLinkPaymentsLocalePlPl` | pl_PL                              |
| `GetPaymentLinkPaymentsLocaleLvLv` | lv_LV                              |
| `GetPaymentLinkPaymentsLocaleLtLt` | lt_LT                              |