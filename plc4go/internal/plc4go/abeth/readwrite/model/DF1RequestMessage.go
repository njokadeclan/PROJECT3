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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type DF1RequestMessage struct {
	DestinationAddress uint8
	SourceAddress      uint8
	Status             uint8
	TransactionCounter uint16
	Child              IDF1RequestMessageChild
}

// The corresponding interface
type IDF1RequestMessage interface {
	CommandCode() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IDF1RequestMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IDF1RequestMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type IDF1RequestMessageChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *DF1RequestMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16)
	GetTypeName() string
	IDF1RequestMessage
}

func NewDF1RequestMessage(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) *DF1RequestMessage {
	return &DF1RequestMessage{DestinationAddress: destinationAddress, SourceAddress: sourceAddress, Status: status, TransactionCounter: transactionCounter}
}

func CastDF1RequestMessage(structType interface{}) *DF1RequestMessage {
	castFunc := func(typ interface{}) *DF1RequestMessage {
		if casted, ok := typ.(DF1RequestMessage); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1RequestMessage); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1RequestMessage) GetTypeName() string {
	return "DF1RequestMessage"
}

func (m *DF1RequestMessage) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DF1RequestMessage) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *DF1RequestMessage) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *DF1RequestMessage) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DF1RequestMessageParse(readBuffer utils.ReadBuffer) (*DF1RequestMessage, error) {
	if pullErr := readBuffer.PullContext("DF1RequestMessage"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (destinationAddress)
	destinationAddress, _destinationAddressErr := readBuffer.ReadUint8("destinationAddress", 8)
	if _destinationAddressErr != nil {
		return nil, errors.Wrap(_destinationAddressErr, "Error parsing 'destinationAddress' field")
	}

	// Simple Field (sourceAddress)
	sourceAddress, _sourceAddressErr := readBuffer.ReadUint8("sourceAddress", 8)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode, _commandCodeErr := readBuffer.ReadUint8("commandCode", 8)
	if _commandCodeErr != nil {
		return nil, errors.Wrap(_commandCodeErr, "Error parsing 'commandCode' field")
	}

	// Simple Field (status)
	status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Simple Field (transactionCounter)
	transactionCounter, _transactionCounterErr := readBuffer.ReadUint16("transactionCounter", 16)
	if _transactionCounterErr != nil {
		return nil, errors.Wrap(_transactionCounterErr, "Error parsing 'transactionCounter' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *DF1RequestMessage
	var typeSwitchError error
	switch {
	case commandCode == 0x0F: // DF1CommandRequestMessage
		_parent, typeSwitchError = DF1CommandRequestMessageParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("DF1RequestMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, destinationAddress, sourceAddress, status, transactionCounter)
	return _parent, nil
}

func (m *DF1RequestMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *DF1RequestMessage) SerializeParent(writeBuffer utils.WriteBuffer, child IDF1RequestMessage, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("DF1RequestMessage"); pushErr != nil {
		return pushErr
	}

	// Simple Field (destinationAddress)
	destinationAddress := uint8(m.DestinationAddress)
	_destinationAddressErr := writeBuffer.WriteUint8("destinationAddress", 8, (destinationAddress))
	if _destinationAddressErr != nil {
		return errors.Wrap(_destinationAddressErr, "Error serializing 'destinationAddress' field")
	}

	// Simple Field (sourceAddress)
	sourceAddress := uint8(m.SourceAddress)
	_sourceAddressErr := writeBuffer.WriteUint8("sourceAddress", 8, (sourceAddress))
	if _sourceAddressErr != nil {
		return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(child.CommandCode())
	_commandCodeErr := writeBuffer.WriteUint8("commandCode", 8, (commandCode))

	if _commandCodeErr != nil {
		return errors.Wrap(_commandCodeErr, "Error serializing 'commandCode' field")
	}

	// Simple Field (status)
	status := uint8(m.Status)
	_statusErr := writeBuffer.WriteUint8("status", 8, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Simple Field (transactionCounter)
	transactionCounter := uint16(m.TransactionCounter)
	_transactionCounterErr := writeBuffer.WriteUint16("transactionCounter", 16, (transactionCounter))
	if _transactionCounterErr != nil {
		return errors.Wrap(_transactionCounterErr, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1RequestMessage"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DF1RequestMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
