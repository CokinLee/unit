package unit

import (
	"testing"
)

func TestNumber(t *testing.T) {
	// SI
	assertFloatEqual(t, 1e2, (1 * Base).Percent())
	assertFloatEqual(t, 1e0, (1 * Base).Base())

	// CN
	assertFloatEqual(t, 1e-1, (1 * Base).Shi())
	assertFloatEqual(t, 1e-1, (1 * Shi).Bai())
	assertFloatEqual(t, 1e-1, (1 * Bai).Qian())
	assertFloatEqual(t, 1e-1, (1 * Qian).Wan())
	assertFloatEqual(t, 1e-1, (1 * Wan).ShiWan())
	assertFloatEqual(t, 1e-1, (1 * ShiWan).BaiWan())
	assertFloatEqual(t, 1e-1, (1 * BaiWan).QianWan())
	assertFloatEqual(t, 1e-1, (1 * QianWan).Yi())
	assertFloatEqual(t, 1e-1, (1 * Yi).ShiYi())
	assertFloatEqual(t, 1e-1, (1 * ShiYi).BaiYi())
	assertFloatEqual(t, 1e-1, (1 * BaiYi).QianYi())
	assertFloatEqual(t, 1e-1, (1 * QianYi).WanYi())
}
