/**
 * @fileoverview gRPC-Web generated client stub for buildsafe.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: buildsafe/v1/repositories.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var buf_validate_validate_pb = require('../../buf/validate/validate_pb.js')

var google_api_annotations_pb = require('../../google/api/annotations_pb.js')
const proto = {};
proto.buildsafe = {};
proto.buildsafe.v1 = require('./repositories_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.buildsafe.v1.RepositoryServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.buildsafe.v1.RepositoryServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.buildsafe.v1.ListRepositoriesRequest,
 *   !proto.buildsafe.v1.ListRepositoriesResponse>}
 */
const methodDescriptor_RepositoryService_ListRepositories = new grpc.web.MethodDescriptor(
  '/buildsafe.v1.RepositoryService/ListRepositories',
  grpc.web.MethodType.UNARY,
  proto.buildsafe.v1.ListRepositoriesRequest,
  proto.buildsafe.v1.ListRepositoriesResponse,
  /**
   * @param {!proto.buildsafe.v1.ListRepositoriesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.buildsafe.v1.ListRepositoriesResponse.deserializeBinary
);


/**
 * @param {!proto.buildsafe.v1.ListRepositoriesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.buildsafe.v1.ListRepositoriesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.buildsafe.v1.ListRepositoriesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.buildsafe.v1.RepositoryServiceClient.prototype.listRepositories =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/buildsafe.v1.RepositoryService/ListRepositories',
      request,
      metadata || {},
      methodDescriptor_RepositoryService_ListRepositories,
      callback);
};


/**
 * @param {!proto.buildsafe.v1.ListRepositoriesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.buildsafe.v1.ListRepositoriesResponse>}
 *     Promise that resolves to the response
 */
proto.buildsafe.v1.RepositoryServicePromiseClient.prototype.listRepositories =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/buildsafe.v1.RepositoryService/ListRepositories',
      request,
      metadata || {},
      methodDescriptor_RepositoryService_ListRepositories);
};


module.exports = proto.buildsafe.v1;

