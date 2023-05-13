package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns cli app ready to be executed
func Generate() *cli.App{
	app := cli.NewApp()
	app.Name = "Command line app"
	app.Usage = "Search for server IPs and Names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "walisontsx.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search for IPs address",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name: "servers",
			Usage: "Search for severs address",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}