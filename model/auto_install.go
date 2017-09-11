package model

// TestCaseMap is
type TestCaseMap struct {
	Name    string `storm:"id"`
	CaseMap map[string][]string
}
