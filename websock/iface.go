package websock

type WriterIf interface {
    Send(m MethodRequest) error
    Read() <-chan []byte
    Stop()
    Run() error
}
