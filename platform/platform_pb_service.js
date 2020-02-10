// package: roleypoly.platform
// file: platform/platform.proto

var platform_platform_pb = require("../platform/platform_pb");
var shared_shared_pb = require("../shared/shared_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Platform = (function () {
  function Platform() {}
  Platform.serviceName = "roleypoly.platform.Platform";
  return Platform;
}());

Platform.EnumerateMyGuilds = {
  methodName: "EnumerateMyGuilds",
  service: Platform,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: platform_platform_pb.GuildEnumeration
};

Platform.GetGuildSlug = {
  methodName: "GetGuildSlug",
  service: Platform,
  requestStream: false,
  responseStream: false,
  requestType: shared_shared_pb.IDQuery,
  responseType: shared_shared_pb.Guild
};

Platform.GetGuild = {
  methodName: "GetGuild",
  service: Platform,
  requestStream: false,
  responseStream: false,
  requestType: shared_shared_pb.IDQuery,
  responseType: platform_platform_pb.PresentableGuild
};

Platform.UpdateMyRoles = {
  methodName: "UpdateMyRoles",
  service: Platform,
  requestStream: false,
  responseStream: false,
  requestType: platform_platform_pb.UpdateRoles,
  responseType: google_protobuf_empty_pb.Empty
};

Platform.UpdateGuildData = {
  methodName: "UpdateGuildData",
  service: Platform,
  requestStream: false,
  responseStream: false,
  requestType: platform_platform_pb.GuildData,
  responseType: google_protobuf_empty_pb.Empty
};

exports.Platform = Platform;

function PlatformClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PlatformClient.prototype.enumerateMyGuilds = function enumerateMyGuilds(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Platform.EnumerateMyGuilds, {
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

PlatformClient.prototype.getGuildSlug = function getGuildSlug(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Platform.GetGuildSlug, {
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

PlatformClient.prototype.getGuild = function getGuild(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Platform.GetGuild, {
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

PlatformClient.prototype.updateMyRoles = function updateMyRoles(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Platform.UpdateMyRoles, {
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

PlatformClient.prototype.updateGuildData = function updateGuildData(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Platform.UpdateGuildData, {
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

exports.PlatformClient = PlatformClient;

