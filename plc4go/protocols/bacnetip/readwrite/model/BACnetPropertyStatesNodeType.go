/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesNodeType is the corresponding interface of BACnetPropertyStatesNodeType
type BACnetPropertyStatesNodeType interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetNodeType returns NodeType (property field)
	GetNodeType() BACnetNodeTypeTagged
}

// BACnetPropertyStatesNodeTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesNodeType.
// This is useful for switch cases.
type BACnetPropertyStatesNodeTypeExactly interface {
	BACnetPropertyStatesNodeType
	isBACnetPropertyStatesNodeType() bool
}

// _BACnetPropertyStatesNodeType is the data-structure of this message
type _BACnetPropertyStatesNodeType struct {
	*_BACnetPropertyStates
	NodeType BACnetNodeTypeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesNodeType) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesNodeType) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNodeType) GetNodeType() BACnetNodeTypeTagged {
	return m.NodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNodeType factory function for _BACnetPropertyStatesNodeType
func NewBACnetPropertyStatesNodeType(nodeType BACnetNodeTypeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesNodeType {
	_result := &_BACnetPropertyStatesNodeType{
		NodeType:              nodeType,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNodeType(structType interface{}) BACnetPropertyStatesNodeType {
	if casted, ok := structType.(BACnetPropertyStatesNodeType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNodeType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNodeType) GetTypeName() string {
	return "BACnetPropertyStatesNodeType"
}

func (m *_BACnetPropertyStatesNodeType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesNodeType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nodeType)
	lengthInBits += m.NodeType.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesNodeType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesNodeTypeParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesNodeType, error) {
	return BACnetPropertyStatesNodeTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), peekedTagNumber) // TODO: get endianness from mspec
}

func BACnetPropertyStatesNodeTypeParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesNodeType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNodeType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeType)
	if pullErr := readBuffer.PullContext("nodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeType")
	}
	_nodeType, _nodeTypeErr := BACnetNodeTypeTaggedParseWithBuffer(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _nodeTypeErr != nil {
		return nil, errors.Wrap(_nodeTypeErr, "Error parsing 'nodeType' field of BACnetPropertyStatesNodeType")
	}
	nodeType := _nodeType.(BACnetNodeTypeTagged)
	if closeErr := readBuffer.CloseContext("nodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeType")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNodeType")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesNodeType{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		NodeType:              nodeType,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesNodeType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNodeType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNodeType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNodeType")
		}

		// Simple Field (nodeType)
		if pushErr := writeBuffer.PushContext("nodeType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeType")
		}
		_nodeTypeErr := writeBuffer.WriteSerializable(m.GetNodeType())
		if popErr := writeBuffer.PopContext("nodeType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeType")
		}
		if _nodeTypeErr != nil {
			return errors.Wrap(_nodeTypeErr, "Error serializing 'nodeType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNodeType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNodeType")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNodeType) isBACnetPropertyStatesNodeType() bool {
	return true
}

func (m *_BACnetPropertyStatesNodeType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}