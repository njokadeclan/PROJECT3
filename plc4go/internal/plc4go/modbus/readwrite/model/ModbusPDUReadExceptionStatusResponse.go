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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ModbusPDUReadExceptionStatusResponse struct {
	Value  uint8
	Parent *ModbusPDU
}

// The corresponding interface
type IModbusPDUReadExceptionStatusResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadExceptionStatusResponse) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUReadExceptionStatusResponse) FunctionFlag() uint8 {
	return 0x07
}

func (m *ModbusPDUReadExceptionStatusResponse) Response() bool {
	return true
}

func (m *ModbusPDUReadExceptionStatusResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReadExceptionStatusResponse(value uint8) *ModbusPDU {
	child := &ModbusPDUReadExceptionStatusResponse{
		Value:  value,
		Parent: NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUReadExceptionStatusResponse(structType interface{}) *ModbusPDUReadExceptionStatusResponse {
	castFunc := func(typ interface{}) *ModbusPDUReadExceptionStatusResponse {
		if casted, ok := typ.(ModbusPDUReadExceptionStatusResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadExceptionStatusResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadExceptionStatusResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadExceptionStatusResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadExceptionStatusResponse) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusResponse"
}

func (m *ModbusPDUReadExceptionStatusResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReadExceptionStatusResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *ModbusPDUReadExceptionStatusResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadExceptionStatusResponseParse(readBuffer utils.ReadBuffer) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	value, _valueErr := readBuffer.ReadUint8("value", 8)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadExceptionStatusResponse{
		Value:  value,
		Parent: &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUReadExceptionStatusResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := uint8(m.Value)
		_valueErr := writeBuffer.WriteUint8("value", 8, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadExceptionStatusResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}