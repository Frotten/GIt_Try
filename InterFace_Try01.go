package main

//type IOL interface {
//	Area() float64
//	Perim() float64
//}
//
//type Circle struct {
//	R float64
//}
//
//type Rectangle struct {
//	C, W float64
//}
//
//func (C Circle) Area() float64 {
//	return math.Pi * C.R * C.R
//}
//
//func (R Rectangle) Area() float64 {
//	return R.C * R.W
//}
//
//func (C Circle) Perim() float64 {
//	return 2 * math.Pi * C.R
//}
//
//func (R Rectangle) Perim() float64 {
//	return (R.C + R.W) * 2
//}
//func PrintArea(I IOL) {
//	fmt.Println("Area:", I.Area())
//	fmt.Println("Perim:", I.Perim())
//}
//
//func TestShape(I IOL) {
//	if Ins, ok := I.(Circle); ok {
//		fmt.Println("这是个圆，接下来输出它的面积和周长：")
//		PrintArea(Ins)
//	} else {
//		fmt.Println("这不是个圆，接下来输出它的面积和周长：")
//		PrintArea(I)
//	}
//}

//func main() {
//	Fang := Rectangle{4.5, 6.0}
//	Yuan := Circle{6.6}
//	TestShape(Yuan)
//	TestShape(Fang)
//}
