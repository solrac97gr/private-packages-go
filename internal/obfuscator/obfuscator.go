package obfuscator

type Obfuscator struct{}

func NewObfuscator() *Obfuscator {
	return &Obfuscator{}
}

func (o *Obfuscator) Obfuscate(msg string) string {
	// Here you should put the logic for hide the information in the Log message
	return msg
}
