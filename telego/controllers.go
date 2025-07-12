package telego

type Handler struct {
	Condition func(template string) bool
	Action Action
}