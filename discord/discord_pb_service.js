// package: 
// file: discord.proto

var discord_pb = require("./discord_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Discord = (function () {
  function Discord() {}
  Discord.serviceName = "Discord";
  return Discord;
}());

Discord.ListServers = {
  methodName: "ListServers",
  service: Discord,
  requestStream: false,
  responseStream: false,
  requestType: discord_pb.Empty,
  responseType: discord_pb.GuildList
};

exports.Discord = Discord;

function DiscordClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

DiscordClient.prototype.listServers = function listServers(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Discord.ListServers, {
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

exports.DiscordClient = DiscordClient;

