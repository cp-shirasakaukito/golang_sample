package holiday

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"flag"
)

const (
	HOLIDAY_URL = "https://www.googleapis.com/calendar/v3/calendars/japanese__ja@holiday.calendar.google.com/events"
)

var (
	key = flag.String("key","","api key for google api")
)


/*************************
 * start 取得したjsonを格納する構造体を定義
 *************************/
type Date struct {
	Date string `json:"date"`
}

type Holiday struct {
	Name string `json:"summary"`
	Updated string `json:"updated"`
	Date Date `json:"start"`
}

type Holidays struct {
	Name string `json:"summary"`
	Updated string `json:"updated"`
	Items []Holiday `json:"items"`
}
/*************************
 * end 取得したjsonを格納する構造体を定義
 *************************/

func GetHoliday(){
	flag.Parse()

	//GoogleAPIからjsonを取得
	values := url.Values{}
	values.Add("key",*key)
	url := HOLIDAY_URL + "?" + values.Encode()


	//JSONをパースした構造体を取得
	HolidayInterface, err := GetJson(url, new(Holidays))
	//構造体Holidayにキャストし、ポインタから値に変換
	respHoliday := *(HolidayInterface.(*Holidays))

	if err != nil {
		panic(err)
		return
	}
	fmt.Print(respHoliday)
}


/**
UrlからJsonを取得して指定した構造体の形式に整形して返す
 */
func GetJson(url string, structure interface{}) (interface{}, error) {

	resp, err := http.Get(url)

	if err != nil {
		return structure, err
	}

	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
	defer resp.Body.Close()

	// Responseの内容を取得
	body, err := ioutil.ReadAll(resp.Body)
	if  err != nil {
		return structure, err
	}

	//JSON→構造体変換
	err = json.Unmarshal(body, structure)

	return structure, err
}