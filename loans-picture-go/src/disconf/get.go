package disconf

import "time"

// Get returns the expanded value for the given key if exists.
// Otherwise, ok is false.
func Get(key string) (value string, ok bool) {
	return prop.Get(key)
}

// MustGet returns the expanded value for the given key if exists.
// Otherwise, it panics.
func MustGet(key string) string {
	return prop.MustGet(key)
}

// GetBool checks if the expanded value is one of '1', 'yes',
// 'true' or 'on' if the key exists. The comparison is case-insensitive.
// If the key does not exist the default value is returned.
func GetBool(key string, def bool) bool {
	return prop.GetBool(key, def)
}

// MustGetBool checks if the expanded value is one of '1', 'yes',
// 'true' or 'on' if the key exists. The comparison is case-insensitive.
// If the key does not exist the function panics.
func MustGetBool(key string) bool {
	return prop.MustGetBool(key)
}

// ----------------------------------------------------------------------------

// GetDuration parses the expanded value as an time.Duration (in ns) if the
// key exists. If key does not exist or the value cannot be parsed the default
// value is returned. In almost all cases you want to use GetParsedDuration().
func GetDuration(key string, def time.Duration) time.Duration {

	return prop.GetDuration(key, def)
}

// MustGetDuration parses the expanded value as an time.Duration (in ns) if
// the key exists. If key does not exist or the value cannot be parsed the
// function panics. In almost all cases you want to use MustGetParsedDuration().
func MustGetDuration(key string) time.Duration {
	return prop.MustGetDuration(key)
}

// ----------------------------------------------------------------------------

// GetParsedDuration parses the expanded value with time.ParseDuration() if the key exists.
// If key does not exist or the value cannot be parsed the default
// value is returned.
func GetParsedDuration(key string, def time.Duration) time.Duration {
	return prop.GetParsedDuration(key, def)
}

// MustGetParsedDuration parses the expanded value with time.ParseDuration() if the key exists.
// If key does not exist or the value cannot be parsed the function panics.
func MustGetParsedDuration(key string) time.Duration {
	return prop.MustGetParsedDuration(key)
}

// ----------------------------------------------------------------------------

// GetFloat64 parses the expanded value as a float64 if the key exists.
// If key does not exist or the value cannot be parsed the default
// value is returned.
func GetFloat64(key string, def float64) float64 {
	return prop.GetFloat64(key, def)
}

// MustGetFloat64 parses the expanded value as a float64 if the key exists.
// If key does not exist or the value cannot be parsed the function panics.
func MustGetFloat64(key string) float64 {
	return prop.MustGetFloat64(key)
}

// ----------------------------------------------------------------------------

// GetInt parses the expanded value as an int if the key exists.
// If key does not exist or the value cannot be parsed the default
// value is returned. If the value does not fit into an int the
// function panics with an out of range error.
func GetInt(key string, def int) int {

	return prop.GetInt(key, def)
}

// MustGetInt parses the expanded value as an int if the key exists.
// If key does not exist or the value cannot be parsed the function panics.
// If the value does not fit into an int the function panics with
// an out of range error.
func MustGetInt(key string) int {
	return prop.MustGetInt(key)
}

// ----------------------------------------------------------------------------

// GetInt64 parses the expanded value as an int64 if the key exists.
// If key does not exist or the value cannot be parsed the default
// value is returned.
func GetInt64(key string, def int64) int64 {
	return prop.GetInt64(key, def)
}

// MustGetInt64 parses the expanded value as an int if the key exists.
// If key does not exist or the value cannot be parsed the function panics.
func MustGetInt64(key string) int64 {
	return prop.MustGetInt64(key)
}

// ----------------------------------------------------------------------------

// GetUint parses the expanded value as an uint if the key exists.
// If key does not exist or the value cannot be parsed the default
// value is returned. If the value does not fit into an int the
// function panics with an out of range error.
func GetUint(key string, def uint) uint {
	return prop.GetUint(key, def)
}

// MustGetUint parses the expanded value as an int if the key exists.
// If key does not exist or the value cannot be parsed the function panics.
// If the value does not fit into an int the function panics with
// an out of range error.
func MustGetUint(key string) uint {
	return prop.MustGetUint(key)
}

// ----------------------------------------------------------------------------

// GetUint64 parses the expanded value as an uint64 if the key exists.
// If key does not exist or the value cannot be parsed the default
// value is returned.
func GetUint64(key string, def uint64) uint64 {
	return prop.GetUint64(key, def)
}

// MustGetUint64 parses the expanded value as an int if the key exists.
// If key does not exist or the value cannot be parsed the function panics.
func MustGetUint64(key string) uint64 {
	return prop.MustGetUint64(key)
}

// ----------------------------------------------------------------------------

// GetString returns the expanded value for the given key if exists or
// the default value otherwise.
func GetString(key, def string) string {
	return prop.GetString(key, def)
}

// MustGetString returns the expanded value for the given key if exists or
// panics otherwise.
func MustGetString(key string) string {
	return prop.MustGetString(key)
}

// Keys returns all keys in the same order as in the input.
func Keys() []string {

	return prop.Keys()
}
