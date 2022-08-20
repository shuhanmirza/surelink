package util

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

type hashSample struct {
	in  []byte
	out string
}

func TestGetHash(t *testing.T) {
	var samples []hashSample
	samples = append(samples, hashSample{[]byte(""), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"})
	samples = append(samples, hashSample{[]byte("abc"), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"})
	samples = append(samples, hashSample{[]byte("hello"), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"})

	for i := 0; i < len(samples); i++ {
		result := GetHash(samples[i].in)
		require.Equal(t, hex.EncodeToString(result), samples[i].out)
	}
}

func TestCheckHash(t *testing.T) {
	var match bool

	match = CheckHash("hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
	require.Equal(t, true, match)

	match = CheckHash("hello ", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
	require.Equal(t, false, match)

	match = CheckHash(" hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
	require.Equal(t, false, match)

	match = CheckHash("hello", "1cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
	require.Equal(t, false, match)
}
