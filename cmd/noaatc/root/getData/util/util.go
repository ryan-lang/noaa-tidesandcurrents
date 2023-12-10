package util

import (
	"log"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi"
)

func ParseDateParam(beginDate, endDate, rangeHours, relative string) dataApi.DateParam {

	if beginDate != "" && endDate != "" {
		return &dataApi.DateParamBeginAndEnd{
			BeginDate: parseDate(beginDate),
			EndDate:   parseDate(endDate),
		}
	} else if beginDate != "" && rangeHours != "" {
		return &dataApi.DateParamBeginAndRange{
			BeginDate:  parseDate(beginDate),
			RangeHours: parseInt(rangeHours),
		}
	} else if endDate != "" && rangeHours != "" {
		return &dataApi.DateParamEndAndRange{
			EndDate:    parseDate(endDate),
			RangeHours: parseInt(rangeHours),
		}
	} else if relative != "" {
		return &dataApi.DateRelative{
			Relative: dataApi.RelativeDateOpt(relative),
		}
	}

	return nil
}

func ParseIntervalParam(interval string) dataApi.IntervalParam {
	if interval != "" {
		return dataApi.IntervalParam(interval)
	}
	return ""
}

func parseInt(s string) int32 {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to parse int: %v", err)
	}
	return int32(i)
}

func parseDate(date string) time.Time {

	// setup the parser
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	if date != "" {
		d := whenParseTry(w, date)
		if !d.IsZero() {
			return d
		}

		return dateParseFatal(date)
	}

	return time.Time{}
}

func whenParseTry(w *when.Parser, s string) time.Time {
	parsed, err := w.Parse(s, time.Now())
	if err != nil {
		return time.Time{}
	} else if parsed == nil {
		return time.Time{}
	}
	return parsed.Time
}

func dateParseFatal(s string) time.Time {
	d, err := dateparse.ParseLocal(s)
	if err != nil {
		log.Fatalf("Failed to parse date: %v", err)
	}
	return d
}
