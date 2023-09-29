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

// OpcuaNodeIdServicesVariableBroker is an enum
type OpcuaNodeIdServicesVariableBroker int32

type IOpcuaNodeIdServicesVariableBroker interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableBroker_BrokerTransportQualityOfService_EnumStrings                 OpcuaNodeIdServicesVariableBroker = 15009
	OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_ResourceUri                   OpcuaNodeIdServicesVariableBroker = 15156
	OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_AuthenticationProfileUri      OpcuaNodeIdServicesVariableBroker = 15178
	OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_ResourceUri                  OpcuaNodeIdServicesVariableBroker = 15246
	OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_AuthenticationProfileUri     OpcuaNodeIdServicesVariableBroker = 15247
	OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_RequestedDeliveryGuarantee   OpcuaNodeIdServicesVariableBroker = 15249
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_ResourceUri                OpcuaNodeIdServicesVariableBroker = 15250
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_AuthenticationProfileUri   OpcuaNodeIdServicesVariableBroker = 15251
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_RequestedDeliveryGuarantee OpcuaNodeIdServicesVariableBroker = 15330
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_ResourceUri                OpcuaNodeIdServicesVariableBroker = 15334
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_AuthenticationProfileUri   OpcuaNodeIdServicesVariableBroker = 15419
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_RequestedDeliveryGuarantee OpcuaNodeIdServicesVariableBroker = 15420
	OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_QueueName                    OpcuaNodeIdServicesVariableBroker = 21137
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_QueueName                  OpcuaNodeIdServicesVariableBroker = 21139
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataQueueName          OpcuaNodeIdServicesVariableBroker = 21140
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataUpdateTime         OpcuaNodeIdServicesVariableBroker = 21141
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_QueueName                  OpcuaNodeIdServicesVariableBroker = 21143
	OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_MetaDataQueueName          OpcuaNodeIdServicesVariableBroker = 21144
)

var OpcuaNodeIdServicesVariableBrokerValues []OpcuaNodeIdServicesVariableBroker

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableBrokerValues = []OpcuaNodeIdServicesVariableBroker{
		OpcuaNodeIdServicesVariableBroker_BrokerTransportQualityOfService_EnumStrings,
		OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_ResourceUri,
		OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_AuthenticationProfileUri,
		OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_ResourceUri,
		OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_AuthenticationProfileUri,
		OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_RequestedDeliveryGuarantee,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_ResourceUri,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_AuthenticationProfileUri,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_RequestedDeliveryGuarantee,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_ResourceUri,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_AuthenticationProfileUri,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_RequestedDeliveryGuarantee,
		OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_QueueName,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_QueueName,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataQueueName,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataUpdateTime,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_QueueName,
		OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_MetaDataQueueName,
	}
}

func OpcuaNodeIdServicesVariableBrokerByValue(value int32) (enum OpcuaNodeIdServicesVariableBroker, ok bool) {
	switch value {
	case 15009:
		return OpcuaNodeIdServicesVariableBroker_BrokerTransportQualityOfService_EnumStrings, true
	case 15156:
		return OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_ResourceUri, true
	case 15178:
		return OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_AuthenticationProfileUri, true
	case 15246:
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_ResourceUri, true
	case 15247:
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_AuthenticationProfileUri, true
	case 15249:
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_RequestedDeliveryGuarantee, true
	case 15250:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_ResourceUri, true
	case 15251:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_AuthenticationProfileUri, true
	case 15330:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_RequestedDeliveryGuarantee, true
	case 15334:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_ResourceUri, true
	case 15419:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_AuthenticationProfileUri, true
	case 15420:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_RequestedDeliveryGuarantee, true
	case 21137:
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_QueueName, true
	case 21139:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_QueueName, true
	case 21140:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataQueueName, true
	case 21141:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataUpdateTime, true
	case 21143:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_QueueName, true
	case 21144:
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_MetaDataQueueName, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBrokerByName(value string) (enum OpcuaNodeIdServicesVariableBroker, ok bool) {
	switch value {
	case "BrokerTransportQualityOfService_EnumStrings":
		return OpcuaNodeIdServicesVariableBroker_BrokerTransportQualityOfService_EnumStrings, true
	case "BrokerConnectionTransportType_ResourceUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_ResourceUri, true
	case "BrokerConnectionTransportType_AuthenticationProfileUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_AuthenticationProfileUri, true
	case "BrokerWriterGroupTransportType_ResourceUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_ResourceUri, true
	case "BrokerWriterGroupTransportType_AuthenticationProfileUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_AuthenticationProfileUri, true
	case "BrokerWriterGroupTransportType_RequestedDeliveryGuarantee":
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_RequestedDeliveryGuarantee, true
	case "BrokerDataSetWriterTransportType_ResourceUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_ResourceUri, true
	case "BrokerDataSetWriterTransportType_AuthenticationProfileUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_AuthenticationProfileUri, true
	case "BrokerDataSetWriterTransportType_RequestedDeliveryGuarantee":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_RequestedDeliveryGuarantee, true
	case "BrokerDataSetReaderTransportType_ResourceUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_ResourceUri, true
	case "BrokerDataSetReaderTransportType_AuthenticationProfileUri":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_AuthenticationProfileUri, true
	case "BrokerDataSetReaderTransportType_RequestedDeliveryGuarantee":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_RequestedDeliveryGuarantee, true
	case "BrokerWriterGroupTransportType_QueueName":
		return OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_QueueName, true
	case "BrokerDataSetWriterTransportType_QueueName":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_QueueName, true
	case "BrokerDataSetWriterTransportType_MetaDataQueueName":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataQueueName, true
	case "BrokerDataSetWriterTransportType_MetaDataUpdateTime":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataUpdateTime, true
	case "BrokerDataSetReaderTransportType_QueueName":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_QueueName, true
	case "BrokerDataSetReaderTransportType_MetaDataQueueName":
		return OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_MetaDataQueueName, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBrokerKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableBrokerValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableBroker(structType any) OpcuaNodeIdServicesVariableBroker {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableBroker {
		if sOpcuaNodeIdServicesVariableBroker, ok := typ.(OpcuaNodeIdServicesVariableBroker); ok {
			return sOpcuaNodeIdServicesVariableBroker
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableBroker) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableBroker) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableBrokerParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableBroker, error) {
	return OpcuaNodeIdServicesVariableBrokerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableBrokerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableBroker, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableBroker", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableBroker")
	}
	if enum, ok := OpcuaNodeIdServicesVariableBrokerByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableBroker")
		return OpcuaNodeIdServicesVariableBroker(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableBroker) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableBroker) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableBroker", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableBroker) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableBroker_BrokerTransportQualityOfService_EnumStrings:
		return "BrokerTransportQualityOfService_EnumStrings"
	case OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_ResourceUri:
		return "BrokerConnectionTransportType_ResourceUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerConnectionTransportType_AuthenticationProfileUri:
		return "BrokerConnectionTransportType_AuthenticationProfileUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_ResourceUri:
		return "BrokerWriterGroupTransportType_ResourceUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_AuthenticationProfileUri:
		return "BrokerWriterGroupTransportType_AuthenticationProfileUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_RequestedDeliveryGuarantee:
		return "BrokerWriterGroupTransportType_RequestedDeliveryGuarantee"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_ResourceUri:
		return "BrokerDataSetWriterTransportType_ResourceUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_AuthenticationProfileUri:
		return "BrokerDataSetWriterTransportType_AuthenticationProfileUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_RequestedDeliveryGuarantee:
		return "BrokerDataSetWriterTransportType_RequestedDeliveryGuarantee"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_ResourceUri:
		return "BrokerDataSetReaderTransportType_ResourceUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_AuthenticationProfileUri:
		return "BrokerDataSetReaderTransportType_AuthenticationProfileUri"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_RequestedDeliveryGuarantee:
		return "BrokerDataSetReaderTransportType_RequestedDeliveryGuarantee"
	case OpcuaNodeIdServicesVariableBroker_BrokerWriterGroupTransportType_QueueName:
		return "BrokerWriterGroupTransportType_QueueName"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_QueueName:
		return "BrokerDataSetWriterTransportType_QueueName"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataQueueName:
		return "BrokerDataSetWriterTransportType_MetaDataQueueName"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetWriterTransportType_MetaDataUpdateTime:
		return "BrokerDataSetWriterTransportType_MetaDataUpdateTime"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_QueueName:
		return "BrokerDataSetReaderTransportType_QueueName"
	case OpcuaNodeIdServicesVariableBroker_BrokerDataSetReaderTransportType_MetaDataQueueName:
		return "BrokerDataSetReaderTransportType_MetaDataQueueName"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableBroker) String() string {
	return e.PLC4XEnumName()
}