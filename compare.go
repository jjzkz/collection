package collection

import (
	"fmt"
)

type Compare struct {
	data any
}

func nc(data any) Compare {
	return Compare{data: data}
}

func (c Compare) isNumber(n any) bool {
	switch n.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}

}

func (c *Compare) compareEnable(n any) bool {
	return c.isNumber(c.data) && c.isNumber(n)
}

func (c Compare) Gt(f any) bool {
	if !c.compareEnable(f) {
		return false
	}
	i, err := compare(c.data, f)
	if err != nil {
		return false
	}
	if i == 1 {
		return true
	}
	return false
}

func (c Compare) GtE(f any) bool {
	return c.Gt(f) || c.Eq(f)
}

func (c Compare) Lt(f any) bool {
	if !c.compareEnable(f) {
		return false
	}
	i, err := compare(c.data, f)
	if err != nil {
		return false
	}
	if i == -1 {
		return true
	}
	return false
}

func (c Compare) LtE(f any) bool {
	return c.Lt(f) || c.Eq(f)
}

func (c Compare) Eq(f any) bool {
	if !c.compareEnable(f) {
		return false
	}
	i, err := compare(c.data, f)
	if err != nil {
		return false
	}
	if i == 0 {
		return true
	}
	return false
}

func (c Compare) NEq(f any) bool {
	return !c.Eq(f)
}

// compare 比较两个 any 类型的值
// 返回值：
// -1: v1 < v2
//
//	0: v1 == v2
//	1: v1 > v2
//
// error: 无法比较或类型不支持比较
func compare(v1, v2 any) (int, error) {
	// 处理基本数值类型
	switch v1 := v1.(type) {
	case int:
		v2 := v2.(int)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case int8:
		v2 := v2.(int8)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case int16:
		v2 := v2.(int16)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case int32:
		v2 := v2.(int32)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case int64:
		v2 := v2.(int64)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case uint:
		v2 := v2.(uint)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case uint8:
		v2 := v2.(uint8)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case uint16:
		v2 := v2.(uint16)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case uint32:
		v2 := v2.(uint32)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case uint64:
		v2 := v2.(uint64)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case float32:
		v2 := v2.(float32)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case float64:
		v2 := v2.(float64)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	case string:
		v2 := v2.(string)
		switch {
		case v1 < v2:
			return -1, nil
		case v1 == v2:
			return 0, nil
		default:
			return 1, nil
		}
	default:
		return 0, fmt.Errorf("unsupported type")
	}
}
