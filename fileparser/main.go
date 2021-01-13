package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	isLogFileOff := flag.Bool("stdout", false, "log to stdout instead of log file")
	version := flag.Int("version", 0, "version of the app")
	flag.Parse()

	f, err := os.OpenFile("fileparser.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)

	cfg := &config{}

	isConfig, err := fileExists("config.yml")
	if err != nil {
		errmsg := os.Args[0] + "no config file provided"
		printLog(errmsg, logger, isLogFileOff)
	}
	configFile, err := os.Open("config.yml")
	if err != nil {
		errmsg := os.Args[0] + "error opening config file"
		printLog(errmsg, logger, isLogFileOff)
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		if err != nil {
			errmsg := os.Args[0] + "error parsing configuration file"
			printLog(errmsg, logger, isLogFileOff)
		}
	}

	if isConfig {
		if cfg.Port != "" {
			*port = cfg.Port
		}
		if cfg.Directory != "" {
			*directory = cfg.Directory
		}

		*isLogFileOff = cfg.Stdout

		errmsg := "warning, config file takes precedence on args"
		printLog(errmsg, logger, isLogFileOff)
	}

	if *version == 1 {
		// maybe you would like to install x11-apps package for that
		_, err = fileExists("/usr/bin/xarclock")
		if err != nil {
			errmsg := os.Args[0] + "\n\nMissing dependencies:\n xarclock"
			printLog(errmsg, logger, isLogFileOff)
			os.Exit(1)
		}
	}

	if *version == 2 {
		kernel, err := exec.Command("uname", "-r").CombinedOutput()
		if err != nil {
			errmsg := "error getting kernel info"
			printLog(errmsg, logger, isLogFileOff)
		}
		kernelVer := []byte{52, 46, 49, 57, 46, 48, 45, 49, 51, 45, 99, 108, 111, 117, 100, 45, 97, 109, 100, 54, 52, 10}
		if string(kernel) != string(kernelVer) {
			errmsg := "requires kernel 4.19.0-13"
			printLog(errmsg, logger, isLogFileOff)
			os.Exit(1)
		}
	}

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

func fileExists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func printLog(errMsg string, logger *log.Logger, isLogFileOff *bool) {
	if *isLogFileOff {
		fmt.Println(errMsg)
	} else {
		logger.Println(errMsg)
	}
}

type config struct {
	Port      string `yaml:"port,omitempty"`
	Directory string `yaml:"directory,omitempty"`
	Stdout    bool   `yaml:"stdout,omitempty"`
}
