package middleware

import (
	"cloud_go/Disk/common"
	"cloud_go/Disk/common/redis"
	"cloud_go/Disk/define"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token = r.Header.Get("Authorization")
		if token == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, "请先登录!😼")
			return
		}

		if redis.Redis == nil {
			return
		}

		claim, err := common.AnalyzeToken(token)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, "身份认证错误或过期，请重新登录!")
			return
		}

		id := claim.Id
		key := redis.UserLogin + strconv.FormatInt(id, 10)

		redisToken, err := redis.Redis.Get(r.Context(), key).Result()
		if redisToken != token {
			httpx.WriteJson(w, http.StatusUnauthorized, "身份认证过期，请重新登录!")
			return
		}

		ctx := context.WithValue(r.Context(), define.UserIdKey, id)
		ctx = context.WithValue(ctx, define.UserNameKey, claim.Name)
		request := r.WithContext(ctx)
		next(w, request)
	}
}
