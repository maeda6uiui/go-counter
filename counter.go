package counter

import "sort"

type Counter struct {
	counts map[string]int

	keys  []string
	freqs []int
}

func newCounter(slice []string) *Counter {
	c := new(Counter)
	c.keys = make([]string, 0)
	c.freqs = make([]int, 0)

	counts := make(map[string]int)
	for _, k := range slice {
		counts[k] += 1
	}
	c.counts = counts

	return c
}

func (c *Counter) len() int {
	return len(c.counts)
}

func (c *Counter) count(k string) int {
	return c.counts[k]
}

func (c *Counter) contains(k string) bool {
	_, ok := c.counts[k]
	return ok
}

func (c *Counter) mostCommon() ([]string, []int) {
	if len(c.keys) == 0 {
		type kv struct {
			key   string
			value int
		}
		var kvs []kv
		for k, v := range c.counts {
			kvs = append(kvs, kv{k, v})
		}

		sort.Slice(kvs, func(i, j int) bool {
			if kvs[i].value > kvs[j].value {
				return true
			}
			if kvs[i].value < kvs[j].value {
				return false
			}

			return kvs[i].key > kvs[j].key
		})

		for _, kv := range kvs {
			c.keys = append(c.keys, kv.key)
			c.freqs = append(c.freqs, kv.value)
		}
	}

	return c.keys, c.freqs
}
