package metric

import "testing"

func TestNameBuilder(t *testing.T) {
	var nb NameBuilder

	res := nb.Name("test_metric").Add("handler", "/").Add("method", "GET").Add("Path", "/carts").String()

	expected := "test_metric{handler=\"/\",method=\"GET\"}"

	if res != expected {
		t.Log(expected)
		t.Log(res)
		t.Error("not equal")
	}
}
