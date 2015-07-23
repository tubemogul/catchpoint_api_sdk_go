package catchpoint

import (
	"bytes"
	"fmt"
	"io"
	//	"io/ioutil"
	// "log"
	//	"net/http"
)

// Those correspond to the return code of the nagions NSCA
const (
	OK = iota
	WARNING
	CRITICAL
	UNKNOWN
)

// RequestHandler is the entry point for parsing the incoming http requests no
// matter what their format is (provided it's XML, orj Json in a future release)
// Returns:
// - uint8: the criticity level of the alert (translated NSCA standards)
// - *string: the name of the alert
// - *[]string: the list of the failures
// - error: any potential error encountered during the processing of the request
func (a *Alert) RequestHandler(r *io.ReadCloser) (uint8, *string, *[]string, error) {

	//  ctype, _ := http.DetectContentType(ioutil.ReadAll(r))
	//  if *verbose {
	//  	log.Printf("Processing content type: %s", ctype)
	//  }

	// val, _ := ioutil.ReadAll(*r)
	// ctype := http.DetectContentType(val)
	// log.Printf("DEBUG: Request body = %v ", string(val[:]))
	// log.Printf("DEBUG: Content type = %v ", ctype)

	err := a.ParseXMLRequest(r)
	if err != nil {
		return uint8(CRITICAL), nil, nil, err
	} else {
		return a.PrettyPrintMessage()
	}
}

// getAlertReturnCode translates the message criticity level to an alert leve
// (replacing the INFO level by OK and other unsupported by UNKONWN)
func (a *Alert) getAlertReturnCode(level uint8) uint8 {
	switch level {
	case 0:
		return WARNING
	case 1:
		return CRITICAL
	case 2: // it is INFO level on Catchpoint considered as a OK return code
		return OK
	case 3:
		return OK
	default:
		return UNKNOWN
	}
}

// PrettyPrintMessage retrieve the informations from the current struct and
// returns:
// - int16: the criticity level of the alert (translated NSCA standards)
// - *string: the name of the alert (Product + test name concatenated by a dash)
// - *[]string: the list of the failures
// - error: any potential error encountered during the processing of the request
func (a *Alert) PrettyPrintMessage() (uint8, *string, *[]string, error) {
	var returnCode uint8
	var alerts []string

	fullName := fmt.Sprintf("%s-%s", a.TestDetail.ProductName, a.TestDetail.Name)

	returnCode = a.getAlertReturnCode(a.NotificationLevelId)

	for _, v := range a.Condition.Nodes {
		var buffer bytes.Buffer

		buffer.WriteString(AlertTypeIdLabel[a.Setting.AlertTypeId])

		// Only print the alert subtype when set
		if a.Setting.AlertSubTypeId != 0 {
			buffer.WriteString("-")
			buffer.WriteString(AlertSubTypeIdLabel[a.Setting.AlertSubTypeId])
		}

		// Only display the exit code if provided
		if v.PageFailure.ErrorCode != 0 {
			buffer.WriteString(" Exit code: ")
			buffer.WriteString(fmt.Sprint(v.PageFailure.ErrorCode))
		}

		// Only display the HTTP return code if provided
		if v.PageFailure.HttpStatusCode != 0 {
			buffer.WriteString(" HTTP code ")
			buffer.WriteString(fmt.Sprint(v.PageFailure.HttpStatusCode))
		}

		buffer.WriteString(" failed from \"")
		buffer.WriteString(v.Name)
		buffer.WriteString("\" (")
		buffer.WriteString(v.IpAddress)
		buffer.WriteString(") to ")
		buffer.WriteString(v.RemoteIpAddress)
		buffer.WriteString(" at ")
		buffer.WriteString(a.Timestamp.ProcessingUtc)
		buffer.WriteString(" UTC")
		alerts = append(alerts, buffer.String())
	}

	return returnCode, &fullName, &alerts, nil
}
