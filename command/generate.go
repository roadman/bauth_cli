package command

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// ApplyCommand is a Command that generates a new cli project
type GenerateCommand struct {
	Meta
}

func (c *GenerateCommand) Run(args []string) int {

	cmdsHtpasswd := []string{
		"/usr/sbin/htpasswd",
		"/usr/bin/htpasswd",
	}

	var generateNum int

	argsNum := len(args)
	if argsNum == 1 {
		generateNum = 1
	} else if argsNum == 2 {
		generateNum, _ = strconv.Atoi(args[1])
		if generateNum == 0 {
			c.Ui.Error("arg error")
			return 1
		}
	} else {
		c.Ui.Error("arg error")
		return 1
	}

	userString := args[0]

	var cmdHtpasswd string
	for _, cmd := range cmdsHtpasswd {
		if Exists(cmd) {
			cmdHtpasswd = cmd
		}
	}
	if cmdHtpasswd == "" {
		c.Ui.Error("htpassword command search miss.")
		return 1
	}

	fmt.Printf("# group: %s, created: %s\n", userString, time.Now())

	var user string
	if generateNum > 1 {
		for cnt := 0; cnt < generateNum; cnt++ {
			user = fmt.Sprintf("%s%03d", userString, cnt+1)
			genStringNum(cmdHtpasswd, userString, user)
		}
	} else {
		genStringNum(cmdHtpasswd, userString, userString)
	}

	return 0
}

func (c *GenerateCommand) Synopsis() string {
	return ""
}

func (c *GenerateCommand) Help() string {
	helpText := `
generate htpasswd command basic auth setting.

Usage:
	bauth_cli generate login num
`
	return strings.TrimSpace(helpText)
}

func randomString(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func genStringNum(cmd string, userString string, user string) {
	pass := randomString(8)
	out, err := exec.Command(cmd, "-nb", user, pass).Output()
	if err == nil {
		fmt.Printf("# login: %s , pass: %s\n", user, pass)
		fmt.Printf("%s\n", strings.TrimRight(string(out), "\n"))
	}
}
