package bloomfilter

import "github.com/spaolacci/murmur3"

// BloomFilter provides a structure which implements a bloom filter.
type BloomFilter struct {
	numHashes int
	bitSet    []bool
}

// New returns an instance of BloomFilter initialized with provided parameters.
func New(size int, numHashes int) *BloomFilter {
	return &BloomFilter{
		numHashes: numHashes,
		bitSet:    make([]bool, size),
	}
}

func hash(data []byte) (uint64, uint64) {
	return murmur3.Sum128(data)
}

func nthHash(n int, h1 uint64, h2 uint64, bfSize int) uint64 {
	return (h1 + uint64(n)*h2) % uint64(bfSize)
}

// Add inserts data to the bloom filter.
func (bf *BloomFilter) Add(data []byte) {
	h1, h2 := hash(data)
	filterSize := len(bf.bitSet)
	for n := 0; n < bf.numHashes; n++ {
		i := nthHash(n, h1, h2, filterSize)
		bf.bitSet[i] = true
	}
}

// MayContain indicates if provided data has already been set in the bloom
// filter. Returns false if data is not in the filter, true is data may be in
// the filter.
func (bf *BloomFilter) MayContain(data []byte) bool {
	h1, h2 := hash(data)
	filterSize := len(bf.bitSet)
	for n := 0; n < bf.numHashes; n++ {
		i := nthHash(n, h1, h2, filterSize)
		if !bf.bitSet[i] {
			return false
		}
	}

	return true
}
