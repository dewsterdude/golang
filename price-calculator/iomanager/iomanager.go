package iomanager

type IOManager interface { // interface to support filemanager AND cmdmanager
	ReadLines() ([]string, error) // GO will look at various structs to see if they have these methods
	WriteResult(data interface{}) error
}
