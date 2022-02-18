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
type BACnetTagPayloadUnsignedInteger struct {
	ValueUint8  *uint8
	ValueUint16 *uint16
	ValueUint24 *uint32
	ValueUint32 *uint32

	// Arguments.
	ActualLength uint32
}

// The corresponding interface
type IBACnetTagPayloadUnsignedInteger interface {
	// GetValueUint8 returns ValueUint8
	GetValueUint8() *uint8
	// GetValueUint16 returns ValueUint16
	GetValueUint16() *uint16
	// GetValueUint24 returns ValueUint24
	GetValueUint24() *uint32
	// GetValueUint32 returns ValueUint32
	GetValueUint32() *uint32
	// GetIsUint8 returns IsUint8
	GetIsUint8() bool
	// GetIsUint16 returns IsUint16
	GetIsUint16() bool
	// GetIsUint24 returns IsUint24
	GetIsUint24() bool
	// GetIsUint32 returns IsUint32
	GetIsUint32() bool
	// GetActualValue returns ActualValue
	GetActualValue() uint32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadUnsignedInteger) GetValueUint8() *uint8 {
	return m.ValueUint8
}

func (m *BACnetTagPayloadUnsignedInteger) GetValueUint16() *uint16 {
	return m.ValueUint16
}

func (m *BACnetTagPayloadUnsignedInteger) GetValueUint24() *uint32 {
	return m.ValueUint24
}

func (m *BACnetTagPayloadUnsignedInteger) GetValueUint32() *uint32 {
	return m.ValueUint32
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadUnsignedInteger) GetIsUint8() bool {
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	return bool((m.ActualLength) == (1))
}

func (m *BACnetTagPayloadUnsignedInteger) GetIsUint16() bool {
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	return bool((m.ActualLength) == (2))
}

func (m *BACnetTagPayloadUnsignedInteger) GetIsUint24() bool {
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	return bool((m.ActualLength) == (3))
}

func (m *BACnetTagPayloadUnsignedInteger) GetIsUint32() bool {
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	return bool((m.ActualLength) == (4))
}

func (m *BACnetTagPayloadUnsignedInteger) GetActualValue() uint32 {
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	return utils.InlineIf(m.GetIsUint8(), func() interface{} { return uint32((*m.GetValueUint8())) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(m.GetIsUint16(), func() interface{} { return uint32((*m.GetValueUint16())) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(m.GetIsUint24(), func() interface{} { return uint32((*m.GetValueUint24())) }, func() interface{} {
				return uint32(uint32(utils.InlineIf(m.GetIsUint32(), func() interface{} { return uint32((*m.GetValueUint32())) }, func() interface{} { return uint32(uint32(0)) }).(uint32)))
			}).(uint32)))
		}).(uint32)))
	}).(uint32)
}

// NewBACnetTagPayloadUnsignedInteger factory function for BACnetTagPayloadUnsignedInteger
func NewBACnetTagPayloadUnsignedInteger(valueUint8 *uint8, valueUint16 *uint16, valueUint24 *uint32, valueUint32 *uint32, actualLength uint32) *BACnetTagPayloadUnsignedInteger {
	return &BACnetTagPayloadUnsignedInteger{ValueUint8: valueUint8, ValueUint16: valueUint16, ValueUint24: valueUint24, ValueUint32: valueUint32, ActualLength: actualLength}
}

func CastBACnetTagPayloadUnsignedInteger(structType interface{}) *BACnetTagPayloadUnsignedInteger {
	castFunc := func(typ interface{}) *BACnetTagPayloadUnsignedInteger {
		if casted, ok := typ.(BACnetTagPayloadUnsignedInteger); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagPayloadUnsignedInteger); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagPayloadUnsignedInteger) GetTypeName() string {
	return "BACnetTagPayloadUnsignedInteger"
}

func (m *BACnetTagPayloadUnsignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadUnsignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint8)
	if m.ValueUint8 != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint16)
	if m.ValueUint16 != nil {
		lengthInBits += 16
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint24)
	if m.ValueUint24 != nil {
		lengthInBits += 24
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint32)
	if m.ValueUint32 != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetTagPayloadUnsignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadUnsignedIntegerParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTagPayloadUnsignedInteger, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadUnsignedInteger"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_isUint8 := bool((actualLength) == (1))
	isUint8 := bool(_isUint8)
	_ = isUint8

	// Optional Field (valueUint8) (Can be skipped, if a given expression evaluates to false)
	var valueUint8 *uint8 = nil
	if isUint8 {
		_val, _err := readBuffer.ReadUint8("valueUint8", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint8' field")
		}
		valueUint8 = &_val
	}

	// Virtual field
	_isUint16 := bool((actualLength) == (2))
	isUint16 := bool(_isUint16)
	_ = isUint16

	// Optional Field (valueUint16) (Can be skipped, if a given expression evaluates to false)
	var valueUint16 *uint16 = nil
	if isUint16 {
		_val, _err := readBuffer.ReadUint16("valueUint16", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint16' field")
		}
		valueUint16 = &_val
	}

	// Virtual field
	_isUint24 := bool((actualLength) == (3))
	isUint24 := bool(_isUint24)
	_ = isUint24

	// Optional Field (valueUint24) (Can be skipped, if a given expression evaluates to false)
	var valueUint24 *uint32 = nil
	if isUint24 {
		_val, _err := readBuffer.ReadUint32("valueUint24", 24)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint24' field")
		}
		valueUint24 = &_val
	}

	// Virtual field
	_isUint32 := bool((actualLength) == (4))
	isUint32 := bool(_isUint32)
	_ = isUint32

	// Optional Field (valueUint32) (Can be skipped, if a given expression evaluates to false)
	var valueUint32 *uint32 = nil
	if isUint32 {
		_val, _err := readBuffer.ReadUint32("valueUint32", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint32' field")
		}
		valueUint32 = &_val
	}

	// Virtual field
	_actualValue := utils.InlineIf(isUint8, func() interface{} { return uint32((*valueUint8)) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(isUint16, func() interface{} { return uint32((*valueUint16)) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(isUint24, func() interface{} { return uint32((*valueUint24)) }, func() interface{} {
				return uint32(uint32(utils.InlineIf(isUint32, func() interface{} { return uint32((*valueUint32)) }, func() interface{} { return uint32(uint32(0)) }).(uint32)))
			}).(uint32)))
		}).(uint32)))
	}).(uint32)
	actualValue := uint32(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadUnsignedInteger"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadUnsignedInteger(valueUint8, valueUint16, valueUint24, valueUint32, actualLength), nil
}

func (m *BACnetTagPayloadUnsignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadUnsignedInteger"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _isUint8Err := writeBuffer.WriteVirtual("isUint8", m.GetIsUint8()); _isUint8Err != nil {
		return errors.Wrap(_isUint8Err, "Error serializing 'isUint8' field")
	}

	// Optional Field (valueUint8) (Can be skipped, if the value is null)
	var valueUint8 *uint8 = nil
	if m.ValueUint8 != nil {
		valueUint8 = m.ValueUint8
		_valueUint8Err := writeBuffer.WriteUint8("valueUint8", 8, *(valueUint8))
		if _valueUint8Err != nil {
			return errors.Wrap(_valueUint8Err, "Error serializing 'valueUint8' field")
		}
	}
	// Virtual field
	if _isUint16Err := writeBuffer.WriteVirtual("isUint16", m.GetIsUint16()); _isUint16Err != nil {
		return errors.Wrap(_isUint16Err, "Error serializing 'isUint16' field")
	}

	// Optional Field (valueUint16) (Can be skipped, if the value is null)
	var valueUint16 *uint16 = nil
	if m.ValueUint16 != nil {
		valueUint16 = m.ValueUint16
		_valueUint16Err := writeBuffer.WriteUint16("valueUint16", 16, *(valueUint16))
		if _valueUint16Err != nil {
			return errors.Wrap(_valueUint16Err, "Error serializing 'valueUint16' field")
		}
	}
	// Virtual field
	if _isUint24Err := writeBuffer.WriteVirtual("isUint24", m.GetIsUint24()); _isUint24Err != nil {
		return errors.Wrap(_isUint24Err, "Error serializing 'isUint24' field")
	}

	// Optional Field (valueUint24) (Can be skipped, if the value is null)
	var valueUint24 *uint32 = nil
	if m.ValueUint24 != nil {
		valueUint24 = m.ValueUint24
		_valueUint24Err := writeBuffer.WriteUint32("valueUint24", 24, *(valueUint24))
		if _valueUint24Err != nil {
			return errors.Wrap(_valueUint24Err, "Error serializing 'valueUint24' field")
		}
	}
	// Virtual field
	if _isUint32Err := writeBuffer.WriteVirtual("isUint32", m.GetIsUint32()); _isUint32Err != nil {
		return errors.Wrap(_isUint32Err, "Error serializing 'isUint32' field")
	}

	// Optional Field (valueUint32) (Can be skipped, if the value is null)
	var valueUint32 *uint32 = nil
	if m.ValueUint32 != nil {
		valueUint32 = m.ValueUint32
		_valueUint32Err := writeBuffer.WriteUint32("valueUint32", 32, *(valueUint32))
		if _valueUint32Err != nil {
			return errors.Wrap(_valueUint32Err, "Error serializing 'valueUint32' field")
		}
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadUnsignedInteger"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}