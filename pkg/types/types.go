package types

type PostCode struct {
	PostCode string
}

type DateAndTime struct {
	Date byte
	Time byte
}

type WaterLevelOnTime struct {
	DateAndTime
	WaterLevel uint16
}

type DeltaWaterLevel struct {
	Delta int16
}

type WaterLevelOn20h struct {
	WaterLevel uint16
}

type Temperature struct {
	WaterTemperature float32
	AirTemperature   int8
}

type Phenomenia struct {
	Phenomen    byte
	isUntensity bool
	Intensity   byte
}

type IcePhenomenia struct {
	State       byte
	Phenomenias []Phenomenia
}

type Reservoir struct {
}

type FullTelegram struct {
	PostCode
	IsDangerous bool
	WaterLevelOnTime
	DeltaWaterLevel
	WaterLevelOn20h
	Temperature     Temperature
	IsIcePhenomenia bool
	IcePhenomenia   IcePhenomenia
	IsReservoir     bool
	Reservoir       Reservoir
}
