package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type Pair struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func main() {

	listOf := loadConfiguration("list.json")

	var wg sync.WaitGroup

	wg.Add(len(listOf))

	fmt.Println("#################################")
	for _, v := range listOf {
		go checkConnection(&wg, v.IP, v.Port)
	}
	wg.Wait()
	fmt.Println("#################################")
}

func checkConnection(wg *sync.WaitGroup, ip string, port int) {

	defer wg.Done()
	addr := fmt.Sprintf("%s:%d", ip, port)
	t := 4 * time.Second
	_, err := net.DialTimeout("tcp", addr, t)
	if err != nil {
		fmt.Printf("%v\n\n", err)
	} else {
		fmt.Printf("Successful connected to: %s\n\n", addr)
	}
}

func loadConfiguration(file string) []Pair {
	var l []Pair
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&l)
	return l
}
