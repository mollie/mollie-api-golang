package hooks

import (
	"sync"
	"crypto/rand"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

var (
	_ beforeRequestHook = (*MollieHooks)(nil)
)

type MollieHooks struct {
	mu sync.Mutex
}

func (h *MollieHooks) generateIdempotencyKey() (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	u := make([]byte, 16)

	// Read 16 random bytes
	if _, err := rand.Read(u); err != nil {
		return "", err
	}
	
	u[6] = (u[6] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16]), nil
}

func (h *MollieHooks) customizeUserAgent(req *http.Request, hookCtx BeforeRequestContext) {
    const userAgentKey = "User-Agent"

    customUserAgent := hookCtx.SDKConfiguration.Globals.CustomUserAgent

    // Parse from existing UserAgent string: "speakeasy-sdk/go 0.8.1 2.730.5 1.0.0 github.com/mollie/mollie-api-golang"
    userAgentParts := strings.Split(hookCtx.SDKConfiguration.UserAgent, " ")
    
    sdkVersion := userAgentParts[1]   // "0.8.1"
    genVersion := userAgentParts[2]   // "2.730.5"
    packageName := userAgentParts[4]  // "github.com/mollie/mollie-api-golang"

    // Get Go runtime version (e.g., "go1.21.0")
    goVersion := runtime.Version()

    mollieUserAgent := fmt.Sprintf("Speakeasy/%s Golang/%s %s/%s", genVersion, goVersion, packageName, sdkVersion)
    
    if customUserAgent != nil && *customUserAgent != "" {
        mollieUserAgent = fmt.Sprintf("%s %s", mollieUserAgent, *customUserAgent)
    }

    req.Header.Set(userAgentKey, mollieUserAgent)
}

func (h *MollieHooks) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	const (
		idempotencyKey = "idempotency-key"
	)

	if req.Header.Get(idempotencyKey) == "" {
		key, err := h.generateIdempotencyKey()

		if err != nil {
			return req, nil
		}
		
		req.Header.Set(idempotencyKey, key)
	}

	// Customize the User-Agent header
	h.customizeUserAgent(req, hookCtx)

	return req, nil
}