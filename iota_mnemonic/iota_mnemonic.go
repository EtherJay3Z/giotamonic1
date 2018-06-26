package iota_mnemonic

import (
	"fmt"

	"github.com/iotaledger/giota"
	"github.com/tyler-smith/go-bip39"
)

// Generates IOTA seed trits from a 64 bytes seed
// The algorithm was adapted from
// https://github.com/iota-trezor/trezor-mcu/blob/25292640b560a644ebf88d0dae848e8928e68127/firmware/iota.c#L70
// Absorb 4 times using sliding window:
// Divide 64 byte bip39 seed in 4 sections of 48 bytes.
// 1: [123.] first 48 bytes
// 2: [.123] last 48 bytes
// 3: [3.12] last 32 bytes + first 16 bytes
// 4: [23.1] last 16 bytes + first 32 bytes
func ByteSeedToTrits(seed []byte) (giota.Trits, error) {
	seedLength := len(seed)
	if seedLength != 64 {
		return nil, fmt.Errorf("Seed must have a length of 64 bytes!")
	}

	bytes := make([]byte, giota.ByteLength)

	sponge := giota.NewKerl()
	if sponge == nil {
		return nil, fmt.Errorf("Could not initialize Kerl instance.")
	}

	// Step 1
	for j := 0; j < giota.ByteLength; j++ { // 48
		bytes[j] = seed[j]
	}
	ti, err := giota.BytesToTrits(bytes)
	if err != nil {
		return nil, err
	}
	err = sponge.Absorb(ti)
	if err != nil {
		return nil, err
	}

	// Step 2
	offset := seedLength - giota.ByteLength //64 - 48 = 16
	for j := 0; j < giota.ByteLength; j++ {
		bytes[j] = seed[j+offset]
	}
	ti, err = giota.BytesToTrits(bytes)
	if err != nil {
		return nil, err
	}
	err = sponge.Absorb(ti)
	if err != nil {
		return nil, err
	}

	// Step 3
	offset = seedLength / 2 // 64 / 2 = 32
	for j := 0; j < offset; j++ {
		bytes[j] = seed[j+offset]
	}
	for j := offset; j < giota.ByteLength; j++ {
		bytes[j] = seed[j-offset]
	}
	ti, err = giota.BytesToTrits(bytes)
	if err != nil {
		return nil, err
	}
	err = sponge.Absorb(ti)
	if err != nil {
		return nil, err
	}

	// Step 4
	offset = seedLength - giota.ByteLength //64 - 48 = 16
	for j := 0; j < offset; j++ {
		bytes[j] = seed[j+seedLength-offset]
	}
	for j := offset; j < giota.ByteLength; j++ {
		bytes[j] = seed[j-offset]
	}
	ti, err = giota.BytesToTrits(bytes)
	if err != nil {
		return nil, err
	}
	err = sponge.Absorb(ti)
	if err != nil {
		return nil, err
	}

	// Squeeze out the seed
	s_trits, err := sponge.Squeeze(giota.TritHashLength)
	if err != nil {
		return nil, err
	}
	return s_trits, nil
}

// Generates IOTA seed trits from a BIP0039 space delimited word list
func ToTrits(mnemonic string, passphrase string) (giota.Trits, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, passphrase)
	if err != nil {
		return nil, err
	}
	return ByteSeedToTrits(seed)
}

// Generates IOTA seed trytes from a BIP0039 space delimited word list
func ToTrytes(mnemonic string, passphrase string) (giota.Trytes, error) {
	t, err := ToTrits(mnemonic, passphrase)
	if err != nil {
		return giota.Trytes(""), err
	}
	return t.Trytes(), err
}

// Generates IOTA seed string from a BIP0039 space delimited word list
func ToSeed(mnemonic string, passphrase string) (string, error) {
	t, err := ToTrytes(mnemonic, passphrase)
	return string(t), err
}
