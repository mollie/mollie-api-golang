package hooks

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/mollie/mollie-api-golang/models/components"
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

	// Handle OAuth, testmode and profileId population
	if h.isOAuthRequest(req, hookCtx) {
		modifiedReq, err := h.populateProfileIdAndTestmode(req, hookCtx)
		if err != nil {
			return req, err
		}
		req = modifiedReq
	}

	return req, nil
}

// isOAuthRequest checks if the request is using OAuth authentication
func (h *MollieHooks) isOAuthRequest(req *http.Request, hookCtx BeforeRequestContext) bool {
	if hookCtx.SecuritySource == nil {
		return false
	}

	security, err := hookCtx.SecuritySource(context.Background())
	if err != nil || security == nil {
		return false
	}
	
	// Try both pointer and value types
	var securityStruct *components.Security
	
	if ptr, ok := security.(*components.Security); ok {
		securityStruct = ptr
	} else if val, ok := security.(components.Security); ok {
		securityStruct = &val
	} else {
		return false
	}

	if securityStruct == nil {
		return false
	}

	oauth := securityStruct.GetOAuth()
	if oauth == nil {
		return false
	}

	authHeader := req.Header.Get("Authorization")
	expectedAuth := fmt.Sprintf("Bearer %s", *oauth)
	
	return authHeader == expectedAuth
}

// populateProfileIdAndTestmode adds profileId and testmode to requests when using OAuth
func (h *MollieHooks) populateProfileIdAndTestmode(req *http.Request, hookCtx BeforeRequestContext) (*http.Request, error) {
	clientProfileId := hookCtx.SDKConfiguration.Globals.GetProfileID()
	clientTestmode := hookCtx.SDKConfiguration.Globals.GetTestmode()

	method := req.Method

	if method == "GET" {
		return h.addToQueryParams(req, clientProfileId, clientTestmode)
	}

	// For POST, PATCH, DELETE - update JSON body
	return h.addToJSONBody(req, clientProfileId, clientTestmode)
}

// addToQueryParams adds profileId and testmode to URL query parameters for GET requests
func (h *MollieHooks) addToQueryParams(req *http.Request, clientProfileId *string, clientTestmode *bool) (*http.Request, error) {
	u, err := url.Parse(req.URL.String())
	if err != nil {
		return req, err
	}

	queryParams := u.Query()

	// Add profileId if not already present
	if clientProfileId != nil && !queryParams.Has("profileId") {
		queryParams.Set("profileId", *clientProfileId)
	}

	// Add testmode if not already present  
	if clientTestmode != nil && !queryParams.Has("testmode") {
		queryParams.Set("testmode", strconv.FormatBool(*clientTestmode))
	}

	u.RawQuery = queryParams.Encode()

	// Create new request with updated URL
	newReq, err := http.NewRequest(req.Method, u.String(), req.Body)
	if err != nil {
		return req, err
	}

	// Copy headers
	newReq.Header = req.Header.Clone()
	
	return newReq, nil
}

// addToJSONBody adds profileId and testmode to JSON request body for POST/PATCH/DELETE requests
func (h *MollieHooks) addToJSONBody(req *http.Request, clientProfileId *string, clientTestmode *bool) (*http.Request, error) {
	var body map[string]interface{}

	// Read existing body
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return req, err
		}

		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				// If it's not JSON, return the request unchanged
				req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
				return req, nil
			}
		}
	}

	if body == nil {
		body = make(map[string]interface{})
	}

	// Add profileId if not already present
	if clientProfileId != nil {
		if _, exists := body["profileId"]; !exists {
			body["profileId"] = *clientProfileId
		}
	}

	// Add testmode if not already present
	if clientTestmode != nil {
		if _, exists := body["testmode"]; !exists {
			body["testmode"] = *clientTestmode
		}
	}

	// Marshal back to JSON
	newBodyBytes, err := json.Marshal(body)
	if err != nil {
		return req, err
	}

	// Create new request with updated body
	newReq, err := http.NewRequest(req.Method, req.URL.String(), bytes.NewReader(newBodyBytes))
	if err != nil {
		return req, err
	}

	// Copy and update headers
	newReq.Header = req.Header.Clone()
	newReq.Header.Set("Content-Length", strconv.Itoa(len(newBodyBytes)))

	return newReq, nil
}