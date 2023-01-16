package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	println(os.Args[1])
	input := os.Args[1:]
	if len(input) == 0 {
		panic("No input")
	}
	append := false
	if input[0] == "-a" {
		append = true
		input = input[1:]
		println("-a is present")
	}
	fds := make([]int, len(input))

	go (func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
		}
	})("a")

	openFlags := 0
	if append {
		openFlags = syscall.O_WRONLY | syscall.O_APPEND | syscall.O_CREAT
	} else {
		openFlags = syscall.O_WRONLY | syscall.O_TRUNC | syscall.O_CREAT
	}

	for i, filename := range input {
		println(filename)
		fd, err := syscall.Open(filename, openFlags, 0666)
		if err != nil {
			panic(err)
		}
		fds[i] = fd
	}

	// keep reading from stdin untill eof
	buf := make([]byte, 1024)
	for {
		// read at most n bytes from stdin with the syscall read function
		n, err := syscall.Read(0, buf) // e se eu tentar escrever n>1024? Will read at most n from docs
		if err != nil {
			panic(err)
		}
		if n == 0 {
			break
		}
		for _, fd := range fds {
			writtenBytes, err := syscall.Write(fd, buf[:n]) // e se eu escrever menos que n? nao eh garantido
			if err != nil {
				panic(err)
			}
			if writtenBytes != n {
				panic("could not write all bytes")
			}
		}
	}

	for _, fd := range fds {
		syscall.Close(fd)
	}

}
