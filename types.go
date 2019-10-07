package libpff

// #include <libpff.h>
import "C"

type ValueType uint32

const (
	ValueTypeUnspecified                  ValueType = C.LIBPFF_VALUE_TYPE_UNSPECIFIED
	ValueTypeNull                                   = C.LIBPFF_VALUE_TYPE_NULL
	ValueTypeInteger16bitSigned                     = C.LIBPFF_VALUE_TYPE_INTEGER_16BIT_SIGNED
	ValueTypeInteger32bitSigned                     = C.LIBPFF_VALUE_TYPE_INTEGER_32BIT_SIGNED
	ValueTypeFloat32bit                             = C.LIBPFF_VALUE_TYPE_FLOAT_32BIT
	ValueTypeDouble64bit                            = C.LIBPFF_VALUE_TYPE_DOUBLE_64BIT
	ValueTypeCurrency                               = C.LIBPFF_VALUE_TYPE_CURRENCY
	ValueTypeFloatingTime                           = C.LIBPFF_VALUE_TYPE_FLOATINGTIME
	ValueTypeError                                  = C.LIBPFF_VALUE_TYPE_ERROR
	ValueTypeBoolean                                = C.LIBPFF_VALUE_TYPE_BOOLEAN
	ValueTypeObject                                 = C.LIBPFF_VALUE_TYPE_OBJECT
	ValueTypeInteger_64bit_Signed                   = C.LIBPFF_VALUE_TYPE_INTEGER_64BIT_SIGNED
	ValueTypeStringASCII                            = C.LIBPFF_VALUE_TYPE_STRING_ASCII
	ValueTypeStringUnicode                          = C.LIBPFF_VALUE_TYPE_STRING_UNICODE
	ValueTypeFiletime                               = C.LIBPFF_VALUE_TYPE_FILETIME
	ValueTypeGUID                                   = C.LIBPFF_VALUE_TYPE_GUID
	ValueTypeServerIdentifier                       = C.LIBPFF_VALUE_TYPE_SERVER_IDENTIFIER
	ValueTypeRestriction                            = C.LIBPFF_VALUE_TYPE_RESTRICTION
	ValueTypeRuleAction                             = C.LIBPFF_VALUE_TYPE_RULE_ACTION
	ValueTypeBinaryData                             = C.LIBPFF_VALUE_TYPE_BINARY_DATA
	ValueTypeMultiValueInteger16bitSigned           = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_INTEGER_16BIT_SIGNED
	ValueTypeMultiValueInteger32bitSigned           = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_INTEGER_32BIT_SIGNED
	ValueTypeMultiValueFloat32bit                   = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_FLOAT_32BIT
	ValueTypeMultiValueDouble64bit                  = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_DOUBLE_64BIT
	ValueTypeMultiValueCurrency                     = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_CURRENCY
	ValueTypeMultiValueFloatingtime                 = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_FLOATINGTIME
	ValueTypeMultiValueInteger64bitSigned           = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_INTEGER_64BIT_SIGNED
	ValueTypeMultiValueStringASCII                  = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_STRING_ASCII
	ValueTypeMultiValueStringUnicode                = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_STRING_UNICODE
	ValueTypeMultiValueFiletime                     = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_FILETIME
	ValueTypeMultiValue_GUID                        = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_GUID
	ValueTypeMultiValueBinaryData                   = C.LIBPFF_VALUE_TYPE_MULTI_VALUE_BINARY_DATA
)
