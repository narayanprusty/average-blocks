package tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/narayanprusty/average-blocks/config"
)

type Chainhead struct {
	HeadSlot string `mapstructure:"headSlot"`
}

var AverageBlocksPerMinute int = 0
var lastSlotNumber int = 0
var iterations int = 0
var totalSlots int = 0

func trackBlockHeight() {
	resp, err := http.Get(config.Config.BeaconURL + "/eth/v1alpha1/beacon/chainhead")
	if err != nil {
		log.Println(err.Error())
	}

	if resp.StatusCode != 200 {
		log.Println("got status code " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	chainHead := Chainhead{}
	json.Unmarshal([]byte(string(body)), &chainHead)

	headSlot, err := strconv.Atoi(chainHead.HeadSlot)
	if err != nil {
		log.Println(err)
	}

	if lastSlotNumber != 0 {
		totalSlots = totalSlots + (headSlot - lastSlotNumber)
		iterations++
		AverageBlocksPerMinute = totalSlots / iterations
		fmt.Println("set average rate to: " + strconv.Itoa(AverageBlocksPerMinute))
	}

	lastSlotNumber = headSlot
}

func GetRate() int {
	return AverageBlocksPerMinute
}

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(trackBlockHeight)
	s.StartBlocking()
}
