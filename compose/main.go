package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/unixpickle/poeturn/model"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var netStart int
	var netFile string
	flag.IntVar(&netStart, "netstart", 0, "initial lines to generate")
	flag.StringVar(&netFile, "network", "", "path to the network")

	flag.Parse()
	if netFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	net, err := model.LoadModel(netFile)
	if err != nil {
		die("Load model:", err)
	}

	sess := model.NewSession(net)

	for netStart > 0 {
		fmt.Println(sess.Query())
		netStart--
	}
	for {
		sess.Dictate(readLine())
		fmt.Println(sess.Query())
	}
}

func readLine() string {
	fmt.Print("> ")
	var res string
	for {
		var b [1]byte
		if n, err := os.Stdin.Read(b[:]); err != nil {
			panic(err)
		} else if n == 1 {
			if b[0] == '\n' {
				break
			} else if b[0] != '\r' {
				res += string(b[:])
			}
		}
	}
	return res
}

func die(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}
