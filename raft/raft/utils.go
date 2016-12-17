package raft

import (
	"math/big"
	"crypto/sha1"
	"time"
	"net/rpc"
	"fmt"
	"os"
)

var connectionMap = make(map[string] *rpc.Client)

func AddrToId(addr string, length int) string {
	h := sha1.New()
	h.Write([]byte(addr))
	v := h.Sum(nil)
	keyInt := big.Int{}
	keyInt.SetBytes(v[:length])
	return keyInt.String()
}

func (r *RaftNode) electionTimeOut() <- chan time.Time {
	return time.After(r.conf.ElectionTimeout)
}

func (r *RaftNode) heartBeats() <- chan time.Time {
	return time.After(r.conf.HeartbeatFrequency)
}

func makeRemoteCall(remoteNode *NodeAddr, method string, req interface{}, resp interface{}) error {
	var err error
	remoteAddStr := remoteNode.address
	client, ok := connectionMap[remoteAddStr]
	if !ok {
		client, err = rpc.Dial("tcp", remoteAddStr)
		if err != nil {
			return nil
		}
		connectionMap[remoteAddStr] = client
	}
	methodPath := fmt.Sprintf("%v.%v", remoteAddStr, method)
	err = client.Call(methodPath, req, resp)
	if err != nil {
		return err
	}
	return nil
}

func getFileStats(filename string) (uint64, bool) {
	stats, err := os.Stat(filename)
	if err == nil {
		return stats.Size(), true
	} else if os.IsNotExist(err) {
		return 0, false
	} else {
		panic(err)
	}
}