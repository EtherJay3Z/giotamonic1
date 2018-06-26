package iota_mnemonic

import (
	"testing"

	"github.com/tyler-smith/go-bip39"
)

/*
func WriteDataToFile(name string, data []byte) error {
	if data == nil {
		return errors.New("data is nil")
	}
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	if _, err := file.Write(data); err != nil {
		file.Close()
		os.Remove(name)
		return err
	}
	file.Close()
	return nil
}
*/

func TestMnemonicToIotaSeed(t *testing.T) {
	tests := []struct {
		name       string
		mnemonic   string
		passphrase string
		expected   string
	}{
		{
			name:       "english 24 word mnemonic without passphrase",
			mnemonic:   "come grocery cube calm void liberty increase pigeon captain appear employ among float fancy cargo faith seek buzz argue lift agent split bachelor judge",
			passphrase: "",
			expected:   "PAESQTSDNLRAKLJGPFGOXBLKPMHTGVXLQ9URFJIYMLOQCYNYUVXFSFYABWQHCBFNTAAIJWYLKAWWXNZZY",
		},
		{
			name:       "english 24 word mnemonic with passphrase",
			mnemonic:   "come grocery cube calm void liberty increase pigeon captain appear employ among float fancy cargo faith seek buzz argue lift agent split bachelor judge",
			passphrase: "pA$$w0rD",
			expected:   "GVGYCIACT9MFT9KYBJDGCSRZXNNKERZCPFNGDEEPTNERGFKYDCKGTIHAUUEYEDV9EEXRGDDIULRMDNTGB",
		},
		{
			name:       "english 12 word mnemonic without passphrase",
			mnemonic:   "broccoli merry lucky milk lizard cannon area utility jelly click bag clever",
			passphrase: "",
			expected:   "T9EW9SDNBZVUMBIGLKWSYOPIIOEXETJPLCOVXBTIWXVERZCKWVTKLYTIJZFGTAARJLMRLWRVSKLHRRJZX",
		},
		{
			name:       "english 12 word mnemonic with passphrase",
			mnemonic:   "broccoli merry lucky milk lizard cannon area utility jelly click bag clever",
			passphrase: "pA$$w0rD",
			expected:   "TDPJRZDHUMLMSQAEC9FEJXVMMJMRHCFUWILZSVHG9GRMAPWWYNJHGFUZGOSPKOG9TMWGFLLMASYWHBNBA",
		},
	}

	for _, tt := range tests {

		bseed, err := bip39.NewSeedWithErrorChecking(tt.mnemonic, tt.passphrase)
		if err != nil {
			t.Errorf("Test '%s' (bip39) failed:\n\tmnemonic: '%s'\n\tpassphrase: '%s'\n\terror: %s", tt.name, tt.mnemonic, tt.passphrase, err)
		}
		if len(bseed) != 64 {
			t.Errorf("Test '%s' (bip39) failed:\n\terror: Length of resulting bip39 seed diffs from expected\n\tresult len:   '%d'\n\texpected len: '%d'", tt.name, len(bseed), 64)
		}

		seed, err := ToSeed(tt.mnemonic, tt.passphrase)
		if err != nil {
			t.Errorf("Test '%s' failed:\n\tmnemonic: '%s'\n\tpassphrase: '%s'\n\terror: %s", tt.name, tt.mnemonic, tt.passphrase, err)
		}

		if tt.expected != seed {
			t.Errorf("Test '%s' failed:\n\terror: Result diffs from expected\n\tresult:   '%s'\n\texpected: '%s'", tt.name, seed, tt.expected)
		}
	}
}
