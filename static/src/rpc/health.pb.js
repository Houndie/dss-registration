/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars, strict, no-lone-blocks, default-case*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots.health || ($protobuf.roots.health = {});

$root.dss = (function() {

    /**
     * Namespace dss.
     * @exports dss
     * @namespace
     */
    var dss = {};

    dss.Info = (function() {

        /**
         * Constructs a new Info service.
         * @memberof dss
         * @classdesc Represents an Info
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Info(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Info.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Info;

        /**
         * Creates new Info service using the specified rpc implementation.
         * @function create
         * @memberof dss.Info
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {Info} RPC service. Useful where requests and/or responses are streamed.
         */
        Info.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link dss.Info#health}.
         * @memberof dss.Info
         * @typedef HealthCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.InfoHealthRes} [response] InfoHealthRes
         */

        /**
         * Calls Health.
         * @function health
         * @memberof dss.Info
         * @instance
         * @param {dss.IInfoHealthReq} request InfoHealthReq message or plain object
         * @param {dss.Info.HealthCallback} callback Node-style callback called with the error, if any, and InfoHealthRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Info.prototype.health = function health(request, callback) {
            return this.rpcCall(health, $root.dss.InfoHealthReq, $root.dss.InfoHealthRes, request, callback);
        }, "name", { value: "Health" });

        /**
         * Calls Health.
         * @function health
         * @memberof dss.Info
         * @instance
         * @param {dss.IInfoHealthReq} request InfoHealthReq message or plain object
         * @returns {Promise<dss.InfoHealthRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Info#version}.
         * @memberof dss.Info
         * @typedef VersionCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.InfoVersionRes} [response] InfoVersionRes
         */

        /**
         * Calls Version.
         * @function version
         * @memberof dss.Info
         * @instance
         * @param {dss.IInfoVersionReq} request InfoVersionReq message or plain object
         * @param {dss.Info.VersionCallback} callback Node-style callback called with the error, if any, and InfoVersionRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Info.prototype.version = function version(request, callback) {
            return this.rpcCall(version, $root.dss.InfoVersionReq, $root.dss.InfoVersionRes, request, callback);
        }, "name", { value: "Version" });

        /**
         * Calls Version.
         * @function version
         * @memberof dss.Info
         * @instance
         * @param {dss.IInfoVersionReq} request InfoVersionReq message or plain object
         * @returns {Promise<dss.InfoVersionRes>} Promise
         * @variation 2
         */

        return Info;
    })();

    dss.InfoHealthReq = (function() {

        /**
         * Properties of an InfoHealthReq.
         * @memberof dss
         * @interface IInfoHealthReq
         */

        /**
         * Constructs a new InfoHealthReq.
         * @memberof dss
         * @classdesc Represents an InfoHealthReq.
         * @implements IInfoHealthReq
         * @constructor
         * @param {dss.IInfoHealthReq=} [properties] Properties to set
         */
        function InfoHealthReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new InfoHealthReq instance using the specified properties.
         * @function create
         * @memberof dss.InfoHealthReq
         * @static
         * @param {dss.IInfoHealthReq=} [properties] Properties to set
         * @returns {dss.InfoHealthReq} InfoHealthReq instance
         */
        InfoHealthReq.create = function create(properties) {
            return new InfoHealthReq(properties);
        };

        /**
         * Encodes the specified InfoHealthReq message. Does not implicitly {@link dss.InfoHealthReq.verify|verify} messages.
         * @function encode
         * @memberof dss.InfoHealthReq
         * @static
         * @param {dss.IInfoHealthReq} message InfoHealthReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoHealthReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified InfoHealthReq message, length delimited. Does not implicitly {@link dss.InfoHealthReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.InfoHealthReq
         * @static
         * @param {dss.IInfoHealthReq} message InfoHealthReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoHealthReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an InfoHealthReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.InfoHealthReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.InfoHealthReq} InfoHealthReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoHealthReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.InfoHealthReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an InfoHealthReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.InfoHealthReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.InfoHealthReq} InfoHealthReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoHealthReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an InfoHealthReq message.
         * @function verify
         * @memberof dss.InfoHealthReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        InfoHealthReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates an InfoHealthReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.InfoHealthReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.InfoHealthReq} InfoHealthReq
         */
        InfoHealthReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.InfoHealthReq)
                return object;
            return new $root.dss.InfoHealthReq();
        };

        /**
         * Creates a plain object from an InfoHealthReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.InfoHealthReq
         * @static
         * @param {dss.InfoHealthReq} message InfoHealthReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InfoHealthReq.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this InfoHealthReq to JSON.
         * @function toJSON
         * @memberof dss.InfoHealthReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        InfoHealthReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return InfoHealthReq;
    })();

    dss.InfoHealthRes = (function() {

        /**
         * Properties of an InfoHealthRes.
         * @memberof dss
         * @interface IInfoHealthRes
         * @property {dss.InfoHealthRes.Healthiness|null} [healthiness] InfoHealthRes healthiness
         */

        /**
         * Constructs a new InfoHealthRes.
         * @memberof dss
         * @classdesc Represents an InfoHealthRes.
         * @implements IInfoHealthRes
         * @constructor
         * @param {dss.IInfoHealthRes=} [properties] Properties to set
         */
        function InfoHealthRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * InfoHealthRes healthiness.
         * @member {dss.InfoHealthRes.Healthiness} healthiness
         * @memberof dss.InfoHealthRes
         * @instance
         */
        InfoHealthRes.prototype.healthiness = 0;

        /**
         * Creates a new InfoHealthRes instance using the specified properties.
         * @function create
         * @memberof dss.InfoHealthRes
         * @static
         * @param {dss.IInfoHealthRes=} [properties] Properties to set
         * @returns {dss.InfoHealthRes} InfoHealthRes instance
         */
        InfoHealthRes.create = function create(properties) {
            return new InfoHealthRes(properties);
        };

        /**
         * Encodes the specified InfoHealthRes message. Does not implicitly {@link dss.InfoHealthRes.verify|verify} messages.
         * @function encode
         * @memberof dss.InfoHealthRes
         * @static
         * @param {dss.IInfoHealthRes} message InfoHealthRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoHealthRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.healthiness != null && Object.hasOwnProperty.call(message, "healthiness"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.healthiness);
            return writer;
        };

        /**
         * Encodes the specified InfoHealthRes message, length delimited. Does not implicitly {@link dss.InfoHealthRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.InfoHealthRes
         * @static
         * @param {dss.IInfoHealthRes} message InfoHealthRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoHealthRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an InfoHealthRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.InfoHealthRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.InfoHealthRes} InfoHealthRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoHealthRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.InfoHealthRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.healthiness = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an InfoHealthRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.InfoHealthRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.InfoHealthRes} InfoHealthRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoHealthRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an InfoHealthRes message.
         * @function verify
         * @memberof dss.InfoHealthRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        InfoHealthRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.healthiness != null && message.hasOwnProperty("healthiness"))
                switch (message.healthiness) {
                default:
                    return "healthiness: enum value expected";
                case 0:
                case 1:
                case 2:
                    break;
                }
            return null;
        };

        /**
         * Creates an InfoHealthRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.InfoHealthRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.InfoHealthRes} InfoHealthRes
         */
        InfoHealthRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.InfoHealthRes)
                return object;
            var message = new $root.dss.InfoHealthRes();
            switch (object.healthiness) {
            case "Unknown":
            case 0:
                message.healthiness = 0;
                break;
            case "Healthy":
            case 1:
                message.healthiness = 1;
                break;
            case "Unhealthy":
            case 2:
                message.healthiness = 2;
                break;
            }
            return message;
        };

        /**
         * Creates a plain object from an InfoHealthRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.InfoHealthRes
         * @static
         * @param {dss.InfoHealthRes} message InfoHealthRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InfoHealthRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.healthiness = options.enums === String ? "Unknown" : 0;
            if (message.healthiness != null && message.hasOwnProperty("healthiness"))
                object.healthiness = options.enums === String ? $root.dss.InfoHealthRes.Healthiness[message.healthiness] : message.healthiness;
            return object;
        };

        /**
         * Converts this InfoHealthRes to JSON.
         * @function toJSON
         * @memberof dss.InfoHealthRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        InfoHealthRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Healthiness enum.
         * @name dss.InfoHealthRes.Healthiness
         * @enum {number}
         * @property {number} Unknown=0 Unknown value
         * @property {number} Healthy=1 Healthy value
         * @property {number} Unhealthy=2 Unhealthy value
         */
        InfoHealthRes.Healthiness = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "Unknown"] = 0;
            values[valuesById[1] = "Healthy"] = 1;
            values[valuesById[2] = "Unhealthy"] = 2;
            return values;
        })();

        return InfoHealthRes;
    })();

    dss.InfoVersionReq = (function() {

        /**
         * Properties of an InfoVersionReq.
         * @memberof dss
         * @interface IInfoVersionReq
         */

        /**
         * Constructs a new InfoVersionReq.
         * @memberof dss
         * @classdesc Represents an InfoVersionReq.
         * @implements IInfoVersionReq
         * @constructor
         * @param {dss.IInfoVersionReq=} [properties] Properties to set
         */
        function InfoVersionReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new InfoVersionReq instance using the specified properties.
         * @function create
         * @memberof dss.InfoVersionReq
         * @static
         * @param {dss.IInfoVersionReq=} [properties] Properties to set
         * @returns {dss.InfoVersionReq} InfoVersionReq instance
         */
        InfoVersionReq.create = function create(properties) {
            return new InfoVersionReq(properties);
        };

        /**
         * Encodes the specified InfoVersionReq message. Does not implicitly {@link dss.InfoVersionReq.verify|verify} messages.
         * @function encode
         * @memberof dss.InfoVersionReq
         * @static
         * @param {dss.IInfoVersionReq} message InfoVersionReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoVersionReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified InfoVersionReq message, length delimited. Does not implicitly {@link dss.InfoVersionReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.InfoVersionReq
         * @static
         * @param {dss.IInfoVersionReq} message InfoVersionReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoVersionReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an InfoVersionReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.InfoVersionReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.InfoVersionReq} InfoVersionReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoVersionReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.InfoVersionReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an InfoVersionReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.InfoVersionReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.InfoVersionReq} InfoVersionReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoVersionReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an InfoVersionReq message.
         * @function verify
         * @memberof dss.InfoVersionReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        InfoVersionReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates an InfoVersionReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.InfoVersionReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.InfoVersionReq} InfoVersionReq
         */
        InfoVersionReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.InfoVersionReq)
                return object;
            return new $root.dss.InfoVersionReq();
        };

        /**
         * Creates a plain object from an InfoVersionReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.InfoVersionReq
         * @static
         * @param {dss.InfoVersionReq} message InfoVersionReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InfoVersionReq.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this InfoVersionReq to JSON.
         * @function toJSON
         * @memberof dss.InfoVersionReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        InfoVersionReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return InfoVersionReq;
    })();

    dss.InfoVersionRes = (function() {

        /**
         * Properties of an InfoVersionRes.
         * @memberof dss
         * @interface IInfoVersionRes
         * @property {string|null} [version] InfoVersionRes version
         */

        /**
         * Constructs a new InfoVersionRes.
         * @memberof dss
         * @classdesc Represents an InfoVersionRes.
         * @implements IInfoVersionRes
         * @constructor
         * @param {dss.IInfoVersionRes=} [properties] Properties to set
         */
        function InfoVersionRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * InfoVersionRes version.
         * @member {string} version
         * @memberof dss.InfoVersionRes
         * @instance
         */
        InfoVersionRes.prototype.version = "";

        /**
         * Creates a new InfoVersionRes instance using the specified properties.
         * @function create
         * @memberof dss.InfoVersionRes
         * @static
         * @param {dss.IInfoVersionRes=} [properties] Properties to set
         * @returns {dss.InfoVersionRes} InfoVersionRes instance
         */
        InfoVersionRes.create = function create(properties) {
            return new InfoVersionRes(properties);
        };

        /**
         * Encodes the specified InfoVersionRes message. Does not implicitly {@link dss.InfoVersionRes.verify|verify} messages.
         * @function encode
         * @memberof dss.InfoVersionRes
         * @static
         * @param {dss.IInfoVersionRes} message InfoVersionRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoVersionRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.version != null && Object.hasOwnProperty.call(message, "version"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.version);
            return writer;
        };

        /**
         * Encodes the specified InfoVersionRes message, length delimited. Does not implicitly {@link dss.InfoVersionRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.InfoVersionRes
         * @static
         * @param {dss.IInfoVersionRes} message InfoVersionRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InfoVersionRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an InfoVersionRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.InfoVersionRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.InfoVersionRes} InfoVersionRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoVersionRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.InfoVersionRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.version = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an InfoVersionRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.InfoVersionRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.InfoVersionRes} InfoVersionRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InfoVersionRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an InfoVersionRes message.
         * @function verify
         * @memberof dss.InfoVersionRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        InfoVersionRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.version != null && message.hasOwnProperty("version"))
                if (!$util.isString(message.version))
                    return "version: string expected";
            return null;
        };

        /**
         * Creates an InfoVersionRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.InfoVersionRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.InfoVersionRes} InfoVersionRes
         */
        InfoVersionRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.InfoVersionRes)
                return object;
            var message = new $root.dss.InfoVersionRes();
            if (object.version != null)
                message.version = String(object.version);
            return message;
        };

        /**
         * Creates a plain object from an InfoVersionRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.InfoVersionRes
         * @static
         * @param {dss.InfoVersionRes} message InfoVersionRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InfoVersionRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.version = "";
            if (message.version != null && message.hasOwnProperty("version"))
                object.version = message.version;
            return object;
        };

        /**
         * Converts this InfoVersionRes to JSON.
         * @function toJSON
         * @memberof dss.InfoVersionRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        InfoVersionRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return InfoVersionRes;
    })();

    return dss;
})();

module.exports = $root;
