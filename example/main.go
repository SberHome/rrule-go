package main

import (
	"fmt"
	"time"

	"github.com/sberhome/rrule-go"
)

func exampleRRule() {
	fmt.Println("Daily, for 10 occurrences.")
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   10,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	fmt.Println(r.All())
	// [1997-09-02 09:00:00 +0000 UTC
	//  1997-09-03 09:00:00 +0000 UTC
	//  1997-09-04 09:00:00 +0000 UTC
	//  ...
	//  1997-09-11 09:00:00 +0000 UTC]

	fmt.Println(r.Between(
		time.Date(1997, 9, 6, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 8, 0, 0, 0, 0, time.UTC), true))
	// [1997-09-06 09:00:00 +0000 UTC
	//  1997-09-07 09:00:00 +0000 UTC]

	fmt.Println(r)
	// DTSTART:19970902T090000Z
	// FREQ=DAILY;COUNT=10

	fmt.Println("\nEvery four years, the first Tuesday after a Monday in November, 3 occurrences (U.S. Presidential Election day).")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.YEARLY,
		Interval:   4,
		Count:      3,
		Bymonth:    []int{11},
		Byweekday:  []rrule.Weekday{rrule.TU},
		Bymonthday: []int{2, 3, 4, 5, 6, 7, 8},
		Dtstart:    time.Date(1996, 11, 5, 9, 0, 0, 0, time.UTC)})
	fmt.Println(r.All())
	// [1996-11-05 09:00:00 +0000 UTC
	//  2000-11-07 09:00:00 +0000 UTC
	//  2004-11-02 09:00:00 +0000 UTC]
}

func exampleRRuleSet() {
	fmt.Println("\nDaily, for 7 days, jumping Saturday and Sunday occurrences.")
	set := rrule.Set{}
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	fmt.Println(set.All())
	// [1997-09-02 09:00:00 +0000 UTC
	//  1997-09-03 09:00:00 +0000 UTC
	//  1997-09-04 09:00:00 +0000 UTC
	//  1997-09-05 09:00:00 +0000 UTC
	//  1997-09-06 09:00:00 +0000 UTC
	//  1997-09-07 09:00:00 +0000 UTC
	//  1997-09-08 09:00:00 +0000 UTC]

	fmt.Println("\nWeekly, for 4 weeks, plus one time on day 7, and not on day 16.")
	set = rrule.Set{}
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.WEEKLY,
		Count:   4,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.RDate(time.Date(1997, 9, 7, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC))
	fmt.Println(set.All())
	// [1997-09-02 09:00:00 +0000 UTC
	//  1997-09-07 09:00:00 +0000 UTC
	//  1997-09-09 09:00:00 +0000 UTC
	//  1997-09-23 09:00:00 +0000 UTC]
}

func exampleStrToRRule() {
	fmt.Println()
	r, _ := rrule.StrToRRule("FREQ=DAILY;INTERVAL=10;COUNT=5")
	fmt.Println(r.All())
	// [2020-10-12 02:26:35 +0000 UTC
	//  2020-10-22 02:26:35 +0000 UTC
	//  2020-11-01 02:26:35 +0000 UTC
	//  2020-11-11 02:26:35 +0000 UTC
	//  2020-11-21 02:26:35 +0000 UTC]
}

func exampleStrToRRuleSet() {
	s, err := rrule.StrToRRuleSet(`DTSTART;TZID=Europe/Moscow:20210901T130255
RRULE:FREQ=DAILY;INTERVAL=10;COUNT=5;X-FILTER-DAY=ANY;X-COLOR=#00ffa
RDATE:20060102T150405Z`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s.All())
	// [2006-01-02 15:04:05 +0000 UTC
	//  2020-10-12 02:26:35 +0000 UTC
	//  2020-10-22 02:26:35 +0000 UTC
	//  2020-11-01 02:26:35 +0000 UTC
	//  2020-11-11 02:26:35 +0000 UTC
	//  2020-11-21 02:26:35 +0000 UTC]

	rruleField := s.GetRRule()
	if rruleField == nil {
		fmt.Println("rrule is empty")
		return
	}
	attrs := rruleField.GetCustomAttributes()
	for key, val := range attrs {
		fmt.Printf("%s = %s\n", key, val)
	}
	// X-FILTER-DAY = ANY
	// X-COLOR = #00ffa
}

func main() {
	// exampleRRule()
	// exampleRRuleSet()
	// exampleStrToRRule()
	exampleStrToRRuleSet()
}
