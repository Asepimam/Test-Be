package model

type Summary struct {
    OfficeName     string       `json:"office_name"`
    RoomName       string       `json:"room_name"`
    Capacity       int          `json:"capacity"`
    UsagePercent   float64      `json:"usage_percent"`
    ConsumeNominal float64      `json:"consume_nominal"`
    TypeConsume    []TypeDetail `json:"type_consume"`
}

type TypeDetail struct {
    Name   string  `json:"name"`
    Count  int     `json:"count"`
    Amount float64 `json:"amount"`
}


