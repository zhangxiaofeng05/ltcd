// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chainhash

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"
)

// mainNetGenesisHash is the hash of the first block in the block chain for the
// main network (genesis block).
var mainNetGenesisHash = Hash([HashSize]byte{ // Make go vet happy.
	0xe2, 0xbf, 0x04, 0x7e, 0x7e, 0x5a, 0x19, 0x1a,
	0xa4, 0xef, 0x34, 0xd3, 0x14, 0x97, 0x9d, 0xc9,
	0x98, 0x6e, 0x0f, 0x19, 0x25, 0x1e, 0xda, 0xba,
	0x59, 0x40, 0xfd, 0x1f, 0xe3, 0x65, 0xa7, 0x12,
})

// TestHash tests the Hash API.
func TestHash(t *testing.T) {
	// Hash of block 234439.
	blockHashStr := "8cbcc392e1b003fed4240f3b5add3d0843584a7c9d3693c14068ff0b3e10630b"
	blockHash, err := NewHashFromStr(blockHashStr)
	if err != nil {
		t.Errorf("NewHashFromStr: %v", err)
	}

	// Hash of block 234440 as byte slice.
	buf := []byte{
		0xd3, 0xae, 0x72, 0xde, 0xf7, 0xbe, 0xc3, 0x93,
		0xf5, 0x4f, 0x38, 0xa7, 0x02, 0xbe, 0x0c, 0x8b,
		0x2d, 0x7f, 0xee, 0xcf, 0x07, 0x31, 0x56, 0x06,
		0x49, 0x42, 0x61, 0x9e, 0x08, 0x32, 0x3a, 0x94,
	}

	hash, err := NewHash(buf)
	if err != nil {
		t.Errorf("NewHash: unexpected error %v", err)
	}

	// Ensure proper size.
	if len(hash) != HashSize {
		t.Errorf("NewHash: hash length mismatch - got: %v, want: %v",
			len(hash), HashSize)
	}

	// Ensure contents match.
	if !bytes.Equal(hash[:], buf) {
		t.Errorf("NewHash: hash contents mismatch - got: %v, want: %v",
			hash[:], buf)
	}

	// Ensure contents of hash of block 234440 don't match 234439.
	if hash.IsEqual(blockHash) {
		t.Errorf("IsEqual: hash contents should not match - got: %v, want: %v",
			hash, blockHash)
	}

	// Set hash from byte slice and ensure contents match.
	err = hash.SetBytes(blockHash.CloneBytes())
	if err != nil {
		t.Errorf("SetBytes: %v", err)
	}
	if !hash.IsEqual(blockHash) {
		t.Errorf("IsEqual: hash contents mismatch - got: %v, want: %v",
			hash, blockHash)
	}

	// Ensure nil hashes are handled properly.
	if !(*Hash)(nil).IsEqual(nil) {
		t.Error("IsEqual: nil hashes should match")
	}
	if hash.IsEqual(nil) {
		t.Error("IsEqual: non-nil hash matches nil hash")
	}

	// Invalid size for SetBytes.
	err = hash.SetBytes([]byte{0x00})
	if err == nil {
		t.Errorf("SetBytes: failed to received expected err - got: nil")
	}

	// Invalid size for NewHash.
	invalidHash := make([]byte, HashSize+1)
	_, err = NewHash(invalidHash)
	if err == nil {
		t.Errorf("NewHash: failed to received expected err - got: nil")
	}
}

// TestHashString  tests the stringized output for hashes.
func TestHashString(t *testing.T) {
	// Block 100000 hash.
	wantStr := "e11049dfa5858be3809f285685e12a5d6f84b936b0f8e8272b5363bf3946ce60"
	hash := Hash([HashSize]byte{ // Make go vet happy.
		0x60, 0xce, 0x46, 0x39, 0xbf, 0x63, 0x53, 0x2b,
		0x27, 0xe8, 0xf8, 0xb0, 0x36, 0xb9, 0x84, 0x6f,
		0x5d, 0x2a, 0xe1, 0x85, 0x56, 0x28, 0x9f, 0x80,
		0xe3, 0x8b, 0x85, 0xa5, 0xdf, 0x49, 0x10, 0xe1,
	})

	hashStr := hash.String()
	if hashStr != wantStr {
		t.Errorf("String: wrong hash string - got %v, want %v",
			hashStr, wantStr)
	}
}

// TestNewHashFromStr executes tests against the NewHashFromStr function.
func TestNewHashFromStr(t *testing.T) {
	tests := []struct {
		in   string
		want Hash
		err  error
	}{
		// Genesis hash.
		{
			"12a765e31ffd4059bada1e25190f6e98c99d9714d334efa41a195a7e7e04bfe2",
			mainNetGenesisHash,
			nil,
		},

		// Empty string.
		{
			"",
			Hash{},
			nil,
		},

		// Single digit hash.
		{
			"1",
			Hash([HashSize]byte{ // Make go vet happy.
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			}),
			nil,
		},

		// Block 203707 with stripped leading zeros.
		{
			"5f926a8bb8a71bd72a8fa413591c5524695b8fdde5fd139582e633ac729ad746",
			Hash([HashSize]byte{ // Make go vet happy.
				0x46, 0xd7, 0x9a, 0x72, 0xac, 0x33, 0xe6, 0x82,
				0x95, 0x13, 0xfd, 0xe5, 0xdd, 0x8f, 0x5b, 0x69,
				0x24, 0x55, 0x1c, 0x59, 0x13, 0xa4, 0x8f, 0x2a,
				0xd7, 0x1b, 0xa7, 0xb8, 0x8b, 0x6a, 0x92, 0x5f,
			}),
			nil,
		},

		// Hash string that is too long.
		{
			"01234567890123456789012345678901234567890123456789012345678912345",
			Hash{},
			ErrHashStrSize,
		},

		// Hash string that is contains non-hex chars.
		{
			"abcdefg",
			Hash{},
			hex.InvalidByteError('g'),
		},
	}

	unexpectedErrStr := "NewHashFromStr #%d failed to detect expected error - got: %v want: %v"
	unexpectedResultStr := "NewHashFromStr #%d got: %v want: %v"
	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result, err := NewHashFromStr(test.in)
		if err != test.err {
			t.Errorf(unexpectedErrStr, i, err, test.err)
			continue
		} else if err != nil {
			// Got expected error. Move on to the next test.
			continue
		}
		if !test.want.IsEqual(result) {
			t.Errorf(unexpectedResultStr, i, result, &test.want)
			continue
		}
	}
}

// TestHashJsonMarshal tests json marshal and unmarshal.
func TestHashJsonMarshal(t *testing.T) {
	hashStr := "000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506"
	legacyHashStr := []byte("[6,229,51,253,26,218,134,57,31,63,108,52,50,4,176,210,120,212,170,236,28,11,32,170,39,186,3,0,0,0,0,0]")

	hash, err := NewHashFromStr(hashStr)
	if err != nil {
		t.Errorf("NewHashFromStr error:%v, hashStr:%s", err, hashStr)
	}

	hashBytes, err := json.Marshal(hash)
	if err != nil {
		t.Errorf("Marshal json error:%v, hash:%v", err, hashBytes)
	}

	var newHash Hash
	err = json.Unmarshal(hashBytes, &newHash)
	if err != nil {
		t.Errorf("Unmarshal json error:%v, hash:%v", err, hashBytes)
	}

	if !hash.IsEqual(&newHash) {
		t.Errorf("String: wrong hash string - got %v, want %v", newHash.String(), hashStr)
	}

	err = newHash.UnmarshalJSON(legacyHashStr)
	if err != nil {
		t.Errorf("Unmarshal legacy json error:%v, hash:%v", err, legacyHashStr)
	}

	if !hash.IsEqual(&newHash) {
		t.Errorf("String: wrong hash string - got %v, want %v", newHash.String(), hashStr)
	}
}
