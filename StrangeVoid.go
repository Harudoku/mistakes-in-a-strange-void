package main

import "fmt"

type TestType struct{
    X int
}

type TileProperties struct {
    // The innate properties of a tile, whether or not an object is there
    // This probably belongs in its own package but I do not know how file
    // structures are supposed to look so I'm going to move fast and break
    // things.
    passable bool
    movable bool
    isPit bool
}

type Tile struct {
    // The full information about a tile, including whether it contains things;
    // once there are more objects it's probably correct to have an array
    // instead of "hasBoulder" and "hasPlayer" but this is a quick mockup.
    locX int
    locY int
    properties TileProperties
    // can you tell that I come from OOP
    // I might as well do stupid shit now to see whether it breaks
    hasBoulder bool
    hasPlayer bool
}

func main() {
    // wallTile := TileProperties{false, false, false}
    clearTile := TileProperties{true, true, false}
    // pitTile := TileProperties{false, true, true}
    // I assume these things would be defined wherever TileProperties is defined
    // but I have no idea what the data structure looks like lmao
    example := Tile{0, 0, clearTile, false, false}
    fmt.Println(example.properties.passable)
    fmt.Println(example.locX)
}
