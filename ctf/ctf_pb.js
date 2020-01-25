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
goog.exportSymbol('proto.roleypoly.ctf.Flag', null, global);
goog.exportSymbol('proto.roleypoly.ctf.Flags', null, global);
goog.exportSymbol('proto.roleypoly.ctf.Ring', null, global);
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
proto.roleypoly.ctf.Flags = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.roleypoly.ctf.Flags.repeatedFields_, null);
};
goog.inherits(proto.roleypoly.ctf.Flags, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.roleypoly.ctf.Flags.displayName = 'proto.roleypoly.ctf.Flags';
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
proto.roleypoly.ctf.Flag = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.roleypoly.ctf.Flag, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.roleypoly.ctf.Flag.displayName = 'proto.roleypoly.ctf.Flag';
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
proto.roleypoly.ctf.Ring = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.roleypoly.ctf.Ring, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.roleypoly.ctf.Ring.displayName = 'proto.roleypoly.ctf.Ring';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.roleypoly.ctf.Flags.repeatedFields_ = [1];



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
proto.roleypoly.ctf.Flags.prototype.toObject = function(opt_includeInstance) {
  return proto.roleypoly.ctf.Flags.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.roleypoly.ctf.Flags} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Flags.toObject = function(includeInstance, msg) {
  var f, obj = {
    flagsList: jspb.Message.toObjectList(msg.getFlagsList(),
    proto.roleypoly.ctf.Flag.toObject, includeInstance)
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
 * @return {!proto.roleypoly.ctf.Flags}
 */
proto.roleypoly.ctf.Flags.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.roleypoly.ctf.Flags;
  return proto.roleypoly.ctf.Flags.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.roleypoly.ctf.Flags} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.roleypoly.ctf.Flags}
 */
proto.roleypoly.ctf.Flags.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.roleypoly.ctf.Flag;
      reader.readMessage(value,proto.roleypoly.ctf.Flag.deserializeBinaryFromReader);
      msg.addFlags(value);
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
proto.roleypoly.ctf.Flags.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.roleypoly.ctf.Flags.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.roleypoly.ctf.Flags} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Flags.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFlagsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.roleypoly.ctf.Flag.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Flag flags = 1;
 * @return {!Array<!proto.roleypoly.ctf.Flag>}
 */
proto.roleypoly.ctf.Flags.prototype.getFlagsList = function() {
  return /** @type{!Array<!proto.roleypoly.ctf.Flag>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.roleypoly.ctf.Flag, 1));
};


/** @param {!Array<!proto.roleypoly.ctf.Flag>} value */
proto.roleypoly.ctf.Flags.prototype.setFlagsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.roleypoly.ctf.Flag=} opt_value
 * @param {number=} opt_index
 * @return {!proto.roleypoly.ctf.Flag}
 */
proto.roleypoly.ctf.Flags.prototype.addFlags = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.roleypoly.ctf.Flag, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.roleypoly.ctf.Flags.prototype.clearFlagsList = function() {
  this.setFlagsList([]);
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
proto.roleypoly.ctf.Flag.prototype.toObject = function(opt_includeInstance) {
  return proto.roleypoly.ctf.Flag.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.roleypoly.ctf.Flag} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Flag.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    ring: jspb.Message.getFieldWithDefault(msg, 2, 0)
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
 * @return {!proto.roleypoly.ctf.Flag}
 */
proto.roleypoly.ctf.Flag.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.roleypoly.ctf.Flag;
  return proto.roleypoly.ctf.Flag.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.roleypoly.ctf.Flag} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.roleypoly.ctf.Flag}
 */
proto.roleypoly.ctf.Flag.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRing(value);
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
proto.roleypoly.ctf.Flag.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.roleypoly.ctf.Flag.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.roleypoly.ctf.Flag} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Flag.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRing();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.roleypoly.ctf.Flag.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.roleypoly.ctf.Flag.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional int32 ring = 2;
 * @return {number}
 */
proto.roleypoly.ctf.Flag.prototype.getRing = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.roleypoly.ctf.Flag.prototype.setRing = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
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
proto.roleypoly.ctf.Ring.prototype.toObject = function(opt_includeInstance) {
  return proto.roleypoly.ctf.Ring.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.roleypoly.ctf.Ring} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Ring.toObject = function(includeInstance, msg) {
  var f, obj = {
    ring: jspb.Message.getFieldWithDefault(msg, 1, 0)
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
 * @return {!proto.roleypoly.ctf.Ring}
 */
proto.roleypoly.ctf.Ring.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.roleypoly.ctf.Ring;
  return proto.roleypoly.ctf.Ring.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.roleypoly.ctf.Ring} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.roleypoly.ctf.Ring}
 */
proto.roleypoly.ctf.Ring.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRing(value);
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
proto.roleypoly.ctf.Ring.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.roleypoly.ctf.Ring.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.roleypoly.ctf.Ring} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.roleypoly.ctf.Ring.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRing();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
};


/**
 * optional int32 ring = 1;
 * @return {number}
 */
proto.roleypoly.ctf.Ring.prototype.getRing = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.roleypoly.ctf.Ring.prototype.setRing = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


goog.object.extend(exports, proto.roleypoly.ctf);
