// source: landcryptid.proto
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

goog.exportSymbol('proto.landcryptid.CryptidInfo', null, global);
goog.exportSymbol('proto.landcryptid.CryptidParam', null, global);
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
proto.landcryptid.CryptidInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.landcryptid.CryptidInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.landcryptid.CryptidInfo.displayName = 'proto.landcryptid.CryptidInfo';
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
proto.landcryptid.CryptidParam = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.landcryptid.CryptidParam, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.landcryptid.CryptidParam.displayName = 'proto.landcryptid.CryptidParam';
}



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
proto.landcryptid.CryptidInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.landcryptid.CryptidInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.landcryptid.CryptidInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.landcryptid.CryptidInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    landType: jspb.Message.getFieldWithDefault(msg, 1, 0),
    energy: jspb.Message.getFieldWithDefault(msg, 2, 0),
    lv: jspb.Message.getFieldWithDefault(msg, 3, 0),
    season: jspb.Message.getFieldWithDefault(msg, 4, 0),
    laboCe: jspb.Message.getFieldWithDefault(msg, 5, 0),
    hp: (f = msg.getHp()) && proto.landcryptid.CryptidParam.toObject(includeInstance, f),
    phy: (f = msg.getPhy()) && proto.landcryptid.CryptidParam.toObject(includeInstance, f),
    pb_int: (f = msg.getInt()) && proto.landcryptid.CryptidParam.toObject(includeInstance, f),
    agi: (f = msg.getAgi()) && proto.landcryptid.CryptidParam.toObject(includeInstance, f),
    recovery: (f = msg.getRecovery()) && proto.landcryptid.CryptidParam.toObject(includeInstance, f)
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
 * @return {!proto.landcryptid.CryptidInfo}
 */
proto.landcryptid.CryptidInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.landcryptid.CryptidInfo;
  return proto.landcryptid.CryptidInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.landcryptid.CryptidInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.landcryptid.CryptidInfo}
 */
proto.landcryptid.CryptidInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setLandType(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setEnergy(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLv(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSeason(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLaboCe(value);
      break;
    case 6:
      var value = new proto.landcryptid.CryptidParam;
      reader.readMessage(value,proto.landcryptid.CryptidParam.deserializeBinaryFromReader);
      msg.setHp(value);
      break;
    case 7:
      var value = new proto.landcryptid.CryptidParam;
      reader.readMessage(value,proto.landcryptid.CryptidParam.deserializeBinaryFromReader);
      msg.setPhy(value);
      break;
    case 8:
      var value = new proto.landcryptid.CryptidParam;
      reader.readMessage(value,proto.landcryptid.CryptidParam.deserializeBinaryFromReader);
      msg.setInt(value);
      break;
    case 9:
      var value = new proto.landcryptid.CryptidParam;
      reader.readMessage(value,proto.landcryptid.CryptidParam.deserializeBinaryFromReader);
      msg.setAgi(value);
      break;
    case 10:
      var value = new proto.landcryptid.CryptidParam;
      reader.readMessage(value,proto.landcryptid.CryptidParam.deserializeBinaryFromReader);
      msg.setRecovery(value);
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
proto.landcryptid.CryptidInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.landcryptid.CryptidInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.landcryptid.CryptidInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.landcryptid.CryptidInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLandType();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getEnergy();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getLv();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getSeason();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getLaboCe();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getHp();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.landcryptid.CryptidParam.serializeBinaryToWriter
    );
  }
  f = message.getPhy();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      proto.landcryptid.CryptidParam.serializeBinaryToWriter
    );
  }
  f = message.getInt();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      proto.landcryptid.CryptidParam.serializeBinaryToWriter
    );
  }
  f = message.getAgi();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      proto.landcryptid.CryptidParam.serializeBinaryToWriter
    );
  }
  f = message.getRecovery();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      proto.landcryptid.CryptidParam.serializeBinaryToWriter
    );
  }
};


/**
 * optional uint32 land_type = 1;
 * @return {number}
 */
proto.landcryptid.CryptidInfo.prototype.getLandType = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.setLandType = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 energy = 2;
 * @return {number}
 */
proto.landcryptid.CryptidInfo.prototype.getEnergy = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.setEnergy = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 lv = 3;
 * @return {number}
 */
proto.landcryptid.CryptidInfo.prototype.getLv = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.setLv = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 season = 4;
 * @return {number}
 */
proto.landcryptid.CryptidInfo.prototype.getSeason = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.setSeason = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int32 labo_ce = 5;
 * @return {number}
 */
proto.landcryptid.CryptidInfo.prototype.getLaboCe = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.setLaboCe = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional CryptidParam hp = 6;
 * @return {?proto.landcryptid.CryptidParam}
 */
proto.landcryptid.CryptidInfo.prototype.getHp = function() {
  return /** @type{?proto.landcryptid.CryptidParam} */ (
    jspb.Message.getWrapperField(this, proto.landcryptid.CryptidParam, 6));
};


/**
 * @param {?proto.landcryptid.CryptidParam|undefined} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
*/
proto.landcryptid.CryptidInfo.prototype.setHp = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.clearHp = function() {
  return this.setHp(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.landcryptid.CryptidInfo.prototype.hasHp = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional CryptidParam phy = 7;
 * @return {?proto.landcryptid.CryptidParam}
 */
proto.landcryptid.CryptidInfo.prototype.getPhy = function() {
  return /** @type{?proto.landcryptid.CryptidParam} */ (
    jspb.Message.getWrapperField(this, proto.landcryptid.CryptidParam, 7));
};


/**
 * @param {?proto.landcryptid.CryptidParam|undefined} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
*/
proto.landcryptid.CryptidInfo.prototype.setPhy = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.clearPhy = function() {
  return this.setPhy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.landcryptid.CryptidInfo.prototype.hasPhy = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional CryptidParam int = 8;
 * @return {?proto.landcryptid.CryptidParam}
 */
proto.landcryptid.CryptidInfo.prototype.getInt = function() {
  return /** @type{?proto.landcryptid.CryptidParam} */ (
    jspb.Message.getWrapperField(this, proto.landcryptid.CryptidParam, 8));
};


/**
 * @param {?proto.landcryptid.CryptidParam|undefined} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
*/
proto.landcryptid.CryptidInfo.prototype.setInt = function(value) {
  return jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.clearInt = function() {
  return this.setInt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.landcryptid.CryptidInfo.prototype.hasInt = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional CryptidParam agi = 9;
 * @return {?proto.landcryptid.CryptidParam}
 */
proto.landcryptid.CryptidInfo.prototype.getAgi = function() {
  return /** @type{?proto.landcryptid.CryptidParam} */ (
    jspb.Message.getWrapperField(this, proto.landcryptid.CryptidParam, 9));
};


/**
 * @param {?proto.landcryptid.CryptidParam|undefined} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
*/
proto.landcryptid.CryptidInfo.prototype.setAgi = function(value) {
  return jspb.Message.setWrapperField(this, 9, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.clearAgi = function() {
  return this.setAgi(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.landcryptid.CryptidInfo.prototype.hasAgi = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional CryptidParam recovery = 10;
 * @return {?proto.landcryptid.CryptidParam}
 */
proto.landcryptid.CryptidInfo.prototype.getRecovery = function() {
  return /** @type{?proto.landcryptid.CryptidParam} */ (
    jspb.Message.getWrapperField(this, proto.landcryptid.CryptidParam, 10));
};


/**
 * @param {?proto.landcryptid.CryptidParam|undefined} value
 * @return {!proto.landcryptid.CryptidInfo} returns this
*/
proto.landcryptid.CryptidInfo.prototype.setRecovery = function(value) {
  return jspb.Message.setWrapperField(this, 10, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.landcryptid.CryptidInfo} returns this
 */
proto.landcryptid.CryptidInfo.prototype.clearRecovery = function() {
  return this.setRecovery(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.landcryptid.CryptidInfo.prototype.hasRecovery = function() {
  return jspb.Message.getField(this, 10) != null;
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
proto.landcryptid.CryptidParam.prototype.toObject = function(opt_includeInstance) {
  return proto.landcryptid.CryptidParam.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.landcryptid.CryptidParam} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.landcryptid.CryptidParam.toObject = function(includeInstance, msg) {
  var f, obj = {
    divideRate100: jspb.Message.getFieldWithDefault(msg, 1, 0),
    ce: jspb.Message.getFieldWithDefault(msg, 2, 0),
    lv: jspb.Message.getFieldWithDefault(msg, 3, 0),
    buffRate1000: jspb.Message.getFieldWithDefault(msg, 4, 0)
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
 * @return {!proto.landcryptid.CryptidParam}
 */
proto.landcryptid.CryptidParam.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.landcryptid.CryptidParam;
  return proto.landcryptid.CryptidParam.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.landcryptid.CryptidParam} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.landcryptid.CryptidParam}
 */
proto.landcryptid.CryptidParam.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setDivideRate100(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCe(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLv(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setBuffRate1000(value);
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
proto.landcryptid.CryptidParam.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.landcryptid.CryptidParam.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.landcryptid.CryptidParam} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.landcryptid.CryptidParam.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDivideRate100();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getCe();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getLv();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getBuffRate1000();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
};


/**
 * optional int32 divide_rate100 = 1;
 * @return {number}
 */
proto.landcryptid.CryptidParam.prototype.getDivideRate100 = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidParam} returns this
 */
proto.landcryptid.CryptidParam.prototype.setDivideRate100 = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 ce = 2;
 * @return {number}
 */
proto.landcryptid.CryptidParam.prototype.getCe = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidParam} returns this
 */
proto.landcryptid.CryptidParam.prototype.setCe = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 lv = 3;
 * @return {number}
 */
proto.landcryptid.CryptidParam.prototype.getLv = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidParam} returns this
 */
proto.landcryptid.CryptidParam.prototype.setLv = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 buff_rate1000 = 4;
 * @return {number}
 */
proto.landcryptid.CryptidParam.prototype.getBuffRate1000 = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.landcryptid.CryptidParam} returns this
 */
proto.landcryptid.CryptidParam.prototype.setBuffRate1000 = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


goog.object.extend(exports, proto.landcryptid);
