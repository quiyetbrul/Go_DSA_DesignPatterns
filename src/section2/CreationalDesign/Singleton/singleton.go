package creational

/*
- provides a single instance of an object
- guarantees that there is only one instance of an object in the system
- provides a global point of access to the object
*/

type singleton struct {
	count int
}

// can init a pointer to a struct as nil but cannot init a struct to nil
// equivalent to static variable in C++
var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
