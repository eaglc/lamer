package deliver

type NewDeliver func() Deliver

// used fo service to service/service to user deliver message
type Deliver interface {

    Init(...Option)

    Options() Options

    // broadcast with a category
    Broadcast(m interface{}, category interface{}) error

    // multicast with multiple indexes
    Multicast(m interface{}, indexes ...interface{}) error

    // unicast with one index
    Unicast(m interface{}, index interface{}) error
}

