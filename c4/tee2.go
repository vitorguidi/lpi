package main

import (
	"fmt"
	"os"
	"syscall"
)

func tee() {
	// Ler todos os arquivos vindo do argumento da linha de comando
	fileNames := os.Args[1:]
	fmt.Println("PID:", os.Getpid())

	shouldAppend := false
	// Procurar flag -a no argumento da linha de comando, se encontrar, setar shouldAppend para true e remover de
	// fileNames
	for i, fileName := range fileNames {
		if fileName == "-a" {
			shouldAppend = true
			fileNames = append(fileNames[:i], fileNames[i+1:]...)
		}
	}

	// Abrir um fd para cada arquivo
	fds := make([]int, len(fileNames))
	for i, fileName := range fileNames {
		mode := syscall.O_WRONLY | syscall.O_CREAT
		if shouldAppend {
			mode |= syscall.O_APPEND
		} else {
			mode |= syscall.O_TRUNC
		}

		fd, err := syscall.Open(fileName, mode, 0666)
		if err != nil {
			fmt.Errorf("Error opening file %s: %v", fileName, err)
		}
		defer syscall.Close(fd)
		fds[i] = fd
	}

	// Ler do stdin usando apenas syscalls
	var input string
	for {
		var buf [1]byte
		isEof, err := syscall.Read(0, buf[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		if isEof == 0 {
			break
		}
		input += string(buf[:])
	}

	// Escrever nos arquivos
	for _, fd := range fds {
		_, err := syscall.Write(fd, []byte(input))
		if err != nil {
			fmt.Errorf("Error writing to file: %v", err)
		}
	}
}

func main() {
	tee()
	os.Exit(0)
}
