package best4

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

//func (i Fruit) String() string {
//	switch i {
//	case Apple:
//		return "Apple"
//	case Orange:
//		return "Orange"
//	case Banana:
//		return "Banana"
//	default:
//		return fmt.Sprintf("Unknown Fruit(%d)", i)
//	}
//}
