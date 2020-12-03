/**
 * @fileoverview gRPC-Web generated client stub for tendermint.rpc.grpc
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var tendermint_abci_types_pb = require('../../../tendermint/abci/types_pb.js')
const proto = {};
proto.tendermint = {};
proto.tendermint.rpc = {};
proto.tendermint.rpc.grpc = require('./types_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tendermint.rpc.grpc.BroadcastAPIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tendermint.rpc.grpc.BroadcastAPIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tendermint.rpc.grpc.RequestPing,
 *   !proto.tendermint.rpc.grpc.ResponsePing>}
 */
const methodDescriptor_BroadcastAPI_Ping = new grpc.web.MethodDescriptor(
  '/tendermint.rpc.grpc.BroadcastAPI/Ping',
  grpc.web.MethodType.UNARY,
  proto.tendermint.rpc.grpc.RequestPing,
  proto.tendermint.rpc.grpc.ResponsePing,
  /**
   * @param {!proto.tendermint.rpc.grpc.RequestPing} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tendermint.rpc.grpc.ResponsePing.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tendermint.rpc.grpc.RequestPing,
 *   !proto.tendermint.rpc.grpc.ResponsePing>}
 */
const methodInfo_BroadcastAPI_Ping = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tendermint.rpc.grpc.ResponsePing,
  /**
   * @param {!proto.tendermint.rpc.grpc.RequestPing} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tendermint.rpc.grpc.ResponsePing.deserializeBinary
);


/**
 * @param {!proto.tendermint.rpc.grpc.RequestPing} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tendermint.rpc.grpc.ResponsePing)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tendermint.rpc.grpc.ResponsePing>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tendermint.rpc.grpc.BroadcastAPIClient.prototype.ping =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tendermint.rpc.grpc.BroadcastAPI/Ping',
      request,
      metadata || {},
      methodDescriptor_BroadcastAPI_Ping,
      callback);
};


/**
 * @param {!proto.tendermint.rpc.grpc.RequestPing} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tendermint.rpc.grpc.ResponsePing>}
 *     Promise that resolves to the response
 */
proto.tendermint.rpc.grpc.BroadcastAPIPromiseClient.prototype.ping =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tendermint.rpc.grpc.BroadcastAPI/Ping',
      request,
      metadata || {},
      methodDescriptor_BroadcastAPI_Ping);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tendermint.rpc.grpc.RequestBroadcastTx,
 *   !proto.tendermint.rpc.grpc.ResponseBroadcastTx>}
 */
const methodDescriptor_BroadcastAPI_BroadcastTx = new grpc.web.MethodDescriptor(
  '/tendermint.rpc.grpc.BroadcastAPI/BroadcastTx',
  grpc.web.MethodType.UNARY,
  proto.tendermint.rpc.grpc.RequestBroadcastTx,
  proto.tendermint.rpc.grpc.ResponseBroadcastTx,
  /**
   * @param {!proto.tendermint.rpc.grpc.RequestBroadcastTx} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tendermint.rpc.grpc.ResponseBroadcastTx.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tendermint.rpc.grpc.RequestBroadcastTx,
 *   !proto.tendermint.rpc.grpc.ResponseBroadcastTx>}
 */
const methodInfo_BroadcastAPI_BroadcastTx = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tendermint.rpc.grpc.ResponseBroadcastTx,
  /**
   * @param {!proto.tendermint.rpc.grpc.RequestBroadcastTx} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tendermint.rpc.grpc.ResponseBroadcastTx.deserializeBinary
);


/**
 * @param {!proto.tendermint.rpc.grpc.RequestBroadcastTx} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tendermint.rpc.grpc.ResponseBroadcastTx)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tendermint.rpc.grpc.ResponseBroadcastTx>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tendermint.rpc.grpc.BroadcastAPIClient.prototype.broadcastTx =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tendermint.rpc.grpc.BroadcastAPI/BroadcastTx',
      request,
      metadata || {},
      methodDescriptor_BroadcastAPI_BroadcastTx,
      callback);
};


/**
 * @param {!proto.tendermint.rpc.grpc.RequestBroadcastTx} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tendermint.rpc.grpc.ResponseBroadcastTx>}
 *     Promise that resolves to the response
 */
proto.tendermint.rpc.grpc.BroadcastAPIPromiseClient.prototype.broadcastTx =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tendermint.rpc.grpc.BroadcastAPI/BroadcastTx',
      request,
      metadata || {},
      methodDescriptor_BroadcastAPI_BroadcastTx);
};


module.exports = proto.tendermint.rpc.grpc;

