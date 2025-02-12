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

// OpcuaNodeIdServicesVariableSemantic is an enum
type OpcuaNodeIdServicesVariableSemantic int32

type IOpcuaNodeIdServicesVariableSemantic interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes               OpcuaNodeIdServicesVariableSemantic = 2739
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventId               OpcuaNodeIdServicesVariableSemantic = 3689
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventType             OpcuaNodeIdServicesVariableSemantic = 3690
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceNode            OpcuaNodeIdServicesVariableSemantic = 3691
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceName            OpcuaNodeIdServicesVariableSemantic = 3692
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Time                  OpcuaNodeIdServicesVariableSemantic = 3693
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ReceiveTime           OpcuaNodeIdServicesVariableSemantic = 3694
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_LocalTime             OpcuaNodeIdServicesVariableSemantic = 3695
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Message               OpcuaNodeIdServicesVariableSemantic = 3696
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Severity              OpcuaNodeIdServicesVariableSemantic = 3697
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassId      OpcuaNodeIdServicesVariableSemantic = 31895
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassName    OpcuaNodeIdServicesVariableSemantic = 31896
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassId   OpcuaNodeIdServicesVariableSemantic = 31897
	OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassName OpcuaNodeIdServicesVariableSemantic = 31898
)

var OpcuaNodeIdServicesVariableSemanticValues []OpcuaNodeIdServicesVariableSemantic

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableSemanticValues = []OpcuaNodeIdServicesVariableSemantic{
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventId,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventType,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceNode,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceName,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Time,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_LocalTime,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Message,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Severity,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassName,
	}
}

func OpcuaNodeIdServicesVariableSemanticByValue(value int32) (enum OpcuaNodeIdServicesVariableSemantic, ok bool) {
	switch value {
	case 2739:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes, true
	case 31895:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassId, true
	case 31896:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassName, true
	case 31897:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassId, true
	case 31898:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassName, true
	case 3689:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventId, true
	case 3690:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventType, true
	case 3691:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceNode, true
	case 3692:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceName, true
	case 3693:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Time, true
	case 3694:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ReceiveTime, true
	case 3695:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_LocalTime, true
	case 3696:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Message, true
	case 3697:
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSemanticByName(value string) (enum OpcuaNodeIdServicesVariableSemantic, ok bool) {
	switch value {
	case "SemanticChangeEventType_Changes":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes, true
	case "SemanticChangeEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassId, true
	case "SemanticChangeEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassName, true
	case "SemanticChangeEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassId, true
	case "SemanticChangeEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassName, true
	case "SemanticChangeEventType_EventId":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventId, true
	case "SemanticChangeEventType_EventType":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventType, true
	case "SemanticChangeEventType_SourceNode":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceNode, true
	case "SemanticChangeEventType_SourceName":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceName, true
	case "SemanticChangeEventType_Time":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Time, true
	case "SemanticChangeEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ReceiveTime, true
	case "SemanticChangeEventType_LocalTime":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_LocalTime, true
	case "SemanticChangeEventType_Message":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Message, true
	case "SemanticChangeEventType_Severity":
		return OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSemanticKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableSemanticValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableSemantic(structType any) OpcuaNodeIdServicesVariableSemantic {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableSemantic {
		if sOpcuaNodeIdServicesVariableSemantic, ok := typ.(OpcuaNodeIdServicesVariableSemantic); ok {
			return sOpcuaNodeIdServicesVariableSemantic
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableSemantic) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableSemantic) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableSemanticParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableSemantic, error) {
	return OpcuaNodeIdServicesVariableSemanticParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableSemanticParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableSemantic, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableSemantic", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableSemantic")
	}
	if enum, ok := OpcuaNodeIdServicesVariableSemanticByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableSemantic")
		return OpcuaNodeIdServicesVariableSemantic(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableSemantic) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableSemantic) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableSemantic", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableSemantic) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Changes:
		return "SemanticChangeEventType_Changes"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassId:
		return "SemanticChangeEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionClassName:
		return "SemanticChangeEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassId:
		return "SemanticChangeEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ConditionSubClassName:
		return "SemanticChangeEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventId:
		return "SemanticChangeEventType_EventId"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_EventType:
		return "SemanticChangeEventType_EventType"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceNode:
		return "SemanticChangeEventType_SourceNode"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_SourceName:
		return "SemanticChangeEventType_SourceName"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Time:
		return "SemanticChangeEventType_Time"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_ReceiveTime:
		return "SemanticChangeEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_LocalTime:
		return "SemanticChangeEventType_LocalTime"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Message:
		return "SemanticChangeEventType_Message"
	case OpcuaNodeIdServicesVariableSemantic_SemanticChangeEventType_Severity:
		return "SemanticChangeEventType_Severity"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableSemantic) String() string {
	return e.PLC4XEnumName()
}
