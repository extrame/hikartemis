package hikartemis

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestGetRootRegion(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	list, err := hk.GetRootRegion()

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(list)

}

func TestGetRegion(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	root, err := hk.GetRootRegion()

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(root)

	list, err := hk.GetSubRegion(root.IndexCode)

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(list)

}

func TestGetRegion2(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	list, err := hk.GetSubRegion("83b599305b9d4642b9656d1bfc4180c9")

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(list)

}

func TestGetRegionAndCamera(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	list, err := hk.GetSubRegion("64937291a89d43d580181f333d039c13")

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(list)

	var ids []string

	for _, region := range list {
		ids = append(ids, region.IndexCode)
	}

	cameras, err := hk.GetCameraList(ids...)

	if err != nil {
		t.Fatal(err)
		return
	}

	// t.Log(cameras)

	url, err := hk.GetCameraUrl(&cameras[0])

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(url)
}

func TestGetResource(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	list, err := hk.GetResourceList()

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(list)

}

func TestGetCamera(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	list, err := hk.GetCameraList()

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(list)

}

func readConfigFromFile(filename string) (*HKConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &HKConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	fmt.Println("config:", config)

	config.Init(15)

	return config, nil
}
