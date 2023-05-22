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

public class ComplexOfComplexUsageB implements Message {

  // Properties.
  protected final ComplexOfComplexUsageC c;
  protected final List<ComplexOfComplexUsageD> ds;

  public ComplexOfComplexUsageB(ComplexOfComplexUsageC c, List<ComplexOfComplexUsageD> ds) {
    super();
    this.c = c;
    this.ds = ds;
  }

  public ComplexOfComplexUsageC getC() {
    return c;
  }

  public List<ComplexOfComplexUsageD> getDs() {
    return ds;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ComplexOfComplexUsageB");

    // Optional Field (c) (Can be skipped, if the value is null)
    writeOptionalField("c", c, new DataWriterComplexDefault<>(writeBuffer));

    // Array Field (ds)
    writeComplexTypeArrayField("ds", ds, writeBuffer);

    writeBuffer.popContext("ComplexOfComplexUsageB");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ComplexOfComplexUsageB _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (c)
    if (c != null) {
      lengthInBits += c.getLengthInBits();
    }

    // Array field
    if (ds != null) {
      int i = 0;
      for (ComplexOfComplexUsageD element : ds) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= ds.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ComplexOfComplexUsageB staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static ComplexOfComplexUsageB staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ComplexOfComplexUsageB");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ComplexOfComplexUsageC c =
        readOptionalField(
            "c",
            new DataReaderComplexDefault<>(
                () -> ComplexOfComplexUsageC.staticParse(readBuffer), readBuffer));

    List<ComplexOfComplexUsageD> ds =
        readCountArrayField(
            "ds",
            new DataReaderComplexDefault<>(
                () -> ComplexOfComplexUsageD.staticParse(readBuffer), readBuffer),
            5);

    readBuffer.closeContext("ComplexOfComplexUsageB");
    // Create the instance
    ComplexOfComplexUsageB _complexOfComplexUsageB;
    _complexOfComplexUsageB = new ComplexOfComplexUsageB(c, ds);
    return _complexOfComplexUsageB;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ComplexOfComplexUsageB)) {
      return false;
    }
    ComplexOfComplexUsageB that = (ComplexOfComplexUsageB) o;
    return (getC() == that.getC()) && (getDs() == that.getDs()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getC(), getDs());
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
