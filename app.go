package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

type PublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type Tweet struct {
	Id      string        `json:"id"`
	Text    string        `json:"text"`
	Metrics PublicMetrics `json:"public_metrics"`
}

type Meta struct {
	OldestId    string `json:"oldest_id"`
	NewestId    string `json:"newest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type twitterResponse struct {
	Data []Tweet `json:"data"`
	Meta `json:"meta"`
}

type response struct {
	TweetCount  int     `json:"tweet_count"`
	Tweets      []Tweet `json:"tweets"`
	MostLiked   string  `json:"most_liked"`
	MostReplied string  `json:"most_replied"`
}

// Types needed for sorting
type byLikes []Tweet

func (s byLikes) Len() int {
	return len(s)
}

func (s byLikes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLikes) Less(i, j int) bool {
	return s[i].Metrics.LikeCount < s[j].Metrics.LikeCount
}

type byReplies []Tweet

func (s byReplies) Len() int {
	return len(s)
}

func (s byReplies) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byReplies) Less(i, j int) bool {
	return s[i].Metrics.ReplyCount < s[j].Metrics.ReplyCount
}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func mostLikedTweet(tweets []Tweet) string {
	liked := byLikes(tweets)
	sort.Sort(liked)
	if len(liked) > 0 {
		return liked[len(liked)-1].Id
	}
	return ""
}

func mostRepliedTweet(tweets []Tweet) string {
	replied := byReplies(tweets)
	sort.Sort(replied)
	if len(replied) > 0 {
		return replied[len(replied)-1].Id
	}
	return ""
}

func getTweets(c *gin.Context) {

	const MO_ID = "22176791"

	API_URL := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets", MO_ID)

	CMT_BEARER_TOKEN := os.Getenv("CMT_BEARER_TOKEN")

	twitterClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, API_URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", CMT_BEARER_TOKEN))

	// Setup query params to Twitter API
	q := req.URL.Query()
	q.Add("exclude", "retweets")
	q.Add("max_results", "100")
	q.Add("tweet.fields", "public_metrics")

	if endTime := c.Query("end_time"); endTime != "" {
		q.Add("end_time", endTime)
	}
	if startTime := c.Query("start_time"); startTime != "" {
		q.Add("start_time", startTime)
	}

	paginationToken := ""

	resp := response{}
	tweets := make([]Tweet, 0)

	for {
		q.Del("pagination_token")
		if paginationToken != "" {
			q.Add("pagination_token", paginationToken)
		}

		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())

		res, getErr := twitterClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		tResp := twitterResponse{}
		jsonErr := json.Unmarshal(body, &tResp)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		tweets = append(tweets, tResp.Data...)

		paginationToken = tResp.Meta.NextToken
		if paginationToken == "" {
			break
		}

	}

	resp.TweetCount = len(tweets)
	resp.Tweets = tweets

	resp.MostLiked = mostLikedTweet(tweets)
	resp.MostReplied = mostRepliedTweet(tweets)

	c.JSON(http.StatusOK, resp)

}

func main() {
	router := gin.New()

	// Serve client and related assets
	router.Static("/css", "./dist/css")
	router.Static("/js", "./dist/js")
	router.Static("/img", "./dist/img")

	router.StaticFile("/apple-touch-icon.png", "./dist/apple-touch-icon.png")
	router.StaticFile("/favicon-32x32.png", "./dist/favicon-32x32.png")
	router.StaticFile("/favicon-16x16.png", "./dist/favicon-16x16.png")
	router.StaticFile("/site.webmanifest", "./dist/site.webmanifest")

	router.LoadHTMLGlob("dist/index.html")

	router.GET("/api/tweets", getTweets)

	router.GET("/", indexPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
