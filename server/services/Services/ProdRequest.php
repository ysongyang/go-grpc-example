<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prod.proto

namespace Services;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>services.ProdRequest</code>
 */
class ProdRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 prodId = 1;</code>
     */
    protected $prodId = 0;
    /**
     * Generated from protobuf field <code>.services.ProAreas proArea = 2;</code>
     */
    protected $proArea = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $prodId
     *     @type int $proArea
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Prod::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 prodId = 1;</code>
     * @return int
     */
    public function getProdId()
    {
        return $this->prodId;
    }

    /**
     * Generated from protobuf field <code>int32 prodId = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setProdId($var)
    {
        GPBUtil::checkInt32($var);
        $this->prodId = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.services.ProAreas proArea = 2;</code>
     * @return int
     */
    public function getProArea()
    {
        return $this->proArea;
    }

    /**
     * Generated from protobuf field <code>.services.ProAreas proArea = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setProArea($var)
    {
        GPBUtil::checkEnum($var, \Services\ProAreas::class);
        $this->proArea = $var;

        return $this;
    }

}
