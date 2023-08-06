package middleware

import (
	"GopherTok/common/consts"
	"GopherTok/common/utils"
	"GopherTok/server/video/api/internal/config"
	"context"
	"net/http"
	"strings"
)

type VideoJWTMiddleware struct {
	Config config.Config
}

func NewVideoJWTMiddleware(c config.Config) *VideoJWTMiddleware {
	return &VideoJWTMiddleware{
		Config: c,
	}
}

func (m *VideoJWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// TODO generate middleware implement function, delete after code implementation
		// JWTAuthMiddleware implementation
		token := r.FormValue("token")
		if token == "" {
			token = r.PostFormValue("token")
			if token == "" {

			}
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 {

		}
		parseToken, isExpire, err := utils.ParseToken(parts[0], parts[1], m.Config.Token.AccessToken, m.Config.Token.RefreshToken)
		if err != nil {

		}
		if isExpire {
			parts[0], parts[1] = utils.GetToken(parseToken.ID, parseToken.State, m.Config.Token.AccessToken, m.Config.Token.RefreshToken)
			//w.Header().Set("Authorization", fmt.Sprintf("Bearer %s %s", parts[0], parts[1]))

		}
		token = parts[0] + " " + parts[1]
		r = r.WithContext(context.WithValue(r.Context(), consts.UserId, parseToken.ID))
		r = r.WithContext(context.WithValue(r.Context(), consts.Token, token))
		next(w, r)

	}
}
