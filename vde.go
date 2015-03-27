// Package vde the interface to send package to vde2 transport
package govde

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

//Initialize a Listener to receive packets
//TODO: make this a ConnectionVde like object
func Listener(netType string, address string, conf map[string]string) (listen net.Listener, err error) {
	listen, err = net.Listen(netType, address)
	return
}

//Send packets
func (conn *ConnectionVde) Send(payload []byte) (err error) {
	_, err = conn.connection.Write(payload)
	return err
}

//Receive packets
func (conn *ConnectionVde) Receive() (payload []byte, err error) {
	_, err = conn.connection.Read(payload)
	return payload, err
}
