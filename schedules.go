package gonavitia

type RouteScheduleResponse struct {
	RouteSchedules []*RouteSchedule `json:"route_schedules"`
	Pagination     *Pagination      `json:"pagination"`
	FeedPublishers []*FeedPublisher `json:"feed_publishers"`
	Error          *Error           `json:"error,omitempty"`
	Exceptions     []struct{}       `json:"exceptions"`
}

type RouteSchedule struct {
	DisplayInfo    *PtDisplayInfo          `json:"display_informations"`
	Table          *Table                  `json:"table"`
	AdditionalInfo *string                 `json:"additional_informations"`
	GeoJson        *GeoJsonMultilineString `json:"geojson"`
	Links          []Link                  `json:"links"`
}

type Table struct {
	Headers []*Header `json:"headers"`
	Rows    []Row     `json:"rows"`
}

type Header struct {
	DisplayInfo     *PtDisplayInfo `json:"display_informations"`
	AdditionalInfos []string       `json:"additional_informations"`
	Links           []Link         `json:"links"`
}

type Row struct {
	StopPoint *StopPoint `json:"stop_point"`
	DateTimes []DateTime `json:"date_times"`
}

type DateTime struct {
	//We don't handle time only schedules from calendar
	DateTime       NavitiaDatetime `json:"date_time"`
	BaseDateTime   NavitiaDatetime `json:"base_date_time,omitempty"`
	AdditionalInfo []string        `json:"additional_informations"`
	DataFreshness  string          `json:"data_freshness"`
	Links          []Link          `json:"links"`
}
