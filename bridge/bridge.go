// 21 october 2015
package bridge

import (
	"fmt"
	"crypto/cipher"
)

var ErrWrongKEK = fmt.Errorf("wrong KEK")

type Bridge interface {
	Name() string
	Is(keySector []byte) bool
	NeedsKEK() bool
	// must return ks == nil if err != nil
	// do not check if KEK is wrong; that will be done in KeySector.DEK()
	// this way, we can still use KeySector.Raw() for research and debugging
	DecryptKeySector(keySector []byte, kek []byte) (ks KeySector, err error)
	Decrypt(c cipher.Block, b []byte)
}

type KeySector interface {
	Raw() []byte
	// must return dek == nil (not an empty slice!) if err != nil
	DEK() (dek []byte, err error)
}

var Bridges []Bridge

func IdentifyKeySector(possibleKeySector []byte) Bridge {
	for _, b := range Bridges {
		if b.Is(possibleKeySector) {
			return b
		}
	}
	return nil // not a (known) key sector
}

type IncompleteImplementationError string

func IncompleteImplementation(format string, args ...interface{}) IncompleteImplementationError {
	return IncompleteImplementationError(fmt.Sprintf(format, args...))
}

func (i IncompleteImplementationError) Error() string {
	return string(i)
}
