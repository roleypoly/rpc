"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
var empty_pb_1 = require("google-protobuf/google/protobuf/empty_pb");
exports.Empty = empty_pb_1.Empty;
const Auth = __importStar(require("./auth"));
exports.Auth = Auth;
const AuthClient = __importStar(require("./auth/client"));
exports.AuthClient = AuthClient;
const Ctf = __importStar(require("./ctf"));
exports.Ctf = Ctf;
const Discord = __importStar(require("./discord"));
exports.Discord = Discord;
const Platform = __importStar(require("./platform"));
exports.Platform = Platform;
const Shared = __importStar(require("./shared"));
exports.Shared = Shared;
//# sourceMappingURL=index.js.map