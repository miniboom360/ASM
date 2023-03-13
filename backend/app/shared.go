package app

import "time"

const (
	GreetingTaskQueue = "GREETING_TASK_QUEUE"
	ScanTaskQueue     = "SCAN_TASK_QUEUE"
)

const (
	DB_TABLE_NOT_EXIST string = "DB_TABLE_NOT_EXIST"
)

type SubdomainS struct {
	MainDomain string
	// change to map
	Subdomains       map[string]*SubdomainItem
	SubdomainsSclice []string
	TaskId           string
	OrgName          string
	// MainNd           *NaabuData
	// MainHxd          *HttpXData

	// SIS            []*SubdomainItem
}

// 单个子域名
type SubdomainItem struct {
	SubDomain string
	TaskId    string
	// 这个不应该只在子域名当中，因为有的只可能是一个ip
	Nd  []*NaabuData
	Hxd *HttpXData
}

type TaskItem struct {
	// Id string `json:"id"  xorm:"notnull pk UUID 'id'"`
	Id               int      `json:"id"  xorm:"pk autoincr"`
	TaskId           string   `json:"task_id"  xorm:"TEXT 'task_id'"`
	OrganizationName string   `json:"organization"  xorm:"TEXT 'organization'"`
	Domains          []string `json:"domains"  xorm:"TEXT 'domains'"`
	Staus            string   `json:"staus"  xorm:"TEXT 'staus'"`
	EntryId          int      `json:"entry_id"  xorm:"int 'EntryId'"`
	ScanPolice       string   `json:"scan_policy"  xorm:"TEXT 'scan_policy'"`
	// for example every monday 8:00 execution something
	Period string `json:"period"  xorm:"TEXT 'period'"`
}
type ScanTaskItem struct {
	Domains []string
	TaskId  string
	OrgName string
	ScanOpt ScanOption
	HttpxReq
}

type ScanOption struct {
	PortTag      string `json:"portTag"`
	PortScanOnly bool   `json:"portScanOnly"`
}

type NucleiReq struct {
	Domains []string
	Tags    string
}

type PortScanReq struct {
	Targets []string
	Tag     string
}

type HttpxReq struct {
	Domain     string
	ThreadsNum int
	Targets    []string
}

type NaabuData struct {
	Host string `json:"host"`
	IP   string `json:"ip"`
	Port struct {
		Port     int  `json:"Port"`
		Protocol int  `json:"Protocol"`
		TLS      bool `json:"TLS"`
	} `json:"port"`
	Timestamp time.Time `json:"timestamp"`
}

type HttpXData struct {
	Timestamp time.Time `json:"timestamp"`
	Hash      struct {
		BodyMd5       string `json:"body_md5"`
		BodyMmh3      string `json:"body_mmh3"`
		BodySha256    string `json:"body_sha256"`
		BodySimhash   string `json:"body_simhash"`
		HeaderMd5     string `json:"header_md5"`
		HeaderMmh3    string `json:"header_mmh3"`
		HeaderSha256  string `json:"header_sha256"`
		HeaderSimhash string `json:"header_simhash"`
	} `json:"hash"`
	Port             string   `json:"port"`
	URL              string   `json:"url"`
	Input            string   `json:"input"`
	Title            string   `json:"title"`
	Scheme           string   `json:"scheme"`
	Webserver        string   `json:"webserver"`
	ContentType      string   `json:"content_type"`
	Method           string   `json:"method"`
	Host             string   `json:"host"`
	Path             string   `json:"path"`
	FinalURL         string   `json:"final_url"`
	Time             string   `json:"time"`
	ChainStatusCodes []int    `json:"chain_status_codes"`
	A                []string `json:"a"`
	Tech             []string `json:"tech"`
	Words            int      `json:"words"`
	Lines            int      `json:"lines"`
	StatusCode       int      `json:"status_code"`
	ContentLength    int      `json:"content_length"`
	Failed           bool     `json:"failed"`
}

type Nucleivulns struct {
	Template         string    `json:"template"  xorm:"TEXT 'template'"`
	TemplateURL      string    `json:"template-url"  xorm:"TEXT 'template_url'"`
	TemplateID       string    `json:"template-id"  xorm:"TEXT 'template_id'"`
	TemplatePath     string    `json:"template-path"  xorm:"TEXT 'template_path'"`
	Info             Info      `json:"info"  xorm:"extends TEXT 'info'"`
	Type             string    `json:"type"  xorm:"TEXT 'type'"`
	Host             string    `json:"host"  xorm:"TEXT 'host'"`
	MatchedAt        string    `json:"matched-at"  xorm:"TEXT 'matched_at'"`
	ExtractedResults []string  `json:"extracted-results"  xorm:"TEXT 'extracted_results'"`
	IP               string    `json:"ip"  xorm:"TEXT 'ip'"`
	Timestamp        time.Time `json:"timestamp"  xorm:"DateTime 'timestamp'"`
	CurlCommand      string    `json:"curl-command"  xorm:"TEXT 'curl_command'"`
	MatcherStatus    bool      `json:"matcher-status"  xorm:"TEXT 'matcher_status'"`
	MatchedLine      string    `json:"matched-line"  xorm:"TEXT 'matched_line'"`
	TaskId           string    `json:"task-id"  xorm:"varchar(255) notnull unique 'task_id'"`
}

type Metadata struct {
	Id string `json:"id"  xorm:"notnull pk UUID 'id'"`
	// Id          int    `json:"id"  xorm:"pk autoincr"`
	Verified    bool   `json:"verified"`
	ShodanQuery string `json:"shodan-query"`
}
type Info struct {
	// Name string `json:"name"  xorm:"TEXT 'name'"`
	Name   string   `json:"name"  xorm:"varchar(255) 'name'"`
	Author []string `json:"author"  xorm:"TEXT 'author'"`
	Tags   []string `json:"tags"  xorm:"TEXT 'tags'"`
	// Reference interface{} `json:"reference"  xorm:"TEXT 'reference'"`
	Severity string   `json:"severity"  xorm:"TEXT 'severity'"`
	Metadata Metadata `json:"metadata"  xorm:"extends TEXT 'metadata'"`
}

type ScanTask struct {
	Domains  []string   `json:"domains"`
	ScanName string     `json:"scan_name"`
	Orgname  string     `json:"orgname"`
	ScanOpt  ScanOption `json:"scanOpt"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	// Id int `json:"id"  xorm:"pk autoincr"`
	// UserId   string  `json:"userId,omitempty"`
	// Username string  `json:"username,omitempty"`
	// RealName string  `json:"realName,omitempty"`
	Avatar string `json:"avatar"`
	// Desc     string  `json:"desc,omitempty"`
	Password string `json:"password,omitempty"`
	// Token    string  `json:"token,omitempty"`
	HomePath string `json:"homePath,omitempty"`
	// Roles    []*Role `json:"roles,omitempty"`
	Result `xorm:"json"`
}

type Role struct {
	RoleName string `json:"roleName,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Result struct {
	Desc     string  `json:"desc"`
	RealName string  `json:"realName"`
	Roles    []*Role `json:"roles" xorm:"json"`
	Token    string  `json:"token"`
	UserID   string  `json:"userId"`
	Username string  `json:"username"`
}

type LoginResp struct {
	Result Result `json:"result"`
	RespNormal
}

type RespNormal struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type GetUserInfoResp struct {
	RespNormal
	Result *User `json:"result"`
}
