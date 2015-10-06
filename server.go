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
)

// global parameter :D
var counter int = 0
var help_msg string = `<simple-server-watch-openfd-s cleint|server>`

// this func will check for open file discriptor of the pid
func check_open_fd(pid int) uint64 {
    //get the process id
    // pid := os.Getpid()
    // fmt.Println(pid)

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

    // log.Println("contents: ", contents)

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
    l, err := net.Listen("tcp", ":9999")
    if err != nil {
        log.Fatal("net.Listen | ", err)
    }
    return l
}

func server_func() {
    // if server run the server
    l := start_listener()
    defer l.Close()
    log.Println("the listener started ")

    // now we need to accept the comming req
    // but B4 that we need to check for openfd's
    for {
        // refresh the number of openfd's
        nofile = check_open_fd(pid)
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

    // initialize
    pid := os.Getpid()

    // first let's start the server
    // but before that we need to watch the open fd's
    // we have to check the folder /proc/$PID/fd
    // and count the links there
    // so lets check the open fd's
    nofile := check_open_fd(pid)

    // check for max_allowed ?
    var rlimit syscall.Rlimit
    err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
    if err != nil {
        log.Fatal( "syscall.Getrlimit | ", err )
    }

    // the rlimit.cur is the one we need :D
    log.Println("max allowed : ", rlimit.Cur)

    // get what the user need
    // before that what are the args ?
    log.Println(os.Args)
    // os args should not be more than 2 or less than 2
    if len(os.Args) > 2 || len(os.Args) < 2 {
        log.Fatal(help_msg)
    }
    if os.Args[1] == "client" {
        // if client , run test the client
        client()
    } else if os.Args[1] == "server" {
        server_func()
    } else {
        log.Println("Worng argument")
        log.Println("use like this ")
        log.Println("simple* <client | server>")
    }
}
