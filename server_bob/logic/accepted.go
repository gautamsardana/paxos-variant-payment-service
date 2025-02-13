package logic

import (
	"context"
	"fmt"
	"sync"

	common "GolandProjects/apaxos-gautamsardana/api_common"
	"GolandProjects/apaxos-gautamsardana/server_bob/config"
	"GolandProjects/apaxos-gautamsardana/server_bob/utils"
)

// i am a follower, i accepted the value. Now I need to send this accepted back to the
//leader and wait for commit

func SendAccepted(ctx context.Context, conf *config.Config, req *common.Accepted) {
	if req.BallotNum.ServerNumber == conf.ServerNumber {
		return
	}
	fmt.Printf("Server %d: sending accepted with request: %v\n", conf.ServerNumber, req)

	leaderAddress := utils.MapServerNumberToAddress[req.BallotNum.ServerNumber]
	server, err := conf.Pool.GetServer(leaderAddress)
	if err != nil {
		fmt.Println(err)
	}

	_, err = server.Accepted(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
}

// i am the leader - i got the accepted from followers. Now I need to send them commit messages

func ReceiveAccepted(ctx context.Context, conf *config.Config, req *common.Accepted) error {
	fmt.Printf("Server %d: received accepted with req: %v\n", conf.ServerNumber, req)

	if req.BallotNum.TermNumber != conf.CurrBallot.TermNumber ||
		req.BallotNum.ServerNumber != conf.CurrBallot.ServerNumber {
		return fmt.Errorf("invalid ballot")
	}

	if conf.MajorityHandler == nil {
		return fmt.Errorf("MajorityHandler is not initialized")
	}

	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	if conf.AcceptedServers == nil {
		config.NewAcceptedServersInfo(conf)
	}

	if conf.AcceptVal == nil {
		conf.AcceptVal = config.NewAcceptVal()
		conf.AcceptVal.BallotNumber = req.BallotNum
		conf.AcceptVal.Transactions = req.AcceptVal
	}

	conf.AcceptedServers.CurrAcceptedCount++
	conf.AcceptedServers.ServerAddresses = append(conf.AcceptedServers.ServerAddresses,
		utils.MapServerNumberToAddress[req.ServerNumber])

	if conf.AcceptedServers.CurrAcceptedCount >= (conf.ServerTotal/2)+1 {
		conf.MajorityHandler.MajorityCh <- true
	}

	return nil
}
