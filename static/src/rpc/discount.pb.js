/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars, strict, no-lone-blocks, default-case*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots.discount || ($protobuf.roots.discount = {});

$root.dss = (function() {

    /**
     * Namespace dss.
     * @exports dss
     * @namespace
     */
    var dss = {};

    dss.Discount = (function() {

        /**
         * Constructs a new Discount service.
         * @memberof dss
         * @classdesc Represents a Discount
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Discount(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Discount.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Discount;

        /**
         * Creates new Discount service using the specified rpc implementation.
         * @function create
         * @memberof dss.Discount
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {Discount} RPC service. Useful where requests and/or responses are streamed.
         */
        Discount.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link dss.Discount#get}.
         * @memberof dss.Discount
         * @typedef GetCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.DiscountGetRes} [response] DiscountGetRes
         */

        /**
         * Calls Get.
         * @function get
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountGetReq} request DiscountGetReq message or plain object
         * @param {dss.Discount.GetCallback} callback Node-style callback called with the error, if any, and DiscountGetRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Discount.prototype.get = function get(request, callback) {
            return this.rpcCall(get, $root.dss.DiscountGetReq, $root.dss.DiscountGetRes, request, callback);
        }, "name", { value: "Get" });

        /**
         * Calls Get.
         * @function get
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountGetReq} request DiscountGetReq message or plain object
         * @returns {Promise<dss.DiscountGetRes>} Promise
         * @variation 2
         */

        return Discount;
    })();

    /**
     * PurchaseItem enum.
     * @name dss.PurchaseItem
     * @enum {number}
     * @property {number} Unknown=0 Unknown value
     * @property {number} FullWeekendPassPurchaseItem=1 FullWeekendPassPurchaseItem value
     * @property {number} DanceOnlyPassPurchaseItem=2 DanceOnlyPassPurchaseItem value
     * @property {number} MixAndMatchPurchaseItem=3 MixAndMatchPurchaseItem value
     * @property {number} SoloJazzPurchaseItem=4 SoloJazzPurchaseItem value
     * @property {number} TeamCompetitionPurchaseItem=5 TeamCompetitionPurchaseItem value
     * @property {number} TShirtPurchaseItem=6 TShirtPurchaseItem value
     */
    dss.PurchaseItem = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Unknown"] = 0;
        values[valuesById[1] = "FullWeekendPassPurchaseItem"] = 1;
        values[valuesById[2] = "DanceOnlyPassPurchaseItem"] = 2;
        values[valuesById[3] = "MixAndMatchPurchaseItem"] = 3;
        values[valuesById[4] = "SoloJazzPurchaseItem"] = 4;
        values[valuesById[5] = "TeamCompetitionPurchaseItem"] = 5;
        values[valuesById[6] = "TShirtPurchaseItem"] = 6;
        return values;
    })();

    dss.DiscountAmount = (function() {

        /**
         * Properties of a DiscountAmount.
         * @memberof dss
         * @interface IDiscountAmount
         * @property {number|Long|null} [dollar] DiscountAmount dollar
         * @property {string|null} [percent] DiscountAmount percent
         */

        /**
         * Constructs a new DiscountAmount.
         * @memberof dss
         * @classdesc Represents a DiscountAmount.
         * @implements IDiscountAmount
         * @constructor
         * @param {dss.IDiscountAmount=} [properties] Properties to set
         */
        function DiscountAmount(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountAmount dollar.
         * @member {number|Long} dollar
         * @memberof dss.DiscountAmount
         * @instance
         */
        DiscountAmount.prototype.dollar = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * DiscountAmount percent.
         * @member {string} percent
         * @memberof dss.DiscountAmount
         * @instance
         */
        DiscountAmount.prototype.percent = "";

        // OneOf field names bound to virtual getters and setters
        var $oneOfFields;

        /**
         * DiscountAmount amount.
         * @member {"dollar"|"percent"|undefined} amount
         * @memberof dss.DiscountAmount
         * @instance
         */
        Object.defineProperty(DiscountAmount.prototype, "amount", {
            get: $util.oneOfGetter($oneOfFields = ["dollar", "percent"]),
            set: $util.oneOfSetter($oneOfFields)
        });

        /**
         * Creates a new DiscountAmount instance using the specified properties.
         * @function create
         * @memberof dss.DiscountAmount
         * @static
         * @param {dss.IDiscountAmount=} [properties] Properties to set
         * @returns {dss.DiscountAmount} DiscountAmount instance
         */
        DiscountAmount.create = function create(properties) {
            return new DiscountAmount(properties);
        };

        /**
         * Encodes the specified DiscountAmount message. Does not implicitly {@link dss.DiscountAmount.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountAmount
         * @static
         * @param {dss.IDiscountAmount} message DiscountAmount message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountAmount.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.dollar != null && Object.hasOwnProperty.call(message, "dollar"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.dollar);
            if (message.percent != null && Object.hasOwnProperty.call(message, "percent"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.percent);
            return writer;
        };

        /**
         * Encodes the specified DiscountAmount message, length delimited. Does not implicitly {@link dss.DiscountAmount.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountAmount
         * @static
         * @param {dss.IDiscountAmount} message DiscountAmount message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountAmount.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountAmount message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountAmount
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountAmount} DiscountAmount
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountAmount.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountAmount();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.dollar = reader.int64();
                    break;
                case 2:
                    message.percent = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DiscountAmount message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountAmount
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountAmount} DiscountAmount
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountAmount.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountAmount message.
         * @function verify
         * @memberof dss.DiscountAmount
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountAmount.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            var properties = {};
            if (message.dollar != null && message.hasOwnProperty("dollar")) {
                properties.amount = 1;
                if (!$util.isInteger(message.dollar) && !(message.dollar && $util.isInteger(message.dollar.low) && $util.isInteger(message.dollar.high)))
                    return "dollar: integer|Long expected";
            }
            if (message.percent != null && message.hasOwnProperty("percent")) {
                if (properties.amount === 1)
                    return "amount: multiple values";
                properties.amount = 1;
                if (!$util.isString(message.percent))
                    return "percent: string expected";
            }
            return null;
        };

        /**
         * Creates a DiscountAmount message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountAmount
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountAmount} DiscountAmount
         */
        DiscountAmount.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountAmount)
                return object;
            var message = new $root.dss.DiscountAmount();
            if (object.dollar != null)
                if ($util.Long)
                    (message.dollar = $util.Long.fromValue(object.dollar)).unsigned = false;
                else if (typeof object.dollar === "string")
                    message.dollar = parseInt(object.dollar, 10);
                else if (typeof object.dollar === "number")
                    message.dollar = object.dollar;
                else if (typeof object.dollar === "object")
                    message.dollar = new $util.LongBits(object.dollar.low >>> 0, object.dollar.high >>> 0).toNumber();
            if (object.percent != null)
                message.percent = String(object.percent);
            return message;
        };

        /**
         * Creates a plain object from a DiscountAmount message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountAmount
         * @static
         * @param {dss.DiscountAmount} message DiscountAmount
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountAmount.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (message.dollar != null && message.hasOwnProperty("dollar")) {
                if (typeof message.dollar === "number")
                    object.dollar = options.longs === String ? String(message.dollar) : message.dollar;
                else
                    object.dollar = options.longs === String ? $util.Long.prototype.toString.call(message.dollar) : options.longs === Number ? new $util.LongBits(message.dollar.low >>> 0, message.dollar.high >>> 0).toNumber() : message.dollar;
                if (options.oneofs)
                    object.amount = "dollar";
            }
            if (message.percent != null && message.hasOwnProperty("percent")) {
                object.percent = message.percent;
                if (options.oneofs)
                    object.amount = "percent";
            }
            return object;
        };

        /**
         * Converts this DiscountAmount to JSON.
         * @function toJSON
         * @memberof dss.DiscountAmount
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountAmount.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountAmount;
    })();

    dss.SingleDiscount = (function() {

        /**
         * Properties of a SingleDiscount.
         * @memberof dss
         * @interface ISingleDiscount
         * @property {dss.IDiscountAmount|null} [amount] SingleDiscount amount
         * @property {dss.PurchaseItem|null} [appliedTo] SingleDiscount appliedTo
         */

        /**
         * Constructs a new SingleDiscount.
         * @memberof dss
         * @classdesc Represents a SingleDiscount.
         * @implements ISingleDiscount
         * @constructor
         * @param {dss.ISingleDiscount=} [properties] Properties to set
         */
        function SingleDiscount(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * SingleDiscount amount.
         * @member {dss.IDiscountAmount|null|undefined} amount
         * @memberof dss.SingleDiscount
         * @instance
         */
        SingleDiscount.prototype.amount = null;

        /**
         * SingleDiscount appliedTo.
         * @member {dss.PurchaseItem} appliedTo
         * @memberof dss.SingleDiscount
         * @instance
         */
        SingleDiscount.prototype.appliedTo = 0;

        /**
         * Creates a new SingleDiscount instance using the specified properties.
         * @function create
         * @memberof dss.SingleDiscount
         * @static
         * @param {dss.ISingleDiscount=} [properties] Properties to set
         * @returns {dss.SingleDiscount} SingleDiscount instance
         */
        SingleDiscount.create = function create(properties) {
            return new SingleDiscount(properties);
        };

        /**
         * Encodes the specified SingleDiscount message. Does not implicitly {@link dss.SingleDiscount.verify|verify} messages.
         * @function encode
         * @memberof dss.SingleDiscount
         * @static
         * @param {dss.ISingleDiscount} message SingleDiscount message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SingleDiscount.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.amount != null && Object.hasOwnProperty.call(message, "amount"))
                $root.dss.DiscountAmount.encode(message.amount, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.appliedTo != null && Object.hasOwnProperty.call(message, "appliedTo"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.appliedTo);
            return writer;
        };

        /**
         * Encodes the specified SingleDiscount message, length delimited. Does not implicitly {@link dss.SingleDiscount.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.SingleDiscount
         * @static
         * @param {dss.ISingleDiscount} message SingleDiscount message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SingleDiscount.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a SingleDiscount message from the specified reader or buffer.
         * @function decode
         * @memberof dss.SingleDiscount
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.SingleDiscount} SingleDiscount
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SingleDiscount.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.SingleDiscount();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.amount = $root.dss.DiscountAmount.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.appliedTo = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a SingleDiscount message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.SingleDiscount
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.SingleDiscount} SingleDiscount
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SingleDiscount.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a SingleDiscount message.
         * @function verify
         * @memberof dss.SingleDiscount
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        SingleDiscount.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.amount != null && message.hasOwnProperty("amount")) {
                var error = $root.dss.DiscountAmount.verify(message.amount);
                if (error)
                    return "amount." + error;
            }
            if (message.appliedTo != null && message.hasOwnProperty("appliedTo"))
                switch (message.appliedTo) {
                default:
                    return "appliedTo: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                case 5:
                case 6:
                    break;
                }
            return null;
        };

        /**
         * Creates a SingleDiscount message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.SingleDiscount
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.SingleDiscount} SingleDiscount
         */
        SingleDiscount.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.SingleDiscount)
                return object;
            var message = new $root.dss.SingleDiscount();
            if (object.amount != null) {
                if (typeof object.amount !== "object")
                    throw TypeError(".dss.SingleDiscount.amount: object expected");
                message.amount = $root.dss.DiscountAmount.fromObject(object.amount);
            }
            switch (object.appliedTo) {
            case "Unknown":
            case 0:
                message.appliedTo = 0;
                break;
            case "FullWeekendPassPurchaseItem":
            case 1:
                message.appliedTo = 1;
                break;
            case "DanceOnlyPassPurchaseItem":
            case 2:
                message.appliedTo = 2;
                break;
            case "MixAndMatchPurchaseItem":
            case 3:
                message.appliedTo = 3;
                break;
            case "SoloJazzPurchaseItem":
            case 4:
                message.appliedTo = 4;
                break;
            case "TeamCompetitionPurchaseItem":
            case 5:
                message.appliedTo = 5;
                break;
            case "TShirtPurchaseItem":
            case 6:
                message.appliedTo = 6;
                break;
            }
            return message;
        };

        /**
         * Creates a plain object from a SingleDiscount message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.SingleDiscount
         * @static
         * @param {dss.SingleDiscount} message SingleDiscount
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        SingleDiscount.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.amount = null;
                object.appliedTo = options.enums === String ? "Unknown" : 0;
            }
            if (message.amount != null && message.hasOwnProperty("amount"))
                object.amount = $root.dss.DiscountAmount.toObject(message.amount, options);
            if (message.appliedTo != null && message.hasOwnProperty("appliedTo"))
                object.appliedTo = options.enums === String ? $root.dss.PurchaseItem[message.appliedTo] : message.appliedTo;
            return object;
        };

        /**
         * Converts this SingleDiscount to JSON.
         * @function toJSON
         * @memberof dss.SingleDiscount
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        SingleDiscount.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return SingleDiscount;
    })();

    dss.DiscountBundle = (function() {

        /**
         * Properties of a DiscountBundle.
         * @memberof dss
         * @interface IDiscountBundle
         * @property {string|null} [code] DiscountBundle code
         * @property {Array.<dss.ISingleDiscount>|null} [discounts] DiscountBundle discounts
         */

        /**
         * Constructs a new DiscountBundle.
         * @memberof dss
         * @classdesc Represents a DiscountBundle.
         * @implements IDiscountBundle
         * @constructor
         * @param {dss.IDiscountBundle=} [properties] Properties to set
         */
        function DiscountBundle(properties) {
            this.discounts = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountBundle code.
         * @member {string} code
         * @memberof dss.DiscountBundle
         * @instance
         */
        DiscountBundle.prototype.code = "";

        /**
         * DiscountBundle discounts.
         * @member {Array.<dss.ISingleDiscount>} discounts
         * @memberof dss.DiscountBundle
         * @instance
         */
        DiscountBundle.prototype.discounts = $util.emptyArray;

        /**
         * Creates a new DiscountBundle instance using the specified properties.
         * @function create
         * @memberof dss.DiscountBundle
         * @static
         * @param {dss.IDiscountBundle=} [properties] Properties to set
         * @returns {dss.DiscountBundle} DiscountBundle instance
         */
        DiscountBundle.create = function create(properties) {
            return new DiscountBundle(properties);
        };

        /**
         * Encodes the specified DiscountBundle message. Does not implicitly {@link dss.DiscountBundle.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountBundle
         * @static
         * @param {dss.IDiscountBundle} message DiscountBundle message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountBundle.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.code != null && Object.hasOwnProperty.call(message, "code"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.code);
            if (message.discounts != null && message.discounts.length)
                for (var i = 0; i < message.discounts.length; ++i)
                    $root.dss.SingleDiscount.encode(message.discounts[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified DiscountBundle message, length delimited. Does not implicitly {@link dss.DiscountBundle.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountBundle
         * @static
         * @param {dss.IDiscountBundle} message DiscountBundle message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountBundle.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountBundle message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountBundle
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountBundle} DiscountBundle
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountBundle.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountBundle();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.code = reader.string();
                    break;
                case 2:
                    if (!(message.discounts && message.discounts.length))
                        message.discounts = [];
                    message.discounts.push($root.dss.SingleDiscount.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DiscountBundle message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountBundle
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountBundle} DiscountBundle
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountBundle.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountBundle message.
         * @function verify
         * @memberof dss.DiscountBundle
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountBundle.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.code != null && message.hasOwnProperty("code"))
                if (!$util.isString(message.code))
                    return "code: string expected";
            if (message.discounts != null && message.hasOwnProperty("discounts")) {
                if (!Array.isArray(message.discounts))
                    return "discounts: array expected";
                for (var i = 0; i < message.discounts.length; ++i) {
                    var error = $root.dss.SingleDiscount.verify(message.discounts[i]);
                    if (error)
                        return "discounts." + error;
                }
            }
            return null;
        };

        /**
         * Creates a DiscountBundle message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountBundle
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountBundle} DiscountBundle
         */
        DiscountBundle.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountBundle)
                return object;
            var message = new $root.dss.DiscountBundle();
            if (object.code != null)
                message.code = String(object.code);
            if (object.discounts) {
                if (!Array.isArray(object.discounts))
                    throw TypeError(".dss.DiscountBundle.discounts: array expected");
                message.discounts = [];
                for (var i = 0; i < object.discounts.length; ++i) {
                    if (typeof object.discounts[i] !== "object")
                        throw TypeError(".dss.DiscountBundle.discounts: object expected");
                    message.discounts[i] = $root.dss.SingleDiscount.fromObject(object.discounts[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a DiscountBundle message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountBundle
         * @static
         * @param {dss.DiscountBundle} message DiscountBundle
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountBundle.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.discounts = [];
            if (options.defaults)
                object.code = "";
            if (message.code != null && message.hasOwnProperty("code"))
                object.code = message.code;
            if (message.discounts && message.discounts.length) {
                object.discounts = [];
                for (var j = 0; j < message.discounts.length; ++j)
                    object.discounts[j] = $root.dss.SingleDiscount.toObject(message.discounts[j], options);
            }
            return object;
        };

        /**
         * Converts this DiscountBundle to JSON.
         * @function toJSON
         * @memberof dss.DiscountBundle
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountBundle.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountBundle;
    })();

    dss.DiscountGetReq = (function() {

        /**
         * Properties of a DiscountGetReq.
         * @memberof dss
         * @interface IDiscountGetReq
         * @property {string|null} [code] DiscountGetReq code
         */

        /**
         * Constructs a new DiscountGetReq.
         * @memberof dss
         * @classdesc Represents a DiscountGetReq.
         * @implements IDiscountGetReq
         * @constructor
         * @param {dss.IDiscountGetReq=} [properties] Properties to set
         */
        function DiscountGetReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountGetReq code.
         * @member {string} code
         * @memberof dss.DiscountGetReq
         * @instance
         */
        DiscountGetReq.prototype.code = "";

        /**
         * Creates a new DiscountGetReq instance using the specified properties.
         * @function create
         * @memberof dss.DiscountGetReq
         * @static
         * @param {dss.IDiscountGetReq=} [properties] Properties to set
         * @returns {dss.DiscountGetReq} DiscountGetReq instance
         */
        DiscountGetReq.create = function create(properties) {
            return new DiscountGetReq(properties);
        };

        /**
         * Encodes the specified DiscountGetReq message. Does not implicitly {@link dss.DiscountGetReq.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountGetReq
         * @static
         * @param {dss.IDiscountGetReq} message DiscountGetReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountGetReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.code != null && Object.hasOwnProperty.call(message, "code"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.code);
            return writer;
        };

        /**
         * Encodes the specified DiscountGetReq message, length delimited. Does not implicitly {@link dss.DiscountGetReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountGetReq
         * @static
         * @param {dss.IDiscountGetReq} message DiscountGetReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountGetReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountGetReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountGetReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountGetReq} DiscountGetReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountGetReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountGetReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.code = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DiscountGetReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountGetReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountGetReq} DiscountGetReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountGetReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountGetReq message.
         * @function verify
         * @memberof dss.DiscountGetReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountGetReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.code != null && message.hasOwnProperty("code"))
                if (!$util.isString(message.code))
                    return "code: string expected";
            return null;
        };

        /**
         * Creates a DiscountGetReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountGetReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountGetReq} DiscountGetReq
         */
        DiscountGetReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountGetReq)
                return object;
            var message = new $root.dss.DiscountGetReq();
            if (object.code != null)
                message.code = String(object.code);
            return message;
        };

        /**
         * Creates a plain object from a DiscountGetReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountGetReq
         * @static
         * @param {dss.DiscountGetReq} message DiscountGetReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountGetReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.code = "";
            if (message.code != null && message.hasOwnProperty("code"))
                object.code = message.code;
            return object;
        };

        /**
         * Converts this DiscountGetReq to JSON.
         * @function toJSON
         * @memberof dss.DiscountGetReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountGetReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountGetReq;
    })();

    dss.DiscountGetRes = (function() {

        /**
         * Properties of a DiscountGetRes.
         * @memberof dss
         * @interface IDiscountGetRes
         * @property {dss.IDiscountBundle|null} [bundle] DiscountGetRes bundle
         */

        /**
         * Constructs a new DiscountGetRes.
         * @memberof dss
         * @classdesc Represents a DiscountGetRes.
         * @implements IDiscountGetRes
         * @constructor
         * @param {dss.IDiscountGetRes=} [properties] Properties to set
         */
        function DiscountGetRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountGetRes bundle.
         * @member {dss.IDiscountBundle|null|undefined} bundle
         * @memberof dss.DiscountGetRes
         * @instance
         */
        DiscountGetRes.prototype.bundle = null;

        /**
         * Creates a new DiscountGetRes instance using the specified properties.
         * @function create
         * @memberof dss.DiscountGetRes
         * @static
         * @param {dss.IDiscountGetRes=} [properties] Properties to set
         * @returns {dss.DiscountGetRes} DiscountGetRes instance
         */
        DiscountGetRes.create = function create(properties) {
            return new DiscountGetRes(properties);
        };

        /**
         * Encodes the specified DiscountGetRes message. Does not implicitly {@link dss.DiscountGetRes.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountGetRes
         * @static
         * @param {dss.IDiscountGetRes} message DiscountGetRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountGetRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.bundle != null && Object.hasOwnProperty.call(message, "bundle"))
                $root.dss.DiscountBundle.encode(message.bundle, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified DiscountGetRes message, length delimited. Does not implicitly {@link dss.DiscountGetRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountGetRes
         * @static
         * @param {dss.IDiscountGetRes} message DiscountGetRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountGetRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountGetRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountGetRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountGetRes} DiscountGetRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountGetRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountGetRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.bundle = $root.dss.DiscountBundle.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DiscountGetRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountGetRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountGetRes} DiscountGetRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountGetRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountGetRes message.
         * @function verify
         * @memberof dss.DiscountGetRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountGetRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.bundle != null && message.hasOwnProperty("bundle")) {
                var error = $root.dss.DiscountBundle.verify(message.bundle);
                if (error)
                    return "bundle." + error;
            }
            return null;
        };

        /**
         * Creates a DiscountGetRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountGetRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountGetRes} DiscountGetRes
         */
        DiscountGetRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountGetRes)
                return object;
            var message = new $root.dss.DiscountGetRes();
            if (object.bundle != null) {
                if (typeof object.bundle !== "object")
                    throw TypeError(".dss.DiscountGetRes.bundle: object expected");
                message.bundle = $root.dss.DiscountBundle.fromObject(object.bundle);
            }
            return message;
        };

        /**
         * Creates a plain object from a DiscountGetRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountGetRes
         * @static
         * @param {dss.DiscountGetRes} message DiscountGetRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountGetRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.bundle = null;
            if (message.bundle != null && message.hasOwnProperty("bundle"))
                object.bundle = $root.dss.DiscountBundle.toObject(message.bundle, options);
            return object;
        };

        /**
         * Converts this DiscountGetRes to JSON.
         * @function toJSON
         * @memberof dss.DiscountGetRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountGetRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountGetRes;
    })();

    return dss;
})();

module.exports = $root;
