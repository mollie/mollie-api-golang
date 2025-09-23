package hooks

import (
	"sync"
	"crypto/rand"
	"fmt"
	"net/http"
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

	return req, nil
}