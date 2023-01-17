package common

import (
	"time"
)

type Subdomains struct {
	// ID        int    `json:"id"`
	Alive     int    `json:"alive"`
	Request   int    `json:"request"`
	Resolve   int    `json:"resolve"`
	URL       string `json:"url"`
	Subdomain string `json:"subdomain"`
	Level     int    `json:"level"`
	Cname     string `json:"cname"`
	IP        string `json:"ip"`
	// Public    int    `json:"public"`
	Cdn    int `json:"cdn"`
	Port   int `json:"port"`
	Status int `json:"status"`
	// Reason string `json:"reason"`
	Title     string `json:"title"`
	Banner    string `json:"banner"`
	Cidr      string `json:"cidr"`
	Asn       string `json:"asn"`
	Org       string `json:"org"`
	Addr      string `json:"addr"`
	Isp       string `json:"isp"`
	Source    string `json:"source"`
	FirstTime string `json:"first_time"`
	TaskId    string `json:"task_id"`
	Domain    string `json:"domain"`
}
type DirItems struct {
	Dir []*DirItem `json:"dir"`
}
type DirItem struct {
	Dir  string
	Code string
	Size int
}

type RustScanItems struct {
	Items []*RustScanItem
}

type RustScanItem struct {
	Port     int
	Domain   string
	Ip       string
	Service  string
	Protocol string
	Status   string
	Reason   string
}

type Nucleivulns struct {
	Id       int64  `gorm:"column:username;not null;type:int(4) primary key auto_increment;comment:'标识'"`
	Password string `gorm:"column:password;type:varchar(30);index:idx_name"`
	// 创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`

	Template         string      `gorm:"column:template;type:varchar(255)"`
	TemplateURL      string      `gorm:"column:templateURL;type:varchar(255)"`
	TemplateID       string      `gorm:"column:template_id;type:varchar(255)"`
	TemplatePath     string      `gorm:"column:template_id;type:varchar(255)"`
	Info             Info        `gorm:"column:template_id;type:varchar(255)"`
	Type             string      `json:"type"`
	Host             string      `json:"host"`
	MatchedAt        string      `json:"matched-at"`
	ExtractedResults []string    `json:"extracted-results"`
	IP               string      `json:"ip"`
	Timestamp        time.Time   `json:"timestamp"`
	CurlCommand      string      `json:"curl-command"`
	MatcherStatus    bool        `json:"matcher-status"`
	MatchedLine      interface{} `json:"matched-line"`
	TaskId           string      `json:"task_id"`
}
type Metadata struct {
	Verified    bool   `json:"verified"`
	ShodanQuery string `json:"shodan-query"`
}
type Info struct {
	Name      string      `json:"name"`
	Author    []string    `json:"author"`
	Tags      []string    `json:"tags"`
	Reference interface{} `json:"reference"`
	Severity  string      `json:"severity"`
	Metadata  Metadata    `json:"metadata"`
}
