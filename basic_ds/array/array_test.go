package array

import "testing"

func Test(t *testing.T) {

	a := NewArray(0, 3)
	a.Print()

	a.Append(10)
	a.Print()

	a.Append(9)
	a.Print()

	a.AppendMany(8, 7, 6)
	a.Print()

}
