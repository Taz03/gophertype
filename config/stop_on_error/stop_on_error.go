package stoponerror

type StopOnError string

const (
    Off StopOnError = "off"
    Letter StopOnError = "letter"
    Word StopOnError = "word"
)
