package region

import "time"

type Region struct {
	IndexCode         string    `json:"indexCode"`
	Name              string    `json:"name"`
	ParentIndexCode   string    `json:"parentIndexCode"`
	Available         bool      `json:"available"`
	Leaf              bool      `json:"leaf"`
	CascadeCode       string    `json:"cascadeCode"`
	CascadeType       int       `json:"cascadeType"`
	CatalogType       int       `json:"catalogType"`
	ExternalIndexCode string    `json:"externalIndexCode"`
	Sort              int       `json:"sort"`
	RegionPath        string    `json:"regionPath"`
	CreateTime        time.Time `json:"createTime"`
	UpdateTime        time.Time `json:"updateTime"`
}

type RegionList []Region
