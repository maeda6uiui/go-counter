package counter

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/xerrors"
)

func TestLen(t *testing.T) {
	tests := map[string]struct {
		items []string
		want  int
	}{
		"test_1": {
			items: []string{"a", "a", "a", "b", "c", "d", "a", "a", "d", "c"},
			want:  10,
		},
		"test_2": {
			items: []string{"こんにちは", "世界", "あいうえお", "Hello", "世界", "こんにちは", "Hello"},
			want:  7,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCounter(tt.items)
			if got := c.Len(); got != tt.want {
				t.Errorf("Len()=%v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := map[string]struct {
		items []string
		key   string
		want  int
	}{
		"test_1": {
			items: []string{"a", "a", "a", "b", "c", "d", "a", "a", "d", "c"},
			key:   "a",
			want:  5,
		},
		"test_2": {
			items: []string{"こんにちは", "世界", "あいうえお", "Hello", "世界", "こんにちは", "Hello"},
			key:   "こんにちは",
			want:  2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCounter(tt.items)
			if got := c.Count(tt.key); got != tt.want {
				t.Errorf("Count(%v)=%v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := map[string]struct {
		items []string
		keys  []string
		wants []bool
	}{
		"test_1": {
			items: []string{"a", "a", "a", "b", "c", "d", "a", "a", "d", "c"},
			keys:  []string{"a", "b", "c", "d", "e", "f", "g"},
			wants: []bool{true, true, true, true, false, false, false},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCounter(tt.items)
			for i := 0; i < len(tt.keys); i++ {
				key := tt.keys[i]
				want := tt.wants[i]
				if got := c.Contains(key); got != want {
					t.Errorf("Contains(%v)=%v, want %v", key, got, want)
				}
			}
		})
	}
}

func TestMostCommon(t *testing.T) {
	tests := map[string]struct {
		items     []string
		wantKeys  []string
		wantFreqs []int
	}{
		"test_1": {
			items:     []string{"a", "a", "a", "b", "c", "d", "a", "a", "d", "c"},
			wantKeys:  []string{"a", "d", "c", "b"},
			wantFreqs: []int{5, 2, 2, 1},
		},
		"test_2": {
			items:     []string{"aa", "aa", "aa", "bb", "cc", "dd", "aa", "aa", "dd", "cc"},
			wantKeys:  []string{"aa", "dd", "cc", "bb"},
			wantFreqs: []int{5, 2, 2, 1},
		},
		"test_3": {
			items:     []string{"あ", "あ", "い", "う", "え", "え", "え", "お"},
			wantKeys:  []string{"え", "あ", "い", "う", "お"},
			wantFreqs: []int{3, 2, 1, 1, 1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCounter(tt.items)
			gotKeys, gotFreqs := c.MostCommon()
			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("MostCommon().keys=%v want %v", gotKeys, tt.wantKeys)
			}
			if !reflect.DeepEqual(gotFreqs, tt.wantFreqs) {
				t.Errorf("MostCommon().freqs=%v want %v", gotFreqs, tt.wantFreqs)
			}
		})
	}
}

func loadCounts(filepath string) (map[string]int, error) {
	fp, err := os.Open(filepath)
	if err != nil {
		return nil, xerrors.Errorf("Failed to load file %v: %w", filepath, err)
	}
	defer fp.Close()

	counts := make(map[string]int)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")

		key := splits[1]
		freq, err := strconv.Atoi(splits[0])
		if err != nil {
			return nil, xerrors.Errorf("Failed to convert string to int: %v: %w", splits[0], err)
		}

		counts[key] = freq
	}
	if err := scanner.Err(); err != nil {
		return nil, xerrors.Errorf("Failed to load file %v: %w", filepath, err)
	}

	return counts, nil
}

func TestMostCommon2(t *testing.T) {
	tests := map[string]struct {
		sortedCountsFilepath   string
		shuffledCountsFilepath string
		loadCounts             func(filepath string) (map[string]int, error)
	}{
		"test_1": {
			sortedCountsFilepath:   "./Data/string_counts_sorted.txt",
			shuffledCountsFilepath: "./Data/string_counts_shuffled.txt",
			loadCounts:             loadCounts,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sortedCounts, err := tt.loadCounts(tt.sortedCountsFilepath)
			if err != nil {
				t.Errorf("Failed to load counts file: %v", tt.sortedCountsFilepath)
			}

			shuffledCounts, err := tt.loadCounts(tt.shuffledCountsFilepath)
			if err != nil {
				t.Errorf("Failed to load counts file: %v", tt.shuffledCountsFilepath)
			}

			c := NewCounterFromMap(shuffledCounts)
			gotKeys, gotFreqs := c.MostCommon()

			got := make(map[string]int, len(gotKeys))
			for i := 0; i < len(gotKeys); i++ {
				got[gotKeys[i]] = gotFreqs[i]
			}

			if !reflect.DeepEqual(got, sortedCounts) {
				t.Errorf("MostCommon() does not match with %v", tt.sortedCountsFilepath)
			}
		})
	}
}
