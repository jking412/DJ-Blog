package config

func LoadConfig(key string, defaultValue ...interface{}) interface{} {
	if !internalViper.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return internalViper.Get(key)
}

func LoadString(key string, defaultValue ...interface{}) string {
	value := LoadConfig(key, defaultValue...)
	if value == nil {
		return ""
	}
	return value.(string)
}

func LoadInt(key string, defaultValue ...interface{}) int {
	value := LoadConfig(key, defaultValue...)
	if value == nil {
		return 0
	}
	return value.(int)
}
