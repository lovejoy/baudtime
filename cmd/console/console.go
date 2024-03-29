/*
 * Copyright 2019 The Baudtime Authors
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/baudtime/baudtime/promql"
	"github.com/peterh/liner"
)

var (
	currentUser, _ = user.Current()
	historyFile    = filepath.Join(currentUser.HomeDir, ".baudtime")
	ip             = flag.String("h", "127.0.0.1", "baudtime server ip (default 127.0.0.1)")
	port           = flag.Int("p", 8088, "baudtime server port (default 8088)")
	queryTimeout   = 120 * time.Second
)

var line *liner.State

func main() {
	flag.Parse()

	line = liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	setAutoCompletionHandler()
	loadHistory()

	defer saveHistory()

	var addr = fmt.Sprintf("%s:%d", *ip, *port)

	reg, _ := regexp.Compile(`'.*?'|".*?"|\S+`)
	prompt := fmt.Sprintf("%s> ", addr)

	exec := &executor{
		addr:        addr,
		queryEngine: promql.NewEngine(nil, 20, queryTimeout),
	}
	err := exec.reconnect()
	if err != nil {
		fmt.Println(err)
		return
	}

	for !exec.closed {
		cmd, err := line.Prompt(prompt)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		cmds := reg.FindAllString(cmd, -1)
		if len(cmds) == 0 {
			continue
		} else {
			line.AppendHistory(cmd)

			args := make([]string, len(cmds[1:]))

			for i := range args {
				args[i] = strings.Trim(cmds[1+i], "\"'")
			}

			cmd := strings.ToLower(cmds[0])
			err = exec.execCommand(cmd, args...)
			if checkConnBroken(err) {
				fmt.Print("\n\nTry to reconnect...\n\n")
				exec.reconnect()
				exec.execCommand(cmd, args...)
			}
		}
	}
}

func printGenericHelp() {
	msg :=
		`baudtime-cli
Type:	"help <command>" for help on <command>
	`
	fmt.Println(msg)
}

func printCommandHelp(cmd string) {
	cmd = strings.ToUpper(cmd)
	for i := 0; i < len(helpCommands); i++ {
		if helpCommands[i][0] == cmd {
			fmt.Println()
			fmt.Printf("Usage:\n\t%s %s \n", helpCommands[i][0], helpCommands[i][1])
			fmt.Printf("Description:\n\t %s \n", helpCommands[i][2])
			fmt.Println()
		}
	}
}

func printHelp(args []string) {
	if len(args) == 0 {
		printGenericHelp()
	} else if len(args) > 1 {
		fmt.Println()
	} else {
		printCommandHelp(args[0])
	}
}

func setAutoCompletionHandler() {
	line.SetCompleter(func(line string) (c []string) {
		for _, i := range helpCommands {
			cmd := strings.ToLower(i[0])
			if strings.HasPrefix(cmd, strings.ToLower(line)) {
				c = append(c, cmd)
			}
		}
		return
	})
}

func loadHistory() {
	if f, err := os.Open(historyFile); err == nil {
		line.ReadHistory(f)
		f.Close()
	}
}

func saveHistory() {
	if f, err := os.Create(historyFile); err != nil {
		fmt.Printf("Error writing history file, err: %v", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}

func checkConnBroken(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(net.Error)
	return ok
}
