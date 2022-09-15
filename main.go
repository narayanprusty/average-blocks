package main

import (
	"github.com/narayanprusty/average-blocks/api"
	"github.com/narayanprusty/average-blocks/tracker"
)

func main() {
	go tracker.RunCronJobs()
	api.StartAPIServer()
}
