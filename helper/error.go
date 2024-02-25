package helper

func ReturnIfError(err error) error {
	if err != nil {
		return err
	}

	return nil
}

func PanicIfError(err error) {
	panic(err)
}