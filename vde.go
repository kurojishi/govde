// Package vde the interface to send package to vde2 transport
package vde

import (
	"net"
)

//ConnectionVde is the object that hold all information about the connection with vde
type ConnectionVde struct {
	name          string
	connection    net.Conn
	configuration map[string]string
}

//Connect initialize a connection
//netType can be  any net.Conn available type
//TODO: add possibility to use an interface as an address
func Connect(netType string, address string, conf map[string]string) (*ConnectionVde, error) {
	var err error
	var connection *ConnectionVde
	connection.name = address
	connection.connection, err = net.Dial(netType, address)
	if err != nil {
		return nil, err
	}
	if conf != nil {
		connection.configuration = conf
	}
	return connection, err
}

//Send packets
func Send(conn *ConnectionVde, payload []byte) (err error) {
	_, err = conn.connection.Write(payload)
	return err
}

//Receive packets
func Receive(conn *ConnectionVde) (payload []byte, err error) {
	_, err = conn.connection.Read(payload)
	return payload, err
}
