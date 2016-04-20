# GOCARINA - simple Optical Character Recognition in Go

Gocarina uses a neural network to do simple Optical Character Recognition (OCR).
It's trained on [Letterpress®](http://www.atebits.com/letterpress) game boards.

![logo](https://github.com/armhold/gocarina/blob/master/gocarina-logo.png "gocarina Logo")


## Usage

First, install the software:

`$ go get github.com/armhold/gocarina/...`

Next, we need to create and train a network. Be sure to first connect to the source directory
(`train` expects the game boards to appear in `board-images/`):

```
$ cd $GOPATH/src/github.com/armhold/gocarina
$ train
creating new network...
success took 63 iterations
success rate: 26/26 => %100.00
```

You now have a trained neural network in `ocr.save`. If you got a failure message, simply try running it again;
sometimes it takes a few attempts to get a successful training.

Once you have a success network, you can ask it to decipher game boards like this:

`$ recognize board-images/board3.png`
```
 L H F L M
 R V P U K
 V O E E X
 I N R I T
 V N S I Q
```

You can also ask it to give you a list of words that can be formed with the board:

`$ recognize -w board-images/board3.png`
```
 L H F L M
 R V P U K
 V O E E X
 I N R I T
 V N S I Q


overmultiplies
relinquishment
feuilletonism
fluorimetries
interinvolves
pluviometries
reptiliferous
[etc...]
```


## How it works

We start with three "known" game boards. We split them up into individual tiles, one per letter.
This covers the entire alphabet, and gives us our training set. We feed the training tiles into the network
one at a time, and calculate the error value for expected vs. the actual result. We do this repeatedly,
until the network is trained (typically requires < 100 iterations).


## Representation & Encoding for the Neural Network

The tiles are quantized to black & white, bounding boxed, and finally scaled down to a small rectangular bitmap.
These bits are then fed directly into the inputs of the network.

We use a bit string to represent a given letter. 8 bits allows us to represent up to 256 different characters,
which is more than sufficient to cover the 26 characters used in Letterpress (we could certainly get away
with using only 5 bits, but I wanted to hold the door open for potentially doing more than just A-Z).

For convenience, we use the natural ASCII/Unicode mapping where 'A' = 65, aka 01000001. So our network has 8
outputs, corresponding to the 8 bits of our letters.


## Can I use this as a production-ready OCR package?

Doubtful. This is more or less a toy implementation of OCR that operates on a very restricted set of input.
It was created by an AI-hobbyist (not an expert), for fun and for educational purposes. However there's nothing
stopping you from building something more robust, based on what you've learned here.

A further caveat: this software expects game boards to be 640x1136 pixels, as that is the size generated by
my iPhone5. Your mobile device likely uses a different board size, based on its screen. Gocarina automatically
scales the boards to the expected size, but I haven't tested it exhaustively with every mobile device; you might
have more success in adjusting the geometry values such as `LetterPressExpectedWidth`, than with scaling alone.


## What's with the name?

This is a Golang port of the [Ruby project](https://github.com/armhold/ocarina) I did a few years back.
Original project: "Ocarina", **OC**a**R**ina, i.e. OCR. Go + Ocarina => Gocarina.


## Credits

The file `words-en.txt` is in the Public Domain, licensed under CC0 thanks to https://github.com/atebits/Words.

Letterpress® is a registered mark of Atebits/Solebon. The Gocarina open-source software is in no way
affiliated with, nor is it endorsed by, the trademark holder.
