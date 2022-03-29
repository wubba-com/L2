package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/**
8. Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
*/

const (
	CommandEcho = "echo"
	CommandCd   = "cd"
	CommandKill = "kill"
	CommandPwd  = "pwd"
	CommandExit = "quit"
	CommandPs   = "ps"
	ExitText    = "Exit"
)

/**
Пакет exec запускает внешние команды. Он обертывает os.StartProcess,
чтобы сделать его проще переназначить stdin и stdout,
соединить ввод /вывод с помощью каналов и сделать другие корректировки.
*/

// Echo выполняет unix-команду echo и возвращает результат в байтах
func Echo(args ...string) ([]byte, error) {
	return exec.Command("echo", args...).Output()
}

// Pwd - выводит путь директории в которой находится терминал
func Pwd() ([]byte, error) {
	dir, err := os.Getwd()
	return []byte(dir), err
}

// Cd - изменяет директорию
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

// Ps - Выводит работающие процессы
func Ps() ([]byte, error) {
	return exec.Command("ps").Output()
}

// Kill - убивает запущенный процесс
func Kill(args ...string) ([]byte, error) {
	return exec.Command("kill", args...).Output()
}

// ExecuteCommands Исполняет команды, которые ввел пользователь
func ExecuteCommands(cmds []string, w io.Writer) {
	for _, cmd := range cmds {
		args := strings.Split(cmd, " ")

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
		case CommandExit: // завершение программы
			_, err := fmt.Fprintln(w, ExitText)
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
			os.Exit(1)
		}
	}
}

func main() {
	// Читает из стандартного ввода
	scan := bufio.NewScanner(os.Stdin)

	// устанавливается общий вывод результата команд
	var output = os.Stdout

	for {
		fmt.Print("command: ")

		if scan.Scan() {
			line := scan.Text()
			cmds := strings.Split(line, " | ")
			ExecuteCommands(cmds, output)
		}
	}
}
