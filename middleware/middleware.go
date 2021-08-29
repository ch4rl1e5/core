package middleware

import (
	"github.com/ch4rl1e5/go-common/constants"
	"github.com/ch4rl1e5/go-common/httphelper"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func ValidateUUID(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		uuidString := chi.URLParam(r, constants.URLIdParam)
		_, err := uuid.Parse(uuidString)
		if err != nil {
			httphelper.HandleError(w, httphelper.ErrInvalidUUID)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
