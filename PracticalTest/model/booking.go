package model

type Booking struct {
    OfficeName      string           `json:"officeName"`
    RoomName        string           `json:"roomName"`
    Participants     int              `json:"participants"`
    ListConsumption []Consumption    `json:"listConsumption"`
}

type Consumption struct {
    Name string `json:"name"`
}
