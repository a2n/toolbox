package toolbox

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/beevik/ntp"
)

// TimeService The time service
type TimeService struct {
	tz   *time.Location
	diff int64
}

// NewTimeService Creating
func NewTimeService() *TimeService {
	tz, e := time.LoadLocation("Asia/Taipei")
	if e != nil {
		log.Fatalf("%+v", e)
	}
	s := &TimeService{tz: tz}
	s.sync()
	go s.loop()
	return s
}

// loop synchronization loop
func (ts *TimeService) loop() {
	for range time.Tick(time.Minute) {
		ts.sync()
	}
}

// sync synchronization
func (ts *TimeService) sync() {
	t, e := ntp.Time("time.google.com")
	if e != nil {
		log.Printf("%+v", e)
		return
	}
	atomic.StoreInt64(&ts.diff, int64(t.Sub(time.Now())))
}

// Now Getting now time
func (ts *TimeService) Now() time.Time {
	d := time.Duration(atomic.LoadInt64(&ts.diff))
	return time.Now().In(ts.tz).Add(d)
}
