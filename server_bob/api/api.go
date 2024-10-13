package api

import (
	common "GolandProjects/apaxos-gautamsardana/api_common"
	"GolandProjects/apaxos-gautamsardana/server_bob/config"
	"GolandProjects/apaxos-gautamsardana/server_bob/logic"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type Server struct {
	common.UnimplementedPaxosServer
	Config *config.Config
}

func (s *Server) ProcessTxn(ctx context.Context, req *common.ProcessTxnRequest) (*emptypb.Empty, error) {
	s.Config.CurrRetryCount = 0
	err := logic.ProcessTxn(ctx, req, s.Config, false)
	if err != nil {
		log.Printf("Error processing txn: %v", err)
		return nil, err
	}
	log.Printf("txn successful!")

	return nil, nil
}

func (s *Server) Prepare(ctx context.Context, req *common.Prepare) (*emptypb.Empty, error) {
	err := logic.ReceivePrepare(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error processing prepare from leader: %v", err)
		return nil, err
	}
	return nil, nil
}

func (s *Server) Promise(ctx context.Context, req *common.Promise) (*emptypb.Empty, error) {
	err := logic.ReceivePromise(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error processing promise request from follower: %v", err)
		return nil, err
	}
	return nil, nil
}

func (s *Server) Accept(ctx context.Context, req *common.Accept) (*emptypb.Empty, error) {
	err := logic.ReceiveAccept(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error processing accept request from leader: %v", err)
		return nil, err
	}
	return nil, nil
}

func (s *Server) Accepted(ctx context.Context, req *common.Accepted) (*emptypb.Empty, error) {
	err := logic.ReceiveAccepted(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error processing accepted request from follower: %v", err)
		return nil, err
	}
	return nil, nil
}

func (s *Server) Commit(ctx context.Context, req *common.Commit) (*emptypb.Empty, error) {
	err := logic.ReceiveCommit(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error processing commit request from leader: %v", err)
		return nil, err
	}
	return nil, nil
}

func (s *Server) Sync(ctx context.Context, req *common.SyncRequest) (*emptypb.Empty, error) {
	err := logic.SyncRequest(ctx, s.Config, req)
	if err != nil {
		log.Printf("Error processing sync request from slow follower: %v", err)
		return nil, err
	}
	return nil, nil
}
