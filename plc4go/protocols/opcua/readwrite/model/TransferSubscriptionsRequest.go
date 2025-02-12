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

// TransferSubscriptionsRequest is the corresponding interface of TransferSubscriptionsRequest
type TransferSubscriptionsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfSubscriptionIds returns NoOfSubscriptionIds (property field)
	GetNoOfSubscriptionIds() int32
	// GetSubscriptionIds returns SubscriptionIds (property field)
	GetSubscriptionIds() []uint32
	// GetSendInitialValues returns SendInitialValues (property field)
	GetSendInitialValues() bool
}

// TransferSubscriptionsRequestExactly can be used when we want exactly this type and not a type which fulfills TransferSubscriptionsRequest.
// This is useful for switch cases.
type TransferSubscriptionsRequestExactly interface {
	TransferSubscriptionsRequest
	isTransferSubscriptionsRequest() bool
}

// _TransferSubscriptionsRequest is the data-structure of this message
type _TransferSubscriptionsRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader       ExtensionObjectDefinition
	NoOfSubscriptionIds int32
	SubscriptionIds     []uint32
	SendInitialValues   bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransferSubscriptionsRequest) GetIdentifier() string {
	return "841"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransferSubscriptionsRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_TransferSubscriptionsRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransferSubscriptionsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_TransferSubscriptionsRequest) GetNoOfSubscriptionIds() int32 {
	return m.NoOfSubscriptionIds
}

func (m *_TransferSubscriptionsRequest) GetSubscriptionIds() []uint32 {
	return m.SubscriptionIds
}

func (m *_TransferSubscriptionsRequest) GetSendInitialValues() bool {
	return m.SendInitialValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTransferSubscriptionsRequest factory function for _TransferSubscriptionsRequest
func NewTransferSubscriptionsRequest(requestHeader ExtensionObjectDefinition, noOfSubscriptionIds int32, subscriptionIds []uint32, sendInitialValues bool) *_TransferSubscriptionsRequest {
	_result := &_TransferSubscriptionsRequest{
		RequestHeader:              requestHeader,
		NoOfSubscriptionIds:        noOfSubscriptionIds,
		SubscriptionIds:            subscriptionIds,
		SendInitialValues:          sendInitialValues,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTransferSubscriptionsRequest(structType any) TransferSubscriptionsRequest {
	if casted, ok := structType.(TransferSubscriptionsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TransferSubscriptionsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TransferSubscriptionsRequest) GetTypeName() string {
	return "TransferSubscriptionsRequest"
}

func (m *_TransferSubscriptionsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfSubscriptionIds)
	lengthInBits += 32

	// Array field
	if len(m.SubscriptionIds) > 0 {
		lengthInBits += 32 * uint16(len(m.SubscriptionIds))
	}

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (sendInitialValues)
	lengthInBits += 1

	return lengthInBits
}

func (m *_TransferSubscriptionsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TransferSubscriptionsRequestParse(ctx context.Context, theBytes []byte, identifier string) (TransferSubscriptionsRequest, error) {
	return TransferSubscriptionsRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func TransferSubscriptionsRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (TransferSubscriptionsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TransferSubscriptionsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransferSubscriptionsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of TransferSubscriptionsRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfSubscriptionIds)
	_noOfSubscriptionIds, _noOfSubscriptionIdsErr := readBuffer.ReadInt32("noOfSubscriptionIds", 32)
	if _noOfSubscriptionIdsErr != nil {
		return nil, errors.Wrap(_noOfSubscriptionIdsErr, "Error parsing 'noOfSubscriptionIds' field of TransferSubscriptionsRequest")
	}
	noOfSubscriptionIds := _noOfSubscriptionIds

	// Array field (subscriptionIds)
	if pullErr := readBuffer.PullContext("subscriptionIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscriptionIds")
	}
	// Count array
	subscriptionIds := make([]uint32, utils.Max(noOfSubscriptionIds, 0))
	// This happens when the size is set conditional to 0
	if len(subscriptionIds) == 0 {
		subscriptionIds = nil
	}
	{
		_numItems := uint16(utils.Max(noOfSubscriptionIds, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint32("", 32)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subscriptionIds' field of TransferSubscriptionsRequest")
			}
			subscriptionIds[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("subscriptionIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscriptionIds")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TransferSubscriptionsRequest")
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

	// Simple Field (sendInitialValues)
	_sendInitialValues, _sendInitialValuesErr := readBuffer.ReadBit("sendInitialValues")
	if _sendInitialValuesErr != nil {
		return nil, errors.Wrap(_sendInitialValuesErr, "Error parsing 'sendInitialValues' field of TransferSubscriptionsRequest")
	}
	sendInitialValues := _sendInitialValues

	if closeErr := readBuffer.CloseContext("TransferSubscriptionsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransferSubscriptionsRequest")
	}

	// Create a partially initialized instance
	_child := &_TransferSubscriptionsRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfSubscriptionIds:        noOfSubscriptionIds,
		SubscriptionIds:            subscriptionIds,
		SendInitialValues:          sendInitialValues,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_TransferSubscriptionsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransferSubscriptionsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransferSubscriptionsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransferSubscriptionsRequest")
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

		// Simple Field (noOfSubscriptionIds)
		noOfSubscriptionIds := int32(m.GetNoOfSubscriptionIds())
		_noOfSubscriptionIdsErr := writeBuffer.WriteInt32("noOfSubscriptionIds", 32, (noOfSubscriptionIds))
		if _noOfSubscriptionIdsErr != nil {
			return errors.Wrap(_noOfSubscriptionIdsErr, "Error serializing 'noOfSubscriptionIds' field")
		}

		// Array Field (subscriptionIds)
		if pushErr := writeBuffer.PushContext("subscriptionIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriptionIds")
		}
		for _curItem, _element := range m.GetSubscriptionIds() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint32("", 32, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'subscriptionIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("subscriptionIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriptionIds")
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

		// Simple Field (sendInitialValues)
		sendInitialValues := bool(m.GetSendInitialValues())
		_sendInitialValuesErr := writeBuffer.WriteBit("sendInitialValues", (sendInitialValues))
		if _sendInitialValuesErr != nil {
			return errors.Wrap(_sendInitialValuesErr, "Error serializing 'sendInitialValues' field")
		}

		if popErr := writeBuffer.PopContext("TransferSubscriptionsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransferSubscriptionsRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransferSubscriptionsRequest) isTransferSubscriptionsRequest() bool {
	return true
}

func (m *_TransferSubscriptionsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
