package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	version := os.Getenv("VER_TAG")
	message := os.Getenv("CONFIG_MESSAGE")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		host, _ := os.Hostname()
		num, err := loadRandomNum()
		out := fmt.Sprintf(
			"Messsage:\t%s\nVersion:\t%s\nHost:\t%s\nRandom num:\t%d\nCPU:\t%d\nGOMAXPROCS:\t%d\n",
			message, version, host, num, runtime.NumCPU(), runtime.GOMAXPROCS(0))

		if err != nil {
			out += "Error:\t" + err.Error()
		}

		_, _ = writer.Write([]byte(out))
	})

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		rand.Seed(time.Now().UnixNano())

		if rand.Intn(5) == 0 {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.WriteHeader(http.StatusAccepted)
		}
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}

func loadRandomNum() (int, error) {
	c := http.DefaultClient
	resp, err := c.Get("http://num-svc:8081/number")
	if err != nil {
		return 0, fmt.Errorf("failed to call num-svc:8081: %w", err)
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response from num-svc:8081: %w", err)
	}

	s := string(all)

	atoi, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to covert response to int: %w", err)
	}

	return atoi, nil
}
