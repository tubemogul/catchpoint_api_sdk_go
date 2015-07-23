package alertsAPI

import (
	"encoding/xml"
	"fmt"
	"io"
)

// ParseXMLRequest parse the xml data you send it and populate the Alert struct
// it belongs to
func (a *Alert) ParseXMLRequest(xmlStream *io.ReadCloser) error {

	dec := xml.NewDecoder(*xmlStream)

	if err := dec.Decode(a); err == io.EOF {
		// Case when hitting the EOF before decoding anything
		return fmt.Errorf("[INFO] Getting en empty object.")

	} else if err != nil {
		return err
	}

	// If the decoding went well...
	return nil
}
