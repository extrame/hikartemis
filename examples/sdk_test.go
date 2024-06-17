package examples

import (
	"encoding/json"
	"testing"

	"github.com/extrame/hikartemis"
)

func TestSDK(t *testing.T) {
	hk := hikartemis.HKConfig{
		Ip:      "127.0.0.1",
		Port:    443,
		AppKey:  "28057000",
		Secret:  "dZztQSS0000kLpURG000",
		IsHttps: true,
	}

	body := map[string]string{
		"pageNo":   "1",
		"pageSize": "100",
	}
	result, err := hk.HttpPost("/artemis/api/resource/v1/cameras", body, 15)
	if err != nil {
		t.Fatal(err)
		return
	}
	resJson, err := json.Marshal(result)
	t.Log("OK", string(resJson))

	/*body := map[string]string{
		"cameraIndexCode": "71c1e8bd1b0d406a94e7cdf88a251f9b",
		"protocol":        "rtmp",
	}
	result, err := hk.HttpPost("/artemis/api/video/v2/cameras/previewURLs", body, 15)
	if err != nil {
		t.Fatal(err)
		return
	}
	resJson, err := json.Marshal(result)
	t.Log("OK", string(resJson))*/
}
