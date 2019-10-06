// package: com.roleypoly.platform
// file: platform.proto

var platform_pb = require("./platform_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Platform = (function () {
  function Platform() {}
  Platform.serviceName = "com.roleypoly.platform.Platform";
  return Platform;
}());

exports.Platform = Platform;

function PlatformClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

exports.PlatformClient = PlatformClient;

