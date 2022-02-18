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
type BACnetTagPayloadBoolean struct {

	// Arguments.
	ActualLength uint32
}

// The corresponding interface
type IBACnetTagPayloadBoolean interface {
	// GetValue returns Value
	GetValue() bool
	// GetIsTrue returns IsTrue
	GetIsTrue() bool
	// GetIsFalse returns IsFalse
	GetIsFalse() bool
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

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadBoolean) GetValue() bool {
	return bool((m.ActualLength) == (1))
}

func (m *BACnetTagPayloadBoolean) GetIsTrue() bool {
	return m.GetValue()
}

func (m *BACnetTagPayloadBoolean) GetIsFalse() bool {
	return !(m.GetValue())
}

// NewBACnetTagPayloadBoolean factory function for BACnetTagPayloadBoolean
func NewBACnetTagPayloadBoolean(actualLength uint32) *BACnetTagPayloadBoolean {
	return &BACnetTagPayloadBoolean{ActualLength: actualLength}
}

func CastBACnetTagPayloadBoolean(structType interface{}) *BACnetTagPayloadBoolean {
	castFunc := func(typ interface{}) *BACnetTagPayloadBoolean {
		if casted, ok := typ.(BACnetTagPayloadBoolean); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagPayloadBoolean); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagPayloadBoolean) GetTypeName() string {
	return "BACnetTagPayloadBoolean"
}

func (m *BACnetTagPayloadBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetTagPayloadBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadBooleanParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTagPayloadBoolean, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadBoolean"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_value := bool((actualLength) == (1))
	value := bool(_value)
	_ = value

	// Virtual field
	_isTrue := value
	isTrue := bool(_isTrue)
	_ = isTrue

	// Virtual field
	_isFalse := !(value)
	isFalse := bool(_isFalse)
	_ = isFalse

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadBoolean"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadBoolean(actualLength), nil
}

func (m *BACnetTagPayloadBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadBoolean"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _valueErr := writeBuffer.WriteVirtual("value", m.GetValue()); _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	if _isTrueErr := writeBuffer.WriteVirtual("isTrue", m.GetIsTrue()); _isTrueErr != nil {
		return errors.Wrap(_isTrueErr, "Error serializing 'isTrue' field")
	}
	// Virtual field
	if _isFalseErr := writeBuffer.WriteVirtual("isFalse", m.GetIsFalse()); _isFalseErr != nil {
		return errors.Wrap(_isFalseErr, "Error serializing 'isFalse' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadBoolean"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}