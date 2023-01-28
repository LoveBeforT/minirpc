package serialize

type InfDefine interface {
	GetData() []byte
	SetData(data []byte)
	WriteString(string)
	GetInt() int
	WriteInt(int)
	GetInterface() interface{}
	WriteInterface(interface{})
	GetUint64() uint64
	WriteUint64(uint64)
	GetUint32() uint32
	WriteUint32(uint32)
	GetUint16() uint16
	WriteUint16(uint16)
	GetUint8() uint8
	WriteUint8(uint8)
}
