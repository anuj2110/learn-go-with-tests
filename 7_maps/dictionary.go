package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	errNotFound      = DictionaryErr("could not find the word")
	errWordExists    = DictionaryErr("word already exists")
	errWordNotExists = DictionaryErr("word does not exist, cannot perform operation")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
		return nil
	case nil:
		return errWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errWordNotExists
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errWordNotExists
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}

