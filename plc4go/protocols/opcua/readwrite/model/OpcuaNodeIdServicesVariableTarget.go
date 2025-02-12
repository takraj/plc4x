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

// OpcuaNodeIdServicesVariableTarget is an enum
type OpcuaNodeIdServicesVariableTarget int32

type IOpcuaNodeIdServicesVariableTarget interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableTarget_TargetVariablesType_TargetVariables                                OpcuaNodeIdServicesVariableTarget = 15114
	OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_InputArguments              OpcuaNodeIdServicesVariableTarget = 15116
	OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_OutputArguments             OpcuaNodeIdServicesVariableTarget = 15117
	OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_InputArguments           OpcuaNodeIdServicesVariableTarget = 15119
	OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_OutputArguments          OpcuaNodeIdServicesVariableTarget = 15120
	OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_InputArguments     OpcuaNodeIdServicesVariableTarget = 15122
	OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_OutputArguments    OpcuaNodeIdServicesVariableTarget = 15123
	OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_InputArguments  OpcuaNodeIdServicesVariableTarget = 15125
	OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_OutputArguments OpcuaNodeIdServicesVariableTarget = 15126
)

var OpcuaNodeIdServicesVariableTargetValues []OpcuaNodeIdServicesVariableTarget

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTargetValues = []OpcuaNodeIdServicesVariableTarget{
		OpcuaNodeIdServicesVariableTarget_TargetVariablesType_TargetVariables,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_InputArguments,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_OutputArguments,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_InputArguments,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_OutputArguments,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_InputArguments,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_InputArguments,
		OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableTargetByValue(value int32) (enum OpcuaNodeIdServicesVariableTarget, ok bool) {
	switch value {
	case 15114:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_TargetVariables, true
	case 15116:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_InputArguments, true
	case 15117:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_OutputArguments, true
	case 15119:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_InputArguments, true
	case 15120:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_OutputArguments, true
	case 15122:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_InputArguments, true
	case 15123:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_OutputArguments, true
	case 15125:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_InputArguments, true
	case 15126:
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTargetByName(value string) (enum OpcuaNodeIdServicesVariableTarget, ok bool) {
	switch value {
	case "TargetVariablesType_TargetVariables":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_TargetVariables, true
	case "TargetVariablesType_AddTargetVariables_InputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_InputArguments, true
	case "TargetVariablesType_AddTargetVariables_OutputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_OutputArguments, true
	case "TargetVariablesType_RemoveTargetVariables_InputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_InputArguments, true
	case "TargetVariablesType_RemoveTargetVariables_OutputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_OutputArguments, true
	case "TargetVariablesTypeAddTargetVariablesMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_InputArguments, true
	case "TargetVariablesTypeAddTargetVariablesMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_OutputArguments, true
	case "TargetVariablesTypeRemoveTargetVariablesMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_InputArguments, true
	case "TargetVariablesTypeRemoveTargetVariablesMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTargetKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTargetValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableTarget(structType any) OpcuaNodeIdServicesVariableTarget {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTarget {
		if sOpcuaNodeIdServicesVariableTarget, ok := typ.(OpcuaNodeIdServicesVariableTarget); ok {
			return sOpcuaNodeIdServicesVariableTarget
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTarget) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTarget) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTargetParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTarget, error) {
	return OpcuaNodeIdServicesVariableTargetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTargetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTarget, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTarget", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTarget")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTargetByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTarget")
		return OpcuaNodeIdServicesVariableTarget(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTarget) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTarget) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTarget", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTarget) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesType_TargetVariables:
		return "TargetVariablesType_TargetVariables"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_InputArguments:
		return "TargetVariablesType_AddTargetVariables_InputArguments"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesType_AddTargetVariables_OutputArguments:
		return "TargetVariablesType_AddTargetVariables_OutputArguments"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_InputArguments:
		return "TargetVariablesType_RemoveTargetVariables_InputArguments"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesType_RemoveTargetVariables_OutputArguments:
		return "TargetVariablesType_RemoveTargetVariables_OutputArguments"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_InputArguments:
		return "TargetVariablesTypeAddTargetVariablesMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeAddTargetVariablesMethodType_OutputArguments:
		return "TargetVariablesTypeAddTargetVariablesMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_InputArguments:
		return "TargetVariablesTypeRemoveTargetVariablesMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableTarget_TargetVariablesTypeRemoveTargetVariablesMethodType_OutputArguments:
		return "TargetVariablesTypeRemoveTargetVariablesMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTarget) String() string {
	return e.PLC4XEnumName()
}
