package types_test

import (
	"testing"

	. "github.com/kadyrov-dev/observation-zones-importer/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_Zone_HasID(t *testing.T) {
	z := NewZone(42, "", 0, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.Equal(t, z.ID, int64(42))
}

func Test_Zone_HasAddress(t *testing.T) {
	z := NewZone(0, "ул. Генерала Иванова", 0, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.Equal(t, z.Address, "ул. Генерала Иванова")
}

func Test_Zone_HasCoordinates(t *testing.T) {
	z := NewZone(
		0,
		"ул. Генерала Иванова",
		0,
		false,
		[]Point{
			{
				X: 56.01053301310651,
				Y: 37.16082706530175,
			},
			{
				X: 56.01038381620671,
				Y: 37.16117038805566,
			},
		},
	)

	expected := make([][][]float64, 1)
	expected[0] = [][]float64{
		{37.16082706530175, 56.01053301310651},
		{37.16117038805566, 56.01038381620671},
	}

	assert.Equal(t, z.BuildCoordinates(), expected)
}

func Test_Zone_IsProcessable_whenTypeIdIsInForbiddenList_returnsTrue(t *testing.T) {
	for i := 1; i <= 4; i++ {
		z := NewZone(0, "", i, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

		assert.True(t, z.IsProcessable())
	}
}

func Test_Zone_IsProcessable_whenTypeIdIsInPaidParkingList_returnsTrue(t *testing.T) {
	z := NewZone(0, "", 7, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.True(t, z.IsProcessable())
}

func Test_Zone_IsProcessable_whenTypeIdIsNotInForbiddenAndPaidParkingLists_returnsFalse(t *testing.T) {
	z := NewZone(0, "", 11, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.False(t, z.IsProcessable())
}

func Test_Zone_IsForbidden_whenTypeIdIsInForbiddenList_returnsTrue(t *testing.T) {
	for i := 1; i <= 4; i++ {
		z := NewZone(0, "", i, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

		assert.True(t, z.IsForbidden())
	}
}

func Test_Zone_IsForbidden_whenTypeIdIsNotInForbiddenList_returnsFalse(t *testing.T) {
	z := NewZone(0, "", 7, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.False(t, z.IsForbidden())
}

func Test_Zone_IsPaidParking_whenTypeIdIsInPaidParkingList_returnsTrue(t *testing.T) {
	z := NewZone(0, "", 7, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.True(t, z.IsPaidParking())
}

func Test_Zone_IsPaidParking_whenTypeIdIsNotInPaidParkingList_returnsFalse(t *testing.T) {
	z := NewZone(0, "", 5, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.False(t, z.IsPaidParking())
}

func Test_Zone_IsDisabled_whenZoneIsDisabled_returnsTrue(t *testing.T) {
	z := NewZone(0, "", 7, true, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.True(t, z.IsDisabled)
}

func Test_Zone_IsDisabled_whenZoneIsEnabled_returnsFalse(t *testing.T) {
	z := NewZone(0, "", 5, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.False(t, z.IsDisabled)
}

func Test_Zone_Type_whenZoneIsDisabled(t *testing.T) {
	z := NewZone(0, "", 7, true, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.Equal(t, z.Type(), "disabled")
}

func Test_Zone_Type_whenZoneIsForbidden(t *testing.T) {
	z := NewZone(0, "", 4, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.Equal(t, z.Type(), "forbidden")
}

func Test_Zone_Type_whenZoneIsPaidParking(t *testing.T) {
	z := NewZone(0, "", 7, false, []Point{{X: 56.01140569096833, Y: 37.192046637126275}})

	assert.Equal(t, z.Type(), "paid_parking")
}
