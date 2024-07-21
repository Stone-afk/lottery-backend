// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserAddressReq          = pb.AddUserAddressReq
	AddUserAddressResp         = pb.AddUserAddressResp
	AddUserAuthReq             = pb.AddUserAuthReq
	AddUserAuthResp            = pb.AddUserAuthResp
	AddUserContactReq          = pb.AddUserContactReq
	AddUserContactResp         = pb.AddUserContactResp
	AddUserDynamicReq          = pb.AddUserDynamicReq
	AddUserDynamicResp         = pb.AddUserDynamicResp
	AddUserReq                 = pb.AddUserReq
	AddUserResp                = pb.AddUserResp
	AddUserShopReq             = pb.AddUserShopReq
	AddUserShopResp            = pb.AddUserShopResp
	AddUserSponsorReq          = pb.AddUserSponsorReq
	AddUserSponsorResp         = pb.AddUserSponsorResp
	CheckIsAdminReq            = pb.CheckIsAdminReq
	CheckIsAdminResp           = pb.CheckIsAdminResp
	DelUserAddressReq          = pb.DelUserAddressReq
	DelUserAddressResp         = pb.DelUserAddressResp
	DelUserAuthReq             = pb.DelUserAuthReq
	DelUserAuthResp            = pb.DelUserAuthResp
	DelUserContactReq          = pb.DelUserContactReq
	DelUserContactResp         = pb.DelUserContactResp
	DelUserDynamicReq          = pb.DelUserDynamicReq
	DelUserDynamicResp         = pb.DelUserDynamicResp
	DelUserReq                 = pb.DelUserReq
	DelUserResp                = pb.DelUserResp
	DelUserShopReq             = pb.DelUserShopReq
	DelUserShopResp            = pb.DelUserShopResp
	DelUserSponsorReq          = pb.DelUserSponsorReq
	DelUserSponsorResp         = pb.DelUserSponsorResp
	EditUserContactReq         = pb.EditUserContactReq
	EditUserContactResp        = pb.EditUserContactResp
	GenerateTokenReq           = pb.GenerateTokenReq
	GenerateTokenResp          = pb.GenerateTokenResp
	GetUserAddressByIdReq      = pb.GetUserAddressByIdReq
	GetUserAddressByIdResp     = pb.GetUserAddressByIdResp
	GetUserAuthByAuthKeyReq    = pb.GetUserAuthByAuthKeyReq
	GetUserAuthByAuthKeyResp   = pb.GetUserAuthByAuthKeyResp
	GetUserAuthByIdReq         = pb.GetUserAuthByIdReq
	GetUserAuthByIdResp        = pb.GetUserAuthByIdResp
	GetUserAuthByUserIdReq     = pb.GetUserAuthByUserIdReq
	GetUserAuthyUserIdResp     = pb.GetUserAuthyUserIdResp
	GetUserByIdReq             = pb.GetUserByIdReq
	GetUserByIdResp            = pb.GetUserByIdResp
	GetUserContactByIdReq      = pb.GetUserContactByIdReq
	GetUserContactByIdResp     = pb.GetUserContactByIdResp
	GetUserDynamicByIdReq      = pb.GetUserDynamicByIdReq
	GetUserDynamicByIdResp     = pb.GetUserDynamicByIdResp
	GetUserDynamicByUserIdReq  = pb.GetUserDynamicByUserIdReq
	GetUserDynamicByUserIdResp = pb.GetUserDynamicByUserIdResp
	GetUserInfoByUserIdsReq    = pb.GetUserInfoByUserIdsReq
	GetUserInfoByUserIdsResp   = pb.GetUserInfoByUserIdsResp
	GetUserInfoReq             = pb.GetUserInfoReq
	GetUserInfoResp            = pb.GetUserInfoResp
	GetUserShopByIdReq         = pb.GetUserShopByIdReq
	GetUserShopByIdResp        = pb.GetUserShopByIdResp
	GetUserSponsorByIdReq      = pb.GetUserSponsorByIdReq
	GetUserSponsorByIdResp     = pb.GetUserSponsorByIdResp
	LoginReq                   = pb.LoginReq
	LoginResp                  = pb.LoginResp
	RegisterReq                = pb.RegisterReq
	RegisterResp               = pb.RegisterResp
	SearchUserAddressReq       = pb.SearchUserAddressReq
	SearchUserAddressResp      = pb.SearchUserAddressResp
	SearchUserAuthReq          = pb.SearchUserAuthReq
	SearchUserAuthResp         = pb.SearchUserAuthResp
	SearchUserContactReq       = pb.SearchUserContactReq
	SearchUserContactResp      = pb.SearchUserContactResp
	SearchUserDynamicReq       = pb.SearchUserDynamicReq
	SearchUserDynamicResp      = pb.SearchUserDynamicResp
	SearchUserReq              = pb.SearchUserReq
	SearchUserResp             = pb.SearchUserResp
	SearchUserShopReq          = pb.SearchUserShopReq
	SearchUserShopResp         = pb.SearchUserShopResp
	SearchUserSponsorReq       = pb.SearchUserSponsorReq
	SearchUserSponsorResp      = pb.SearchUserSponsorResp
	SetAdminReq                = pb.SetAdminReq
	SetAdminResp               = pb.SetAdminResp
	SponsorDetailReq           = pb.SponsorDetailReq
	SponsorDetailResp          = pb.SponsorDetailResp
	UpdateUserAddressReq       = pb.UpdateUserAddressReq
	UpdateUserAddressResp      = pb.UpdateUserAddressResp
	UpdateUserAuthReq          = pb.UpdateUserAuthReq
	UpdateUserAuthResp         = pb.UpdateUserAuthResp
	UpdateUserBaseInfoReq      = pb.UpdateUserBaseInfoReq
	UpdateUserBaseInfoResp     = pb.UpdateUserBaseInfoResp
	UpdateUserContactReq       = pb.UpdateUserContactReq
	UpdateUserContactResp      = pb.UpdateUserContactResp
	UpdateUserDynamicReq       = pb.UpdateUserDynamicReq
	UpdateUserDynamicResp      = pb.UpdateUserDynamicResp
	UpdateUserReq              = pb.UpdateUserReq
	UpdateUserResp             = pb.UpdateUserResp
	UpdateUserShopReq          = pb.UpdateUserShopReq
	UpdateUserShopResp         = pb.UpdateUserShopResp
	UpdateUserSponsorReq       = pb.UpdateUserSponsorReq
	UpdateUserSponsorResp      = pb.UpdateUserSponsorResp
	User                       = pb.User
	UserAddress                = pb.UserAddress
	UserAuth                   = pb.UserAuth
	UserContact                = pb.UserContact
	UserDynamic                = pb.UserDynamic
	UserInfoForComment         = pb.UserInfoForComment
	UserShop                   = pb.UserShop
	UserSponsor                = pb.UserSponsor
	WXMiniRegisterReq          = pb.WXMiniRegisterReq
	WXMiniRegisterResp         = pb.WXMiniRegisterResp

	Usercenter interface {
		// 自定义的服务
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error)
		GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error)
		SetAdmin(ctx context.Context, in *SetAdminReq, opts ...grpc.CallOption) (*SetAdminResp, error)
		// -----------------------用户表-----------------------
		GetUserInfoByUserIds(ctx context.Context, in *GetUserInfoByUserIdsReq, opts ...grpc.CallOption) (*GetUserInfoByUserIdsResp, error)
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error)
		GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
		SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error)
		CheckIsAdmin(ctx context.Context, in *CheckIsAdminReq, opts ...grpc.CallOption) (*CheckIsAdminResp, error)
		WxMiniRegister(ctx context.Context, in *WXMiniRegisterReq, opts ...grpc.CallOption) (*WXMiniRegisterResp, error)
		// -----------------------用户收货地址表-----------------------
		AddUserAddress(ctx context.Context, in *AddUserAddressReq, opts ...grpc.CallOption) (*AddUserAddressResp, error)
		UpdateUserAddress(ctx context.Context, in *UpdateUserAddressReq, opts ...grpc.CallOption) (*UpdateUserAddressResp, error)
		DelUserAddress(ctx context.Context, in *DelUserAddressReq, opts ...grpc.CallOption) (*DelUserAddressResp, error)
		GetUserAddressById(ctx context.Context, in *GetUserAddressByIdReq, opts ...grpc.CallOption) (*GetUserAddressByIdResp, error)
		SearchUserAddress(ctx context.Context, in *SearchUserAddressReq, opts ...grpc.CallOption) (*SearchUserAddressResp, error)
		// -----------------------用户授权表-----------------------
		AddUserAuth(ctx context.Context, in *AddUserAuthReq, opts ...grpc.CallOption) (*AddUserAuthResp, error)
		UpdateUserAuth(ctx context.Context, in *UpdateUserAuthReq, opts ...grpc.CallOption) (*UpdateUserAuthResp, error)
		DelUserAuth(ctx context.Context, in *DelUserAuthReq, opts ...grpc.CallOption) (*DelUserAuthResp, error)
		GetUserAuthById(ctx context.Context, in *GetUserAuthByIdReq, opts ...grpc.CallOption) (*GetUserAuthByIdResp, error)
		SearchUserAuth(ctx context.Context, in *SearchUserAuthReq, opts ...grpc.CallOption) (*SearchUserAuthResp, error)
		// -----------------------抽奖发起人联系方式-----------------------
		AddUserContact(ctx context.Context, in *AddUserContactReq, opts ...grpc.CallOption) (*AddUserContactResp, error)
		EditUserContact(ctx context.Context, in *EditUserContactReq, opts ...grpc.CallOption) (*EditUserContactResp, error)
		UpdateUserContact(ctx context.Context, in *UpdateUserContactReq, opts ...grpc.CallOption) (*UpdateUserContactResp, error)
		DelUserContact(ctx context.Context, in *DelUserContactReq, opts ...grpc.CallOption) (*DelUserContactResp, error)
		GetUserContactById(ctx context.Context, in *GetUserContactByIdReq, opts ...grpc.CallOption) (*GetUserContactByIdResp, error)
		SearchUserContact(ctx context.Context, in *SearchUserContactReq, opts ...grpc.CallOption) (*SearchUserContactResp, error)
		// -----------------------userShop-----------------------
		AddUserShop(ctx context.Context, in *AddUserShopReq, opts ...grpc.CallOption) (*AddUserShopResp, error)
		UpdateUserShop(ctx context.Context, in *UpdateUserShopReq, opts ...grpc.CallOption) (*UpdateUserShopResp, error)
		DelUserShop(ctx context.Context, in *DelUserShopReq, opts ...grpc.CallOption) (*DelUserShopResp, error)
		GetUserShopById(ctx context.Context, in *GetUserShopByIdReq, opts ...grpc.CallOption) (*GetUserShopByIdResp, error)
		SearchUserShop(ctx context.Context, in *SearchUserShopReq, opts ...grpc.CallOption) (*SearchUserShopResp, error)
		// -----------------------抽奖发起人联系方式（抽奖赞助商）-----------------------
		AddUserSponsor(ctx context.Context, in *AddUserSponsorReq, opts ...grpc.CallOption) (*AddUserSponsorResp, error)
		UpdateUserSponsor(ctx context.Context, in *UpdateUserSponsorReq, opts ...grpc.CallOption) (*UpdateUserSponsorResp, error)
		DelUserSponsor(ctx context.Context, in *DelUserSponsorReq, opts ...grpc.CallOption) (*DelUserSponsorResp, error)
		GetUserSponsorById(ctx context.Context, in *GetUserSponsorByIdReq, opts ...grpc.CallOption) (*GetUserSponsorByIdResp, error)
		SearchUserSponsor(ctx context.Context, in *SearchUserSponsorReq, opts ...grpc.CallOption) (*SearchUserSponsorResp, error)
		SponsorDetail(ctx context.Context, in *SponsorDetailReq, opts ...grpc.CallOption) (*SponsorDetailResp, error)
		// -----------------------用户动态 UserDynamic-----------------------
		AddUserDynamic(ctx context.Context, in *AddUserDynamicReq, opts ...grpc.CallOption) (*AddUserDynamicResp, error)
		UpdateUserDynamic(ctx context.Context, in *UpdateUserDynamicReq, opts ...grpc.CallOption) (*UpdateUserDynamicResp, error)
		DelUserDynamic(ctx context.Context, in *DelUserDynamicReq, opts ...grpc.CallOption) (*DelUserDynamicResp, error)
		GetUserDynamicById(ctx context.Context, in *GetUserDynamicByIdReq, opts ...grpc.CallOption) (*GetUserDynamicByIdResp, error)
		SearchUserDynamic(ctx context.Context, in *SearchUserDynamicReq, opts ...grpc.CallOption) (*SearchUserDynamicResp, error)
		GetUserDynamicByUserId(ctx context.Context, in *GetUserDynamicByUserIdReq, opts ...grpc.CallOption) (*GetUserDynamicByUserIdResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

// 自定义的服务
func (m *defaultUsercenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByAuthKey(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByUserId(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserBaseInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) SetAdmin(ctx context.Context, in *SetAdminReq, opts ...grpc.CallOption) (*SetAdminResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SetAdmin(ctx, in, opts...)
}

// -----------------------用户表-----------------------
func (m *defaultUsercenter) GetUserInfoByUserIds(ctx context.Context, in *GetUserInfoByUserIdsReq, opts ...grpc.CallOption) (*GetUserInfoByUserIdsResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfoByUserIds(ctx, in, opts...)
}

func (m *defaultUsercenter) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUser(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUser(ctx, in, opts...)
}

func (m *defaultUsercenter) CheckIsAdmin(ctx context.Context, in *CheckIsAdminReq, opts ...grpc.CallOption) (*CheckIsAdminResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.CheckIsAdmin(ctx, in, opts...)
}

func (m *defaultUsercenter) WxMiniRegister(ctx context.Context, in *WXMiniRegisterReq, opts ...grpc.CallOption) (*WXMiniRegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.WxMiniRegister(ctx, in, opts...)
}

// -----------------------用户收货地址表-----------------------
func (m *defaultUsercenter) AddUserAddress(ctx context.Context, in *AddUserAddressReq, opts ...grpc.CallOption) (*AddUserAddressResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUserAddress(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserAddress(ctx context.Context, in *UpdateUserAddressReq, opts ...grpc.CallOption) (*UpdateUserAddressResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserAddress(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUserAddress(ctx context.Context, in *DelUserAddressReq, opts ...grpc.CallOption) (*DelUserAddressResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUserAddress(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAddressById(ctx context.Context, in *GetUserAddressByIdReq, opts ...grpc.CallOption) (*GetUserAddressByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAddressById(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUserAddress(ctx context.Context, in *SearchUserAddressReq, opts ...grpc.CallOption) (*SearchUserAddressResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUserAddress(ctx, in, opts...)
}

// -----------------------用户授权表-----------------------
func (m *defaultUsercenter) AddUserAuth(ctx context.Context, in *AddUserAuthReq, opts ...grpc.CallOption) (*AddUserAuthResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUserAuth(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserAuth(ctx context.Context, in *UpdateUserAuthReq, opts ...grpc.CallOption) (*UpdateUserAuthResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserAuth(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUserAuth(ctx context.Context, in *DelUserAuthReq, opts ...grpc.CallOption) (*DelUserAuthResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUserAuth(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthById(ctx context.Context, in *GetUserAuthByIdReq, opts ...grpc.CallOption) (*GetUserAuthByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthById(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUserAuth(ctx context.Context, in *SearchUserAuthReq, opts ...grpc.CallOption) (*SearchUserAuthResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUserAuth(ctx, in, opts...)
}

// -----------------------抽奖发起人联系方式-----------------------
func (m *defaultUsercenter) AddUserContact(ctx context.Context, in *AddUserContactReq, opts ...grpc.CallOption) (*AddUserContactResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUserContact(ctx, in, opts...)
}

func (m *defaultUsercenter) EditUserContact(ctx context.Context, in *EditUserContactReq, opts ...grpc.CallOption) (*EditUserContactResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.EditUserContact(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserContact(ctx context.Context, in *UpdateUserContactReq, opts ...grpc.CallOption) (*UpdateUserContactResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserContact(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUserContact(ctx context.Context, in *DelUserContactReq, opts ...grpc.CallOption) (*DelUserContactResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUserContact(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserContactById(ctx context.Context, in *GetUserContactByIdReq, opts ...grpc.CallOption) (*GetUserContactByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserContactById(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUserContact(ctx context.Context, in *SearchUserContactReq, opts ...grpc.CallOption) (*SearchUserContactResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUserContact(ctx, in, opts...)
}

// -----------------------userShop-----------------------
func (m *defaultUsercenter) AddUserShop(ctx context.Context, in *AddUserShopReq, opts ...grpc.CallOption) (*AddUserShopResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUserShop(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserShop(ctx context.Context, in *UpdateUserShopReq, opts ...grpc.CallOption) (*UpdateUserShopResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserShop(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUserShop(ctx context.Context, in *DelUserShopReq, opts ...grpc.CallOption) (*DelUserShopResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUserShop(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserShopById(ctx context.Context, in *GetUserShopByIdReq, opts ...grpc.CallOption) (*GetUserShopByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserShopById(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUserShop(ctx context.Context, in *SearchUserShopReq, opts ...grpc.CallOption) (*SearchUserShopResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUserShop(ctx, in, opts...)
}

// -----------------------抽奖发起人联系方式（抽奖赞助商）-----------------------
func (m *defaultUsercenter) AddUserSponsor(ctx context.Context, in *AddUserSponsorReq, opts ...grpc.CallOption) (*AddUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserSponsor(ctx context.Context, in *UpdateUserSponsorReq, opts ...grpc.CallOption) (*UpdateUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUserSponsor(ctx context.Context, in *DelUserSponsorReq, opts ...grpc.CallOption) (*DelUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserSponsorById(ctx context.Context, in *GetUserSponsorByIdReq, opts ...grpc.CallOption) (*GetUserSponsorByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserSponsorById(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUserSponsor(ctx context.Context, in *SearchUserSponsorReq, opts ...grpc.CallOption) (*SearchUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) SponsorDetail(ctx context.Context, in *SponsorDetailReq, opts ...grpc.CallOption) (*SponsorDetailResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SponsorDetail(ctx, in, opts...)
}

// -----------------------用户动态 UserDynamic-----------------------
func (m *defaultUsercenter) AddUserDynamic(ctx context.Context, in *AddUserDynamicReq, opts ...grpc.CallOption) (*AddUserDynamicResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUserDynamic(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserDynamic(ctx context.Context, in *UpdateUserDynamicReq, opts ...grpc.CallOption) (*UpdateUserDynamicResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserDynamic(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUserDynamic(ctx context.Context, in *DelUserDynamicReq, opts ...grpc.CallOption) (*DelUserDynamicResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUserDynamic(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserDynamicById(ctx context.Context, in *GetUserDynamicByIdReq, opts ...grpc.CallOption) (*GetUserDynamicByIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserDynamicById(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUserDynamic(ctx context.Context, in *SearchUserDynamicReq, opts ...grpc.CallOption) (*SearchUserDynamicResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUserDynamic(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserDynamicByUserId(ctx context.Context, in *GetUserDynamicByUserIdReq, opts ...grpc.CallOption) (*GetUserDynamicByUserIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserDynamicByUserId(ctx, in, opts...)
}
