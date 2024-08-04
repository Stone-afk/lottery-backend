// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: comment.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
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

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error) {
	out := new(UpdateCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error) {
	out := new(DelCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/DelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetCommentById(ctx context.Context, in *GetCommentByIdReq, opts ...grpc.CallOption) (*GetCommentByIdResp, error) {
	out := new(GetCommentByIdResp)
	err := c.cc.Invoke(ctx, "/pb.comment/GetCommentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) SearchComment(ctx context.Context, in *SearchCommentReq, opts ...grpc.CallOption) (*SearchCommentResp, error) {
	out := new(SearchCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/SearchComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) IsPraise(ctx context.Context, in *IsPraiseReq, opts ...grpc.CallOption) (*IsPraiseResp, error) {
	out := new(IsPraiseResp)
	err := c.cc.Invoke(ctx, "/pb.comment/IsPraise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) PraiseComment(ctx context.Context, in *PraiseCommentReq, opts ...grpc.CallOption) (*PraiseCommentResp, error) {
	out := new(PraiseCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/PraiseComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetCommentLastId(ctx context.Context, in *GetCommentLastIdReq, opts ...grpc.CallOption) (*GetCommentLastIdResp, error) {
	out := new(GetCommentLastIdResp)
	err := c.cc.Invoke(ctx, "/pb.comment/GetCommentLastId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) IsPraiseList(ctx context.Context, in *IsPraiseListReq, opts ...grpc.CallOption) (*IsPraiseListResp, error) {
	out := new(IsPraiseListResp)
	err := c.cc.Invoke(ctx, "/pb.comment/IsPraiseList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetUserComment(ctx context.Context, in *GetUserCommentReq, opts ...grpc.CallOption) (*GetUserCommentResp, error) {
	out := new(GetUserCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/GetUserComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) AddPraise(ctx context.Context, in *AddPraiseReq, opts ...grpc.CallOption) (*AddPraiseResp, error) {
	out := new(AddPraiseResp)
	err := c.cc.Invoke(ctx, "/pb.comment/AddPraise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) UpdatePraise(ctx context.Context, in *UpdatePraiseReq, opts ...grpc.CallOption) (*UpdatePraiseResp, error) {
	out := new(UpdatePraiseResp)
	err := c.cc.Invoke(ctx, "/pb.comment/UpdatePraise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) DelPraise(ctx context.Context, in *DelPraiseReq, opts ...grpc.CallOption) (*DelPraiseResp, error) {
	out := new(DelPraiseResp)
	err := c.cc.Invoke(ctx, "/pb.comment/DelPraise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetPraiseById(ctx context.Context, in *GetPraiseByIdReq, opts ...grpc.CallOption) (*GetPraiseByIdResp, error) {
	out := new(GetPraiseByIdResp)
	err := c.cc.Invoke(ctx, "/pb.comment/GetPraiseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) SearchPraise(ctx context.Context, in *SearchPraiseReq, opts ...grpc.CallOption) (*SearchPraiseResp, error) {
	out := new(SearchPraiseResp)
	err := c.cc.Invoke(ctx, "/pb.comment/SearchPraise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) CheckUserPraise(ctx context.Context, in *CheckUserPraiseReq, opts ...grpc.CallOption) (*CheckUserPraiseResp, error) {
	out := new(CheckUserPraiseResp)
	err := c.cc.Invoke(ctx, "/pb.comment/CheckUserPraise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	// -----------------------comment-----------------------
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error)
	DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error)
	GetCommentById(context.Context, *GetCommentByIdReq) (*GetCommentByIdResp, error)
	SearchComment(context.Context, *SearchCommentReq) (*SearchCommentResp, error)
	IsPraise(context.Context, *IsPraiseReq) (*IsPraiseResp, error)
	PraiseComment(context.Context, *PraiseCommentReq) (*PraiseCommentResp, error)
	GetCommentLastId(context.Context, *GetCommentLastIdReq) (*GetCommentLastIdResp, error)
	IsPraiseList(context.Context, *IsPraiseListReq) (*IsPraiseListResp, error)
	GetUserComment(context.Context, *GetUserCommentReq) (*GetUserCommentResp, error)
	// -----------------------praise-----------------------
	AddPraise(context.Context, *AddPraiseReq) (*AddPraiseResp, error)
	UpdatePraise(context.Context, *UpdatePraiseReq) (*UpdatePraiseResp, error)
	DelPraise(context.Context, *DelPraiseReq) (*DelPraiseResp, error)
	GetPraiseById(context.Context, *GetPraiseByIdReq) (*GetPraiseByIdResp, error)
	SearchPraise(context.Context, *SearchPraiseReq) (*SearchPraiseResp, error)
	// -----------------------others-----------------------
	CheckUserPraise(context.Context, *CheckUserPraiseReq) (*CheckUserPraiseResp, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCommentServer) UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentServer) DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelComment not implemented")
}
func (UnimplementedCommentServer) GetCommentById(context.Context, *GetCommentByIdReq) (*GetCommentByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedCommentServer) SearchComment(context.Context, *SearchCommentReq) (*SearchCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchComment not implemented")
}
func (UnimplementedCommentServer) IsPraise(context.Context, *IsPraiseReq) (*IsPraiseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPraise not implemented")
}
func (UnimplementedCommentServer) PraiseComment(context.Context, *PraiseCommentReq) (*PraiseCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PraiseComment not implemented")
}
func (UnimplementedCommentServer) GetCommentLastId(context.Context, *GetCommentLastIdReq) (*GetCommentLastIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentLastId not implemented")
}
func (UnimplementedCommentServer) IsPraiseList(context.Context, *IsPraiseListReq) (*IsPraiseListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPraiseList not implemented")
}
func (UnimplementedCommentServer) GetUserComment(context.Context, *GetUserCommentReq) (*GetUserCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserComment not implemented")
}
func (UnimplementedCommentServer) AddPraise(context.Context, *AddPraiseReq) (*AddPraiseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPraise not implemented")
}
func (UnimplementedCommentServer) UpdatePraise(context.Context, *UpdatePraiseReq) (*UpdatePraiseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePraise not implemented")
}
func (UnimplementedCommentServer) DelPraise(context.Context, *DelPraiseReq) (*DelPraiseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelPraise not implemented")
}
func (UnimplementedCommentServer) GetPraiseById(context.Context, *GetPraiseByIdReq) (*GetPraiseByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPraiseById not implemented")
}
func (UnimplementedCommentServer) SearchPraise(context.Context, *SearchPraiseReq) (*SearchPraiseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPraise not implemented")
}
func (UnimplementedCommentServer) CheckUserPraise(context.Context, *CheckUserPraiseReq) (*CheckUserPraiseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserPraise not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).UpdateComment(ctx, req.(*UpdateCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_DelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).DelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/DelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).DelComment(ctx, req.(*DelCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/GetCommentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetCommentById(ctx, req.(*GetCommentByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_SearchComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).SearchComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/SearchComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).SearchComment(ctx, req.(*SearchCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_IsPraise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPraiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).IsPraise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/IsPraise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).IsPraise(ctx, req.(*IsPraiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_PraiseComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PraiseCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).PraiseComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/PraiseComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).PraiseComment(ctx, req.(*PraiseCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetCommentLastId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentLastIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetCommentLastId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/GetCommentLastId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetCommentLastId(ctx, req.(*GetCommentLastIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_IsPraiseList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPraiseListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).IsPraiseList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/IsPraiseList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).IsPraiseList(ctx, req.(*IsPraiseListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetUserComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetUserComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/GetUserComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetUserComment(ctx, req.(*GetUserCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_AddPraise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPraiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).AddPraise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/AddPraise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).AddPraise(ctx, req.(*AddPraiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_UpdatePraise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePraiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).UpdatePraise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/UpdatePraise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).UpdatePraise(ctx, req.(*UpdatePraiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_DelPraise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelPraiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).DelPraise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/DelPraise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).DelPraise(ctx, req.(*DelPraiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetPraiseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPraiseByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetPraiseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/GetPraiseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetPraiseById(ctx, req.(*GetPraiseByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_SearchPraise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPraiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).SearchPraise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/SearchPraise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).SearchPraise(ctx, req.(*SearchPraiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_CheckUserPraise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserPraiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).CheckUserPraise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/CheckUserPraise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).CheckUserPraise(ctx, req.(*CheckUserPraiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _Comment_AddComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Comment_UpdateComment_Handler,
		},
		{
			MethodName: "DelComment",
			Handler:    _Comment_DelComment_Handler,
		},
		{
			MethodName: "GetCommentById",
			Handler:    _Comment_GetCommentById_Handler,
		},
		{
			MethodName: "SearchComment",
			Handler:    _Comment_SearchComment_Handler,
		},
		{
			MethodName: "IsPraise",
			Handler:    _Comment_IsPraise_Handler,
		},
		{
			MethodName: "PraiseComment",
			Handler:    _Comment_PraiseComment_Handler,
		},
		{
			MethodName: "GetCommentLastId",
			Handler:    _Comment_GetCommentLastId_Handler,
		},
		{
			MethodName: "IsPraiseList",
			Handler:    _Comment_IsPraiseList_Handler,
		},
		{
			MethodName: "GetUserComment",
			Handler:    _Comment_GetUserComment_Handler,
		},
		{
			MethodName: "AddPraise",
			Handler:    _Comment_AddPraise_Handler,
		},
		{
			MethodName: "UpdatePraise",
			Handler:    _Comment_UpdatePraise_Handler,
		},
		{
			MethodName: "DelPraise",
			Handler:    _Comment_DelPraise_Handler,
		},
		{
			MethodName: "GetPraiseById",
			Handler:    _Comment_GetPraiseById_Handler,
		},
		{
			MethodName: "SearchPraise",
			Handler:    _Comment_SearchPraise_Handler,
		},
		{
			MethodName: "CheckUserPraise",
			Handler:    _Comment_CheckUserPraise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
