package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type scraperResult struct {
	companyName string
	title       string
	lookingFor  string
	degree      string
	location    string
	hireType    string
	salary      string
	description string
}

func (s scraperResult) String() string {
	return "[" + s.companyName + "] " + s.title
}

func main() {
	baseURL := "https://www.jobkorea.co.kr/recruit/joblist?menucode=duty&dutyCtgr=10016#anchorGICnt_"
	jobkoreaJobs := scraper(baseURL)

	for _, jobkoreaJob := range jobkoreaJobs {
		fmt.Println(jobkoreaJob)
	}
}

func scraper(url string) []scraperResult {
	// Request the HTML page.
	res, err := http.Get(url)
	errorCheck(err)
	respCheck(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	errorCheck(err)

	// get total page of url
	totalPageOfUrl := getTotalPages(doc)

	// scrapping
	var jobs []scraperResult
	for i := 11; i <= totalPageOfUrl; i++ {
		targetPage := url + strconv.Itoa(i)
		fmt.Println(targetPage, "page scrapping...")
		res, err := http.Get(targetPage)
		errorCheck(err)
		respCheck(res)

		doc, err := goquery.NewDocumentFromReader(res.Body)
		errorCheck(err)

		doc.Find(".tplList.tplJobList").Find(".devloopArea").Each(func(i int, hireCard *goquery.Selection) {
			var job = scraperResult{}

			job.companyName = cleanString(hireCard.Find(".tplCo>a").Text())

			rightPart := hireCard.Find(".tplTit")
			job.title = cleanString(rightPart.Find(".link.normalLog").Text())

			rightPart.Find(".cell").Each(func(i int, s *goquery.Selection) {
				//fmt.Println(strings.Join(strings.Fields(strings.TrimSpace(s.Find(".cell").Text())), " "))
				//fmt.Println(i, s.Text())
				switch i {
				case 0:
					job.lookingFor = s.Text()
				case 1:
					job.degree = s.Text()
				case 2:
					job.location = s.Text()
				case 3:
					job.hireType = s.Text()
				case 4:
					job.salary = s.Text()
				}
			})
			job.description = rightPart.Find(".desc").Text()

			jobs = append(jobs, job)
		})
	}

	return jobs
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "")
}

func getTotalPages(doc *goquery.Document) int {
	pageStr, exists := doc.Find(".tplPagination.newVer").Find(".tplBtn.btnPgnNext").Attr("data-page")
	if exists {
		page, err := strconv.Atoi(pageStr)
		errorCheck(err)
		return page
	} else {
		return 1
	}
}

func respCheck(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
