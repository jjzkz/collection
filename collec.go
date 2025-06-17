package collection

import (
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Collect[T any] struct {
	o    []T // original data
	n    []T // new data
	data any
}

func NewCollect[T any](data ...T) *Collect[T] {
	newData := make([]T, len(data))
	copy(newData, data)
	return &Collect[T]{
		o: data,
		n: newData,
	}
}

func (c *Collect[T]) All() []T {
	return c.n
}

func (c *Collect[T]) First() T {
	var zero T
	if len(c.n) == 0 {
		return zero
	}
	return c.n[0]
}

func (c *Collect[T]) Last() T {
	var last T
	if len(c.n) == 0 {
		return last
	}
	return c.n[len(c.n)-1]
}

// Where 仅仅支持 map 类型的数据
func (c *Collect[T]) Where(key any, value ...any) *Collect[T] {
	var result []T
	for _, v0 := range c.n {
		if len(value) == 1 {
			val := value[0]
			switch any(v0).(type) {
			case map[any]any:
				if v, ok := any(v0).(map[any]any)[key]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[string]any:
				if v, ok := any(v0).(map[string]any)[key.(string)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[int8]any:
				if v, ok := any(v0).(map[int8]any)[key.(int8)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[uint8]any:
				if v, ok := any(v0).(map[int8]any)[key.(int8)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[int16]any:
				if v, ok := any(v0).(map[int16]any)[key.(int16)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[uint16]any:
				if v, ok := any(v0).(map[uint16]any)[key.(uint16)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[int32]any:
				if v, ok := any(v0).(map[int32]any)[key.(int32)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[uint32]any:
				if v, ok := any(v0).(map[uint32]any)[key.(uint32)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[int64]any:
				if v, ok := any(v0).(map[int64]any)[key.(int64)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			case map[uint64]any:
				if v, ok := any(v0).(map[uint64]any)[key.(uint64)]; ok && reflect.DeepEqual(v, val) {
					result = append(result, v0)
				}
			default:

			}
		} else if len(value) == 2 {
			// = > < >= <= !=
			compareTag, ok := value[0].(string)
			val := value[1]

			fn := func(ok bool, compareTag string, v1, v2 any) {
				if ok && compareTag == ">" && nc(v1).Gt(v2) {
					result = append(result, v0)
				} else if ok && compareTag == "<" && nc(v1).Lt(v2) {
					result = append(result, v0)
				} else if ok && compareTag == ">=" && nc(v1).GtE(v2) {
					result = append(result, v0)
				} else if ok && compareTag == "<=" && nc(v1).LtE(v2) {
					result = append(result, v0)
				} else if ok && compareTag == "=" && nc(v1).Eq(v2) {
					result = append(result, v0)
				} else if ok && compareTag == "!=" && nc(v1).NEq(v2) {
					result = append(result, v0)
				} else if ok && compareTag == "like" {
					// "%you" "you%" "%you%"
					v1Str, ok1 := v1.(string)
					v2Str, ok2 := v2.(string)
					if ok1 && ok2 {

						if matchedAny, _ := regexp.MatchString("^%.+%$", v2Str); matchedAny {
							// %you%
							if strings.Contains(v1Str, strings.TrimLeft(strings.TrimRight(v2Str, "%"), "%")) {
								result = append(result, v0)
							}
						} else if matchedPrefix, _ := regexp.MatchString("^%.+", v2Str); matchedPrefix {
							// %you
							if strings.HasSuffix(v1Str, strings.TrimLeft(v2Str, "%")) {
								result = append(result, v0)
							}
						} else if matchedSufix, _ := regexp.MatchString(".+%$", v2Str); matchedSufix {
							// you%
							if strings.HasPrefix(v1Str, strings.TrimRight(v2Str, "%")) {
								result = append(result, v0)
							}
						}
					}
				}
			}

			switch any(v0).(type) {
			case map[any]any:
				if v, ok2 := any(v0).(map[any]any)[key]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[string]any:
				if v, ok2 := any(v0).(map[string]any)[key.(string)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[int8]any:
				if v, ok2 := any(v0).(map[int8]any)[key.(int8)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[uint8]any:
				if v, ok2 := any(v0).(map[uint8]any)[key.(uint8)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[int16]any:
				if v, ok2 := any(v0).(map[int16]any)[key.(int16)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[uint16]any:
				if v, ok2 := any(v0).(map[uint16]any)[key.(uint16)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[int32]any:
				if v, ok2 := any(v0).(map[int32]any)[key.(int32)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[uint32]any:
				if v, ok2 := any(v0).(map[uint32]any)[key.(uint32)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[int64]any:
				if v, ok2 := any(v0).(map[int64]any)[key.(int64)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			case map[uint64]any:
				if v, ok2 := any(v0).(map[uint64]any)[key.(uint64)]; ok2 {
					fn(ok, compareTag, v, val)
				}
			default:

			}

		}
	}
	c.n = result
	return c
}

// Filter 过滤数据 对于切片结构体过滤
func (c *Collect[T]) Filter(fn func(item, value interface{}) bool) *Collect[T] {
	var result []T
	for k, v := range c.n {
		if fn(k, v) {
			result = append(result, v)
		}
	}
	c.n = result
	return c
}

// Pop 弹出最后一个元素
func (c *Collect[T]) Pop() T {
	last := c.Last()
	var result []T
	d := c.n[:len(c.n)-1]
	result = d
	c.n = result
	return last
}

// Shift 弹出第一个元素
func (c *Collect[T]) Shift() T {
	first := c.First()
	var result []T
	d := c.n[1:len(c.n)]
	result = d
	c.n = result
	return first
}

// 打乱 支持字符串、切片
func (c *Collect[T]) Shuffle() *Collect[T] {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c.n), func(i, j int) {
		c.n[i], c.n[j] = c.n[j], c.n[i]
	})
	return c
}

// Sort 排序 []string []int []string []float64 []float32...
func (c *Collect[T]) Sort() *Collect[T] {
	// []string []int []string []float64 []float32...
	if len(c.n) == 0 {
		return c
	}
	var result []T
	switch any(c.n[0]).(type) {
	case string:
		var strSlice []string
		for _, v := range c.n {
			s, ok := any(v).(string)
			if ok {
				strSlice = append(strSlice, s)
			}
		}
		sort.Strings(strSlice)
		for _, vv := range strSlice {
			result = append(result, any(vv).(T))
		}
	case float64:
		var float64Slice []float64
		for _, v := range c.n {
			i, ok := any(v).(float64)
			if ok {
				float64Slice = append(float64Slice, i)
			}
		}
		sort.Float64s(float64Slice)
		for _, vv := range float64Slice {
			result = append(result, any(vv).(T))
		}
	case float32:
		var float32Slice []float32
		for _, v := range c.n {
			i, ok := any(v).(float32)
			if ok {
				float32Slice = append(float32Slice, i)
			}
		}
		var float64Slice []float64
		for _, v := range float32Slice {
			float64Slice = append(float64Slice, float64(v))
		}

		sort.Float64s(float64Slice)
		for _, vv := range float64Slice {
			vvv := float32(vv)
			result = append(result, any(vvv).(T))
		}

	case int:
		var intSlice []int
		for _, v := range c.n {
			i, ok := any(v).(int)
			if ok {
				intSlice = append(intSlice, i)
			}
		}
		sort.Ints(intSlice)
		for _, vv := range intSlice {
			result = append(result, any(vv).(T))
		}

	case uint:
		var uintSlice []uint
		for _, v := range c.n {
			i, ok := any(v).(uint)
			if ok {
				uintSlice = append(uintSlice, i)
			}
		}

		var intSlice []int
		for _, v := range uintSlice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := uint(vv)
			result = append(result, any(vvv).(T))
		}

	case int8:
		var int8Slice []int8
		for _, v := range c.n {
			i, ok := any(v).(int8)
			if ok {
				int8Slice = append(int8Slice, i)
			}
		}

		var intSlice []int
		for _, v := range int8Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := int8(vv)
			result = append(result, any(vvv).(T))
		}
	case uint8:
		var uint8Slice []uint8
		for _, v := range c.n {
			i, ok := any(v).(uint8)
			if ok {
				uint8Slice = append(uint8Slice, i)
			}
		}

		var intSlice []int
		for _, v := range uint8Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := uint8(vv)
			result = append(result, any(vvv).(T))
		}
	case int16:
		var int16Slice []int16
		for _, v := range c.n {
			i, ok := any(v).(int16)
			if ok {
				int16Slice = append(int16Slice, i)
			}
		}

		var intSlice []int
		for _, v := range int16Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := int16(vv)
			result = append(result, any(vvv).(T))
		}
	case uint16:
		var uint16Slice []uint16
		for _, v := range c.n {
			i, ok := any(v).(uint16)
			if ok {
				uint16Slice = append(uint16Slice, i)
			}
		}

		var intSlice []int
		for _, v := range uint16Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := uint16(vv)
			result = append(result, any(vvv).(T))
		}
	case int32:
		var int32Slice []int32
		for _, v := range c.n {
			i, ok := any(v).(int32)
			if ok {
				int32Slice = append(int32Slice, i)
			}
		}

		var intSlice []int
		for _, v := range int32Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := int32(vv)
			result = append(result, any(vvv).(T))
		}
	case uint32:
		var uint32Slice []uint32
		for _, v := range c.n {
			i, ok := any(v).(uint32)
			if ok {
				uint32Slice = append(uint32Slice, i)
			}
		}

		var intSlice []int
		for _, v := range uint32Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := uint32(vv)
			result = append(result, any(vvv).(T))
		}
	case int64:
		var int64Slice []int64
		for _, v := range c.n {
			i, ok := any(v).(int64)
			if ok {
				int64Slice = append(int64Slice, i)
			}
		}

		var intSlice []int
		for _, v := range int64Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := int64(vv)
			result = append(result, any(vvv).(T))
		}
	case uint64:
		var uint64Slice []uint64
		for _, v := range c.n {
			i, ok := any(v).(uint64)
			if ok {
				uint64Slice = append(uint64Slice, i)
			}
		}

		var intSlice []int
		for _, v := range uint64Slice {
			intSlice = append(intSlice, int(v))
		}

		sort.Ints(intSlice)
		for _, vv := range intSlice {
			vvv := uint64(vv)
			result = append(result, any(vvv).(T))
		}

	}

	c.n = result
	return c
}

// SortDesc
func (c *Collect[T]) SortDesc() *Collect[T] {
	c.Sort()
	for i, j := 0, len(c.n)-1; i < j; i, j = i+1, j-1 {
		c.n[i], c.n[j] = c.n[j], c.n[i]
	}
	return c
}

// SortBy 排序 []map[any]any []map[string]any
func (c *Collect[T]) SortBy(keyOrFunc ...any) *Collect[T] {

	return c
}

// SortByDesc
func (c *Collect[T]) SortByDesc(keyOrFunc string) *Collect[T] {

	return c
}

// Values 获取指定 key 的值 []map[any]any []map[string]any []map[string]string []map[string]string
func (c *Collect[T]) Values(key any) []any {
	var result = make([]any, 0)
	for _, v := range c.n {
		switch any(v).(type) {
		case map[any]any:
			if vMap, ok := any(v).(map[any]any); ok {
				if vv, ok2 := vMap[key]; ok2 {
					result = append(result, vv)
				}
			}
		case map[string]any:
			if _, ok := key.(string); !ok {
				return result
			}
			if vMap, ok := any(v).(map[string]any); ok {
				if vv, ok2 := vMap[key.(string)]; ok2 {
					result = append(result, vv)
				}
			}
		case map[string]string:
			if _, ok := key.(string); !ok {
				return result
			}
			if vMap, ok := any(v).(map[string]string); ok {
				if vv, ok2 := vMap[key.(string)]; ok2 {
					result = append(result, vv)
				}
			}
		case map[string]int:
			if _, ok := key.(string); !ok {
				return result
			}
			if vMap, ok := any(v).(map[string]int); ok {
				if vv, ok2 := vMap[key.(string)]; ok2 {
					result = append(result, vv)
				}
			}

		}

	}
	return result
}
