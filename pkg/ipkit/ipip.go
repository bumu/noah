package ipkit

import (
	"fmt"
	"log"
	"noah/pkg/configkit"
	"os"

	"github.com/ipipdotnet/ipdb-go"
)

func CheckIpV4File() bool {
	_, err := os.Stat(configkit.GlobalConfig.Ipdb.Ipv4)
	return !os.IsNotExist(err)
}

func CheckIpV6File() bool {
	_, err := os.Stat(configkit.GlobalConfig.Ipdb.Ipv6)
	return !os.IsNotExist(err)
}

func IpdbInit() {
	if !CheckIpV4File() {
		log.Println("ipdb config not found")
		return
	}

	ipv4Conf := configkit.GlobalConfig.Ipdb.Ipv4
	ipv4db, err := ipdb.NewCity(ipv4Conf)
	if err != nil {
		log.Fatal(err)
	}

	ipv4db.Reload(ipv4Conf)
	fmt.Println("ipv4: ", ipv4db.IsIPv4())
	fmt.Println(ipv4db.Find("1.1.1.1", "CN"))
}

func IpdbFind(findIp string) (ret *ipdb.CityInfo, err error) {
	if !CheckIpV4File() {
		log.Println("ipdb config not found")
		return
	}

	ipv4Conf := configkit.GlobalConfig.Ipdb.Ipv4
	ipv4db, err := ipdb.NewCity(ipv4Conf)
	if err != nil {
		log.Fatal(err)
	}

	ipv4db.Reload(ipv4Conf)
	fmt.Println("ipv4: ", ipv4db.IsIPv4())

	ret, err = ipv4db.FindInfo(findIp, "CN")
	return
}
