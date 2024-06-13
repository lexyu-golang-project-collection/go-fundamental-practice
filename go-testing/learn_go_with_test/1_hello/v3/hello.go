package hello

const (
	PREFIX        = "Hello, "
	spanish       = "Spanish"
	spanishPrefix = "Hola, "
	french        = "French"
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = PREFIX
	}

	return prefix
}
