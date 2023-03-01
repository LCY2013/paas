package config

import (
	"os"
	"strings"
)

var (
	conf *Config
)

func parseCommandLine() map[string]string {
	args := os.Args

	commandLine := make(map[string]string, len(args))
	for _, arg := range args {
		if !strings.Contains(arg, "=") {
			commandLine[arg] = arg
			continue
		}
		commandLineSplit := strings.Split(arg, "=")
		if len(commandLineSplit) < 2 {
			commandLine[commandLineSplit[0]] = commandLineSplit[0]
			continue
		}
		if !strings.Contains(arg, "--") {
			commandLine[commandLineSplit[0]] = commandLineSplit[1]
			continue
		}
		commandLine[strings.ReplaceAll(commandLineSplit[0], "--", "")] = commandLineSplit[1]
	}

	return commandLine
}

func init() {
	// using standard library "flag" package
	//flag.String("--registry", "consul", "registry type: consul, etcd...")
	//flag.String("--registry_address", "127.0.0.1:5000", "registry address: 127.0.0.1:5000,127.0.0.1:5000")
	//flag.String("--handler", "", "registry handler, eg: api, proto, http")
	//flag.String("--namespace", "go.micro.api", "registry namespace, default go.micro.api")
	//flag.String("--addr", "go.micro.api", "registry namespace, default go.micro.api")

	commandLine := parseCommandLine()

	registry := commandLine["registry"]
	registryAddress := commandLine["registry_address"]
	handler := commandLine["handler"]
	namespace := commandLine["namespace"]
	addr := commandLine["addr"]

	conf = &Config{
		RegistryType:      registry,
		RegisterAddresses: strings.Split(registryAddress, ","),
		Handler:           handler,
		Namespace:         namespace,
		Addr:              addr,
	}
}

type Config struct {
	RegistryType      string   `json:"registry,omitempty"`
	RegisterAddresses []string `json:"registry_address,omitempty"`
	Handler           string   `json:"handler,omitempty"`
	Namespace         string   `json:"namespace,omitempty"`
	Addr              string   `json:"addr,omitempty"`
}

func Conf() *Config {
	return conf
}
