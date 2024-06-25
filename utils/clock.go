package utils

import "time"

var Clock clock = NewRealTimeClock()

type clock interface {
	Now() time.Time // 現在時刻を返す
}

// RealTimeClock 通常のアプリケーションで使う時計
type RealTimeClock struct{}

func NewRealTimeClock() *RealTimeClock {
	return &RealTimeClock{}
}

func (c *RealTimeClock) Now() time.Time {
	return time.Now().Local()
}

func (c *RealTimeClock) Freeze() {
	// RealTimeClockでは何も行わない
}

// DummyClock テスト時に使う時計
type DummyClock struct {
	now *time.Time
}

func NewDummyClock(frozenTime time.Time) *DummyClock {
	return &DummyClock{
		now: &frozenTime,
	}
}

func (d *DummyClock) Now() time.Time {
	if d.now == nil {
		return time.Now().Local()
	} else {
		return (*d.now).Local()
	}
}
