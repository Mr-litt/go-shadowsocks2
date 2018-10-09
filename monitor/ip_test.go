package monitor

import (
	"testing"
)

func TestGetArtisanGoods(t *testing.T) {

	InitMonitorIp(false, "137057181@qq.com","","36.157.164.139,223.104.10.1")
	GMonitorIp.Add("36.157.164.139")
	GMonitorIp.Add("223.104.10.1")
	GMonitorIp.Add("223.104.10.2")
	GMonitorIp.Add("223.104.10.3")
	GMonitorIp.Add("223.104.10.4")
	GMonitorIp.Add("223.104.10.5")
	GMonitorIp.Add("223.104.10.6")
	GMonitorIp.Add("223.104.10.7")
	GMonitorIp.Add("223.104.10.8")
	GMonitorIp.Add("223.104.10.9")
	GMonitorIp.Add("223.104.10.10")
	GMonitorIp.Add("223.104.10.11")
	GMonitorIp.Add("223.104.10.12")
	t.Log("ok")

}