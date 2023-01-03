package common

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
