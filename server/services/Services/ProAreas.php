<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prod.proto

namespace Services;

use UnexpectedValueException;

/**
 * Protobuf type <code>services.ProAreas</code>
 */
class ProAreas
{
    /**
     *必须是0 表示默认值
     *
     * Generated from protobuf enum <code>A = 0;</code>
     */
    const A = 0;
    /**
     * Generated from protobuf enum <code>B = 1;</code>
     */
    const B = 1;
    /**
     * Generated from protobuf enum <code>C = 2;</code>
     */
    const C = 2;

    private static $valueToName = [
        self::A => 'A',
        self::B => 'B',
        self::C => 'C',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}
