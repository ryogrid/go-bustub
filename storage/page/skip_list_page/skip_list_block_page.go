package skip_list_page

import "github.com/ryogrid/SamehadaDB/types"

type SkipListPair struct {
	key   uint32
	value uint32
}

const sizeOfHashTablePair = 16
const BlockArraySize = 4 * 4096 / (4*sizeOfHashTablePair + 1)

// TODO: (SDB) SkipList should be able to be store variable length size key data
/**
 * Store indexed key and value together within block page. Supports
 * non-unique keys.
 *
 * Block page format (keys are stored in order):
 *  ----------------------------------------------------------------
 * | KEY(1) + VALUE(1) | KEY(2) + VALUE(2) | ... | KEY(n) + VALUE(n)
 *  ----------------------------------------------------------------
 *
 *  Here '+' means concatenation.
 *
 */

// TODO: (SDB) not implemented yet skip_list_block_page.go

type SkipListBlockPageOnMem struct {
	//occuppied [(BlockArraySize-1)/8 + 1]byte // 256 bits
	//readable  [(BlockArraySize-1)/8 + 1]byte // 256 bits
	//array     [BlockArraySize]SkipListPair   // 252 * 16 bits
	key   []byte
	value uint32
}

type SkipListBlockPage struct {
	occuppied [(BlockArraySize-1)/8 + 1]byte // 256 bits
	readable  [(BlockArraySize-1)/8 + 1]byte // 256 bits
	array     [BlockArraySize]SkipListPair   // 252 * 16 bits
}

// Gets the key at an index in the block
func (page_ *SkipListBlockPage) KeyAt(index uint32) *types.Value {
	//return page_.array[index].key
	return nil
}

// Gets the value at an index in the block
func (page_ *SkipListBlockPage) ValueAt(index uint32) uint32 {
	return page_.array[index].value
}

// Attempts to insert a key and value into an index in the baccess.
func (page_ *SkipListBlockPage) Insert(index uint32, key uint32, value uint32) bool {
	if page_.IsOccupied(index) {
		return false
	}

	page_.array[index] = SkipListPair{key, value}
	page_.occuppied[index/8] |= (1 << (index % 8))
	page_.readable[index/8] |= (1 << (index % 8))
	return true
}

func (page_ *SkipListBlockPage) Remove(index uint32) {
	if !page_.IsReadable(index) {
		return
	}

	page_.readable[index/8] &= ^(1 << (index % 8))
}

// Returns whether or not an index is occuppied (valid key/value pair)
func (page_ *SkipListBlockPage) IsOccupied(index uint32) bool {
	return (page_.occuppied[index/8] & (1 << (index % 8))) != 0
}

// Returns whether or not an index is readable (valid key/value pair)
func (page_ *SkipListBlockPage) IsReadable(index uint32) bool {
	return (page_.readable[index/8] & (1 << (index % 8))) != 0
}