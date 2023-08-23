package value_object

import "time"

type Sks struct {
	sks uint
}

func (sks *Sks) NewSks(datasks uint) *Sks {
	sks.sks = datasks
	return sks
}

func (sks *Sks) NewSksTime() time.Duration {
	skstime := 45 * time.Minute
	return time.Duration(sks.sks * uint(skstime))
}
