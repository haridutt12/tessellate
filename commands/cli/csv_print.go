package main

import (
	"fmt"
	"os"
)

func csvPrintHeader(headers []string) {
	headersStr := ""
	for _, header := range headers {
		headersStr = headersStr + header + "|"
	}
	if len(headersStr) > 0 {
		fmt.Fprintf(os.Stdout, headersStr[:len(headersStr)-1]+"\n")
	}
}

func csvPrint(data map[string]string) {
	dataStr := ""
	for _, value := range data {
		dataStr = dataStr + value + "|"
	}
	if len(dataStr) > 0 {
		fmt.Fprintf(os.Stdout, dataStr[:len(dataStr)-1]+"\n")
	}
}
