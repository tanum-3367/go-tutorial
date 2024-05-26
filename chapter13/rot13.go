package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func rot13(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		switch {
		case char >= 'a' && char <= 'z':
			result[i] = 'a' + (char-'a'+13)%26
		case char >= 'A' && char <= 'Z':
			result[i] = 'A' + (char-'A'+13)%26
		default:
			result[i] = char
		}
	}
	return string(result)
}

func processStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("An error occurred: ", err)
			return
		}
		encoded := rot13(input)
		fmt.Print(encoded)
	}
}

func processFileOrInput() {
	var inputReader io.Reader

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("An error occurred: ", err)
			return
		}
		defer file.Close()
	} else {
		// No file provided, use input instead
		fmt.Println("Please enter text to encode/decode: ")
		inputReader = os.Stdin
	}

	// process input and apply rot13 encoding
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		encoded := rot13(scanner.Text())
		fmt.Println("Encoded: ", encoded)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("An error occurred: ", err)
	}
}

func main() {
	// check if data is available on stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data available on stdin
		processStdin()
	} else {
		processFileOrInput()
	}
}
