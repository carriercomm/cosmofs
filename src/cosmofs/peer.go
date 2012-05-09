/**

Copyright (C) 2012  Roberto Costumero Moreno <roberto@costumero.es>

This file is part of Cosmofs.

Cosmofs is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Cosmofs is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Cosmofs.  If not, see <http://www.gnu.org/licenses/>.

**/

package cosmofs

import (
	"bytes"
	"crypto/rsa"
	"encoding/binary"
	"encoding/base64"
	"log"
	"math/big"
)

const (
	hostAlgoRSA = "ssh-rsa"
)

var (
	PeerList map[string]*Peer = make(map[string]*Peer)
)

type localPeer struct {
	ID string
	Key *rsa.PrivateKey
}

type Peer struct {
	ID string
	PubKey *rsa.PubKey
}

// ParsePubKey parses a Public SSH-RSA Key encoded in Base64 format
func ParsePubKey(in []byte) (out interface{}, rest, id []byte, ok bool) {
	algo, key, id, ok := parseString(in)

	if !ok {
		return
	}

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(key)))

	_, err := base64.StdEncoding.Decode(dst, key)

	if err != nil {
		log.Println("Error decoding rsa key:", err)
		return
	}

	algo, key, ok = parseKey(dst)

	if !ok {
		return
	}


	switch string(algo) {
	case hostAlgoRSA:
		pubkey, rest, ok := parseRSA(key)
		return pubkey, rest, id, ok
	}
	panic("ssh: unknown public key type")
}

// parseRSA parses an RSA key according to RFC 4253, section 6.6.
func parseRSA(in []byte) (out *rsa.PublicKey, rest []byte, ok bool) {
	key := new(rsa.PublicKey)

	bigE, in, ok := parseInt(in)
	if !ok || bigE.BitLen() > 24 {
		return
	}
	e := bigE.Int64()
	if e < 3 || e&1 == 0 {
		ok = false
		return
	}
	key.E = int(e)

	if key.N, in, ok = parseInt(in); !ok {
		return
	}

	ok = true
	return key, in, ok
}

func parseString(in []byte) (kind, key, id []byte, ok bool) {
	parts := bytes.Split(in, []byte(" "))

	kind = parts[0]
	key = parts[1]
	id = parts[2]
	ok = true
	return
}

func parseKey(in []byte) (out, rest []byte, ok bool) {
	if len(in) < 4 {
		return
	}

	length := binary.BigEndian.Uint32(in)
	if uint32(len(in)) < 4+length {
		return
	}
	out = in[4 : 4+length]
	rest = in[4+length:]
	ok = true
	return
}

func parseInt(in []byte) (out *big.Int, rest []byte, ok bool) {
	contents, rest, ok := parseKey(in)
	if !ok {
		return
	}
	out = new(big.Int)

	if len(contents) > 0 && contents[0]&0x80 == 0x80 {
		// This is a negative number
		notBytes := make([]byte, len(contents))
		for i := range notBytes {
			notBytes[i] = ^contents[i]
		}
		out.SetBytes(notBytes)
		out.Add(out, big.NewInt(1))
		out.Neg(out)
	} else {
		// Positive number
		out.SetBytes(contents)
	}
	ok = true
	return
}
