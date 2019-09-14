// package: com.roleypoly.discord
// file: discord.proto

var discord_pb = require("./discord_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Discord = (function () {
  function Discord() {}
  Discord.serviceName = "com.roleypoly.discord.Discord";
  return Discord;
}());

Discord.ListGuilds = {
  methodName: "ListGuilds",
  service: Discord,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
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

Discord.GetGuildRoles = {
  methodName: "GetGuildRoles",
  service: Discord,
  requestStream: false,
  responseStream: false,
  requestType: discord_pb.IDQuery,
  responseType: discord_pb.GuildRoles
};

Discord.GetGuildsByMember = {
  methodName: "GetGuildsByMember",
  service: Discord,
  requestStream: false,
  responseStream: false,
  requestType: discord_pb.IDQuery,
  responseType: discord_pb.GuildList
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

DiscordClient.prototype.listGuilds = function listGuilds(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Discord.ListGuilds, {
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

DiscordClient.prototype.getGuild = function getGuild(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Discord.GetGuild, {
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

DiscordClient.prototype.getGuildRoles = function getGuildRoles(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Discord.GetGuildRoles, {
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

DiscordClient.prototype.getGuildsByMember = function getGuildsByMember(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Discord.GetGuildsByMember, {
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

DiscordClient.prototype.getMember = function getMember(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Discord.GetMember, {
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

DiscordClient.prototype.updateMember = function updateMember(requestMessage, metadata) {
  return new Promise((resolve, reject) => {
    grpc.unary(Discord.UpdateMember, {
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

