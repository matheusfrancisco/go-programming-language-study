package transport

import (
	"net"

	"github.com/xico-labs/go-programming-language-study/testingclass/interfacesmocks/inventory"
)

type InventoryTransporter interface {
	inventory.Service
	Serve(net.Listener) error
}
