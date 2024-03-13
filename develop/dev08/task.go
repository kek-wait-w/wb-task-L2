package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

//функция обработки команд
func CoamndHandleer(stringCommand string) string {
	var id uintptr
	var res string
	//при наличии флага -d создается daemon процесс
	daemon := strings.Contains(stringCommand, "-d")
	if daemon {
		stringCommand = strings.Replace(stringCommand, "-d", "", 1)

		id, _, _ = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if id == 0 {

		} else {
			fmt.Println("Daemon created, PID :", id)
			return ""
		}
	}

	switch strings.Split(stringCommand, " ")[0] {
	//в зависимости от парсинга выбираем команду
	case "cd":
		res = chDirCommand(stringCommand)
	case "pwd":
		res = pwdCommand()
	case "echo":
		res = echoCommand(stringCommand)
	case "kill":
		killPsCommand(stringCommand)
	case "ps":
		res = psCommand()
	case `\quit`:
		exitCommand()
	default:
		fmt.Println("Invalid command")
		if daemon {
			fmt.Printf("Process %v killed\n", os.Getpid())
			killPsCommand(strconv.Itoa(os.Getpid()))
		}
	}
	return res
}

//функция команды директории
func chDirCommand(stringCommand string) string {

	err := os.Chdir(strings.TrimSpace(strings.Replace(stringCommand, "cd", "", 1)))
	if err != nil {
		fmt.Println(err)
	}
	return pwdCommand()
}

//функция получение текущей директории
func pwdCommand() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	return dir
}

//функция эхо
func echoCommand(stringCommand string) string {
	fmt.Println(strings.TrimSpace(strings.Replace(stringCommand, "echo", "", 1)))
	return strings.TrimSpace(strings.Replace(stringCommand, "echo", "", 1))
}

//функция убийства процесса
func killPsCommand(stringCommand string) {
	pid, err := strconv.Atoi(strings.Replace(stringCommand, "kill ", "", 1))
	if err != nil {
		fmt.Println(err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	err = proc.Kill()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Process %v killed\n", pid)
}

//функция получения запущенных процессов
func psCommand() string {
	c := exec.Command("ps")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	//scanner := bufio.NewScanner(c.Stdin)
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	return ""

}

//функция выхода
func exitCommand() {
	fmt.Println("Exit")
	os.Exit(0)
}

//основная функция консоли
func CMD() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		CoamndHandleer(scanner.Text())
	}
}
func main() {
	CMD()
}
