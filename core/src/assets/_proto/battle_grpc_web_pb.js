/**
 * @fileoverview gRPC-Web generated client stub for battle
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')
const proto = {};
proto.battle = require('./battle_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.battle.BattleServiceClient =
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
proto.battle.BattleServicePromiseClient =
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
 *   !proto.battle.BattleId,
 *   !proto.battle.Battle>}
 */
const methodDescriptor_BattleService_GetBattle = new grpc.web.MethodDescriptor(
  '/battle.BattleService/GetBattle',
  grpc.web.MethodType.UNARY,
  proto.battle.BattleId,
  proto.battle.Battle,
  /**
   * @param {!proto.battle.BattleId} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.battle.Battle.deserializeBinary
);


/**
 * @param {!proto.battle.BattleId} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.battle.Battle)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.battle.Battle>|undefined}
 *     The XHR Node Readable Stream
 */
proto.battle.BattleServiceClient.prototype.getBattle =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/battle.BattleService/GetBattle',
      request,
      metadata || {},
      methodDescriptor_BattleService_GetBattle,
      callback);
};


/**
 * @param {!proto.battle.BattleId} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.battle.Battle>}
 *     Promise that resolves to the response
 */
proto.battle.BattleServicePromiseClient.prototype.getBattle =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/battle.BattleService/GetBattle',
      request,
      metadata || {},
      methodDescriptor_BattleService_GetBattle);
};


module.exports = proto.battle;

