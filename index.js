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
const AuthClient = __importStar(require("./auth/client"));
exports.AuthClient = AuthClient;
const Discord = __importStar(require("./discord"));
exports.Discord = Discord;
const Platform = __importStar(require("./platform"));
exports.Platform = Platform;
//# sourceMappingURL=index.js.map