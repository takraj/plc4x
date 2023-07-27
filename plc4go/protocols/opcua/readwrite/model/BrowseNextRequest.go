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
)

// Code generated by code-generation. DO NOT EDIT.

// BrowseNextRequest is the corresponding interface of BrowseNextRequest
type BrowseNextRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetReleaseContinuationPoints returns ReleaseContinuationPoints (property field)
	GetReleaseContinuationPoints() bool
	// GetNoOfContinuationPoints returns NoOfContinuationPoints (property field)
	GetNoOfContinuationPoints() int32
	// GetContinuationPoints returns ContinuationPoints (property field)
	GetContinuationPoints() []PascalByteString
}

// BrowseNextRequestExactly can be used when we want exactly this type and not a type which fulfills BrowseNextRequest.
// This is useful for switch cases.
type BrowseNextRequestExactly interface {
	BrowseNextRequest
	isBrowseNextRequest() bool
}

// _BrowseNextRequest is the data-structure of this message
type _BrowseNextRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader             ExtensionObjectDefinition
	ReleaseContinuationPoints bool
	NoOfContinuationPoints    int32
	ContinuationPoints        []PascalByteString
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowseNextRequest) GetIdentifier() string {
	return "533"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowseNextRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BrowseNextRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowseNextRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_BrowseNextRequest) GetReleaseContinuationPoints() bool {
	return m.ReleaseContinuationPoints
}

func (m *_BrowseNextRequest) GetNoOfContinuationPoints() int32 {
	return m.NoOfContinuationPoints
}

func (m *_BrowseNextRequest) GetContinuationPoints() []PascalByteString {
	return m.ContinuationPoints
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowseNextRequest factory function for _BrowseNextRequest
func NewBrowseNextRequest(requestHeader ExtensionObjectDefinition, releaseContinuationPoints bool, noOfContinuationPoints int32, continuationPoints []PascalByteString) *_BrowseNextRequest {
	_result := &_BrowseNextRequest{
		RequestHeader:              requestHeader,
		ReleaseContinuationPoints:  releaseContinuationPoints,
		NoOfContinuationPoints:     noOfContinuationPoints,
		ContinuationPoints:         continuationPoints,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowseNextRequest(structType any) BrowseNextRequest {
	if casted, ok := structType.(BrowseNextRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BrowseNextRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BrowseNextRequest) GetTypeName() string {
	return "BrowseNextRequest"
}

func (m *_BrowseNextRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (releaseContinuationPoints)
	lengthInBits += 1

	// Simple field (noOfContinuationPoints)
	lengthInBits += 32

	// Array field
	if len(m.ContinuationPoints) > 0 {
		for _curItem, element := range m.ContinuationPoints {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ContinuationPoints), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowseNextRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowseNextRequestParse(ctx context.Context, theBytes []byte, identifier string) (BrowseNextRequest, error) {
	return BrowseNextRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BrowseNextRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BrowseNextRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BrowseNextRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowseNextRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of BrowseNextRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of BrowseNextRequest")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (releaseContinuationPoints)
	_releaseContinuationPoints, _releaseContinuationPointsErr := readBuffer.ReadBit("releaseContinuationPoints")
	if _releaseContinuationPointsErr != nil {
		return nil, errors.Wrap(_releaseContinuationPointsErr, "Error parsing 'releaseContinuationPoints' field of BrowseNextRequest")
	}
	releaseContinuationPoints := _releaseContinuationPoints

	// Simple Field (noOfContinuationPoints)
	_noOfContinuationPoints, _noOfContinuationPointsErr := readBuffer.ReadInt32("noOfContinuationPoints", 32)
	if _noOfContinuationPointsErr != nil {
		return nil, errors.Wrap(_noOfContinuationPointsErr, "Error parsing 'noOfContinuationPoints' field of BrowseNextRequest")
	}
	noOfContinuationPoints := _noOfContinuationPoints

	// Array field (continuationPoints)
	if pullErr := readBuffer.PullContext("continuationPoints", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for continuationPoints")
	}
	// Count array
	continuationPoints := make([]PascalByteString, noOfContinuationPoints)
	// This happens when the size is set conditional to 0
	if len(continuationPoints) == 0 {
		continuationPoints = nil
	}
	{
		_numItems := uint16(noOfContinuationPoints)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalByteStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'continuationPoints' field of BrowseNextRequest")
			}
			continuationPoints[_curItem] = _item.(PascalByteString)
		}
	}
	if closeErr := readBuffer.CloseContext("continuationPoints", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for continuationPoints")
	}

	if closeErr := readBuffer.CloseContext("BrowseNextRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowseNextRequest")
	}

	// Create a partially initialized instance
	_child := &_BrowseNextRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		ReleaseContinuationPoints:  releaseContinuationPoints,
		NoOfContinuationPoints:     noOfContinuationPoints,
		ContinuationPoints:         continuationPoints,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BrowseNextRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowseNextRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowseNextRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowseNextRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (releaseContinuationPoints)
		releaseContinuationPoints := bool(m.GetReleaseContinuationPoints())
		_releaseContinuationPointsErr := writeBuffer.WriteBit("releaseContinuationPoints", (releaseContinuationPoints))
		if _releaseContinuationPointsErr != nil {
			return errors.Wrap(_releaseContinuationPointsErr, "Error serializing 'releaseContinuationPoints' field")
		}

		// Simple Field (noOfContinuationPoints)
		noOfContinuationPoints := int32(m.GetNoOfContinuationPoints())
		_noOfContinuationPointsErr := writeBuffer.WriteInt32("noOfContinuationPoints", 32, (noOfContinuationPoints))
		if _noOfContinuationPointsErr != nil {
			return errors.Wrap(_noOfContinuationPointsErr, "Error serializing 'noOfContinuationPoints' field")
		}

		// Array Field (continuationPoints)
		if pushErr := writeBuffer.PushContext("continuationPoints", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for continuationPoints")
		}
		for _curItem, _element := range m.GetContinuationPoints() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetContinuationPoints()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'continuationPoints' field")
			}
		}
		if popErr := writeBuffer.PopContext("continuationPoints", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for continuationPoints")
		}

		if popErr := writeBuffer.PopContext("BrowseNextRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowseNextRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowseNextRequest) isBrowseNextRequest() bool {
	return true
}

func (m *_BrowseNextRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}