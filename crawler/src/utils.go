package main

import (
	"log"
	"os"
	"strings"
)

func checkError(msg string, err error, fatal bool) bool {
	if err != nil {
		if fatal {
			log.Println("FATAL ERROR:", msg)
			os.Exit(1)
		} else {
			log.Println("ERROR:", msg)
		}
	}
	return err == nil
}

func getDomain(host string) string {
	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return host
	}
	return parts[len(parts)-2] + "." + parts[len(parts)-1]
}
