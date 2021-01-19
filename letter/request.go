package letter

type Request interface {
    Body() interface{}

    // addr or other. defined by yourself
    Endpoint() string
    // endpoint name
    Name() string
    // endpoint id
    Id() string
}