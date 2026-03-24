package hooks

import (
	_ "embed"
	"bytes"
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

	"github.com/mollie/mollie-api-golang/internal/utils"
)

//go:embed global_usage.json
var globalUsageJSON []byte

var globalUsage map[string][]string

func init() {
	globalUsage = make(map[string][]string)
	_ = json.Unmarshal(globalUsageJSON, &globalUsage)
}

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

func (h *MollieHooks) validatePathParameters(req *http.Request, hookCtx BeforeRequestContext) error {
	pathSegments := strings.Split(req.URL.Path, "/")
	
	for i, segment := range pathSegments {
		if i == 0 && segment == "" {
			continue
		}
		if segment == "" || strings.TrimSpace(segment) == "" {
			return fmt.Errorf("invalid request: empty path parameter detected in [%s] '%s'", req.Method, req.URL.Path)
		}
	}
	return nil
}

func (h *MollieHooks) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	// Prevent Empty Path Parameters
	if err := h.validatePathParameters(req, hookCtx); err != nil {
		return nil, err
	}
	
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

	// Inject global fields (profileId, testmode) for OAuth requests per operation
	if utils.CanHaveGlobalFields(hookCtx.SecuritySource) {
		modifiedReq, err := h.injectGlobalFields(req, hookCtx)
		if err != nil {
			return req, err
		}
		req = modifiedReq
	}

	return req, nil
}

// injectGlobalFields injects profileId and/or testmode into the request for operations
// listed in global_usage.json, when using an OAuth access token.
func (h *MollieHooks) injectGlobalFields(req *http.Request, hookCtx BeforeRequestContext) (*http.Request, error) {
	operationID := hookCtx.OperationID

	// Build globals as a flat map (JSON field name → value)
	globalsMap := map[string]interface{}{}
	if hookCtx.SDKConfiguration.Globals.ProfileID != nil {
		globalsMap["profileId"] = *hookCtx.SDKConfiguration.Globals.ProfileID
	}
	if hookCtx.SDKConfiguration.Globals.Testmode != nil {
		globalsMap["testmode"] = *hookCtx.SDKConfiguration.Globals.Testmode
	}

	// Filter: only inject fields whose operation list contains this operationID
	fieldsToInject := map[string]interface{}{}
	for field, ops := range globalUsage {
		if _, hasValue := globalsMap[field]; !hasValue {
			continue
		}
		for _, op := range ops {
			if op == operationID {
				fieldsToInject[field] = globalsMap[field]
				break
			}
		}
	}

	if len(fieldsToInject) == 0 {
		return req, nil
	}

	if req.Method == "GET" {
		return h.addToQueryParams(req, fieldsToInject)
	}
	return h.addToJSONBody(req, fieldsToInject)
}

// addToQueryParams adds fields to URL query parameters for GET requests
func (h *MollieHooks) addToQueryParams(req *http.Request, fields map[string]interface{}) (*http.Request, error) {
	u, err := url.Parse(req.URL.String())
	if err != nil {
		return req, err
	}

	queryParams := u.Query()

	for field, value := range fields {
		if !queryParams.Has(field) {
			switch v := value.(type) {
			case bool:
				queryParams.Set(field, strconv.FormatBool(v))
			default:
				queryParams.Set(field, fmt.Sprintf("%v", v))
			}
		}
	}

	u.RawQuery = queryParams.Encode()

	newReq, err := http.NewRequest(req.Method, u.String(), req.Body)
	if err != nil {
		return req, err
	}
	newReq.Header = req.Header.Clone()
	return newReq, nil
}

// addToJSONBody adds fields to JSON request body for POST/PATCH/DELETE requests
func (h *MollieHooks) addToJSONBody(req *http.Request, fields map[string]interface{}) (*http.Request, error) {
	var body map[string]interface{}

	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return req, err
		}

		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				// Not JSON — return unchanged
				req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
				return req, nil
			}
		}
	}

	if body == nil {
		body = make(map[string]interface{})
	}

	for field, value := range fields {
		if _, exists := body[field]; !exists {
			body[field] = value
		}
	}

	newBodyBytes, err := json.Marshal(body)
	if err != nil {
		return req, err
	}

	newReq, err := http.NewRequest(req.Method, req.URL.String(), bytes.NewReader(newBodyBytes))
	if err != nil {
		return req, err
	}

	newReq.Header = req.Header.Clone()
	newReq.Header.Set("Content-Length", strconv.Itoa(len(newBodyBytes)))
	return newReq, nil
}
