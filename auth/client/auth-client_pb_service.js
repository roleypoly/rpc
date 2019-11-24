// package: roleypoly.auth.client
// file: auth/client/auth-client.proto

var auth_client_auth_client_pb = require("../../auth/client/auth-client_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var AuthClient = (function () {
  function AuthClient() {}
  AuthClient.serviceName = "roleypoly.auth.client.AuthClient";
  return AuthClient;
}());

exports.AuthClient = AuthClient;

function AuthClientClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

exports.AuthClientClient = AuthClientClient;

