// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: keyprovider.proto
package attestation_agent

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
)

type KeyProviderServiceService interface {
	WrapKey(context.Context, *KeyProviderKeyWrapProtocolInput) (*KeyProviderKeyWrapProtocolOutput, error)
	UnWrapKey(context.Context, *KeyProviderKeyWrapProtocolInput) (*KeyProviderKeyWrapProtocolOutput, error)
}

func RegisterKeyProviderServiceService(srv *ttrpc.Server, svc KeyProviderServiceService) {
	srv.RegisterService("keyprovider.KeyProviderService", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"WrapKey": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req KeyProviderKeyWrapProtocolInput
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.WrapKey(ctx, &req)
			},
			"UnWrapKey": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req KeyProviderKeyWrapProtocolInput
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.UnWrapKey(ctx, &req)
			},
		},
	})
}

type keyproviderserviceClient struct {
	client *ttrpc.Client
}

func NewKeyProviderServiceClient(client *ttrpc.Client) KeyProviderServiceService {
	return &keyproviderserviceClient{
		client: client,
	}
}

func (c *keyproviderserviceClient) WrapKey(ctx context.Context, req *KeyProviderKeyWrapProtocolInput) (*KeyProviderKeyWrapProtocolOutput, error) {
	var resp KeyProviderKeyWrapProtocolOutput
	if err := c.client.Call(ctx, "keyprovider.KeyProviderService", "WrapKey", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *keyproviderserviceClient) UnWrapKey(ctx context.Context, req *KeyProviderKeyWrapProtocolInput) (*KeyProviderKeyWrapProtocolOutput, error) {
	var resp KeyProviderKeyWrapProtocolOutput
	if err := c.client.Call(ctx, "keyprovider.KeyProviderService", "UnWrapKey", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
