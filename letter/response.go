package letter

type ResponseWriter interface {
    Write(m interface{}) error
}