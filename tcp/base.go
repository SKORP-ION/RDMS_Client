package tcp

import (
	"RDMS_Client/structures"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const BUFFERSIZE = 4096

type Connection struct {
	conn net.Conn
}

func (this *Connection) Connect() error {
	conn, err := net.Dial("tcp", os.Getenv("tcp_host"))

	if err != nil {
		return err
	}

	this.conn = conn

	return nil
}

func (this *Connection) sendSessionKey(key string) error {
	conn := this.conn
	bytesKey := []byte(key + "\n")
	_, err := conn.Write(bytesKey)

	return err
}

func (this *Connection) ReceivePackage(pkg *structures.Package, sessionKey string) error {
	err := this.sendSessionKey(sessionKey)

	if err != nil {
		return err
	}


	filename := fmt.Sprintf("%s_%s.deb", pkg.Name, pkg.Version)
	path := "packages/" + filename
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	conn := this.conn
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		buffer := make([]byte, BUFFERSIZE)
		n, err := reader.Read(buffer)
		if err == io.EOF{
			break
		}
		_, err = file.Write(buffer[:n])

		if err != nil {
			return err
		}
	}

	return nil
}