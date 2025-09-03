# CreateCustomerPaymentMethodRequest

Normally, a payment method screen is shown. However, when using this parameter, you can choose a specific payment
method and your customer will skip the selection screen and is sent directly to the chosen payment method. The
parameter enables you to fully integrate the payment method selection into your website.

You can also specify the methods in an array. By doing so we will still show the payment method selection screen
but will only show the methods specified in the array. For example, you can use this functionality to only show
payment methods from a specific country to your customer `['bancontact', 'belfius']`.


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `CreateCustomerPaymentMethodRequestAlma`           | alma                                               |
| `CreateCustomerPaymentMethodRequestApplepay`       | applepay                                           |
| `CreateCustomerPaymentMethodRequestBacs`           | bacs                                               |
| `CreateCustomerPaymentMethodRequestBancomatpay`    | bancomatpay                                        |
| `CreateCustomerPaymentMethodRequestBancontact`     | bancontact                                         |
| `CreateCustomerPaymentMethodRequestBanktransfer`   | banktransfer                                       |
| `CreateCustomerPaymentMethodRequestBelfius`        | belfius                                            |
| `CreateCustomerPaymentMethodRequestBillie`         | billie                                             |
| `CreateCustomerPaymentMethodRequestBizum`          | bizum                                              |
| `CreateCustomerPaymentMethodRequestBlik`           | blik                                               |
| `CreateCustomerPaymentMethodRequestCreditcard`     | creditcard                                         |
| `CreateCustomerPaymentMethodRequestDirectdebit`    | directdebit                                        |
| `CreateCustomerPaymentMethodRequestEps`            | eps                                                |
| `CreateCustomerPaymentMethodRequestGiftcard`       | giftcard                                           |
| `CreateCustomerPaymentMethodRequestIdeal`          | ideal                                              |
| `CreateCustomerPaymentMethodRequestIn3`            | in3                                                |
| `CreateCustomerPaymentMethodRequestKbc`            | kbc                                                |
| `CreateCustomerPaymentMethodRequestKlarna`         | klarna                                             |
| `CreateCustomerPaymentMethodRequestKlarnapaylater` | klarnapaylater                                     |
| `CreateCustomerPaymentMethodRequestKlarnapaynow`   | klarnapaynow                                       |
| `CreateCustomerPaymentMethodRequestKlarnasliceit`  | klarnasliceit                                      |
| `CreateCustomerPaymentMethodRequestMbway`          | mbway                                              |
| `CreateCustomerPaymentMethodRequestMultibanco`     | multibanco                                         |
| `CreateCustomerPaymentMethodRequestMybank`         | mybank                                             |
| `CreateCustomerPaymentMethodRequestPaybybank`      | paybybank                                          |
| `CreateCustomerPaymentMethodRequestPayconiq`       | payconiq                                           |
| `CreateCustomerPaymentMethodRequestPaypal`         | paypal                                             |
| `CreateCustomerPaymentMethodRequestPaysafecard`    | paysafecard                                        |
| `CreateCustomerPaymentMethodRequestPointofsale`    | pointofsale                                        |
| `CreateCustomerPaymentMethodRequestPrzelewy24`     | przelewy24                                         |
| `CreateCustomerPaymentMethodRequestRiverty`        | riverty                                            |
| `CreateCustomerPaymentMethodRequestSatispay`       | satispay                                           |
| `CreateCustomerPaymentMethodRequestSwish`          | swish                                              |
| `CreateCustomerPaymentMethodRequestTrustly`        | trustly                                            |
| `CreateCustomerPaymentMethodRequestTwint`          | twint                                              |
| `CreateCustomerPaymentMethodRequestVoucher`        | voucher                                            |