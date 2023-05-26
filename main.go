package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/containerd/ttrpc"

	"github.com/katexochen/coco-ttrpc-cli/attestation_agent"
)

func main() {
	if err := handle(); err != nil {
		log.Fatal(err)
	}
}

func handle() error {
	system := os.Args[1]
	switch system {
	case "attestation-agent":
		role := os.Args[2]
		switch role {
		case "server":
			return startAttestationAgentServer()
		case "client":
			return attestationAgentClient()
		default:
			return errors.New("invalid role")
		}
	default:
		return errors.New("invalid system")
	}
}

const socket = "example-ttrpc-server"

func startAttestationAgentServer() error {
	s, err := ttrpc.NewServer(ttrpc.WithServerHandshaker(ttrpc.UnixSocketRequireSameUser()))
	if err != nil {
		return err
	}
	defer s.Close()
	attestation_agent.RegisterGetResourceServiceService(s, &attestationAgentServer{})
	attestation_agent.RegisterKeyProviderServiceService(s, &attestationAgentServer{})

	l, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}
	defer func() {
		l.Close()
		os.Remove(socket)
	}()
	return s.Serve(context.Background(), l)
}

func attestationAgentClient() error {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	defer conn.Close()

	tc := ttrpc.NewClient(conn)
	getresourceClient := attestation_agent.NewGetResourceServiceClient(tc)
	keyproviderClient := attestation_agent.NewKeyProviderServiceClient(tc)

	method := os.Args[3]
	payload := os.Args[4]
	switch method {
	case "GetResource":
		r := &attestation_agent.GetResourceRequest{}
		if err := json.Unmarshal([]byte(payload), r); err != nil {
			return err
		}
		resp, err := getresourceClient.GetResource(context.Background(), r)
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(resp)
	case "WrapKey":
		r := &attestation_agent.KeyProviderKeyWrapProtocolInput{}
		if err := json.Unmarshal([]byte(payload), r); err != nil {
			return err
		}
		resp, err := keyproviderClient.WrapKey(context.Background(), r)
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(resp)
	case "UnWrapKey":
		r := &attestation_agent.KeyProviderKeyWrapProtocolInput{}
		if err := json.Unmarshal([]byte(payload), r); err != nil {
			return err
		}
		resp, err := keyproviderClient.UnWrapKey(context.Background(), r)
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(resp)
	}

	return nil
}

type attestationAgentServer struct{}

func (s *attestationAgentServer) WrapKey(ctx context.Context, r *attestation_agent.KeyProviderKeyWrapProtocolInput) (*attestation_agent.KeyProviderKeyWrapProtocolOutput, error) {
	fmt.Printf("receive wrap key request: %v\n", r)
	return &attestation_agent.KeyProviderKeyWrapProtocolOutput{}, nil
}

func (s *attestationAgentServer) UnWrapKey(ctx context.Context, r *attestation_agent.KeyProviderKeyWrapProtocolInput) (*attestation_agent.KeyProviderKeyWrapProtocolOutput, error) {
	fmt.Printf("receive unwrap key request: %v\n", r)
	return &attestation_agent.KeyProviderKeyWrapProtocolOutput{}, nil
}

func (s *attestationAgentServer) GetResource(ctx context.Context, r *attestation_agent.GetResourceRequest) (*attestation_agent.GetResourceResponse, error) {
	fmt.Printf("receive get resource request: %v\n", r)
	return &attestation_agent.GetResourceResponse{}, nil
}
