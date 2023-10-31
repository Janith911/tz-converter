package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func getTimeZones(timeZone string) string {
	timeZone = strings.ToUpper(timeZone)
	var timeGap string
	switch timeZone {
	case "IST":
		timeGap = "+5:30"
	case "AEST":
		timeGap = "+10:00"
	case "AEDT":
		timeGap = "+11:00"
	case "MYT":
		timeGap = "+8:00"
	case "UTC":
		timeGap = "+0:00"
	case "EST":
		timeGap = "-5:00"
	case "EDT":
		timeGap = "-4:00"
	default:
		fmt.Printf(Red+"Invalid Timezone : %v\n"+Reset, timeZone)
		os.Exit(0)
	}

	return timeGap
}

func prefferedTime(targetTimeGap int, fromTimeGap int, fromTime int) (string, string) {
	var targetTime int = fromTime + (targetTimeGap - fromTimeGap)
	var prefferedTime1 string
	var day string
	if targetTime >= 0 && targetTime <= 1440 {
		hours := targetTime / 60
		minutes := targetTime % 60
		formattedHours := fmt.Sprintf("%02d", hours)
		formattedMinutes := fmt.Sprintf("%02d", minutes)
		prefferedTime1 = formattedHours + ":" + formattedMinutes
		day = "Same Day"
	} else if targetTime < 0 {
		targetTime = 1440 + targetTime
		hours := targetTime / 60
		minutes := targetTime % 60
		formattedHours := fmt.Sprintf("%02d", hours)
		formattedMinutes := fmt.Sprintf("%02d", minutes)
		prefferedTime1 = formattedHours + ":" + formattedMinutes
		day = "Previous Day"
	} else if targetTime > 1440 {
		targetTime = targetTime - 1440
		hours := targetTime / 60
		minutes := targetTime % 60
		formattedHours := fmt.Sprintf("%02d", hours)
		formattedMinutes := fmt.Sprintf("%02d", minutes)
		prefferedTime1 = formattedHours + ":" + formattedMinutes
		day = "Next Day"
	}
	return prefferedTime1, day
}

func splitTime(time_str string) int {
	test := strings.Split(time_str, ":")
	hours, _ := strconv.Atoi(test[0])
	minutes, _ := strconv.Atoi(test[1])

	inMinutes := hours*60 + minutes

	return inMinutes
}

func timenow() (string, string) {
	now_time_zone, _ := time.Now().Zone()
	now_time_zone = now_time_zone + "00"
	time_gap := now_time_zone[0:3] + ":" + now_time_zone[3:5]

	now_time_hour, now_time_min, _ := time.Now().Clock()
	now_time := fmt.Sprintf("%02d", now_time_hour) + ":" + fmt.Sprintf("%02d", now_time_min)
	return time_gap, now_time
}

func main() {
	cliArgs := os.Args
	if len(cliArgs) == 4 {
		targetGapMinutes := splitTime(getTimeZones(cliArgs[3]))
		fromGapMinutes := splitTime(getTimeZones(cliArgs[2]))
		fromTimeMinutes := splitTime(cliArgs[1])
		calculatedTime, tartget_day := prefferedTime(targetGapMinutes, fromGapMinutes, fromTimeMinutes)
		fmt.Printf(White+"%v Time \t: %v\n", strings.ToUpper(cliArgs[3]), calculatedTime)
		fmt.Printf(White+"Day \t\t: %v\n", tartget_day)
	} else if len(cliArgs) == 3 && cliArgs[1] == "now" {
		local_timegap, time_now := timenow()
		targetGapMinutes := splitTime(getTimeZones(cliArgs[2]))
		fromGapMinutes := splitTime(local_timegap)
		fromTimeMinutes := splitTime(time_now)
		calculatedTime, tartget_day := prefferedTime(targetGapMinutes, fromGapMinutes, fromTimeMinutes)
		t := time.Now()
		date := fmt.Sprintf("%d-%02d-%02d", t.Day(), t.Month(), t.Year())
		fmt.Printf(White+"Local Time \t: %v (%v)\n", time_now, date)
		fmt.Printf(Yellow+"%v Time \t: %v\n"+Reset, strings.ToUpper(cliArgs[2]), calculatedTime)
		fmt.Printf(White+"Day \t\t: %v\n", tartget_day)
	} else {
		fmt.Printf(Red + "No required arguments !\n" + Reset)
	}
}
