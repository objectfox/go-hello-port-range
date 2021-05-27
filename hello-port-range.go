package main

import (
    "fmt"
    "net/http"
    "strconv"
    "syscall"
    "flag"
)

func hello_world(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

func server(port int) {
    serverMux := http.NewServeMux()
    serverMux.HandleFunc("/", hello_world)
    go func() {
        http.ListenAndServe(":"+strconv.Itoa(port), serverMux)
    }()
}

func main() {

    // Declare our args.
    end := flag.Int("end", 9000, "end of the port range")
    start := flag.Int("start", 8000, "start of the port range")
    flag.Parse()

    // Set the limit on the number of open files so we can listen
    // on this many ports.
    var rLimit syscall.Rlimit
    rLimit.Max = 1024
    rLimit.Cur = 1024
    err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        fmt.Println("Error Setting Rlimit ", err)
    }

    // Create new http servers on ports 8000-8999
    for i := *start; i < *end; i++ {
        server(i)
    }

    fmt.Println("Running server on port "+strconv.Itoa(*start)+" to "+strconv.Itoa(*end))

    // Listen on our last port so we don't exit.
    serverMux := http.NewServeMux()
    serverMux.HandleFunc("/", hello_world)
    http.ListenAndServe(":"+strconv.Itoa(*end), serverMux)
}
