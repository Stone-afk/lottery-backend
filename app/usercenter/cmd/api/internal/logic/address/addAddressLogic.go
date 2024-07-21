package address

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/ctxdata"
	"looklook/common/xerr"
)

type AddAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddressLogic) AddAddress(req *types.AddAddressReq) (resp *types.AddAddressResp, err error) {
	pbAddressReq := new(pb.AddUserAddressReq)
	err = copier.Copy(pbAddressReq, req)
	pbAddressReq.UserId = ctxdata.GetUidFromCtx(l.ctx)
	districtByte, err := json.Marshal(req.District)
	if err != nil {
		return nil, err
	}
	pbAddressReq.District = string(districtByte)
	addAddress, err := l.svcCtx.UsercenterRpc.AddUserAddress(l.ctx, pbAddressReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("add address fail"), "add address rpc AddUserAddress fail req: %+v , err : %v ", req, err)
	}
	return &types.AddAddressResp{Id: addAddress.Id}, nil
}
