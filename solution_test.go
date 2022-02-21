package square

import "testing"

func TestCalcSquare(t *testing.T) {
	tables := []struct {
		sideLen float64
		sides   side
		want    float64
	}{

		{2.0, 4, 4.0},
		{3.0, 0, 28.274333882308138},
		{4.0, 3, 6.928203230275509},
	}
	for _, table := range tables {
		area := CalcSquare(table.sideLen, table.sides)
		if area != table.want {
			t.Errorf("area are not same got %v want %v", area, table.want)
		}
	}
}
