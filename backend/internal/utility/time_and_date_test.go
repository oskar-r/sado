package utility

import (
	"testing"
	"time"
)

func TestInFuture(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{t: "2019-01-01 21:21:21"},
			true,
		},
		{
			"test2",
			args{t: "2016-01-01 21:21:21"},
			false,
		},
		{
			"test3",
			args{t: "2018-03-31 21:21:21"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InFuture(tt.args.t); got != tt.want {
				t.Errorf("InFuture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTime(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"TEST1",
			args{0, 0, 0},
			CurrentTime(),
		},
		{
			"TEST2",
			args{1, 0, 0},
			time.Now().AddDate(1, 0, 0).Format("2006-01-02 15:04:05"),
		},
		{
			"TEST3",
			args{-1, 0, 0},
			time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTime(tt.args.year, tt.args.month, tt.args.day); got != tt.want {
				t.Errorf("AddTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetweenHours(t *testing.T) {
	type args struct {
		startH int
		endH   int
		t      time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{2, 3, time.Date(2018, 04, 23, 23, 0, 0, 0, time.UTC)}, // 2018-04-24 01:00:00
			false,
		},
		{
			"Test 2",
			args{2, 3, time.Date(2018, 04, 23, 0, 0, 1, 0, time.UTC)}, // 2018-04-24 01:00:01
			true,
		},
		{
			"Test 3",
			args{23, 0, time.Date(2018, 04, 01, 2, 0, 0, 0, time.UTC)}, //2018-03-31 04:00:00
			false,
		},
		{
			"Test 4",
			args{23, 0, time.Date(2018, 04, 23, 21, 0, 1, 0, time.UTC)}, //2018-04-23 23:00:01
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BetweenHours(tt.args.startH, tt.args.endH, tt.args.t); got != tt.want {
				t.Errorf("BetweenHours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrentTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"TEST1",
			time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrentTime(); got != tt.want {
				t.Errorf("CurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareDates(t *testing.T) {
	type args struct {
		t1 string
		t2 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"TEST 1",
			args{"2018-01-02 00:00:00", "2018-01-01 00:00:00"},
			true,
			false,
		},
		{
			"TEST 2",
			args{"2018-01-01 00:00:00", "2018-01-02 00:00:00"},
			false,
			false,
		},
		{
			"TEST 3",
			args{"2018-01-01 01:00:02", "2018-01-01 01:00:01"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareDates(tt.args.t1, tt.args.t2)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
