// +build !windows

package forwarder

import "github.com/logmatic/beats-forwarder/output"

func init() {
	Registry = make(map[string]output.Output)
	Registry["logmatic"] = &output.LogmaticClient{}
	Registry["udp_tcp"] = &output.SocketClient{}
	Registry["syslog"] = &output.SyslogClient{}
}
