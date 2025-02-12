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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnIoCm_Block_IAndM0 extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.I_AND_M_0;
  }

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final int vendorId;
  protected final String orderId;
  protected final String serialNumber;
  protected final int hardwareRevision;
  protected final String softwareRevision;
  protected final int revisionCounter;
  protected final int profileId;
  protected final int profileSpecificType;
  protected final short versionMajor;
  protected final short versionMinor;
  protected final int supported;

  public PnIoCm_Block_IAndM0(
      short blockVersionHigh,
      short blockVersionLow,
      int vendorId,
      String orderId,
      String serialNumber,
      int hardwareRevision,
      String softwareRevision,
      int revisionCounter,
      int profileId,
      int profileSpecificType,
      short versionMajor,
      short versionMinor,
      int supported) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.vendorId = vendorId;
    this.orderId = orderId;
    this.serialNumber = serialNumber;
    this.hardwareRevision = hardwareRevision;
    this.softwareRevision = softwareRevision;
    this.revisionCounter = revisionCounter;
    this.profileId = profileId;
    this.profileSpecificType = profileSpecificType;
    this.versionMajor = versionMajor;
    this.versionMinor = versionMinor;
    this.supported = supported;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public int getVendorId() {
    return vendorId;
  }

  public String getOrderId() {
    return orderId;
  }

  public String getSerialNumber() {
    return serialNumber;
  }

  public int getHardwareRevision() {
    return hardwareRevision;
  }

  public String getSoftwareRevision() {
    return softwareRevision;
  }

  public int getRevisionCounter() {
    return revisionCounter;
  }

  public int getProfileId() {
    return profileId;
  }

  public int getProfileSpecificType() {
    return profileSpecificType;
  }

  public short getVersionMajor() {
    return versionMajor;
  }

  public short getVersionMinor() {
    return versionMinor;
  }

  public int getSupported() {
    return supported;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnIoCm_Block_IAndM0");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (vendorId)
    writeSimpleField(
        "vendorId",
        vendorId,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (orderId)
    writeSimpleField(
        "orderId",
        orderId,
        writeString(writeBuffer, 160),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (serialNumber)
    writeSimpleField(
        "serialNumber",
        serialNumber,
        writeString(writeBuffer, 128),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (hardwareRevision)
    writeSimpleField(
        "hardwareRevision",
        hardwareRevision,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (softwareRevision)
    writeSimpleField(
        "softwareRevision",
        softwareRevision,
        writeString(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (revisionCounter)
    writeSimpleField(
        "revisionCounter",
        revisionCounter,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (profileId)
    writeSimpleField(
        "profileId",
        profileId,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (profileSpecificType)
    writeSimpleField(
        "profileSpecificType",
        profileSpecificType,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (versionMajor)
    writeSimpleField(
        "versionMajor",
        versionMajor,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (versionMinor)
    writeSimpleField(
        "versionMinor",
        versionMinor,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (supported)
    writeSimpleField(
        "supported",
        supported,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_Block_IAndM0");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Block_IAndM0 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Simple field (vendorId)
    lengthInBits += 16;

    // Simple field (orderId)
    lengthInBits += 160;

    // Simple field (serialNumber)
    lengthInBits += 128;

    // Simple field (hardwareRevision)
    lengthInBits += 16;

    // Simple field (softwareRevision)
    lengthInBits += 32;

    // Simple field (revisionCounter)
    lengthInBits += 16;

    // Simple field (profileId)
    lengthInBits += 16;

    // Simple field (profileSpecificType)
    lengthInBits += 16;

    // Simple field (versionMajor)
    lengthInBits += 8;

    // Simple field (versionMinor)
    lengthInBits += 8;

    // Simple field (supported)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Block_IAndM0");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int vendorId =
        readSimpleField(
            "vendorId",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    String orderId =
        readSimpleField(
            "orderId", readString(readBuffer, 160), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    String serialNumber =
        readSimpleField(
            "serialNumber",
            readString(readBuffer, 128),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int hardwareRevision =
        readSimpleField(
            "hardwareRevision",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    String softwareRevision =
        readSimpleField(
            "softwareRevision",
            readString(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int revisionCounter =
        readSimpleField(
            "revisionCounter",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int profileId =
        readSimpleField(
            "profileId",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int profileSpecificType =
        readSimpleField(
            "profileSpecificType",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short versionMajor =
        readSimpleField(
            "versionMajor",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short versionMinor =
        readSimpleField(
            "versionMinor",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int supported =
        readSimpleField(
            "supported",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_Block_IAndM0");
    // Create the instance
    return new PnIoCm_Block_IAndM0BuilderImpl(
        blockVersionHigh,
        blockVersionLow,
        vendorId,
        orderId,
        serialNumber,
        hardwareRevision,
        softwareRevision,
        revisionCounter,
        profileId,
        profileSpecificType,
        versionMajor,
        versionMinor,
        supported);
  }

  public static class PnIoCm_Block_IAndM0BuilderImpl implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final int vendorId;
    private final String orderId;
    private final String serialNumber;
    private final int hardwareRevision;
    private final String softwareRevision;
    private final int revisionCounter;
    private final int profileId;
    private final int profileSpecificType;
    private final short versionMajor;
    private final short versionMinor;
    private final int supported;

    public PnIoCm_Block_IAndM0BuilderImpl(
        short blockVersionHigh,
        short blockVersionLow,
        int vendorId,
        String orderId,
        String serialNumber,
        int hardwareRevision,
        String softwareRevision,
        int revisionCounter,
        int profileId,
        int profileSpecificType,
        short versionMajor,
        short versionMinor,
        int supported) {
      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.vendorId = vendorId;
      this.orderId = orderId;
      this.serialNumber = serialNumber;
      this.hardwareRevision = hardwareRevision;
      this.softwareRevision = softwareRevision;
      this.revisionCounter = revisionCounter;
      this.profileId = profileId;
      this.profileSpecificType = profileSpecificType;
      this.versionMajor = versionMajor;
      this.versionMinor = versionMinor;
      this.supported = supported;
    }

    public PnIoCm_Block_IAndM0 build() {
      PnIoCm_Block_IAndM0 pnIoCm_Block_IAndM0 =
          new PnIoCm_Block_IAndM0(
              blockVersionHigh,
              blockVersionLow,
              vendorId,
              orderId,
              serialNumber,
              hardwareRevision,
              softwareRevision,
              revisionCounter,
              profileId,
              profileSpecificType,
              versionMajor,
              versionMinor,
              supported);
      return pnIoCm_Block_IAndM0;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block_IAndM0)) {
      return false;
    }
    PnIoCm_Block_IAndM0 that = (PnIoCm_Block_IAndM0) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getVendorId() == that.getVendorId())
        && (getOrderId() == that.getOrderId())
        && (getSerialNumber() == that.getSerialNumber())
        && (getHardwareRevision() == that.getHardwareRevision())
        && (getSoftwareRevision() == that.getSoftwareRevision())
        && (getRevisionCounter() == that.getRevisionCounter())
        && (getProfileId() == that.getProfileId())
        && (getProfileSpecificType() == that.getProfileSpecificType())
        && (getVersionMajor() == that.getVersionMajor())
        && (getVersionMinor() == that.getVersionMinor())
        && (getSupported() == that.getSupported())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getBlockVersionHigh(),
        getBlockVersionLow(),
        getVendorId(),
        getOrderId(),
        getSerialNumber(),
        getHardwareRevision(),
        getSoftwareRevision(),
        getRevisionCounter(),
        getProfileId(),
        getProfileSpecificType(),
        getVersionMajor(),
        getVersionMinor(),
        getSupported());
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
