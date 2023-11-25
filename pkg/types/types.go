package types

type PostCode string

type IsDangerous bool

type DateAndTime struct {
	Date byte
	Time byte
}

type WaterLevelOnTime int16

type DeltaWaterLevel int16

type WaterLevelOn20h uint16

type Temperature struct {
	WaterTemperature float32
	AirTemperature   int8
}

type Phenomenia struct {
	Phenomen    byte
	IsUntensity bool
	Intensity   *byte
}

type Precipitation struct {
	Value    float32
	Duration PrecipitationDuration
}

type IsReservoir struct {
	State bool
	Date  byte
}

type HeadwaterLevel uint32

type AverageReservoirLevel uint32

type DownstreamLevel uint32

type ReservoirVolume uint32

type IsReservoirWaterInflow struct {
	IsReservoirWaterInflow bool
	Date                   byte
}

type Inflow uint32

type Reset uint32

type PrevDay struct {
	IsNextDay bool
	Date      byte
}
