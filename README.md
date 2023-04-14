# k8s-hello-world
Hello world HTTP server app for Kubernetes/Openshift

## Run
```bash
go run main.go

curl http://localhost:8080
```

You can change the server port using "PORT" env variable. Default port is 8080.
```bash
export PORT=8080
```
