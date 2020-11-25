// source: cosmos/evidence/v1beta1/tx.proto
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

var gogoproto_gogo_pb = require('../../../gogoproto/gogo_pb.js');
goog.object.extend(proto, gogoproto_gogo_pb);
var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js');
goog.object.extend(proto, google_protobuf_any_pb);
var cosmos_proto_cosmos_pb = require('../../../cosmos_proto/cosmos_pb.js');
goog.object.extend(proto, cosmos_proto_cosmos_pb);
goog.exportSymbol('proto.cosmos.evidence.v1beta1.MsgSubmitEvidence', null, global);
goog.exportSymbol('proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse', null, global);
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
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.cosmos.evidence.v1beta1.MsgSubmitEvidence, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.displayName = 'proto.cosmos.evidence.v1beta1.MsgSubmitEvidence';
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
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.displayName = 'proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse';
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
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.toObject = function(opt_includeInstance) {
  return proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.toObject = function(includeInstance, msg) {
  var f, obj = {
    submitter: jspb.Message.getFieldWithDefault(msg, 1, ""),
    evidence: (f = msg.getEvidence()) && google_protobuf_any_pb.Any.toObject(includeInstance, f)
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
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.cosmos.evidence.v1beta1.MsgSubmitEvidence;
  return proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setSubmitter(value);
      break;
    case 2:
      var value = new google_protobuf_any_pb.Any;
      reader.readMessage(value,google_protobuf_any_pb.Any.deserializeBinaryFromReader);
      msg.setEvidence(value);
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
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSubmitter();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getEvidence();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_any_pb.Any.serializeBinaryToWriter
    );
  }
};


/**
 * optional string submitter = 1;
 * @return {string}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.getSubmitter = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence} returns this
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.setSubmitter = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.Any evidence = 2;
 * @return {?proto.google.protobuf.Any}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.getEvidence = function() {
  return /** @type{?proto.google.protobuf.Any} */ (
    jspb.Message.getWrapperField(this, google_protobuf_any_pb.Any, 2));
};


/**
 * @param {?proto.google.protobuf.Any|undefined} value
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence} returns this
*/
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.setEvidence = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidence} returns this
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.clearEvidence = function() {
  return this.setEvidence(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidence.prototype.hasEvidence = function() {
  return jspb.Message.getField(this, 2) != null;
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
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    hash: msg.getHash_asB64()
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
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse;
  return proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 4:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setHash(value);
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
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHash_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      4,
      f
    );
  }
};


/**
 * optional bytes hash = 4;
 * @return {string}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.prototype.getHash = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * optional bytes hash = 4;
 * This is a type-conversion wrapper around `getHash()`
 * @return {string}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.prototype.getHash_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getHash()));
};


/**
 * optional bytes hash = 4;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getHash()`
 * @return {!Uint8Array}
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.prototype.getHash_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getHash()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse} returns this
 */
proto.cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.prototype.setHash = function(value) {
  return jspb.Message.setProto3BytesField(this, 4, value);
};


goog.object.extend(exports, proto.cosmos.evidence.v1beta1);
