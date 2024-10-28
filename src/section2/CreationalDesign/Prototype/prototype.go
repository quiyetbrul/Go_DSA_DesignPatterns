package prototype

import (
	"errors"
	"fmt"
)

/*
- comes handy when creating objects is expensive
- utilize instance of existing object to create new object
- avoids repeated initialization of objects
- aims to have an obj or set of objs that are already created at compile time

- have a shirt-cloner obj and interface to ask for different types of shirts
- when you ask for a white shirt, a clone of the white shirt must be made
- the SKU of the created obj shouldn't affect new obj creation
- an info method must give all the information available on the instance fields
*/

type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
