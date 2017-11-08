package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/leekchan/timeutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var opts struct {
		Format      string `short:"F" long:"format" description:"time format" default:"%Y-%m-%d"`
		Placeholder string `short:"I" long:"replace" description:"placeholder" default:"{}"`
		Wday        int    `short:"w" long:"wday" optional:"true" optional-value:"1" default:"-1"`
		Mday        int    `short:"m" long:"mday" optional:"true" optional-value:"1" default:"0"`
		Args        struct {
			StartDate string   `positional-arg-name:"START_DATE"`
			EndDate   string   `positional-arg-name:"END_DATE"`
			Commands  []string `positional-arg-name:"[COMMAND]"`
		} `positional-args:"yes" required:"2"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	start, err := time.Parse("2006-01-02", opts.Args.StartDate)
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid date format:", opts.Args.StartDate)
		os.Exit(1)
	}

	end, err := time.Parse("2006-01-02", opts.Args.EndDate)
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid date format:", opts.Args.EndDate)
		os.Exit(1)
	}

	op := buildOperator(opts.Args.Commands, opts.Format, opts.Placeholder)
	op = buildFilterOperator(op, opts.Wday, opts.Mday)
	iterDates(start, end, op)
}

func iterDates(start time.Time, end time.Time, callback func(time.Time)) {
	if start.Before(end) {
		t := start
		for !t.After(end) {
			callback(t)
			t = t.AddDate(0, 0, 1)
		}
	} else {
		t := start
		for !t.Before(end) {
			callback(t)
			t = t.AddDate(0, 0, -1)
		}
	}
}

func buildOperator(commands []string, format string, placeholder string) func(time.Time) {
	if len(commands) == 0 {
		return func(t time.Time) {
			fmt.Println(timeutil.Strftime(&t, format))
		}
	} else {
		cmds := mapString(commands, func(v string) string {
			return strings.Replace(v, placeholder, format, -1)
		})
		return func(t time.Time) {
			cs := mapString(cmds, func(v string) string {
				return timeutil.Strftime(&t, v)
			})
			c := exec.Command(cs[0], cs[1:]...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			e := c.Run()
			if e != nil {
				fmt.Fprintln(os.Stderr, e)
				os.Exit(1)
			}
		}
	}
}

func buildFilterOperator(op func(time.Time), wday int, mday int) func(time.Time) {
	var filter func(time.Time) bool
	if wday >= 0 {
		filter = func(t time.Time) bool {
			return t.Weekday() == time.Weekday(wday)
		}
	} else if mday > 0 {
		filter = func(t time.Time) bool {
			return t.Day() == mday
		}
	} else if mday < 0 {
		filter = func(t time.Time) bool {
			return t.AddDate(0, 0, -mday).Day() == 1
		}
	} else {
		filter = func(_ time.Time) bool {
			return true
		}
	}

	return func(t time.Time) {
		if filter(t) {
			op(t)
		}
	}
}

func mapString(vs []string, f func(string) string) []string {
	ret := make([]string, len(vs))
	for i, v := range vs {
		ret[i] = f(v)
	}
	return ret
}
