# Diablo II character file parser
[![Go Report Card](https://goreportcard.com/badge/github.com/nokka/d2s)](https://goreportcard.com/report/github.com/nokka/d2s)
[![GoDoc](https://godoc.org/github.com/nokka/d2s?status.svg)](https://godoc.org/github.com/nokka/d2s)

D2s is a binary parser written in Go that's used to parse `.d2s` files. This is the binary format that the game [Diablo II](http://eu.blizzard.com/en-gb/games/d2/) uses to save all information about a certain character. 

## Motivation
This package was built for a private server of Diablo II called [Slashgaming](https://reddit.com/r/slashgaming) to build an Armory for all characters on the server. Where anyone could see everything about a particular character at any given point in time. Here's my sorceress for example [Nokkasorc](https://armory.slashgaming.net/character/nokkasorc#equipped) .

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

## Binary layout

### Header

Offset | Bytes | Description
-----|-------- |------------
0    | 4       | Identifier
4    | 4       | Version ID
8    | 4       | File size
12   | 4       | Checksum
16   | 4       | Active weapon
20   | 16      | [Character Name](#character-name)
36   | 1       | [Character Status](#character-status)
37   | 1       | [Character progression](#character-progression)
38   | 2       | Unknown
40   | 1       | [Character Class](#character-class)
41   | 2       | Unknown
43   | 1       | Character Level
44   | 4       | Unknown
48   | 4       | [Last played](#last-played)
52   | 4       | Unknown
56   | 64      | [Assigned skills](#assigned-skills)
120   | 4      | Left mouse button skill ID
124   | 4      | Right mouse button skill ID
128   | 4      | Left swap mouse button skill ID
132   | 4      | Right swap mouse button skill ID
136   | 32     | Character menu appearance
168   | 3      | Difficulty
171   | 4      | Map ID
175   | 2      | Unknown
177   | 2      | Mercenary dead
179   | 4      | Mercenary ID
183   | 2      | Mercenary Name ID
185   | 2      | Mercenary type
187   | 4      | Mercenary experience
191   | 144      | Unknown
335   | 298      | Quests
633   | 81      | Waypoints
714   | 51      | NPC Introductions
765   |  -     | Attributes

#### Character name

Character names are storted as a `[16]byte` which will contain the name, one letter per `byte`. The name can be 16 characters long, and a name that's shorter will have padded `0x00`'s behind the name until we reach `16 bytes`.

 * Must be 2-15 in length
 * Must begin with a letter
 * May contain up to one `-` or `_`.

#### Character status
Character status is a `byte` where different bits will be set, depending on the status of the character. Still haven't figured them all out, but here's the most important ones.

##### Bit position

| 7 | 6 | 5         | 4 | 3    | 2        | 1 | 0 |
|---|---|-----------|---|------|----------|---|---|
| ? | ? | Expansion | ? | Died | Hardcore | ? | ? |

#### Character progression 
>Not implemented yet.

The value is incremented every time you kill an act boss.

##### Classic
| Value | Standard       | Hardcore       |
|-------|----------------|----------------|
| 0-3   | -              | -              |
| 4-7   | Sir/Dame       | Count/Countess |
| 8-11  | Lord/Lady      | Duke/Duchess   |
| 12    | Baron/Baroness | King/Queen     |

##### Expansion
| Value | Expansion           | Expansion hardcore |
|-------|---------------------|--------------------|
| 0-3   | -                   | -                  |
| 5-8   | Slayer              | Destroyer          |
| 10-13 | Champion            | Conqueror          |
| 15    | Patriarch/Matriarch | Guardian           |


#### Character class
Character class is a `byte` where different values represent a class.

|    Class    |  Value |
|:-----------:|:------:|
| Amazon      | `0x00` |
| Sorceress   | `0x01` |
| Necromancer | `0x02` |
| Paladin     | `0x03` |
| Barbarian   | `0x04` |
| Druid       | `0x05` |
| Assassin    | `0x06` |

#### Last played
Last played is saved as a `unit32` [unix timestamp](https://en.wikipedia.org/wiki/Unix_time) e.g `1495882861`.

#### Assigned skills
Assigned skills are a `32 byte` section containing a `2 byte` header with the value `if` and `30 byte` of skill data. Each class has 30 skills available to them, so each skill get `1 byte` each. The tricky part about the skill mapping is that each class has a different offset into the [skill map](skills.go#L29) where their class specific skills start, and then go 30 indexes into the map. So for example Assassin has an offset of `251`. Which means Assassin skills are  between the indexes of `251` and `281` which is exactly 30 indexes.

##### Layout
|  Type  | Bytes | Value                     |
|:------:|:-----:|---------------------------|
| Header | `2`   | `if`                      |
| Skills | `30`  | [[30]skill](skills.go#L29) |

##### Skill offset map
|    Class    | Offset |
|:-----------:|:------:|
| Amazon      | `6`    |
| Sorceress   | `36`   |
| Necromancer | `66`   |
| Paladin     | `96`   |
| Barbarian   | `126`  |
| Druid       | `221`  |
| Assassin    | `251`  |


#### Quests
The quests struct is `298 byte` section that describes all quests in the game but also contains data about act traveling and NPC introductions. Each quest is `2 byte` long.

##### Header
| Offset | Bytes | Content |
|--------|-------|---------|
| 335    | 4     | Woo!    |
| 339    | 6     | Unknown |

##### Quest structs
A quest is `2 byte` long, I've created a general `quest` struct that holds the most important data of a quest, if it's completed or not. Each quest has a lot of unique bits set depending on different milestones of the quest. For example if you've consumed the Scroll of Resistance from the quest "Prison of Ice" or not.

###### General 
| Bit | Description     |
|-----|-----------------|
| 0   | Quest completed |

##### Prison of Ice
[Prison of Ice ](quest.go#L26)is the only quest I bothered to implement, because I needed to know if the character has increased resistances from the scroll or not.

| Bit | Description     |
|-----|-----------------|
| 0   | Quest completed |
| 7   | Consumed scroll |


| Offset | Bytes      | Description                                                                             |
|--------|------------|-----------------------------------------------------------------------------------------|
| 0      | `2`        | Set to `1` if you have been introduced to Warriv in Act I.                              |
| 2      | `[6]quest` | All six quests for Act I.                                                               |
| 14     | `2`        | Set to `1` if you have traveled to Act II.                                              |
| 16     | `2`        | Set to `1` if you have been introduced to Jerhyn.                                       |
| 18     | `[6]quest` | All six quests for Act II.                                                              |
| 30     | `2`        | Set to `1` if you have traveled to Act III.                                             |
| 32     | `2`        | Set to `1` if you have been introduced to Hratli.                                       |
| 34     | `[6]quest` | All six quests for Act III.                                                             |
| 46     | `2`        | Set to `1` if you have traveled to Act IV.                                              |
| 48     | `2`        | Set to `1` if you have been introduced to Act IV. (which you have if you have traveled) |
| 50     | `[6]quest` | Act IV only has 3 quests, so the struct has 6 empty bytes here.                         |
| 62     | `2`        | Set to `1` if you have traveled to Act V.                                               |
| 64     | `2`        | Seems to be set to 1 after completing Terror's End and talking to Cain in act IV.       |
| 66     | `4`        | Seems to be some kind of padding.                                                       |
| 70     | `[6]quest` | All six quests for Act V.                                                               |
| 82     | 14         | Some kind of padding after all the quest data.                                          |


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

All magical properties are mapped in the [items.go](items.go) file.

### 5. Corpse data
If your character is currently dead, and a corpse is on the ground when you enter a game, your equipped items will be in this item struct. It's a corpse header `16 bytes` containing the header `JM` followed by an item count similar to the item list.

Reading the corpse items are done in the exact way as the previous section of items.

### 6. Expansion data
If your character is created in the expansion Lord of Destruction if will contain 2 more sections.

#### 7. Mercenary items
The mercenary sections starts of with a `2 byte` header with the value `jf` and is followed by a `4 byte` item header containing the number of items the mercenary is currently wearing. The items are read like any other item list.

#### 8. Golem
If your character is both a Necromancer and an expansion character, this section starts of with a `3 byte` header, where the first two bytes are the header `kf` followed by a boolean called `hasGolem`, if this value is true, there's an item list with the length 1 following the header.


----------
