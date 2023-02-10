package wire

import "github.com/shravan9912/mpquic_ml_va/internal/protocol"

// AckRange is an ACK range
type AckRange struct {
	First protocol.PacketNumber
	Last  protocol.PacketNumber
}
