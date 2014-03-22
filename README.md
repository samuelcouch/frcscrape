# FRC Scrape

This is just my little Golang library for scraping event data from FRC event pages.

## Setting up

Just get goquery and then frcscrape

	github.com/PuerkitoBio/goquery
	github.com/ZachOrr/frcscrape

## Examples

Get all the teams for an event

	teams, _ := frcscrape.ScrapeTeamsForEvent("migbl", 2012)
	fmt.Println("Teams: ", strings.Join(teams, ", "))

Get all the awards for an event

	awards, _ := frcscrape.ScrapeAwardsForEvent("migbl", 2012)
	for _, a := range awards {
		fmt.Println(fmt.Sprintf("%s won the %s", a.Team, a.Name))
	}

Get the alliance selection results for an event

	alliances, _ := frcscrape.ScrapeAllianceSelections("migbl", 2012)
	for i, a := range alliances {
		fmt.Println(fmt.Sprintf("%s are on the %d alliance", strings.Join(a, ", "), i))
	}

## To Do

* Add some other useful scraping functionss
* Make some Godocs