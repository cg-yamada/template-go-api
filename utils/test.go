package utils

/// IgnoreFailed エラーを返り値に持つ関数の、エラー返り値を無視する
/// WARN: テストでのみ使用すること
func IgnoreFailed[T any, U any](value T, _ U) T {
	return value
}

/// GetPtr 値のポインタを取得する
/// WARN: テストでのみ使用すること
func GetPtr[T any](value T) *T {
	return &value
}
