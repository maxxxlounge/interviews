package numbermanager

// errors constant to manage tests.
const ErrorMissingPartialPrefix = "some number are missing added prefix"
const ErrorWrongPrefix = "wrong prefix, replace with the correct one"
const ErrorCutExtraDigits = "digit number more than wanted format, the exceeding was cut "
const ErrorNotNumericDigits = "found not numeric digits, removed"
const ErrMsgLessThanPrefix string = "digits are less than wanted prefix '%s'"
const ErrMsgLessThanCore string = "digits are less than 'core' digits format '%v digits'"

// configuration parameters.
const CoreLen int = 7
const prefixLen int = 4
const RightPrefix string = "2783"

// types enum.
type NumberType string

const ValidFirstAttempt NumberType = "ValidFirstAttempt"
const InvalidCritical NumberType = "InvalidCritical"
const InvalidButFixable NumberType = "InvalidButFixable"

const regexpNotNumber = "[^0-9]+"
