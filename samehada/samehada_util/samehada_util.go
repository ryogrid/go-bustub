package samehada_util

import (
	"bytes"
	"encoding/binary"
	"github.com/ryogrid/SamehadaDB/storage/page"
	"github.com/ryogrid/SamehadaDB/types"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func PackRIDtoUint32(value *page.RID) uint32 {
	buf1 := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)
	pack_buf := make([]byte, 4)
	binary.Write(buf1, binary.LittleEndian, value.PageId)
	binary.Write(buf2, binary.LittleEndian, value.SlotNum)
	pageIdInBytes := buf1.Bytes()
	slotNumInBytes := buf2.Bytes()
	copy(pack_buf[:2], pageIdInBytes[:2])
	copy(pack_buf[2:], slotNumInBytes[:2])
	return binary.LittleEndian.Uint32(pack_buf)
}

func UnpackUint32toRID(value uint32) page.RID {
	packed_buf := new(bytes.Buffer)
	binary.Write(packed_buf, binary.LittleEndian, value)
	packedDataInBytes := packed_buf.Bytes()
	var PageId types.PageID
	var SlotNum uint32
	buf := make([]byte, 4)
	copy(buf[:2], packedDataInBytes[:2])
	PageId = types.PageID(binary.LittleEndian.Uint32(buf))
	copy(buf[:2], packedDataInBytes[2:])
	SlotNum = binary.LittleEndian.Uint32(buf)
	ret := new(page.RID)
	ret.PageId = PageId
	ret.SlotNum = SlotNum
	return *ret
}

func GetPonterOfValue(value types.Value) *types.Value {
	val := value
	return &val
}