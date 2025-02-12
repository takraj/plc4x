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

// SetTriggeringRequest is the corresponding interface of SetTriggeringRequest
type SetTriggeringRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetTriggeringItemId returns TriggeringItemId (property field)
	GetTriggeringItemId() uint32
	// GetNoOfLinksToAdd returns NoOfLinksToAdd (property field)
	GetNoOfLinksToAdd() int32
	// GetLinksToAdd returns LinksToAdd (property field)
	GetLinksToAdd() []uint32
	// GetNoOfLinksToRemove returns NoOfLinksToRemove (property field)
	GetNoOfLinksToRemove() int32
	// GetLinksToRemove returns LinksToRemove (property field)
	GetLinksToRemove() []uint32
}

// SetTriggeringRequestExactly can be used when we want exactly this type and not a type which fulfills SetTriggeringRequest.
// This is useful for switch cases.
type SetTriggeringRequestExactly interface {
	SetTriggeringRequest
	isSetTriggeringRequest() bool
}

// _SetTriggeringRequest is the data-structure of this message
type _SetTriggeringRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader     ExtensionObjectDefinition
	SubscriptionId    uint32
	TriggeringItemId  uint32
	NoOfLinksToAdd    int32
	LinksToAdd        []uint32
	NoOfLinksToRemove int32
	LinksToRemove     []uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetTriggeringRequest) GetIdentifier() string {
	return "775"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetTriggeringRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SetTriggeringRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SetTriggeringRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_SetTriggeringRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_SetTriggeringRequest) GetTriggeringItemId() uint32 {
	return m.TriggeringItemId
}

func (m *_SetTriggeringRequest) GetNoOfLinksToAdd() int32 {
	return m.NoOfLinksToAdd
}

func (m *_SetTriggeringRequest) GetLinksToAdd() []uint32 {
	return m.LinksToAdd
}

func (m *_SetTriggeringRequest) GetNoOfLinksToRemove() int32 {
	return m.NoOfLinksToRemove
}

func (m *_SetTriggeringRequest) GetLinksToRemove() []uint32 {
	return m.LinksToRemove
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSetTriggeringRequest factory function for _SetTriggeringRequest
func NewSetTriggeringRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, triggeringItemId uint32, noOfLinksToAdd int32, linksToAdd []uint32, noOfLinksToRemove int32, linksToRemove []uint32) *_SetTriggeringRequest {
	_result := &_SetTriggeringRequest{
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		TriggeringItemId:           triggeringItemId,
		NoOfLinksToAdd:             noOfLinksToAdd,
		LinksToAdd:                 linksToAdd,
		NoOfLinksToRemove:          noOfLinksToRemove,
		LinksToRemove:              linksToRemove,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetTriggeringRequest(structType any) SetTriggeringRequest {
	if casted, ok := structType.(SetTriggeringRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SetTriggeringRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SetTriggeringRequest) GetTypeName() string {
	return "SetTriggeringRequest"
}

func (m *_SetTriggeringRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (triggeringItemId)
	lengthInBits += 32

	// Simple field (noOfLinksToAdd)
	lengthInBits += 32

	// Array field
	if len(m.LinksToAdd) > 0 {
		lengthInBits += 32 * uint16(len(m.LinksToAdd))
	}

	// Simple field (noOfLinksToRemove)
	lengthInBits += 32

	// Array field
	if len(m.LinksToRemove) > 0 {
		lengthInBits += 32 * uint16(len(m.LinksToRemove))
	}

	return lengthInBits
}

func (m *_SetTriggeringRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetTriggeringRequestParse(ctx context.Context, theBytes []byte, identifier string) (SetTriggeringRequest, error) {
	return SetTriggeringRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SetTriggeringRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SetTriggeringRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SetTriggeringRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetTriggeringRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of SetTriggeringRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (subscriptionId)
	_subscriptionId, _subscriptionIdErr := readBuffer.ReadUint32("subscriptionId", 32)
	if _subscriptionIdErr != nil {
		return nil, errors.Wrap(_subscriptionIdErr, "Error parsing 'subscriptionId' field of SetTriggeringRequest")
	}
	subscriptionId := _subscriptionId

	// Simple Field (triggeringItemId)
	_triggeringItemId, _triggeringItemIdErr := readBuffer.ReadUint32("triggeringItemId", 32)
	if _triggeringItemIdErr != nil {
		return nil, errors.Wrap(_triggeringItemIdErr, "Error parsing 'triggeringItemId' field of SetTriggeringRequest")
	}
	triggeringItemId := _triggeringItemId

	// Simple Field (noOfLinksToAdd)
	_noOfLinksToAdd, _noOfLinksToAddErr := readBuffer.ReadInt32("noOfLinksToAdd", 32)
	if _noOfLinksToAddErr != nil {
		return nil, errors.Wrap(_noOfLinksToAddErr, "Error parsing 'noOfLinksToAdd' field of SetTriggeringRequest")
	}
	noOfLinksToAdd := _noOfLinksToAdd

	// Array field (linksToAdd)
	if pullErr := readBuffer.PullContext("linksToAdd", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for linksToAdd")
	}
	// Count array
	linksToAdd := make([]uint32, utils.Max(noOfLinksToAdd, 0))
	// This happens when the size is set conditional to 0
	if len(linksToAdd) == 0 {
		linksToAdd = nil
	}
	{
		_numItems := uint16(utils.Max(noOfLinksToAdd, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint32("", 32)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'linksToAdd' field of SetTriggeringRequest")
			}
			linksToAdd[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("linksToAdd", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for linksToAdd")
	}

	// Simple Field (noOfLinksToRemove)
	_noOfLinksToRemove, _noOfLinksToRemoveErr := readBuffer.ReadInt32("noOfLinksToRemove", 32)
	if _noOfLinksToRemoveErr != nil {
		return nil, errors.Wrap(_noOfLinksToRemoveErr, "Error parsing 'noOfLinksToRemove' field of SetTriggeringRequest")
	}
	noOfLinksToRemove := _noOfLinksToRemove

	// Array field (linksToRemove)
	if pullErr := readBuffer.PullContext("linksToRemove", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for linksToRemove")
	}
	// Count array
	linksToRemove := make([]uint32, utils.Max(noOfLinksToRemove, 0))
	// This happens when the size is set conditional to 0
	if len(linksToRemove) == 0 {
		linksToRemove = nil
	}
	{
		_numItems := uint16(utils.Max(noOfLinksToRemove, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint32("", 32)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'linksToRemove' field of SetTriggeringRequest")
			}
			linksToRemove[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("linksToRemove", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for linksToRemove")
	}

	if closeErr := readBuffer.CloseContext("SetTriggeringRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetTriggeringRequest")
	}

	// Create a partially initialized instance
	_child := &_SetTriggeringRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		SubscriptionId:             subscriptionId,
		TriggeringItemId:           triggeringItemId,
		NoOfLinksToAdd:             noOfLinksToAdd,
		LinksToAdd:                 linksToAdd,
		NoOfLinksToRemove:          noOfLinksToRemove,
		LinksToRemove:              linksToRemove,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SetTriggeringRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetTriggeringRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetTriggeringRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetTriggeringRequest")
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

		// Simple Field (subscriptionId)
		subscriptionId := uint32(m.GetSubscriptionId())
		_subscriptionIdErr := writeBuffer.WriteUint32("subscriptionId", 32, (subscriptionId))
		if _subscriptionIdErr != nil {
			return errors.Wrap(_subscriptionIdErr, "Error serializing 'subscriptionId' field")
		}

		// Simple Field (triggeringItemId)
		triggeringItemId := uint32(m.GetTriggeringItemId())
		_triggeringItemIdErr := writeBuffer.WriteUint32("triggeringItemId", 32, (triggeringItemId))
		if _triggeringItemIdErr != nil {
			return errors.Wrap(_triggeringItemIdErr, "Error serializing 'triggeringItemId' field")
		}

		// Simple Field (noOfLinksToAdd)
		noOfLinksToAdd := int32(m.GetNoOfLinksToAdd())
		_noOfLinksToAddErr := writeBuffer.WriteInt32("noOfLinksToAdd", 32, (noOfLinksToAdd))
		if _noOfLinksToAddErr != nil {
			return errors.Wrap(_noOfLinksToAddErr, "Error serializing 'noOfLinksToAdd' field")
		}

		// Array Field (linksToAdd)
		if pushErr := writeBuffer.PushContext("linksToAdd", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for linksToAdd")
		}
		for _curItem, _element := range m.GetLinksToAdd() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint32("", 32, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'linksToAdd' field")
			}
		}
		if popErr := writeBuffer.PopContext("linksToAdd", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for linksToAdd")
		}

		// Simple Field (noOfLinksToRemove)
		noOfLinksToRemove := int32(m.GetNoOfLinksToRemove())
		_noOfLinksToRemoveErr := writeBuffer.WriteInt32("noOfLinksToRemove", 32, (noOfLinksToRemove))
		if _noOfLinksToRemoveErr != nil {
			return errors.Wrap(_noOfLinksToRemoveErr, "Error serializing 'noOfLinksToRemove' field")
		}

		// Array Field (linksToRemove)
		if pushErr := writeBuffer.PushContext("linksToRemove", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for linksToRemove")
		}
		for _curItem, _element := range m.GetLinksToRemove() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint32("", 32, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'linksToRemove' field")
			}
		}
		if popErr := writeBuffer.PopContext("linksToRemove", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for linksToRemove")
		}

		if popErr := writeBuffer.PopContext("SetTriggeringRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetTriggeringRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetTriggeringRequest) isSetTriggeringRequest() bool {
	return true
}

func (m *_SetTriggeringRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
