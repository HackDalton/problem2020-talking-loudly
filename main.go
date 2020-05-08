package main

import (
	"encoding/base64"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const flag = "hackDalton{gr3p_1s_4_us3ful_t00l_si1EldtNcQ}"

var parts = splitString(flag, 6)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	currentPart := 0
	for currentPart < len(parts) {
		if rand.Intn(400) == 42 {
			// yay they get another part of the flag
			connection.Write([]byte("part " + strconv.Itoa(currentPart) + " of " + strconv.Itoa(len(parts)) + " of the flag is: " + parts[currentPart] + "\n"))
			currentPart++
		} else {
			str, err := generateRandomString(40)
			if err != nil {
				panic(err)
			}
			connection.Write([]byte(str + "\n"))
		}
	}

	connection.Close()
}

func splitString(s string, parts int) []string {
	result := make([]string, parts)
	partLen := len(s) / parts
	for i := 0; i < parts; i++ {
		result[i] = s[i*partLen : ((i + 1) * partLen)]
	}
	result[parts-1] = s[(parts-1)*partLen:]
	return result
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)[:s], err
}
