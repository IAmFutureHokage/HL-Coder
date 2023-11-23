package types

type IcePhenomeniaState byte

const (
	None IcePhenomeniaState = iota
	True
	Empty
)

func (ct IcePhenomeniaState) ToByte() byte {
	return byte(ct)
}

func (ct *IcePhenomeniaState) FromByte(b byte) {
	*ct = IcePhenomeniaState(b)
}

type PrecipitationDuration byte

const (
	Less1 PrecipitationDuration = iota
	From1To3
	From3To6
	From6To12
	More12
)

func (ct PrecipitationDuration) ToByte() byte {
	return byte(ct)
}

func (ct *PrecipitationDuration) FromByte(b byte) {
	*ct = PrecipitationDuration(b)
}
