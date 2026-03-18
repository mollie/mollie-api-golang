# BalanceFeeType

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceFeeTypePaymentFee

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceFeeType("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BalanceFeeTypePaymentFee`                                        | payment-fee                                                       |
| `BalanceFeeTypeDirectDebitFailureFee`                             | direct-debit-failure-fee                                          |
| `BalanceFeeTypeUnauthorizedDirectDebitFee`                        | unauthorized-direct-debit-fee                                     |
| `BalanceFeeTypeBankChargedDirectDebitFailureFee`                  | bank-charged-direct-debit-failure-fee                             |
| `BalanceFeeTypePartnerCommission`                                 | partner-commission                                                |
| `BalanceFeeTypeApplicationFee`                                    | application-fee                                                   |
| `BalanceFeeTypeCaptureFee`                                        | capture-fee                                                       |
| `BalanceFeeTypeRefundFee`                                         | refund-fee                                                        |
| `BalanceFeeTypeChargebackFee`                                     | chargeback-fee                                                    |
| `BalanceFeeTypePaymentNotificationFee`                            | payment-notification-fee                                          |
| `BalanceFeeTypeTransferNotificationFee`                           | transfer-notification-fee                                         |
| `BalanceFeeTypePayoutFee`                                         | payout-fee                                                        |
| `BalanceFeeTypeFeeDiscount`                                       | fee-discount                                                      |
| `BalanceFeeTypeFeeReimbursement`                                  | fee-reimbursement                                                 |
| `BalanceFeeTypePlatformVolumeFee`                                 | platform-volume-fee                                               |
| `BalanceFeeTypePlatformConnectedOrganizationsFee`                 | platform-connected-organizations-fee                              |
| `BalanceFeeTypeBalanceChargeFee`                                  | balance-charge-fee                                                |
| `BalanceFeeTypeThreedsAuthenticationAttemptFee`                   | 3ds-authentication-attempt-fee                                    |
| `BalanceFeeTypeTerminalMonthlyFee`                                | terminal-monthly-fee                                              |
| `BalanceFeeTypeAcceptanceRiskFee`                                 | acceptance-risk-fee                                               |
| `BalanceFeeTypeTopUpFee`                                          | top-up-fee                                                        |
| `BalanceFeeTypePaymentGatewayFee`                                 | payment-gateway-fee                                               |
| `BalanceFeeTypeMastercardSpecialtyMerchantProgramProcessingFee`   | mastercard-specialty-merchant-program-processing-fee              |
| `BalanceFeeTypeMastercardSpecialtyMerchantProgramRegistrationFee` | mastercard-specialty-merchant-program-registration-fee            |
| `BalanceFeeTypeVisaIntegrityRiskProgramProcessingFee`             | visa-integrity-risk-program-processing-fee                        |
| `BalanceFeeTypeVisaIntegrityRiskProgramRegistrationFee`           | visa-integrity-risk-program-registration-fee                      |
| `BalanceFeeTypeMinimumInvoiceAmountFee`                           | minimum-invoice-amount-fee                                        |