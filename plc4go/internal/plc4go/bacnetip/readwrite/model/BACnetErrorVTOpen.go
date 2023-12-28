/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetErrorVTOpen struct {
	Parent *BACnetError
}

// The corresponding interface
type IBACnetErrorVTOpen interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetErrorVTOpen) ServiceChoice() uint8 {
	return 0x15
}

func (m *BACnetErrorVTOpen) InitializeParent(parent *BACnetError) {
}

func NewBACnetErrorVTOpen() *BACnetError {
	child := &BACnetErrorVTOpen{
		Parent: NewBACnetError(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetErrorVTOpen(structType interface{}) *BACnetErrorVTOpen {
	castFunc := func(typ interface{}) *BACnetErrorVTOpen {
		if casted, ok := typ.(BACnetErrorVTOpen); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetErrorVTOpen); ok {
			return casted
		}
		if casted, ok := typ.(BACnetError); ok {
			return CastBACnetErrorVTOpen(casted.Child)
		}
		if casted, ok := typ.(*BACnetError); ok {
			return CastBACnetErrorVTOpen(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetErrorVTOpen) GetTypeName() string {
	return "BACnetErrorVTOpen"
}

func (m *BACnetErrorVTOpen) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetErrorVTOpen) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorVTOpen) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorVTOpenParse(readBuffer utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := readBuffer.PullContext("BACnetErrorVTOpen"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetErrorVTOpen"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorVTOpen{
		Parent: &BACnetError{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetErrorVTOpen) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorVTOpen"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetErrorVTOpen"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetErrorVTOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}