package hello

const PREFIX = "Hello, "
const spanish = "Spanish"
const spanishPrefix = "Hola, "
const french = "French"
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := PREFIX

	// 1.
	/*
		if language == spanish {
			return spanishPrefix + name
		}

		if language == french {
			return frenchPrefix + name
		}
	*/

	// 2.
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}

	return prefix + name
}
