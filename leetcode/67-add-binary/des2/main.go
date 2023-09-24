package des2

import (
	"math/big"
)

func addBinary(a string, b string) string {
	bigA := big.Int{}
	bigB := big.Int{}
	bigA.SetString(a, 2)
	bigB.SetString(b, 2)
	bigR := big.Int{}
	bigR.Add(&bigA, &bigB)
	return bigR.Text(2)
}
