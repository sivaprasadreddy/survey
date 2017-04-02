package main

import (
	"fmt"

	"github.com/alecaivazis/survey"
)

// the questions to ask
var simpleQs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What is your name?",
			Default: "Johnny Appleseed",
		},
	},
	{
		Name: "color",
		Prompt: &survey.Choice{
			Message: "Choose a color:",
			Choices: []string{"red", "blue", "green", "yellow"},
			Default: "yellow",
		},
		Validate: survey.Required,
	},
	{
		Name: "days",
		Prompt: &survey.Checkbox{
			Message:  "What days do you prefer:",
			Options:  []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			Defaults: []string{"Saturday", "Sunday"},
		},
	},
}

func main() {

	answers, err := survey.Ask(simpleQs)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s chose %s and prefers %s.\n", answers["name"], answers["color"], answers["days"])
}
