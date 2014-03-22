// frcscrape_test.go
package frcscrape

import (
	"testing"
	"reflect"
	"strings"
)

func TestScrapeAdvanceMatchQualsNoData(t *testing.T) {
	_, err := ScrapeAdvanceMatchQuals("migbl", "79", 2013)
	if err != ErrNoData {
		t.Fatalf("Not handling scraping with big number properly | %v", err)
	}
}

func TestScrapeAdvanceMatchQuals(t *testing.T) {
	m1, err := ScrapeAdvanceMatchQuals("migbl", "1", 2013)
	if err != nil {
		t.Fatalf("Couldn't get advanced match for quals! %v", err)
	}
	m2 := []string{"1506", "302", "1025"}
	m3 := []string{"2604", "1322", "3667"}
	if !strings.EqualFold(m1.MatchNumber, "Match 3") {
		t.Fatalf("Pulling wrong match for advanced!")
	}
	eq1 := reflect.DeepEqual(m1.RedAlliance, m2)
	if !eq1 {
		t.Fatalf("Scraping advanced match alliances wrong!")
	}
	eq2 := reflect.DeepEqual(m1.BlueAlliance, m3)
	if !eq2 {
		t.Fatalf("Scraping advanced match alliances wrong!")
	}
}

func TestScrapeAdvanceMatchElimsNoData(t *testing.T) {
	_, err := ScrapeAdvanceMatchElims("migbl", "19", 2013)
	if err != ErrNoData {
		t.Fatalf("Not handling scraping with big number properly elims | %v", err)
	}
}

func TestScrapeAdvanceMatchElims(t *testing.T) {
	m1, err := ScrapeAdvanceMatchElims("migbl", "3", 2013)
	if err != nil {
		t.Fatalf("Couldn't get advanced match for elims! %v", err)
	}
	m2 := []string{"1718", "33", "247"}
	m3 := []string{"3570", "3667", "3535"}
	if !strings.EqualFold(m1.MatchNumber, "Qtr 1-2") {
		t.Fatalf("Pulling wrong match for advanced elims!")
	}
	eq1 := reflect.DeepEqual(m1.RedAlliance, m2)
	if !eq1 {
		t.Fatalf("Scraping advanced match alliances wrong elims!")
	}
	eq2 := reflect.DeepEqual(m1.BlueAlliance, m3)
	if !eq2 {
		t.Fatalf("Scraping advanced match alliances wrong elims!")
	}
}

func TestScrapeAllianceSelections(t *testing.T) {
	m1, err := ScrapeAllianceSelections("migbl", 2013)
	if err != nil {
		t.Fatalf("Couldn't scrape alliance selections! %v", err)
	}
	m2 := map[int][]string {
		1:[]string{"33", "1718", "247"}, 
		8:[]string{"3570", "3535", "3667"}, 
		4:[]string{"2145", "2612", "3620"}, 
		5:[]string{"51", "4810", "703"}, 
		2:[]string{"2619", "3302", "245"}, 
		7:[]string{"3322", "1025", "548"}, 
		3:[]string{"573", "3098", "1684"}, 
		6:[]string{"4382", "4405", "1504"},
	}
	eq := reflect.DeepEqual(m1, m2)
	if !eq {
		t.Fatalf("Scrape elim data has changed!")
	}
}

func TestScrapeAllianceSelectionsCMP(t *testing.T) {
	m1, err := ScrapeAllianceSelections("gal", 2013)
	if err != nil {
		t.Fatalf("Couldn't scrape elimination alliance selections! %v", err)
	}
	m2 := map[int][]string {
		1:[]string{"4039", "1114", "118"}, 
		8:[]string{"1732", "245", "27"}, 
		4:[]string{"2630", "3641", "111"}, 
		5:[]string{"1241", "610", "1477"}, 
		2:[]string{"3284", "2175", "2169"}, 
		7:[]string{"2338", "1323", "2512"}, 
		3:[]string{"2337", "1425", "2474"}, 
		6:[]string{"1806", "3656", "4334"},
	}
	eq := reflect.DeepEqual(m1, m2)
	if !eq {
		t.Fatalf("Scrape cmp elim data has changed!")
	}
}

func TestScrapeAwardsForEvent(t *testing.T) {
	m1, err := ScrapeAwardsForEvent("migbl", 2013)
	if err != nil {
		t.Fatalf("Couldn't scrape awards! %v", err)
	}
	m2 := []Award{
		Award{"1718", "District Chairman's Award"}, 
		Award{"2604", "Engineering Inspiration Award"}, 
		Award{"247", "District Winners #3"}, 
		Award{"1718", "District Winners #2"}, 
		Award{"33", "District Winners #1"}, 
		Award{"3302", "District Finalists #3"}, 
		Award{"245", "District Finalists #2"}, 
		Award{"2619", "District Finalists #1"}, 
		Award{"2145", "Industrial Safety Award sponsored by Underwriters Laboratories"}, 
		Award{"4810", "Highest Rookie Seed"}, 
		Award{"302", "Judges' Award"}, 
		Award{"4810", "Rookie All Star Award"}, 
		Award{"4507", "Rookie Inspiration Award"}, 
		Award{"245", "Entrepreneurship Award sponsored by Kleiner Perkins Caufield and Byers"}, 
		Award{"573", "Team Spirit Award sponsored by Chrysler"}, 
		Award{"1684", "Excellence in Engineering Award sponsored by Delphi"}, 
		Award{"2619", "Gracious Professionalism Award sponsored by Johnson & Johnson"}, 
		Award{"1025", "Creativity Award sponsored by Xerox"}, 
		Award{"33", "Quality Award sponsored by Motorola"}, 
		Award{"1504", "Innovation in Control Award sponsored by Rockwell Automation"}, 
		Award{"2145", "Industrial Design Award sponsored by General Motors"}, 
		Award{"3322", "Imagery Award in honor of Jack Kamen"},
	}
	eq := reflect.DeepEqual(m1, m2)
	if !eq {
		t.Fatalf("Scrape awards data has changed!")
	}
}

func TestScrapeAwardsForEventCMP(t *testing.T) {
	m1, err := ScrapeAwardsForEvent("gal", 2013)
	if err != nil {
		t.Fatalf("Couldn't scrape cmp awards! %v", err)
	}
	m2 := []Award{
		Award{"1241", "Championship Division Winners - Galileo #1"}, 
		Award{"1477", "Championship Division Winners - Galileo #2"}, 
		Award{"610", "Championship Division Winners - Galileo #3"},
		Award{"2169", "Championship Division Finalists - Galileo #1"}, 
		Award{"3284", "Championship Division Finalists - Galileo #2"}, 
		Award{"2175", "Championship Division Finalists - Galileo #3"}, 
		Award{"4557", "Highest Rookie Seed - Galileo"}, 
		Award{"4472", "Rookie All Star Award - Galileo"}, 
		Award{"4601", "Rookie Inspiration Award - Galileo"}, 
		Award{"3211", "Judges' Award - Galileo"}, 
	}
	eq := reflect.DeepEqual(m1, m2)
	if !eq {
		t.Fatalf("Scrape cmp awards data has changed!")
	}
}

func TestScrapeTeamsForEvent(t *testing.T) {
	m1, err := ScrapeTeamsForEvent("migbl", 2013)
	if err != nil {
		t.Fatalf("Couldn't scrape teams! %v", err)
	}
	m2 := []string{"33","51","66","245","247","302","548","573","703","894","1025","1322","1504","1506","1684","1718","2145","2604","2612","2619","3098","3302","3322","3415","3534","3535","3536","3568","3570","3617","3620","3667","4327","4375","4382","4405","4507","4810","4827","4839"}
	eq := reflect.DeepEqual(m1, m2)
	if !eq {
		t.Fatalf("Scrape teams data has changed!")
	}
}

func TestScrapeTeamsForEventCMP(t *testing.T) {
	m1, err := ScrapeTeamsForEvent("gal", 2013)
	if err != nil {
		t.Fatalf("Couldn't scrape cmp teams! %v", err)
	}
	m2 := []string{"27", "45", "70", "95", "111", "118", "125", "151", "192", "222", "245", "295", "329", "337", "358", "384", "422", "447", "467", "578", "610", "744", "842", "1086", "1114", "1218", "1241", "1323", "1325", "1378", "1405", "1425", "1429", "1477", "1629", "1675", "1710", "1726", "1732", "1772", "1801", "1806", "1912", "1987", "2000", "2046", "2169", "2175", "2199", "2259", "2337", "2338", "2341", "2403", "2474", "2481", "2485", "2502", "2512", "2630", "2648", "2729", "2809", "2834", "2907", "2978", "3018", "3132", "3189", "3211", "3284", "3459", "3481", "3528", "3641", "3656", "3941", "3944", "4011", "4026", "4039", "4069", "4158", "4334", "4452", "4462", "4472", "4481", "4492", "4502", "4522", "4541", "4557", "4567", "4579", "4601", "4607", "4627", "4641", "4797"}
	eq := reflect.DeepEqual(m1, m2)
	if !eq {
		t.Fatalf("Scrape cmp teams data has changed!")
	}
}