package webframework

type Settings struct {
	Addr       string
	Urls       []Url
	Middleware []Middleware
}
