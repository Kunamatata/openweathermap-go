package openweathermapgo

//Unit of measurement
type Unit string

const (
	IMPERIAL Unit = "imperial"
	CELSIUS  Unit = "metric"
	STANDARD Unit = "standard"
)

//Language
type Language string

const (
	ENGLISH Language = "en"
	FRENCH  Language = "fr"
)
