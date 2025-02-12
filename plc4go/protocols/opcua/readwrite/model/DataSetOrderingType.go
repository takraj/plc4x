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

// DataSetOrderingType is an enum
type DataSetOrderingType uint32

type IDataSetOrderingType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DataSetOrderingType_dataSetOrderingTypeUndefined               DataSetOrderingType = 0
	DataSetOrderingType_dataSetOrderingTypeAscendingWriterId       DataSetOrderingType = 1
	DataSetOrderingType_dataSetOrderingTypeAscendingWriterIdSingle DataSetOrderingType = 2
)

var DataSetOrderingTypeValues []DataSetOrderingType

func init() {
	_ = errors.New
	DataSetOrderingTypeValues = []DataSetOrderingType{
		DataSetOrderingType_dataSetOrderingTypeUndefined,
		DataSetOrderingType_dataSetOrderingTypeAscendingWriterId,
		DataSetOrderingType_dataSetOrderingTypeAscendingWriterIdSingle,
	}
}

func DataSetOrderingTypeByValue(value uint32) (enum DataSetOrderingType, ok bool) {
	switch value {
	case 0:
		return DataSetOrderingType_dataSetOrderingTypeUndefined, true
	case 1:
		return DataSetOrderingType_dataSetOrderingTypeAscendingWriterId, true
	case 2:
		return DataSetOrderingType_dataSetOrderingTypeAscendingWriterIdSingle, true
	}
	return 0, false
}

func DataSetOrderingTypeByName(value string) (enum DataSetOrderingType, ok bool) {
	switch value {
	case "dataSetOrderingTypeUndefined":
		return DataSetOrderingType_dataSetOrderingTypeUndefined, true
	case "dataSetOrderingTypeAscendingWriterId":
		return DataSetOrderingType_dataSetOrderingTypeAscendingWriterId, true
	case "dataSetOrderingTypeAscendingWriterIdSingle":
		return DataSetOrderingType_dataSetOrderingTypeAscendingWriterIdSingle, true
	}
	return 0, false
}

func DataSetOrderingTypeKnows(value uint32) bool {
	for _, typeValue := range DataSetOrderingTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDataSetOrderingType(structType any) DataSetOrderingType {
	castFunc := func(typ any) DataSetOrderingType {
		if sDataSetOrderingType, ok := typ.(DataSetOrderingType); ok {
			return sDataSetOrderingType
		}
		return 0
	}
	return castFunc(structType)
}

func (m DataSetOrderingType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m DataSetOrderingType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataSetOrderingTypeParse(ctx context.Context, theBytes []byte) (DataSetOrderingType, error) {
	return DataSetOrderingTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DataSetOrderingTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DataSetOrderingType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("DataSetOrderingType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DataSetOrderingType")
	}
	if enum, ok := DataSetOrderingTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for DataSetOrderingType")
		return DataSetOrderingType(val), nil
	} else {
		return enum, nil
	}
}

func (e DataSetOrderingType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DataSetOrderingType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("DataSetOrderingType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DataSetOrderingType) PLC4XEnumName() string {
	switch e {
	case DataSetOrderingType_dataSetOrderingTypeUndefined:
		return "dataSetOrderingTypeUndefined"
	case DataSetOrderingType_dataSetOrderingTypeAscendingWriterId:
		return "dataSetOrderingTypeAscendingWriterId"
	case DataSetOrderingType_dataSetOrderingTypeAscendingWriterIdSingle:
		return "dataSetOrderingTypeAscendingWriterIdSingle"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e DataSetOrderingType) String() string {
	return e.PLC4XEnumName()
}
