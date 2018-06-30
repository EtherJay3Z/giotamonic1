[![Build Status](https://travis-ci.org/mcpride/giotamonic.svg?branch=master)](https://travis-ci.org/mcpride/giotamonic) [![license](https://img.shields.io/github/license/mashape/apistatus.svg?style=flat)](https://raw.githubusercontent.com/mcpride/giotamonic/master/LICENSE)


# GIOTAmonic

This is a command line tool to help user generate IOTA seed from Bitcoin BIP39 mnenomic

## Install

``` shell
$ go install github.com/mcpride/giotamonic
```

## Build

GIOTAmonic has a dependency to the tool `glide` which manages related vendor packages - so it must be installed.

To install the vendor packages run:

``` shell
$ glide install
```

then you can build it with

``` shell
$ go build
```

or test the sub package `iota_mnemonic` with following command:

``` shell
$ go test ./iota_mnemonic
```



## Usage

With `giotamonic help` you will get a help how to use giotamonic:

``` shell
$ giotamonic help

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

``` shell
$ giotamonic help new

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
``` shell
$ giotamonic new

Mnemonic: abuse episode network recall cement connect left sport nose claw reveal certain struggle north strike surprise tennis luxury begin pole trap quote labor collect
Seed: HYRQZWPCRFIQ9BZOJBJGMNBRKAHWRSZKYZWS9NTTGPCWLRFUFUGHORONUHWXGNUQZUQGNWPPKUUEEXJQY
```

* 12 word mnemonic + seed:
``` shell
$ giotamonic new --bitsize 128

Mnemonic: ugly scorpion hour trial blue forum glass life click feature mean sentence
Seed: IUSOCINQJUYBBGWDKBIKWU9YYCFHJFRALIPPHGIQHMYRWUNVZLWFEAHZDKFRFGZAAEWMKEI9YTRRGCYZA
```

* Word mnemonic + seed (passphrase encrypted):
``` shell
$ giotamonic new --passphrase "Mei Pa$$frA$e"

Mnemonic: rifle rhythm zebra practice pet fish general accuse virtual traffic history blanket visit gaze leave city alpha injury myself pizza upgrade trade detect awake
Seed: KIVQGMYMDXXOXMDKIKOXHDYMWFTRHXHWZPEPHHIWONSZLLCQMVTVHLICMUUEIERAQFZB9ZPKKGHFBWCP9
```

### Generate seed from existing mnemonic

``` shell
$ giotamonic help to-seed

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

``` shell
$ giotamonic to-seed "ugly scorpion hour trial blue forum glass life click feature mean sentence"

IUSOCINQJUYBBGWDKBIKWU9YYCFHJFRALIPPHGIQHMYRWUNVZLWFEAHZDKFRFGZAAEWMKEI9YTRRGCYZA
```

* with passphrase:
``` shell
$ giotamonic to-seed "ugly scorpion hour trial blue forum glass life click feature mean sentence" --passphrase "Mei Pa$$frA$e"

XLQHKPKWBASMVJVWRDPDIKFOGTEFQGSCUXCKHSQXTTSEAXIL9JJGEWQHHGXFWKPBUUBNJTQEGEKPDSOKX
```

* You can also use pipes:
``` shell
$ giotamonic to-seed --passphrase "Mei Pa$$frA$e" < mnemonic.txt > seed.txt
```

### How it works

The mnemonic and passphrase will be used to generate a 64 byte seed according to the [BIP-0039](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) standard. IOTA seeds consist of 81 trytes. Deterministic conversions can done between 81 trytes and 48 bytes. For the extraction of 4 x 48 byte slices from a 64 byte block I implemented the [simple algorithm from Bart Slinger](https://github.com/iota-trezor/trezor-mcu/blob/25292640b560a644ebf88d0dae848e8928e68127/firmware/iota.c#L70). These slices will be absorbed by IOTA Kerl and then squeezed out to an IOTA hash seed.

#### More detailed

The following steps sketches the algrorithm in prosa:

* Generate the 64 bytes seed from mnemonic words and password according to BIP-39
* Slide the 64 bytes of mnemonic seed into 4 x 16 bytes blocks `[1|2|3|4]`.
* Get the first 48 bytes (first 3 blocks): `[1|2|3]` and absorb them with IOTA's `Kerl`.
* Get the last 48 bytes (last 3 blocks): `[2|3|4]` and absorb them with IOTA's `Kerl`.
* Get the last 32 bytes (last 2 blocks) + first 16 bytes (1st block): `[3|4|1]` and absorb them with IOTA's `Kerl`.
* Get tshe last 16 bytes (last block) + first 32 bytes (first 2 blocks): `[4|1|2]` and absorb them with IOTA's `Kerl`.
* Call `Squeeze` from `Kerl` to get the IOTA seed.

### License

Giotamonic and it's sub packages are under the MIT license. See the [LICENSE](LICENSE) file for details.
