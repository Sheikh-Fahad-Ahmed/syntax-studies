# go-review

## Description

A review of Golang with a few simple practical applications

## Programs

1. Character class selector
2. Playlist

# 1. Character Class Selector

## Description

A lightweight, terminal-based character class selector where you can also create your own custom class.

## Features

1. Listing all the classes
2. Choosing your desired class
3. Creating your own custom class

## How to run

1. Clone the repo
2. run the app using `go run characterClass/main.go` command

## Output

```text
Welcome to Class Selector:
1. List all classes
2. Choose A Class
3. Create your custom class
Enter your choice
1

rogue: [80 knives dexterity]
knight: [100 sword strength]
mage: [60 staff intelligence]

1. List all classes
2. Choose A Class
3. Create your custom class
Enter your choice
3


name Your class:Barbarian
How much health does your class name:
200
Whats your weapon:
Axe
Whats your attribute:
Rage
1. List all classes
2. Choose A Class
3. Create your custom class
Enter your choice
1

Barbarian: [200 Axe Rage]
knight: [100 sword strength]
mage: [60 staff intelligence]
rogue: [80 knives dexterity]

1. List all classes
2. Choose A Class
3. Create your custom class
Enter your choice
2
1. Barbarian: [200 Axe Rage]
2. knight: [100 sword strength]
3. mage: [60 staff intelligence]
4. rogue: [80 knives dexterity]
Choose Your class
1



You chose Barbarian!
Your health is: 200
Your weapon: Axe
Your attribute: Rage


exit status 1


```

# 2. Playlist

## Description

a simple music playlist where you create songs and add them to a playlist

## Features

- Create songs
- Add them to playlist
- Play the Songs one by one
- Skip to Next song on the playlist

## How to run

1. Clone the repo
2. run the command `go run playlist/main.go`

## Output

```text
Now playing: Yellow by Coldplay
Next song is Sweden by C418
Next song is Drag Path by Twenty one Pilots
You are at the end of the playlist

```
