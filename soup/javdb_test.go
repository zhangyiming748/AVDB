package soup
import "testing"
func TestJavdb(t *testing.T) {
	result,err := Javdb("OEA-002")
	if err!=nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("Result: %v", result)
}
