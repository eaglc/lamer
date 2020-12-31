package session

type Cache interface {
    // return all session in the cache
    All() ([]Session, error)

    // return multiple session in a category
    Category(category interface{}) ([]Session, error)

    // return multiple session by multiple indexes
    Indexes(indexes ...interface{}) ([]Session, error)

    // return one session by one index
    Index(index interface{}) (Session, error)

    // return multiple session by multiple remote addrs
    Addrs(addrs ...string) ([]Session, error)

    // return one session by one remote addr
    Addr(addr string) (Session, error)
}