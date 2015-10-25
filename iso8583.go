package main

import (
    "fmt"
)

func analyze_iso_msg(data []byte) {
    DumpHex(data)
}

func DumpHex(data []byte) error {
	if !DEBUG {
		return nil
	}
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
