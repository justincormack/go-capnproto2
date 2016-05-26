// Package canonical provides functions to convert a Cap'n Proto message into its canonical form.
// See https://capnproto.org/encoding.html#canonicalization for more details.
package canonical // import "zombiezen.com/go/capnproto2/canonical"

import (
	"errors"

	"zombiezen.com/go/capnproto2"
)

// Convert canonicalizes the contents of m in a new single-segment message.
func Convert(m *capnp.Message) (*capnp.Message, error) {
	return nil, errors.New("not implemented")
}

// Marshal canonicalizes m and marshals it to the standard framing format.
func Marshal(m *capnp.Message) ([]byte, error) {
	m, err := Convert(m)
	if err != nil {
		return nil, err
	}
	return m.Marshal()
}
