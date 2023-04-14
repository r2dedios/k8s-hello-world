package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	ips, _ := net.LookupIP(hostname)
	ipList := ""
	for _, ip := range ips {
		ipList += "<div> Host IPs: " + ip.String() + "</div>"
	}
	response := `
<html>
<head><title>Hello Pod</title></head>
<body>
<h1>Hello World Pod!</h1>
<div> Hostname: ` + hostname + ` </div>
` + ipList + `
<div> User Name: ` + user.Username + ` </div>
<div> User ID: ` + user.Uid + ` </div>
</body>
</html>
`

	fmt.Fprintln(w, response)
	fmt.Println("Servicing an impatient beginner's request.")
}

func listenAndServe(port string) {
	fmt.Printf("Serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}
