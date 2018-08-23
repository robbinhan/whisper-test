package main

import (
	"crypto/ecdsa"
	"fmt"
	mrand "math/rand"
	"net"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
	"github.com/robbinhan/whisper-test/utils"
)

func main() {
	var priKey *ecdsa.PrivateKey
	keyFile := "node2.key"
	if utils.ExistsFile(keyFile) {
		priKey, _ = crypto.LoadECDSA(keyFile)
	} else {
		priKey, _ = crypto.GenerateKey()
		crypto.SaveECDSA(keyFile, priKey)
	}

	// set the log level to Trace
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stdout, log.TerminalFormat(true))))

	whisper := whisperv6.New(&whisperv6.DefaultConfig)

	p2pConfig := &p2p.Config{
		PrivateKey: priKey,
		MaxPeers:   10,
		ListenAddr: ":8001",
		Protocols:  whisper.Protocols(),
		Logger:     log.Root(),
	}
	srv := p2p.Server{Config: *p2pConfig}
	if err := srv.Start(); err != nil {
		fmt.Println("could not start server:", err)
		os.Exit(1)
	}

	log.Info("Node", "info", srv.NodeInfo())

	whisper.Start(&srv)

	// 连接node1
	id := discover.MustHexID("fdf241bcdd60ad1ab2e53518b43bf585d14f016c273595b65e4daea86151e2f839fa7a99ca1e188ae8d2263c42227b51512a8687da2f686ebf59559500e9dde3")
	n1 := discover.NewNode(id, net.ParseIP("127.0.0.1"), 8000, 8000)
	srv.AddPeer(n1)

	// 生成消息
	params, err := generateMessageParams()
	if err != nil {
		srv.Logger.Crit("failed generateMessageParams", "err", err)
	}

	params.TTL = 1
	msg, err := whisperv6.NewSentMessage(params)
	if err != nil {
		srv.Logger.Crit("failed to create new message", "err", err)
	}

	// 装到信封
	env, err := msg.Wrap(params)
	if err != nil {
		srv.Logger.Crit("failed Wrap", "err", err)
	}

	// 循环发送消息
	go func() {
		for {
			time.Sleep(time.Second)
			for _, peer := range srv.Peers() {
				srv.Logger.Info("print peer info", "id", peer.ID(), "name", peer.String())

				err = whisper.SendP2PMessage(peer.ID().Bytes(), env)
				if err != nil {
					srv.Logger.Crit("failed SendP2PDirect", "err", err)
				}
				srv.Logger.Info("sent message done", "to", peer.String())
			}
		}
	}()

	select {}
}

func generateMessageParams() (*whisperv6.MessageParams, error) {
	// set all the parameters except p.Dst and p.Padding
	var p whisperv6.MessageParams

	p.PoW = 0
	p.WorkTime = 0
	p.TTL = uint32(mrand.Intn(1024))
	p.Payload = []byte("This is whisper test message from node2")
	p.KeySym = []byte("whisperv6 message test..........") // 必须32字节
	p.Topic = whisperv6.BytesToTopic([]byte("node2"))

	// var err error
	// p.Src, err = crypto.GenerateKey()
	// if err != nil {
	// 	return nil, err
	// }

	return &p, nil
}
