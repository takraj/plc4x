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
package org.apache.plc4x.java.bacnetip.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetEngineeringUnits {
  METERS_PER_SECOND_PER_SECOND((long) 166L),
  SQUARE_METERS((long) 0L),
  SQUARE_CENTIMETERS((long) 116L),
  SQUARE_FEET((long) 1L),
  SQUARE_INCHES((long) 115L),
  CURRENCY1((long) 105L),
  CURRENCY2((long) 106L),
  CURRENCY3((long) 107L),
  CURRENCY4((long) 108L),
  CURRENCY5((long) 109L),
  CURRENCY6((long) 110L),
  CURRENCY7((long) 111L),
  CURRENCY8((long) 112L),
  CURRENCY9((long) 113L),
  CURRENCY10((long) 114L),
  MILLIAMPERES((long) 2L),
  AMPERES((long) 3L),
  AMPERES_PER_METER((long) 167L),
  AMPERES_PER_SQUARE_METER((long) 168L),
  AMPERE_SQUARE_METERS((long) 169L),
  DECIBELS((long) 199L),
  DECIBELS_MILLIVOLT((long) 200L),
  DECIBELS_VOLT((long) 201L),
  FARADS((long) 170L),
  HENRYS((long) 171L),
  OHMS((long) 4L),
  OHM_METER_SQUARED_PER_METER((long) 237L),
  OHM_METERS((long) 172L),
  MILLIOHMS((long) 145L),
  KILOHMS((long) 122L),
  MEGOHMS((long) 123L),
  MICROSIEMENS((long) 190L),
  MILLISIEMENS((long) 202L),
  SIEMENS((long) 173L),
  SIEMENS_PER_METER((long) 174L),
  TESLAS((long) 175L),
  VOLTS((long) 5L),
  MILLIVOLTS((long) 124L),
  KILOVOLTS((long) 6L),
  MEGAVOLTS((long) 7L),
  VOLT_AMPERES((long) 8L),
  KILOVOLT_AMPERES((long) 9L),
  MEGAVOLT_AMPERES((long) 10L),
  VOLT_AMPERES_REACTIVE((long) 11L),
  KILOVOLT_AMPERES_REACTIVE((long) 12L),
  MEGAVOLT_AMPERES_REACTIVE((long) 13L),
  VOLTS_PER_DEGREE_KELVIN((long) 176L),
  VOLTS_PER_METER((long) 177L),
  DEGREES_PHASE((long) 14L),
  POWER_FACTOR((long) 15L),
  WEBERS((long) 178L),
  AMPERE_SECONDS((long) 238L),
  VOLT_AMPERE_HOURS((long) 239L),
  KILOVOLT_AMPERE_HOURS((long) 240L),
  MEGAVOLT_AMPERE_HOURS((long) 241L),
  VOLT_AMPERE_HOURS_REACTIVE((long) 242L),
  KILOVOLT_AMPERE_HOURS_REACTIVE((long) 243L),
  MEGAVOLT_AMPERE_HOURS_REACTIVE((long) 244L),
  VOLT_SQUARE_HOURS((long) 245L),
  AMPERE_SQUARE_HOURS((long) 246L),
  JOULES((long) 16L),
  KILOJOULES((long) 17L),
  KILOJOULES_PER_KILOGRAM((long) 125L),
  MEGAJOULES((long) 126L),
  WATT_HOURS((long) 18L),
  KILOWATT_HOURS((long) 19L),
  MEGAWATT_HOURS((long) 146L),
  WATT_HOURS_REACTIVE((long) 203L),
  KILOWATT_HOURS_REACTIVE((long) 204L),
  MEGAWATT_HOURS_REACTIVE((long) 205L),
  BTUS((long) 20L),
  KILO_BTUS((long) 147L),
  MEGA_BTUS((long) 148L),
  THERMS((long) 21L),
  TON_HOURS((long) 22L),
  JOULES_PER_KILOGRAM_DRY_AIR((long) 23L),
  KILOJOULES_PER_KILOGRAM_DRY_AIR((long) 149L),
  MEGAJOULES_PER_KILOGRAM_DRY_AIR((long) 150L),
  BTUS_PER_POUND_DRY_AIR((long) 24L),
  BTUS_PER_POUND((long) 117L),
  GRAMS_OF_WATER_PER_KILOGRAM_DRY_AIR((long) 28L),
  PERCENT_RELATIVE_HUMIDITY((long) 29L),
  MICROMETERS((long) 194L),
  MILLIMETERS((long) 30L),
  CENTIMETERS((long) 118L),
  KILOMETERS((long) 193L),
  METERS((long) 31L),
  INCHES((long) 32L),
  FEET((long) 33L),
  CANDELAS((long) 179L),
  CANDELAS_PER_SQUARE_METER((long) 180L),
  WATTS_PER_SQUARE_FOOT((long) 34L),
  WATTS_PER_SQUARE_METER((long) 35L),
  LUMENS((long) 36L),
  LUXES((long) 37L),
  FOOT_CANDLES((long) 38L),
  MILLIGRAMS((long) 196L),
  GRAMS((long) 195L),
  KILOGRAMS((long) 39L),
  POUNDS_MASS((long) 40L),
  TONS((long) 41L),
  GRAMS_PER_SECOND((long) 154L),
  GRAMS_PER_MINUTE((long) 155L),
  KILOGRAMS_PER_SECOND((long) 42L),
  KILOGRAMS_PER_MINUTE((long) 43L),
  KILOGRAMS_PER_HOUR((long) 44L),
  POUNDS_MASS_PER_SECOND((long) 119L),
  POUNDS_MASS_PER_MINUTE((long) 45L),
  POUNDS_MASS_PER_HOUR((long) 46L),
  TONS_PER_HOUR((long) 156L),
  IWATTS((long) 132L),
  WATTS((long) 47L),
  KILOWATTS((long) 48L),
  MEGAWATTS((long) 49L),
  BTUS_PER_HOUR((long) 50L),
  KILO_BTUS_PER_HOUR((long) 157L),
  JOULE_PER_HOURS((long) 247L),
  HORSEPOWER((long) 51L),
  TONS_REFRIGERATION((long) 52L),
  PASCALS((long) 53L),
  HECTOPASCALS((long) 133L),
  KILOPASCALS((long) 54L),
  MILLIBARS((long) 134L),
  BARS((long) 55L),
  POUNDS_FORCE_PER_SQUARE_INCH((long) 56L),
  MILLIMETERS_OF_WATER((long) 206L),
  CENTIMETERS_OF_WATER((long) 57L),
  INCHES_OF_WATER((long) 58L),
  MILLIMETERS_OF_MERCURY((long) 59L),
  CENTIMETERS_OF_MERCURY((long) 60L),
  INCHES_OF_MERCURY((long) 61L),
  DEGREES_CELSIUS((long) 62L),
  DEGREES_KELVIN((long) 63L),
  DEGREES_KELVIN_PER_HOUR((long) 181L),
  DEGREES_KELVIN_PER_MINUTE((long) 182L),
  DEGREES_FAHRENHEIT((long) 64L),
  DEGREE_DAYS_CELSIUS((long) 65L),
  DEGREE_DAYS_FAHRENHEIT((long) 66L),
  DELTA_DEGREES_FAHRENHEIT((long) 120L),
  DELTA_DEGREES_KELVIN((long) 121L),
  YEARS((long) 67L),
  MONTHS((long) 68L),
  WEEKS((long) 69L),
  DAYS((long) 70L),
  HOURS((long) 71L),
  MINUTES((long) 72L),
  SECONDS((long) 73L),
  HUNDREDTHS_SECONDS((long) 158L),
  MILLISECONDS((long) 159L),
  NEWTON_METERS((long) 160L),
  MILLIMETERS_PER_SECOND((long) 161L),
  MILLIMETERS_PER_MINUTE((long) 162L),
  METERS_PER_SECOND((long) 74L),
  METERS_PER_MINUTE((long) 163L),
  METERS_PER_HOUR((long) 164L),
  KILOMETERS_PER_HOUR((long) 75L),
  FEET_PER_SECOND((long) 76L),
  FEET_PER_MINUTE((long) 77L),
  MILES_PER_HOUR((long) 78L),
  CUBIC_FEET((long) 79L),
  CUBIC_METERS((long) 80L),
  IMPERIAL_GALLONS((long) 81L),
  MILLILITERS((long) 197L),
  LITERS((long) 82L),
  US_GALLONS((long) 83L),
  CUBIC_FEET_PER_SECOND((long) 142L),
  CUBIC_FEET_PER_MINUTE((long) 84L),
  MILLION_STANDARD_CUBIC_FEET_PER_MINUTE((long) 254L),
  CUBIC_FEET_PER_HOUR((long) 191L),
  CUBIC_FEET_PER_DAY((long) 248L),
  STANDARD_CUBIC_FEET_PER_DAY((long) 47808L),
  MILLION_STANDARD_CUBIC_FEET_PER_DAY((long) 47809L),
  THOUSAND_CUBIC_FEET_PER_DAY((long) 47810L),
  THOUSAND_STANDARD_CUBIC_FEET_PER_DAY((long) 47811L),
  POUNDS_MASS_PER_DAY((long) 47812L),
  CUBIC_METERS_PER_SECOND((long) 85L),
  CUBIC_METERS_PER_MINUTE((long) 165L),
  CUBIC_METERS_PER_HOUR((long) 135L),
  CUBIC_METERS_PER_DAY((long) 249L),
  IMPERIAL_GALLONS_PER_MINUTE((long) 86L),
  MILLILITERS_PER_SECOND((long) 198L),
  LITERS_PER_SECOND((long) 87L),
  LITERS_PER_MINUTE((long) 88L),
  LITERS_PER_HOUR((long) 136L),
  US_GALLONS_PER_MINUTE((long) 89L),
  US_GALLONS_PER_HOUR((long) 192L),
  DEGREES_ANGULAR((long) 90L),
  DEGREES_CELSIUS_PER_HOUR((long) 91L),
  DEGREES_CELSIUS_PER_MINUTE((long) 92L),
  DEGREES_FAHRENHEIT_PER_HOUR((long) 93L),
  DEGREES_FAHRENHEIT_PER_MINUTE((long) 94L),
  JOULE_SECONDS((long) 183L),
  KILOGRAMS_PER_CUBIC_METER((long) 186L),
  KILOWATT_HOURS_PER_SQUARE_METER((long) 137L),
  KILOWATT_HOURS_PER_SQUARE_FOOT((long) 138L),
  WATT_HOURS_PER_CUBIC_METER((long) 250L),
  JOULES_PER_CUBIC_METER((long) 251L),
  MEGAJOULES_PER_SQUARE_METER((long) 139L),
  MEGAJOULES_PER_SQUARE_FOOT((long) 140L),
  MOLE_PERCENT((long) 252L),
  NO_UNITS((long) 95L),
  NEWTON_SECONDS((long) 187L),
  NEWTONS_PER_METER((long) 188L),
  PARTS_PER_MILLION((long) 96L),
  PARTS_PER_BILLION((long) 97L),
  PASCAL_SECONDS((long) 253L),
  PERCENT((long) 98L),
  PERCENT_OBSCURATION_PER_FOOT((long) 143L),
  PERCENT_OBSCURATION_PER_METER((long) 144L),
  PERCENT_PER_SECOND((long) 99L),
  PER_MINUTE((long) 100L),
  PER_SECOND((long) 101L),
  PSI_PER_DEGREE_FAHRENHEIT((long) 102L),
  RADIANS((long) 103L),
  RADIANS_PER_SECOND((long) 184L),
  REVOLUTIONS_PER_MINUTE((long) 104L),
  SQUARE_METERS_PER_NEWTON((long) 185L),
  WATTS_PER_METER_PER_DEGREE_KELVIN((long) 189L),
  WATTS_PER_SQUARE_METER_DEGREE_KELVIN((long) 141L),
  PER_MILLE((long) 207L),
  GRAMS_PER_GRAM((long) 208L),
  KILOGRAMS_PER_KILOGRAM((long) 209L),
  GRAMS_PER_KILOGRAM((long) 210L),
  MILLIGRAMS_PER_GRAM((long) 211L),
  MILLIGRAMS_PER_KILOGRAM((long) 212L),
  GRAMS_PER_MILLILITER((long) 213L),
  GRAMS_PER_LITER((long) 214L),
  MILLIGRAMS_PER_LITER((long) 215L),
  MICROGRAMS_PER_LITER((long) 216L),
  GRAMS_PER_CUBIC_METER((long) 217L),
  MILLIGRAMS_PER_CUBIC_METER((long) 218L),
  MICROGRAMS_PER_CUBIC_METER((long) 219L),
  NANOGRAMS_PER_CUBIC_METER((long) 220L),
  GRAMS_PER_CUBIC_CENTIMETER((long) 221L),
  BECQUERELS((long) 222L),
  KILOBECQUERELS((long) 223L),
  MEGABECQUERELS((long) 224L),
  GRAY((long) 225L),
  MILLIGRAY((long) 226L),
  MICROGRAY((long) 227L),
  SIEVERTS((long) 228L),
  MILLISIEVERTS((long) 229L),
  MICROSIEVERTS((long) 230L),
  MICROSIEVERTS_PER_HOUR((long) 231L),
  MILLIREMS((long) 47814L),
  MILLIREMS_PER_HOUR((long) 47815L),
  DECIBELS_A((long) 232L),
  NEPHELOMETRIC_TURBIDITY_UNIT((long) 233L),
  P_H((long) 234L),
  GRAMS_PER_SQUARE_METER((long) 235L),
  MINUTES_PER_DEGREE_KELVIN((long) 236L),
  VENDOR_PROPRIETARY_VALUE((long) 0XFFL);
  private static final Map<Long, BACnetEngineeringUnits> map;

  static {
    map = new HashMap<>();
    for (BACnetEngineeringUnits value : BACnetEngineeringUnits.values()) {
      map.put((long) value.getValue(), value);
    }
  }

  private final long value;

  BACnetEngineeringUnits(long value) {
    this.value = value;
  }

  public long getValue() {
    return value;
  }

  public static BACnetEngineeringUnits enumForValue(long value) {
    return map.get(value);
  }

  public static Boolean isDefined(long value) {
    return map.containsKey(value);
  }
}
