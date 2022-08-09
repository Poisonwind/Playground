package electronic

/// INTERFACE ////

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonCount() int
}

type Smartphone interface {
	OS() string
}


//// STRUCTS + CONSTRUCTORS ////

type ApplePhone struct{
	model string
	os string
}

func NewApplePhone(m, os string) *ApplePhone {
	return &ApplePhone{m, os}
}

type AndroidPhone struct{
	brand string
	model string
	os string
}

func NewAndroidPhone(b, m, os string) *AndroidPhone {
	return &AndroidPhone{b, m, os}
}

type RadioPhone struct{
	brand string
	model string
	buttonCount int
}

func NewRadioPhone(b, m string, bc int) *RadioPhone {
	return &RadioPhone{b, m, bc}
}

//// METHODS ////

	// applePhone //

func (ap *ApplePhone) Brand() string {
	return "apple"
}

func (ap *ApplePhone) Model() string {
	return ap.model
}

func (ap *ApplePhone) Type() string {
	return "smartphone"
}

func (ap *ApplePhone) OS() string {
	return ap.os
}

	// androidPhone //

func (an *AndroidPhone) Brand() string {
	return an.brand
}

func (an *AndroidPhone) Model() string {
	return an.model
}

func (an *AndroidPhone) Type() string {
	return "smartphone"
}

func (an *AndroidPhone) OS() string {
	return an.os
}


	// radioPhone //


func (rp *RadioPhone) Brand() string {
	return rp.brand
}

func (rp *RadioPhone) Model() string {
	return rp.model
}

func (rp *RadioPhone) Type() string {
	return "station"
}

func (rp *RadioPhone) ButtonCount() int {
	return rp.buttonCount
}

