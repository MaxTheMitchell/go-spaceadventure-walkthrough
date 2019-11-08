package spaceadventure

import "fmt"

func Start(planetarySystem PlanetarySystem) {
	printWelcome(planetarySystem)
	printGreeting(getResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	travel(promptForRandomOrSpecificDestination("want to go to a random planet Y/N?"))
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Println("Welcome to the %s!", planetarySystem.Name)
	fmt.Println("There are 8 planets to explore.")
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel(randomDestination bool) {
	if randomDestination {
		travelToRandomPlanet()
	} else {
		travelToPlanet(getResponseToPrompt("what planet?"))
	}
}

func promptForRandomOrSpecificDestination(promp string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getResponseToPrompt(promp)
		if choice == "Y" {
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
	return false
}

func getResponseToPrompt(prompt string) string {
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return response
}

func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
