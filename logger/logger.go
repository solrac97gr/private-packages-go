package logger

import (
	"fmt"

	"private-packages-go/internal/obfuscator"
)

type Logger interface {
	Debug(msg string)
}

type CompanyLogger struct {
	obfuscator *obfuscator.Obfuscator
}

func NewCompanyLogger() *CompanyLogger {
	o := obfuscator.NewObfuscator()

	return &CompanyLogger{
		obfuscator: o,
	}
}

func (c *CompanyLogger) Debug(msg string) {
	nMsg := c.obfuscator.Obfuscate(msg)
	fmt.Println(nMsg)
}
