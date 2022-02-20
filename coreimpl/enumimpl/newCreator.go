package enumimpl

type newCreator struct {
	BasicByte   newBasicByteCreator
	BasicInt8   newBasicInt8Creator
	BasicInt16  newBasicInt16Creator
	BasicInt32  newBasicInt32Creator
	BasicString newBasicStringCreator
	BasicUInt16 newBasicUInt16Creator
}
