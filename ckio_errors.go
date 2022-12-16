package ckio

type FileCloseError struct {
	errorMessage string
}

func NewFileCloseError(errorMessage string) *FileCloseError {
	return &FileCloseError{errorMessage: errorMessage}
}

func (f *FileCloseError) Error() string { return f.errorMessage }

type FileOpenError struct {
	errorMessage string
}

func NewFileOpenError(errorMessage string) *FileOpenError {
	return &FileOpenError{errorMessage: errorMessage}
}

func (f *FileOpenError) Error() string { return f.errorMessage }

type ReadRuneError struct {
	errorMessage string
}

func NewReadRuneError(errorMessage string) *ReadRuneError {
	return &ReadRuneError{errorMessage: errorMessage}
}

func (r *ReadRuneError) Error() string { return r.errorMessage }
