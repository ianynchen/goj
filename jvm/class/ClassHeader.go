package class

import "errors"

const (
	CONSTANT_Class              uint16 = 7
	CONSTANT_Fieldref           uint16 = 9
	CONSTANT_Methodref          uint16 = 10
	CONSTANT_InterfaceMethodref uint16 = 11
	CONSTANT_String             uint16 = 8
	CONSTANT_Integer            uint16 = 3
	CONSTANT_Float              uint16 = 4
	CONSTANT_Long               uint16 = 5
	CONSTANT_Double             uint16 = 6
	CONSTANT_NameAndType        uint16 = 12
	CONSTANT_Utf8               uint16 = 1
	CONSTANT_MethodHandle       uint16 = 15
	CONSTANT_MethodType         uint16 = 16
	CONSTANT_InvokeDynamic      uint16 = 18
)

const (
	ACC_PUBLIC     uint16 = 0x0001 //Declared public; may be accessed from outside its package.
	ACC_FINAL      uint16 = 0x0010 //Declared final; no subclasses allowed.
	ACC_SUPER      uint16 = 0x0020 //Treat superclass methods specially when invoked by the invokespecial instruction.
	ACC_INTERFACE  uint16 = 0x0200 //Is an interface, not a class.
	ACC_ABSTRACT   uint16 = 0x0400 //Declared abstract; must not be instantiated.
	ACC_SYNTHETIC  uint16 = 0x1000 //Declared synthetic; not present in the source code.
	ACC_ANNOTATION uint16 = 0x2000 //Declared as an annotation type.
	ACC_ENUM       uint16 = 0x4000 //Declared as an enum type.
)

type ConstantPoolInfo struct {
	Tag  uint16
	Info []uint16
}

type FieldInfo struct {
}

type MethodInfo struct {
}

type AttributeInfo struct {
}

type ClassHeader struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPoolInfo  []ConstantPoolInfo
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []uint16
	FieldsCount       uint16
	Fields            []FieldInfo
	MethodCount       uint16
	Methods           []MethodInfo
	AttributesCount   uint16
	Attributes        []AttributeInfo
}

func (this *ClassHeader) IsValid() bool {

	return this.Magic == 0xCAFEBABE &&
		this.ConstantPoolCount >= 0 &&
		this.MethodCount >= 0 &&
		this.InterfacesCount >= 0 &&
		this.FieldsCount >= 0 &&
		this.AttributesCount >= 0
}

func (this *ClassHeader) Initialize() (err error) {

	if this.Magic != 0xCAFEBABE {
		return errors.New("Magic number no match")
	}

	if this.ConstantPoolCount < 0 {
		return errors.New("Constant pool size illegal")
	}

	if this.MethodCount < 0 {
		return errors.New("Method count illegal")
	}

	if this.InterfacesCount < 0 {
		return errors.New("Interface count illegal")
	}

	if this.FieldsCount < 0 {
		return errors.New("Field count illegal")
	}

	if this.AttributesCount < 0 {
		return errors.New("Attribute count illegal")
	}

	this.ConstantPoolInfo = make([]ConstantPoolInfo, this.ConstantPoolCount)
	this.Methods = make([]MethodInfo, this.MethodCount)
	this.Interfaces = make([]uint16, this.InterfacesCount)
	this.Fields = make([]FieldInfo, this.FieldsCount)
	this.Attributes = make([]AttributeInfo, this.AttributesCount)

	return err
}
