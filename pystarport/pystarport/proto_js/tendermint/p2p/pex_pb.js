// source: tendermint/p2p/pex.proto
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

var tendermint_p2p_types_pb = require('../../tendermint/p2p/types_pb.js');
goog.object.extend(proto, tendermint_p2p_types_pb);
var gogoproto_gogo_pb = require('../../gogoproto/gogo_pb.js');
goog.object.extend(proto, gogoproto_gogo_pb);
goog.exportSymbol('proto.tendermint.p2p.Message', null, global);
goog.exportSymbol('proto.tendermint.p2p.Message.SumCase', null, global);
goog.exportSymbol('proto.tendermint.p2p.PexAddrs', null, global);
goog.exportSymbol('proto.tendermint.p2p.PexRequest', null, global);
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
proto.tendermint.p2p.PexRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.p2p.PexRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.p2p.PexRequest.displayName = 'proto.tendermint.p2p.PexRequest';
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
proto.tendermint.p2p.PexAddrs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.tendermint.p2p.PexAddrs.repeatedFields_, null);
};
goog.inherits(proto.tendermint.p2p.PexAddrs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.p2p.PexAddrs.displayName = 'proto.tendermint.p2p.PexAddrs';
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
proto.tendermint.p2p.Message = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.tendermint.p2p.Message.oneofGroups_);
};
goog.inherits(proto.tendermint.p2p.Message, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.p2p.Message.displayName = 'proto.tendermint.p2p.Message';
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
proto.tendermint.p2p.PexRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.p2p.PexRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.p2p.PexRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.p2p.PexRequest.toObject = function(includeInstance, msg) {
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
 * @return {!proto.tendermint.p2p.PexRequest}
 */
proto.tendermint.p2p.PexRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.p2p.PexRequest;
  return proto.tendermint.p2p.PexRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.p2p.PexRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.p2p.PexRequest}
 */
proto.tendermint.p2p.PexRequest.deserializeBinaryFromReader = function(msg, reader) {
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
proto.tendermint.p2p.PexRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.p2p.PexRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.p2p.PexRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.p2p.PexRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.tendermint.p2p.PexAddrs.repeatedFields_ = [1];



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
proto.tendermint.p2p.PexAddrs.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.p2p.PexAddrs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.p2p.PexAddrs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.p2p.PexAddrs.toObject = function(includeInstance, msg) {
  var f, obj = {
    addrsList: jspb.Message.toObjectList(msg.getAddrsList(),
    tendermint_p2p_types_pb.NetAddress.toObject, includeInstance)
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
 * @return {!proto.tendermint.p2p.PexAddrs}
 */
proto.tendermint.p2p.PexAddrs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.p2p.PexAddrs;
  return proto.tendermint.p2p.PexAddrs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.p2p.PexAddrs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.p2p.PexAddrs}
 */
proto.tendermint.p2p.PexAddrs.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new tendermint_p2p_types_pb.NetAddress;
      reader.readMessage(value,tendermint_p2p_types_pb.NetAddress.deserializeBinaryFromReader);
      msg.addAddrs(value);
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
proto.tendermint.p2p.PexAddrs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.p2p.PexAddrs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.p2p.PexAddrs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.p2p.PexAddrs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAddrsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      tendermint_p2p_types_pb.NetAddress.serializeBinaryToWriter
    );
  }
};


/**
 * repeated NetAddress addrs = 1;
 * @return {!Array<!proto.tendermint.p2p.NetAddress>}
 */
proto.tendermint.p2p.PexAddrs.prototype.getAddrsList = function() {
  return /** @type{!Array<!proto.tendermint.p2p.NetAddress>} */ (
    jspb.Message.getRepeatedWrapperField(this, tendermint_p2p_types_pb.NetAddress, 1));
};


/**
 * @param {!Array<!proto.tendermint.p2p.NetAddress>} value
 * @return {!proto.tendermint.p2p.PexAddrs} returns this
*/
proto.tendermint.p2p.PexAddrs.prototype.setAddrsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.tendermint.p2p.NetAddress=} opt_value
 * @param {number=} opt_index
 * @return {!proto.tendermint.p2p.NetAddress}
 */
proto.tendermint.p2p.PexAddrs.prototype.addAddrs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.tendermint.p2p.NetAddress, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.tendermint.p2p.PexAddrs} returns this
 */
proto.tendermint.p2p.PexAddrs.prototype.clearAddrsList = function() {
  return this.setAddrsList([]);
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.tendermint.p2p.Message.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.tendermint.p2p.Message.SumCase = {
  SUM_NOT_SET: 0,
  PEX_REQUEST: 1,
  PEX_ADDRS: 2
};

/**
 * @return {proto.tendermint.p2p.Message.SumCase}
 */
proto.tendermint.p2p.Message.prototype.getSumCase = function() {
  return /** @type {proto.tendermint.p2p.Message.SumCase} */(jspb.Message.computeOneofCase(this, proto.tendermint.p2p.Message.oneofGroups_[0]));
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
proto.tendermint.p2p.Message.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.p2p.Message.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.p2p.Message} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.p2p.Message.toObject = function(includeInstance, msg) {
  var f, obj = {
    pexRequest: (f = msg.getPexRequest()) && proto.tendermint.p2p.PexRequest.toObject(includeInstance, f),
    pexAddrs: (f = msg.getPexAddrs()) && proto.tendermint.p2p.PexAddrs.toObject(includeInstance, f)
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
 * @return {!proto.tendermint.p2p.Message}
 */
proto.tendermint.p2p.Message.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.p2p.Message;
  return proto.tendermint.p2p.Message.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.p2p.Message} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.p2p.Message}
 */
proto.tendermint.p2p.Message.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.tendermint.p2p.PexRequest;
      reader.readMessage(value,proto.tendermint.p2p.PexRequest.deserializeBinaryFromReader);
      msg.setPexRequest(value);
      break;
    case 2:
      var value = new proto.tendermint.p2p.PexAddrs;
      reader.readMessage(value,proto.tendermint.p2p.PexAddrs.deserializeBinaryFromReader);
      msg.setPexAddrs(value);
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
proto.tendermint.p2p.Message.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.p2p.Message.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.p2p.Message} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.p2p.Message.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPexRequest();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.tendermint.p2p.PexRequest.serializeBinaryToWriter
    );
  }
  f = message.getPexAddrs();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.tendermint.p2p.PexAddrs.serializeBinaryToWriter
    );
  }
};


/**
 * optional PexRequest pex_request = 1;
 * @return {?proto.tendermint.p2p.PexRequest}
 */
proto.tendermint.p2p.Message.prototype.getPexRequest = function() {
  return /** @type{?proto.tendermint.p2p.PexRequest} */ (
    jspb.Message.getWrapperField(this, proto.tendermint.p2p.PexRequest, 1));
};


/**
 * @param {?proto.tendermint.p2p.PexRequest|undefined} value
 * @return {!proto.tendermint.p2p.Message} returns this
*/
proto.tendermint.p2p.Message.prototype.setPexRequest = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.tendermint.p2p.Message.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.p2p.Message} returns this
 */
proto.tendermint.p2p.Message.prototype.clearPexRequest = function() {
  return this.setPexRequest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.p2p.Message.prototype.hasPexRequest = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional PexAddrs pex_addrs = 2;
 * @return {?proto.tendermint.p2p.PexAddrs}
 */
proto.tendermint.p2p.Message.prototype.getPexAddrs = function() {
  return /** @type{?proto.tendermint.p2p.PexAddrs} */ (
    jspb.Message.getWrapperField(this, proto.tendermint.p2p.PexAddrs, 2));
};


/**
 * @param {?proto.tendermint.p2p.PexAddrs|undefined} value
 * @return {!proto.tendermint.p2p.Message} returns this
*/
proto.tendermint.p2p.Message.prototype.setPexAddrs = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.tendermint.p2p.Message.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.p2p.Message} returns this
 */
proto.tendermint.p2p.Message.prototype.clearPexAddrs = function() {
  return this.setPexAddrs(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.p2p.Message.prototype.hasPexAddrs = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.tendermint.p2p);
