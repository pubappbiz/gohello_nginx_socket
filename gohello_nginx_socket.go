package main

import (
    "syscall"
    "fmt"
    "net"
    "net/http"
    "net/http/fcgi"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Hello World! on Nginx socket")
}

func main() {
    sock := "/tmp/gohello-nginx-socket.sock"

    umaskDefault := syscall.Umask(0)
    l, err := net.Listen("unix", sock)
    syscall.Umask(umaskDefault)
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
    http.HandleFunc("/", mainHandler)
    fcgi.Serve(l, nil)
}
