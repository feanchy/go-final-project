package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	data, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", errors.New("invalid dstart")
	}

	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 || len(parts) > 2 {
		return "", errors.New("invalid repeat format")
	}

	switch parts[0] {
	case "d":

		if len(parts) != 2 {
			return "", errors.New("invalid repeat")
		}

		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("error converting to a number")
		}

		if days <= 0 || days > 400 {
			return "", errors.New("invalid count days")
		}

		for {
			data = data.AddDate(0, 0, days)
			if data.After(now) {
				break
			}
		}

	case "y":
		for {
			data = data.AddDate(1, 0, 0)

			if data.After(now) {
				break
			}
		}

	default:
		return "", errors.New("nothing fits")
	}

	//---
	return data.Format(dateFormat), nil
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")

	var nowTime time.Time
	var err error

	if nowStr == "" {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	start := r.FormValue("date")
	repeat := r.FormValue("repeat")

	date, err := NextDate(nowTime, start, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(date))

}
