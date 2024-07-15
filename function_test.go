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

	list, err := hk.GetSubRegion("bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f")

	//9eb7b355-2a3c-4ba9-8646-8e7346821d34
	//bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f
	//dd59cb7d-331e-4793-9e81-6ad25a09b92c
	// [{7bd23377-c8ba-41e2-af0c-38d2299ff751 正缘种猪场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {dd59cb7d-331e-4793-9e81-6ad25a09b92c 正缘育肥一场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {1d7e97be-ebd8-456c-9610-05e82076356a 正缘育肥二场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {0d606016-1515-4f9a-8179-beb9851183b9 正缘育肥三场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {a0e539ec-773c-4e04-972d-05d2c1efa449 正缘育肥四场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {b915a89b-6f43-46f2-a71b-bec1c70714dd 正缘育肥五场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {7100af55-8507-4991-8b88-4745262e62a8 正缘育肥六场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {0342e85d-7215-47f9-be24-c0a003e3b10f 正缘育肥七场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {32ff3fd6-9ee1-4404-847b-7226638f9a9b 正缘育肥八场 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {137e404a-4dcc-4a79-bd5e-e699da3d0652 正缘公猪站 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {c8b16be7-85d7-4e55-90fb-1ce30b61d1af 正缘洗消二 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	// {42017acf-a0bb-474e-8a5c-ef1796e5e7e5 正缘洗消三 bea6f327-7cfa-4a4b-bf55-b23bfc4ca41f false false  0 0  0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}]
	//09701687-23cf-426a-875e-7aabf185df66

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

	list, err := hk.GetSubRegion("137e404a-4dcc-4a79-bd5e-e699da3d0652")

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
