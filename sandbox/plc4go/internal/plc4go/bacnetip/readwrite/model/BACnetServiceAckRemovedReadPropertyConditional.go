//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
	"plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetServiceAckRemovedReadPropertyConditional struct {
    BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckRemovedReadPropertyConditional interface {
    IBACnetServiceAck
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetServiceAckRemovedReadPropertyConditional) ServiceChoice() uint8 {
    return 0x0D
}

func (m BACnetServiceAckRemovedReadPropertyConditional) initialize() spi.Message {
    return m
}

func NewBACnetServiceAckRemovedReadPropertyConditional() BACnetServiceAckInitializer {
    return &BACnetServiceAckRemovedReadPropertyConditional{}
}

func CastIBACnetServiceAckRemovedReadPropertyConditional(structType interface{}) IBACnetServiceAckRemovedReadPropertyConditional {
    castFunc := func(typ interface{}) IBACnetServiceAckRemovedReadPropertyConditional {
        if iBACnetServiceAckRemovedReadPropertyConditional, ok := typ.(IBACnetServiceAckRemovedReadPropertyConditional); ok {
            return iBACnetServiceAckRemovedReadPropertyConditional
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetServiceAckRemovedReadPropertyConditional(structType interface{}) BACnetServiceAckRemovedReadPropertyConditional {
    castFunc := func(typ interface{}) BACnetServiceAckRemovedReadPropertyConditional {
        if sBACnetServiceAckRemovedReadPropertyConditional, ok := typ.(BACnetServiceAckRemovedReadPropertyConditional); ok {
            return sBACnetServiceAckRemovedReadPropertyConditional
        }
        if sBACnetServiceAckRemovedReadPropertyConditional, ok := typ.(*BACnetServiceAckRemovedReadPropertyConditional); ok {
            return *sBACnetServiceAckRemovedReadPropertyConditional
        }
        return BACnetServiceAckRemovedReadPropertyConditional{}
    }
    return castFunc(structType)
}

func (m BACnetServiceAckRemovedReadPropertyConditional) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetServiceAck.LengthInBits()

    return lengthInBits
}

func (m BACnetServiceAckRemovedReadPropertyConditional) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetServiceAckRemovedReadPropertyConditionalParse(io *utils.ReadBuffer) (BACnetServiceAckInitializer, error) {

    // Create the instance
    return NewBACnetServiceAckRemovedReadPropertyConditional(), nil
}

func (m BACnetServiceAckRemovedReadPropertyConditional) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BACnetServiceAckSerialize(io, m.BACnetServiceAck, CastIBACnetServiceAck(m), ser)
}