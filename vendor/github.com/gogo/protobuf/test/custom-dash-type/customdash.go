// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

/*
	Package custom contains custom types for test and example purposes.
	These types are used by the test structures generated by gogoprotobuf.
*/
package custom

import (
	"bytes"
	"encoding/json"
)

type Bytes []byte

func (b Bytes) Marshal() ([]byte, error) {
	buffer := make([]byte, len(b))
	_, err := b.MarshalTo(buffer)
	return buffer, err
}

func (b Bytes) MarshalTo(data []byte) (n int, err error) {
	copy(data, b)
	return len(b), nil
}

func (b *Bytes) Unmarshal(data []byte) error {
	if data == nil {
		b = nil
		return nil
	}
	pb := make([]byte, len(data))
	copy(pb, data)
	*b = pb
	return nil
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	data, err := b.Marshal()
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}

func (b *Bytes) Size() int {
	return len(*b)
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	v := new([]byte)
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return b.Unmarshal(*v)
}

func (this Bytes) Equal(that Bytes) bool {
	return bytes.Equal(this, that)
}

func (this Bytes) Compare(that Bytes) int {
	return bytes.Compare(this, that)
}

type randy interface {
	Intn(n int) int
}

func NewPopulatedBytes(r randy) *Bytes {
	l := r.Intn(100)
	data := Bytes(make([]byte, l))
	for i := 0; i < l; i++ {
		data[i] = byte(r.Intn(255))
	}
	return &data
}
