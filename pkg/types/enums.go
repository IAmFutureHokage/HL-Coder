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

type PrecipitationInterval byte

const (
	Less1 PrecipitationInterval = iota
	From1To3
	From3To6
	From6To12
	More12
)

func (ct PrecipitationInterval) ToByte() byte {
	return byte(ct)
}

func (ct *PrecipitationInterval) FromByte(b byte) {
	*ct = PrecipitationInterval(b)
}
