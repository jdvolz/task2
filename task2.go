package main

type DoubleSet map[int]int

// func (ds1 DoubleSet) Add(ds2 DoubleSet) DoubleSet {
//   ds := DoubleSet
//   for k, v := range ds1 {
//
//   }
//
//   return ds
// }
//

func (ds1 DoubleSet) Subtract(ds2 DoubleSet) DoubleSet {
	ds := make(DoubleSet)
	for k, v := range ds1 {
		var cv int
		v2, ok := ds2[k]
		if ok {
			cv = v - v2
		} else {
			cv = v
		}

		if cv == 2 || cv == 1 {
			ds[k] = cv
		}
	}

	return ds
}

func main() {
}
