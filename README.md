# Diablo II character file parser
[![Go Report Card](https://goreportcard.com/badge/github.com/nokka/d2s)](https://goreportcard.com/report/github.com/nokka/d2s)

D2s is a binary parser written in Go that's used to parse `.d2s` files. This is the binary format that the game [Diablo II](http://eu.blizzard.com/en-gb/games/d2/) uses to save all information about a certain character. 

## Motivation
This package was built for a private server of Diablo II called [Slashgaming](https://reddit.com/r/slashgaming) to build an Armory for all characters on the server. Where anyone could see everything about a particular character at any given point in time. Here's my sorceress for example [Nokkasorc](https://armory.slashgaming.net/character/nokkasorc#equipped) .

## Examples
The examples folder contains a [JSON](examples/nokkasorc.json) example of an unmarshalled character struct, the amount of data that is received from a binary is quite immense.

## Install

```bash
$ go get github.com/nokka/d2s
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nokka/d2s"
)

func main() {
	path := "nokka.d2s"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening .d2s file", err)
	}

	defer file.Close()

	char, err := d2s.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// Prints character name and class.
	fmt.Println(char.Header.Name)
	fmt.Println(char.Header.Class)
}

```


## Sections
The sections is a rough explanation of the different sections that exists, but each section is pretty complex to read and was even more complex to reverse engineer.

### 1. Header
Each file starts with a `767 bytes` long header, this header contains all the basic information about your character like Level, Class, Game version, Name and so on.

### 2. Attributes
Following the header is the attributes section, this sections layout consists of an array of  `9 bit` attribute id, followed by a `n bit` length attribute value. The bit length of all attributes are documented in the [attributes.go](attributes.go) file.

The section is terminated by a `9 bit` value of `0x1ff`.

### 3. Skills
Skills contains a section of `32 bytes` and starts of with a `2 byte` header value of `if`.  

Each class has 30 skills, and the map of skills are offseted by class. For example the Amazon class starts at index `6`, while the Sorceress starts at `36` and so on. After finding the correct index to start from, the skills section will read 30 skills into an array. Each skill is `1 byte` long.

### 4. Items
This is by far the most tricky part to read. The items section starts of with a `4 byte` header, containing the value `JM`, and a `uint16` value which is the item count your character currently has. Equipped, inventory, stash, cube and belt all included.

The byte length of the section is unknown before reading it in it's entirety, because the bit length of each item varies depending on its quality, number of sockets and magical attributes it possess.

Each item follows a certain pattern though, which is:

#### Simple items
Each item starts of with `111` bits of simple data, which all items contain. This is information like item type, if it's socketed, position id like equipped or stash and so on.

Each item also has a boolean called `SimpleItem` which is `1` bit long,  if this is set to `1`, the item contains no more bits, and the next item starts.

#### Advanced items
If the item is not a simple items, this means it will have tons of data following the initial `111` bits. A few examples of this is the rarity level, magical suffix, magical affix, if it's a runeword, personalized, part of a set, class specific and so on. 

Last but not least if the item has will have lists of magical properties depending on if it's a runeword, magical, rare, crafted, unique part of set and so on.

These lists are similar to the attributes section where we will read:

1. `9 bit id`
2. `n bits of magical properties`
3. `0x1ff terminator`

When we hit the terminator `0x1ff` the next item begins.

#### Magical properties
A magical property is a unique property that can occur on an item, each property has a different bit length, and the map is huge.

##### Example

This is the magical property with id `83` which contains 2 bit fields each `3` bits long.
```go
83: {Bits: []uint{3, 3}, Name: "+{1} to {0} Skill Levels"},
```

All magical properties are mapped in the [item.go](item.go) file.

### 5. Corpse data
If your character is currently dead, and a corpse is on the ground when you enter a game, your equipped items will be in this item struct. It's a corpse header `16 bytes` containing the header `JM` followed by an item count similar to the item list.

Reading the corpse items are done in the exact way as the previous section of items.

### 6. Expansion data
If your character is created in the expansion Lord of Destruction if will contain 2 more sections.

#### 7. Mercenary items
The mercenary sections starts of with a `2 byte` header with the value `jf` and is followed by a `4 byte` item header containing the number of items the mercenary is currently wearing. The items are read like any other item list.

#### 8. Golem
If your character is both a Necromancer and an expansion character, this section starts of with a `3 byte` header, where the first two bytes are the header `kf` followed by a boolean called `hasGolem`, if this value is true, there's an item list with the length 1 following the header.
