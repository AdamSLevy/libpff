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

func (v ValueType) String() string {
	switch v {
	case ValueTypeUnspecified:
		return "ValueTypeUnspecified"
	case ValueTypeNull:
		return "ValueTypeNull"
	case ValueTypeInteger16bitSigned:
		return "ValueTypeInteger16bitSigned"
	case ValueTypeInteger32bitSigned:
		return "ValueTypeInteger32bitSigned"
	case ValueTypeFloat32bit:
		return "ValueTypeFloat32bit"
	case ValueTypeDouble64bit:
		return "ValueTypeDouble64bit"
	case ValueTypeCurrency:
		return "ValueTypeCurrency"
	case ValueTypeFloatingTime:
		return "ValueTypeFloatingTime"
	case ValueTypeError:
		return "ValueTypeError"
	case ValueTypeBoolean:
		return "ValueTypeBoolean"
	case ValueTypeObject:
		return "ValueTypeObject"
	case ValueTypeInteger_64bit_Signed:
		return "ValueTypeInteger_64bit_Signed"
	case ValueTypeStringASCII:
		return "ValueTypeStringASCII"
	case ValueTypeStringUnicode:
		return "ValueTypeStringUnicode"
	case ValueTypeFiletime:
		return "ValueTypeFiletime"
	case ValueTypeGUID:
		return "ValueTypeGUID"
	case ValueTypeServerIdentifier:
		return "ValueTypeServerIdentifier"
	case ValueTypeRestriction:
		return "ValueTypeRestriction"
	case ValueTypeRuleAction:
		return "ValueTypeRuleAction"
	case ValueTypeBinaryData:
		return "ValueTypeBinaryData"
	case ValueTypeMultiValueInteger16bitSigned:
		return "ValueTypeMultiValueInteger16bitSigned"
	case ValueTypeMultiValueInteger32bitSigned:
		return "ValueTypeMultiValueInteger32bitSigned"
	case ValueTypeMultiValueFloat32bit:
		return "ValueTypeMultiValueFloat32bit"
	case ValueTypeMultiValueDouble64bit:
		return "ValueTypeMultiValueDouble64bit"
	case ValueTypeMultiValueCurrency:
		return "ValueTypeMultiValueCurrency"
	case ValueTypeMultiValueFloatingtime:
		return "ValueTypeMultiValueFloatingtime"
	case ValueTypeMultiValueInteger64bitSigned:
		return "ValueTypeMultiValueInteger64bitSigned"
	case ValueTypeMultiValueStringASCII:
		return "ValueTypeMultiValueStringASCII"
	case ValueTypeMultiValueStringUnicode:
		return "ValueTypeMultiValueStringUnicode"
	case ValueTypeMultiValueFiletime:
		return "ValueTypeMultiValueFiletime"
	case ValueTypeMultiValue_GUID:
		return "ValueTypeMultiValue_GUID"
	case ValueTypeMultiValueBinaryData:
		return "ValueTypeMultiValueBinaryData"
	default:
		return fmt.Sprintf("unrecognized ValueType: %v", uint32(v))
	}

}

//LIBPFF_ITEM_TYPE_UNDEFINED,
//LIBPFF_ITEM_TYPE_ACTIVITY,
//LIBPFF_ITEM_TYPE_APPOINTMENT,
//LIBPFF_ITEM_TYPE_ATTACHMENT,
//LIBPFF_ITEM_TYPE_ATTACHMENTS,
//LIBPFF_ITEM_TYPE_COMMON,
//LIBPFF_ITEM_TYPE_CONFIGURATION,
//LIBPFF_ITEM_TYPE_CONFLICT_MESSAGE,
//LIBPFF_ITEM_TYPE_CONTACT,
//LIBPFF_ITEM_TYPE_DISTRIBUTION_LIST,
//LIBPFF_ITEM_TYPE_DOCUMENT,
//LIBPFF_ITEM_TYPE_EMAIL,
//LIBPFF_ITEM_TYPE_EMAIL_SMIME,
//LIBPFF_ITEM_TYPE_FAX,
//LIBPFF_ITEM_TYPE_FOLDER,
//LIBPFF_ITEM_TYPE_MEETING,
//LIBPFF_ITEM_TYPE_MMS,
//LIBPFF_ITEM_TYPE_NOTE,
//LIBPFF_ITEM_TYPE_POSTING_NOTE,
//LIBPFF_ITEM_TYPE_RECIPIENTS,
//LIBPFF_ITEM_TYPE_RSS_FEED,
//LIBPFF_ITEM_TYPE_SHARING,
//LIBPFF_ITEM_TYPE_SMS,
//LIBPFF_ITEM_TYPE_SUB_ASSOCIATED_CONTENTS,
//LIBPFF_ITEM_TYPE_SUB_FOLDERS,
//LIBPFF_ITEM_TYPE_SUB_MESSAGES,
//LIBPFF_ITEM_TYPE_TASK,
//LIBPFF_ITEM_TYPE_TASK_REQUEST,
//LIBPFF_ITEM_TYPE_VOICEMAIL,
//LIBPFF_ITEM_TYPE_UNKNOWN
