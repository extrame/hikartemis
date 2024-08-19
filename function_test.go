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

	list, err := hk.GetSubRegion("88fabbb0c57943e8bb6dadb272403f97")

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

	// slog.SetLevel(slog.DebugLevel)

	list, err := hk.GetSubRegion("88fabbb0c57943e8bb6dadb272403f97")

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

	fmt.Println("\nindex_code,name,source_id,unit_id")
	for _, camera := range cameras {
		fmt.Printf("%s,%s,2,\n", camera.IndexCode, camera.Name)
	}

	url, err := hk.GetCameraUrl(&cameras[0])

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(url)
}

func TestGetCameraOfRegion(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	// slog.SetLevel(slog.DebugLevel)

	cameras, err := hk.GetCameraList("88fabbb0c57943e8bb6dadb272403f97")

	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println("\nindex_code,name,source_id,unit_id")
	for _, camera := range cameras {
		fmt.Printf("%s,%s,2,\n", camera.IndexCode, camera.Name)
	}

	url, err := hk.GetCameraUrl(&cameras[0])

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(url)
}

func TestGetCameraById(t *testing.T) {
	hk, err := readConfigFromFile("./artemis.yaml")

	if err != nil {
		t.Fatal(err)
		return
	}

	// list, err := hk.GetSubRegion("64937291a89d43d580181f333d039c13")

	// if err != nil {
	// 	t.Fatal(err)
	// 	return
	// }

	// t.Log(list)

	var ids = []string{"7bd23377-c8ba-41e2-af0c-38d2299ff751"}

	// for _, region := range list {
	// 	ids = append(ids, region.IndexCode)
	// }

	cameras, err := hk.GetCameraList(ids...)

	if err != nil {
		t.Fatal(err)
		return
	}

	for _, camera := range cameras {
		fmt.Println(camera.RegionIndexCode, ",", camera.Name, ",", camera.IndexCode)
	}

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
