package middleware

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/user/api/internal/config"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	Config config.Config
}

func NewJWTMiddleware(c config.Config) *JWTMiddleware {
	return &JWTMiddleware{
		Config: c,
	}
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// JWTAuthMiddleware implementation
		token := r.FormValue("token")
		if token == "" {
			token = r.PostFormValue("token")
			if token == "" {
				http.Error(w, errorx.ErrTokenEmpty, http.StatusUnauthorized)
				return
			}
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			err, _ := json.Marshal(errorx.NewCodeError(30002, errorx.ErrHeadFormat))
			w.Write(err)
			return
		}
		parseToken, isExpire, err := utils.ParseToken(parts[0], parts[1], m.Config.Token.AccessToken, m.Config.Token.RefreshToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			err, _ := json.Marshal(errorx.NewCodeError(30003, errorx.ErrTokenProve))
			w.Write(err)
			return
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
