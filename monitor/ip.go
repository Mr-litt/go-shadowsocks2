package monitor

import (
	"sync"
	"fmt"
	"time"
	"io/ioutil"
	"strings"
)

type ValidIpMap map[string]bool

type Ip struct {
	IpMap sync.Map
	Email * Email
	validIp ValidIpMap
	sendCount map[string]uint32
	mutex sync.Mutex
}

var GMonitorIp *Ip

func InitMonitorIp(emailSwitch bool, emailUserName string, emailPassword string, validIps string) {
	monitorIp := &Ip{}
	monitorIp.Email = InitEmail(emailSwitch, emailUserName, emailPassword)

	monitorIp.validIp = make(map[string]bool)
	validIpArr := strings.Split(validIps, ",")
	for _, ip  := range validIpArr {
		monitorIp.validIp[ip] = true
	}

	monitorIp.sendCount = make(map[string]uint32)
	GMonitorIp = monitorIp
}

func (ip * Ip) Add (ipAddr string){

	// 检测是否属于合法Ip
	_, ok := ip.validIp[ipAddr]
	if !ok {

		// 检测是否已经发送过邮件通知
		_, ok := ip.IpMap.Load(ipAddr)
		if !ok {
			ip.IpMap.Store(ipAddr, true)

			// 1天最多发10次
			currentDate := time.Now().Format("2006-01-02")
			ip.mutex.Lock()
			_, ok := ip.sendCount[currentDate]
			if !ok {
				ip.sendCount[currentDate] = 0
			}
			ip.sendCount[currentDate]++
			ip.mutex.Unlock()
			if ip.sendCount[currentDate] <= 10 {
				ip.NewIpSendEmail(ipAddr)
			}

			// 记录日志
			nowDate := time.Now().Format("2006-01-02 15:04:05")
			writeData := "[" + nowDate + "] " + ipAddr
			err := ioutil.WriteFile("/tmp/newIpMonitor.log" ,[]byte(writeData),0644)
			if err != nil {
				fmt.Printf("write error: %v", err)
			}
		}
	}
}

func (ip * Ip) NewIpSendEmail(ipAddr string) {
	sendBody := "[waning] [" + time.Now().Format("2006-01-02 15:04:05") + "] new Ip:" + ipAddr
	ip.Email.Send(sendBody)
}
