package connection

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func OpenConnection(masterHostAddress string, port string) net.Conn {
	conn, err := net.Dial("tcp", masterHostAddress)
	if err != nil {
		fmt.Println("Faild to connect to master " + masterHostAddress)
		return conn
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// initial handshake
	conn.Write(WriteBulkString([]string{"PING"}))
	reader.ReadString('\n')
	conn.Write(WriteBulkString([]string{"REPLCONF", "listening-port", port}))
	reader.ReadString('\n')
	conn.Write(WriteBulkString([]string{"REPLCONF", "capa", "psync2"}))
	reader.ReadString('\n')
	conn.Write(WriteBulkString([]string{"PSYNC", "?", "-1"}))
	return conn
}

func WriteBulkString(elements []string) []byte {
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
