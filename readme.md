# GIOTAmonic

This is a command line tool to help user generate IOTA seed from Bitcoin BIP39 mnenomic

## Build

GIOTAmonic has a dependency to `glide` which manages related vendor packages - so it must be installed.

To install the vendor packages run:

```
> glide install
```

then you can build it with

```
> go build
```

or test the sub package `iota_mnemonic` with following command:

```
> go test ./iota_mnemonic
```



## Usage

With `giotamonic help` you will get a help how to use giotamonic:

```
> giotamonic help

NAME:
   giotamonic - Generate, restore IOTA seed from Bitcoin BIP39 mnemonic word list.

USAGE:
   giotamonic [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     new      Create new mnemonic and iota seed
     to-seed  Convert mnemonic to iota seed
```

### Generate new seed

To output a new seed and the corresponding BIP39 mnemonic word list, execute giotamonic mit the `new` command:

```
> giotamonic help new

NAME:
   giotamonic new - Create new mnemonic and iota seed

USAGE:
   giotamonic new [command options] [arguments...]

DESCRIPTION:
   Creates new BIP0039 mnemonic word list and corresponding - optional passphrase encrypted - iota seed.

OPTIONS:
   --passphrase value  Additional passphrase
   --bitsize value     Bit length of the mnemonic entropy. Must be [128, 256] and a multiple of 32 (default: 256)
```

#### Examples

* 24 word mnemonic + seed:
```
> giotamonic new

Mnemonic: abuse episode network recall cement connect left sport nose claw reveal certain struggle north strike surprise tennis luxury begin pole trap quote labor collect
Seed: HYRQZWPCRFIQ9BZOJBJGMNBRKAHWRSZKYZWS9NTTGPCWLRFUFUGHORONUHWXGNUQZUQGNWPPKUUEEXJQY
```

* 12 word mnemonic + seed:
```
> giotamonic new --bitsize 128

Mnemonic: ugly scorpion hour trial blue forum glass life click feature mean sentence
Seed: IUSOCINQJUYBBGWDKBIKWU9YYCFHJFRALIPPHGIQHMYRWUNVZLWFEAHZDKFRFGZAAEWMKEI9YTRRGCYZA
```

* Word mnemonic + seed (passphrase encrypted):
```
> giotamonic new --passphrase "Mei Pa$$frA$e"

Mnemonic: rifle rhythm zebra practice pet fish general accuse virtual traffic history blanket visit gaze leave city alpha injury myself pizza upgrade trade detect awake
Seed: KIVQGMYMDXXOXMDKIKOXHDYMWFTRHXHWZPEPHHIWONSZLLCQMVTVHLICMUUEIERAQFZB9ZPKKGHFBWCP9
```

### Generate seed from existing mnemonic

```
> giotamonic help to-seed

NAME:
   giotamonic to-seed - Convert mnemonic to iota seed

USAGE:
   giotamonic to-seed [command options] [arguments...]

DESCRIPTION:
   Converts mnemonic words to iota seed.

OPTIONS:
   --passphrase value  Additional passphrase
```

#### Examples

```
> giotamonic to-seed "ugly scorpion hour trial blue forum glass life click feature mean sentence"

IUSOCINQJUYBBGWDKBIKWU9YYCFHJFRALIPPHGIQHMYRWUNVZLWFEAHZDKFRFGZAAEWMKEI9YTRRGCYZA
```

* with passphrase:
```
> giotamonic to-seed "ugly scorpion hour trial blue forum glass life click feature mean sentence" --passphrase "Mei Pa$$frA$e"

XLQHKPKWBASMVJVWRDPDIKFOGTEFQGSCUXCKHSQXTTSEAXIL9JJGEWQHHGXFWKPBUUBNJTQEGEKPDSOKX
```

* You can also use pipes:

```
> giotamonic to-seed --passphrase "Mei Pa$$frA$e" < mnemonic.txt > seed.txt
```

### How it works

The mnemonic and passphrase will be used to generate a 64 byte seed according to the (BIP-0039)[https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki] standard. IOTA seeds consist of 81 trytes. Deterministic conversions can done between 81 trytes and 48 bytes which IOTAs `Kerl` does. For the extraction of 4 48 byte slices from a 64 byte block I implemented the simple algorithm from (Bart Slinger)[https://github.com/iota-trezor/trezor-mcu/blob/25292640b560a644ebf88d0dae848e8928e68127/firmware/iota.c#L70].
