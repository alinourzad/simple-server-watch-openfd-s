package main

import (
    "fmt"
)

func analyze_iso_msg(data []byte) {
    DumpHex(data)
    // if it is bcd we need to have different approach
    // if it is ascii we need to have another approach
    // so we need 2 function bcd to ascii and the other way around .
}

func DumpHex(data []byte) error {
	for i := 0; i < len(data); i++ {
		if (i+1)%25 != 0 {
			fmt.Printf("%02x ", data[i])
		} else {
			fmt.Printf("%02x\n", data[i])
		}
	}
	fmt.Printf("\n")
	return nil
}
