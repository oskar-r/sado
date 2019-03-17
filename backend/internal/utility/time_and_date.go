package utility

import (
	"log"
	"time"
)

func InFuture(t string) bool {
	start, _ := time.Parse("2006-01-02 15:04:05", t)
	end := time.Now()
	return start.After(end)
}

/*CompareDates compares two dates and returns true of date1 is more recent than date2 or false if date 2 is more recent*/
func CompareDates(t1 string, t2 string) (bool, error) {
	first, err := time.Parse("2006-01-02 15:04:05", t1)
	if err != nil {
		log.Println(err)
		return false, err
	}
	second, err := time.Parse("2006-01-02 15:04:05", t2)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return first.After(second), nil
}
func CurrentTime() string {
	loc, _ := time.LoadLocation("Europe/Stockholm")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func FormatTime(t time.Time) string {
	loc, _ := time.LoadLocation("Europe/Stockholm")
	return t.In(loc).Format("2006-01-02 15:04:05")
}

func CurrentTimeUnix() int64 {
	return time.Now().Unix()
}

func AddTime(year int, month int, day int) string {
	return time.Now().AddDate(year, month, day).Format("2006-01-02 15:04:05")
}

func BetweenHours(startH int, endH int, t time.Time) bool {
	loc, _ := time.LoadLocation("Europe/Stockholm")
	tt := t.In(loc)

	if startH == endH {
		return false
	}
	if startH > 23 || endH > 23 {
		return false
	}
	if endH == 0 {
		endH = 24
	}

	//Cross dates
	if startH > endH {
		if tt.Hour() >= startH {
			return true
		}

		if tt.Hour() < endH {
			return true
		}
	}

	if tt.Hour() >= startH && tt.Hour() < endH {
		return true
	}
	return false
}

func UnixTimeToTimestamp(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}
