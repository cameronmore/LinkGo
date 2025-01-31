package link

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cameronmore/LinkGo/LinkGo/funcs"
)

type Link struct {
	URL                string `json:"url"`
	ResponseCode       int    `json:"response_status_code"`
	ResponseCodeHealth string `json:"response_status"`
	TimeOfCheck        string `json:"time_of_validation"`
}

func New(s string) *Link {
	return &Link{
		URL: s,
	}
}

// we need to take a link, request it, parse the response,

func (l *Link) Validate() error {
	r, err := http.Get(l.URL)
	if err != nil {
		return err
	}
	l.TimeOfCheck = time.Now().UTC().Format(time.RFC3339)
	responseInt, err := strconv.Atoi(funcs.FirstN(r.Status, 3))
	if err != nil {
		return err
	}
	l.ResponseCodeHealth = strings.Join(strings.Split(r.Status, " ")[1:], " ")
	l.ResponseCode = responseInt
	/*
		if responseInt > 199 && responseInt < 300 {
		}
	*/
	return nil
}

func (l Link) Display() {
	fmt.Println(l)
}
