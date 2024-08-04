package event

import (
	"context"
	"encoding/xml"
	"github.com/pkg/errors"
	"looklook/app/notice/cmd/rpc/pb"
	"looklook/common/constants"
	"net/http"

	"looklook/app/notice/cmd/api/internal/svc"
	"looklook/app/notice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrReceiveEventFail = errors.New("receive event fail")

type ReceiveEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiveEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveEventLogic {
	return &ReceiveEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveEventLogic) ReceiveEvent(_ *types.ReceiveEventReq, r *http.Request) (resp *types.ReceiveEventResp, err error) {
	// 接收回调事件
	_, callbackMsgHeader, err := l.svcCtx.WxMiniProgram.Server.GetEvent(r)
	if err != nil {
		return nil, errors.Wrapf(ErrReceiveEventFail, "ReceiveEventLogic get event err:%v", err)
	}

	logx.WithContext(l.ctx).Infow("ReceiveEventLogic received an event",
		logx.Field("content", string(callbackMsgHeader.Content)),
	)
	// 解析事件内容
	var msg types.MsgEvent
	err = xml.Unmarshal(callbackMsgHeader.Content, &msg)
	if err != nil {
		return nil, errors.Wrapf(ErrReceiveEventFail, "ReceiveEventLogic event xml unmarshal err:%v", err)
	}

	// 处理事件
	userSubscribeSettings := msg.SubscribeMsgPopupEvent.List
	for _, setting := range userSubscribeSettings {
		// 落库
		var typ int64
		switch setting.SubscribeStatusString {
		case "accept":
			typ = constants.TypeAcceptSubscribeMessage
		case "reject":
			typ = constants.TypeRejectSubscribeMessage
		default:
			return nil, errors.Wrapf(ErrReceiveEventFail, "ReceiveEventLogic event invalid setting, setting:%+v", setting)
		}

		_, err = l.svcCtx.NoticeRpc.SaveNoticeSubscribePreference(l.ctx, &pb.SaveNoticeSubscribePreferenceReq{
			Openid:     msg.ToUserName,
			TemplateId: setting.TemplateId,
			Type:       typ,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrReceiveEventFail, "ReceiveEventLogic NoticeRpc.SaveNoticeSubscribePreference err:%v, setting:%+v", err, setting)
		}
	}
	return
}
