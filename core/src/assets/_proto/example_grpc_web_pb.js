/**
 * @fileoverview gRPC-Web generated client stub for example
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.example = require('./example_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.example.ExampleServiceClient =
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
proto.example.ExampleServicePromiseClient =
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
 *   !proto.example.HelloRequest,
 *   !proto.example.HelloResponse>}
 */
const methodDescriptor_ExampleService_SayHello = new grpc.web.MethodDescriptor(
  '/example.ExampleService/SayHello',
  grpc.web.MethodType.UNARY,
  proto.example.HelloRequest,
  proto.example.HelloResponse,
  /**
   * @param {!proto.example.HelloRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.example.HelloResponse.deserializeBinary
);


/**
 * @param {!proto.example.HelloRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.example.HelloResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.example.HelloResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.example.ExampleServiceClient.prototype.sayHello =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/example.ExampleService/SayHello',
      request,
      metadata || {},
      methodDescriptor_ExampleService_SayHello,
      callback);
};


/**
 * @param {!proto.example.HelloRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.example.HelloResponse>}
 *     Promise that resolves to the response
 */
proto.example.ExampleServicePromiseClient.prototype.sayHello =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/example.ExampleService/SayHello',
      request,
      metadata || {},
      methodDescriptor_ExampleService_SayHello);
};


module.exports = proto.example;

