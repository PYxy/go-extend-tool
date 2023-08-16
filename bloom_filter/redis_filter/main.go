package redis_filter

//
//import "gopkg.in/bculberson/bloom.v2"
//
//func main() {
//	m, k := bloom.EstimateParameters(1000, .01)
//	b := bloom.New(m, k, bloom.NewBitSet(m))
//	b.Add([]byte("some key"))
//	//exists, _ := b.Exists([]byte("some key"))
//	//doesNotExist, _ := b.Exists([]byte("some other key"))
//}
