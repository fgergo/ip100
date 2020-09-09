// Command ip100 prints all local ip addresses in 100.64.0.0/10 subnet.
// That's between 100.64.0.0 and 100.127.255.255.
// Background: https://tailscale.com/kb/1015/100.x-addresses
//
// Command ip100 is supposed to promote naive servers
// without real authentication or authorization.
// Background: https://crawshaw.io/blog/remembering-the-lan
//
// Also see
//
// $ ip100 -ideas
//
// for ideas you might want to try with or without ip100.
package main
