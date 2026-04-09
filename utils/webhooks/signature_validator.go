package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const SignatureHeader = "X-Mollie-Signature"

const signaturePrefix = "sha256="

type InvalidSignatureError struct {
	message string
}

func (e *InvalidSignatureError) Error() string {
	return e.message
}

type SignatureValidator struct {
	signingSecrets []string
}

func NewSignatureValidator(signingSecrets ...string) *SignatureValidator {
	secrets := make([]string, len(signingSecrets))
	copy(secrets, signingSecrets)

	return &SignatureValidator{signingSecrets: secrets}
}

func Validate(payload string, signingSecrets []string, signatures ...string) (bool, error) {
	return NewSignatureValidator(signingSecrets...).ValidatePayload(payload, signatures...)
}

func (v *SignatureValidator) ValidatePayload(payload string, signatures ...string) (bool, error) {
	signatureList := normalizeSignatures(signatures)
	if len(signatureList) == 0 {
		return false, nil
	}

	for _, signature := range signatureList {
		if v.isValidSignature(extractSignature(signature), payload) {
			return true, nil
		}
	}

	return false, &InvalidSignatureError{message: "Invalid webhook signature"}
}

func CreateSignature(payload, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(payload))

	return hex.EncodeToString(hash.Sum(nil))
}

func normalizeSignatures(signatures []string) []string {
	result := make([]string, 0, len(signatures))
	for _, signature := range signatures {
		if signature != "" {
			result = append(result, signature)
		}
	}

	return result
}

func extractSignature(signature string) string {
	if strings.HasPrefix(signature, signaturePrefix) {
		return strings.TrimPrefix(signature, signaturePrefix)
	}

	return signature
}

func (v *SignatureValidator) isValidSignature(providedSignature, payload string) bool {
	for _, secret := range v.signingSecrets {
		expectedSignature := CreateSignature(payload, secret)
		if hmac.Equal([]byte(expectedSignature), []byte(providedSignature)) {
			return true
		}
	}

	return false
}
