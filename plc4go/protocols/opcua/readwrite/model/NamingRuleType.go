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

// NamingRuleType is an enum
type NamingRuleType uint32

type INamingRuleType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	NamingRuleType_namingRuleTypeMandatory  NamingRuleType = 1
	NamingRuleType_namingRuleTypeOptional   NamingRuleType = 2
	NamingRuleType_namingRuleTypeConstraint NamingRuleType = 3
)

var NamingRuleTypeValues []NamingRuleType

func init() {
	_ = errors.New
	NamingRuleTypeValues = []NamingRuleType{
		NamingRuleType_namingRuleTypeMandatory,
		NamingRuleType_namingRuleTypeOptional,
		NamingRuleType_namingRuleTypeConstraint,
	}
}

func NamingRuleTypeByValue(value uint32) (enum NamingRuleType, ok bool) {
	switch value {
	case 1:
		return NamingRuleType_namingRuleTypeMandatory, true
	case 2:
		return NamingRuleType_namingRuleTypeOptional, true
	case 3:
		return NamingRuleType_namingRuleTypeConstraint, true
	}
	return 0, false
}

func NamingRuleTypeByName(value string) (enum NamingRuleType, ok bool) {
	switch value {
	case "namingRuleTypeMandatory":
		return NamingRuleType_namingRuleTypeMandatory, true
	case "namingRuleTypeOptional":
		return NamingRuleType_namingRuleTypeOptional, true
	case "namingRuleTypeConstraint":
		return NamingRuleType_namingRuleTypeConstraint, true
	}
	return 0, false
}

func NamingRuleTypeKnows(value uint32) bool {
	for _, typeValue := range NamingRuleTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNamingRuleType(structType any) NamingRuleType {
	castFunc := func(typ any) NamingRuleType {
		if sNamingRuleType, ok := typ.(NamingRuleType); ok {
			return sNamingRuleType
		}
		return 0
	}
	return castFunc(structType)
}

func (m NamingRuleType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m NamingRuleType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NamingRuleTypeParse(ctx context.Context, theBytes []byte) (NamingRuleType, error) {
	return NamingRuleTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NamingRuleTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NamingRuleType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("NamingRuleType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NamingRuleType")
	}
	if enum, ok := NamingRuleTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for NamingRuleType")
		return NamingRuleType(val), nil
	} else {
		return enum, nil
	}
}

func (e NamingRuleType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e NamingRuleType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("NamingRuleType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NamingRuleType) PLC4XEnumName() string {
	switch e {
	case NamingRuleType_namingRuleTypeMandatory:
		return "namingRuleTypeMandatory"
	case NamingRuleType_namingRuleTypeOptional:
		return "namingRuleTypeOptional"
	case NamingRuleType_namingRuleTypeConstraint:
		return "namingRuleTypeConstraint"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e NamingRuleType) String() string {
	return e.PLC4XEnumName()
}
