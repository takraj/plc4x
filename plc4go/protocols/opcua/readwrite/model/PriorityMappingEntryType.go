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

// PriorityMappingEntryType is the corresponding interface of PriorityMappingEntryType
type PriorityMappingEntryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetMappingUri returns MappingUri (property field)
	GetMappingUri() PascalString
	// GetPriorityLabel returns PriorityLabel (property field)
	GetPriorityLabel() PascalString
	// GetPriorityValue_PCP returns PriorityValue_PCP (property field)
	GetPriorityValue_PCP() uint8
	// GetPriorityValue_DSCP returns PriorityValue_DSCP (property field)
	GetPriorityValue_DSCP() uint32
}

// PriorityMappingEntryTypeExactly can be used when we want exactly this type and not a type which fulfills PriorityMappingEntryType.
// This is useful for switch cases.
type PriorityMappingEntryTypeExactly interface {
	PriorityMappingEntryType
	isPriorityMappingEntryType() bool
}

// _PriorityMappingEntryType is the data-structure of this message
type _PriorityMappingEntryType struct {
	*_ExtensionObjectDefinition
	MappingUri         PascalString
	PriorityLabel      PascalString
	PriorityValue_PCP  uint8
	PriorityValue_DSCP uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PriorityMappingEntryType) GetIdentifier() string {
	return "25222"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PriorityMappingEntryType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PriorityMappingEntryType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PriorityMappingEntryType) GetMappingUri() PascalString {
	return m.MappingUri
}

func (m *_PriorityMappingEntryType) GetPriorityLabel() PascalString {
	return m.PriorityLabel
}

func (m *_PriorityMappingEntryType) GetPriorityValue_PCP() uint8 {
	return m.PriorityValue_PCP
}

func (m *_PriorityMappingEntryType) GetPriorityValue_DSCP() uint32 {
	return m.PriorityValue_DSCP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPriorityMappingEntryType factory function for _PriorityMappingEntryType
func NewPriorityMappingEntryType(mappingUri PascalString, priorityLabel PascalString, priorityValue_PCP uint8, priorityValue_DSCP uint32) *_PriorityMappingEntryType {
	_result := &_PriorityMappingEntryType{
		MappingUri:                 mappingUri,
		PriorityLabel:              priorityLabel,
		PriorityValue_PCP:          priorityValue_PCP,
		PriorityValue_DSCP:         priorityValue_DSCP,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPriorityMappingEntryType(structType any) PriorityMappingEntryType {
	if casted, ok := structType.(PriorityMappingEntryType); ok {
		return casted
	}
	if casted, ok := structType.(*PriorityMappingEntryType); ok {
		return *casted
	}
	return nil
}

func (m *_PriorityMappingEntryType) GetTypeName() string {
	return "PriorityMappingEntryType"
}

func (m *_PriorityMappingEntryType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (mappingUri)
	lengthInBits += m.MappingUri.GetLengthInBits(ctx)

	// Simple field (priorityLabel)
	lengthInBits += m.PriorityLabel.GetLengthInBits(ctx)

	// Simple field (priorityValue_PCP)
	lengthInBits += 8

	// Simple field (priorityValue_DSCP)
	lengthInBits += 32

	return lengthInBits
}

func (m *_PriorityMappingEntryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PriorityMappingEntryTypeParse(ctx context.Context, theBytes []byte, identifier string) (PriorityMappingEntryType, error) {
	return PriorityMappingEntryTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PriorityMappingEntryTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PriorityMappingEntryType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PriorityMappingEntryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PriorityMappingEntryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (mappingUri)
	if pullErr := readBuffer.PullContext("mappingUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for mappingUri")
	}
	_mappingUri, _mappingUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _mappingUriErr != nil {
		return nil, errors.Wrap(_mappingUriErr, "Error parsing 'mappingUri' field of PriorityMappingEntryType")
	}
	mappingUri := _mappingUri.(PascalString)
	if closeErr := readBuffer.CloseContext("mappingUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for mappingUri")
	}

	// Simple Field (priorityLabel)
	if pullErr := readBuffer.PullContext("priorityLabel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priorityLabel")
	}
	_priorityLabel, _priorityLabelErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _priorityLabelErr != nil {
		return nil, errors.Wrap(_priorityLabelErr, "Error parsing 'priorityLabel' field of PriorityMappingEntryType")
	}
	priorityLabel := _priorityLabel.(PascalString)
	if closeErr := readBuffer.CloseContext("priorityLabel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priorityLabel")
	}

	// Simple Field (priorityValue_PCP)
	_priorityValue_PCP, _priorityValue_PCPErr := readBuffer.ReadUint8("priorityValue_PCP", 8)
	if _priorityValue_PCPErr != nil {
		return nil, errors.Wrap(_priorityValue_PCPErr, "Error parsing 'priorityValue_PCP' field of PriorityMappingEntryType")
	}
	priorityValue_PCP := _priorityValue_PCP

	// Simple Field (priorityValue_DSCP)
	_priorityValue_DSCP, _priorityValue_DSCPErr := readBuffer.ReadUint32("priorityValue_DSCP", 32)
	if _priorityValue_DSCPErr != nil {
		return nil, errors.Wrap(_priorityValue_DSCPErr, "Error parsing 'priorityValue_DSCP' field of PriorityMappingEntryType")
	}
	priorityValue_DSCP := _priorityValue_DSCP

	if closeErr := readBuffer.CloseContext("PriorityMappingEntryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PriorityMappingEntryType")
	}

	// Create a partially initialized instance
	_child := &_PriorityMappingEntryType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		MappingUri:                 mappingUri,
		PriorityLabel:              priorityLabel,
		PriorityValue_PCP:          priorityValue_PCP,
		PriorityValue_DSCP:         priorityValue_DSCP,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PriorityMappingEntryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PriorityMappingEntryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PriorityMappingEntryType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PriorityMappingEntryType")
		}

		// Simple Field (mappingUri)
		if pushErr := writeBuffer.PushContext("mappingUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for mappingUri")
		}
		_mappingUriErr := writeBuffer.WriteSerializable(ctx, m.GetMappingUri())
		if popErr := writeBuffer.PopContext("mappingUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for mappingUri")
		}
		if _mappingUriErr != nil {
			return errors.Wrap(_mappingUriErr, "Error serializing 'mappingUri' field")
		}

		// Simple Field (priorityLabel)
		if pushErr := writeBuffer.PushContext("priorityLabel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priorityLabel")
		}
		_priorityLabelErr := writeBuffer.WriteSerializable(ctx, m.GetPriorityLabel())
		if popErr := writeBuffer.PopContext("priorityLabel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priorityLabel")
		}
		if _priorityLabelErr != nil {
			return errors.Wrap(_priorityLabelErr, "Error serializing 'priorityLabel' field")
		}

		// Simple Field (priorityValue_PCP)
		priorityValue_PCP := uint8(m.GetPriorityValue_PCP())
		_priorityValue_PCPErr := writeBuffer.WriteUint8("priorityValue_PCP", 8, (priorityValue_PCP))
		if _priorityValue_PCPErr != nil {
			return errors.Wrap(_priorityValue_PCPErr, "Error serializing 'priorityValue_PCP' field")
		}

		// Simple Field (priorityValue_DSCP)
		priorityValue_DSCP := uint32(m.GetPriorityValue_DSCP())
		_priorityValue_DSCPErr := writeBuffer.WriteUint32("priorityValue_DSCP", 32, (priorityValue_DSCP))
		if _priorityValue_DSCPErr != nil {
			return errors.Wrap(_priorityValue_DSCPErr, "Error serializing 'priorityValue_DSCP' field")
		}

		if popErr := writeBuffer.PopContext("PriorityMappingEntryType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PriorityMappingEntryType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PriorityMappingEntryType) isPriorityMappingEntryType() bool {
	return true
}

func (m *_PriorityMappingEntryType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
