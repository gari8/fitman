package main

func main() {
	c := NewHttpClient()
	cfg := NewConfig(c)
	cfg.setup()
}
