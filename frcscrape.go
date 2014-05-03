// frcscrape.go
package frcscrape

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

var cmpStrings = map[string]string{
	"arc": "archimedes",
	"cur": "curie",
	"gal": "galileo",
	"new": "newton",
}

type Award struct {
	Team string
	Name string
}

type Match struct {
	MatchNumber  string
	RedAlliance  []string
	BlueAlliance []string
}

type NoData struct{}

var ErrNoData = errors.New("No data for event")

var allianceNums = map[int]int{
	0: 1,
	1: 4,
	2: 2,
	3: 3,
	8: 8,
	7: 5,
	6: 7,
	5: 6,
}

func filter(s []byte, fn func(byte) bool) []byte {
	var p []byte // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func trimWhitespace(s string) string {
	return strings.Replace(strings.TrimSpace(s), "\n", "", -1)
}

func removeUnicode(s string) string {
	q := filter([]byte(s), func(b byte) bool {
		return b <= 127
	})
	return string(q)
}

func getCodeForEvent(eventCode string) string {
	for event := range cmpStrings {
		if strings.EqualFold(strings.ToLower(eventCode), event) {
			return cmpStrings[event]
		}
	}
	return eventCode
}

func ScrapeAllianceSelections(eventCode string, year int) (map[int][]string, error) {
	url := fmt.Sprintf("http://www2.usfirst.org/%dcomp/events/%s/scheduleelim.html", year, getCodeForEvent(eventCode))

	doc, err := getDoc(url)
	if err != nil {
		return nil, err
	}

	tds := doc.Find("tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		t := s.Find("td")
		return t.Length() == 9
	})

	if tds.Length() <= 1 {
		return nil, ErrNoData
	}

	m := make(map[int][]string)

	tds.Slice(1, 5).Each(func(i int, s *goquery.Selection) {
		info := s.Find("td").Slice(3, 9)
		m[allianceNums[i]] = info.Slice(0, 3).Map(func(j int, q *goquery.Selection) string {
			return trimWhitespace(q.Text())
		})
		m[allianceNums[8-i]] = info.Slice(3, 6).Map(func(j int, q *goquery.Selection) string {
			return trimWhitespace(q.Text())
		})
	})
	return m, nil
}

func ScrapeAwardsForEvent(eventCode string, year int) ([]Award, error) {
	url := fmt.Sprintf("http://www2.usfirst.org/%dcomp/events/%s/awards.html", year, getCodeForEvent(eventCode))

	doc, err := getDoc(url)
	if err != nil {
		return nil, err
	}

	tds := doc.Find("tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		t := s.Find("td")
		return t.Length() == 5 && !strings.EqualFold(trimWhitespace(t.Eq(1).Text()), "")
	})

	if tds.Length() <= 2 {
		return nil, ErrNoData
	}

	m := []Award{}

	tds.Slice(2, tds.Length()).Each(func(i int, s *goquery.Selection) {
		info := s.Find("td")
		team := trimWhitespace(info.Eq(1).Text())
		award := strings.Replace(trimWhitespace(info.Eq(0).Text()), "  ", " ", -1)
		award = removeUnicode(award)
		a := Award{team, award}
		m = append(m, a)
	})
	return m, nil
}

func ScrapeTeamsForEvent(eventCode string, year int) ([]string, error) {
	var url string
	for event := range cmpStrings {
		if strings.EqualFold(strings.ToLower(eventCode), event) {
			url = fmt.Sprintf("https://my.usfirst.org/myarea/index.lasso?page=teamlist&event_type=FRC&sort_teams=number&year=%d&event=cmp&division=%s", year, cmpStrings[event])
			break
		}
	}
	if strings.EqualFold(url, "") {
		url = fmt.Sprintf("https://my.usfirst.org/myarea/index.lasso?page=teamlist&event_type=FRC&sort_teams=number&year=%d&event=%s", year, eventCode)
	}

	doc, err := getDoc(url)
	if err != nil {
		return nil, err
	}

	tds := doc.Find("tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		t := s.Find("td")
		return t.Length() > 0
	})

	if tds.Length() <= 3 {
		return nil, ErrNoData
	}

	m := []string{}

	tds.Slice(3, tds.Length()).Each(func(i int, s *goquery.Selection) {
		info := s.Find("td")
		team := trimWhitespace(info.Eq(2).Text())
		m = append(m, team)
	})
	return m, nil
}

func ScrapeAdvanceMatchQuals(eventCode, matchNumber string, year int) (Match, error) {
	return scrapeAdvanceMatch(eventCode, matchNumber, year, 0)
}

func ScrapeAdvanceMatchElims(eventCode, matchNumber string, year int) (Match, error) {
	return scrapeAdvanceMatch(eventCode, matchNumber, year, 1)
}

func scrapeAdvanceMatch(eventCode, matchNumber string, year, round int) (Match, error) {
	// type
	// 	0 = Quals
	// 	1 = Elims
	match, err := strconv.Atoi(matchNumber)
	if err != nil {
		return Match{}, err
	}

	url := fmt.Sprintf("http://www2.usfirst.org/%dcomp/events/%s/matchresults.html", year, getCodeForEvent(eventCode))

	doc, err := getDoc(url)
	if err != nil {
		return Match{}, err
	}

	tableBodies := doc.Find("tbody")
	if tableBodies.Length() == 0 {
		return Match{}, ErrNoData
	}

	table := tableBodies.Eq(round + 2)

	// Quals table is tableBodies.Eq(2), Elims table is tableBodies.Eq(3)
	trs := table.Find("tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		return s.Find("td").Length() > 1
	})

	if trs.Length() <= 1 || (trs.Length()-1) < (match+2) {
		return Match{}, ErrNoData
	}

	tds := trs.Eq(match + 2).Find("td")

	m := Match{tds.Eq(1).Text(),
		[]string{tds.Eq(2 + round).Text(), tds.Eq(3 + round).Text(), tds.Eq(4 + round).Text()},
		[]string{tds.Eq(5 + round).Text(), tds.Eq(6 + round).Text(), tds.Eq(7 + round).Text()}}

	if round == 0 {
		m.MatchNumber = fmt.Sprintf("Match %s", m.MatchNumber)
	}

	return m, nil
}

func getDoc(url string) (*goquery.Document, error) {
	docChan := make(chan *goquery.Document, 1)
	errChan := make(chan error, 1)
	go func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Referer", "usfirst.org")
		res, _ := client.Do(req)
		doc, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			errChan <- err
		}
		docChan <- doc
	}()
	select {
	case res := <-docChan:
		return res, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(time.Second * 10):
		return nil, errors.New("Timed out getting " + url)
	}
}
