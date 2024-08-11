package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/shop/model"
	"looklook/common/xerr"

	"looklook/app/shop/cmd/rpc/internal/svc"
	"looklook/app/shop/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsListLogic) GetGoodsList(in *pb.GoodsListReq) (*pb.GoodsListResp, error) {
	list, err := l.svcCtx.GoodsModel.List(l.ctx, in.PageSize, 5)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "PageSize:%d,err:%v", in.PageSize, err)
	}
	//根据响应,先创建一个指针数组,用于存储每一个赋值的商品信息指针。
	goods := make([]*pb.Goods, 0)
	//list为查询的分页列表
	for _, goodsInfo := range list {
		//创建pbGoods指针
		pbGoods := new(pb.Goods)
		//使用copy函数进行copy
		err := copier.Copy(pbGoods, goodsInfo)
		if err != nil {
			return nil, err
		}
		//再把指针pbGoods添加到goods指针数组中
		goods = append(goods, pbGoods)
	}
	//最后返回指针数组即可
	return &pb.GoodsListResp{
		Goods: goods,
	}, nil
}
