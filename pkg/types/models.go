package types

type Reservoir struct {
	HeadwaterLevel        HeadwaterLevel
	AverageReservoirLevel AverageReservoirLevel
	DownstreamLevel       DownstreamLevel
	ReservoirVolume       ReservoirVolume
	IsReservoirWaterInflow
	ReservoirWaterInflow ReservoirWaterInflow
}

type ReservoirWaterInflow struct {
	Inflow Inflow
	Reset  Reset
}

type Telegram struct {
	PostCode
	DateAndTime
	IsDangerous        IsDangerous
	WaterLevelOnTime   WaterLevelOnTime
	DeltaWaterLevel    DeltaWaterLevel
	WaterLevelOn20h    WaterLevelOn20h
	Temperature        Temperature
	IcePhenomeniaState IcePhenomeniaState
	IcePhenomenia      []Phenomenia
	Precipitation      Precipitation
	IsReservoir        IsReservoir
	Reservoir          Reservoir
}
