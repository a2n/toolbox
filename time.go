package deg_

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/beevik/ntp"
)

type TimeService struct {
	tz   *time.Location
	diff *int64
}

func NewTimeService() *TimeService {
	tz, e := time.LoadLocation("Asia/Taipei")
	if e != nil {
		log.Fatalf("%+v", e)
	}
	s := &TimeService{tz: tz}
	go s.loop()
	return s
}

func (this *TimeService) loop() {
	this.sync()
	for _ = range time.Tick(time.Minute) {
		this.sync()
	}
}

func (this *TimeService) sync() {
	t, e := ntp.Time("time.google.com")
	if e != nil {
		log.Printf("%+v", e)
		return
	}
	atomic.StoreInt64(this.diff, int64(t.Sub(time.Now())))
}

func (this *TimeService) Now() time.Time {
	d := time.Duration(atomic.LoadInt64(this.diff))
	return time.Now().In(this.tz).Add(d)
}
