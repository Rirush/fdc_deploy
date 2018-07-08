package fdc_protocol

import "strings"

type Command byte

// Command opcodes
const (
	_ = Command(iota)
	// Checks connection with server
	COMMAND_PING
	// Not an actual command, server should reply with COMMAND_ERROR if it can't process a command
	COMMAND_ERROR
	// Request all available apps for deployment
	COMMAND_LIST_APPS
	// Create and deploy new application
	COMMAND_DEPLOY_NEW
	// Deploy new version for existing app
	COMMAND_DEPLOY_UPDATE
	// Rollback to previous version of existing app
	COMMAND_DEPLOY_ROLLBACK
	// Delete application from server and stop all containers
	COMMAND_DEPLOY_DELETE
	// Request application version
	COMMAND_APP_VERSION
	// Stop application containers
	COMMAND_APP_STOP
	// Start application containers
	COMMAND_APP_START
	// Restart application containers
	COMMAND_APP_RESTART
	// Request application containers status
	COMMAND_APP_STATUS
)

type Package struct {
	Command   Command
	Arguments []string
}

func NewPackage(command Command, args ...string) Package {
	p := Package{
		Command:   command,
		Arguments: make([]string, 0),
	}
	for _, arg := range args {
		p.Arguments = append(p.Arguments, arg)
	}
	return p
}

func Marshal(p Package) []byte {
	res := make([]byte, 0)
	res = append(res, byte(p.Command))
	for _, arg := range p.Arguments {
		res = append(res, []byte(arg)...)
		res = append(res, 0)
	}
	return append(res[:len(res)-1], byte('\n'))
}

func Unmarshal(data []byte, p *Package) {
	if len(data) == 0 {
		return
	}
	// Package must end with \n
	if data[len(data)-1] != '\n' {
		return
	}
	p.Command = Command(data[0])
	argsData := string(data[1:])
	p.Arguments = strings.Split(argsData, "\x00")
	lr := p.Arguments[len(p.Arguments)-1]
	p.Arguments[len(p.Arguments)-1] = lr[:len(lr)-1]
}
