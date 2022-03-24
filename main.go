package main

func main() {
	c := NewHttpClient()
	h := NewIoHandler()
	cfg := NewConfig(c, h)
	cfg.setup()
}
