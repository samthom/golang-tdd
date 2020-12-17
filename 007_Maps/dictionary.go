package maps

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary type
type Dictionary map[string]string

// Search function to search in dict
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil

}

// Add function adds new value to the dict
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
