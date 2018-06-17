/**
#############################################################
#											                #
#		           Made By: DigitalQubit		            #
#  (https://github.com/DigitalQubit/Monero-Status-Checker)  #
#############################################################

**/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/dariubs/percent"
	"github.com/gizak/termui"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	currData    Stats
	refresh     = false
	hashLC      = termui.NewLineChart()
	sharesRight = termui.NewGauge()
	amountDue   = termui.NewPar("")
	amountPaid  = termui.NewPar("")
	currRate    = termui.NewPar("")
)

type Stats struct {
	Hash          float64 `json:"hash"`
	ValidShares   int     `json:"validShares"`
	InvalidShares int     `json:"invalidShares"`
	AmtPaid       float64 `json:"amtPaid"`
	AmtDue        float64 `json:"amtDue"`
	Identifier    string
	LastHash      int `json:"lastHash"`
	TotalHashes   int `json:"totalHashes"`
	Expiry        int
	TxnCount      int `json:"txnCount"`
}

func SetAddress(naddress string, pool int) string {
	var addressURL string
	switch pool {
	case 1:
		addressURL = "https://www.supportxmr.com/api/"
	case 2:
		addressURL = "https://api.xmrpool.net/"
	case 3:
		addressURL = "https://api.viaxmr.com/"
	case 4:
		addressURL = "https://monero.hashvault.pro/api/"
	case 5:
		addressURL = "https://api.moriaxmr.com/"
	case 6:
		addressURL = "https://api.moneroocean.stream/"
	}
	addressURL += "miner/" + naddress + "/stats"
	return addressURL
}

func GetStats(c *colly.Collector, saddress string) {
	for {
		c.Visit(saddress)
		time.Sleep(10000 * time.Millisecond)
	}
}

func SettingsColl(cd *colly.Collector) {

	cd.WithTransport(&http.Transport{
		DisableKeepAlives: false,
	})
	cd.OnResponse(func(r *colly.Response) {
		var currData Stats
		json.Unmarshal(r.Body, &currData)
		setStats(&currData)
	})
	extensions.RandomUserAgent(cd)
	cd.OnError(func(r *colly.Response, err error) {
		fmt.Println("An error occurred while requesting site!")
		os.Exit(-1)
	})
	cd.AllowURLRevisit = true
}

func settingsTerm() {
	hashLC.BorderLabel = "Hashrate History"
	hashLC.Height = 10
	hashLC.Width = termui.TermWidth() * 2
	hashLC.Mode = "dot"
	hashLC.Data = []float64{}

	sharesRight.BorderLabel = "Percentage of Valid Shares"
	sharesRight.Width = termui.TermWidth() * 2
	sharesRight.Height = 5
	sharesRight.PercentColor = termui.ColorBlack

	amountDue.BorderLabel = "Amount Due"
	amountDue.Width = (termui.TermWidth() / 3) * 2
	amountDue.Height = 3
	amountDue.TextFgColor = termui.ColorBlack
	amountDue.TextBgColor = termui.ColorYellow

	amountPaid.BorderLabel = "Amount Paid"
	amountPaid.Width = (termui.TermWidth() / 3) * 2
	amountPaid.Height = 3
	amountPaid.TextFgColor = termui.ColorBlack
	amountPaid.TextBgColor = termui.ColorGreen

	currRate.BorderLabel = "Current Hashrate"
	currRate.Width = (termui.TermWidth() / 3) * 2
	currRate.Height = 3

	row1 := termui.NewRow(termui.NewCol(6, 0, hashLC))
	row2 := termui.NewRow(termui.NewCol(6, 0, sharesRight))
	row3 := termui.NewRow(termui.NewCol(2, 0, amountDue), termui.NewCol(2, 0, amountPaid), termui.NewCol(2, 0, currRate))
	row1.SetWidth(termui.TermWidth())
	row2.SetWidth(termui.TermWidth())
	row3.SetWidth(termui.TermWidth())
	termui.Body.AddRows(row1, row2, row3)
	termui.Body.Width = termui.TermWidth() * 2
	termui.Body.Align()

}

func setStats(cData *Stats) {
	if len(hashLC.Data) > 60 {
		hashLC.Data = hashLC.Data[1:]
	}
	hashLC.Data = append(hashLC.Data, cData.Hash)
	currRate.Text = strconv.FormatFloat(cData.Hash, 'f', 0, 64) + " H/s"
	switch {
	case cData.Hash <= 500:
		hashLC.LineColor = termui.ColorRed
		currRate.TextBgColor = termui.ColorRed
		currRate.TextFgColor = termui.ColorBlue
	case cData.Hash <= 2000:
		hashLC.LineColor = termui.ColorYellow
		currRate.TextBgColor = termui.ColorYellow
		currRate.TextFgColor = termui.ColorBlack
	default:
		hashLC.LineColor = termui.ColorGreen
		currRate.TextBgColor = termui.ColorGreen
		currRate.TextFgColor = termui.ColorRed
	}
	amountDue.Text = FormatValue(strconv.FormatFloat(cData.AmtDue, 'f', 0, 64), 11)
	amountPaid.Text = FormatValue(strconv.FormatFloat(cData.AmtPaid, 'f', 0, 64), 11)
	sharesRight.Percent = int(percent.PercentOf(cData.ValidShares, (cData.InvalidShares + cData.ValidShares)))
	if sharesRight.Percent >= 98 {
		sharesRight.BarColor = termui.ColorGreen
	} else {
		sharesRight.BarColor = termui.ColorRed
	}
	termui.Render(termui.Body)
}

func DynamicResize(millisecs time.Duration) {
	for {
		termui.Body.Width = termui.TermWidth() * 2
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
		time.Sleep(millisecs * time.Millisecond)
	}
}

func FormatValue(d string, place int) string {
	dd := len(d)
	//log.Println(d)
	switch {
	case dd <= place || dd == (place+1):
		for i := 0; i < ((place + 1) - dd); i++ {
			d = "0" + d
		}
		d = "0." + d
	case dd > place:
		//log.Println("here")
		d = d[:dd-(place+1)] + "." + d[dd-(place+1):]
	}
	return d
}

func Usage() {
	fmt.Println(`Usage: 
	--------------------------------------------
	CheckStats <Monero Address> <Pool #> 
	--------------------------------------------

	Pool Numbers: 
	--------------------------------------------
		1. supportxmr
		2. xmrpool
		3. viaxmr
		4. hashvault
		5. moriaxmr
		6. moneroocean
	--------------------------------------------
	
	Example: 
	--------------------------------------------
		./CheckStats "4hx5kldf..." 2
	--------------------------------------------
	`)
	os.Exit(0)
}

//The magic happens here
func main() {
	if len(os.Args) < 3 {
		Usage()
	} else if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()
	c := colly.NewCollector()
	SettingsColl(c)
	settingsTerm()
	var addressURL string
	if value, err := strconv.Atoi(os.Args[2]); err == nil {
		addressURL = SetAddress(string(os.Args[1]), value)
	} else {
		fmt.Println("Couldn't parse pool number!")
		os.Exit(-1)
	}
	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})
	go GetStats(c, addressURL)
	go DynamicResize(250)
	termui.Loop()
}
