package define

import (
	"encoding/binary"
	"log"
	"strconv"
)

const (
	initDataLength   = uint64(1024) // 初始长度
	appendDataLength = uint64(1024) // 增长量
)

type DataBuffer struct {
	buf               []byte // 缓冲内存
	bufferLength      uint64 // 缓冲总长度
	currentReadPos    uint64 // 当前读取位置
	currentDataLength uint64 // 当前数据具体长度
}

func (buff *DataBuffer) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		panic(any("length of dest buffer can't be 0."))
	}
	n = copy(p, buff.buf[buff.currentReadPos:])
	buff.currentReadPos += uint64(n)
	return n, nil
}

func (buff *DataBuffer) Write(p []byte) (n int, err error) {
	varSize := uint64(len(p))
	if varSize+buff.currentDataLength > buff.bufferLength {
		appendSize := buff.currentDataLength + varSize - buff.bufferLength
		appendSize = TernaryConditionalOperator(appendSize > appendDataLength, appendSize, appendDataLength).(uint64)
		tmpBuff := make([]byte, buff.bufferLength+appendSize)
		copy(tmpBuff, buff.buf[:])
		buff.buf = tmpBuff
		buff.bufferLength = uint64(len(buff.buf))
	}
	copiedLen := copy(buff.buf[buff.currentDataLength:], p)
	buff.currentDataLength += uint64(copiedLen)
	return copiedLen, nil
}

func (buff *DataBuffer) Pack(in interface{}) *DataBuffer {
	if err := binary.Write(buff, binary.BigEndian, in); nil != err {
		log.Panic("Pack data failed:", err.Error())
	}
	return buff
}

func (buff *DataBuffer) Unpack(out interface{}) *DataBuffer {
	if err := binary.Read(buff, binary.BigEndian, out); nil != err {
		log.Panic("Unpack data failed:", err.Error())
	}
	return buff
}

func (buff *DataBuffer) ResetReadPos() {
	buff.currentReadPos = 0
	return
}

func (buff *DataBuffer) WriteByteArray(varArray []byte) *DataBuffer {
	_, _ = buff.Write(varArray)
	return buff
}

func (buff *DataBuffer) ReadByteArray(outArray []byte) *DataBuffer {
	_, _ = buff.Read(outArray)
	return buff
}

func (buff *DataBuffer) GetData() []byte {
	return buff.buf
}

func (buff *DataBuffer) GetDataLength() uint64 {
	return buff.currentDataLength
}

func (buff *DataBuffer) GetBufferLength() uint64 {
	return buff.bufferLength
}

func (buff *DataBuffer) GetReadPos() uint64 {
	return buff.currentReadPos
}

func (buff *DataBuffer) SetReadPos(pos uint64) {
	if pos < buff.currentReadPos {
		buff.currentReadPos = pos
	}
}

func (buff *DataBuffer) ClearData(startPos uint64, count uint64) {
	// 如果两个参数都为0， 则删除全部数据
	if startPos == 0 && count == 0 {
		buff.buf = nil
		buff.ResetReadPos()
		buff.currentDataLength = 0
		buff.bufferLength = 0
	}

	// 判断是否超出范围
	if startPos+count > buff.currentDataLength {
		panic(any("can't erase buffer out of range: (startPos+count=" + strconv.Itoa(int(startPos+count)) + ") > (currentLength=" + strconv.Itoa(int(buff.currentDataLength)) + ")"))
	}

	//删除指定范围数据
	buff.buf = append(buff.buf[:startPos], buff.buf[startPos+count:]...)
	buff.bufferLength = uint64(len(buff.buf))

	// 更新数据长度
	buff.currentDataLength -= count

	// 将当前读取位置重置
	buff.ResetReadPos()
}

// TernaryConditionalOperator 三目运算
func TernaryConditionalOperator(condition bool, trueExp, falseExp interface{}) interface{} {
	if condition {
		return trueExp
	}

	return falseExp
}
