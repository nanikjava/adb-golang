/**
The app demonstrate on how to write adb client in golang to communicate with adb server.
The app communicate with local adbd server.

adb server have to be running before running the app, use command `adb server`
*/
package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

func main() {
	// TCP connect to the adb server
	netConn, err := net.Dial("tcp", "localhost:5037")
	if err != nil {
		fmt.Println("Connection error %v", err)
	}

	// send message to the adb server to get version information
	n, err := sendMsg("host:version", err, netConn)
	if err != nil {
		fmt.Println("Sending information error %v", err)
	}
	fmt.Println("Number of byte ", n)

	// Check whether the response is OKAY
	octet, err := checkOkayReply(err, netConn)
	if err != nil {
		fmt.Println("Checking OKAY error %v", err)
	}
	fmt.Println("Response result : ", string(octet))

	// The next will be the response length
	err, length := readResponseLength(err, netConn)
	if err != nil {
		fmt.Println("Reading response length error %v", err)
	}

	// Read the response sent by adb server based on the length given in
	// the previous step
	dataToRead, err := readRealResponse(length, err, netConn)
	if err != nil {
		fmt.Println("Reading real response error %v", err)
	}
	fmt.Println("Real response value :", string(dataToRead))

	// Close connection
	netConn.Close()


	// Open connection
	// TCP connect to the adb server
	netConn, err = net.Dial("tcp", "localhost:5037")
	if err != nil {
		fmt.Println("Connection error %v", err)
	}

	// send message to the adb server to use any transport to perform
	// the next operation
	n, err = sendMsg("host:transport-any", err, netConn)
	if err != nil {
		fmt.Println("Sending information error %v", err)
	}
	fmt.Println("Number of byte ", n)

	// check OKAY
	// Check whether the response is OKAY
	octet, err = checkOkayReply(err, netConn)
	if err != nil {
		fmt.Println("Checking OKAY error %v", err)
	}
	fmt.Println("Response result : ", string(octet))


	// send message to the adb server to get uname -a
	n, err = sendMsg("shell:uname -a", err, netConn)
	if err != nil {
		fmt.Println("Sending information error %v", err)
	}
	fmt.Println("Number of byte ", n)


	// check OKAY
	// Check whether the response is OKAY
	octet, err = checkOkayReply(err, netConn)
	if err != nil {
		fmt.Println("Checking OKAY error %v", err)
	}
	fmt.Println("Response result : ", string(octet))

	// Read the response sent by adb server based on the length given in
	// the previous step
	dataToRead, err = readRealResponse(2048, err, netConn)
	if err != nil {
		fmt.Println("Reading real response error %v", err)
	}
	fmt.Println("Real response value :", string(dataToRead))

	// Close connection
	netConn.Close()
}

func readRealResponse(length int64, err error, netConn net.Conn) ([]byte, error) {
	dataToRead := make([]byte, length)

	_, err = netConn.Read(dataToRead)

	return dataToRead, err
}

func readResponseLength(err error, netConn net.Conn) (error, int64) {
	version := make([]byte, 4)
	_, err = netConn.Read(version)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
	}

	length, err := strconv.ParseInt(string(version), 16, 64)
	return err, length
}

func checkOkayReply(err error, netConn net.Conn) ([]byte, error) {
	octet := make([]byte, 4)
	_, err = netConn.Read(octet)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
	}
	return octet, err
}

func sendMsg(msg string, err error, netConn net.Conn) (int, error) {
	data := []byte(msg)

	lengthAndMsg := []byte(fmt.Sprintf("%04x%s", len(data), data))

	n, err := netConn.Write(lengthAndMsg[0:])
	return n, err
}
