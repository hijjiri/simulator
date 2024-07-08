/**
 * @fileoverview gRPC-Web generated client stub for simulator
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.simulator = require('./simulator_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.simulator.SimulatorServiceClient =
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
proto.simulator.SimulatorServicePromiseClient =
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
 *   !proto.simulator.SimulateRequest,
 *   !proto.simulator.SimulateResponse>}
 */
const methodDescriptor_SimulatorService_SimulateBattle = new grpc.web.MethodDescriptor(
  '/simulator.SimulatorService/SimulateBattle',
  grpc.web.MethodType.UNARY,
  proto.simulator.SimulateRequest,
  proto.simulator.SimulateResponse,
  /**
   * @param {!proto.simulator.SimulateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.simulator.SimulateResponse.deserializeBinary
);


/**
 * @param {!proto.simulator.SimulateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.simulator.SimulateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.simulator.SimulateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.simulator.SimulatorServiceClient.prototype.simulateBattle =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/simulator.SimulatorService/SimulateBattle',
      request,
      metadata || {},
      methodDescriptor_SimulatorService_SimulateBattle,
      callback);
};


/**
 * @param {!proto.simulator.SimulateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.simulator.SimulateResponse>}
 *     Promise that resolves to the response
 */
proto.simulator.SimulatorServicePromiseClient.prototype.simulateBattle =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/simulator.SimulatorService/SimulateBattle',
      request,
      metadata || {},
      methodDescriptor_SimulatorService_SimulateBattle);
};


module.exports = proto.simulator;

