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

// MonitoredItemModifyRequest is the corresponding interface of MonitoredItemModifyRequest
type MonitoredItemModifyRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetMonitoredItemId returns MonitoredItemId (property field)
	GetMonitoredItemId() uint32
	// GetRequestedParameters returns RequestedParameters (property field)
	GetRequestedParameters() ExtensionObjectDefinition
}

// MonitoredItemModifyRequestExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemModifyRequest.
// This is useful for switch cases.
type MonitoredItemModifyRequestExactly interface {
	MonitoredItemModifyRequest
	isMonitoredItemModifyRequest() bool
}

// _MonitoredItemModifyRequest is the data-structure of this message
type _MonitoredItemModifyRequest struct {
	*_ExtensionObjectDefinition
	MonitoredItemId     uint32
	RequestedParameters ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemModifyRequest) GetIdentifier() string {
	return "757"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemModifyRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoredItemModifyRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemModifyRequest) GetMonitoredItemId() uint32 {
	return m.MonitoredItemId
}

func (m *_MonitoredItemModifyRequest) GetRequestedParameters() ExtensionObjectDefinition {
	return m.RequestedParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemModifyRequest factory function for _MonitoredItemModifyRequest
func NewMonitoredItemModifyRequest(monitoredItemId uint32, requestedParameters ExtensionObjectDefinition) *_MonitoredItemModifyRequest {
	_result := &_MonitoredItemModifyRequest{
		MonitoredItemId:            monitoredItemId,
		RequestedParameters:        requestedParameters,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemModifyRequest(structType any) MonitoredItemModifyRequest {
	if casted, ok := structType.(MonitoredItemModifyRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemModifyRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemModifyRequest) GetTypeName() string {
	return "MonitoredItemModifyRequest"
}

func (m *_MonitoredItemModifyRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (monitoredItemId)
	lengthInBits += 32

	// Simple field (requestedParameters)
	lengthInBits += m.RequestedParameters.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemModifyRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredItemModifyRequestParse(ctx context.Context, theBytes []byte, identifier string) (MonitoredItemModifyRequest, error) {
	return MonitoredItemModifyRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoredItemModifyRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoredItemModifyRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoredItemModifyRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemModifyRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (monitoredItemId)
	_monitoredItemId, _monitoredItemIdErr := readBuffer.ReadUint32("monitoredItemId", 32)
	if _monitoredItemIdErr != nil {
		return nil, errors.Wrap(_monitoredItemIdErr, "Error parsing 'monitoredItemId' field of MonitoredItemModifyRequest")
	}
	monitoredItemId := _monitoredItemId

	// Simple Field (requestedParameters)
	if pullErr := readBuffer.PullContext("requestedParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestedParameters")
	}
	_requestedParameters, _requestedParametersErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("742"))
	if _requestedParametersErr != nil {
		return nil, errors.Wrap(_requestedParametersErr, "Error parsing 'requestedParameters' field of MonitoredItemModifyRequest")
	}
	requestedParameters := _requestedParameters.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestedParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestedParameters")
	}

	if closeErr := readBuffer.CloseContext("MonitoredItemModifyRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemModifyRequest")
	}

	// Create a partially initialized instance
	_child := &_MonitoredItemModifyRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		MonitoredItemId:            monitoredItemId,
		RequestedParameters:        requestedParameters,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredItemModifyRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemModifyRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemModifyRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemModifyRequest")
		}

		// Simple Field (monitoredItemId)
		monitoredItemId := uint32(m.GetMonitoredItemId())
		_monitoredItemIdErr := writeBuffer.WriteUint32("monitoredItemId", 32, (monitoredItemId))
		if _monitoredItemIdErr != nil {
			return errors.Wrap(_monitoredItemIdErr, "Error serializing 'monitoredItemId' field")
		}

		// Simple Field (requestedParameters)
		if pushErr := writeBuffer.PushContext("requestedParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestedParameters")
		}
		_requestedParametersErr := writeBuffer.WriteSerializable(ctx, m.GetRequestedParameters())
		if popErr := writeBuffer.PopContext("requestedParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestedParameters")
		}
		if _requestedParametersErr != nil {
			return errors.Wrap(_requestedParametersErr, "Error serializing 'requestedParameters' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemModifyRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemModifyRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemModifyRequest) isMonitoredItemModifyRequest() bool {
	return true
}

func (m *_MonitoredItemModifyRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
