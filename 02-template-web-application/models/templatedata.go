package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // catch all
	CSRFToken string
	Flash     string // success message
	Warning   string
	Error     string
}
