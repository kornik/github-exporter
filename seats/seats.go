package seats

import (
	"context"
	"github.com/google/go-github/v56/github"
	"os"
)

func GetFreeSeats(client *github.Client) int {
	orgName := os.Getenv("GITHUB_ORG")
	if orgName == "" {
		panic("GITHUB_ORG environment variable not set")
	}
	org, _, err := client.Organizations.Get(context.Background(), orgName)
	if err != nil {
		panic(err)
	}
	plan := org.GetPlan()
	freeSeats := plan.GetSeats() - plan.GetFilledSeats()
	return freeSeats
}

func GetTotalSeats(client *github.Client) int {

	orgName := os.Getenv("GITHUB_ORG")
	if orgName == "" {
		panic("GITHUB_ORG environment variable not set")
	}
	org, _, err := client.Organizations.Get(context.Background(), orgName)
	if err != nil {
		panic(err)
	}
	plan := org.GetPlan()
	totalSeats := plan.GetSeats()
	return totalSeats
}

func GetFilledSeats(client *github.Client) int {

	orgName := os.Getenv("GITHUB_ORG")
	if orgName == "" {
		panic("GITHUB_ORG environment variable not set")
	}
	org, _, err := client.Organizations.Get(context.Background(), orgName)
	if err != nil {
		panic(err)
	}
	plan := org.GetPlan()
	filledSeats := plan.GetFilledSeats()
	return filledSeats
}
