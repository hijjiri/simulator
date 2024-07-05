/**
 * @fileoverview gRPC-Web generated client stub for janken
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.janken = require('./janken_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.janken.JankenServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.janken.JankenServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 *   !proto.janken.JankenRequest,
 *   !proto.janken.JankenResponse>}
 */
const methodDescriptor_JankenService_PlayJanken = new grpc.web.MethodDescriptor(
  '/janken.JankenService/PlayJanken',
  grpc.web.MethodType.UNARY,
  proto.janken.JankenRequest,
  proto.janken.JankenResponse,
  /**
   * @param {!proto.janken.JankenRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.janken.JankenResponse.deserializeBinary
);


/**
 * @param {!proto.janken.JankenRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.janken.JankenResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.janken.JankenResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.janken.JankenServiceClient.prototype.playJanken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/janken.JankenService/PlayJanken',
      request,
      metadata || {},
      methodDescriptor_JankenService_PlayJanken,
      callback);
};


/**
 * @param {!proto.janken.JankenRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.janken.JankenResponse>}
 *     Promise that resolves to the response
 */
proto.janken.JankenServicePromiseClient.prototype.playJanken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/janken.JankenService/PlayJanken',
      request,
      metadata || {},
      methodDescriptor_JankenService_PlayJanken);
};


module.exports = proto.janken;

