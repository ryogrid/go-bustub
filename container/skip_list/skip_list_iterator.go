// TODO: (SDB) not implemented yet skip_list_iterator.go

package skip_list

import (
	"github.com/ryogrid/SamehadaDB/storage/page/skip_list_page"
	"math"
	"unsafe"

	"github.com/ryogrid/SamehadaDB/storage/buffer"
	"github.com/ryogrid/SamehadaDB/types"
)

type SkipListIteratorOnMem struct {
	//bpm        *buffer.BufferPoolManager
	//headerPage *skip_list_page.SkipListHeaderPage
	//bucket     uint32
	//offset     uint32
	//blockId    types.PageID
	//blockPage  *skip_list_page.SkipListBlockPage
	curNode       *SkipListOnMem
	rangeStartKey *types.Value
	rangeEndKey   *types.Value
}

type skipListIterator struct {
	bpm        *buffer.BufferPoolManager
	headerPage *skip_list_page.SkipListHeaderPage
	bucket     uint32
	offset     uint32
	blockId    types.PageID
	blockPage  *skip_list_page.SkipListBlockPage
}

func newSkipListIterator(bpm *buffer.BufferPoolManager, header *skip_list_page.SkipListHeaderPage, bucket uint32, offset uint32) *skipListIterator {
	blockPageId := header.GetBlockPageId(bucket)

	bPageData := bpm.FetchPage(blockPageId).Data()
	blockPage := (*skip_list_page.SkipListBlockPage)(unsafe.Pointer(bPageData))

	return &skipListIterator{bpm, header, bucket, offset, blockPageId, blockPage}
}

func (itr *SkipListIteratorOnMem) Next() (done bool, err error, key *types.Value, val uint32) {
	itr.curNode = itr.curNode.Forward[0]
	if (itr.rangeEndKey != nil && itr.curNode.Key.CompareGreaterThan(*itr.rangeEndKey)) || itr.curNode.Key.IsInf() {
		return true, nil, nil, math.MaxUint32
	}

	return false, nil, itr.curNode.Key, itr.curNode.Val
}

func (itr *skipListIterator) next() {
	itr.offset++
	// the current block page is full, we need to go to the next one
	if itr.offset >= skip_list_page.BlockArraySize {
		itr.bucket += 1
		itr.offset = 0

		// we need to go to the first block
		if itr.bucket >= itr.headerPage.NumBlocks() {
			itr.bucket = 0
		}

		itr.bpm.UnpinPage(itr.blockId, true)
		itr.blockId = itr.headerPage.GetBlockPageId(itr.bucket)

		bPageData := itr.bpm.FetchPage(itr.blockId).Data()
		itr.blockPage = (*skip_list_page.SkipListBlockPage)(unsafe.Pointer(bPageData))
	}
}