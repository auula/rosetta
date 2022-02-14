// MIT License

// Copyright (c) 2022 Leon Ding

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// encoding.go
// File data encryption and decryption encoder implementation.

package bigmap

import (
	"encoding/binary"
	"errors"
)

var (
	ErrSourceDataEncodeFail = errors.New("big map error: error encrypting source data to write")
)

// Encoder bytes data encoder
type Encoder struct {
	Encryptor      // encryptor concrete implementation
	enable    bool // whether to enable data encryption and decryption
}

// AESEncoder enable the AES encryption encoder
func AESEncoder() *Encoder {
	return &Encoder{
		enable:    true,
		Encryptor: new(AESEncryptor),
	}
}

// DefaultEncoder disable the AES encryption encoder
func DefaultEncoder() *Encoder {
	return &Encoder{
		enable:    false,
		Encryptor: nil,
	}
}

type metaData struct {
}

// ToWrite write to entity's current activation file
func (e *Encoder) ToWrite(entity *Entity) error {

	// whether encryption is enabled
	if e.enable && e.Encryptor != nil {
		// building source data
		sd := &SourceData{
			Secret: secret,
			Data:   entity.Value,
		}
		if err := e.Encode(sd); err != nil {
			return ErrSourceDataEncodeFail
		}
		// valueSize := len(sd.Data)
		// keySize := len(entity.Key)

		//currentFile.WriteAt(sd.Data, lastOffset)
		// bufToWrite(sd.Data)
		//var buf []byte
		binary.LittleEndian.PutUint32(buf[:])
		return nil
	}
	return nil
}

func bufToWrite(entity []byte) {

}
