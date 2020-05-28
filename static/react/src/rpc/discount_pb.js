// source: discount.proto
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

goog.exportSymbol('proto.dss.DiscountAddReq', null, global);
goog.exportSymbol('proto.dss.DiscountAddRes', null, global);
goog.exportSymbol('proto.dss.DiscountAmount', null, global);
goog.exportSymbol('proto.dss.DiscountAmount.AmountCase', null, global);
goog.exportSymbol('proto.dss.DiscountBundle', null, global);
goog.exportSymbol('proto.dss.DiscountGetReq', null, global);
goog.exportSymbol('proto.dss.DiscountGetRes', null, global);
goog.exportSymbol('proto.dss.SingleDiscount', null, global);
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
proto.dss.DiscountAmount = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.dss.DiscountAmount.oneofGroups_);
};
goog.inherits(proto.dss.DiscountAmount, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dss.DiscountAmount.displayName = 'proto.dss.DiscountAmount';
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
proto.dss.SingleDiscount = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dss.SingleDiscount, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dss.SingleDiscount.displayName = 'proto.dss.SingleDiscount';
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
proto.dss.DiscountBundle = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dss.DiscountBundle.repeatedFields_, null);
};
goog.inherits(proto.dss.DiscountBundle, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dss.DiscountBundle.displayName = 'proto.dss.DiscountBundle';
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
proto.dss.DiscountAddReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dss.DiscountAddReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dss.DiscountAddReq.displayName = 'proto.dss.DiscountAddReq';
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
proto.dss.DiscountAddRes = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dss.DiscountAddRes, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dss.DiscountAddRes.displayName = 'proto.dss.DiscountAddRes';
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
proto.dss.DiscountGetReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dss.DiscountGetReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dss.DiscountGetReq.displayName = 'proto.dss.DiscountGetReq';
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
proto.dss.DiscountGetRes = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dss.DiscountGetRes, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dss.DiscountGetRes.displayName = 'proto.dss.DiscountGetRes';
}

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.dss.DiscountAmount.oneofGroups_ = [[8,9]];

/**
 * @enum {number}
 */
proto.dss.DiscountAmount.AmountCase = {
  AMOUNT_NOT_SET: 0,
  DOLLAR: 8,
  PERCENT: 9
};

/**
 * @return {proto.dss.DiscountAmount.AmountCase}
 */
proto.dss.DiscountAmount.prototype.getAmountCase = function() {
  return /** @type {proto.dss.DiscountAmount.AmountCase} */(jspb.Message.computeOneofCase(this, proto.dss.DiscountAmount.oneofGroups_[0]));
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
proto.dss.DiscountAmount.prototype.toObject = function(opt_includeInstance) {
  return proto.dss.DiscountAmount.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dss.DiscountAmount} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountAmount.toObject = function(includeInstance, msg) {
  var f, obj = {
    dollar: jspb.Message.getFieldWithDefault(msg, 8, 0),
    percent: jspb.Message.getFieldWithDefault(msg, 9, "")
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
 * @return {!proto.dss.DiscountAmount}
 */
proto.dss.DiscountAmount.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dss.DiscountAmount;
  return proto.dss.DiscountAmount.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dss.DiscountAmount} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dss.DiscountAmount}
 */
proto.dss.DiscountAmount.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 8:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDollar(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setPercent(value);
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
proto.dss.DiscountAmount.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dss.DiscountAmount.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dss.DiscountAmount} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountAmount.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeInt64(
      8,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 9));
  if (f != null) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * optional int64 dollar = 8;
 * @return {number}
 */
proto.dss.DiscountAmount.prototype.getDollar = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/**
 * @param {number} value
 * @return {!proto.dss.DiscountAmount} returns this
 */
proto.dss.DiscountAmount.prototype.setDollar = function(value) {
  return jspb.Message.setOneofField(this, 8, proto.dss.DiscountAmount.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.dss.DiscountAmount} returns this
 */
proto.dss.DiscountAmount.prototype.clearDollar = function() {
  return jspb.Message.setOneofField(this, 8, proto.dss.DiscountAmount.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dss.DiscountAmount.prototype.hasDollar = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional string percent = 9;
 * @return {string}
 */
proto.dss.DiscountAmount.prototype.getPercent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.dss.DiscountAmount} returns this
 */
proto.dss.DiscountAmount.prototype.setPercent = function(value) {
  return jspb.Message.setOneofField(this, 9, proto.dss.DiscountAmount.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.dss.DiscountAmount} returns this
 */
proto.dss.DiscountAmount.prototype.clearPercent = function() {
  return jspb.Message.setOneofField(this, 9, proto.dss.DiscountAmount.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dss.DiscountAmount.prototype.hasPercent = function() {
  return jspb.Message.getField(this, 9) != null;
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
proto.dss.SingleDiscount.prototype.toObject = function(opt_includeInstance) {
  return proto.dss.SingleDiscount.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dss.SingleDiscount} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.SingleDiscount.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    amount: (f = msg.getAmount()) && proto.dss.DiscountAmount.toObject(includeInstance, f)
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
 * @return {!proto.dss.SingleDiscount}
 */
proto.dss.SingleDiscount.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dss.SingleDiscount;
  return proto.dss.SingleDiscount.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dss.SingleDiscount} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dss.SingleDiscount}
 */
proto.dss.SingleDiscount.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = new proto.dss.DiscountAmount;
      reader.readMessage(value,proto.dss.DiscountAmount.deserializeBinaryFromReader);
      msg.setAmount(value);
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
proto.dss.SingleDiscount.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dss.SingleDiscount.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dss.SingleDiscount} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.SingleDiscount.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAmount();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.dss.DiscountAmount.serializeBinaryToWriter
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.dss.SingleDiscount.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.dss.SingleDiscount} returns this
 */
proto.dss.SingleDiscount.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional DiscountAmount amount = 2;
 * @return {?proto.dss.DiscountAmount}
 */
proto.dss.SingleDiscount.prototype.getAmount = function() {
  return /** @type{?proto.dss.DiscountAmount} */ (
    jspb.Message.getWrapperField(this, proto.dss.DiscountAmount, 2));
};


/**
 * @param {?proto.dss.DiscountAmount|undefined} value
 * @return {!proto.dss.SingleDiscount} returns this
*/
proto.dss.SingleDiscount.prototype.setAmount = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.dss.SingleDiscount} returns this
 */
proto.dss.SingleDiscount.prototype.clearAmount = function() {
  return this.setAmount(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dss.SingleDiscount.prototype.hasAmount = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dss.DiscountBundle.repeatedFields_ = [2];



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
proto.dss.DiscountBundle.prototype.toObject = function(opt_includeInstance) {
  return proto.dss.DiscountBundle.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dss.DiscountBundle} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountBundle.toObject = function(includeInstance, msg) {
  var f, obj = {
    code: jspb.Message.getFieldWithDefault(msg, 1, ""),
    discountsList: jspb.Message.toObjectList(msg.getDiscountsList(),
    proto.dss.SingleDiscount.toObject, includeInstance)
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
 * @return {!proto.dss.DiscountBundle}
 */
proto.dss.DiscountBundle.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dss.DiscountBundle;
  return proto.dss.DiscountBundle.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dss.DiscountBundle} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dss.DiscountBundle}
 */
proto.dss.DiscountBundle.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCode(value);
      break;
    case 2:
      var value = new proto.dss.SingleDiscount;
      reader.readMessage(value,proto.dss.SingleDiscount.deserializeBinaryFromReader);
      msg.addDiscounts(value);
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
proto.dss.DiscountBundle.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dss.DiscountBundle.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dss.DiscountBundle} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountBundle.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCode();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDiscountsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.dss.SingleDiscount.serializeBinaryToWriter
    );
  }
};


/**
 * optional string code = 1;
 * @return {string}
 */
proto.dss.DiscountBundle.prototype.getCode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.dss.DiscountBundle} returns this
 */
proto.dss.DiscountBundle.prototype.setCode = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated SingleDiscount discounts = 2;
 * @return {!Array<!proto.dss.SingleDiscount>}
 */
proto.dss.DiscountBundle.prototype.getDiscountsList = function() {
  return /** @type{!Array<!proto.dss.SingleDiscount>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.dss.SingleDiscount, 2));
};


/**
 * @param {!Array<!proto.dss.SingleDiscount>} value
 * @return {!proto.dss.DiscountBundle} returns this
*/
proto.dss.DiscountBundle.prototype.setDiscountsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.dss.SingleDiscount=} opt_value
 * @param {number=} opt_index
 * @return {!proto.dss.SingleDiscount}
 */
proto.dss.DiscountBundle.prototype.addDiscounts = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.dss.SingleDiscount, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.dss.DiscountBundle} returns this
 */
proto.dss.DiscountBundle.prototype.clearDiscountsList = function() {
  return this.setDiscountsList([]);
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
proto.dss.DiscountAddReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dss.DiscountAddReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dss.DiscountAddReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountAddReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    bundle: (f = msg.getBundle()) && proto.dss.DiscountBundle.toObject(includeInstance, f)
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
 * @return {!proto.dss.DiscountAddReq}
 */
proto.dss.DiscountAddReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dss.DiscountAddReq;
  return proto.dss.DiscountAddReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dss.DiscountAddReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dss.DiscountAddReq}
 */
proto.dss.DiscountAddReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.dss.DiscountBundle;
      reader.readMessage(value,proto.dss.DiscountBundle.deserializeBinaryFromReader);
      msg.setBundle(value);
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
proto.dss.DiscountAddReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dss.DiscountAddReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dss.DiscountAddReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountAddReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBundle();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.dss.DiscountBundle.serializeBinaryToWriter
    );
  }
};


/**
 * optional DiscountBundle bundle = 1;
 * @return {?proto.dss.DiscountBundle}
 */
proto.dss.DiscountAddReq.prototype.getBundle = function() {
  return /** @type{?proto.dss.DiscountBundle} */ (
    jspb.Message.getWrapperField(this, proto.dss.DiscountBundle, 1));
};


/**
 * @param {?proto.dss.DiscountBundle|undefined} value
 * @return {!proto.dss.DiscountAddReq} returns this
*/
proto.dss.DiscountAddReq.prototype.setBundle = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.dss.DiscountAddReq} returns this
 */
proto.dss.DiscountAddReq.prototype.clearBundle = function() {
  return this.setBundle(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dss.DiscountAddReq.prototype.hasBundle = function() {
  return jspb.Message.getField(this, 1) != null;
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
proto.dss.DiscountAddRes.prototype.toObject = function(opt_includeInstance) {
  return proto.dss.DiscountAddRes.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dss.DiscountAddRes} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountAddRes.toObject = function(includeInstance, msg) {
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
 * @return {!proto.dss.DiscountAddRes}
 */
proto.dss.DiscountAddRes.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dss.DiscountAddRes;
  return proto.dss.DiscountAddRes.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dss.DiscountAddRes} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dss.DiscountAddRes}
 */
proto.dss.DiscountAddRes.deserializeBinaryFromReader = function(msg, reader) {
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
proto.dss.DiscountAddRes.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dss.DiscountAddRes.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dss.DiscountAddRes} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountAddRes.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
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
proto.dss.DiscountGetReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dss.DiscountGetReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dss.DiscountGetReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountGetReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    code: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.dss.DiscountGetReq}
 */
proto.dss.DiscountGetReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dss.DiscountGetReq;
  return proto.dss.DiscountGetReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dss.DiscountGetReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dss.DiscountGetReq}
 */
proto.dss.DiscountGetReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCode(value);
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
proto.dss.DiscountGetReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dss.DiscountGetReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dss.DiscountGetReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountGetReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCode();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string code = 1;
 * @return {string}
 */
proto.dss.DiscountGetReq.prototype.getCode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.dss.DiscountGetReq} returns this
 */
proto.dss.DiscountGetReq.prototype.setCode = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
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
proto.dss.DiscountGetRes.prototype.toObject = function(opt_includeInstance) {
  return proto.dss.DiscountGetRes.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dss.DiscountGetRes} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountGetRes.toObject = function(includeInstance, msg) {
  var f, obj = {
    bundle: (f = msg.getBundle()) && proto.dss.DiscountBundle.toObject(includeInstance, f)
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
 * @return {!proto.dss.DiscountGetRes}
 */
proto.dss.DiscountGetRes.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dss.DiscountGetRes;
  return proto.dss.DiscountGetRes.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dss.DiscountGetRes} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dss.DiscountGetRes}
 */
proto.dss.DiscountGetRes.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.dss.DiscountBundle;
      reader.readMessage(value,proto.dss.DiscountBundle.deserializeBinaryFromReader);
      msg.setBundle(value);
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
proto.dss.DiscountGetRes.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dss.DiscountGetRes.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dss.DiscountGetRes} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dss.DiscountGetRes.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBundle();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.dss.DiscountBundle.serializeBinaryToWriter
    );
  }
};


/**
 * optional DiscountBundle bundle = 1;
 * @return {?proto.dss.DiscountBundle}
 */
proto.dss.DiscountGetRes.prototype.getBundle = function() {
  return /** @type{?proto.dss.DiscountBundle} */ (
    jspb.Message.getWrapperField(this, proto.dss.DiscountBundle, 1));
};


/**
 * @param {?proto.dss.DiscountBundle|undefined} value
 * @return {!proto.dss.DiscountGetRes} returns this
*/
proto.dss.DiscountGetRes.prototype.setBundle = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.dss.DiscountGetRes} returns this
 */
proto.dss.DiscountGetRes.prototype.clearBundle = function() {
  return this.setBundle(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dss.DiscountGetRes.prototype.hasBundle = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.dss);