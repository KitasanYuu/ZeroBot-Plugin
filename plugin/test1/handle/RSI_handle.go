package handle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
)

var UserID = ""
var handle = ""
var Display_Title = ""
var MainOg = ""
var SID = ""
var Rank = ""
var date = ""

type UserHandle struct{}

func (h *UserHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Errorf("doc.err.", err)
	}
	doc.Find(".account-profile").Find(".info").Find(".value").Each(func(i int, s *goquery.Selection) {
		//User := s.Text()
		//fmt.Printf("%s\n", User)
		if i == 0 {
			UserID = s.Text()
		}
		if i == 1 {
			handle = s.Text()
		}
		if i == 2 {
			Display_Title = s.Text()
		}
		if i == 3 {
			MainOg = s.Text()
		}
		if i == 4 {
			SID = s.Text()
		}
		if i == 5 {
			Rank = s.Text()
		}
	})

	doc.Find(".left-col").Find(".inner").Find(".value").Each(func(i int, s *goquery.Selection) {
		if i == 3 {
			date = s.Text()
		}

	})

	if MainOg == "" {
		MainOg = "无"
		SID = "无"
		Rank = "无"
	}

	fmt.Printf("%s\n", UserID)
	//fmt.Printf("%s\n", handle)
	fmt.Printf("%s\n", Display_Title)
	fmt.Printf("%s\n", MainOg)
	fmt.Printf("%s\n", SID)
	fmt.Printf("%s\n", Rank)
	fmt.Printf("%s\n", date)

}
