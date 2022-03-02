package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

//type myInt int8

const SidesTriangle int = 3
const SidesSquare int = 4
const SidesCircle int = 1

func CalcSquare(sideLen float64, sidesNum int) float64 {
	switch sidesNum {
	case 0:
		return 0
	case SidesTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(sideLen) / 4
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return math.Pow(sideLen, 2) * math.Pi
	}
	return 1
}
