package core

//config.yaml
type Monitor_yaml struct { //[*]type首字母必须大写才能读取到。
	Monitor Monitor `yaml:"monitor"`
}
type Monitor struct {
	Domain      []string `yaml:"domain"`
	ServerJiang ServerJiang `yaml:"serverJiang"`
	Chinaz struct {
		Apikey string `yaml:"apikey"`
		Url    string `yaml:"url"`
	}
	Slackbot struct {
		Apikey string `yaml:"apikey"`
		Url    string `yaml:"url"`
	}
	AutoUpadteNewParma string `yaml:"auto_update_newsubdomain_to_old"`
	//Tgbot struct {
	//	Apikey string `yaml:"apikey"`
	//	Url    string `yaml:"url"`
	//}
	//Dingtalkbot struct {
	//	Apikey string `yaml:"apikey"`
	//	Url    string `yaml:"url"`
	//}
	Github struct {
		Account  string `yaml:"account"`
		Password string `yaml:"password"`
	}
	Tools Tools `yaml:"tools"`
}
type Tools struct {
	SubdomainBrute_tools SubdomainBrute_tools `yaml:"subdomainBrute_tools"`
	Web_check_tools      Web_check_tools      `yaml:"web_check_tools"`
}
type ServerJiang struct {
	Enable string `yaml:"enable"`
	ServerJiangApi string `yaml:"server_jiang_api"`
}
type SubdomainBrute_tools struct {
	Subfinder_enable  string `yaml:"subfinder_enable"`
	Ksubdomain_enable string `yaml:"ksubdomain_enable"`
}
type Web_check_tools struct {
	Httpx_enable string `yaml:"httpx_enable"`
}
type Subdomain struct {
	Domain        string
	SubdomainName string
	Http_status   int
	New           int //1代表新增，0代表不是新增，-1代表域名已经找不到了
	Ip            string
	Resource      string
}
type tgRobot struct {
	//后期完成
}
type slackRobot struct {
}
type sqliteConfig struct {
}
type chinazApi struct {
}
type DomainList struct {
	Domain    string
	Subdomain []string
}

//type SubfinderResult struct {
//	Host   string `json:"host"`
//	Source string `json:"source"`
//}
type HttpxResult struct {
	IPs    []string `json:"ips"`
	CNAMEs []string `json:"cnames,omitempty"`
	//raw           string
	URL string `json:"url"`
	//Location      string `json:"location"`
	Title         string `json:"title"`
	WebServer     string `json:"webserver"`
	Response      string `json:"serverResponse,omitempty"`
	ContentType   string `json:"content-type,omitempty"`
	Method        string `json:"method"`
	IP            string `json:"ip"`
	ContentLength int    `json:"content-length"`
	StatusCode    int    `json:"status-code"`
	//VHost         bool           `json:"vhost"`
	//WebSocket     bool           `json:"websocket,omitempty"`
	//Pipeline      bool           `json:"pipeline,omitempty"`
	HTTP2 bool `json:"http2"`
	CDN   bool `json:"cdn,omitempty"`
	//ResponseTime  string         `json:"response-time"`
}
