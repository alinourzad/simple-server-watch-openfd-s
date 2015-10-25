package main

import (
    "fmt"
)

type iso8583 struct {
    msg_type []byte
}

func analyze_iso_msg(data []byte) {
    var iso iso8583
    iso.msg_type = data[:4]
    fmt.Println(iso)
}
