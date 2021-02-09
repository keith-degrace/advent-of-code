package main

import (
	"au"
	"fmt"
	"regexp"
	"sort"
	"time"
)

func testInputs() ([]string) {
	return []string {
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}
}

type GuardSleepLog struct {
	id int
	minuteStats [60]int
}

func (this * GuardSleepLog) getSleepTotal() (int) {
	total := 0
	for _, minuteStat := range this.minuteStats {
		total += minuteStat
	}
	return total
}

func (this * GuardSleepLog) getLaziestMinute() (int) {
	var laziestMinute int
	var laziestMinuteStat int

	for minute, minuteStat := range this.minuteStats {
		if (minuteStat > laziestMinuteStat) {
			laziestMinute = minute
			laziestMinuteStat = minuteStat
		}
	}

	return laziestMinute
}

func parseGuardLogs(inputs []string) (map [int] GuardSleepLog) {
	sort.Strings(inputs)

	inputRegex := regexp.MustCompile("\\[(.*)] (.*)")
	shiftBeginsRegex := regexp.MustCompile("Guard \\#([0-9]*)")

	logs := map [int] GuardSleepLog {}

	var currentGuardId int
	var sleepStartMinute int

	for _, input := range inputs {
		inputRegexMatches := inputRegex.FindStringSubmatch(input)

		eventTime, err := time.Parse("2006-01-02 15:04",  inputRegexMatches[1])
		au.FatalOnError(err)

		event := inputRegexMatches[2]

		shiftBeginsMatches := shiftBeginsRegex.FindStringSubmatch(event)
		if (len(shiftBeginsMatches) > 0) {
			currentGuardId = au.ToNumber(shiftBeginsMatches[1])
			continue
		}

		if event == "falls asleep" {
			sleepStartMinute = eventTime.Minute()
		}

		if event == "wakes up" {
			sleepEndMinute := eventTime.Minute()

			log := logs[currentGuardId]

			log.id = currentGuardId
			for minute := sleepStartMinute; minute < sleepEndMinute; minute++ {
				log.minuteStats[minute]++
			}

			logs[currentGuardId] = log
		}
	}

	return logs
}

func getLaziestGuard(logs map [int] GuardSleepLog) (GuardSleepLog) {
	var laziestGuard GuardSleepLog
	var laziestSleepTime int

	for _, log := range logs {
		sleepTime := log.getSleepTotal()
		if sleepTime > laziestSleepTime {
			laziestGuard = log
			laziestSleepTime = sleepTime
		}
	}

	return laziestGuard
}

func main() {
	inputs := au.ReadInputAsStringArray("04")
	//inputs := testInputs()

  logs := parseGuardLogs(inputs)

	laziestGuard := getLaziestGuard(logs)

	fmt.Println(laziestGuard.id * laziestGuard.getLaziestMinute())
}
