// package: roleypoly.auth.client
// file: auth/client/auth-client.proto

var auth_client_auth_client_pb = require("../../auth/client/auth-client_pb");
var shared_internal_pb = require("../../shared/internal_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var auth_shared_pb = require("../../auth/shared_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var AuthClient = (function () {
  function AuthClient() {}
  AuthClient.serviceName = "roleypoly.auth.client.AuthClient";
  return AuthClient;
}());

AuthClient.GetClientToken = {
  methodName: "GetClientToken",
  service: AuthClient,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: auth_shared_pb.Token
};

AuthClient.GetUserSession = {
  methodName: "GetUserSession",
  service: AuthClient,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: shared_internal_pb.RoleypolySession
};

AuthClient.ResolveSessionKey = {
  methodName: "ResolveSessionKey",
  service: AuthClient,
  requestStream: false,
  responseStream: false,
  requestType: auth_shared_pb.Token,
  responseType: auth_shared_pb.Token
};

AuthClient.AuthorizeChallenge = {
  methodName: "AuthorizeChallenge",
  service: AuthClient,
  requestStream: false,
  responseStream: false,
  requestType: auth_shared_pb.AuthChallenge,
  responseType: auth_shared_pb.Token
};

exports.AuthClient = AuthClient;

function AuthClientClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AuthClientClient.prototype.getClientToken = function getClientToken(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(AuthClient.GetClientToken, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

AuthClientClient.prototype.getUserSession = function getUserSession(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(AuthClient.GetUserSession, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

AuthClientClient.prototype.resolveSessionKey = function resolveSessionKey(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(AuthClient.ResolveSessionKey, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

AuthClientClient.prototype.authorizeChallenge = function authorizeChallenge(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(AuthClient.AuthorizeChallenge, {
      request: requestMessage,
      host: this.serviceHost,
      metadata: metadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: function (response) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          reject(err);
        } else {
          resolve(response.message);
        }
      }
    });
  });
};

exports.AuthClientClient = AuthClientClient;

