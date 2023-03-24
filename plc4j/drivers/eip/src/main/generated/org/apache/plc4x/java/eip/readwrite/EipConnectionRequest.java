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
package org.apache.plc4x.java.eip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class EipConnectionRequest extends EipPacket implements Message {

  // Accessors for discriminator values.
  public Integer getCommand() {
    return (int) 0x0065;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  public Integer getPacketLength() {
    return 0;
  }

  // Constant values.
  public static final Integer PROTOCOLVERSION = 0x01;
  public static final Integer FLAGS = 0x00;

  public EipConnectionRequest(long sessionHandle, long status, byte[] senderContext, long options) {
    super(sessionHandle, status, senderContext, options);
  }

  public int getProtocolVersion() {
    return PROTOCOLVERSION;
  }

  public int getFlags() {
    return FLAGS;
  }

  @Override
  protected void serializeEipPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("EipConnectionRequest");

    // Const Field (protocolVersion)
    writeConstField("protocolVersion", PROTOCOLVERSION, writeUnsignedInt(writeBuffer, 16));

    // Const Field (flags)
    writeConstField("flags", FLAGS, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("EipConnectionRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EipConnectionRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (protocolVersion)
    lengthInBits += 16;

    // Const Field (flags)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static EipPacketBuilder staticParseEipPacketBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("EipConnectionRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int protocolVersion =
        readConstField(
            "protocolVersion",
            readUnsignedInt(readBuffer, 16),
            EipConnectionRequest.PROTOCOLVERSION);

    int flags =
        readConstField("flags", readUnsignedInt(readBuffer, 16), EipConnectionRequest.FLAGS);

    readBuffer.closeContext("EipConnectionRequest");
    // Create the instance
    return new EipConnectionRequestBuilderImpl();
  }

  public static class EipConnectionRequestBuilderImpl implements EipPacket.EipPacketBuilder {

    public EipConnectionRequestBuilderImpl() {}

    public EipConnectionRequest build(
        long sessionHandle, long status, byte[] senderContext, long options) {
      EipConnectionRequest eipConnectionRequest =
          new EipConnectionRequest(sessionHandle, status, senderContext, options);
      return eipConnectionRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EipConnectionRequest)) {
      return false;
    }
    EipConnectionRequest that = (EipConnectionRequest) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}