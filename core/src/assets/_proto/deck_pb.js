// source: deck.proto
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

goog.exportSymbol('proto.deck.Deck', null, global);
goog.exportSymbol('proto.deck.DeckUnit', null, global);
goog.exportSymbol('proto.deck.Decks', null, global);
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
proto.deck.Decks = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.deck.Decks, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.deck.Decks.displayName = 'proto.deck.Decks';
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
proto.deck.Deck = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.deck.Deck.repeatedFields_, null);
};
goog.inherits(proto.deck.Deck, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.deck.Deck.displayName = 'proto.deck.Deck';
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
proto.deck.DeckUnit = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.deck.DeckUnit.repeatedFields_, null);
};
goog.inherits(proto.deck.DeckUnit, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.deck.DeckUnit.displayName = 'proto.deck.DeckUnit';
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
proto.deck.Decks.prototype.toObject = function(opt_includeInstance) {
  return proto.deck.Decks.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.deck.Decks} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.deck.Decks.toObject = function(includeInstance, msg) {
  var f, obj = {
    simulateOffenseDeck: (f = msg.getSimulateOffenseDeck()) && proto.deck.Deck.toObject(includeInstance, f),
    simulateDefenseDeck: (f = msg.getSimulateDefenseDeck()) && proto.deck.Deck.toObject(includeInstance, f)
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
 * @return {!proto.deck.Decks}
 */
proto.deck.Decks.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.deck.Decks;
  return proto.deck.Decks.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.deck.Decks} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.deck.Decks}
 */
proto.deck.Decks.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.deck.Deck;
      reader.readMessage(value,proto.deck.Deck.deserializeBinaryFromReader);
      msg.setSimulateOffenseDeck(value);
      break;
    case 2:
      var value = new proto.deck.Deck;
      reader.readMessage(value,proto.deck.Deck.deserializeBinaryFromReader);
      msg.setSimulateDefenseDeck(value);
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
proto.deck.Decks.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.deck.Decks.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.deck.Decks} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.deck.Decks.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSimulateOffenseDeck();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.deck.Deck.serializeBinaryToWriter
    );
  }
  f = message.getSimulateDefenseDeck();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.deck.Deck.serializeBinaryToWriter
    );
  }
};


/**
 * optional Deck simulate_offense_deck = 1;
 * @return {?proto.deck.Deck}
 */
proto.deck.Decks.prototype.getSimulateOffenseDeck = function() {
  return /** @type{?proto.deck.Deck} */ (
    jspb.Message.getWrapperField(this, proto.deck.Deck, 1));
};


/**
 * @param {?proto.deck.Deck|undefined} value
 * @return {!proto.deck.Decks} returns this
*/
proto.deck.Decks.prototype.setSimulateOffenseDeck = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.deck.Decks} returns this
 */
proto.deck.Decks.prototype.clearSimulateOffenseDeck = function() {
  return this.setSimulateOffenseDeck(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.deck.Decks.prototype.hasSimulateOffenseDeck = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Deck simulate_defense_deck = 2;
 * @return {?proto.deck.Deck}
 */
proto.deck.Decks.prototype.getSimulateDefenseDeck = function() {
  return /** @type{?proto.deck.Deck} */ (
    jspb.Message.getWrapperField(this, proto.deck.Deck, 2));
};


/**
 * @param {?proto.deck.Deck|undefined} value
 * @return {!proto.deck.Decks} returns this
*/
proto.deck.Decks.prototype.setSimulateDefenseDeck = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.deck.Decks} returns this
 */
proto.deck.Decks.prototype.clearSimulateDefenseDeck = function() {
  return this.setSimulateDefenseDeck(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.deck.Decks.prototype.hasSimulateDefenseDeck = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.deck.Deck.repeatedFields_ = [1];



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
proto.deck.Deck.prototype.toObject = function(opt_includeInstance) {
  return proto.deck.Deck.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.deck.Deck} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.deck.Deck.toObject = function(includeInstance, msg) {
  var f, obj = {
    unitsList: jspb.Message.toObjectList(msg.getUnitsList(),
    proto.deck.DeckUnit.toObject, includeInstance)
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
 * @return {!proto.deck.Deck}
 */
proto.deck.Deck.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.deck.Deck;
  return proto.deck.Deck.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.deck.Deck} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.deck.Deck}
 */
proto.deck.Deck.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.deck.DeckUnit;
      reader.readMessage(value,proto.deck.DeckUnit.deserializeBinaryFromReader);
      msg.addUnits(value);
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
proto.deck.Deck.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.deck.Deck.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.deck.Deck} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.deck.Deck.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUnitsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.deck.DeckUnit.serializeBinaryToWriter
    );
  }
};


/**
 * repeated DeckUnit units = 1;
 * @return {!Array<!proto.deck.DeckUnit>}
 */
proto.deck.Deck.prototype.getUnitsList = function() {
  return /** @type{!Array<!proto.deck.DeckUnit>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.deck.DeckUnit, 1));
};


/**
 * @param {!Array<!proto.deck.DeckUnit>} value
 * @return {!proto.deck.Deck} returns this
*/
proto.deck.Deck.prototype.setUnitsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.deck.DeckUnit=} opt_value
 * @param {number=} opt_index
 * @return {!proto.deck.DeckUnit}
 */
proto.deck.Deck.prototype.addUnits = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.deck.DeckUnit, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.deck.Deck} returns this
 */
proto.deck.Deck.prototype.clearUnitsList = function() {
  return this.setUnitsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.deck.DeckUnit.repeatedFields_ = [2,3];



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
proto.deck.DeckUnit.prototype.toObject = function(opt_includeInstance) {
  return proto.deck.DeckUnit.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.deck.DeckUnit} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.deck.DeckUnit.toObject = function(includeInstance, msg) {
  var f, obj = {
    heroId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    extensionIdsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    skillOrdersList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    heroActiveIndex: jspb.Message.getFieldWithDefault(msg, 4, 0)
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
 * @return {!proto.deck.DeckUnit}
 */
proto.deck.DeckUnit.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.deck.DeckUnit;
  return proto.deck.DeckUnit.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.deck.DeckUnit} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.deck.DeckUnit}
 */
proto.deck.DeckUnit.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setHeroId(value);
      break;
    case 2:
      var value = /** @type {!Array<number>} */ (reader.readPackedUint32());
      msg.setExtensionIdsList(value);
      break;
    case 3:
      var value = /** @type {!Array<number>} */ (reader.readPackedInt32());
      msg.setSkillOrdersList(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setHeroActiveIndex(value);
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
proto.deck.DeckUnit.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.deck.DeckUnit.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.deck.DeckUnit} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.deck.DeckUnit.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHeroId();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getExtensionIdsList();
  if (f.length > 0) {
    writer.writePackedUint32(
      2,
      f
    );
  }
  f = message.getSkillOrdersList();
  if (f.length > 0) {
    writer.writePackedInt32(
      3,
      f
    );
  }
  f = message.getHeroActiveIndex();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
};


/**
 * optional uint32 hero_id = 1;
 * @return {number}
 */
proto.deck.DeckUnit.prototype.getHeroId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.setHeroId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated uint32 extension_ids = 2;
 * @return {!Array<number>}
 */
proto.deck.DeckUnit.prototype.getExtensionIdsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.setExtensionIdsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.addExtensionIds = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.clearExtensionIdsList = function() {
  return this.setExtensionIdsList([]);
};


/**
 * repeated int32 skill_orders = 3;
 * @return {!Array<number>}
 */
proto.deck.DeckUnit.prototype.getSkillOrdersList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.setSkillOrdersList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.addSkillOrders = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.clearSkillOrdersList = function() {
  return this.setSkillOrdersList([]);
};


/**
 * optional int32 hero_active_index = 4;
 * @return {number}
 */
proto.deck.DeckUnit.prototype.getHeroActiveIndex = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.deck.DeckUnit} returns this
 */
proto.deck.DeckUnit.prototype.setHeroActiveIndex = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


goog.object.extend(exports, proto.deck);
