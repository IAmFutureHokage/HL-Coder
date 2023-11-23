package types

type Reservoir struct {
	HeadwaterLevel        HeadwaterLevel
	AverageReservoirLevel AverageReservoirLevel
	DownstreamLevel       DownstreamLevel
	ReservoirVolume       ReservoirVolume
	IsReservoirWaterInflow
	ReservoirWaterInflow ReservoirWaterInflow
	//1. Уровень верхнего бьефа 19543 (1 - группа, 2-5 значение)
	//2. Уровень среднего водохранилища 29555 ( 1 - группа, 2-5 значение)
	//3. Уровень нижнего бьефа 40308 ( 1-2 группа, 3-5 уровень)
	//4. Обем водохранилища 75671 (1-2 группа, 3-5 значение * 100)
	//5  955(дата)
}

type ReservoirWaterInflow struct {
	ObPs  ObPs
	Reset Reset
	//1. Об.пс 43390 (43 - группа, 390 - значение)
	//2. Сброс 73759 (73 - число группа, 759 - значение)
}

type Telegram struct {
	PostCode
	IsDangerous        bool
	DateAndTime        DateAndTime
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

type FullTelegram struct {
	Telegram []Telegram
}
