package types

var forbiddenZoneTypeIds = []int{1, 2, 3, 4} //nolint:gochecknoglobals
var paidParkingZoneTypeIds = []int{7}        //nolint:gochecknoglobals

type Zone struct {
	ID          int64
	Address     string
	TypeID      int
	IsDisabled  bool
	Coordinates []Point
}

type Zones struct {
	Disabled    []Zone
	Forbidden   []Zone
	PaidParking []Zone
}

func NewZone(
	id int64,
	address string,
	typeID int,
	isDisabled bool,
	coordinates []Point,
) *Zone {
	return &Zone{
		ID:          id,
		Address:     address,
		TypeID:      typeID,
		IsDisabled:  isDisabled,
		Coordinates: coordinates,
	}
}

func (c *Zone) IsProcessable() bool {
	return c.IsForbidden() || c.IsPaidParking()
}

func (c *Zone) IsForbidden() bool {
	for _, i := range forbiddenZoneTypeIds {
		if i == c.TypeID {
			return true
		}
	}

	return false
}

func (c *Zone) IsPaidParking() bool {
	for _, i := range paidParkingZoneTypeIds {
		if i == c.TypeID {
			return true
		}
	}

	return false
}

func (c *Zone) BuildCoordinates() [][][]float64 {
	result := make([][]float64, 0)

	for _, coords := range c.Coordinates {
		result = append(result, []float64{coords.Y, coords.X})
	}

	r := make([][][]float64, 1)
	r[0] = result

	return r
}

func (c *Zone) Type() string {
	switch {
	case c.IsDisabled:
		return "disabled"
	case c.IsForbidden():
		return "forbidden"
	case c.IsPaidParking():
		return "paid_parking"
	default:
		return "unknown"
	}
}
