package letter

type Request interface {
    Header() map[string]string
    Body() interface{}
}