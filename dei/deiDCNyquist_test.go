// Copyright Raul Vera 2022


package dei

import (
    //"fmt"
	"reflect"
    "testing"

    . "github.com/Causticity/sipp/sipptesting"
    "github.com/Causticity/sipp/sipptesting/sipptestcore"
)

// Give some test images. Perhaps use the ones in the testing infrastructure?

// Compute row and column and compare to golden

var spRowAvg []DCvalue
var spColAvg []DCvalue

func init() {
	spRowAvg = make ([]DCvalue, len(sipptestcore.SmallPicRowAverages))
	for i, v := range sipptestcore.SmallPicRowAverages {
		spRowAvg[i] = DCvalue(v)
	}
	spColAvg = make ([]DCvalue, len(sipptestcore.SmallPicColumnAverages))
	for i, v := range sipptestcore.SmallPicColumnAverages {
		spColAvg[i] = DCvalue(v)
	}
}

func TestDC(t *testing.T) {
	got := DCcolumn(Sgray)
	if !reflect.DeepEqual(got, spRowAvg) {
		t.Errorf("Error: Incorrect DC row averages. Expected:\n %v\nGot:%v\n",
			    sipptestcore.SmallPicRowAverages, got)
	}
	got = DCrow(Sgray)
	if !reflect.DeepEqual(got, spColAvg) {
		t.Errorf("Error: Incorrect DC row averages. Expected:\n %v\nGot:%v\n",
			    sipptestcore.SmallPicColumnAverages, got)
	}
}
