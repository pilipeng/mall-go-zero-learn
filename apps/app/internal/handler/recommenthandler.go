package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall-go-zero-learn.com/apps/app/internal/logic"
	"mall-go-zero-learn.com/apps/app/internal/svc"
	"mall-go-zero-learn.com/apps/app/internal/types"
)

func RecommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRecommentLogic(r.Context(), svcCtx)
		resp, err := l.Recomment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
