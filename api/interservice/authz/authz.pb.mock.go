// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: api/interservice/authz/authz.proto

package authz

import (
	"context"

	version "github.com/chef/automate/api/external/common/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the AuthorizationServer interface (at compile time)
var _ AuthorizationServer = &AuthorizationServerMock{}

// NewAuthorizationServerMock gives you a fresh instance of AuthorizationServerMock.
func NewAuthorizationServerMock() *AuthorizationServerMock {
	return &AuthorizationServerMock{validateRequests: true}
}

// NewAuthorizationServerMockWithoutValidation gives you a fresh instance of
// AuthorizationServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewAuthorizationServerMockWithoutValidation() *AuthorizationServerMock {
	return &AuthorizationServerMock{}
}

// AuthorizationServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type AuthorizationServerMock struct {
	validateRequests          bool
	GetVersionFunc            func(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	IsAuthorizedFunc          func(context.Context, *IsAuthorizedReq) (*IsAuthorizedResp, error)
	FilterAuthorizedPairsFunc func(context.Context, *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error)
}

func (m *AuthorizationServerMock) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetVersionFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetVersion' not implemented")
}

func (m *AuthorizationServerMock) IsAuthorized(ctx context.Context, req *IsAuthorizedReq) (*IsAuthorizedResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.IsAuthorizedFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'IsAuthorized' not implemented")
}

func (m *AuthorizationServerMock) FilterAuthorizedPairs(ctx context.Context, req *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.FilterAuthorizedPairsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'FilterAuthorizedPairs' not implemented")
}

// Reset resets all overridden functions
func (m *AuthorizationServerMock) Reset() {
	m.GetVersionFunc = nil
	m.IsAuthorizedFunc = nil
	m.FilterAuthorizedPairsFunc = nil
}
