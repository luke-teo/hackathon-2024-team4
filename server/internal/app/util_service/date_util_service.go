package util_service

import "time"

type JSONDate string

func (jd JSONDate) Parse() time.Time {
	layout := "2006-01-02"

	t, err := time.Parse(layout, string(jd))
	if err != nil {
		panic(err)
	}

	return t
}
