package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	exitChan := make(chan os.Signal, 1)
	//При нажатии ctrl + D (SIGQUIT) в канал sigCh будет отправлено сообщение
	signal.Notify(exitChan, syscall.SIGQUIT)
	//горутинка функции выхода
	go Exit(exitChan)

	//Устанавливаем флаги на задержку и аргументы на хост и порт
	timeout := flag.String("timeout", "10s", "timeout for a connection")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Example: task.go host port")
		return
	}
	//Переводим задержку в нужный формат
	timeoutDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		return
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	hostPort := host + ":" + port
	start := time.Now()
	start = start.Add(timeoutDuration)
	// Подключаемся к сокету. Задаем таймаут подключения
	var conn net.Conn
	fmt.Printf("Пытаемся подключится к %s:%s...\n", host, port)
	for start.After(time.Now()) {
		conn, err = net.DialTimeout("tcp", hostPort, timeoutDuration)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		fmt.Println("Ошибка подключения к ", hostPort)
		return
	}
	fmt.Println("Соединение с сервером установлено.")
	defer conn.Close()

	//горутинка печати ответа с сервера
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("Сообщение с сервера: " + message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		_, err := fmt.Fprintf(conn, in+"\n")
		if err != nil {
			log.Fatal("Connection close")
		}
	}
}

//Ждем нужного сигнала для выхода из программы
func Exit(exitChan chan os.Signal) {
	for {
		switch <-exitChan {
		case syscall.SIGQUIT:
			os.Exit(0)
		default:
		}
	}
}
