Command ip100 prints all local ip addresses in 100.64.0.0/10 subnet (details: rfc6598).
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
$ ip100=`ip100` ufs -addr $ip100:5640

# improvise a radio station from your own music:
$ go get github.com/fgergo/boringstreamer
$ ip100=`ip100` boringstreamer -addr $ip100:4444