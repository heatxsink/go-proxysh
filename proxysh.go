package proxysh

import (
	"fmt"
	"encoding/xml"
	"net/url"
	"github.com/golang/glog"
	"github.com/heatxsink/go-httprequest"
)

const (
	SERVER_LOAD_URL = "https://proxy.sh/api.php"
	DEBUG_MODE = true
)

type XmlResponse struct {
	XMLName xml.Name `xml:"account_info"`
	Credentials Credentials `xml:"credentials,omitempty"`
	ServerList []Server `xml:"server_list>server,omitempty"`
}

type Credentials struct {
	XMLName xml.Name `xml:"credentials,omitempty"`
	Username string `xml:"username,omitempty"`
	Password string `xml:"password,omitempty"`
}

type Server struct {
	XMLName xml.Name `xml:"server,omitempty"`
	Address string `xml:"address,omitempty"`
	Location string `xml:"location,omitempty"`
	ServerLoad float32 `xml:"server_load,omitempty"`
}

type ProxySh struct {
	Username string
	Password string
}

func New(username string, password string) *ProxySh {
	p := new (ProxySh)
	p.Username = username
	p.Password = password
	return p
}

func (p *ProxySh) GetServerLoad() (XmlResponse, error) {
	server_load_url := fmt.Sprintf(SERVER_LOAD_URL)
	data := url.Values{}
	data.Set("u", p.Username)
	data.Add("p", p.Password)
	return get_root_object(server_load_url, data)
}

func create_root_object(body []byte) (XmlResponse, error) {
	var feed XmlResponse
	err := xml.Unmarshal(body, &feed)
	if err != nil {
		glog.Errorln("create_root_object xml.Unmarshal()")
		glog.Errorln(err)
		glog.Flush()
	}
	return feed, err
}

func get_root_object(url string, data url.Values) (XmlResponse, error) {
	var root XmlResponse
	hr := httprequest.NewWithDefaults()
	headers := make(map[string]string)
	body, status_code, err := hr.PostUrlEncoded(url, headers, data)
	if err != nil {
		return root, err
	}
	if status_code == 200 {
		root, err = create_root_object(body)
	} else {
		glog.Warning("***********")
		glog.Warning("NON HTTP 200 Status Code")
		glog.Warning("URL:         ", url)
		glog.Warning("Status Code: ", status_code)
		glog.Warning("Body:        ", string(body))
		glog.Warning("***********")
		glog.Flush()
	}
	if DEBUG_MODE {
		glog.Infoln(url)
		glog.Infoln(string(body))
		glog.Flush()
	}
	return root, err
}
