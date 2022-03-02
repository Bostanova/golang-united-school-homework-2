package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type myInt int8
type SidesTriangle myInt
type SidesSquare myInt
type SidesCircle myInt

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	switch sidesNum {
	case 0:
		return 0
	case 3:
		return math.Pow(sideLen, 2) * math.Sqrt(sideLen) / 4
	case 4:
		return math.Pow(sideLen, 2)
	}
	return 1
}
