package main

import "github.com/MaxTheMitchell/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(spaceadventure.PlanetarySystem{Name: "solar system", Planets: mockPlanets()})
}

func mockPlanets() []spaceadventure.Planet {
	return []spaceadventure.Planet{
		spaceadventure.Planet{Name: "Mars", Description: "the red one"},
		spaceadventure.Planet{Name: "Earth", Description: "The 70% blue one"},
		spaceadventure.Planet{Name: "Venus", Description: "The yellowish one"},
		spaceadventure.Planet{Name: "Saturn", Description: "The other yellow one"},
		spaceadventure.Planet{Name: "Jupitur", Description: "get more studier"},
		spaceadventure.Planet{Name: "Moon", Description: "Idk if this one is actualy a planet"},
		spaceadventure.Planet{Name: "pluto", Description: "Micky mouse's dog"},
	}
}
