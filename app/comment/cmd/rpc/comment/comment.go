// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package comment

import (
	"context"

	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCommentReq        = pb.AddCommentReq
	AddCommentResp       = pb.AddCommentResp
	AddPraiseReq         = pb.AddPraiseReq
	AddPraiseResp        = pb.AddPraiseResp
	CheckUserPraiseReq   = pb.CheckUserPraiseReq
	CheckUserPraiseResp  = pb.CheckUserPraiseResp
	Comment              = pb.Comment
	DelCommentReq        = pb.DelCommentReq
	DelCommentResp       = pb.DelCommentResp
	DelPraiseReq         = pb.DelPraiseReq
	DelPraiseResp        = pb.DelPraiseResp
	GetCommentByIdReq    = pb.GetCommentByIdReq
	GetCommentByIdResp   = pb.GetCommentByIdResp
	GetCommentLastIdReq  = pb.GetCommentLastIdReq
	GetCommentLastIdResp = pb.GetCommentLastIdResp
	GetPraiseByIdReq     = pb.GetPraiseByIdReq
	GetPraiseByIdResp    = pb.GetPraiseByIdResp
	GetUserCommentReq    = pb.GetUserCommentReq
	GetUserCommentResp   = pb.GetUserCommentResp
	IsPraiseListReq      = pb.IsPraiseListReq
	IsPraiseListResp     = pb.IsPraiseListResp
	IsPraiseReq          = pb.IsPraiseReq
	IsPraiseResp         = pb.IsPraiseResp
	Praise               = pb.Praise
	PraiseCommentReq     = pb.PraiseCommentReq
	PraiseCommentResp    = pb.PraiseCommentResp
	SearchCommentReq     = pb.SearchCommentReq
	SearchCommentResp    = pb.SearchCommentResp
	SearchPraiseReq      = pb.SearchPraiseReq
	SearchPraiseResp     = pb.SearchPraiseResp
	UpdateCommentReq     = pb.UpdateCommentReq
	UpdateCommentResp    = pb.UpdateCommentResp
	UpdatePraiseReq      = pb.UpdatePraiseReq
	UpdatePraiseResp     = pb.UpdatePraiseResp

	CommentZrpcClient interface {
		// -----------------------comment-----------------------
		AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
		UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error)
		DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error)
		GetCommentById(ctx context.Context, in *GetCommentByIdReq, opts ...grpc.CallOption) (*GetCommentByIdResp, error)
		SearchComment(ctx context.Context, in *SearchCommentReq, opts ...grpc.CallOption) (*SearchCommentResp, error)
		IsPraise(ctx context.Context, in *IsPraiseReq, opts ...grpc.CallOption) (*IsPraiseResp, error)
		PraiseComment(ctx context.Context, in *PraiseCommentReq, opts ...grpc.CallOption) (*PraiseCommentResp, error)
		GetCommentLastId(ctx context.Context, in *GetCommentLastIdReq, opts ...grpc.CallOption) (*GetCommentLastIdResp, error)
		IsPraiseList(ctx context.Context, in *IsPraiseListReq, opts ...grpc.CallOption) (*IsPraiseListResp, error)
		GetUserComment(ctx context.Context, in *GetUserCommentReq, opts ...grpc.CallOption) (*GetUserCommentResp, error)
		// -----------------------praise-----------------------
		AddPraise(ctx context.Context, in *AddPraiseReq, opts ...grpc.CallOption) (*AddPraiseResp, error)
		UpdatePraise(ctx context.Context, in *UpdatePraiseReq, opts ...grpc.CallOption) (*UpdatePraiseResp, error)
		DelPraise(ctx context.Context, in *DelPraiseReq, opts ...grpc.CallOption) (*DelPraiseResp, error)
		GetPraiseById(ctx context.Context, in *GetPraiseByIdReq, opts ...grpc.CallOption) (*GetPraiseByIdResp, error)
		SearchPraise(ctx context.Context, in *SearchPraiseReq, opts ...grpc.CallOption) (*SearchPraiseResp, error)
		// -----------------------others-----------------------
		CheckUserPraise(ctx context.Context, in *CheckUserPraiseReq, opts ...grpc.CallOption) (*CheckUserPraiseResp, error)
	}

	defaultCommentZrpcClient struct {
		cli zrpc.Client
	}
)

func NewCommentZrpcClient(cli zrpc.Client) CommentZrpcClient {
	return &defaultCommentZrpcClient{
		cli: cli,
	}
}

// -----------------------comment-----------------------
func (m *defaultCommentZrpcClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.AddComment(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.UpdateComment(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.DelComment(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) GetCommentById(ctx context.Context, in *GetCommentByIdReq, opts ...grpc.CallOption) (*GetCommentByIdResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.GetCommentById(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) SearchComment(ctx context.Context, in *SearchCommentReq, opts ...grpc.CallOption) (*SearchCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.SearchComment(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) IsPraise(ctx context.Context, in *IsPraiseReq, opts ...grpc.CallOption) (*IsPraiseResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.IsPraise(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) PraiseComment(ctx context.Context, in *PraiseCommentReq, opts ...grpc.CallOption) (*PraiseCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.PraiseComment(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) GetCommentLastId(ctx context.Context, in *GetCommentLastIdReq, opts ...grpc.CallOption) (*GetCommentLastIdResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.GetCommentLastId(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) IsPraiseList(ctx context.Context, in *IsPraiseListReq, opts ...grpc.CallOption) (*IsPraiseListResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.IsPraiseList(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) GetUserComment(ctx context.Context, in *GetUserCommentReq, opts ...grpc.CallOption) (*GetUserCommentResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.GetUserComment(ctx, in, opts...)
}

// -----------------------praise-----------------------
func (m *defaultCommentZrpcClient) AddPraise(ctx context.Context, in *AddPraiseReq, opts ...grpc.CallOption) (*AddPraiseResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.AddPraise(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) UpdatePraise(ctx context.Context, in *UpdatePraiseReq, opts ...grpc.CallOption) (*UpdatePraiseResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.UpdatePraise(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) DelPraise(ctx context.Context, in *DelPraiseReq, opts ...grpc.CallOption) (*DelPraiseResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.DelPraise(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) GetPraiseById(ctx context.Context, in *GetPraiseByIdReq, opts ...grpc.CallOption) (*GetPraiseByIdResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.GetPraiseById(ctx, in, opts...)
}

func (m *defaultCommentZrpcClient) SearchPraise(ctx context.Context, in *SearchPraiseReq, opts ...grpc.CallOption) (*SearchPraiseResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.SearchPraise(ctx, in, opts...)
}

// -----------------------others-----------------------
func (m *defaultCommentZrpcClient) CheckUserPraise(ctx context.Context, in *CheckUserPraiseReq, opts ...grpc.CallOption) (*CheckUserPraiseResp, error) {
	client := pb.NewCommentClient(m.cli.Conn())
	return client.CheckUserPraise(ctx, in, opts...)
}
