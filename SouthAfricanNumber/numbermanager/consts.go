package numbermanager

// errors constant to manage tests.
const (
	ErrorMissingPartialPrefix        = "some number are missing added prefix"
	ErrorWrongPrefix                 = "wrong prefix, replace with the correct one"
	ErrorCutExtraDigits              = "digit number more than wanted format, the exceeding was cut "
	ErrorNotNumericDigits            = "found not numeric digits, removed"
	ErrMsgLessThanPrefix      string = "digits are less than wanted prefix '%s'"
	ErrMsgLessThanCore        string = "digits are less than 'core' digits format '%v digits'"
)

// configuration parameters.
const (
	CoreLen     int    = 7
	prefixLen   int    = 4
	RightPrefix string = "2783"
)

// types enum.
type NumberType string

const (
	NotEvaluated      NumberType = "NotEvaluated"
	ValidFirstAttempt NumberType = "ValidFirstAttempt"
	InvalidCritical   NumberType = "InvalidCritical"
	InvalidButFixable NumberType = "InvalidButFixable"
)

const regexpNotNumber = "[^0-9]+"
