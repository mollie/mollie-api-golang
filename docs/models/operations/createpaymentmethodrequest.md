# CreatePaymentMethodRequest

Normally, a payment method screen is shown. However, when using this parameter, you can choose a specific payment
method and your customer will skip the selection screen and is sent directly to the chosen payment method. The
parameter enables you to fully integrate the payment method selection into your website.

You can also specify the methods in an array. By doing so we will still show the payment method selection screen
but will only show the methods specified in the array. For example, you can use this functionality to only show
payment methods from a specific country to your customer `['bancontact', 'belfius']`.


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreatePaymentMethodRequestAlma`           | alma                                       |
| `CreatePaymentMethodRequestApplepay`       | applepay                                   |
| `CreatePaymentMethodRequestBacs`           | bacs                                       |
| `CreatePaymentMethodRequestBancomatpay`    | bancomatpay                                |
| `CreatePaymentMethodRequestBancontact`     | bancontact                                 |
| `CreatePaymentMethodRequestBanktransfer`   | banktransfer                               |
| `CreatePaymentMethodRequestBelfius`        | belfius                                    |
| `CreatePaymentMethodRequestBillie`         | billie                                     |
| `CreatePaymentMethodRequestBizum`          | bizum                                      |
| `CreatePaymentMethodRequestBlik`           | blik                                       |
| `CreatePaymentMethodRequestCreditcard`     | creditcard                                 |
| `CreatePaymentMethodRequestDirectdebit`    | directdebit                                |
| `CreatePaymentMethodRequestEps`            | eps                                        |
| `CreatePaymentMethodRequestGiftcard`       | giftcard                                   |
| `CreatePaymentMethodRequestIdeal`          | ideal                                      |
| `CreatePaymentMethodRequestIn3`            | in3                                        |
| `CreatePaymentMethodRequestKbc`            | kbc                                        |
| `CreatePaymentMethodRequestKlarna`         | klarna                                     |
| `CreatePaymentMethodRequestKlarnapaylater` | klarnapaylater                             |
| `CreatePaymentMethodRequestKlarnapaynow`   | klarnapaynow                               |
| `CreatePaymentMethodRequestKlarnasliceit`  | klarnasliceit                              |
| `CreatePaymentMethodRequestMbway`          | mbway                                      |
| `CreatePaymentMethodRequestMultibanco`     | multibanco                                 |
| `CreatePaymentMethodRequestMybank`         | mybank                                     |
| `CreatePaymentMethodRequestPaybybank`      | paybybank                                  |
| `CreatePaymentMethodRequestPayconiq`       | payconiq                                   |
| `CreatePaymentMethodRequestPaypal`         | paypal                                     |
| `CreatePaymentMethodRequestPaysafecard`    | paysafecard                                |
| `CreatePaymentMethodRequestPointofsale`    | pointofsale                                |
| `CreatePaymentMethodRequestPrzelewy24`     | przelewy24                                 |
| `CreatePaymentMethodRequestRiverty`        | riverty                                    |
| `CreatePaymentMethodRequestSatispay`       | satispay                                   |
| `CreatePaymentMethodRequestSwish`          | swish                                      |
| `CreatePaymentMethodRequestTrustly`        | trustly                                    |
| `CreatePaymentMethodRequestTwint`          | twint                                      |
| `CreatePaymentMethodRequestVoucher`        | voucher                                    |