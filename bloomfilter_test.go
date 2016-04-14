package bloomfilter

import "testing"

type TestData struct {
	word       string
	isInFilter bool
}

func TestBloomFilter(t *testing.T) {
	size, numHashes := 1000, 4
	bloom := New(size, numHashes)

	words := []string{"foo", "bar", "foobar"}
	for _, word := range words {
		bloom.Add([]byte(word))
	}

	tests := []TestData{
		TestData{"foo", true},
		TestData{"bar", true},
		TestData{"foobar", true},
		TestData{"not_there", false},
	}

	for _, data := range tests {
		actual := bloom.MayContain([]byte(data.word))
		if actual != data.isInFilter {
			t.Errorf("MayContain(%s) should return %v, got %v",
				data.word, data.isInFilter, actual)
		}
	}
}
