package main

import "github.com/MaxTheMitchell/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(spaceadventure.PlanetarySystem{Name: "solar system", Planets: []spaceadventure.Planet{
		spaceadventure.Planet{Name: "Mars", Description: "the red one"}}})
}
