package c1device

var wsPrefix string = "ws://"
var wsSufix string = "/wscomm.cgi"
var localhost string = "http://localhost"
var port string = "8080"

// C1Device describes a C1 device
type C1Device struct {
	Name      string
	IP        string
	TCPPort   string
	WsChannel chan []byte
}
