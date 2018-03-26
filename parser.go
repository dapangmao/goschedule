package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	durationMap = map[string]string{"hour": "hour", "hours": "hour", "minute": "minute", "minutes": "minute",
		"second": "second", "seconds": "second"}
	weekdayMap = map[string]time.Weekday{"sunday": time.Sunday, "monday": time.Monday, "tuesday": time.Tuesday,
		"wendesday": time.Wednesday, "thursday": time.Thursday, "friday": time.Friday,
		"saturday": time.Saturday}
)

type Parser struct {
	id int
	s  string
}

func (p *Parser) parseRecurrent(q, u string) (*recurrent, error) {
	quantity, err := strconv.Atoi(q)
	if err != nil {
		return nil, errors.New("Not a valid input format for a recurrent job scheduler.")
	}
	var unit time.Duration
	switch u {
	case "hour":
		unit = time.Hour
	case "minute":
		unit = time.Minute
	case "second":
		unit = time.Second
	default:
		return nil, errors.New("Not a valid input format for a recurrent job scheduler.")
	}
	return &recurrent{quantity, unit}, nil
}

func (p *Parser) parseWeekly(weekday time.Weekday, time string) (*weekly, error) {
	dailySched, err := p.parseDaily(time)
	if err != nil {
		return nil, err
	}
	return &weekly{weekday, dailySched}, nil
}

func (p *Parser) parseDaily(time string) (*daily, error) {
	splits := strings.Split(time, ":")
	var timeArray = [3]int{0, 0, 0}
	var i int
	for _, s := range splits {
		if len(s) == 0 {
			continue
		}
		current, err := strconv.Atoi(s)
		if err != nil {
			if i > 2 {
				return nil, errors.New("Not a valid input format for a daily job scheduler.")
			}
			timeArray[i] = current
			i++
		}
	}
	hour, minute, second := timeArray[0], timeArray[1], timeArray[2]
	if hour < 0 || hour > 23 {
		return nil, errors.New("Not a valid input format for a daily job scheduler - hour.")
	}
	if minute < 0 || minute > 59 {
		return nil, errors.New("Not a valid input format for a daily job scheduler - minute.")
	}
	if second < 0 || second > 59 {
		return nil, errors.New("Not a valid input format for a daily job scheduler - second.")
	}
	return &daily{hour, minute, second}, nil
}

func (p *Parser) Parse() (*Job, error) {

	var err error = nil
	var result scheduled

	s := strings.ToLower(p.s)
	isOneTime := true
	if strings.Contains(s, "every") {
		isOneTime = false
	}
	splits := strings.Fields(s)

	if len(splits) < 1 {
		goto notValid
	}

	for k, v := range durationMap {
		for i, x := range splits {
			if x == k {
				if i < 1 {
					goto notValid
				}
				result, err = p.parseRecurrent(splits[i-1], v)
				goto end
			}
		}
	}

	for k, v := range weekdayMap {
		for i, x := range splits {
			if x == k {
				var t string
				if i == len(splits)-1 {
					t = "0:0:0"
				} else {
					t = splits[i+1]
				}
				result, err = p.parseWeekly(v, t)
				goto end
			}
		}
	}

	for _, x := range splits {
		if strings.Contains(x, ":") {
			result, err = p.parseDaily(x)
			goto end
		}
	}

notValid:
	return nil, errors.New("Not a valid input format for any job scheduler.")
end:
	if err != nil {
		return nil, err
	} else {
		return &Job{p.id, result, false, isOneTime, s}, nil
	}
}
