package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
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
	howManyPage := 11
	mainChannel := make(chan []scraperResult)

	go scraper(baseURL, howManyPage, mainChannel)

	scraperResultsToCsvFile(<-mainChannel)
}

func scraper(url string, howManyPage int, mainChannel chan<- []scraperResult) {
	// Request the HTML page.
	//res, err := http.Get(url)
	//errorCheck(err)
	//respCheck(res)

	//defer res.Body.Close()

	//doc, err := goquery.NewDocumentFromReader(res.Body)
	//errorCheck(err)

	// get total page of url
	//totalPageOfUrl := getTotalPages(doc)

	// scrapping
	c := make(chan []scraperResult)
	for i := 2; i <= howManyPage; i++ {
		targetUrl := url + strconv.Itoa(i)
		go scrapping(targetUrl, c)
	}

	mainChannel <- <-c
}

func scrapping(targetUrl string, c chan<- []scraperResult) {
	//fmt.Println(targetUrl, "page scrapping...")
	res, err := http.Get(targetUrl)
	errorCheck(err)
	respCheck(res)
	fmt.Println(targetUrl, "scrapping result :", res.Status)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	errorCheck(err)

	var jobs []scraperResult
	doc.Find(".tplList.tplJobList").Find(".devloopArea").Each(func(i int, hireCard *goquery.Selection) {
		var job scraperResult

		job.companyName = cleanString(hireCard.Find(".tplCo>a").Text())

		rightPart := hireCard.Find(".tplTit")
		job.title = cleanString(rightPart.Find(".link.normalLog").Text())

		rightPart.Find(".cell").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				job.lookingFor = cleanString(s.Text())
			case 1:
				job.degree = cleanString(s.Text())
			case 2:
				job.location = cleanString(s.Text())
			case 3:
				job.hireType = cleanString(s.Text())
			case 4:
				job.salary = cleanString(s.Text())
			}
		})

		job.description = rightPart.Find(".desc").Text()

		if job.title != "" { // filtering invalid data
			jobs = append(jobs, job)
		}
	})

	fmt.Println(jobs)
	c <- jobs
}

func scraperResultsToCsvFile(jobs []scraperResult) {
	file, err := os.Create("practice/job_scraper/csv/jobkorea.csv")
	errorCheck(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"companyName", "title", "lookingFor", "degree", "location", "hireType", "salary", "description"}
	wErr := w.Write(headers)
	errorCheck(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.companyName, job.title, job.lookingFor, job.degree, job.location, job.hireType, job.salary, job.description}
		wErr := w.Write(jobSlice)
		errorCheck(wErr)
	}
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
