# ListSettlementPaymentsLocale

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
| `ListSettlementPaymentsLocaleEnUs` | en_US                              |
| `ListSettlementPaymentsLocaleEnGb` | en_GB                              |
| `ListSettlementPaymentsLocaleNlNl` | nl_NL                              |
| `ListSettlementPaymentsLocaleNlBe` | nl_BE                              |
| `ListSettlementPaymentsLocaleDeDe` | de_DE                              |
| `ListSettlementPaymentsLocaleDeAt` | de_AT                              |
| `ListSettlementPaymentsLocaleDeCh` | de_CH                              |
| `ListSettlementPaymentsLocaleFrFr` | fr_FR                              |
| `ListSettlementPaymentsLocaleFrBe` | fr_BE                              |
| `ListSettlementPaymentsLocaleEsEs` | es_ES                              |
| `ListSettlementPaymentsLocaleCaEs` | ca_ES                              |
| `ListSettlementPaymentsLocalePtPt` | pt_PT                              |
| `ListSettlementPaymentsLocaleItIt` | it_IT                              |
| `ListSettlementPaymentsLocaleNbNo` | nb_NO                              |
| `ListSettlementPaymentsLocaleSvSe` | sv_SE                              |
| `ListSettlementPaymentsLocaleFiFi` | fi_FI                              |
| `ListSettlementPaymentsLocaleDaDk` | da_DK                              |
| `ListSettlementPaymentsLocaleIsIs` | is_IS                              |
| `ListSettlementPaymentsLocaleHuHu` | hu_HU                              |
| `ListSettlementPaymentsLocalePlPl` | pl_PL                              |
| `ListSettlementPaymentsLocaleLvLv` | lv_LV                              |
| `ListSettlementPaymentsLocaleLtLt` | lt_LT                              |