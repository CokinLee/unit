package unit

import (
	"fmt"
	"reflect"
)

// Converter 转换器
type Converter struct {
	data float64
	from Unit
	to   Unit
}

// NewConverter create a new converter
func NewConverter(data float64, from Unit, to Unit) Converter {
	return Converter{data: data, from: from, to: to}
}

// Convert convert unit by converter
func (c Converter) Convert() (float64, error) {
	st := reflect.TypeOf(c.from)
	dt := reflect.TypeOf(c.to)
	// 解析source
	if st.Name() != dt.Name() {
		return 0, fmt.Errorf("unit type error: source type[%s], destination type[%s]", st.Name(), dt.Name())
	}
	d := c.from.Float64() * c.data / c.to.Float64()
	return d, nil
}

// Convert convert unit
func Convert(data float64, from Unit, to Unit) (float64, error) {
	st := reflect.TypeOf(from)
	dt := reflect.TypeOf(to)
	// 解析source
	if st.Name() != dt.Name() {
		return 0, fmt.Errorf("unit type error: source type[%s], destination type[%s]", st.Name(), dt.Name())
	}
	d := from.Float64() * data / to.Float64()
	return d, nil
}
