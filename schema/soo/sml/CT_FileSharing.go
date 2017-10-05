// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_FileSharing struct {
	// Read Only Recommended
	ReadOnlyRecommendedAttr *bool
	// User Name
	UserNameAttr *string
	// Write Reservation Password
	ReservationPasswordAttr *string
	// Cryptographic Algorithm Name
	AlgorithmNameAttr *string
	// Password Hash Value
	HashValueAttr *string
	// Salt Value for Password Verifier
	SaltValueAttr *string
	// Iterations to Run Hashing Algorithm
	SpinCountAttr *uint32
}

func NewCT_FileSharing() *CT_FileSharing {
	ret := &CT_FileSharing{}
	return ret
}

func (m *CT_FileSharing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ReadOnlyRecommendedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "readOnlyRecommended"},
			Value: fmt.Sprintf("%d", b2i(*m.ReadOnlyRecommendedAttr))})
	}
	if m.UserNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "userName"},
			Value: fmt.Sprintf("%v", *m.UserNameAttr)})
	}
	if m.ReservationPasswordAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "reservationPassword"},
			Value: fmt.Sprintf("%v", *m.ReservationPasswordAttr)})
	}
	if m.AlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "algorithmName"},
			Value: fmt.Sprintf("%v", *m.AlgorithmNameAttr)})
	}
	if m.HashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hashValue"},
			Value: fmt.Sprintf("%v", *m.HashValueAttr)})
	}
	if m.SaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saltValue"},
			Value: fmt.Sprintf("%v", *m.SaltValueAttr)})
	}
	if m.SpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spinCount"},
			Value: fmt.Sprintf("%v", *m.SpinCountAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FileSharing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "readOnlyRecommended" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ReadOnlyRecommendedAttr = &parsed
			continue
		}
		if attr.Name.Local == "userName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UserNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "reservationPassword" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ReservationPasswordAttr = &parsed
			continue
		}
		if attr.Name.Local == "algorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgorithmNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "hashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "saltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "spinCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpinCountAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FileSharing: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FileSharing and its children
func (m *CT_FileSharing) Validate() error {
	return m.ValidateWithPath("CT_FileSharing")
}

// ValidateWithPath validates the CT_FileSharing and its children, prefixing error messages with path
func (m *CT_FileSharing) ValidateWithPath(path string) error {
	return nil
}
