package hashmaps

type HashMap struct {
	size int
	data [][]any
}

func (h *HashMap) Set(key string, value any) {

	element := make([]any, 2)
	element[0] = key
	element[1] = value

	h.data[NaiveHash(key)] = element

}

func (h *HashMap) Get(key string) (int, bool) {

	index := NaiveHash(key)
	if index >= len(h.data) {
		return 0, false
	} else {

		tuple := h.data[index]
		value, ok := 0, false
		if len(tuple) > 1 {
			value, ok = tuple[1].(int), true
		}
		return value, ok
	}

}

func New(length int) *HashMap {

	return &HashMap{

		size: length,
		data: make([][]any, length),
	}

}

func NaiveHash(data string) int {

	runes := []rune(data)

	result := 0

	for k, v := range runes {
		// fmt.Println(int(v))
		result += (result + int(v)*k) % len(data)
	}

	return result

}

func TestingHashMaps() {

}
