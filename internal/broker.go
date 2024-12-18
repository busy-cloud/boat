package internal

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func StartBroker() error {

	server := mqtt.New(nil)

	// Allow all connections.
	err := server.AddHook(new(auth.AllowHook), nil)
	if err != nil {
		return err
	}

	tcp := listeners.NewTCP(listeners.Config{ID: "base", Address: ":1883"})
	err = server.AddListener(tcp)
	if err != nil {
		return err
	}

	sock := listeners.NewUnixSock(listeners.Config{ID: "unix", Address: "boat.sock"})
	err = server.AddListener(sock)
	if err != nil {
		return err
	}

	return server.Serve()
}
