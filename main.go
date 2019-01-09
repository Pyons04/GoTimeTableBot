package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/headzoo/surf"
	"github.com/joho/godotenv"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func postTweet(content string) int {

	envLoad()
	consumerKey := os.Getenv("consumer_key")
	consumerSecret := os.Getenv("consumer_key_secret")
	accessToken := os.Getenv("access_token")
	accessTokenSecret := os.Getenv("access_token_secret")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	_, err := api.PostTweet(content, nil)
	if err != nil {
		fmt.Println(err)
	}

	return 0
}

func generateContent(accessData string, data string, status string, oneSub string, oneInfo string, twoSub string, twoInfo string, threeSub string, threeInfo string, fourSub string, fourInfo string, fiveSub string, fiveInfo string) string {
	accessData = (accessData + "\n")

	dataInfo := ("明日は" + data + status)
	timetableInfo1 := ("１限" + oneSub + oneInfo + "\n")
	timetableInfo2 := ("２限" + twoSub + twoInfo + "\n")
	timetableInfo3 := ("３限" + threeSub + threeInfo + "\n")
	timetableInfo4 := ("４限" + fourSub + fourInfo + "\n")
	timetableInfo5 := ("５限" + fiveSub + fiveInfo + "\n")
	tweet := (accessData + dataInfo + timetableInfo1 + timetableInfo2 + timetableInfo3 + timetableInfo4 + timetableInfo5)

	tweet = strings.Replace(tweet, " ", "", -1)
	tweet = strings.Replace(tweet, "　", "", -1)
	fmt.Println(tweet)

	return tweet
}

func main() {

	browser := surf.NewBrowser()

	browser.Open("https://aoyama-portal.aoyama.ac.jp/aogaku_auth/jsp/AUTH01.jsp?TYPE=33554433&REALMOID=06-2a27689c-587a-42e0-834c-7a6751c24219&GUID=0&SMAUTHREASON=0&METHOD=GET&SMAGENTNAME=W02MLRk5ZzW6lo1jgEk6cBozc3Jz0O9V2D3FFJOoCKp0KOwIGgLV9PDcb8W1gma1&TARGET=-SM-https%3a%2f%2faoyama--portal%2eaoyama%2eac%2ejp%2fprotect%2f")
	fm, err := browser.Form("form")
	if err != nil {
		fmt.Println(err.Error())
	}
	envLoad()

	fm.Input("USER", os.Getenv("studentid"))
	fm.Input("PASSWORD", os.Getenv("password"))

	if fm.Submit() != nil {
		fmt.Println(err.Error())
	}

	browser.Open("https://aguinfo.jm.aoyama.ac.jp/AGUInfo/jikanwari.aspx")

	accessData := browser.Find("#div_asof").Text()

	tomorrowDate := strings.Trim(browser.Find("#cph_content_rptJikanwari_hyp_date2").Text(), "\n")
	tomorrowStatus := strings.Trim(browser.Find("#cph_content_rptJikanwari_Item2_0").Text(), "\n")

	tomorrowOneSubject := strings.Trim(browser.Find("#cph_content_rptJikanwari_Item2_1").Text(), "\n")
	tomorrowOneInfo, _ := browser.Find("#cph_content_rptJikanwari_Item2_1").Attr("class")
	if tomorrowOneInfo == "" {
		tomorrowOneInfo = "１限の教室変更、休講情報はありません。"
	} else if tomorrowOneInfo == "bkgnd-kyuukou" {
		tomorrowOneInfo = "１限に休講情報があります。"
	} else if tomorrowOneInfo == "bkgnd-roomhenkou" {
		tomorrowOneInfo = "１限に教室変更情報があります。"
	}

	tomorrowTwoSubject := strings.Trim(browser.Find("#cph_content_rptJikanwari_Item2_2").Text(), "\n")
	tomorrowTwoInfo, _ := browser.Find("#cph_content_rptJikanwari_Item2_2").Attr("class")

	if tomorrowTwoInfo == "" {
		tomorrowTwoInfo = "２限の教室変更、休講情報はありません。"
	} else if tomorrowTwoInfo == "bkgnd-kyuukou" {
		tomorrowOneInfo = "２限に休講情報があります。"
	} else if tomorrowTwoInfo == "bkgnd-roomhenkou" {
		tomorrowOneInfo = "２限に教室変更情報があります。"
	}

	tomorrowThreeSubject := strings.Trim(browser.Find("#cph_content_rptJikanwari_Item2_3").Text(), "\n")
	tomorrowThreeInfo, _ := browser.Find("#cph_content_rptJikanwari_Item2_3").Attr("class")

	if tomorrowThreeInfo == "" {
		tomorrowThreeInfo = "３限の教室変更、休講情報はありません。"
	} else if tomorrowThreeInfo == "bkgnd-kyuukou" {
		tomorrowOneInfo = "３限に休講情報があります。"
	} else if tomorrowThreeInfo == "bkgnd-roomhenkou" {
		tomorrowOneInfo = "３限に教室変更情報があります。"
	}

	tomorrowFourSubject := strings.Trim(browser.Find("#cph_content_rptJikanwari_Item2_4").Text(), "\n")
	tomorrowFourInfo, _ := browser.Find("#cph_content_rptJikanwari_Item2_4").Attr("class")

	if tomorrowFourInfo == "" {
		tomorrowFourInfo = "４限の教室変更、休講情報はありません。"
	} else if tomorrowFourInfo == "bkgnd-kyuukou" {
		tomorrowOneInfo = "４限に休講情報があります。"
	} else if tomorrowFourInfo == "bkgnd-roomhenkou" {
		tomorrowOneInfo = "４限に教室変更情報があります。"
	}

	tomorrowFiveSubject := strings.Trim(browser.Find("#cph_content_rptJikanwari_Item2_5").Text(), "\n")
	tomorrowFiveInfo, _ := browser.Find("#cph_content_rptJikanwari_Item2_5").Attr("class")

	if tomorrowFiveInfo == "" {
		tomorrowFiveInfo = "5限の教室変更、休講情報はありません。"
	} else if tomorrowFiveInfo == "bkgnd-kyuukou" {
		tomorrowOneInfo = "５限に休講情報があります。"
	} else if tomorrowFiveInfo == "bkgnd-roomhenkou" {
		tomorrowOneInfo = "５限に教室変更情報があります。"
	}

	tweet := generateContent(accessData, tomorrowDate, tomorrowStatus, tomorrowOneSubject, tomorrowOneInfo, tomorrowTwoSubject, tomorrowTwoInfo, tomorrowThreeSubject, tomorrowThreeInfo, tomorrowFourSubject, tomorrowFourInfo, tomorrowFiveSubject, tomorrowFiveInfo)
	postTweet(tweet)
}
