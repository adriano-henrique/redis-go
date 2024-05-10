package connection

import (
	"fmt"
	"net"
	"strings"
)

func OpenConnection(masterHostAddress string) {
	conn, err := net.Dial("tcp", masterHostAddress)
	if err != nil {
		fmt.Println("Faild to connect to master " + masterHostAddress)
		return
	}
	defer conn.Close()
	err = sendPing(conn)
	if err != nil {
		fmt.Println("Failed to send ping to master " + masterHostAddress)
		return
	}
}

func sendPing(conn net.Conn) error {
	_, err := conn.Write(writeBulkString([]string{"PING"}))
	if err != nil {
		return err
	}
	return nil
}

func writeBulkString(elements []string) []byte {
	numElements := len(elements)
	var responseBuilder strings.Builder
	prefix := fmt.Sprintf("*%d\r\n", numElements)
	responseBuilder.WriteString(prefix)
	var responseBodyBuilder strings.Builder
	for _, element := range elements {
		responseBodyBuilder.WriteString(encodeRedisBulkString(element))
	}
	responseBody := responseBodyBuilder.String()
	responseBuilder.WriteString(responseBody)
	return []byte(responseBuilder.String())
}

func encodeRedisBulkString(value string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}
