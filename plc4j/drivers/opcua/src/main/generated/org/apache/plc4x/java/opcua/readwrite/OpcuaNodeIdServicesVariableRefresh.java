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
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableRefresh {
  RefreshStartEventType_EventId((int) 3969L),
  RefreshStartEventType_EventType((int) 3970L),
  RefreshStartEventType_SourceNode((int) 3971L),
  RefreshStartEventType_SourceName((int) 3972L),
  RefreshStartEventType_Time((int) 3973L),
  RefreshStartEventType_ReceiveTime((int) 3974L),
  RefreshStartEventType_LocalTime((int) 3975L),
  RefreshStartEventType_Message((int) 3976L),
  RefreshStartEventType_Severity((int) 3977L),
  RefreshEndEventType_EventId((int) 3978L),
  RefreshEndEventType_EventType((int) 3979L),
  RefreshEndEventType_SourceNode((int) 3980L),
  RefreshEndEventType_SourceName((int) 3981L),
  RefreshEndEventType_Time((int) 3982L),
  RefreshEndEventType_ReceiveTime((int) 3983L),
  RefreshEndEventType_LocalTime((int) 3984L),
  RefreshEndEventType_Message((int) 3985L),
  RefreshEndEventType_Severity((int) 3986L),
  RefreshRequiredEventType_EventId((int) 3987L),
  RefreshRequiredEventType_EventType((int) 3988L),
  RefreshRequiredEventType_SourceNode((int) 3989L),
  RefreshRequiredEventType_SourceName((int) 3990L),
  RefreshRequiredEventType_Time((int) 3991L),
  RefreshRequiredEventType_ReceiveTime((int) 3992L),
  RefreshRequiredEventType_LocalTime((int) 3993L),
  RefreshRequiredEventType_Message((int) 3994L),
  RefreshRequiredEventType_Severity((int) 3995L),
  RefreshStartEventType_ConditionClassId((int) 31975L),
  RefreshStartEventType_ConditionClassName((int) 31976L),
  RefreshStartEventType_ConditionSubClassId((int) 31977L),
  RefreshStartEventType_ConditionSubClassName((int) 31978L),
  RefreshEndEventType_ConditionClassId((int) 31979L),
  RefreshEndEventType_ConditionClassName((int) 31980L),
  RefreshEndEventType_ConditionSubClassId((int) 31981L),
  RefreshEndEventType_ConditionSubClassName((int) 31982L),
  RefreshRequiredEventType_ConditionClassId((int) 31983L),
  RefreshRequiredEventType_ConditionClassName((int) 31984L),
  RefreshRequiredEventType_ConditionSubClassId((int) 31985L),
  RefreshRequiredEventType_ConditionSubClassName((int) 31986L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableRefresh> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableRefresh value : OpcuaNodeIdServicesVariableRefresh.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableRefresh(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableRefresh enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
