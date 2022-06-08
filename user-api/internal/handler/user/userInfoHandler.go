package user

import (
	"net/http"

	"go-zore/user-api/internal/logic/user"
	"go-zore/user-api/internal/svc"
	"go-zore/user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest

		if err := httpx.Parse(r, &req); err != nil { //用户的输入解析到请求体上
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
