// source: ctf/ctf.proto
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

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.object.extend(proto, google_protobuf_empty_pb);
goog.exportSymbol('proto.roleypoly.ctf.Canaries', null, global);
goog.exportSymbol('proto.roleypoly.ctf.Canary', null, global);
goog.exportSymbol('proto.roleypoly.ctf.CanaryQuery', null, global);
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
proto.roleypoly.ctf.Canary = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.roleypoly.ctf.Canary, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.roleypoly.ctf.Canary.displayName = 'proto.roleypoly.ctf.Canary';
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
proto.roleypoly.ctf.Canaries = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.roleypoly.ctf.Canaries.repeatedFields_, null);
};
goog.inherits(proto.roleypoly.ctf.Canaries, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.roleypoly.ctf.Canaries.displayName = 'proto.roleypoly.ctf.Canaries';
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
proto.roleypoly.ctf.CanaryQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.roleypoly.ctf.CanaryQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.roleypoly.ctf.CanaryQuery.displayName = 'proto.roleypoly.ctf.CanaryQuery';
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
proto.roleypoly.ctf.Canary.prototype.toObject = function(opt_includeInstance) {
  return proto.roleypoly.ctf.Canary.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.roleypoly.ctf.Canary} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Canary.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    percent: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0)
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
 * @return {!proto.roleypoly.ctf.Canary}
 */
proto.roleypoly.ctf.Canary.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.roleypoly.ctf.Canary;
  return proto.roleypoly.ctf.Canary.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.roleypoly.ctf.Canary} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.roleypoly.ctf.Canary}
 */
proto.roleypoly.ctf.Canary.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {number} */ (reader.readFloat());
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
proto.roleypoly.ctf.Canary.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.roleypoly.ctf.Canary.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.roleypoly.ctf.Canary} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Canary.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPercent();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.roleypoly.ctf.Canary.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.roleypoly.ctf.Canary.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional float percent = 2;
 * @return {number}
 */
proto.roleypoly.ctf.Canary.prototype.getPercent = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/** @param {number} value */
proto.roleypoly.ctf.Canary.prototype.setPercent = function(value) {
  jspb.Message.setProto3FloatField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.roleypoly.ctf.Canaries.repeatedFields_ = [1];



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
proto.roleypoly.ctf.Canaries.prototype.toObject = function(opt_includeInstance) {
  return proto.roleypoly.ctf.Canaries.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.roleypoly.ctf.Canaries} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Canaries.toObject = function(includeInstance, msg) {
  var f, obj = {
    canariesList: jspb.Message.toObjectList(msg.getCanariesList(),
    proto.roleypoly.ctf.Canary.toObject, includeInstance)
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
 * @return {!proto.roleypoly.ctf.Canaries}
 */
proto.roleypoly.ctf.Canaries.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.roleypoly.ctf.Canaries;
  return proto.roleypoly.ctf.Canaries.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.roleypoly.ctf.Canaries} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.roleypoly.ctf.Canaries}
 */
proto.roleypoly.ctf.Canaries.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.roleypoly.ctf.Canary;
      reader.readMessage(value,proto.roleypoly.ctf.Canary.deserializeBinaryFromReader);
      msg.addCanaries(value);
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
proto.roleypoly.ctf.Canaries.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.roleypoly.ctf.Canaries.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.roleypoly.ctf.Canaries} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Canaries.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCanariesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.roleypoly.ctf.Canary.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Canary canaries = 1;
 * @return {!Array<!proto.roleypoly.ctf.Canary>}
 */
proto.roleypoly.ctf.Canaries.prototype.getCanariesList = function() {
  return /** @type{!Array<!proto.roleypoly.ctf.Canary>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.roleypoly.ctf.Canary, 1));
};


/** @param {!Array<!proto.roleypoly.ctf.Canary>} value */
proto.roleypoly.ctf.Canaries.prototype.setCanariesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.roleypoly.ctf.Canary=} opt_value
 * @param {number=} opt_index
 * @return {!proto.roleypoly.ctf.Canary}
 */
proto.roleypoly.ctf.Canaries.prototype.addCanaries = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.roleypoly.ctf.Canary, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.roleypoly.ctf.Canaries.prototype.clearCanariesList = function() {
  this.setCanariesList([]);
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
proto.roleypoly.ctf.CanaryQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.roleypoly.ctf.CanaryQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.roleypoly.ctf.CanaryQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.CanaryQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    threshold: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0),
    canary: (f = msg.getCanary()) && proto.roleypoly.ctf.Canary.toObject(includeInstance, f)
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
 * @return {!proto.roleypoly.ctf.CanaryQuery}
 */
proto.roleypoly.ctf.CanaryQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.roleypoly.ctf.CanaryQuery;
  return proto.roleypoly.ctf.CanaryQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.roleypoly.ctf.CanaryQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.roleypoly.ctf.CanaryQuery}
 */
proto.roleypoly.ctf.CanaryQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setThreshold(value);
      break;
    case 2:
      var value = new proto.roleypoly.ctf.Canary;
      reader.readMessage(value,proto.roleypoly.ctf.Canary.deserializeBinaryFromReader);
      msg.setCanary(value);
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
proto.roleypoly.ctf.CanaryQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.roleypoly.ctf.CanaryQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.roleypoly.ctf.CanaryQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.CanaryQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getThreshold();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = message.getCanary();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.roleypoly.ctf.Canary.serializeBinaryToWriter
    );
  }
};


/**
 * optional float threshold = 1;
 * @return {number}
 */
proto.roleypoly.ctf.CanaryQuery.prototype.getThreshold = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/** @param {number} value */
proto.roleypoly.ctf.CanaryQuery.prototype.setThreshold = function(value) {
  jspb.Message.setProto3FloatField(this, 1, value);
};


/**
 * optional Canary canary = 2;
 * @return {?proto.roleypoly.ctf.Canary}
 */
proto.roleypoly.ctf.CanaryQuery.prototype.getCanary = function() {
  return /** @type{?proto.roleypoly.ctf.Canary} */ (
    jspb.Message.getWrapperField(this, proto.roleypoly.ctf.Canary, 2));
};


/** @param {?proto.roleypoly.ctf.Canary|undefined} value */
proto.roleypoly.ctf.CanaryQuery.prototype.setCanary = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.roleypoly.ctf.CanaryQuery.prototype.clearCanary = function() {
  this.setCanary(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.roleypoly.ctf.CanaryQuery.prototype.hasCanary = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.roleypoly.ctf);
