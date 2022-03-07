package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {
	var (
		host    = flag.String("host", "", "Hostname to check")
		port    = flag.String("port", "", "Port")
		timeout = flag.Int("timeout", 15, "Timeout in seconds")
		quite   = flag.Bool("quite", false, "Disable logging")
		help    = flag.Bool("help", false, "Print usage")
	)
	flag.Parse()

	var timeoutVar = time.Duration(*timeout) * time.Second

	if *help {
		fmt.Fprintf(os.Stderr, "wait-for [options] Run command with args after the test finishes\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if *quite {
		log.SetOutput(ioutil.Discard)
	}

	// Ensure the minimum arguments were provided.
	if *host == "" || *port == "" {
		log.Fatalln("-host, -port must be provided")
	}

	d := net.Dialer{Timeout: timeoutVar}
	conn, err := d.Dial("tcp", *host+":"+*port)
	if err != nil {
		log.Fatalf("Timeout occurred after waiting %s ", timeoutVar)
	}
	defer conn.Close()

	args := flag.Args()
	// Only call function if args were provided
	if len(args) > 0 {
		runCommand(args)
	}
}

func runCommand(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(":::::::::::::::::::::::::::::: STDERR ::::::::::::::::::::::::::::::")
		log.Printf("Failed to run %v: %v", args, err)
		log.Printf("Command output was:\n%s", out)
		log.Println(":::::::::::::::::::::::::::::: STDERR ::::::::::::::::::::::::::::::")
	} else {
		log.Println(":::::::::::::::::::::::::::::: STDOUT ::::::::::::::::::::::::::::::")
		log.Printf("Running command: %+q\n", args)
		log.Printf(string(out))
		log.Println(":::::::::::::::::::::::::::::: STDOUT ::::::::::::::::::::::::::::::")
	}
	return err
}
