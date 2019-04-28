// fake project fake.go
package fake

type Retrieve struct {
	UserAgent string
	Contents  string
}

func (r Retrieve) Get(url string) string {
	return "this is fake Retrieve"
}
