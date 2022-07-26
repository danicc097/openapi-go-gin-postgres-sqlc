// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/python-ml-app-protos/tfidf/v1/service.proto

package v1

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

// MovieGenreClient is the client API for MovieGenre service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieGenreClient interface {
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictReply, error)
	Train(ctx context.Context, in *TrainRequest, opts ...grpc.CallOption) (*TrainReply, error)
}

type movieGenreClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieGenreClient(cc grpc.ClientConnInterface) MovieGenreClient {
	return &movieGenreClient{cc}
}

func (c *movieGenreClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictReply, error) {
	out := new(PredictReply)
	err := c.cc.Invoke(ctx, "/tfidf.MovieGenre/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieGenreClient) Train(ctx context.Context, in *TrainRequest, opts ...grpc.CallOption) (*TrainReply, error) {
	out := new(TrainReply)
	err := c.cc.Invoke(ctx, "/tfidf.MovieGenre/Train", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieGenreServer is the server API for MovieGenre service.
// All implementations must embed UnimplementedMovieGenreServer
// for forward compatibility
type MovieGenreServer interface {
	Predict(context.Context, *PredictRequest) (*PredictReply, error)
	Train(context.Context, *TrainRequest) (*TrainReply, error)
	mustEmbedUnimplementedMovieGenreServer()
}

// UnimplementedMovieGenreServer must be embedded to have forward compatible implementations.
type UnimplementedMovieGenreServer struct {
}

func (UnimplementedMovieGenreServer) Predict(context.Context, *PredictRequest) (*PredictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (UnimplementedMovieGenreServer) Train(context.Context, *TrainRequest) (*TrainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Train not implemented")
}
func (UnimplementedMovieGenreServer) mustEmbedUnimplementedMovieGenreServer() {}

// UnsafeMovieGenreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieGenreServer will
// result in compilation errors.
type UnsafeMovieGenreServer interface {
	mustEmbedUnimplementedMovieGenreServer()
}

func RegisterMovieGenreServer(s grpc.ServiceRegistrar, srv MovieGenreServer) {
	s.RegisterService(&MovieGenre_ServiceDesc, srv)
}

func _MovieGenre_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieGenreServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfidf.MovieGenre/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieGenreServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieGenre_Train_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieGenreServer).Train(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfidf.MovieGenre/Train",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieGenreServer).Train(ctx, req.(*TrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieGenre_ServiceDesc is the grpc.ServiceDesc for MovieGenre service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieGenre_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tfidf.MovieGenre",
	HandlerType: (*MovieGenreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _MovieGenre_Predict_Handler,
		},
		{
			MethodName: "Train",
			Handler:    _MovieGenre_Train_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/python-ml-app-protos/tfidf/v1/service.proto",
}
