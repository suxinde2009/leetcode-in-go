package ring

import "testing"

func Test(t *testing.T) {
	r := NewRing()
	r.Value = -1

	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	r.Link(&Ring{Value: 5})

	PrintRing(r)
}
