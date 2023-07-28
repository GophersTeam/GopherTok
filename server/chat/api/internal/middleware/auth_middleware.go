package middleware

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"strconv"
)

type AuthMiddleware struct {
	RedisClient *redis.Redis
}

func NewAuthMiddleware(redisClient *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{
		RedisClient: redisClient,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		if token == "" {
			token = r.PostFormValue("token")
			if token == "" {
				http.Error(w, errorx.ErrTokenEmpty, http.StatusUnauthorized)
				return
			}
		}

		// token不为空，从redis中获取用户id
		userIdStr, err := m.RedisClient.GetCtx(r.Context(), consts.TokenPrefix+token)
		if err != nil {
			logx.Errorf("redis get token failed, err:%v", err)
			http.Error(w, errorx.ErrTokenProve, http.StatusUnauthorized)
			return
		}
		if userIdStr == "" {
			http.Error(w, errorx.ErrTokenProve, http.StatusUnauthorized)
			return
		}

		// userId写入上下文
		userId, _ := strconv.Atoi(userIdStr)
		ctx := context.WithValue(r.Context(), consts.UserId, int64(userId))
		r = r.WithContext(ctx)

		next(w, r)
	}
}
