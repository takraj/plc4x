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
package org.apache.plc4x.java.test.readwrite;

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

public class LengthLoopTest implements Message {

  // Properties.
  protected final int commandType;

  public LengthLoopTest(int commandType) {
    super();
    this.commandType = commandType;
  }

  public int getCommandType() {
    return commandType;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("LengthLoopTest");

    // Simple Field (commandType)
    writeSimpleField("commandType", commandType, writeUnsignedInt(writeBuffer, 16));

    // Implicit Field (len) (Used for parsing, but its value is not stored as it's implicitly given
    // by the objects content)
    int len = (int) ((getLengthInBytes()) - (8));
    writeImplicitField("len", len, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("LengthLoopTest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    LengthLoopTest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (commandType)
    lengthInBits += 16;

    // Implicit Field (len)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static LengthLoopTest staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static LengthLoopTest staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("LengthLoopTest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int commandType = readSimpleField("commandType", readUnsignedInt(readBuffer, 16));

    int len = readImplicitField("len", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("LengthLoopTest");
    // Create the instance
    LengthLoopTest _lengthLoopTest;
    _lengthLoopTest = new LengthLoopTest(commandType);
    return _lengthLoopTest;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LengthLoopTest)) {
      return false;
    }
    LengthLoopTest that = (LengthLoopTest) o;
    return (getCommandType() == that.getCommandType()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCommandType());
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
