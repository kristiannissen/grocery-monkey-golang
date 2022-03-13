package models
// The specific model should return something or an error
func (m *Model) GetUser(key string) string {
  return "Hello " + key
}
