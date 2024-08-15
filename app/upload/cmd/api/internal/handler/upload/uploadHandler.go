package upload

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/upload/cmd/api/internal/logic/upload"
	"looklook/app/upload/cmd/api/internal/svc"
	"looklook/app/upload/cmd/api/internal/types"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)

		result.HttpResult(r, w, resp, err)
	}
}
