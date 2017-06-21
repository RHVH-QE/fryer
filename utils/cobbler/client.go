package cobbler

import (
	"bytes"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"fmt"

	"github.com/divan/gorilla-xmlrpc/xml"
)

// Cobbler represent a cobbler instance
type Cobbler struct {
	APIURL   string
	Username string
	Password string
	Token    string
}

func (c Cobbler) xmlRPCCall(method string, args interface{}) (reply struct{ Message string }, err error) {
	buf, _ := xml.EncodeClientRequest(method, args)
	resp, err := http.Post(c.APIURL, "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &reply)
	return
}

// NewCobbler is
func NewCobbler() *Cobbler {
	cb := &Cobbler{
		APIURL:   "http://10.73.60.74/cobbler_api",
		Username: "cobbler",
		Password: "cobbler",
	}
	cb.login()
	return cb
}

func (c *Cobbler) login() {
	args := struct {
		User string
		Pass string
	}{"cobbler", "cobbler"}

	r, err := c.xmlRPCCall("login", &args)
	if err != nil {
		log.Fatal(err)
	}
	c.Token = r.Message
}

// workground for simple reason
func (c Cobbler) modifyNicForNewSystem(sysID, nicName, mac, token string) {
	buf := fmt.Sprintf(modifyNicTpl, sysID, nicName, mac, token)
	http.Post(c.APIURL, "text/xml", bytes.NewBuffer([]byte(buf)))
}

// NewSystem add a new systen into cobbler server
func (c Cobbler) NewSystem(name, profile, comment, status, kargs string, nic []string) {
	r, err := c.xmlRPCCall("new_system", &struct{ Token string }{c.Token})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("get new system id :: %s", r.Message)

	nsp := make(map[string]string)
	nsp["name"],
		nsp["profile"],
		nsp["comment"],
		nsp["status"],
		nsp["kernel_options"] = name, profile, comment, status, kargs

	for k, v := range nsp {
		c.xmlRPCCall("modify_system", &struct {
			ID    string
			Key   string
			Val   string
			Token string
		}{r.Message, k, v, c.Token})
	}

	c.modifyNicForNewSystem(r.Message, nic[0], nic[1], c.Token)

	r2, _ := c.xmlRPCCall("save_system", &struct {
		ID    string
		Token string
	}{r.Message, c.Token})
	log.Info(r2.Message)
}

// RemoveSystem is
func (c Cobbler) RemoveSystem(name string) {
	_, err := c.xmlRPCCall("remove_system", &struct {
		Name  string
		Token string
	}{name, c.Token})
	if err != nil {
		log.Warn(err)
	}
}
