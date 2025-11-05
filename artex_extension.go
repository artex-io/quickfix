package quickfix

import "fmt"

//------------------------ message ------------------------------------------------------

// GetFields returns fields as they appear in the raw message.
func (m *Message) GetFields() []TagValue { return m.fields }

//------------------------ message end ------------------------------------------------------

// ------------------------ field_map -----------------------------------------------------
func (tv TagValue) Value() string {
	return string(tv.value)
}
func (tv TagValue) Tag() Tag {
	return tv.tag
}

//------------------------ field_map end -----------------------------------------------------

// ----------------------- Registry ----------------------------------------------------------

// SendReject is a helper function which allows to send an error outside of the
// quickfix application methods. Useful when doing asynchronous work.
func SendReject(m *Message, sessionID SessionID, rej MessageRejectError) error {
	session, ok := lookupSession(sessionID)
	if !ok {
		return errUnknownSession
	}
	return session.doReject(m, rej)
}

// ----------------------- Registry end----------------------------------------------------------

//-------------------------- errors ----------------------------------------------------------

// ValueIsIncorrectWithValue returns an error indicating a field with value that is not valid.
func ValueIsIncorrectWithValue(tag Tag, value string) MessageRejectError {
	return NewMessageRejectError(fmt.Sprintf("Value (%s) is incorrect for this tag (%d)", value, tag), rejectReasonValueIsIncorrect, &tag)
}

//-------------------------- errors end ----------------------------------------------------------
