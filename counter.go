package counter

import "sort"

type Counter struct {
	counts map[string]int
}

func NewCounter(slice []string) *Counter {
	c := new(Counter)

	counts := make(map[string]int)
	for _, k := range slice {
		counts[k] += 1
	}
	c.counts = counts

	return c
}

func NewCounterFromMap(counts map[string]int) *Counter {
	c := new(Counter)
	c.counts = counts

	return c
}

func (c *Counter) Len() int {
	return len(c.counts)
}

func (c *Counter) Count(k string) int {
	return c.counts[k]
}

func (c *Counter) Contains(k string) bool {
	_, ok := c.counts[k]
	return ok
}

func (c *Counter) getFrequencies(descending bool) ([]string, []int) {
	if len(c.counts) == 0 {
		return []string{}, []int{}
	}

	type kv struct {
		key   string
		value int
	}
	var kvs []kv
	for k, v := range c.counts {
		kvs = append(kvs, kv{k, v})
	}

	if descending {
		sort.Slice(kvs, func(i, j int) bool {
			if kvs[i].value > kvs[j].value {
				return true
			}
			if kvs[i].value < kvs[j].value {
				return false
			}

			return kvs[i].key > kvs[j].key
		})
	} else {
		sort.Slice(kvs, func(i, j int) bool {
			if kvs[i].value < kvs[j].value {
				return true
			}
			if kvs[i].value > kvs[j].value {
				return false
			}

			return kvs[i].key < kvs[j].key
		})
	}

	keys := make([]string, len(kvs))
	freqs := make([]int, len(kvs))
	for i, kv := range kvs {
		keys[i] = kv.key
		freqs[i] = kv.value
	}

	return keys, freqs
}

func (c *Counter) MostCommon() ([]string, []int) {
	return c.getFrequencies(true)
}
