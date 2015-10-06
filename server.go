package main

import (
    "os"
    "log"
    // "io/ioutil"
    // "fmt"
    "strconv"
    "path/filepath"
    "syscall"
    "net"
    "flag"
)

// global parameter :D
var counter int = 0
var pid     int

// this func will check for open file discriptor of the pid
func check_open_fd() uint64 {

    pid_string := strconv.Itoa(pid)
    // fmt.Println(pid_string)

    path := "/proc/" + pid_string + "/fd"
    // log.Println("the path is : ", path)

    err := os.Chdir(path)
    if err != nil {
        log.Fatal( "L:20 | os.Chdir | ", err )
    }

    contents, err := filepath.Glob("*")
    if err != nil {
        log.Fatal( "L:41 | filepath.Glob | ", err )
    }

    return uint64(len(contents))
}

//this func will handle the new connection and the stream they have
func handleConn(c net.Conn) {
    counter++
    log.Println(c.RemoteAddr())
    log.Println(c.LocalAddr())
    log.Println("-----", counter, "-----")
    defer c.Close()
}

// this function will start listener
func start_listener() net.Listener {
    // TODO: try to get port from user
    l, err := net.Listen("tcp", ":9999")
    if err != nil {
        log.Fatal("net.Listen | ", err)
    }
    return l
}

func server_func() {

    // check for max_allowed ?
    var rlimit syscall.Rlimit
    err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
    if err != nil {
        log.Fatal( "syscall.Getrlimit | ", err )
    }

    // if server run the server
    l := start_listener()
    defer l.Close()
    log.Println("the listener started ")

    // now we need to accept the comming req
    // but B4 that we need to check for openfd's
    for {
        // refresh the number of openfd's
        nofile := check_open_fd()
        // log.Printf("the new number of openfd's %d\n", nofile)
        // check them ?
        if rlimit.Cur > nofile {
            // if no problem accept the new connection
            c, err := l.Accept()
            if err != nil {
                log.Fatal("l.Accept " , err)
            }
            // defer c.Close()

            go handleConn(c)
        }
    }
}

// this is the main function
func main() {
    // get pid of this app
    pid = os.Getpid()

    // TODO: add the flags needed ?
    server_pbool := flag.Bool("server", false, "for running as sever")
    client_pbool := flag.Bool("client", false, "for testing the server")

    // parse the flags
    // but remember to add the flags before this
    flag.Parse()

    // we need to know if we should run as server or client
    if flag.NFlag() != 1 {
        flag.Usage()
        os.Exit(1)
    }
    if *client_pbool {
        // if client , run test the client
        client()
    } else if *server_pbool {
        server_func()
    } else {
        flag.Usage()
        os.Exit(1)
    }
}
