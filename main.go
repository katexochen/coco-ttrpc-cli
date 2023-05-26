package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
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
	socket := flag.String("socket", "example-ttrpc-server", "socket path")
	flag.Parse()
	system := flag.Arg(0)
	role := flag.Arg(1)
	fmt.Printf("system: %s, role: %s, socket: %s\n", system, role, *socket)

	if flag.NArg() < 2 {
		return errors.New("invalid args. usage: <system> <role> [method] [payload]")
	}
	switch system {
	case "attestation-agent":
		switch role {
		case "server":
			return startAttestationAgentServer(*socket)
		case "client":
			return attestationAgentClient(*socket)
		default:
			return errors.New("invalid role")
		}
	default:
		return errors.New("invalid system")
	}
}

func startAttestationAgentServer(socket string) error {
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

func attestationAgentClient(socket string) error {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	defer conn.Close()

	tc := ttrpc.NewClient(conn)
	getresourceClient := attestation_agent.NewGetResourceServiceClient(tc)
	keyproviderClient := attestation_agent.NewKeyProviderServiceClient(tc)

	method := flag.Arg(2)
	payload := flag.Arg(3)
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
	default:
		return errors.New("invalid method")
	}
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
