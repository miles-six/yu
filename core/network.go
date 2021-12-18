package core

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/core/types"
	"io"
	"io/ioutil"
	"net/http"
)

func SendHeartbeats(addrs []string, handleAlive func(addr string) error, handleDead func(addr string) error) {
	for _, addr := range addrs {
		_, err := http.Get(addr + HeartbeatPath)
		if err != nil {
			logrus.Errorf("send heartbeat to (%s) error: %s", addr, err.Error())
			err = handleDead(addr)
			if err != nil {
				logrus.Errorf("handle dead node (%s) error: %s", addr, err.Error())
			}
		} else {
			logrus.Debugf("send heartbeat to (%s) succeed!", addr)
			err = handleAlive(addr)
			if err != nil {
				logrus.Errorf("handle alive node (%s) error: %s", addr, err.Error())
			}
		}
	}

}

func PostRequest(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	cli := &http.Client{}
	return cli.Do(req)
}

func DecodeBlockFromHttp(body io.ReadCloser, chain types.IBlockChain) (*types.CompactBlock, error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return types.DecodeCompactBlock(byt)
}
