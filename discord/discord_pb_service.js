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

Discord.GetGuild = {
  methodName: "GetGuild",
  service: Discord,
  requestStream: false,
  responseStream: false,
  requestType: discord_pb.IDQuery,
  responseType: discord_pb.Guild
};

Discord.GetMember = {
  methodName: "GetMember",
  service: Discord,
  requestStream: false,
  responseStream: false,
  requestType: discord_pb.IDQuery,
  responseType: discord_pb.Member
};

Discord.UpdateMember = {
  methodName: "UpdateMember",
  service: Discord,
  requestStream: false,
  responseStream: false,
  requestType: discord_pb.Member,
  responseType: discord_pb.Member
};

exports.Discord = Discord;

function DiscordClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

DiscordClient.prototype.listServers = function listServers(requestMessage, metadata) {
  let cancelled = false;
  const client = grpc.unary(Discord.ListServers, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (cancelled === false) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          Promise.reject(err);
        } else {
          Promise.resolve(response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      cancelled = true;
      Promise.resolve(null);
      client.close();
    }
  };
};

DiscordClient.prototype.getGuild = function getGuild(requestMessage, metadata) {
  let cancelled = false;
  const client = grpc.unary(Discord.GetGuild, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (cancelled === false) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          Promise.reject(err);
        } else {
          Promise.resolve(response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      cancelled = true;
      Promise.resolve(null);
      client.close();
    }
  };
};

DiscordClient.prototype.getMember = function getMember(requestMessage, metadata) {
  let cancelled = false;
  const client = grpc.unary(Discord.GetMember, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (cancelled === false) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          Promise.reject(err);
        } else {
          Promise.resolve(response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      cancelled = true;
      Promise.resolve(null);
      client.close();
    }
  };
};

DiscordClient.prototype.updateMember = function updateMember(requestMessage, metadata) {
  let cancelled = false;
  const client = grpc.unary(Discord.UpdateMember, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (cancelled === false) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          Promise.reject(err);
        } else {
          Promise.resolve(response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      cancelled = true;
      Promise.resolve(null);
      client.close();
    }
  };
};

exports.DiscordClient = DiscordClient;

