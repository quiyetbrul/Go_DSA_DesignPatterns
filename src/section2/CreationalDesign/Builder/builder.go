package builder

/*
- helps to construct complex objects without directly intantiating their struct
- avoids writing the logic to create all the objects in the package
*/

/*
Vehicle Manufacturing

- create relationship between a director, a few builders, and the product they build
- every kind of vehicle, choose vehicle type, assemble the structurem place the wheels, and place the seats
- director is represented by the ManufacturingDirector type

Requirements

- you must have a manufacturing type
- VehicleProduct with four wheels, five seats, and a structure defined as car must be returned
- VehicleProduct with two wheels, two seats, and a structure defined as motorbike must be returned
- VehicleProduct built by any BuildProcess builder must be open to modifications
*/

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// A Builder of type car
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// A Builder of type motorbike
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
