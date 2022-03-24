package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const(
	CommandEcho = "echo"
	CommandCd = "cd"
	CommandKill = "kill"
	CommandPwd = "pwd"
	CommandExit = "quit"
	CommandPs = "ps"
	ExitText = "Exit"
)

func Echo(args... string) ([]byte, error) {
	return exec.Command("echo", args...).Output()
}

func Pwd() ([]byte, error) {
	dir, err := os.Getwd()
	return []byte(dir), err
}

func Cd(dir string) ([]byte, error) {
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return []byte(dir), nil
}

func Ps() ([]byte, error) {
	return exec.Command("ps").Output()
}

func Kill(args... string) ([]byte, error) {
	return exec.Command("kill", args...).Output()
}

func ExecuteCommands(commands []string, w io.Writer)  {
	for _, command := range commands {
		args := strings.Split(command, " ")

		com := args[0]
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case CommandEcho:
			b, err := Echo(args...)
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}

			_, err = fmt.Fprintln(w, string(b))
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
		case CommandCd:
			b, err := Cd(args[0])
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}

			_, err = fmt.Fprintln(w, string(b))
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
		case CommandKill:
			b, err := Kill(args...)
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}

			_, err = fmt.Fprintln(w, string(b))
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
		case CommandPwd:
			b, err := Pwd()
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}

			_, err = fmt.Fprintln(w, string(b))
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
		case CommandPs:
			b, err := Ps()
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}

			_, err = fmt.Fprintln(w, string(b))
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
		case CommandExit:
			_, err := fmt.Fprintln(w, ExitText)
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
			os.Exit(1)
		}
	}
}



func main()  {
	scan := bufio.NewScanner(os.Stdin)
	var output = os.Stdout

	for {
		fmt.Print("command: ")
		if scan.Scan() {
			line := scan.Text()
			commands := strings.Split(line, " | ")
			ExecuteCommands(commands, output)
		}
	}
}