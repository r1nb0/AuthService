package logging

// Logger
// TODO (impl Infof, Debugf, Warnf, Errorf, Fatalf)
type Logger interface {
	Init()

	Info(cat Category, sub SubCategory, msg string, another map[string]interface{})
	Debug(cat Category, sub SubCategory, msg string, another map[string]interface{})
	Warn(cat Category, sub SubCategory, msg string, another map[string]interface{})
	Error(cat Category, sub SubCategory, msg string, another map[string]interface{})
	Fatal(cat Category, sub SubCategory, msg string, another map[string]interface{})
}
