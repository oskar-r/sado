package models

type Query struct {
	Dataset         string `json:"dataset,omitempty"`
	Query           string `json:"query,omitempty"`
	RecordDelimiter string `json:"record_delimiter,omitempty"`
	FieldDelimiter  string `json:"field_delimiter,omitempty"`
	Output          string `json:"output,omitempty"`
}
