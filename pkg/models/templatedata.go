package models

// TemplateData is a map of data sent through
// handler to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string // use go get github.com/justinas/nosurf
	Flash     string
	Warning   string
	Error     string
}
