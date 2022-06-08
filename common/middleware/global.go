package middleware

import (
	"net/http"
)

// CommonJwtAuthMiddleware : with jwt on the verification, no jwt on the verification
type CommonJwtAuthMiddleware struct {
	secret string
}

func NewCommonJwtAuthMiddleware(secret string) *CommonJwtAuthMiddleware {
	return &CommonJwtAuthMiddleware{
		secret: secret,
	}
}

func (m *CommonJwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)

	}
}
