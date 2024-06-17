package resource

import "time"

type Resource struct {
	IndexCode        string    `json:"indexCode"`
	Name             string    `json:"name"`
	ResourceType     string    `json:"resourceType"`
	DoorNo           int       `json:"doorNo"`
	Description      string    `json:"description"`
	ParentIndexCodes string    `json:"parentIndexCodes"`
	RegionIndexCode  string    `json:"regionIndexCode"`
	RegionPath       string    `json:"regionPath"`
	ChannelType      string    `json:"channelType"`
	ChannelNo        string    `json:"channelNo"`
	InstallLocation  string    `json:"installLocation"`
	CapabilitySet    string    `json:"capabilitySet"`
	ControlOneId     string    `json:"controlOneId"`
	ControlTwoId     string    `json:"controlTwoId"`
	ReaderInId       string    `json:"readerInId"`
	ReaderOutId      string    `json:"readerOutId"`
	ComId            string    `json:"comId"`
	CreateTime       time.Time `json:"createTime"`
	UpdateTime       time.Time `json:"updateTime"`
}

type ResourceList []Resource
