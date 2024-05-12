package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type HTTPRequest struct {
	StartLine StartLine
	Headers   map[string]string
}

type StartLine struct {
	Method  string
	Path    string
	Version string
}

func ParseHTTPRequest(r *bufio.Reader) (*HTTPRequest, error) {
	// read the first line of the request
	startLine, err := readStartLine(r)
	if err != nil {
		return nil, err
	}

	// read the headers
	headers, err := readHeaders(r)
	if err != nil {
		return nil, err
	}

	return &HTTPRequest{
		StartLine: startLine,
		Headers:   headers,
	}, nil
}

// readLine reads a line from the input stream and returns the line, the number of bytes read, and any error encountered.
func readLine(br *bufio.Reader) (line []byte, n int, err error) {
	for {
		b, err := br.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n++
		line = append(line, b)
		if len(line) >= 2 &&
			line[len(line)-2] == '\r' &&
			line[len(line)-1] == '\n' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func readStartLine(r *bufio.Reader) (StartLine, error) {
	var startLine StartLine

	br := bufio.NewReader(r)
	line, _, err := readLine(br)
	if err != nil {
		return startLine, err
	}

	parts := strings.Split(string(line), " ")
	if len(parts) != 3 {
		return startLine, fmt.Errorf("Invalid start line: %s", line)
	}

	startLine.Method = parts[0]
	startLine.Path = parts[1]
	startLine.Version = parts[2]

	return startLine, nil
}

func readHeaders(r *bufio.Reader) (map[string]string, error) {
	headers := make(map[string]string)
	for {
		line, _, err := readLine(r)
		if err != nil {
			return headers, err
		}
		if len(line) == 0 {
			break
		}
		parts := strings.Split(string(line), ": ")
		if len(parts) != 2 {
			return headers, fmt.Errorf("Invalid header: %s", line)
		}
		headers[parts[0]] = parts[1]
	}
	return headers, nil
}

type HTTPResponse struct {
	StatusLine string
	Headers    map[string]string
	Body       string
}

// Bytes returns the byte representation of the HTTPResponse
func (r *HTTPResponse) Bytes() []byte {
	var res bytes.Buffer
	res.WriteString(r.StatusLine + "\r\n")
	for k, v := range r.Headers {
		res.WriteString(k + ": " + v + "\r\n")
	}
	// marks the end of the headers
	res.WriteString("\r\n")
	res.WriteString(r.Body)
	return res.Bytes()
}

var StatusOK = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s"

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	req := make([]byte, 1024)

	_, err = conn.Read(req)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}

	httpReq, err := ParseHTTPRequest(bufio.NewReader(bytes.NewReader(req)))
	if err != nil {
		fmt.Println("Error parsing request: ", err.Error())
		os.Exit(1)
	}

	log.Printf("Request: %+v", httpReq)

	if httpReq.StartLine.Path == "/" {
		response := HTTPResponse{
			StatusLine: "HTTP/1.1 200 OK",
		}
		_, err = conn.Write(response.Bytes())
		if err != nil {
			fmt.Println("Error writing response: ", err.Error())
		}
		return
	}

	_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
	}

	conn.Close()

}
