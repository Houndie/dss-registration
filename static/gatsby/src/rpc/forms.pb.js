/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars, strict, no-lone-blocks, default-case*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots.forms || ($protobuf.roots.forms = {});

$root.dss = (function() {

    /**
     * Namespace dss.
     * @exports dss
     * @namespace
     */
    var dss = {};

    dss.Forms = (function() {

        /**
         * Constructs a new Forms service.
         * @memberof dss
         * @classdesc Represents a Forms
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Forms(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Forms.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Forms;

        /**
         * Creates new Forms service using the specified rpc implementation.
         * @function create
         * @memberof dss.Forms
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {Forms} RPC service. Useful where requests and/or responses are streamed.
         */
        Forms.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link dss.Forms#contactUs}.
         * @memberof dss.Forms
         * @typedef ContactUsCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.ContactUsRes} [response] ContactUsRes
         */

        /**
         * Calls ContactUs.
         * @function contactUs
         * @memberof dss.Forms
         * @instance
         * @param {dss.IContactUsReq} request ContactUsReq message or plain object
         * @param {dss.Forms.ContactUsCallback} callback Node-style callback called with the error, if any, and ContactUsRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Forms.prototype.contactUs = function contactUs(request, callback) {
            return this.rpcCall(contactUs, $root.dss.ContactUsReq, $root.dss.ContactUsRes, request, callback);
        }, "name", { value: "ContactUs" });

        /**
         * Calls ContactUs.
         * @function contactUs
         * @memberof dss.Forms
         * @instance
         * @param {dss.IContactUsReq} request ContactUsReq message or plain object
         * @returns {Promise<dss.ContactUsRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Forms#safetyReport}.
         * @memberof dss.Forms
         * @typedef SafetyReportCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.SafetyReportRes} [response] SafetyReportRes
         */

        /**
         * Calls SafetyReport.
         * @function safetyReport
         * @memberof dss.Forms
         * @instance
         * @param {dss.ISafetyReportReq} request SafetyReportReq message or plain object
         * @param {dss.Forms.SafetyReportCallback} callback Node-style callback called with the error, if any, and SafetyReportRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Forms.prototype.safetyReport = function safetyReport(request, callback) {
            return this.rpcCall(safetyReport, $root.dss.SafetyReportReq, $root.dss.SafetyReportRes, request, callback);
        }, "name", { value: "SafetyReport" });

        /**
         * Calls SafetyReport.
         * @function safetyReport
         * @memberof dss.Forms
         * @instance
         * @param {dss.ISafetyReportReq} request SafetyReportReq message or plain object
         * @returns {Promise<dss.SafetyReportRes>} Promise
         * @variation 2
         */

        return Forms;
    })();

    dss.ContactUsReq = (function() {

        /**
         * Properties of a ContactUsReq.
         * @memberof dss
         * @interface IContactUsReq
         * @property {string|null} [name] ContactUsReq name
         * @property {string|null} [email] ContactUsReq email
         * @property {string|null} [msg] ContactUsReq msg
         * @property {string|null} [recaptchaResponse] ContactUsReq recaptchaResponse
         */

        /**
         * Constructs a new ContactUsReq.
         * @memberof dss
         * @classdesc Represents a ContactUsReq.
         * @implements IContactUsReq
         * @constructor
         * @param {dss.IContactUsReq=} [properties] Properties to set
         */
        function ContactUsReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ContactUsReq name.
         * @member {string} name
         * @memberof dss.ContactUsReq
         * @instance
         */
        ContactUsReq.prototype.name = "";

        /**
         * ContactUsReq email.
         * @member {string} email
         * @memberof dss.ContactUsReq
         * @instance
         */
        ContactUsReq.prototype.email = "";

        /**
         * ContactUsReq msg.
         * @member {string} msg
         * @memberof dss.ContactUsReq
         * @instance
         */
        ContactUsReq.prototype.msg = "";

        /**
         * ContactUsReq recaptchaResponse.
         * @member {string} recaptchaResponse
         * @memberof dss.ContactUsReq
         * @instance
         */
        ContactUsReq.prototype.recaptchaResponse = "";

        /**
         * Creates a new ContactUsReq instance using the specified properties.
         * @function create
         * @memberof dss.ContactUsReq
         * @static
         * @param {dss.IContactUsReq=} [properties] Properties to set
         * @returns {dss.ContactUsReq} ContactUsReq instance
         */
        ContactUsReq.create = function create(properties) {
            return new ContactUsReq(properties);
        };

        /**
         * Encodes the specified ContactUsReq message. Does not implicitly {@link dss.ContactUsReq.verify|verify} messages.
         * @function encode
         * @memberof dss.ContactUsReq
         * @static
         * @param {dss.IContactUsReq} message ContactUsReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ContactUsReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
            if (message.email != null && Object.hasOwnProperty.call(message, "email"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.email);
            if (message.msg != null && Object.hasOwnProperty.call(message, "msg"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.msg);
            if (message.recaptchaResponse != null && Object.hasOwnProperty.call(message, "recaptchaResponse"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.recaptchaResponse);
            return writer;
        };

        /**
         * Encodes the specified ContactUsReq message, length delimited. Does not implicitly {@link dss.ContactUsReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.ContactUsReq
         * @static
         * @param {dss.IContactUsReq} message ContactUsReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ContactUsReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ContactUsReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.ContactUsReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.ContactUsReq} ContactUsReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ContactUsReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.ContactUsReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.email = reader.string();
                    break;
                case 3:
                    message.msg = reader.string();
                    break;
                case 4:
                    message.recaptchaResponse = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ContactUsReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.ContactUsReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.ContactUsReq} ContactUsReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ContactUsReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ContactUsReq message.
         * @function verify
         * @memberof dss.ContactUsReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ContactUsReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.email != null && message.hasOwnProperty("email"))
                if (!$util.isString(message.email))
                    return "email: string expected";
            if (message.msg != null && message.hasOwnProperty("msg"))
                if (!$util.isString(message.msg))
                    return "msg: string expected";
            if (message.recaptchaResponse != null && message.hasOwnProperty("recaptchaResponse"))
                if (!$util.isString(message.recaptchaResponse))
                    return "recaptchaResponse: string expected";
            return null;
        };

        /**
         * Creates a ContactUsReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.ContactUsReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.ContactUsReq} ContactUsReq
         */
        ContactUsReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.ContactUsReq)
                return object;
            var message = new $root.dss.ContactUsReq();
            if (object.name != null)
                message.name = String(object.name);
            if (object.email != null)
                message.email = String(object.email);
            if (object.msg != null)
                message.msg = String(object.msg);
            if (object.recaptchaResponse != null)
                message.recaptchaResponse = String(object.recaptchaResponse);
            return message;
        };

        /**
         * Creates a plain object from a ContactUsReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.ContactUsReq
         * @static
         * @param {dss.ContactUsReq} message ContactUsReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ContactUsReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.name = "";
                object.email = "";
                object.msg = "";
                object.recaptchaResponse = "";
            }
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.email != null && message.hasOwnProperty("email"))
                object.email = message.email;
            if (message.msg != null && message.hasOwnProperty("msg"))
                object.msg = message.msg;
            if (message.recaptchaResponse != null && message.hasOwnProperty("recaptchaResponse"))
                object.recaptchaResponse = message.recaptchaResponse;
            return object;
        };

        /**
         * Converts this ContactUsReq to JSON.
         * @function toJSON
         * @memberof dss.ContactUsReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ContactUsReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ContactUsReq;
    })();

    dss.ContactUsRes = (function() {

        /**
         * Properties of a ContactUsRes.
         * @memberof dss
         * @interface IContactUsRes
         */

        /**
         * Constructs a new ContactUsRes.
         * @memberof dss
         * @classdesc Represents a ContactUsRes.
         * @implements IContactUsRes
         * @constructor
         * @param {dss.IContactUsRes=} [properties] Properties to set
         */
        function ContactUsRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new ContactUsRes instance using the specified properties.
         * @function create
         * @memberof dss.ContactUsRes
         * @static
         * @param {dss.IContactUsRes=} [properties] Properties to set
         * @returns {dss.ContactUsRes} ContactUsRes instance
         */
        ContactUsRes.create = function create(properties) {
            return new ContactUsRes(properties);
        };

        /**
         * Encodes the specified ContactUsRes message. Does not implicitly {@link dss.ContactUsRes.verify|verify} messages.
         * @function encode
         * @memberof dss.ContactUsRes
         * @static
         * @param {dss.IContactUsRes} message ContactUsRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ContactUsRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified ContactUsRes message, length delimited. Does not implicitly {@link dss.ContactUsRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.ContactUsRes
         * @static
         * @param {dss.IContactUsRes} message ContactUsRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ContactUsRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ContactUsRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.ContactUsRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.ContactUsRes} ContactUsRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ContactUsRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.ContactUsRes();
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
         * Decodes a ContactUsRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.ContactUsRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.ContactUsRes} ContactUsRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ContactUsRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ContactUsRes message.
         * @function verify
         * @memberof dss.ContactUsRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ContactUsRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a ContactUsRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.ContactUsRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.ContactUsRes} ContactUsRes
         */
        ContactUsRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.ContactUsRes)
                return object;
            return new $root.dss.ContactUsRes();
        };

        /**
         * Creates a plain object from a ContactUsRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.ContactUsRes
         * @static
         * @param {dss.ContactUsRes} message ContactUsRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ContactUsRes.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this ContactUsRes to JSON.
         * @function toJSON
         * @memberof dss.ContactUsRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ContactUsRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ContactUsRes;
    })();

    dss.SafetyReportReq = (function() {

        /**
         * Properties of a SafetyReportReq.
         * @memberof dss
         * @interface ISafetyReportReq
         * @property {google.protobuf.ITimestamp|null} [occurredOn] SafetyReportReq occurredOn
         * @property {string|null} [description] SafetyReportReq description
         * @property {number|null} [severity] SafetyReportReq severity
         * @property {boolean|null} [issuesBefore] SafetyReportReq issuesBefore
         * @property {string|null} [resolution] SafetyReportReq resolution
         * @property {string|null} [name] SafetyReportReq name
         * @property {string|null} [email] SafetyReportReq email
         * @property {string|null} [phoneNumber] SafetyReportReq phoneNumber
         * @property {string|null} [recaptchaResponse] SafetyReportReq recaptchaResponse
         */

        /**
         * Constructs a new SafetyReportReq.
         * @memberof dss
         * @classdesc Represents a SafetyReportReq.
         * @implements ISafetyReportReq
         * @constructor
         * @param {dss.ISafetyReportReq=} [properties] Properties to set
         */
        function SafetyReportReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * SafetyReportReq occurredOn.
         * @member {google.protobuf.ITimestamp|null|undefined} occurredOn
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.occurredOn = null;

        /**
         * SafetyReportReq description.
         * @member {string} description
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.description = "";

        /**
         * SafetyReportReq severity.
         * @member {number} severity
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.severity = 0;

        /**
         * SafetyReportReq issuesBefore.
         * @member {boolean} issuesBefore
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.issuesBefore = false;

        /**
         * SafetyReportReq resolution.
         * @member {string} resolution
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.resolution = "";

        /**
         * SafetyReportReq name.
         * @member {string} name
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.name = "";

        /**
         * SafetyReportReq email.
         * @member {string} email
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.email = "";

        /**
         * SafetyReportReq phoneNumber.
         * @member {string} phoneNumber
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.phoneNumber = "";

        /**
         * SafetyReportReq recaptchaResponse.
         * @member {string} recaptchaResponse
         * @memberof dss.SafetyReportReq
         * @instance
         */
        SafetyReportReq.prototype.recaptchaResponse = "";

        /**
         * Creates a new SafetyReportReq instance using the specified properties.
         * @function create
         * @memberof dss.SafetyReportReq
         * @static
         * @param {dss.ISafetyReportReq=} [properties] Properties to set
         * @returns {dss.SafetyReportReq} SafetyReportReq instance
         */
        SafetyReportReq.create = function create(properties) {
            return new SafetyReportReq(properties);
        };

        /**
         * Encodes the specified SafetyReportReq message. Does not implicitly {@link dss.SafetyReportReq.verify|verify} messages.
         * @function encode
         * @memberof dss.SafetyReportReq
         * @static
         * @param {dss.ISafetyReportReq} message SafetyReportReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SafetyReportReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.occurredOn != null && Object.hasOwnProperty.call(message, "occurredOn"))
                $root.google.protobuf.Timestamp.encode(message.occurredOn, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.description != null && Object.hasOwnProperty.call(message, "description"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.description);
            if (message.severity != null && Object.hasOwnProperty.call(message, "severity"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.severity);
            if (message.issuesBefore != null && Object.hasOwnProperty.call(message, "issuesBefore"))
                writer.uint32(/* id 4, wireType 0 =*/32).bool(message.issuesBefore);
            if (message.resolution != null && Object.hasOwnProperty.call(message, "resolution"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.resolution);
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.name);
            if (message.email != null && Object.hasOwnProperty.call(message, "email"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.email);
            if (message.phoneNumber != null && Object.hasOwnProperty.call(message, "phoneNumber"))
                writer.uint32(/* id 8, wireType 2 =*/66).string(message.phoneNumber);
            if (message.recaptchaResponse != null && Object.hasOwnProperty.call(message, "recaptchaResponse"))
                writer.uint32(/* id 9, wireType 2 =*/74).string(message.recaptchaResponse);
            return writer;
        };

        /**
         * Encodes the specified SafetyReportReq message, length delimited. Does not implicitly {@link dss.SafetyReportReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.SafetyReportReq
         * @static
         * @param {dss.ISafetyReportReq} message SafetyReportReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SafetyReportReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a SafetyReportReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.SafetyReportReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.SafetyReportReq} SafetyReportReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SafetyReportReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.SafetyReportReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.occurredOn = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.severity = reader.int32();
                    break;
                case 4:
                    message.issuesBefore = reader.bool();
                    break;
                case 5:
                    message.resolution = reader.string();
                    break;
                case 6:
                    message.name = reader.string();
                    break;
                case 7:
                    message.email = reader.string();
                    break;
                case 8:
                    message.phoneNumber = reader.string();
                    break;
                case 9:
                    message.recaptchaResponse = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a SafetyReportReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.SafetyReportReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.SafetyReportReq} SafetyReportReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SafetyReportReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a SafetyReportReq message.
         * @function verify
         * @memberof dss.SafetyReportReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        SafetyReportReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.occurredOn != null && message.hasOwnProperty("occurredOn")) {
                var error = $root.google.protobuf.Timestamp.verify(message.occurredOn);
                if (error)
                    return "occurredOn." + error;
            }
            if (message.description != null && message.hasOwnProperty("description"))
                if (!$util.isString(message.description))
                    return "description: string expected";
            if (message.severity != null && message.hasOwnProperty("severity"))
                if (!$util.isInteger(message.severity))
                    return "severity: integer expected";
            if (message.issuesBefore != null && message.hasOwnProperty("issuesBefore"))
                if (typeof message.issuesBefore !== "boolean")
                    return "issuesBefore: boolean expected";
            if (message.resolution != null && message.hasOwnProperty("resolution"))
                if (!$util.isString(message.resolution))
                    return "resolution: string expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.email != null && message.hasOwnProperty("email"))
                if (!$util.isString(message.email))
                    return "email: string expected";
            if (message.phoneNumber != null && message.hasOwnProperty("phoneNumber"))
                if (!$util.isString(message.phoneNumber))
                    return "phoneNumber: string expected";
            if (message.recaptchaResponse != null && message.hasOwnProperty("recaptchaResponse"))
                if (!$util.isString(message.recaptchaResponse))
                    return "recaptchaResponse: string expected";
            return null;
        };

        /**
         * Creates a SafetyReportReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.SafetyReportReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.SafetyReportReq} SafetyReportReq
         */
        SafetyReportReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.SafetyReportReq)
                return object;
            var message = new $root.dss.SafetyReportReq();
            if (object.occurredOn != null) {
                if (typeof object.occurredOn !== "object")
                    throw TypeError(".dss.SafetyReportReq.occurredOn: object expected");
                message.occurredOn = $root.google.protobuf.Timestamp.fromObject(object.occurredOn);
            }
            if (object.description != null)
                message.description = String(object.description);
            if (object.severity != null)
                message.severity = object.severity | 0;
            if (object.issuesBefore != null)
                message.issuesBefore = Boolean(object.issuesBefore);
            if (object.resolution != null)
                message.resolution = String(object.resolution);
            if (object.name != null)
                message.name = String(object.name);
            if (object.email != null)
                message.email = String(object.email);
            if (object.phoneNumber != null)
                message.phoneNumber = String(object.phoneNumber);
            if (object.recaptchaResponse != null)
                message.recaptchaResponse = String(object.recaptchaResponse);
            return message;
        };

        /**
         * Creates a plain object from a SafetyReportReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.SafetyReportReq
         * @static
         * @param {dss.SafetyReportReq} message SafetyReportReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        SafetyReportReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.occurredOn = null;
                object.description = "";
                object.severity = 0;
                object.issuesBefore = false;
                object.resolution = "";
                object.name = "";
                object.email = "";
                object.phoneNumber = "";
                object.recaptchaResponse = "";
            }
            if (message.occurredOn != null && message.hasOwnProperty("occurredOn"))
                object.occurredOn = $root.google.protobuf.Timestamp.toObject(message.occurredOn, options);
            if (message.description != null && message.hasOwnProperty("description"))
                object.description = message.description;
            if (message.severity != null && message.hasOwnProperty("severity"))
                object.severity = message.severity;
            if (message.issuesBefore != null && message.hasOwnProperty("issuesBefore"))
                object.issuesBefore = message.issuesBefore;
            if (message.resolution != null && message.hasOwnProperty("resolution"))
                object.resolution = message.resolution;
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.email != null && message.hasOwnProperty("email"))
                object.email = message.email;
            if (message.phoneNumber != null && message.hasOwnProperty("phoneNumber"))
                object.phoneNumber = message.phoneNumber;
            if (message.recaptchaResponse != null && message.hasOwnProperty("recaptchaResponse"))
                object.recaptchaResponse = message.recaptchaResponse;
            return object;
        };

        /**
         * Converts this SafetyReportReq to JSON.
         * @function toJSON
         * @memberof dss.SafetyReportReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        SafetyReportReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return SafetyReportReq;
    })();

    dss.SafetyReportRes = (function() {

        /**
         * Properties of a SafetyReportRes.
         * @memberof dss
         * @interface ISafetyReportRes
         */

        /**
         * Constructs a new SafetyReportRes.
         * @memberof dss
         * @classdesc Represents a SafetyReportRes.
         * @implements ISafetyReportRes
         * @constructor
         * @param {dss.ISafetyReportRes=} [properties] Properties to set
         */
        function SafetyReportRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new SafetyReportRes instance using the specified properties.
         * @function create
         * @memberof dss.SafetyReportRes
         * @static
         * @param {dss.ISafetyReportRes=} [properties] Properties to set
         * @returns {dss.SafetyReportRes} SafetyReportRes instance
         */
        SafetyReportRes.create = function create(properties) {
            return new SafetyReportRes(properties);
        };

        /**
         * Encodes the specified SafetyReportRes message. Does not implicitly {@link dss.SafetyReportRes.verify|verify} messages.
         * @function encode
         * @memberof dss.SafetyReportRes
         * @static
         * @param {dss.ISafetyReportRes} message SafetyReportRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SafetyReportRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified SafetyReportRes message, length delimited. Does not implicitly {@link dss.SafetyReportRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.SafetyReportRes
         * @static
         * @param {dss.ISafetyReportRes} message SafetyReportRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SafetyReportRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a SafetyReportRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.SafetyReportRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.SafetyReportRes} SafetyReportRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SafetyReportRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.SafetyReportRes();
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
         * Decodes a SafetyReportRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.SafetyReportRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.SafetyReportRes} SafetyReportRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SafetyReportRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a SafetyReportRes message.
         * @function verify
         * @memberof dss.SafetyReportRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        SafetyReportRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a SafetyReportRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.SafetyReportRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.SafetyReportRes} SafetyReportRes
         */
        SafetyReportRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.SafetyReportRes)
                return object;
            return new $root.dss.SafetyReportRes();
        };

        /**
         * Creates a plain object from a SafetyReportRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.SafetyReportRes
         * @static
         * @param {dss.SafetyReportRes} message SafetyReportRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        SafetyReportRes.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this SafetyReportRes to JSON.
         * @function toJSON
         * @memberof dss.SafetyReportRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        SafetyReportRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return SafetyReportRes;
    })();

    return dss;
})();

$root.google = (function() {

    /**
     * Namespace google.
     * @exports google
     * @namespace
     */
    var google = {};

    google.protobuf = (function() {

        /**
         * Namespace protobuf.
         * @memberof google
         * @namespace
         */
        var protobuf = {};

        protobuf.Timestamp = (function() {

            /**
             * Properties of a Timestamp.
             * @memberof google.protobuf
             * @interface ITimestamp
             * @property {number|Long|null} [seconds] Timestamp seconds
             * @property {number|null} [nanos] Timestamp nanos
             */

            /**
             * Constructs a new Timestamp.
             * @memberof google.protobuf
             * @classdesc Represents a Timestamp.
             * @implements ITimestamp
             * @constructor
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             */
            function Timestamp(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Timestamp seconds.
             * @member {number|Long} seconds
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.seconds = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

            /**
             * Timestamp nanos.
             * @member {number} nanos
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.nanos = 0;

            /**
             * Creates a new Timestamp instance using the specified properties.
             * @function create
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             * @returns {google.protobuf.Timestamp} Timestamp instance
             */
            Timestamp.create = function create(properties) {
                return new Timestamp(properties);
            };

            /**
             * Encodes the specified Timestamp message. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.seconds != null && Object.hasOwnProperty.call(message, "seconds"))
                    writer.uint32(/* id 1, wireType 0 =*/8).int64(message.seconds);
                if (message.nanos != null && Object.hasOwnProperty.call(message, "nanos"))
                    writer.uint32(/* id 2, wireType 0 =*/16).int32(message.nanos);
                return writer;
            };

            /**
             * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer.
             * @function decode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.google.protobuf.Timestamp();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.seconds = reader.int64();
                        break;
                    case 2:
                        message.nanos = reader.int32();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Timestamp message.
             * @function verify
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Timestamp.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (!$util.isInteger(message.seconds) && !(message.seconds && $util.isInteger(message.seconds.low) && $util.isInteger(message.seconds.high)))
                        return "seconds: integer|Long expected";
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    if (!$util.isInteger(message.nanos))
                        return "nanos: integer expected";
                return null;
            };

            /**
             * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {google.protobuf.Timestamp} Timestamp
             */
            Timestamp.fromObject = function fromObject(object) {
                if (object instanceof $root.google.protobuf.Timestamp)
                    return object;
                var message = new $root.google.protobuf.Timestamp();
                if (object.seconds != null)
                    if ($util.Long)
                        (message.seconds = $util.Long.fromValue(object.seconds)).unsigned = false;
                    else if (typeof object.seconds === "string")
                        message.seconds = parseInt(object.seconds, 10);
                    else if (typeof object.seconds === "number")
                        message.seconds = object.seconds;
                    else if (typeof object.seconds === "object")
                        message.seconds = new $util.LongBits(object.seconds.low >>> 0, object.seconds.high >>> 0).toNumber();
                if (object.nanos != null)
                    message.nanos = object.nanos | 0;
                return message;
            };

            /**
             * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
             * @function toObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.Timestamp} message Timestamp
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Timestamp.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults) {
                    if ($util.Long) {
                        var long = new $util.Long(0, 0, false);
                        object.seconds = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.seconds = options.longs === String ? "0" : 0;
                    object.nanos = 0;
                }
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (typeof message.seconds === "number")
                        object.seconds = options.longs === String ? String(message.seconds) : message.seconds;
                    else
                        object.seconds = options.longs === String ? $util.Long.prototype.toString.call(message.seconds) : options.longs === Number ? new $util.LongBits(message.seconds.low >>> 0, message.seconds.high >>> 0).toNumber() : message.seconds;
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    object.nanos = message.nanos;
                return object;
            };

            /**
             * Converts this Timestamp to JSON.
             * @function toJSON
             * @memberof google.protobuf.Timestamp
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Timestamp.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Timestamp;
        })();

        return protobuf;
    })();

    return google;
})();

module.exports = $root;
