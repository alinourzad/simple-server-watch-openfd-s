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

// global counter :D
var counter int = 0

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

    // we shold watch for open fd's
    // if the openfd's have been reached we are done
    // and we should not accept any more connection
    // so start_server() will only launch net.Listen()
    // function
    // after we get the current openfd we can run accept
    // and run go()

    //start the listener
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
