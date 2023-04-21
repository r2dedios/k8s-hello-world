package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
)

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("getHostname: " + err.Error())
		return ""
	}
	return hostname
}

func getUser() string {
	user, err := user.Current()
	if err != nil {
		fmt.Println("getUser: " + err.Error())
		return "Not User Found"
	}
	return user.Username + "; (UID: " + user.Uid + "; GID:" + user.Gid + ")"
}

func getIPList(hostname string) string {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println("getIPList: " + err.Error())
		return "Not IPs Found"
	}
	ipList := ""
	for _, ip := range ips {
		ipList += "<div> Host IPs: " + ip.String() + "</div>"
	}
	return ipList
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	response := `
<html>
<head><title>Hello Pod</title></head>
<body>
<div>Ok</div>
</body>
</html>
`
	fmt.Fprintln(w, response)
	fmt.Println("Healthz requested")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname := getHostname()
	user := getUser()
	ipList := getIPList(hostname)
	response := `
<html>
<head><title>Hello Pod</title></head>
<body>
<h1>Hello World Pod!</h1>
<div> Hostname: ` + hostname + ` </div>
` + ipList + `
<div> User: ` + user + ` </div>
</body>
</html>
`

	fmt.Fprintln(w, response)
	fmt.Println("Saying hello!")
}

func listenAndServe(port string) {
	fmt.Printf("Serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}
