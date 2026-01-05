package main

import (
	"encoding/json"
	"fmt"
	"fundamentals/infra-observer/internal/systeminfo"
	"log"
)

func main() {
	info := systeminfo.Collect()
	out, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		log.Fatalf("marshal failes %v", err)
	}
	fmt.Print(string(out))
}
