package portpicker

import (
	"net"
	"strconv"
)

func PickUnusedPort() (int, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer l.Close()
	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(port)
}
