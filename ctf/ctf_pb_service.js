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

CTF.GetCanaries = {
  methodName: "GetCanaries",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.CanaryQuery,
  responseType: ctf_ctf_pb.Canaries
};

CTF.SetSingleCanaryRollout = {
  methodName: "SetSingleCanaryRollout",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.CanaryQuery,
  responseType: google_protobuf_empty_pb.Empty
};

CTF.SetGlobalCanaryRollout = {
  methodName: "SetGlobalCanaryRollout",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.CanaryQuery,
  responseType: google_protobuf_empty_pb.Empty
};

CTF.CreateCanary = {
  methodName: "CreateCanary",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.Canary,
  responseType: ctf_ctf_pb.Canary
};

CTF.DeleteCanary = {
  methodName: "DeleteCanary",
  service: CTF,
  requestStream: false,
  responseStream: false,
  requestType: ctf_ctf_pb.CanaryQuery,
  responseType: google_protobuf_empty_pb.Empty
};

exports.CTF = CTF;

function CTFClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CTFClient.prototype.getCanaries = function getCanaries(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.GetCanaries, {
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

CTFClient.prototype.setSingleCanaryRollout = function setSingleCanaryRollout(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.SetSingleCanaryRollout, {
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

CTFClient.prototype.setGlobalCanaryRollout = function setGlobalCanaryRollout(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.SetGlobalCanaryRollout, {
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

CTFClient.prototype.createCanary = function createCanary(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.CreateCanary, {
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

CTFClient.prototype.deleteCanary = function deleteCanary(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(CTF.DeleteCanary, {
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

