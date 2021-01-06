package service

import (
	"fmt"
	"holycode-task/model"
	"holycode-task/repository/postgres"
	"time"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/gocolly/colly"
)

type Scraper struct {
	store *postgres.Store
	insta *goinsta.Instagram
}

func NewScraper(ig *goinsta.Instagram, db *postgres.Store) *Scraper {

	return &Scraper{insta: ig, store: db}
}

// First check if this instagram account is stored in database, if not, then load from instagram
func (s *Scraper) SearchInstagramByUsername(username string) (*model.InstagramAccount, error) {
	acc, _ := s.store.FindInstagramAccountByUsername(username)
	if acc != nil {
		return acc, nil
	}
	// If account is not present in database, we load inforamtions from instagram
	searchedUser, err := s.insta.Profiles.ByName(username)
	if err != nil {
		return nil, err
	}

	account := &model.InstagramAccount{
		Username:      searchedUser.Username,
		FullName:      searchedUser.FullName,
		Biography:     searchedUser.Biography,
		ProfilePicURL: searchedUser.ProfilePicURL,
		Email:         searchedUser.Email,
		MediaCount:    searchedUser.MediaCount,
		FollowerCount: searchedUser.FollowerCount,
		UsertagsCount: searchedUser.UsertagsCount,
	}

	if err = s.store.SaveInstagramAccount(account); err != nil {
		return nil, fmt.Errorf("Some error occuurred saving account in database: %v", err)
	}

	return account, nil
}

func (s *Scraper) ScrapeFacebookProfile(username string) (*model.FacebookResponse, error) {
	//First loadfrom database if present
	fb, _ := s.store.FindFacebookAccountByUsername(username)
	if fb != nil {
		return fb, nil
	}
	//if not present, scrape and save to database
	c := colly.NewCollector(
		colly.Async(false),
	)
	_ = c.Limit(&colly.LimitRule{
		DomainRegexp: "127.0.0.1:8080",
		Parallelism:  1,
		Delay:        5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode, r.Request.URL)
	})

	accountToSave := fetchUserInfo(c, username)
	err := s.store.SaveFacebookAccount(accountToSave)
	if err != nil {
		return nil, fmt.Errorf("ERROR SAVING TO DATABASE : %v", err)
	}
	return accountToSave, nil
}

func fetchUserInfo(coll *colly.Collector, username string) *model.FacebookResponse {
	infos := []*string{}
	coll.OnHTML("._4bl9", func(e *colly.HTMLElement) {
		temp := e.ChildText("div")
		infos = append(infos, &temp)
	})
	if err := coll.Visit("https://www.facebook.com/" + username); err != nil {
		fmt.Errorf("error visitingt site: %s", err.Error())
	}
	time.Sleep((1 * time.Second))

	for _, i := range infos {
		fmt.Println("\t\t ---->", *i)
	}
	likes := "unkwown"
	followers := "unknown"
	if len(infos) >= 2 {
		likes = *infos[1]
		followers = *infos[2]
	}
	accountToSave := &model.FacebookResponse{
		Username:  username,
		Likes:     likes,
		Followers: followers,
	}
	infos = []*string{}

	return accountToSave
}
