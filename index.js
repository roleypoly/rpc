"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Shared = exports.Platform = exports.Discord = exports.Ctf = exports.AuthClient = exports.Auth = void 0;
var empty_pb_1 = require("google-protobuf/google/protobuf/empty_pb");
Object.defineProperty(exports, "Empty", { enumerable: true, get: function () { return empty_pb_1.Empty; } });
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