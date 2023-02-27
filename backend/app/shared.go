package app

import "time"

const (
	GreetingTaskQueue = "GREETING_TASK_QUEUE"
	ScanTaskQueue     = "SCAN_TASK_QUEUE"
)

type SubdomainS struct {
	Domain         string
	SubdomainsItem []string
	TaskId         string
	OrgName        string
}

type ScanTaskItem struct {
	Domains []string
	TaskId  string
	OrgName string
}

type Subdomain struct {
	UId string `xorm:"pk UUID 'id'"`
	// Id        string `xorm:"notnull pk UUID 'id'"`
	// Id        int    `json:"id"  xorm:"notnull pk INT 'id'"`
	Alive     int    `json:"alive"  xorm:"INT 'alive'"`
	Request   int    `json:"request"  xorm:"INT 'request'"`
	Resolve   int    `json:"resolve"  xorm:"INT 'resolve'"`
	URL       string `json:"url"  xorm:"TEXT 'url'"`
	Subdomain string `json:"subdomain"  xorm:"TEXT 'subdomain'"`
	Level     int    `json:"level"  xorm:"INT 'level'"`
	Cname     string `json:"cname"  xorm:"TEXT 'cname'"`
	IP        string `json:"ip"  xorm:"TEXT 'ip'"`
	// Public    string `json:"public"  xorm:"TEXT 'public'"`
	Cdn       int    `json:"cdn"  xorm:"INT 'cdn'"`
	Port      int    `json:"port"  xorm:"INT 'port'"`
	Status    int    `json:"status"  xorm:"INT 'status'"`
	Title     string `json:"title"  xorm:"TEXT 'title'"`
	Banner    string `json:"banner"  xorm:"TEXT 'banner'"`
	Cidr      string `json:"cidr"  xorm:"TEXT 'cidr'"`
	Asn       string `json:"asn"  xorm:"TEXT 'asn'"`
	Org       string `json:"org"  xorm:"TEXT 'org'"`
	Addr      string `json:"addr"  xorm:"TEXT 'addr'"`
	Isp       string `json:"isp"  xorm:"TEXT 'isp'"`
	Source    string `json:"source"  xorm:"TEXT 'source'"`
	FirstTime string `json:"first_time"  xorm:"TEXT 'first_time'"`
	TaskId    string `json:"task_id"  xorm:"TEXT 'task_id'"`
	Domain    string `json:"domain"  xorm:"TEXT 'domain'"`
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

type NucleiReq struct {
	Domains []string
	Tags    string
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
