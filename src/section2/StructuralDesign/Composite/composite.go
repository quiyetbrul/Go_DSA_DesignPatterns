package composite

/*
The Composite pattern is a structural design pattern that
allows to compose objects into tree structures to
represent part-whole hierarchies.
This pattern lets clients treat individual objects
and compositions of objects uniformly.

- helps to shape applications with commonly used structures and relationships
  - composition "has-a" relationship, inheritance "is-a" relationship
- golang encourages use of composition almost exclusively by its lack of inheritance

- TWO TYPES OF COMPOSITION
  - direct composition
  - embedding composition
*/

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim() {
	println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

type Animal struct{}

func (a *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}
