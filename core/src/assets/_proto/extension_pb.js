// source: extension.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.extension.Empty', null, global);
goog.exportSymbol('proto.extension.ExtensionData', null, global);
goog.exportSymbol('proto.extension.ExtensionData.Lock', null, global);
goog.exportSymbol('proto.extension.Param', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.extension.ExtensionData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.extension.ExtensionData.repeatedFields_, null);
};
goog.inherits(proto.extension.ExtensionData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.extension.ExtensionData.displayName = 'proto.extension.ExtensionData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.extension.Param = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.extension.Param, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.extension.Param.displayName = 'proto.extension.Param';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.extension.Empty = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.extension.Empty, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.extension.Empty.displayName = 'proto.extension.Empty';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.extension.ExtensionData.repeatedFields_ = [7];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.extension.ExtensionData.prototype.toObject = function(opt_includeInstance) {
  return proto.extension.ExtensionData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.extension.ExtensionData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.extension.ExtensionData.toObject = function(includeInstance, msg) {
  var f, obj = {
    extensionId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    extensionType: jspb.Message.getFieldWithDefault(msg, 2, 0),
    param: (f = msg.getParam()) && proto.extension.Param.toObject(includeInstance, f),
    trend: (f = msg.getTrend()) && proto.extension.Param.toObject(includeInstance, f),
    active: jspb.Message.getFieldWithDefault(msg, 5, 0),
    lock: jspb.Message.getFieldWithDefault(msg, 6, 0),
    titleList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
    landType: jspb.Message.getFieldWithDefault(msg, 8, 0),
    lockUntil: jspb.Message.getFieldWithDefault(msg, 9, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.extension.ExtensionData}
 */
proto.extension.ExtensionData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.extension.ExtensionData;
  return proto.extension.ExtensionData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.extension.ExtensionData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.extension.ExtensionData}
 */
proto.extension.ExtensionData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setExtensionId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setExtensionType(value);
      break;
    case 3:
      var value = new proto.extension.Param;
      reader.readMessage(value,proto.extension.Param.deserializeBinaryFromReader);
      msg.setParam(value);
      break;
    case 4:
      var value = new proto.extension.Param;
      reader.readMessage(value,proto.extension.Param.deserializeBinaryFromReader);
      msg.setTrend(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setActive(value);
      break;
    case 6:
      var value = /** @type {!proto.extension.ExtensionData.Lock} */ (reader.readEnum());
      msg.setLock(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.addTitle(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setLandType(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLockUntil(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.extension.ExtensionData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.extension.ExtensionData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.extension.ExtensionData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.extension.ExtensionData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getExtensionId();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getExtensionType();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
  f = message.getParam();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.extension.Param.serializeBinaryToWriter
    );
  }
  f = message.getTrend();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.extension.Param.serializeBinaryToWriter
    );
  }
  f = message.getActive();
  if (f !== 0) {
    writer.writeUint32(
      5,
      f
    );
  }
  f = message.getLock();
  if (f !== 0.0) {
    writer.writeEnum(
      6,
      f
    );
  }
  f = message.getTitleList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      7,
      f
    );
  }
  f = message.getLandType();
  if (f !== 0) {
    writer.writeUint32(
      8,
      f
    );
  }
  f = message.getLockUntil();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.extension.ExtensionData.Lock = {
  NONE: 0,
  ADMIN: 1,
  DECK: 2,
  TRADE: 3,
  CRAFT: 4,
  OUTSOURCING: 5
};

/**
 * optional uint32 extension_id = 1;
 * @return {number}
 */
proto.extension.ExtensionData.prototype.getExtensionId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.setExtensionId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional uint32 extension_type = 2;
 * @return {number}
 */
proto.extension.ExtensionData.prototype.getExtensionType = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.setExtensionType = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional Param param = 3;
 * @return {?proto.extension.Param}
 */
proto.extension.ExtensionData.prototype.getParam = function() {
  return /** @type{?proto.extension.Param} */ (
    jspb.Message.getWrapperField(this, proto.extension.Param, 3));
};


/**
 * @param {?proto.extension.Param|undefined} value
 * @return {!proto.extension.ExtensionData} returns this
*/
proto.extension.ExtensionData.prototype.setParam = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.clearParam = function() {
  return this.setParam(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.extension.ExtensionData.prototype.hasParam = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Param trend = 4;
 * @return {?proto.extension.Param}
 */
proto.extension.ExtensionData.prototype.getTrend = function() {
  return /** @type{?proto.extension.Param} */ (
    jspb.Message.getWrapperField(this, proto.extension.Param, 4));
};


/**
 * @param {?proto.extension.Param|undefined} value
 * @return {!proto.extension.ExtensionData} returns this
*/
proto.extension.ExtensionData.prototype.setTrend = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.clearTrend = function() {
  return this.setTrend(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.extension.ExtensionData.prototype.hasTrend = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional uint32 active = 5;
 * @return {number}
 */
proto.extension.ExtensionData.prototype.getActive = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.setActive = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional Lock lock = 6;
 * @return {!proto.extension.ExtensionData.Lock}
 */
proto.extension.ExtensionData.prototype.getLock = function() {
  return /** @type {!proto.extension.ExtensionData.Lock} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {!proto.extension.ExtensionData.Lock} value
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.setLock = function(value) {
  return jspb.Message.setProto3EnumField(this, 6, value);
};


/**
 * repeated string title = 7;
 * @return {!Array<string>}
 */
proto.extension.ExtensionData.prototype.getTitleList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.setTitleList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.addTitle = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.clearTitleList = function() {
  return this.setTitleList([]);
};


/**
 * optional uint32 land_type = 8;
 * @return {number}
 */
proto.extension.ExtensionData.prototype.getLandType = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.setLandType = function(value) {
  return jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int64 lock_until = 9;
 * @return {number}
 */
proto.extension.ExtensionData.prototype.getLockUntil = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.ExtensionData} returns this
 */
proto.extension.ExtensionData.prototype.setLockUntil = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.extension.Param.prototype.toObject = function(opt_includeInstance) {
  return proto.extension.Param.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.extension.Param} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.extension.Param.toObject = function(includeInstance, msg) {
  var f, obj = {
    hp: jspb.Message.getFieldWithDefault(msg, 1, 0),
    phy: jspb.Message.getFieldWithDefault(msg, 2, 0),
    pb_int: jspb.Message.getFieldWithDefault(msg, 3, 0),
    agi: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.extension.Param}
 */
proto.extension.Param.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.extension.Param;
  return proto.extension.Param.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.extension.Param} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.extension.Param}
 */
proto.extension.Param.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setHp(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPhy(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setInt(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAgi(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.extension.Param.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.extension.Param.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.extension.Param} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.extension.Param.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHp();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getPhy();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getInt();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getAgi();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
};


/**
 * optional int32 hp = 1;
 * @return {number}
 */
proto.extension.Param.prototype.getHp = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.Param} returns this
 */
proto.extension.Param.prototype.setHp = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 phy = 2;
 * @return {number}
 */
proto.extension.Param.prototype.getPhy = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.Param} returns this
 */
proto.extension.Param.prototype.setPhy = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 int = 3;
 * @return {number}
 */
proto.extension.Param.prototype.getInt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.Param} returns this
 */
proto.extension.Param.prototype.setInt = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 agi = 4;
 * @return {number}
 */
proto.extension.Param.prototype.getAgi = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.extension.Param} returns this
 */
proto.extension.Param.prototype.setAgi = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.extension.Empty.prototype.toObject = function(opt_includeInstance) {
  return proto.extension.Empty.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.extension.Empty} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.extension.Empty.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.extension.Empty}
 */
proto.extension.Empty.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.extension.Empty;
  return proto.extension.Empty.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.extension.Empty} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.extension.Empty}
 */
proto.extension.Empty.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.extension.Empty.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.extension.Empty.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.extension.Empty} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.extension.Empty.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


goog.object.extend(exports, proto.extension);
