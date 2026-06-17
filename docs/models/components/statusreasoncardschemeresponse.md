# StatusReasonCardSchemeResponse

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.StatusReasonCardSchemeResponseApprovedOrCompletedSuccessfully

// Open enum: custom values can be created with a direct type cast
custom := components.StatusReasonCardSchemeResponse("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `StatusReasonCardSchemeResponseApprovedOrCompletedSuccessfully`             | approved_or_completed_successfully                                          |
| `StatusReasonCardSchemeResponseReferToCardIssuer`                           | refer_to_card_issuer                                                        |
| `StatusReasonCardSchemeResponseInvalidMerchant`                             | invalid_merchant                                                            |
| `StatusReasonCardSchemeResponseCaptureCard`                                 | capture_card                                                                |
| `StatusReasonCardSchemeResponseDoNotHonor`                                  | do_not_honor                                                                |
| `StatusReasonCardSchemeResponseError`                                       | error                                                                       |
| `StatusReasonCardSchemeResponsePartialApproval`                             | partial_approval                                                            |
| `StatusReasonCardSchemeResponseInvalidTransaction`                          | invalid_transaction                                                         |
| `StatusReasonCardSchemeResponseInvalidAmount`                               | invalid_amount                                                              |
| `StatusReasonCardSchemeResponseInvalidIssuer`                               | invalid_issuer                                                              |
| `StatusReasonCardSchemeResponseLostCard`                                    | lost_card                                                                   |
| `StatusReasonCardSchemeResponseStolenCard`                                  | stolen_card                                                                 |
| `StatusReasonCardSchemeResponseInsufficientFunds`                           | insufficient_funds                                                          |
| `StatusReasonCardSchemeResponseExpiredCard`                                 | expired_card                                                                |
| `StatusReasonCardSchemeResponseInvalidPin`                                  | invalid_pin                                                                 |
| `StatusReasonCardSchemeResponseTransactionNotPermittedToCardholder`         | transaction_not_permitted_to_cardholder                                     |
| `StatusReasonCardSchemeResponseTransactionNotAllowedAtTerminal`             | transaction_not_allowed_at_terminal                                         |
| `StatusReasonCardSchemeResponseExceedsWithdrawalAmountLimit`                | exceeds_withdrawal_amount_limit                                             |
| `StatusReasonCardSchemeResponseRestrictedCard`                              | restricted_card                                                             |
| `StatusReasonCardSchemeResponseSecurityViolation`                           | security_violation                                                          |
| `StatusReasonCardSchemeResponseExceedsWithdrawalCountLimit`                 | exceeds_withdrawal_count_limit                                              |
| `StatusReasonCardSchemeResponseAllowableNumberOfPinTriesExceeded`           | allowable_number_of_pin_tries_exceeded                                      |
| `StatusReasonCardSchemeResponseNoReasonToDecline`                           | no_reason_to_decline                                                        |
| `StatusReasonCardSchemeResponseCannotVerifyPin`                             | cannot_verify_pin                                                           |
| `StatusReasonCardSchemeResponseIssuerUnavailable`                           | issuer_unavailable                                                          |
| `StatusReasonCardSchemeResponseUnableToRouteTransaction`                    | unable_to_route_transaction                                                 |
| `StatusReasonCardSchemeResponseDuplicateTransaction`                        | duplicate_transaction                                                       |
| `StatusReasonCardSchemeResponseSystemMalfunction`                           | system_malfunction                                                          |
| `StatusReasonCardSchemeResponseHonorWithID`                                 | honor_with_id                                                               |
| `StatusReasonCardSchemeResponseInvalidCardNumber`                           | invalid_card_number                                                         |
| `StatusReasonCardSchemeResponseFormatError`                                 | format_error                                                                |
| `StatusReasonCardSchemeResponseContactCardIssuer`                           | contact_card_issuer                                                         |
| `StatusReasonCardSchemeResponsePinNotChanged`                               | pin_not_changed                                                             |
| `StatusReasonCardSchemeResponseInvalidNonexistentToAccountSpecified`        | invalid_nonexistent_to_account_specified                                    |
| `StatusReasonCardSchemeResponseInvalidNonexistentFromAccountSpecified`      | invalid_nonexistent_from_account_specified                                  |
| `StatusReasonCardSchemeResponseInvalidNonexistentAccountSpecified`          | invalid_nonexistent_account_specified                                       |
| `StatusReasonCardSchemeResponseLifecycleRelated`                            | lifecycle_related                                                           |
| `StatusReasonCardSchemeResponseDomesticDebitTransactionNotAllowed`          | domestic_debit_transaction_not_allowed                                      |
| `StatusReasonCardSchemeResponsePolicyRelated`                               | policy_related                                                              |
| `StatusReasonCardSchemeResponseFraudSecurityRelated`                        | fraud_security_related                                                      |
| `StatusReasonCardSchemeResponseInvalidAuthorizationLifeCycle`               | invalid_authorization_life_cycle                                            |
| `StatusReasonCardSchemeResponsePurchaseAmountOnlyNoCashBackAllowed`         | purchase_amount_only_no_cash_back_allowed                                   |
| `StatusReasonCardSchemeResponseCryptographicFailure`                        | cryptographic_failure                                                       |
| `StatusReasonCardSchemeResponseUnacceptablePin`                             | unacceptable_pin                                                            |
| `StatusReasonCardSchemeResponseReferToCardIssuerSpecialCondition`           | refer_to_card_issuer_special_condition                                      |
| `StatusReasonCardSchemeResponsePickUpCardSpecialCondition`                  | pick_up_card_special_condition                                              |
| `StatusReasonCardSchemeResponseVipApproval`                                 | vip_approval                                                                |
| `StatusReasonCardSchemeResponseInvalidAccountNumber`                        | invalid_account_number                                                      |
| `StatusReasonCardSchemeResponseReEnterTransaction`                          | re_enter_transaction                                                        |
| `StatusReasonCardSchemeResponseNoActionTaken`                               | no_action_taken                                                             |
| `StatusReasonCardSchemeResponseUnableToLocateRecord`                        | unable_to_locate_record                                                     |
| `StatusReasonCardSchemeResponseFileTemporarilyUnavailable`                  | file_temporarily_unavailable                                                |
| `StatusReasonCardSchemeResponseNoCreditAccount`                             | no_credit_account                                                           |
| `StatusReasonCardSchemeResponseClosedAccount`                               | closed_account                                                              |
| `StatusReasonCardSchemeResponseNoCheckingAccount`                           | no_checking_account                                                         |
| `StatusReasonCardSchemeResponseNoSavingsAccount`                            | no_savings_account                                                          |
| `StatusReasonCardSchemeResponseSuspectedFraud`                              | suspected_fraud                                                             |
| `StatusReasonCardSchemeResponseTransactionDoesNotFulfillAmlRequirement`     | transaction_does_not_fulfill_aml_requirement                                |
| `StatusReasonCardSchemeResponsePinDataRequired`                             | pin_data_required                                                           |
| `StatusReasonCardSchemeResponseUnableToLocatePreviousMessage`               | unable_to_locate_previous_message                                           |
| `StatusReasonCardSchemeResponsePreviousMessageLocatedInconsistentData`      | previous_message_located_inconsistent_data                                  |
| `StatusReasonCardSchemeResponseBlockedFirstUsed`                            | blocked_first_used                                                          |
| `StatusReasonCardSchemeResponseTransactionReversed`                         | transaction_reversed                                                        |
| `StatusReasonCardSchemeResponseCreditIssuerUnavailable`                     | credit_issuer_unavailable                                                   |
| `StatusReasonCardSchemeResponsePinCryptographicErrorFound`                  | pin_cryptographic_error_found                                               |
| `StatusReasonCardSchemeResponseNegativeOnlineCamResult`                     | negative_online_cam_result                                                  |
| `StatusReasonCardSchemeResponseViolationOfLaw`                              | violation_of_law                                                            |
| `StatusReasonCardSchemeResponseForceStip`                                   | force_stip                                                                  |
| `StatusReasonCardSchemeResponseCashServiceNotAvailable`                     | cash_service_not_available                                                  |
| `StatusReasonCardSchemeResponseCashbackRequestExceedsIssuerLimit`           | cashback_request_exceeds_issuer_limit                                       |
| `StatusReasonCardSchemeResponseDeclineForCvv2Failure`                       | decline_for_cvv2_failure                                                    |
| `StatusReasonCardSchemeResponseTransactionAmountExceedsPreAuthorizedAmount` | transaction_amount_exceeds_pre_authorized_amount                            |
| `StatusReasonCardSchemeResponseInvalidBillerInformation`                    | invalid_biller_information                                                  |
| `StatusReasonCardSchemeResponsePinChangeUnblockRequestDeclined`             | pin_change_unblock_request_declined                                         |
| `StatusReasonCardSchemeResponseUnsafePin`                                   | unsafe_pin                                                                  |
| `StatusReasonCardSchemeResponseCardAuthenticationFailed`                    | card_authentication_failed                                                  |
| `StatusReasonCardSchemeResponseStopPaymentOrder`                            | stop_payment_order                                                          |
| `StatusReasonCardSchemeResponseRevocationOfAuthorization`                   | revocation_of_authorization                                                 |
| `StatusReasonCardSchemeResponseRevocationOfAllAuthorizations`               | revocation_of_all_authorizations                                            |
| `StatusReasonCardSchemeResponseForwardToIssuerXa`                           | forward_to_issuer_xa                                                        |
| `StatusReasonCardSchemeResponseForwardToIssuerXd`                           | forward_to_issuer_xd                                                        |
| `StatusReasonCardSchemeResponseUnableToGoOnline`                            | unable_to_go_online                                                         |
| `StatusReasonCardSchemeResponseAdditionalCustomerAuthenticationRequired`    | additional_customer_authentication_required                                 |