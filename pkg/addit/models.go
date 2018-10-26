package addit

type Info struct {
	Version       string `json:"version"`
	VersionString string `json:"version_string"`
}

type URLRequest struct {
	Url    string  `json:"url" form:"url" query:"url"`
	PreUrl string  `json:"pre_url" form:"pre_url" query:"pre_url"`
}

type URLResponse struct {
	Index   int     `json:"index"`
	Hash    string  `json:"hash"`
	Count   int     `json:"count"`
}

type UrlIndex struct {
	Index   int     `xorm:"index pk autoincr unique"    json:"index"`
	Hash    string  `xorm:"VARCHAR(65) not null unique" json:"hash"`
	Count   int     `xorm:"INT(11) default 0"           json:"count"`
	Url     string  `xorm:"TEXT not null"               json:"url"`
}

