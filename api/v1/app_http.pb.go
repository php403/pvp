// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// context./http.
const _ = http1.SupportPackageIsVersion1

type AppHTTPServer interface {
	DeleteApp(context.Context, *DeleteAppRequest) (*DeleteAppReply, error)

	GetApp(context.Context, *GetAppRequest) (*GetAppReply, error)

	ListApp(context.Context, *ListAppRequest) (*ListAppReply, error)

	LoginApp(context.Context, *LoginAppRequest) (*LoginAppReply, error)

	RegisterApp(context.Context, *RegisterAppRequest) (*RegisterAppReply, error)

	UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppReply, error)
}

func RegisterAppHTTPServer(s http1.ServiceRegistrar, srv AppHTTPServer) {
	s.RegisterService(&_HTTP_App_serviceDesc, srv)
}

func _HTTP_App_RegisterApp_0(srv interface{}, ctx context.Context, req *http.Request) (interface{}, error) {
	var in RegisterAppRequest

	if err := http1.PopulateForm(&in, req); err != nil {
		return nil, err
	}

	out, err := srv.(AppServer).RegisterApp(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_App_UpdateApp_0(srv interface{}, ctx context.Context, req *http.Request) (interface{}, error) {
	var in UpdateAppRequest

	if err := http1.PopulateForm(&in, req); err != nil {
		return nil, err
	}

	out, err := srv.(AppServer).UpdateApp(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_App_DeleteApp_0(srv interface{}, ctx context.Context, req *http.Request) (interface{}, error) {
	var in DeleteAppRequest

	if err := http1.PopulateForm(&in, req); err != nil {
		return nil, err
	}

	out, err := srv.(AppServer).DeleteApp(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_App_GetApp_0(srv interface{}, ctx context.Context, req *http.Request) (interface{}, error) {
	var in GetAppRequest

	if err := http1.PopulateForm(&in, req); err != nil {
		return nil, err
	}

	out, err := srv.(AppServer).GetApp(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_App_LoginApp_0(srv interface{}, ctx context.Context, req *http.Request) (interface{}, error) {
	var in LoginAppRequest

	if err := http1.PopulateForm(&in, req); err != nil {
		return nil, err
	}

	out, err := srv.(AppServer).LoginApp(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _HTTP_App_ListApp_0(srv interface{}, ctx context.Context, req *http.Request) (interface{}, error) {
	var in ListAppRequest

	if err := http1.PopulateForm(&in, req); err != nil {
		return nil, err
	}

	out, err := srv.(AppServer).ListApp(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _HTTP_App_serviceDesc = http1.ServiceDesc{
	ServiceName: "api.v1.App",
	HandlerType: (*AppHTTPServer)(nil),
	Methods: []http1.MethodDesc{

		{
			Path:    "/api.v1.App/RegisterApp",
			Method:  "POST",
			Handler: _HTTP_App_RegisterApp_0,
		},

		{
			Path:    "/api.v1.App/UpdateApp",
			Method:  "POST",
			Handler: _HTTP_App_UpdateApp_0,
		},

		{
			Path:    "/api.v1.App/DeleteApp",
			Method:  "POST",
			Handler: _HTTP_App_DeleteApp_0,
		},

		{
			Path:    "/api.v1.App/GetApp",
			Method:  "POST",
			Handler: _HTTP_App_GetApp_0,
		},

		{
			Path:    "/api.v1.App/LoginApp",
			Method:  "POST",
			Handler: _HTTP_App_LoginApp_0,
		},

		{
			Path:    "/api.v1.App/ListApp",
			Method:  "POST",
			Handler: _HTTP_App_ListApp_0,
		},
	},
	Metadata: "api/v1/app.proto",
}
