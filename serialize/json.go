package serialize

type jsonImpl struct {
	data []byte
}

func NewJsonSerializer() InfDefine {
	return &jsonImpl{}
}

func (j *jsonImpl) GetData() []byte {
	return j.data
}

func (j *jsonImpl) SetData(data []byte) {
	j.data = data
}

func (j *jsonImpl) WriteString(string) {

}

func (j *jsonImpl) GetInt() int {
	return 0
}

func (j *jsonImpl) WriteInt(int) {

}

func (j *jsonImpl) GetInterface() interface{} {
	return nil
}

func (j *jsonImpl) WriteInterface(interface{}) {

}

func (j *jsonImpl) GetUint64() uint64 {
	return 0
}

func (j *jsonImpl) WriteUint64(uint64) {

}

func (j *jsonImpl) GetUint32() uint32 {
	return 0
}

func (j *jsonImpl) WriteUint32(uint32) {

}

func (j *jsonImpl) GetUint16() uint16 {
	return 0
}

func (j *jsonImpl) WriteUint16(uint16) {

}

func (j *jsonImpl) GetUint8() uint8 {
	return 0
}

func (j *jsonImpl) WriteUint8(uint8) {

}
