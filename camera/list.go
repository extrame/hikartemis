package camera

import "time"

type Camera struct {
	IndexCode         string    `json:"indexCode"`
	ResourceType      string    `json:"resourceType"`
	ExternalIndexCode string    `json:"externalIndexCode"`
	Name              string    `json:"name"`
	ChanNum           int       `json:"chanNum"`
	CascadeCode       string    `json:"cascadeCode"`
	ParentIndexCode   string    `json:"parentIndexCode"`
	Longitude         string    `json:"longitude"`
	Latitude          string    `json:"latitude"`
	Elevation         string    `json:"elevation"`
	CameraType        int       `json:"cameraType"`
	Capability        string    `json:"capability"`
	RecordLocation    string    `json:"recordLocation"`
	ChannelType       string    `json:"channelType"`
	RegionIndexCode   string    `json:"regionIndexCode"`
	RegionPath        string    `json:"regionPath"`
	TransType         int       `json:"transType"`
	TreatyType        string    `json:"treatyType"`
	InstallLocation   string    `json:"installLocation"`
	CreateTime        time.Time `json:"createTime"`
	UpdateTime        time.Time `json:"updateTime"`
	DisOrder          int       `json:"disOrder"`
	ResourceIndexCode string    `json:"resourceIndexCode"`
	DecodeTag         string    `json:"decodeTag"`
	CameraRelateTalk  string    `json:"cameraRelateTalk"`
	RegionName        string    `json:"regionName"`
	RegionPathName    string    `json:"regionPathName"`
}

type CameraList []Camera

type Url struct {
	Url string `json:"url"`
}
