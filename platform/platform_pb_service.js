// package: com.roleypoly.platform
// file: platform/platform.proto

var platform_platform_pb = require("../platform/platform_pb");
var discord_discord_pb = require("../discord/discord_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Platform = (function () {
  function Platform() {}
  Platform.serviceName = "com.roleypoly.platform.Platform";
  return Platform;
}());

Platform.GetGuild = {
  methodName: "GetGuild",
  service: Platform,
  requestStream: false,
  responseStream: false,
  requestType: discord_discord_pb.IDQuery,
  responseType: discord_discord_pb.Guild
};

exports.Platform = Platform;

function PlatformClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

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

exports.PlatformClient = PlatformClient;

