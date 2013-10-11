package TSUtil

import (
	"strconv"
)

func ToString(a interface{}) string {
	if v, p := a.(float32); p {
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	}
	if v, p := a.(float64); p {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	if v, p := a.(int); p {
		return strconv.Itoa(v)
	}
	if v, p := a.(uint); p {
		return strconv.Itoa(int(v))
	}

	if v, p := a.(int64); p {
		return strconv.FormatInt(v, 10)
	}
	if v, p := a.(uint64); p {
		return strconv.FormatUint(v, 10)
	}

	if v, p := a.(int32); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(uint32); p {
		return strconv.Itoa(int(v))
	}

	if v, p := a.(int16); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(uint16); p {
		return strconv.Itoa(int(v))
	}

	if v, p := a.(int8); p {
		return strconv.Itoa(int(v))
	}
	if v, p := a.(uint8); p {
		return strconv.Itoa(int(v))
	}

	return "wrong"
}
