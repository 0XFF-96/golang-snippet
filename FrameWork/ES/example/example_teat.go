package example

import (
	"encoding/json"
	"testing"

	elasticv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/require"
)

func Test_mockESClient(t *testing.T) {
	// user, pass := conf.GetESUserPass()
	user := ""
	pass := ""
	addr := "http://127.0.0.1:9200/"

	cli, err := elasticv7.NewClient(elasticv7.Config{
		Addresses: []string{addr},
		Username:  user,
		Password:  pass,
		Transport: &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          20,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}})
	require.NoError(t, err)

	resp, err := cli.Search()
	_ = resp
	//if err != nil {
	//	return nil, err
	//}
	// log.Infof("labeling dao inited with: %v", conf.GetESAddresses())
}


func Test_DefaultESClient(t *testing.T) {
	// 通过验证我们知道，default es client 是没有验证体系的。
	es, err := elasticv7.NewDefaultClient()
	require.NoError(t, err)
	res, err := es.Info()
	require.NoError(t, err)

	var r  map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	require.NoError(t, err)
	t.Log(r)
}


