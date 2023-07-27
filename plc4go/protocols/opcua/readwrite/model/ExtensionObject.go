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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// ExtensionObject is the corresponding interface of ExtensionObject
type ExtensionObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTypeId returns TypeId (property field)
	GetTypeId() ExpandedNodeId
	// GetEncodingMask returns EncodingMask (property field)
	GetEncodingMask() ExtensionObjectEncodingMask
	// GetBody returns Body (property field)
	GetBody() ExtensionObjectDefinition
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// ExtensionObjectExactly can be used when we want exactly this type and not a type which fulfills ExtensionObject.
// This is useful for switch cases.
type ExtensionObjectExactly interface {
	ExtensionObject
	isExtensionObject() bool
}

// _ExtensionObject is the data-structure of this message
type _ExtensionObject struct {
	TypeId       ExpandedNodeId
	EncodingMask ExtensionObjectEncodingMask
	Body         ExtensionObjectDefinition

	// Arguments.
	IncludeEncodingMask bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensionObject) GetTypeId() ExpandedNodeId {
	return m.TypeId
}

func (m *_ExtensionObject) GetEncodingMask() ExtensionObjectEncodingMask {
	return m.EncodingMask
}

func (m *_ExtensionObject) GetBody() ExtensionObjectDefinition {
	return m.Body
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ExtensionObject) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	encodingMask := m.EncodingMask
	_ = encodingMask
	return fmt.Sprintf("%v", m.GetTypeId().GetIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewExtensionObject factory function for _ExtensionObject
func NewExtensionObject(typeId ExpandedNodeId, encodingMask ExtensionObjectEncodingMask, body ExtensionObjectDefinition, includeEncodingMask bool) *_ExtensionObject {
	return &_ExtensionObject{TypeId: typeId, EncodingMask: encodingMask, Body: body, IncludeEncodingMask: includeEncodingMask}
}

// Deprecated: use the interface for direct cast
func CastExtensionObject(structType any) ExtensionObject {
	if casted, ok := structType.(ExtensionObject); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensionObject); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensionObject) GetTypeName() string {
	return "ExtensionObject"
}

func (m *_ExtensionObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeId)
	lengthInBits += m.TypeId.GetLengthInBits(ctx)

	// Optional Field (encodingMask)
	if m.EncodingMask != nil {
		lengthInBits += m.EncodingMask.GetLengthInBits(ctx)
	}

	// A virtual field doesn't have any in- or output.

	// Simple field (body)
	lengthInBits += m.Body.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ExtensionObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ExtensionObjectParse(ctx context.Context, theBytes []byte, includeEncodingMask bool) (ExtensionObject, error) {
	return ExtensionObjectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), includeEncodingMask)
}

func ExtensionObjectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, includeEncodingMask bool) (ExtensionObject, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ExtensionObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensionObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (typeId)
	if pullErr := readBuffer.PullContext("typeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for typeId")
	}
	_typeId, _typeIdErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _typeIdErr != nil {
		return nil, errors.Wrap(_typeIdErr, "Error parsing 'typeId' field of ExtensionObject")
	}
	typeId := _typeId.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("typeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for typeId")
	}

	// Optional Field (encodingMask) (Can be skipped, if a given expression evaluates to false)
	var encodingMask ExtensionObjectEncodingMask = nil
	if includeEncodingMask {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("encodingMask"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for encodingMask")
		}
		_val, _err := ExtensionObjectEncodingMaskParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'encodingMask' field of ExtensionObject")
		default:
			encodingMask = _val.(ExtensionObjectEncodingMask)
			if closeErr := readBuffer.CloseContext("encodingMask"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for encodingMask")
			}
		}
	}

	// Virtual field
	_identifier := typeId.GetIdentifier()
	identifier := fmt.Sprintf("%v", _identifier)
	_ = identifier

	// Simple Field (body)
	if pullErr := readBuffer.PullContext("body"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for body")
	}
	_body, _bodyErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string(identifier))
	if _bodyErr != nil {
		return nil, errors.Wrap(_bodyErr, "Error parsing 'body' field of ExtensionObject")
	}
	body := _body.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("body"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for body")
	}

	if closeErr := readBuffer.CloseContext("ExtensionObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensionObject")
	}

	// Create the instance
	return &_ExtensionObject{
		IncludeEncodingMask: includeEncodingMask,
		TypeId:              typeId,
		EncodingMask:        encodingMask,
		Body:                body,
	}, nil
}

func (m *_ExtensionObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ExtensionObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ExtensionObject"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExtensionObject")
	}

	// Simple Field (typeId)
	if pushErr := writeBuffer.PushContext("typeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for typeId")
	}
	_typeIdErr := writeBuffer.WriteSerializable(ctx, m.GetTypeId())
	if popErr := writeBuffer.PopContext("typeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for typeId")
	}
	if _typeIdErr != nil {
		return errors.Wrap(_typeIdErr, "Error serializing 'typeId' field")
	}

	// Optional Field (encodingMask) (Can be skipped, if the value is null)
	var encodingMask ExtensionObjectEncodingMask = nil
	if m.GetEncodingMask() != nil {
		if pushErr := writeBuffer.PushContext("encodingMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for encodingMask")
		}
		encodingMask = m.GetEncodingMask()
		_encodingMaskErr := writeBuffer.WriteSerializable(ctx, encodingMask)
		if popErr := writeBuffer.PopContext("encodingMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for encodingMask")
		}
		if _encodingMaskErr != nil {
			return errors.Wrap(_encodingMaskErr, "Error serializing 'encodingMask' field")
		}
	}
	// Virtual field
	if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
		return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
	}

	// Simple Field (body)
	if pushErr := writeBuffer.PushContext("body"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for body")
	}
	_bodyErr := writeBuffer.WriteSerializable(ctx, m.GetBody())
	if popErr := writeBuffer.PopContext("body"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for body")
	}
	if _bodyErr != nil {
		return errors.Wrap(_bodyErr, "Error serializing 'body' field")
	}

	if popErr := writeBuffer.PopContext("ExtensionObject"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExtensionObject")
	}
	return nil
}

////
// Arguments Getter

func (m *_ExtensionObject) GetIncludeEncodingMask() bool {
	return m.IncludeEncodingMask
}

//
////

func (m *_ExtensionObject) isExtensionObject() bool {
	return true
}

func (m *_ExtensionObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}