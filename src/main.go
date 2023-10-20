package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strconv"

	gbs "github.com/inhies/go-bytesize"
	"github.com/pbnjay/memory"
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
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("getIPList: " + err.Error())
		return "Not IPs Found"
	}
	ipList := ""
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, a := range addrs {
			ipList += fmt.Sprintf("<li> iface [%v]:  {%s}  (%s) </li>", i.Name, a.String(), a.Network())
		}
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

func getCPUCoresCount() string {
	return strconv.Itoa(runtime.NumCPU())
}

func getRAMCapacity() string {
	bytes := gbs.New(float64(memory.TotalMemory()))
	return bytes.String()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname := getHostname()
	user := getUser()
	ipList := getIPList(hostname)
	cpus := getCPUCoresCount()
	memory := getRAMCapacity()
	response := `
<html>
<head><title>Hello Pod</title></head>
<body>
<h1>Hello World Pod!</h1>
<div> <b>Hostname:</b> ` + hostname + ` </div>
<div> <b>Network Interfaces:</b>  </div>
<ul>
` + ipList + `
</ul>
<div> <b>User:</b> ` + user + ` </div>
<div> <b>CPUs:</b> ` + cpus + ` </div>
<div> <b>Memory:</b> ` + memory + ` </div>
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
