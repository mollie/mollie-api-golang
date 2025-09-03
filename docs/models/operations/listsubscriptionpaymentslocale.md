# ListSubscriptionPaymentsLocale

Allows you to preset the language to be used in the hosted payment pages shown to the customer. Setting a locale
is highly recommended and will greatly improve your conversion rate. When this parameter is omitted the browser
language will be used instead if supported by the payment method. You can provide any `xx_XX` format ISO 15897
locale, but our hosted payment pages currently only support the specified languages.

For bank transfer payments specifically, the locale will determine the target bank account the customer has to
transfer the money to. We have dedicated bank accounts for Belgium, Germany, and The Netherlands. Having the
customer use a local bank account greatly increases the conversion and speed of payment.


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ListSubscriptionPaymentsLocaleEnUs` | en_US                                |
| `ListSubscriptionPaymentsLocaleEnGb` | en_GB                                |
| `ListSubscriptionPaymentsLocaleNlNl` | nl_NL                                |
| `ListSubscriptionPaymentsLocaleNlBe` | nl_BE                                |
| `ListSubscriptionPaymentsLocaleDeDe` | de_DE                                |
| `ListSubscriptionPaymentsLocaleDeAt` | de_AT                                |
| `ListSubscriptionPaymentsLocaleDeCh` | de_CH                                |
| `ListSubscriptionPaymentsLocaleFrFr` | fr_FR                                |
| `ListSubscriptionPaymentsLocaleFrBe` | fr_BE                                |
| `ListSubscriptionPaymentsLocaleEsEs` | es_ES                                |
| `ListSubscriptionPaymentsLocaleCaEs` | ca_ES                                |
| `ListSubscriptionPaymentsLocalePtPt` | pt_PT                                |
| `ListSubscriptionPaymentsLocaleItIt` | it_IT                                |
| `ListSubscriptionPaymentsLocaleNbNo` | nb_NO                                |
| `ListSubscriptionPaymentsLocaleSvSe` | sv_SE                                |
| `ListSubscriptionPaymentsLocaleFiFi` | fi_FI                                |
| `ListSubscriptionPaymentsLocaleDaDk` | da_DK                                |
| `ListSubscriptionPaymentsLocaleIsIs` | is_IS                                |
| `ListSubscriptionPaymentsLocaleHuHu` | hu_HU                                |
| `ListSubscriptionPaymentsLocalePlPl` | pl_PL                                |
| `ListSubscriptionPaymentsLocaleLvLv` | lv_LV                                |
| `ListSubscriptionPaymentsLocaleLtLt` | lt_LT                                |