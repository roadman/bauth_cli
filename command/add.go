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

type AddCommand struct {
	Meta
}

func (c *AddCommand) Run(args []string) int {
	cmdsHtpasswd := []string{
		"/usr/sbin/htpasswd",
		"/usr/bin/htpasswd",
	}

	argsNum := len(args)
	var generateNum int
	if argsNum == 0 {
		return 1
	} else if argsNum == 2 {
		generateNum, _ = strconv.Atoi(args[1])
		if generateNum == 0 {
			generateNum = 1
		}
	} else if argsNum > 2 {
		return 1
	}

	if len(args[0]) == 0 {
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
		return 1
	}

	fmt.Printf("# group: %s, created: %s\n", userString, time.Now())

	var pass string
	var user string
	for cnt := 0; cnt < generateNum; cnt++ {
		pass = randomString(8)
		user = fmt.Sprintf("%s%03d", userString, cnt+1)
		out, err := exec.Command(cmdHtpasswd, "-nb", user, pass).Output()
		if err != nil {
			continue
		}
		fmt.Printf("# pass: %s\n", pass)
		fmt.Printf("%s\n", strings.TrimRight(string(out), "\n"))
	}

	return 0
}

func (c *AddCommand) Synopsis() string {
	return ""
}

func (c *AddCommand) Help() string {
	helpText := `
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
