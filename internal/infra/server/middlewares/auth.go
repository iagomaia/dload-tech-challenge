package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	sharedcontracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/shared"
	"github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	adaptersfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/adapters"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authError := &sharedcontracts.MessageResponse{
			Message: "Authentication failed",
		}
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		jwtAdapter := adaptersfactories.GetJwtAdapter()
		claims, err := jwtAdapter.Verify(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(authError)
			return
		}
		userId := claims["sub"].(string)
		ctx := context.WithValue(r.Context(), utils.UserIDContextKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
