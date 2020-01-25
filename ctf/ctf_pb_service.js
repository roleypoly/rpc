// package: roleypoly.ctf
// file: ctf/ctf.proto

var ctf_ctf_pb = require("../ctf/ctf_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var CTF = (function () {
  function CTF() {}
  CTF.serviceName = "roleypoly.ctf.CTF";
  return CTF;
}());

CTF.GetRingFlags = {
  methodName: "GetRingFlags",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.Ring,
  responseType: ctf_ctf_pb.Flags
};

CTF.CreateFlag = {
  methodName: "CreateFlag",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.Flag,
  responseType: ctf_ctf_pb.Flag
};

CTF.PromoteFlag = {
  methodName: "PromoteFlag",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.Flag,
  responseType: ctf_ctf_pb.Flag
};

CTF.RemoveFlag = {
  methodName: "RemoveFlag",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.Flag,
  responseType: google_protobuf_empty_pb.Empty
};

exports.CTF = CTF;

function CTFClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CTFClient.prototype.getRingFlags = function getRingFlags(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.GetRingFlags, {
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

CTFClient.prototype.createFlag = function createFlag(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.CreateFlag, {
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

CTFClient.prototype.promoteFlag = function promoteFlag(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.PromoteFlag, {
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

CTFClient.prototype.removeFlag = function removeFlag(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.RemoveFlag, {
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

exports.CTFClient = CTFClient;

