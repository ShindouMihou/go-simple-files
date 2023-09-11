package siopao

func read[T any](file *File, fn func() (*T, error)) (*T, error) {
	err := file.openRead()
	if err != nil {
		return nil, err
	}
	defer file.close()
	return fn()
}

func write[T any](file *File, trunc bool, fn func() (*T, error)) (*T, error) {
	if err := file.openWrite(trunc); err != nil {
		return nil, err
	}
	defer file.close()
	return fn()
}
