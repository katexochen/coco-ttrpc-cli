// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: getresource.proto
package attestation_agent

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
)

type GetResourceServiceService interface {
	GetResource(context.Context, *GetResourceRequest) (*GetResourceResponse, error)
}

func RegisterGetResourceServiceService(srv *ttrpc.Server, svc GetResourceServiceService) {
	srv.RegisterService("getresource.GetResourceService", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"GetResource": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req GetResourceRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.GetResource(ctx, &req)
			},
		},
	})
}

type getresourceserviceClient struct {
	client *ttrpc.Client
}

func NewGetResourceServiceClient(client *ttrpc.Client) GetResourceServiceService {
	return &getresourceserviceClient{
		client: client,
	}
}

func (c *getresourceserviceClient) GetResource(ctx context.Context, req *GetResourceRequest) (*GetResourceResponse, error) {
	var resp GetResourceResponse
	if err := c.client.Call(ctx, "getresource.GetResourceService", "GetResource", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
