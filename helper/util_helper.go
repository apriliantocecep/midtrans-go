package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func MessageForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return ""
}
