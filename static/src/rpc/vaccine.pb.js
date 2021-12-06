/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars, strict, no-lone-blocks, default-case*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots.vaccine || ($protobuf.roots.vaccine = {});

$root.dss = (function() {

    /**
     * Namespace dss.
     * @exports dss
     * @namespace
     */
    var dss = {};

    dss.Vaccine = (function() {

        /**
         * Constructs a new Vaccine service.
         * @memberof dss
         * @classdesc Represents a Vaccine
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Vaccine(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Vaccine.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Vaccine;

        /**
         * Creates new Vaccine service using the specified rpc implementation.
         * @function create
         * @memberof dss.Vaccine
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {Vaccine} RPC service. Useful where requests and/or responses are streamed.
         */
        Vaccine.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link dss.Vaccine#upload}.
         * @memberof dss.Vaccine
         * @typedef UploadCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.VaccineUploadRes} [response] VaccineUploadRes
         */

        /**
         * Calls Upload.
         * @function upload
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineUploadReq} request VaccineUploadReq message or plain object
         * @param {dss.Vaccine.UploadCallback} callback Node-style callback called with the error, if any, and VaccineUploadRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Vaccine.prototype.upload = function upload(request, callback) {
            return this.rpcCall(upload, $root.dss.VaccineUploadReq, $root.dss.VaccineUploadRes, request, callback);
        }, "name", { value: "Upload" });

        /**
         * Calls Upload.
         * @function upload
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineUploadReq} request VaccineUploadReq message or plain object
         * @returns {Promise<dss.VaccineUploadRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Vaccine#get}.
         * @memberof dss.Vaccine
         * @typedef GetCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.VaccineGetRes} [response] VaccineGetRes
         */

        /**
         * Calls Get.
         * @function get
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineGetReq} request VaccineGetReq message or plain object
         * @param {dss.Vaccine.GetCallback} callback Node-style callback called with the error, if any, and VaccineGetRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Vaccine.prototype.get = function get(request, callback) {
            return this.rpcCall(get, $root.dss.VaccineGetReq, $root.dss.VaccineGetRes, request, callback);
        }, "name", { value: "Get" });

        /**
         * Calls Get.
         * @function get
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineGetReq} request VaccineGetReq message or plain object
         * @returns {Promise<dss.VaccineGetRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Vaccine#approve}.
         * @memberof dss.Vaccine
         * @typedef ApproveCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.VaccineApproveRes} [response] VaccineApproveRes
         */

        /**
         * Calls Approve.
         * @function approve
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineApproveReq} request VaccineApproveReq message or plain object
         * @param {dss.Vaccine.ApproveCallback} callback Node-style callback called with the error, if any, and VaccineApproveRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Vaccine.prototype.approve = function approve(request, callback) {
            return this.rpcCall(approve, $root.dss.VaccineApproveReq, $root.dss.VaccineApproveRes, request, callback);
        }, "name", { value: "Approve" });

        /**
         * Calls Approve.
         * @function approve
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineApproveReq} request VaccineApproveReq message or plain object
         * @returns {Promise<dss.VaccineApproveRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Vaccine#reject}.
         * @memberof dss.Vaccine
         * @typedef RejectCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.VaccineRejectRes} [response] VaccineRejectRes
         */

        /**
         * Calls Reject.
         * @function reject
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineRejectReq} request VaccineRejectReq message or plain object
         * @param {dss.Vaccine.RejectCallback} callback Node-style callback called with the error, if any, and VaccineRejectRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Vaccine.prototype.reject = function reject(request, callback) {
            return this.rpcCall(reject, $root.dss.VaccineRejectReq, $root.dss.VaccineRejectRes, request, callback);
        }, "name", { value: "Reject" });

        /**
         * Calls Reject.
         * @function reject
         * @memberof dss.Vaccine
         * @instance
         * @param {dss.IVaccineRejectReq} request VaccineRejectReq message or plain object
         * @returns {Promise<dss.VaccineRejectRes>} Promise
         * @variation 2
         */

        return Vaccine;
    })();

    dss.VaxApproved = (function() {

        /**
         * Properties of a VaxApproved.
         * @memberof dss
         * @interface IVaxApproved
         */

        /**
         * Constructs a new VaxApproved.
         * @memberof dss
         * @classdesc Represents a VaxApproved.
         * @implements IVaxApproved
         * @constructor
         * @param {dss.IVaxApproved=} [properties] Properties to set
         */
        function VaxApproved(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new VaxApproved instance using the specified properties.
         * @function create
         * @memberof dss.VaxApproved
         * @static
         * @param {dss.IVaxApproved=} [properties] Properties to set
         * @returns {dss.VaxApproved} VaxApproved instance
         */
        VaxApproved.create = function create(properties) {
            return new VaxApproved(properties);
        };

        /**
         * Encodes the specified VaxApproved message. Does not implicitly {@link dss.VaxApproved.verify|verify} messages.
         * @function encode
         * @memberof dss.VaxApproved
         * @static
         * @param {dss.IVaxApproved} message VaxApproved message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaxApproved.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified VaxApproved message, length delimited. Does not implicitly {@link dss.VaxApproved.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaxApproved
         * @static
         * @param {dss.IVaxApproved} message VaxApproved message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaxApproved.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaxApproved message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaxApproved
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaxApproved} VaxApproved
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaxApproved.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaxApproved();
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
         * Decodes a VaxApproved message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaxApproved
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaxApproved} VaxApproved
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaxApproved.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaxApproved message.
         * @function verify
         * @memberof dss.VaxApproved
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaxApproved.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a VaxApproved message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaxApproved
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaxApproved} VaxApproved
         */
        VaxApproved.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaxApproved)
                return object;
            return new $root.dss.VaxApproved();
        };

        /**
         * Creates a plain object from a VaxApproved message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaxApproved
         * @static
         * @param {dss.VaxApproved} message VaxApproved
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaxApproved.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this VaxApproved to JSON.
         * @function toJSON
         * @memberof dss.VaxApproved
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaxApproved.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaxApproved;
    })();

    dss.VaxApprovalPending = (function() {

        /**
         * Properties of a VaxApprovalPending.
         * @memberof dss
         * @interface IVaxApprovalPending
         * @property {string|null} [url] VaxApprovalPending url
         */

        /**
         * Constructs a new VaxApprovalPending.
         * @memberof dss
         * @classdesc Represents a VaxApprovalPending.
         * @implements IVaxApprovalPending
         * @constructor
         * @param {dss.IVaxApprovalPending=} [properties] Properties to set
         */
        function VaxApprovalPending(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * VaxApprovalPending url.
         * @member {string} url
         * @memberof dss.VaxApprovalPending
         * @instance
         */
        VaxApprovalPending.prototype.url = "";

        /**
         * Creates a new VaxApprovalPending instance using the specified properties.
         * @function create
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {dss.IVaxApprovalPending=} [properties] Properties to set
         * @returns {dss.VaxApprovalPending} VaxApprovalPending instance
         */
        VaxApprovalPending.create = function create(properties) {
            return new VaxApprovalPending(properties);
        };

        /**
         * Encodes the specified VaxApprovalPending message. Does not implicitly {@link dss.VaxApprovalPending.verify|verify} messages.
         * @function encode
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {dss.IVaxApprovalPending} message VaxApprovalPending message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaxApprovalPending.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.url != null && Object.hasOwnProperty.call(message, "url"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.url);
            return writer;
        };

        /**
         * Encodes the specified VaxApprovalPending message, length delimited. Does not implicitly {@link dss.VaxApprovalPending.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {dss.IVaxApprovalPending} message VaxApprovalPending message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaxApprovalPending.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaxApprovalPending message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaxApprovalPending} VaxApprovalPending
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaxApprovalPending.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaxApprovalPending();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.url = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a VaxApprovalPending message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaxApprovalPending} VaxApprovalPending
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaxApprovalPending.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaxApprovalPending message.
         * @function verify
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaxApprovalPending.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.url != null && message.hasOwnProperty("url"))
                if (!$util.isString(message.url))
                    return "url: string expected";
            return null;
        };

        /**
         * Creates a VaxApprovalPending message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaxApprovalPending} VaxApprovalPending
         */
        VaxApprovalPending.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaxApprovalPending)
                return object;
            var message = new $root.dss.VaxApprovalPending();
            if (object.url != null)
                message.url = String(object.url);
            return message;
        };

        /**
         * Creates a plain object from a VaxApprovalPending message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaxApprovalPending
         * @static
         * @param {dss.VaxApprovalPending} message VaxApprovalPending
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaxApprovalPending.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.url = "";
            if (message.url != null && message.hasOwnProperty("url"))
                object.url = message.url;
            return object;
        };

        /**
         * Converts this VaxApprovalPending to JSON.
         * @function toJSON
         * @memberof dss.VaxApprovalPending
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaxApprovalPending.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaxApprovalPending;
    })();

    dss.NoVaxProofSupplied = (function() {

        /**
         * Properties of a NoVaxProofSupplied.
         * @memberof dss
         * @interface INoVaxProofSupplied
         */

        /**
         * Constructs a new NoVaxProofSupplied.
         * @memberof dss
         * @classdesc Represents a NoVaxProofSupplied.
         * @implements INoVaxProofSupplied
         * @constructor
         * @param {dss.INoVaxProofSupplied=} [properties] Properties to set
         */
        function NoVaxProofSupplied(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new NoVaxProofSupplied instance using the specified properties.
         * @function create
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {dss.INoVaxProofSupplied=} [properties] Properties to set
         * @returns {dss.NoVaxProofSupplied} NoVaxProofSupplied instance
         */
        NoVaxProofSupplied.create = function create(properties) {
            return new NoVaxProofSupplied(properties);
        };

        /**
         * Encodes the specified NoVaxProofSupplied message. Does not implicitly {@link dss.NoVaxProofSupplied.verify|verify} messages.
         * @function encode
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {dss.INoVaxProofSupplied} message NoVaxProofSupplied message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        NoVaxProofSupplied.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified NoVaxProofSupplied message, length delimited. Does not implicitly {@link dss.NoVaxProofSupplied.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {dss.INoVaxProofSupplied} message NoVaxProofSupplied message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        NoVaxProofSupplied.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a NoVaxProofSupplied message from the specified reader or buffer.
         * @function decode
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.NoVaxProofSupplied} NoVaxProofSupplied
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        NoVaxProofSupplied.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.NoVaxProofSupplied();
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
         * Decodes a NoVaxProofSupplied message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.NoVaxProofSupplied} NoVaxProofSupplied
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        NoVaxProofSupplied.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a NoVaxProofSupplied message.
         * @function verify
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        NoVaxProofSupplied.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a NoVaxProofSupplied message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.NoVaxProofSupplied} NoVaxProofSupplied
         */
        NoVaxProofSupplied.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.NoVaxProofSupplied)
                return object;
            return new $root.dss.NoVaxProofSupplied();
        };

        /**
         * Creates a plain object from a NoVaxProofSupplied message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.NoVaxProofSupplied
         * @static
         * @param {dss.NoVaxProofSupplied} message NoVaxProofSupplied
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        NoVaxProofSupplied.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this NoVaxProofSupplied to JSON.
         * @function toJSON
         * @memberof dss.NoVaxProofSupplied
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        NoVaxProofSupplied.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return NoVaxProofSupplied;
    })();

    dss.VaccineUploadReq = (function() {

        /**
         * Properties of a VaccineUploadReq.
         * @memberof dss
         * @interface IVaccineUploadReq
         * @property {string|null} [id] VaccineUploadReq id
         * @property {number|Long|null} [filesize] VaccineUploadReq filesize
         */

        /**
         * Constructs a new VaccineUploadReq.
         * @memberof dss
         * @classdesc Represents a VaccineUploadReq.
         * @implements IVaccineUploadReq
         * @constructor
         * @param {dss.IVaccineUploadReq=} [properties] Properties to set
         */
        function VaccineUploadReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * VaccineUploadReq id.
         * @member {string} id
         * @memberof dss.VaccineUploadReq
         * @instance
         */
        VaccineUploadReq.prototype.id = "";

        /**
         * VaccineUploadReq filesize.
         * @member {number|Long} filesize
         * @memberof dss.VaccineUploadReq
         * @instance
         */
        VaccineUploadReq.prototype.filesize = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Creates a new VaccineUploadReq instance using the specified properties.
         * @function create
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {dss.IVaccineUploadReq=} [properties] Properties to set
         * @returns {dss.VaccineUploadReq} VaccineUploadReq instance
         */
        VaccineUploadReq.create = function create(properties) {
            return new VaccineUploadReq(properties);
        };

        /**
         * Encodes the specified VaccineUploadReq message. Does not implicitly {@link dss.VaccineUploadReq.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {dss.IVaccineUploadReq} message VaccineUploadReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineUploadReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            if (message.filesize != null && Object.hasOwnProperty.call(message, "filesize"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.filesize);
            return writer;
        };

        /**
         * Encodes the specified VaccineUploadReq message, length delimited. Does not implicitly {@link dss.VaccineUploadReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {dss.IVaccineUploadReq} message VaccineUploadReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineUploadReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineUploadReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineUploadReq} VaccineUploadReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineUploadReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineUploadReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.filesize = reader.int64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a VaccineUploadReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineUploadReq} VaccineUploadReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineUploadReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineUploadReq message.
         * @function verify
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineUploadReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            if (message.filesize != null && message.hasOwnProperty("filesize"))
                if (!$util.isInteger(message.filesize) && !(message.filesize && $util.isInteger(message.filesize.low) && $util.isInteger(message.filesize.high)))
                    return "filesize: integer|Long expected";
            return null;
        };

        /**
         * Creates a VaccineUploadReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineUploadReq} VaccineUploadReq
         */
        VaccineUploadReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineUploadReq)
                return object;
            var message = new $root.dss.VaccineUploadReq();
            if (object.id != null)
                message.id = String(object.id);
            if (object.filesize != null)
                if ($util.Long)
                    (message.filesize = $util.Long.fromValue(object.filesize)).unsigned = false;
                else if (typeof object.filesize === "string")
                    message.filesize = parseInt(object.filesize, 10);
                else if (typeof object.filesize === "number")
                    message.filesize = object.filesize;
                else if (typeof object.filesize === "object")
                    message.filesize = new $util.LongBits(object.filesize.low >>> 0, object.filesize.high >>> 0).toNumber();
            return message;
        };

        /**
         * Creates a plain object from a VaccineUploadReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineUploadReq
         * @static
         * @param {dss.VaccineUploadReq} message VaccineUploadReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineUploadReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.id = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.filesize = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.filesize = options.longs === String ? "0" : 0;
            }
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.filesize != null && message.hasOwnProperty("filesize"))
                if (typeof message.filesize === "number")
                    object.filesize = options.longs === String ? String(message.filesize) : message.filesize;
                else
                    object.filesize = options.longs === String ? $util.Long.prototype.toString.call(message.filesize) : options.longs === Number ? new $util.LongBits(message.filesize.low >>> 0, message.filesize.high >>> 0).toNumber() : message.filesize;
            return object;
        };

        /**
         * Converts this VaccineUploadReq to JSON.
         * @function toJSON
         * @memberof dss.VaccineUploadReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineUploadReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineUploadReq;
    })();

    dss.VaccineUploadRes = (function() {

        /**
         * Properties of a VaccineUploadRes.
         * @memberof dss
         * @interface IVaccineUploadRes
         * @property {string|null} [url] VaccineUploadRes url
         */

        /**
         * Constructs a new VaccineUploadRes.
         * @memberof dss
         * @classdesc Represents a VaccineUploadRes.
         * @implements IVaccineUploadRes
         * @constructor
         * @param {dss.IVaccineUploadRes=} [properties] Properties to set
         */
        function VaccineUploadRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * VaccineUploadRes url.
         * @member {string} url
         * @memberof dss.VaccineUploadRes
         * @instance
         */
        VaccineUploadRes.prototype.url = "";

        /**
         * Creates a new VaccineUploadRes instance using the specified properties.
         * @function create
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {dss.IVaccineUploadRes=} [properties] Properties to set
         * @returns {dss.VaccineUploadRes} VaccineUploadRes instance
         */
        VaccineUploadRes.create = function create(properties) {
            return new VaccineUploadRes(properties);
        };

        /**
         * Encodes the specified VaccineUploadRes message. Does not implicitly {@link dss.VaccineUploadRes.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {dss.IVaccineUploadRes} message VaccineUploadRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineUploadRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.url != null && Object.hasOwnProperty.call(message, "url"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.url);
            return writer;
        };

        /**
         * Encodes the specified VaccineUploadRes message, length delimited. Does not implicitly {@link dss.VaccineUploadRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {dss.IVaccineUploadRes} message VaccineUploadRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineUploadRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineUploadRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineUploadRes} VaccineUploadRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineUploadRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineUploadRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.url = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a VaccineUploadRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineUploadRes} VaccineUploadRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineUploadRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineUploadRes message.
         * @function verify
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineUploadRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.url != null && message.hasOwnProperty("url"))
                if (!$util.isString(message.url))
                    return "url: string expected";
            return null;
        };

        /**
         * Creates a VaccineUploadRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineUploadRes} VaccineUploadRes
         */
        VaccineUploadRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineUploadRes)
                return object;
            var message = new $root.dss.VaccineUploadRes();
            if (object.url != null)
                message.url = String(object.url);
            return message;
        };

        /**
         * Creates a plain object from a VaccineUploadRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineUploadRes
         * @static
         * @param {dss.VaccineUploadRes} message VaccineUploadRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineUploadRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.url = "";
            if (message.url != null && message.hasOwnProperty("url"))
                object.url = message.url;
            return object;
        };

        /**
         * Converts this VaccineUploadRes to JSON.
         * @function toJSON
         * @memberof dss.VaccineUploadRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineUploadRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineUploadRes;
    })();

    dss.VaccineGetReq = (function() {

        /**
         * Properties of a VaccineGetReq.
         * @memberof dss
         * @interface IVaccineGetReq
         * @property {string|null} [id] VaccineGetReq id
         */

        /**
         * Constructs a new VaccineGetReq.
         * @memberof dss
         * @classdesc Represents a VaccineGetReq.
         * @implements IVaccineGetReq
         * @constructor
         * @param {dss.IVaccineGetReq=} [properties] Properties to set
         */
        function VaccineGetReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * VaccineGetReq id.
         * @member {string} id
         * @memberof dss.VaccineGetReq
         * @instance
         */
        VaccineGetReq.prototype.id = "";

        /**
         * Creates a new VaccineGetReq instance using the specified properties.
         * @function create
         * @memberof dss.VaccineGetReq
         * @static
         * @param {dss.IVaccineGetReq=} [properties] Properties to set
         * @returns {dss.VaccineGetReq} VaccineGetReq instance
         */
        VaccineGetReq.create = function create(properties) {
            return new VaccineGetReq(properties);
        };

        /**
         * Encodes the specified VaccineGetReq message. Does not implicitly {@link dss.VaccineGetReq.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineGetReq
         * @static
         * @param {dss.IVaccineGetReq} message VaccineGetReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineGetReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            return writer;
        };

        /**
         * Encodes the specified VaccineGetReq message, length delimited. Does not implicitly {@link dss.VaccineGetReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineGetReq
         * @static
         * @param {dss.IVaccineGetReq} message VaccineGetReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineGetReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineGetReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineGetReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineGetReq} VaccineGetReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineGetReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineGetReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a VaccineGetReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineGetReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineGetReq} VaccineGetReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineGetReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineGetReq message.
         * @function verify
         * @memberof dss.VaccineGetReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineGetReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            return null;
        };

        /**
         * Creates a VaccineGetReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineGetReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineGetReq} VaccineGetReq
         */
        VaccineGetReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineGetReq)
                return object;
            var message = new $root.dss.VaccineGetReq();
            if (object.id != null)
                message.id = String(object.id);
            return message;
        };

        /**
         * Creates a plain object from a VaccineGetReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineGetReq
         * @static
         * @param {dss.VaccineGetReq} message VaccineGetReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineGetReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.id = "";
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            return object;
        };

        /**
         * Converts this VaccineGetReq to JSON.
         * @function toJSON
         * @memberof dss.VaccineGetReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineGetReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineGetReq;
    })();

    dss.VaccineGetRes = (function() {

        /**
         * Properties of a VaccineGetRes.
         * @memberof dss
         * @interface IVaccineGetRes
         * @property {dss.IVaxApproved|null} [vaxApproved] VaccineGetRes vaxApproved
         * @property {dss.IVaxApprovalPending|null} [vaxApprovalPending] VaccineGetRes vaxApprovalPending
         * @property {dss.INoVaxProofSupplied|null} [noVaxProofSupplied] VaccineGetRes noVaxProofSupplied
         */

        /**
         * Constructs a new VaccineGetRes.
         * @memberof dss
         * @classdesc Represents a VaccineGetRes.
         * @implements IVaccineGetRes
         * @constructor
         * @param {dss.IVaccineGetRes=} [properties] Properties to set
         */
        function VaccineGetRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * VaccineGetRes vaxApproved.
         * @member {dss.IVaxApproved|null|undefined} vaxApproved
         * @memberof dss.VaccineGetRes
         * @instance
         */
        VaccineGetRes.prototype.vaxApproved = null;

        /**
         * VaccineGetRes vaxApprovalPending.
         * @member {dss.IVaxApprovalPending|null|undefined} vaxApprovalPending
         * @memberof dss.VaccineGetRes
         * @instance
         */
        VaccineGetRes.prototype.vaxApprovalPending = null;

        /**
         * VaccineGetRes noVaxProofSupplied.
         * @member {dss.INoVaxProofSupplied|null|undefined} noVaxProofSupplied
         * @memberof dss.VaccineGetRes
         * @instance
         */
        VaccineGetRes.prototype.noVaxProofSupplied = null;

        // OneOf field names bound to virtual getters and setters
        var $oneOfFields;

        /**
         * VaccineGetRes info.
         * @member {"vaxApproved"|"vaxApprovalPending"|"noVaxProofSupplied"|undefined} info
         * @memberof dss.VaccineGetRes
         * @instance
         */
        Object.defineProperty(VaccineGetRes.prototype, "info", {
            get: $util.oneOfGetter($oneOfFields = ["vaxApproved", "vaxApprovalPending", "noVaxProofSupplied"]),
            set: $util.oneOfSetter($oneOfFields)
        });

        /**
         * Creates a new VaccineGetRes instance using the specified properties.
         * @function create
         * @memberof dss.VaccineGetRes
         * @static
         * @param {dss.IVaccineGetRes=} [properties] Properties to set
         * @returns {dss.VaccineGetRes} VaccineGetRes instance
         */
        VaccineGetRes.create = function create(properties) {
            return new VaccineGetRes(properties);
        };

        /**
         * Encodes the specified VaccineGetRes message. Does not implicitly {@link dss.VaccineGetRes.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineGetRes
         * @static
         * @param {dss.IVaccineGetRes} message VaccineGetRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineGetRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.vaxApproved != null && Object.hasOwnProperty.call(message, "vaxApproved"))
                $root.dss.VaxApproved.encode(message.vaxApproved, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.vaxApprovalPending != null && Object.hasOwnProperty.call(message, "vaxApprovalPending"))
                $root.dss.VaxApprovalPending.encode(message.vaxApprovalPending, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.noVaxProofSupplied != null && Object.hasOwnProperty.call(message, "noVaxProofSupplied"))
                $root.dss.NoVaxProofSupplied.encode(message.noVaxProofSupplied, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified VaccineGetRes message, length delimited. Does not implicitly {@link dss.VaccineGetRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineGetRes
         * @static
         * @param {dss.IVaccineGetRes} message VaccineGetRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineGetRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineGetRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineGetRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineGetRes} VaccineGetRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineGetRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineGetRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.vaxApproved = $root.dss.VaxApproved.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.vaxApprovalPending = $root.dss.VaxApprovalPending.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.noVaxProofSupplied = $root.dss.NoVaxProofSupplied.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a VaccineGetRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineGetRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineGetRes} VaccineGetRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineGetRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineGetRes message.
         * @function verify
         * @memberof dss.VaccineGetRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineGetRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            var properties = {};
            if (message.vaxApproved != null && message.hasOwnProperty("vaxApproved")) {
                properties.info = 1;
                {
                    var error = $root.dss.VaxApproved.verify(message.vaxApproved);
                    if (error)
                        return "vaxApproved." + error;
                }
            }
            if (message.vaxApprovalPending != null && message.hasOwnProperty("vaxApprovalPending")) {
                if (properties.info === 1)
                    return "info: multiple values";
                properties.info = 1;
                {
                    var error = $root.dss.VaxApprovalPending.verify(message.vaxApprovalPending);
                    if (error)
                        return "vaxApprovalPending." + error;
                }
            }
            if (message.noVaxProofSupplied != null && message.hasOwnProperty("noVaxProofSupplied")) {
                if (properties.info === 1)
                    return "info: multiple values";
                properties.info = 1;
                {
                    var error = $root.dss.NoVaxProofSupplied.verify(message.noVaxProofSupplied);
                    if (error)
                        return "noVaxProofSupplied." + error;
                }
            }
            return null;
        };

        /**
         * Creates a VaccineGetRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineGetRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineGetRes} VaccineGetRes
         */
        VaccineGetRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineGetRes)
                return object;
            var message = new $root.dss.VaccineGetRes();
            if (object.vaxApproved != null) {
                if (typeof object.vaxApproved !== "object")
                    throw TypeError(".dss.VaccineGetRes.vaxApproved: object expected");
                message.vaxApproved = $root.dss.VaxApproved.fromObject(object.vaxApproved);
            }
            if (object.vaxApprovalPending != null) {
                if (typeof object.vaxApprovalPending !== "object")
                    throw TypeError(".dss.VaccineGetRes.vaxApprovalPending: object expected");
                message.vaxApprovalPending = $root.dss.VaxApprovalPending.fromObject(object.vaxApprovalPending);
            }
            if (object.noVaxProofSupplied != null) {
                if (typeof object.noVaxProofSupplied !== "object")
                    throw TypeError(".dss.VaccineGetRes.noVaxProofSupplied: object expected");
                message.noVaxProofSupplied = $root.dss.NoVaxProofSupplied.fromObject(object.noVaxProofSupplied);
            }
            return message;
        };

        /**
         * Creates a plain object from a VaccineGetRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineGetRes
         * @static
         * @param {dss.VaccineGetRes} message VaccineGetRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineGetRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (message.vaxApproved != null && message.hasOwnProperty("vaxApproved")) {
                object.vaxApproved = $root.dss.VaxApproved.toObject(message.vaxApproved, options);
                if (options.oneofs)
                    object.info = "vaxApproved";
            }
            if (message.vaxApprovalPending != null && message.hasOwnProperty("vaxApprovalPending")) {
                object.vaxApprovalPending = $root.dss.VaxApprovalPending.toObject(message.vaxApprovalPending, options);
                if (options.oneofs)
                    object.info = "vaxApprovalPending";
            }
            if (message.noVaxProofSupplied != null && message.hasOwnProperty("noVaxProofSupplied")) {
                object.noVaxProofSupplied = $root.dss.NoVaxProofSupplied.toObject(message.noVaxProofSupplied, options);
                if (options.oneofs)
                    object.info = "noVaxProofSupplied";
            }
            return object;
        };

        /**
         * Converts this VaccineGetRes to JSON.
         * @function toJSON
         * @memberof dss.VaccineGetRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineGetRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineGetRes;
    })();

    dss.VaccineApproveReq = (function() {

        /**
         * Properties of a VaccineApproveReq.
         * @memberof dss
         * @interface IVaccineApproveReq
         * @property {string|null} [id] VaccineApproveReq id
         */

        /**
         * Constructs a new VaccineApproveReq.
         * @memberof dss
         * @classdesc Represents a VaccineApproveReq.
         * @implements IVaccineApproveReq
         * @constructor
         * @param {dss.IVaccineApproveReq=} [properties] Properties to set
         */
        function VaccineApproveReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * VaccineApproveReq id.
         * @member {string} id
         * @memberof dss.VaccineApproveReq
         * @instance
         */
        VaccineApproveReq.prototype.id = "";

        /**
         * Creates a new VaccineApproveReq instance using the specified properties.
         * @function create
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {dss.IVaccineApproveReq=} [properties] Properties to set
         * @returns {dss.VaccineApproveReq} VaccineApproveReq instance
         */
        VaccineApproveReq.create = function create(properties) {
            return new VaccineApproveReq(properties);
        };

        /**
         * Encodes the specified VaccineApproveReq message. Does not implicitly {@link dss.VaccineApproveReq.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {dss.IVaccineApproveReq} message VaccineApproveReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineApproveReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            return writer;
        };

        /**
         * Encodes the specified VaccineApproveReq message, length delimited. Does not implicitly {@link dss.VaccineApproveReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {dss.IVaccineApproveReq} message VaccineApproveReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineApproveReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineApproveReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineApproveReq} VaccineApproveReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineApproveReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineApproveReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a VaccineApproveReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineApproveReq} VaccineApproveReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineApproveReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineApproveReq message.
         * @function verify
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineApproveReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            return null;
        };

        /**
         * Creates a VaccineApproveReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineApproveReq} VaccineApproveReq
         */
        VaccineApproveReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineApproveReq)
                return object;
            var message = new $root.dss.VaccineApproveReq();
            if (object.id != null)
                message.id = String(object.id);
            return message;
        };

        /**
         * Creates a plain object from a VaccineApproveReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineApproveReq
         * @static
         * @param {dss.VaccineApproveReq} message VaccineApproveReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineApproveReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.id = "";
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            return object;
        };

        /**
         * Converts this VaccineApproveReq to JSON.
         * @function toJSON
         * @memberof dss.VaccineApproveReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineApproveReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineApproveReq;
    })();

    dss.VaccineApproveRes = (function() {

        /**
         * Properties of a VaccineApproveRes.
         * @memberof dss
         * @interface IVaccineApproveRes
         */

        /**
         * Constructs a new VaccineApproveRes.
         * @memberof dss
         * @classdesc Represents a VaccineApproveRes.
         * @implements IVaccineApproveRes
         * @constructor
         * @param {dss.IVaccineApproveRes=} [properties] Properties to set
         */
        function VaccineApproveRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new VaccineApproveRes instance using the specified properties.
         * @function create
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {dss.IVaccineApproveRes=} [properties] Properties to set
         * @returns {dss.VaccineApproveRes} VaccineApproveRes instance
         */
        VaccineApproveRes.create = function create(properties) {
            return new VaccineApproveRes(properties);
        };

        /**
         * Encodes the specified VaccineApproveRes message. Does not implicitly {@link dss.VaccineApproveRes.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {dss.IVaccineApproveRes} message VaccineApproveRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineApproveRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified VaccineApproveRes message, length delimited. Does not implicitly {@link dss.VaccineApproveRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {dss.IVaccineApproveRes} message VaccineApproveRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineApproveRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineApproveRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineApproveRes} VaccineApproveRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineApproveRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineApproveRes();
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
         * Decodes a VaccineApproveRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineApproveRes} VaccineApproveRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineApproveRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineApproveRes message.
         * @function verify
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineApproveRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a VaccineApproveRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineApproveRes} VaccineApproveRes
         */
        VaccineApproveRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineApproveRes)
                return object;
            return new $root.dss.VaccineApproveRes();
        };

        /**
         * Creates a plain object from a VaccineApproveRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineApproveRes
         * @static
         * @param {dss.VaccineApproveRes} message VaccineApproveRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineApproveRes.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this VaccineApproveRes to JSON.
         * @function toJSON
         * @memberof dss.VaccineApproveRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineApproveRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineApproveRes;
    })();

    dss.VaccineRejectReq = (function() {

        /**
         * Properties of a VaccineRejectReq.
         * @memberof dss
         * @interface IVaccineRejectReq
         * @property {string|null} [id] VaccineRejectReq id
         * @property {string|null} [reason] VaccineRejectReq reason
         */

        /**
         * Constructs a new VaccineRejectReq.
         * @memberof dss
         * @classdesc Represents a VaccineRejectReq.
         * @implements IVaccineRejectReq
         * @constructor
         * @param {dss.IVaccineRejectReq=} [properties] Properties to set
         */
        function VaccineRejectReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * VaccineRejectReq id.
         * @member {string} id
         * @memberof dss.VaccineRejectReq
         * @instance
         */
        VaccineRejectReq.prototype.id = "";

        /**
         * VaccineRejectReq reason.
         * @member {string} reason
         * @memberof dss.VaccineRejectReq
         * @instance
         */
        VaccineRejectReq.prototype.reason = "";

        /**
         * Creates a new VaccineRejectReq instance using the specified properties.
         * @function create
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {dss.IVaccineRejectReq=} [properties] Properties to set
         * @returns {dss.VaccineRejectReq} VaccineRejectReq instance
         */
        VaccineRejectReq.create = function create(properties) {
            return new VaccineRejectReq(properties);
        };

        /**
         * Encodes the specified VaccineRejectReq message. Does not implicitly {@link dss.VaccineRejectReq.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {dss.IVaccineRejectReq} message VaccineRejectReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineRejectReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            if (message.reason != null && Object.hasOwnProperty.call(message, "reason"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.reason);
            return writer;
        };

        /**
         * Encodes the specified VaccineRejectReq message, length delimited. Does not implicitly {@link dss.VaccineRejectReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {dss.IVaccineRejectReq} message VaccineRejectReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineRejectReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineRejectReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineRejectReq} VaccineRejectReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineRejectReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineRejectReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.reason = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a VaccineRejectReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineRejectReq} VaccineRejectReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineRejectReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineRejectReq message.
         * @function verify
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineRejectReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            if (message.reason != null && message.hasOwnProperty("reason"))
                if (!$util.isString(message.reason))
                    return "reason: string expected";
            return null;
        };

        /**
         * Creates a VaccineRejectReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineRejectReq} VaccineRejectReq
         */
        VaccineRejectReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineRejectReq)
                return object;
            var message = new $root.dss.VaccineRejectReq();
            if (object.id != null)
                message.id = String(object.id);
            if (object.reason != null)
                message.reason = String(object.reason);
            return message;
        };

        /**
         * Creates a plain object from a VaccineRejectReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineRejectReq
         * @static
         * @param {dss.VaccineRejectReq} message VaccineRejectReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineRejectReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.id = "";
                object.reason = "";
            }
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.reason != null && message.hasOwnProperty("reason"))
                object.reason = message.reason;
            return object;
        };

        /**
         * Converts this VaccineRejectReq to JSON.
         * @function toJSON
         * @memberof dss.VaccineRejectReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineRejectReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineRejectReq;
    })();

    dss.VaccineRejectRes = (function() {

        /**
         * Properties of a VaccineRejectRes.
         * @memberof dss
         * @interface IVaccineRejectRes
         */

        /**
         * Constructs a new VaccineRejectRes.
         * @memberof dss
         * @classdesc Represents a VaccineRejectRes.
         * @implements IVaccineRejectRes
         * @constructor
         * @param {dss.IVaccineRejectRes=} [properties] Properties to set
         */
        function VaccineRejectRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new VaccineRejectRes instance using the specified properties.
         * @function create
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {dss.IVaccineRejectRes=} [properties] Properties to set
         * @returns {dss.VaccineRejectRes} VaccineRejectRes instance
         */
        VaccineRejectRes.create = function create(properties) {
            return new VaccineRejectRes(properties);
        };

        /**
         * Encodes the specified VaccineRejectRes message. Does not implicitly {@link dss.VaccineRejectRes.verify|verify} messages.
         * @function encode
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {dss.IVaccineRejectRes} message VaccineRejectRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineRejectRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified VaccineRejectRes message, length delimited. Does not implicitly {@link dss.VaccineRejectRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {dss.IVaccineRejectRes} message VaccineRejectRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        VaccineRejectRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a VaccineRejectRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.VaccineRejectRes} VaccineRejectRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineRejectRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.VaccineRejectRes();
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
         * Decodes a VaccineRejectRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.VaccineRejectRes} VaccineRejectRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        VaccineRejectRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a VaccineRejectRes message.
         * @function verify
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        VaccineRejectRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a VaccineRejectRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.VaccineRejectRes} VaccineRejectRes
         */
        VaccineRejectRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.VaccineRejectRes)
                return object;
            return new $root.dss.VaccineRejectRes();
        };

        /**
         * Creates a plain object from a VaccineRejectRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.VaccineRejectRes
         * @static
         * @param {dss.VaccineRejectRes} message VaccineRejectRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        VaccineRejectRes.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this VaccineRejectRes to JSON.
         * @function toJSON
         * @memberof dss.VaccineRejectRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        VaccineRejectRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return VaccineRejectRes;
    })();

    return dss;
})();

module.exports = $root;
