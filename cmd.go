package main

import (
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/fgergo/naive"
)

var ideas = flag.Bool("ideas", false, "list ideas, examples")
var all = flag.Bool("all", false, "print all cgnat ip addresses")

const ideastxt = `Command ip100 prints all local ip addresses in 100.64.0.0/10 subnet (details: rfc6598).
100.64.0.0/10 is from 100.64.0.0 to 100.127.255.255.

Ideas for programs to be used with ip100 command:
- access remote filesystems via 9p without authentication (eg. serve with https://github.com/Harvey-OS/ninep/tree/master/cmd/ufs )
- private streamer, listen to your own music (eg. https://github.com/fgergo/boringstreamer )
- small business database listening only in the cgnat range
- some_funky_home_built_ad-hoc_family_server_software_without_any_auth
- any server software without authentication intended for personal use or for a small group of trusted naive people

# examples:
# (for Linux, macOS, Windows etc.)

# install ip100
$ go get github.com/fgergo/ip100

# serve 9p securely (thanks tailscale!) without authentication:
$ go get github.com/Harvey-OS/ninep/cmd/ufs
$ ip100=` + "`" + `ip100` + "`" + ` ufs -addr $ip100:5640

# improvise a radio station from your own music:
$ go get github.com/fgergo/boringstreamer
$ ip100=` + "`" + `ip100` + "`" + ` boringstreamer -addr $ip100:4444
`

func main() {
	flag.Parse()

	if *ideas {
		fmt.Println(ideastxt)
		return
	}

	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	// select all ip addresses in cgnat range
	var ip []string
	for i := 0; i < len(interfaceAddrs); i++ {
		s := interfaceAddrs[i].String()
		addr := strings.Split(s, "/")
		if naive.CgnatIp(addr[0]) {
			ip = append(ip, addr[0])
		}
	}

	if len(ip) == 0 {
		return
	}

	if !*all && len(ip) > 1 {
		fmt.Printf("More than 1 cgnat address is available. Use -all to print.")
		return
	}

	// print all ip addresses in cgnat range.
	// Seemingly awkward iteration to print 1 line result (without \n).
	// This way the result is easier to use in shell scripts.
	fmt.Printf("%v", ip[0])
	i := 1
	for ; i < len(ip); i++ {
		fmt.Printf("\n%v", ip[i])
	}
	if i > 1 {
		fmt.Println()
	}
}
