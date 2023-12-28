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
type COTPParameterCalledTsap struct {
	TsapId uint16
	Parent *COTPParameter
}

// The corresponding interface
type ICOTPParameterCalledTsap interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPParameterCalledTsap) ParameterType() uint8 {
	return 0xC2
}

func (m *COTPParameterCalledTsap) InitializeParent(parent *COTPParameter) {
}

func NewCOTPParameterCalledTsap(tsapId uint16) *COTPParameter {
	child := &COTPParameterCalledTsap{
		TsapId: tsapId,
		Parent: NewCOTPParameter(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCOTPParameterCalledTsap(structType interface{}) *COTPParameterCalledTsap {
	castFunc := func(typ interface{}) *COTPParameterCalledTsap {
		if casted, ok := typ.(COTPParameterCalledTsap); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPParameterCalledTsap); ok {
			return casted
		}
		if casted, ok := typ.(COTPParameter); ok {
			return CastCOTPParameterCalledTsap(casted.Child)
		}
		if casted, ok := typ.(*COTPParameter); ok {
			return CastCOTPParameterCalledTsap(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPParameterCalledTsap) GetTypeName() string {
	return "COTPParameterCalledTsap"
}

func (m *COTPParameterCalledTsap) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *COTPParameterCalledTsap) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (tsapId)
	lengthInBits += 16

	return lengthInBits
}

func (m *COTPParameterCalledTsap) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPParameterCalledTsapParse(readBuffer utils.ReadBuffer) (*COTPParameter, error) {
	if pullErr := readBuffer.PullContext("COTPParameterCalledTsap"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (tsapId)
	tsapId, _tsapIdErr := readBuffer.ReadUint16("tsapId", 16)
	if _tsapIdErr != nil {
		return nil, errors.Wrap(_tsapIdErr, "Error parsing 'tsapId' field")
	}

	if closeErr := readBuffer.CloseContext("COTPParameterCalledTsap"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPParameterCalledTsap{
		TsapId: tsapId,
		Parent: &COTPParameter{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *COTPParameterCalledTsap) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterCalledTsap"); pushErr != nil {
			return pushErr
		}

		// Simple Field (tsapId)
		tsapId := uint16(m.TsapId)
		_tsapIdErr := writeBuffer.WriteUint16("tsapId", 16, (tsapId))
		if _tsapIdErr != nil {
			return errors.Wrap(_tsapIdErr, "Error serializing 'tsapId' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterCalledTsap"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPParameterCalledTsap) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
