package counter

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func loadCounts(filepath string) (map[string]int, error) {
	fp, err := os.Open("./Data/counts.txt")
	if err != nil {
		return nil, err
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
			return nil, err
		}

		counts[key] = freq
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return counts, nil
}

func createCountsFile(outputFilepath string, keys []string, freqs []int) error {
	fp, err := os.Create(outputFilepath)
	if err != nil {
		return err
	}
	defer fp.Close()

	for i := 0; i < len(keys); i++ {
		fp.WriteString(fmt.Sprintf("%v %v\n", freqs[i], keys[i]))
	}

	return nil
}

func getFileMD5Hash(inputFilepath string) (string, error) {
	bs, err := ioutil.ReadFile(inputFilepath)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write(bs)
	v := hash.Sum(nil)

	return string(v), nil
}

func innerTestCounter(t *testing.T, c *Counter, counts map[string]int, countsOutputFilepath string) {
	if c.Len() == len(counts) {
		t.Log("PASS: Len()")
	} else {
		t.Error("ERROR: Len()")
	}

	if c.Contains("yH") {
		t.Log("PASS: Contains()")
	} else {
		t.Error("ERROR: Contains()")
	}

	if c.Count("yH") == counts["yH"] {
		t.Log("PASS: Count()")
	} else {
		t.Error("ERROR: Count()")
	}

	keys, freqs := c.MostCommon()
	if err := createCountsFile(countsOutputFilepath, keys, freqs); err != nil {
		t.Fatal(err)
	}

	correctHash, err := getFileMD5Hash("./Data/counts.txt")
	if err != nil {
		t.Fatal(err)
	}
	cHash, err := getFileMD5Hash(countsOutputFilepath)
	if err != nil {
		t.Fatal(err)
	}

	if cHash == correctHash {
		t.Log("PASS: MostCommon()")
	} else {
		t.Error("ERROR: MostCommon()")
	}
}

func TestCounter(t *testing.T) {
	t.Log("テストに使用するデータを読み込んでいます...")

	counts, err := loadCounts("./Data/counts.txt")
	if err != nil {
		t.Fatal(err)
	}

	bs, err := ioutil.ReadFile("./Data/random_strings.txt")
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(bs), "\n")
	lines = lines[:len(lines)-1]

	t.Log("Counterを作成しています...")
	counter := NewCounter(lines)
	counter2 := NewCounterFromMap(counts)

	t.Log("Counterのテストを行っています...")

	t.Log("Sliceから作成されたCounterのテストを行っています...")
	innerTestCounter(t, counter, counts, "./Data/counts_2.txt")

	t.Log("Mapから作成されたCounterのテストを行っています...")
	innerTestCounter(t, counter2, counts, "./Data/counts_3.txt")

	t.Log("テストが終了しました")
}
