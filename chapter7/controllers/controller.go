package controllers

// Controller passes state to our handlers
type Controller struct {
	storage Storage
}

// New is a Controller 'constructor'
func New(storage Storage) *Controller {
	c := Controller{
		storage: storage,
	}
	return &c
}

// Payload is our common response
type Payload struct {
	Value string `json:"value"`
}
