package services

type Data struct {
	UserID int
	Id     int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}
