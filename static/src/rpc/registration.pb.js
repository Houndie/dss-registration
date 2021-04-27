/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars, strict, no-lone-blocks, default-case*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots.registration || ($protobuf.roots.registration = {});

$root.dss = (function() {

    /**
     * Namespace dss.
     * @exports dss
     * @namespace
     */
    var dss = {};

    dss.Registration = (function() {

        /**
         * Constructs a new Registration service.
         * @memberof dss
         * @classdesc Represents a Registration
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function Registration(rpcImpl, requestDelimited, responseDelimited) {
            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
        }

        (Registration.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = Registration;

        /**
         * Creates new Registration service using the specified rpc implementation.
         * @function create
         * @memberof dss.Registration
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {Registration} RPC service. Useful where requests and/or responses are streamed.
         */
        Registration.create = function create(rpcImpl, requestDelimited, responseDelimited) {
            return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link dss.Registration#add}.
         * @memberof dss.Registration
         * @typedef AddCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.RegistrationAddRes} [response] RegistrationAddRes
         */

        /**
         * Calls Add.
         * @function add
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationAddReq} request RegistrationAddReq message or plain object
         * @param {dss.Registration.AddCallback} callback Node-style callback called with the error, if any, and RegistrationAddRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Registration.prototype.add = function add(request, callback) {
            return this.rpcCall(add, $root.dss.RegistrationAddReq, $root.dss.RegistrationAddRes, request, callback);
        }, "name", { value: "Add" });

        /**
         * Calls Add.
         * @function add
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationAddReq} request RegistrationAddReq message or plain object
         * @returns {Promise<dss.RegistrationAddRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Registration#get}.
         * @memberof dss.Registration
         * @typedef GetCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.RegistrationGetRes} [response] RegistrationGetRes
         */

        /**
         * Calls Get.
         * @function get
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationGetReq} request RegistrationGetReq message or plain object
         * @param {dss.Registration.GetCallback} callback Node-style callback called with the error, if any, and RegistrationGetRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Registration.prototype.get = function get(request, callback) {
            return this.rpcCall(get, $root.dss.RegistrationGetReq, $root.dss.RegistrationGetRes, request, callback);
        }, "name", { value: "Get" });

        /**
         * Calls Get.
         * @function get
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationGetReq} request RegistrationGetReq message or plain object
         * @returns {Promise<dss.RegistrationGetRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Registration#getSummary}.
         * @memberof dss.Registration
         * @typedef GetSummaryCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.RegistrationGetSummaryRes} [response] RegistrationGetSummaryRes
         */

        /**
         * Calls GetSummary.
         * @function getSummary
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationGetSummaryReq} request RegistrationGetSummaryReq message or plain object
         * @param {dss.Registration.GetSummaryCallback} callback Node-style callback called with the error, if any, and RegistrationGetSummaryRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Registration.prototype.getSummary = function getSummary(request, callback) {
            return this.rpcCall(getSummary, $root.dss.RegistrationGetSummaryReq, $root.dss.RegistrationGetSummaryRes, request, callback);
        }, "name", { value: "GetSummary" });

        /**
         * Calls GetSummary.
         * @function getSummary
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationGetSummaryReq} request RegistrationGetSummaryReq message or plain object
         * @returns {Promise<dss.RegistrationGetSummaryRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Registration#prices}.
         * @memberof dss.Registration
         * @typedef PricesCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.RegistrationPricesRes} [response] RegistrationPricesRes
         */

        /**
         * Calls Prices.
         * @function prices
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationPricesReq} request RegistrationPricesReq message or plain object
         * @param {dss.Registration.PricesCallback} callback Node-style callback called with the error, if any, and RegistrationPricesRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Registration.prototype.prices = function prices(request, callback) {
            return this.rpcCall(prices, $root.dss.RegistrationPricesReq, $root.dss.RegistrationPricesRes, request, callback);
        }, "name", { value: "Prices" });

        /**
         * Calls Prices.
         * @function prices
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationPricesReq} request RegistrationPricesReq message or plain object
         * @returns {Promise<dss.RegistrationPricesRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Registration#update}.
         * @memberof dss.Registration
         * @typedef UpdateCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.RegistrationUpdateRes} [response] RegistrationUpdateRes
         */

        /**
         * Calls Update.
         * @function update
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationUpdateReq} request RegistrationUpdateReq message or plain object
         * @param {dss.Registration.UpdateCallback} callback Node-style callback called with the error, if any, and RegistrationUpdateRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Registration.prototype.update = function update(request, callback) {
            return this.rpcCall(update, $root.dss.RegistrationUpdateReq, $root.dss.RegistrationUpdateRes, request, callback);
        }, "name", { value: "Update" });

        /**
         * Calls Update.
         * @function update
         * @memberof dss.Registration
         * @instance
         * @param {dss.IRegistrationUpdateReq} request RegistrationUpdateReq message or plain object
         * @returns {Promise<dss.RegistrationUpdateRes>} Promise
         * @variation 2
         */

        return Registration;
    })();

    dss.RegistrationInfo = (function() {

        /**
         * Properties of a RegistrationInfo.
         * @memberof dss
         * @interface IRegistrationInfo
         * @property {string|null} [id] RegistrationInfo id
         * @property {string|null} [firstName] RegistrationInfo firstName
         * @property {string|null} [lastName] RegistrationInfo lastName
         * @property {string|null} [streetAddress] RegistrationInfo streetAddress
         * @property {string|null} [city] RegistrationInfo city
         * @property {string|null} [state] RegistrationInfo state
         * @property {string|null} [zipCode] RegistrationInfo zipCode
         * @property {string|null} [email] RegistrationInfo email
         * @property {string|null} [homeScene] RegistrationInfo homeScene
         * @property {boolean|null} [isStudent] RegistrationInfo isStudent
         * @property {dss.IFullWeekendPass|null} [fullWeekendPass] RegistrationInfo fullWeekendPass
         * @property {dss.IDanceOnlyPass|null} [danceOnlyPass] RegistrationInfo danceOnlyPass
         * @property {dss.INoPass|null} [noPass] RegistrationInfo noPass
         * @property {dss.IMixAndMatch|null} [mixAndMatch] RegistrationInfo mixAndMatch
         * @property {dss.ISoloJazz|null} [soloJazz] RegistrationInfo soloJazz
         * @property {dss.ITeamCompetition|null} [teamCompetition] RegistrationInfo teamCompetition
         * @property {dss.ITShirt|null} [tshirt] RegistrationInfo tshirt
         * @property {dss.IProvideHousing|null} [provideHousing] RegistrationInfo provideHousing
         * @property {dss.IRequireHousing|null} [requireHousing] RegistrationInfo requireHousing
         * @property {dss.INoHousing|null} [noHousing] RegistrationInfo noHousing
         * @property {Array.<string>|null} [discountCodes] RegistrationInfo discountCodes
         * @property {string|null} [createdAt] RegistrationInfo createdAt
         */

        /**
         * Constructs a new RegistrationInfo.
         * @memberof dss
         * @classdesc Represents a RegistrationInfo.
         * @implements IRegistrationInfo
         * @constructor
         * @param {dss.IRegistrationInfo=} [properties] Properties to set
         */
        function RegistrationInfo(properties) {
            this.discountCodes = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationInfo id.
         * @member {string} id
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.id = "";

        /**
         * RegistrationInfo firstName.
         * @member {string} firstName
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.firstName = "";

        /**
         * RegistrationInfo lastName.
         * @member {string} lastName
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.lastName = "";

        /**
         * RegistrationInfo streetAddress.
         * @member {string} streetAddress
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.streetAddress = "";

        /**
         * RegistrationInfo city.
         * @member {string} city
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.city = "";

        /**
         * RegistrationInfo state.
         * @member {string} state
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.state = "";

        /**
         * RegistrationInfo zipCode.
         * @member {string} zipCode
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.zipCode = "";

        /**
         * RegistrationInfo email.
         * @member {string} email
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.email = "";

        /**
         * RegistrationInfo homeScene.
         * @member {string} homeScene
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.homeScene = "";

        /**
         * RegistrationInfo isStudent.
         * @member {boolean} isStudent
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.isStudent = false;

        /**
         * RegistrationInfo fullWeekendPass.
         * @member {dss.IFullWeekendPass|null|undefined} fullWeekendPass
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.fullWeekendPass = null;

        /**
         * RegistrationInfo danceOnlyPass.
         * @member {dss.IDanceOnlyPass|null|undefined} danceOnlyPass
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.danceOnlyPass = null;

        /**
         * RegistrationInfo noPass.
         * @member {dss.INoPass|null|undefined} noPass
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.noPass = null;

        /**
         * RegistrationInfo mixAndMatch.
         * @member {dss.IMixAndMatch|null|undefined} mixAndMatch
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.mixAndMatch = null;

        /**
         * RegistrationInfo soloJazz.
         * @member {dss.ISoloJazz|null|undefined} soloJazz
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.soloJazz = null;

        /**
         * RegistrationInfo teamCompetition.
         * @member {dss.ITeamCompetition|null|undefined} teamCompetition
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.teamCompetition = null;

        /**
         * RegistrationInfo tshirt.
         * @member {dss.ITShirt|null|undefined} tshirt
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.tshirt = null;

        /**
         * RegistrationInfo provideHousing.
         * @member {dss.IProvideHousing|null|undefined} provideHousing
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.provideHousing = null;

        /**
         * RegistrationInfo requireHousing.
         * @member {dss.IRequireHousing|null|undefined} requireHousing
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.requireHousing = null;

        /**
         * RegistrationInfo noHousing.
         * @member {dss.INoHousing|null|undefined} noHousing
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.noHousing = null;

        /**
         * RegistrationInfo discountCodes.
         * @member {Array.<string>} discountCodes
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.discountCodes = $util.emptyArray;

        /**
         * RegistrationInfo createdAt.
         * @member {string} createdAt
         * @memberof dss.RegistrationInfo
         * @instance
         */
        RegistrationInfo.prototype.createdAt = "";

        // OneOf field names bound to virtual getters and setters
        var $oneOfFields;

        /**
         * RegistrationInfo passType.
         * @member {"fullWeekendPass"|"danceOnlyPass"|"noPass"|undefined} passType
         * @memberof dss.RegistrationInfo
         * @instance
         */
        Object.defineProperty(RegistrationInfo.prototype, "passType", {
            get: $util.oneOfGetter($oneOfFields = ["fullWeekendPass", "danceOnlyPass", "noPass"]),
            set: $util.oneOfSetter($oneOfFields)
        });

        /**
         * RegistrationInfo housing.
         * @member {"provideHousing"|"requireHousing"|"noHousing"|undefined} housing
         * @memberof dss.RegistrationInfo
         * @instance
         */
        Object.defineProperty(RegistrationInfo.prototype, "housing", {
            get: $util.oneOfGetter($oneOfFields = ["provideHousing", "requireHousing", "noHousing"]),
            set: $util.oneOfSetter($oneOfFields)
        });

        /**
         * Creates a new RegistrationInfo instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationInfo
         * @static
         * @param {dss.IRegistrationInfo=} [properties] Properties to set
         * @returns {dss.RegistrationInfo} RegistrationInfo instance
         */
        RegistrationInfo.create = function create(properties) {
            return new RegistrationInfo(properties);
        };

        /**
         * Encodes the specified RegistrationInfo message. Does not implicitly {@link dss.RegistrationInfo.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationInfo
         * @static
         * @param {dss.IRegistrationInfo} message RegistrationInfo message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationInfo.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            if (message.firstName != null && Object.hasOwnProperty.call(message, "firstName"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.firstName);
            if (message.lastName != null && Object.hasOwnProperty.call(message, "lastName"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.lastName);
            if (message.streetAddress != null && Object.hasOwnProperty.call(message, "streetAddress"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.streetAddress);
            if (message.city != null && Object.hasOwnProperty.call(message, "city"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.city);
            if (message.state != null && Object.hasOwnProperty.call(message, "state"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.state);
            if (message.zipCode != null && Object.hasOwnProperty.call(message, "zipCode"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.zipCode);
            if (message.email != null && Object.hasOwnProperty.call(message, "email"))
                writer.uint32(/* id 8, wireType 2 =*/66).string(message.email);
            if (message.homeScene != null && Object.hasOwnProperty.call(message, "homeScene"))
                writer.uint32(/* id 9, wireType 2 =*/74).string(message.homeScene);
            if (message.isStudent != null && Object.hasOwnProperty.call(message, "isStudent"))
                writer.uint32(/* id 10, wireType 0 =*/80).bool(message.isStudent);
            if (message.fullWeekendPass != null && Object.hasOwnProperty.call(message, "fullWeekendPass"))
                $root.dss.FullWeekendPass.encode(message.fullWeekendPass, writer.uint32(/* id 11, wireType 2 =*/90).fork()).ldelim();
            if (message.danceOnlyPass != null && Object.hasOwnProperty.call(message, "danceOnlyPass"))
                $root.dss.DanceOnlyPass.encode(message.danceOnlyPass, writer.uint32(/* id 12, wireType 2 =*/98).fork()).ldelim();
            if (message.noPass != null && Object.hasOwnProperty.call(message, "noPass"))
                $root.dss.NoPass.encode(message.noPass, writer.uint32(/* id 13, wireType 2 =*/106).fork()).ldelim();
            if (message.mixAndMatch != null && Object.hasOwnProperty.call(message, "mixAndMatch"))
                $root.dss.MixAndMatch.encode(message.mixAndMatch, writer.uint32(/* id 14, wireType 2 =*/114).fork()).ldelim();
            if (message.soloJazz != null && Object.hasOwnProperty.call(message, "soloJazz"))
                $root.dss.SoloJazz.encode(message.soloJazz, writer.uint32(/* id 15, wireType 2 =*/122).fork()).ldelim();
            if (message.teamCompetition != null && Object.hasOwnProperty.call(message, "teamCompetition"))
                $root.dss.TeamCompetition.encode(message.teamCompetition, writer.uint32(/* id 16, wireType 2 =*/130).fork()).ldelim();
            if (message.tshirt != null && Object.hasOwnProperty.call(message, "tshirt"))
                $root.dss.TShirt.encode(message.tshirt, writer.uint32(/* id 17, wireType 2 =*/138).fork()).ldelim();
            if (message.provideHousing != null && Object.hasOwnProperty.call(message, "provideHousing"))
                $root.dss.ProvideHousing.encode(message.provideHousing, writer.uint32(/* id 18, wireType 2 =*/146).fork()).ldelim();
            if (message.requireHousing != null && Object.hasOwnProperty.call(message, "requireHousing"))
                $root.dss.RequireHousing.encode(message.requireHousing, writer.uint32(/* id 19, wireType 2 =*/154).fork()).ldelim();
            if (message.noHousing != null && Object.hasOwnProperty.call(message, "noHousing"))
                $root.dss.NoHousing.encode(message.noHousing, writer.uint32(/* id 20, wireType 2 =*/162).fork()).ldelim();
            if (message.discountCodes != null && message.discountCodes.length)
                for (var i = 0; i < message.discountCodes.length; ++i)
                    writer.uint32(/* id 21, wireType 2 =*/170).string(message.discountCodes[i]);
            if (message.createdAt != null && Object.hasOwnProperty.call(message, "createdAt"))
                writer.uint32(/* id 22, wireType 2 =*/178).string(message.createdAt);
            return writer;
        };

        /**
         * Encodes the specified RegistrationInfo message, length delimited. Does not implicitly {@link dss.RegistrationInfo.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationInfo
         * @static
         * @param {dss.IRegistrationInfo} message RegistrationInfo message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationInfo.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationInfo message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationInfo
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationInfo} RegistrationInfo
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationInfo.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationInfo();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.firstName = reader.string();
                    break;
                case 3:
                    message.lastName = reader.string();
                    break;
                case 4:
                    message.streetAddress = reader.string();
                    break;
                case 5:
                    message.city = reader.string();
                    break;
                case 6:
                    message.state = reader.string();
                    break;
                case 7:
                    message.zipCode = reader.string();
                    break;
                case 8:
                    message.email = reader.string();
                    break;
                case 9:
                    message.homeScene = reader.string();
                    break;
                case 10:
                    message.isStudent = reader.bool();
                    break;
                case 11:
                    message.fullWeekendPass = $root.dss.FullWeekendPass.decode(reader, reader.uint32());
                    break;
                case 12:
                    message.danceOnlyPass = $root.dss.DanceOnlyPass.decode(reader, reader.uint32());
                    break;
                case 13:
                    message.noPass = $root.dss.NoPass.decode(reader, reader.uint32());
                    break;
                case 14:
                    message.mixAndMatch = $root.dss.MixAndMatch.decode(reader, reader.uint32());
                    break;
                case 15:
                    message.soloJazz = $root.dss.SoloJazz.decode(reader, reader.uint32());
                    break;
                case 16:
                    message.teamCompetition = $root.dss.TeamCompetition.decode(reader, reader.uint32());
                    break;
                case 17:
                    message.tshirt = $root.dss.TShirt.decode(reader, reader.uint32());
                    break;
                case 18:
                    message.provideHousing = $root.dss.ProvideHousing.decode(reader, reader.uint32());
                    break;
                case 19:
                    message.requireHousing = $root.dss.RequireHousing.decode(reader, reader.uint32());
                    break;
                case 20:
                    message.noHousing = $root.dss.NoHousing.decode(reader, reader.uint32());
                    break;
                case 21:
                    if (!(message.discountCodes && message.discountCodes.length))
                        message.discountCodes = [];
                    message.discountCodes.push(reader.string());
                    break;
                case 22:
                    message.createdAt = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationInfo message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationInfo
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationInfo} RegistrationInfo
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationInfo.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationInfo message.
         * @function verify
         * @memberof dss.RegistrationInfo
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationInfo.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            var properties = {};
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            if (message.firstName != null && message.hasOwnProperty("firstName"))
                if (!$util.isString(message.firstName))
                    return "firstName: string expected";
            if (message.lastName != null && message.hasOwnProperty("lastName"))
                if (!$util.isString(message.lastName))
                    return "lastName: string expected";
            if (message.streetAddress != null && message.hasOwnProperty("streetAddress"))
                if (!$util.isString(message.streetAddress))
                    return "streetAddress: string expected";
            if (message.city != null && message.hasOwnProperty("city"))
                if (!$util.isString(message.city))
                    return "city: string expected";
            if (message.state != null && message.hasOwnProperty("state"))
                if (!$util.isString(message.state))
                    return "state: string expected";
            if (message.zipCode != null && message.hasOwnProperty("zipCode"))
                if (!$util.isString(message.zipCode))
                    return "zipCode: string expected";
            if (message.email != null && message.hasOwnProperty("email"))
                if (!$util.isString(message.email))
                    return "email: string expected";
            if (message.homeScene != null && message.hasOwnProperty("homeScene"))
                if (!$util.isString(message.homeScene))
                    return "homeScene: string expected";
            if (message.isStudent != null && message.hasOwnProperty("isStudent"))
                if (typeof message.isStudent !== "boolean")
                    return "isStudent: boolean expected";
            if (message.fullWeekendPass != null && message.hasOwnProperty("fullWeekendPass")) {
                properties.passType = 1;
                {
                    var error = $root.dss.FullWeekendPass.verify(message.fullWeekendPass);
                    if (error)
                        return "fullWeekendPass." + error;
                }
            }
            if (message.danceOnlyPass != null && message.hasOwnProperty("danceOnlyPass")) {
                if (properties.passType === 1)
                    return "passType: multiple values";
                properties.passType = 1;
                {
                    var error = $root.dss.DanceOnlyPass.verify(message.danceOnlyPass);
                    if (error)
                        return "danceOnlyPass." + error;
                }
            }
            if (message.noPass != null && message.hasOwnProperty("noPass")) {
                if (properties.passType === 1)
                    return "passType: multiple values";
                properties.passType = 1;
                {
                    var error = $root.dss.NoPass.verify(message.noPass);
                    if (error)
                        return "noPass." + error;
                }
            }
            if (message.mixAndMatch != null && message.hasOwnProperty("mixAndMatch")) {
                var error = $root.dss.MixAndMatch.verify(message.mixAndMatch);
                if (error)
                    return "mixAndMatch." + error;
            }
            if (message.soloJazz != null && message.hasOwnProperty("soloJazz")) {
                var error = $root.dss.SoloJazz.verify(message.soloJazz);
                if (error)
                    return "soloJazz." + error;
            }
            if (message.teamCompetition != null && message.hasOwnProperty("teamCompetition")) {
                var error = $root.dss.TeamCompetition.verify(message.teamCompetition);
                if (error)
                    return "teamCompetition." + error;
            }
            if (message.tshirt != null && message.hasOwnProperty("tshirt")) {
                var error = $root.dss.TShirt.verify(message.tshirt);
                if (error)
                    return "tshirt." + error;
            }
            if (message.provideHousing != null && message.hasOwnProperty("provideHousing")) {
                properties.housing = 1;
                {
                    var error = $root.dss.ProvideHousing.verify(message.provideHousing);
                    if (error)
                        return "provideHousing." + error;
                }
            }
            if (message.requireHousing != null && message.hasOwnProperty("requireHousing")) {
                if (properties.housing === 1)
                    return "housing: multiple values";
                properties.housing = 1;
                {
                    var error = $root.dss.RequireHousing.verify(message.requireHousing);
                    if (error)
                        return "requireHousing." + error;
                }
            }
            if (message.noHousing != null && message.hasOwnProperty("noHousing")) {
                if (properties.housing === 1)
                    return "housing: multiple values";
                properties.housing = 1;
                {
                    var error = $root.dss.NoHousing.verify(message.noHousing);
                    if (error)
                        return "noHousing." + error;
                }
            }
            if (message.discountCodes != null && message.hasOwnProperty("discountCodes")) {
                if (!Array.isArray(message.discountCodes))
                    return "discountCodes: array expected";
                for (var i = 0; i < message.discountCodes.length; ++i)
                    if (!$util.isString(message.discountCodes[i]))
                        return "discountCodes: string[] expected";
            }
            if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                if (!$util.isString(message.createdAt))
                    return "createdAt: string expected";
            return null;
        };

        /**
         * Creates a RegistrationInfo message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationInfo
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationInfo} RegistrationInfo
         */
        RegistrationInfo.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationInfo)
                return object;
            var message = new $root.dss.RegistrationInfo();
            if (object.id != null)
                message.id = String(object.id);
            if (object.firstName != null)
                message.firstName = String(object.firstName);
            if (object.lastName != null)
                message.lastName = String(object.lastName);
            if (object.streetAddress != null)
                message.streetAddress = String(object.streetAddress);
            if (object.city != null)
                message.city = String(object.city);
            if (object.state != null)
                message.state = String(object.state);
            if (object.zipCode != null)
                message.zipCode = String(object.zipCode);
            if (object.email != null)
                message.email = String(object.email);
            if (object.homeScene != null)
                message.homeScene = String(object.homeScene);
            if (object.isStudent != null)
                message.isStudent = Boolean(object.isStudent);
            if (object.fullWeekendPass != null) {
                if (typeof object.fullWeekendPass !== "object")
                    throw TypeError(".dss.RegistrationInfo.fullWeekendPass: object expected");
                message.fullWeekendPass = $root.dss.FullWeekendPass.fromObject(object.fullWeekendPass);
            }
            if (object.danceOnlyPass != null) {
                if (typeof object.danceOnlyPass !== "object")
                    throw TypeError(".dss.RegistrationInfo.danceOnlyPass: object expected");
                message.danceOnlyPass = $root.dss.DanceOnlyPass.fromObject(object.danceOnlyPass);
            }
            if (object.noPass != null) {
                if (typeof object.noPass !== "object")
                    throw TypeError(".dss.RegistrationInfo.noPass: object expected");
                message.noPass = $root.dss.NoPass.fromObject(object.noPass);
            }
            if (object.mixAndMatch != null) {
                if (typeof object.mixAndMatch !== "object")
                    throw TypeError(".dss.RegistrationInfo.mixAndMatch: object expected");
                message.mixAndMatch = $root.dss.MixAndMatch.fromObject(object.mixAndMatch);
            }
            if (object.soloJazz != null) {
                if (typeof object.soloJazz !== "object")
                    throw TypeError(".dss.RegistrationInfo.soloJazz: object expected");
                message.soloJazz = $root.dss.SoloJazz.fromObject(object.soloJazz);
            }
            if (object.teamCompetition != null) {
                if (typeof object.teamCompetition !== "object")
                    throw TypeError(".dss.RegistrationInfo.teamCompetition: object expected");
                message.teamCompetition = $root.dss.TeamCompetition.fromObject(object.teamCompetition);
            }
            if (object.tshirt != null) {
                if (typeof object.tshirt !== "object")
                    throw TypeError(".dss.RegistrationInfo.tshirt: object expected");
                message.tshirt = $root.dss.TShirt.fromObject(object.tshirt);
            }
            if (object.provideHousing != null) {
                if (typeof object.provideHousing !== "object")
                    throw TypeError(".dss.RegistrationInfo.provideHousing: object expected");
                message.provideHousing = $root.dss.ProvideHousing.fromObject(object.provideHousing);
            }
            if (object.requireHousing != null) {
                if (typeof object.requireHousing !== "object")
                    throw TypeError(".dss.RegistrationInfo.requireHousing: object expected");
                message.requireHousing = $root.dss.RequireHousing.fromObject(object.requireHousing);
            }
            if (object.noHousing != null) {
                if (typeof object.noHousing !== "object")
                    throw TypeError(".dss.RegistrationInfo.noHousing: object expected");
                message.noHousing = $root.dss.NoHousing.fromObject(object.noHousing);
            }
            if (object.discountCodes) {
                if (!Array.isArray(object.discountCodes))
                    throw TypeError(".dss.RegistrationInfo.discountCodes: array expected");
                message.discountCodes = [];
                for (var i = 0; i < object.discountCodes.length; ++i)
                    message.discountCodes[i] = String(object.discountCodes[i]);
            }
            if (object.createdAt != null)
                message.createdAt = String(object.createdAt);
            return message;
        };

        /**
         * Creates a plain object from a RegistrationInfo message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationInfo
         * @static
         * @param {dss.RegistrationInfo} message RegistrationInfo
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationInfo.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.discountCodes = [];
            if (options.defaults) {
                object.id = "";
                object.firstName = "";
                object.lastName = "";
                object.streetAddress = "";
                object.city = "";
                object.state = "";
                object.zipCode = "";
                object.email = "";
                object.homeScene = "";
                object.isStudent = false;
                object.mixAndMatch = null;
                object.soloJazz = null;
                object.teamCompetition = null;
                object.tshirt = null;
                object.createdAt = "";
            }
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.firstName != null && message.hasOwnProperty("firstName"))
                object.firstName = message.firstName;
            if (message.lastName != null && message.hasOwnProperty("lastName"))
                object.lastName = message.lastName;
            if (message.streetAddress != null && message.hasOwnProperty("streetAddress"))
                object.streetAddress = message.streetAddress;
            if (message.city != null && message.hasOwnProperty("city"))
                object.city = message.city;
            if (message.state != null && message.hasOwnProperty("state"))
                object.state = message.state;
            if (message.zipCode != null && message.hasOwnProperty("zipCode"))
                object.zipCode = message.zipCode;
            if (message.email != null && message.hasOwnProperty("email"))
                object.email = message.email;
            if (message.homeScene != null && message.hasOwnProperty("homeScene"))
                object.homeScene = message.homeScene;
            if (message.isStudent != null && message.hasOwnProperty("isStudent"))
                object.isStudent = message.isStudent;
            if (message.fullWeekendPass != null && message.hasOwnProperty("fullWeekendPass")) {
                object.fullWeekendPass = $root.dss.FullWeekendPass.toObject(message.fullWeekendPass, options);
                if (options.oneofs)
                    object.passType = "fullWeekendPass";
            }
            if (message.danceOnlyPass != null && message.hasOwnProperty("danceOnlyPass")) {
                object.danceOnlyPass = $root.dss.DanceOnlyPass.toObject(message.danceOnlyPass, options);
                if (options.oneofs)
                    object.passType = "danceOnlyPass";
            }
            if (message.noPass != null && message.hasOwnProperty("noPass")) {
                object.noPass = $root.dss.NoPass.toObject(message.noPass, options);
                if (options.oneofs)
                    object.passType = "noPass";
            }
            if (message.mixAndMatch != null && message.hasOwnProperty("mixAndMatch"))
                object.mixAndMatch = $root.dss.MixAndMatch.toObject(message.mixAndMatch, options);
            if (message.soloJazz != null && message.hasOwnProperty("soloJazz"))
                object.soloJazz = $root.dss.SoloJazz.toObject(message.soloJazz, options);
            if (message.teamCompetition != null && message.hasOwnProperty("teamCompetition"))
                object.teamCompetition = $root.dss.TeamCompetition.toObject(message.teamCompetition, options);
            if (message.tshirt != null && message.hasOwnProperty("tshirt"))
                object.tshirt = $root.dss.TShirt.toObject(message.tshirt, options);
            if (message.provideHousing != null && message.hasOwnProperty("provideHousing")) {
                object.provideHousing = $root.dss.ProvideHousing.toObject(message.provideHousing, options);
                if (options.oneofs)
                    object.housing = "provideHousing";
            }
            if (message.requireHousing != null && message.hasOwnProperty("requireHousing")) {
                object.requireHousing = $root.dss.RequireHousing.toObject(message.requireHousing, options);
                if (options.oneofs)
                    object.housing = "requireHousing";
            }
            if (message.noHousing != null && message.hasOwnProperty("noHousing")) {
                object.noHousing = $root.dss.NoHousing.toObject(message.noHousing, options);
                if (options.oneofs)
                    object.housing = "noHousing";
            }
            if (message.discountCodes && message.discountCodes.length) {
                object.discountCodes = [];
                for (var j = 0; j < message.discountCodes.length; ++j)
                    object.discountCodes[j] = message.discountCodes[j];
            }
            if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                object.createdAt = message.createdAt;
            return object;
        };

        /**
         * Converts this RegistrationInfo to JSON.
         * @function toJSON
         * @memberof dss.RegistrationInfo
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationInfo.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationInfo;
    })();

    /**
     * FullWeekendPassTier enum.
     * @name dss.FullWeekendPassTier
     * @enum {number}
     * @property {number} Tier1=0 Tier1 value
     * @property {number} Tier2=1 Tier2 value
     * @property {number} Tier3=2 Tier3 value
     * @property {number} Tier4=3 Tier4 value
     * @property {number} Tier5=4 Tier5 value
     */
    dss.FullWeekendPassTier = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Tier1"] = 0;
        values[valuesById[1] = "Tier2"] = 1;
        values[valuesById[2] = "Tier3"] = 2;
        values[valuesById[3] = "Tier4"] = 3;
        values[valuesById[4] = "Tier5"] = 4;
        return values;
    })();

    /**
     * FullWeekendPassLevel enum.
     * @name dss.FullWeekendPassLevel
     * @enum {number}
     * @property {number} Level1=0 Level1 value
     * @property {number} Level2=1 Level2 value
     * @property {number} Level3=2 Level3 value
     */
    dss.FullWeekendPassLevel = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "Level1"] = 0;
        values[valuesById[1] = "Level2"] = 1;
        values[valuesById[2] = "Level3"] = 2;
        return values;
    })();

    dss.FullWeekendPass = (function() {

        /**
         * Properties of a FullWeekendPass.
         * @memberof dss
         * @interface IFullWeekendPass
         * @property {dss.FullWeekendPassTier|null} [tier] FullWeekendPass tier
         * @property {dss.FullWeekendPassLevel|null} [level] FullWeekendPass level
         * @property {boolean|null} [paid] FullWeekendPass paid
         */

        /**
         * Constructs a new FullWeekendPass.
         * @memberof dss
         * @classdesc Represents a FullWeekendPass.
         * @implements IFullWeekendPass
         * @constructor
         * @param {dss.IFullWeekendPass=} [properties] Properties to set
         */
        function FullWeekendPass(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * FullWeekendPass tier.
         * @member {dss.FullWeekendPassTier} tier
         * @memberof dss.FullWeekendPass
         * @instance
         */
        FullWeekendPass.prototype.tier = 0;

        /**
         * FullWeekendPass level.
         * @member {dss.FullWeekendPassLevel} level
         * @memberof dss.FullWeekendPass
         * @instance
         */
        FullWeekendPass.prototype.level = 0;

        /**
         * FullWeekendPass paid.
         * @member {boolean} paid
         * @memberof dss.FullWeekendPass
         * @instance
         */
        FullWeekendPass.prototype.paid = false;

        /**
         * Creates a new FullWeekendPass instance using the specified properties.
         * @function create
         * @memberof dss.FullWeekendPass
         * @static
         * @param {dss.IFullWeekendPass=} [properties] Properties to set
         * @returns {dss.FullWeekendPass} FullWeekendPass instance
         */
        FullWeekendPass.create = function create(properties) {
            return new FullWeekendPass(properties);
        };

        /**
         * Encodes the specified FullWeekendPass message. Does not implicitly {@link dss.FullWeekendPass.verify|verify} messages.
         * @function encode
         * @memberof dss.FullWeekendPass
         * @static
         * @param {dss.IFullWeekendPass} message FullWeekendPass message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FullWeekendPass.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.tier != null && Object.hasOwnProperty.call(message, "tier"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.tier);
            if (message.level != null && Object.hasOwnProperty.call(message, "level"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.level);
            if (message.paid != null && Object.hasOwnProperty.call(message, "paid"))
                writer.uint32(/* id 3, wireType 0 =*/24).bool(message.paid);
            return writer;
        };

        /**
         * Encodes the specified FullWeekendPass message, length delimited. Does not implicitly {@link dss.FullWeekendPass.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.FullWeekendPass
         * @static
         * @param {dss.IFullWeekendPass} message FullWeekendPass message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FullWeekendPass.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a FullWeekendPass message from the specified reader or buffer.
         * @function decode
         * @memberof dss.FullWeekendPass
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.FullWeekendPass} FullWeekendPass
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FullWeekendPass.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.FullWeekendPass();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.tier = reader.int32();
                    break;
                case 2:
                    message.level = reader.int32();
                    break;
                case 3:
                    message.paid = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a FullWeekendPass message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.FullWeekendPass
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.FullWeekendPass} FullWeekendPass
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FullWeekendPass.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a FullWeekendPass message.
         * @function verify
         * @memberof dss.FullWeekendPass
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        FullWeekendPass.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.tier != null && message.hasOwnProperty("tier"))
                switch (message.tier) {
                default:
                    return "tier: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                    break;
                }
            if (message.level != null && message.hasOwnProperty("level"))
                switch (message.level) {
                default:
                    return "level: enum value expected";
                case 0:
                case 1:
                case 2:
                    break;
                }
            if (message.paid != null && message.hasOwnProperty("paid"))
                if (typeof message.paid !== "boolean")
                    return "paid: boolean expected";
            return null;
        };

        /**
         * Creates a FullWeekendPass message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.FullWeekendPass
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.FullWeekendPass} FullWeekendPass
         */
        FullWeekendPass.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.FullWeekendPass)
                return object;
            var message = new $root.dss.FullWeekendPass();
            switch (object.tier) {
            case "Tier1":
            case 0:
                message.tier = 0;
                break;
            case "Tier2":
            case 1:
                message.tier = 1;
                break;
            case "Tier3":
            case 2:
                message.tier = 2;
                break;
            case "Tier4":
            case 3:
                message.tier = 3;
                break;
            case "Tier5":
            case 4:
                message.tier = 4;
                break;
            }
            switch (object.level) {
            case "Level1":
            case 0:
                message.level = 0;
                break;
            case "Level2":
            case 1:
                message.level = 1;
                break;
            case "Level3":
            case 2:
                message.level = 2;
                break;
            }
            if (object.paid != null)
                message.paid = Boolean(object.paid);
            return message;
        };

        /**
         * Creates a plain object from a FullWeekendPass message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.FullWeekendPass
         * @static
         * @param {dss.FullWeekendPass} message FullWeekendPass
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        FullWeekendPass.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.tier = options.enums === String ? "Tier1" : 0;
                object.level = options.enums === String ? "Level1" : 0;
                object.paid = false;
            }
            if (message.tier != null && message.hasOwnProperty("tier"))
                object.tier = options.enums === String ? $root.dss.FullWeekendPassTier[message.tier] : message.tier;
            if (message.level != null && message.hasOwnProperty("level"))
                object.level = options.enums === String ? $root.dss.FullWeekendPassLevel[message.level] : message.level;
            if (message.paid != null && message.hasOwnProperty("paid"))
                object.paid = message.paid;
            return object;
        };

        /**
         * Converts this FullWeekendPass to JSON.
         * @function toJSON
         * @memberof dss.FullWeekendPass
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        FullWeekendPass.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return FullWeekendPass;
    })();

    dss.DanceOnlyPass = (function() {

        /**
         * Properties of a DanceOnlyPass.
         * @memberof dss
         * @interface IDanceOnlyPass
         * @property {boolean|null} [paid] DanceOnlyPass paid
         */

        /**
         * Constructs a new DanceOnlyPass.
         * @memberof dss
         * @classdesc Represents a DanceOnlyPass.
         * @implements IDanceOnlyPass
         * @constructor
         * @param {dss.IDanceOnlyPass=} [properties] Properties to set
         */
        function DanceOnlyPass(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DanceOnlyPass paid.
         * @member {boolean} paid
         * @memberof dss.DanceOnlyPass
         * @instance
         */
        DanceOnlyPass.prototype.paid = false;

        /**
         * Creates a new DanceOnlyPass instance using the specified properties.
         * @function create
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {dss.IDanceOnlyPass=} [properties] Properties to set
         * @returns {dss.DanceOnlyPass} DanceOnlyPass instance
         */
        DanceOnlyPass.create = function create(properties) {
            return new DanceOnlyPass(properties);
        };

        /**
         * Encodes the specified DanceOnlyPass message. Does not implicitly {@link dss.DanceOnlyPass.verify|verify} messages.
         * @function encode
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {dss.IDanceOnlyPass} message DanceOnlyPass message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DanceOnlyPass.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.paid != null && Object.hasOwnProperty.call(message, "paid"))
                writer.uint32(/* id 1, wireType 0 =*/8).bool(message.paid);
            return writer;
        };

        /**
         * Encodes the specified DanceOnlyPass message, length delimited. Does not implicitly {@link dss.DanceOnlyPass.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {dss.IDanceOnlyPass} message DanceOnlyPass message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DanceOnlyPass.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DanceOnlyPass message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DanceOnlyPass} DanceOnlyPass
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DanceOnlyPass.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DanceOnlyPass();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.paid = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DanceOnlyPass message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DanceOnlyPass} DanceOnlyPass
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DanceOnlyPass.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DanceOnlyPass message.
         * @function verify
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DanceOnlyPass.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.paid != null && message.hasOwnProperty("paid"))
                if (typeof message.paid !== "boolean")
                    return "paid: boolean expected";
            return null;
        };

        /**
         * Creates a DanceOnlyPass message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DanceOnlyPass} DanceOnlyPass
         */
        DanceOnlyPass.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DanceOnlyPass)
                return object;
            var message = new $root.dss.DanceOnlyPass();
            if (object.paid != null)
                message.paid = Boolean(object.paid);
            return message;
        };

        /**
         * Creates a plain object from a DanceOnlyPass message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DanceOnlyPass
         * @static
         * @param {dss.DanceOnlyPass} message DanceOnlyPass
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DanceOnlyPass.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.paid = false;
            if (message.paid != null && message.hasOwnProperty("paid"))
                object.paid = message.paid;
            return object;
        };

        /**
         * Converts this DanceOnlyPass to JSON.
         * @function toJSON
         * @memberof dss.DanceOnlyPass
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DanceOnlyPass.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DanceOnlyPass;
    })();

    dss.NoPass = (function() {

        /**
         * Properties of a NoPass.
         * @memberof dss
         * @interface INoPass
         */

        /**
         * Constructs a new NoPass.
         * @memberof dss
         * @classdesc Represents a NoPass.
         * @implements INoPass
         * @constructor
         * @param {dss.INoPass=} [properties] Properties to set
         */
        function NoPass(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new NoPass instance using the specified properties.
         * @function create
         * @memberof dss.NoPass
         * @static
         * @param {dss.INoPass=} [properties] Properties to set
         * @returns {dss.NoPass} NoPass instance
         */
        NoPass.create = function create(properties) {
            return new NoPass(properties);
        };

        /**
         * Encodes the specified NoPass message. Does not implicitly {@link dss.NoPass.verify|verify} messages.
         * @function encode
         * @memberof dss.NoPass
         * @static
         * @param {dss.INoPass} message NoPass message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        NoPass.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified NoPass message, length delimited. Does not implicitly {@link dss.NoPass.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.NoPass
         * @static
         * @param {dss.INoPass} message NoPass message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        NoPass.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a NoPass message from the specified reader or buffer.
         * @function decode
         * @memberof dss.NoPass
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.NoPass} NoPass
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        NoPass.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.NoPass();
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
         * Decodes a NoPass message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.NoPass
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.NoPass} NoPass
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        NoPass.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a NoPass message.
         * @function verify
         * @memberof dss.NoPass
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        NoPass.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a NoPass message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.NoPass
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.NoPass} NoPass
         */
        NoPass.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.NoPass)
                return object;
            return new $root.dss.NoPass();
        };

        /**
         * Creates a plain object from a NoPass message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.NoPass
         * @static
         * @param {dss.NoPass} message NoPass
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        NoPass.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this NoPass to JSON.
         * @function toJSON
         * @memberof dss.NoPass
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        NoPass.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return NoPass;
    })();

    dss.MixAndMatch = (function() {

        /**
         * Properties of a MixAndMatch.
         * @memberof dss
         * @interface IMixAndMatch
         * @property {dss.MixAndMatch.Role|null} [role] MixAndMatch role
         * @property {boolean|null} [paid] MixAndMatch paid
         */

        /**
         * Constructs a new MixAndMatch.
         * @memberof dss
         * @classdesc Represents a MixAndMatch.
         * @implements IMixAndMatch
         * @constructor
         * @param {dss.IMixAndMatch=} [properties] Properties to set
         */
        function MixAndMatch(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * MixAndMatch role.
         * @member {dss.MixAndMatch.Role} role
         * @memberof dss.MixAndMatch
         * @instance
         */
        MixAndMatch.prototype.role = 0;

        /**
         * MixAndMatch paid.
         * @member {boolean} paid
         * @memberof dss.MixAndMatch
         * @instance
         */
        MixAndMatch.prototype.paid = false;

        /**
         * Creates a new MixAndMatch instance using the specified properties.
         * @function create
         * @memberof dss.MixAndMatch
         * @static
         * @param {dss.IMixAndMatch=} [properties] Properties to set
         * @returns {dss.MixAndMatch} MixAndMatch instance
         */
        MixAndMatch.create = function create(properties) {
            return new MixAndMatch(properties);
        };

        /**
         * Encodes the specified MixAndMatch message. Does not implicitly {@link dss.MixAndMatch.verify|verify} messages.
         * @function encode
         * @memberof dss.MixAndMatch
         * @static
         * @param {dss.IMixAndMatch} message MixAndMatch message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        MixAndMatch.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.role != null && Object.hasOwnProperty.call(message, "role"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.role);
            if (message.paid != null && Object.hasOwnProperty.call(message, "paid"))
                writer.uint32(/* id 2, wireType 0 =*/16).bool(message.paid);
            return writer;
        };

        /**
         * Encodes the specified MixAndMatch message, length delimited. Does not implicitly {@link dss.MixAndMatch.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.MixAndMatch
         * @static
         * @param {dss.IMixAndMatch} message MixAndMatch message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        MixAndMatch.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a MixAndMatch message from the specified reader or buffer.
         * @function decode
         * @memberof dss.MixAndMatch
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.MixAndMatch} MixAndMatch
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        MixAndMatch.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.MixAndMatch();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.role = reader.int32();
                    break;
                case 2:
                    message.paid = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a MixAndMatch message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.MixAndMatch
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.MixAndMatch} MixAndMatch
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        MixAndMatch.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a MixAndMatch message.
         * @function verify
         * @memberof dss.MixAndMatch
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        MixAndMatch.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.role != null && message.hasOwnProperty("role"))
                switch (message.role) {
                default:
                    return "role: enum value expected";
                case 0:
                case 1:
                    break;
                }
            if (message.paid != null && message.hasOwnProperty("paid"))
                if (typeof message.paid !== "boolean")
                    return "paid: boolean expected";
            return null;
        };

        /**
         * Creates a MixAndMatch message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.MixAndMatch
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.MixAndMatch} MixAndMatch
         */
        MixAndMatch.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.MixAndMatch)
                return object;
            var message = new $root.dss.MixAndMatch();
            switch (object.role) {
            case "Follower":
            case 0:
                message.role = 0;
                break;
            case "Leader":
            case 1:
                message.role = 1;
                break;
            }
            if (object.paid != null)
                message.paid = Boolean(object.paid);
            return message;
        };

        /**
         * Creates a plain object from a MixAndMatch message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.MixAndMatch
         * @static
         * @param {dss.MixAndMatch} message MixAndMatch
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        MixAndMatch.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.role = options.enums === String ? "Follower" : 0;
                object.paid = false;
            }
            if (message.role != null && message.hasOwnProperty("role"))
                object.role = options.enums === String ? $root.dss.MixAndMatch.Role[message.role] : message.role;
            if (message.paid != null && message.hasOwnProperty("paid"))
                object.paid = message.paid;
            return object;
        };

        /**
         * Converts this MixAndMatch to JSON.
         * @function toJSON
         * @memberof dss.MixAndMatch
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        MixAndMatch.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Role enum.
         * @name dss.MixAndMatch.Role
         * @enum {number}
         * @property {number} Follower=0 Follower value
         * @property {number} Leader=1 Leader value
         */
        MixAndMatch.Role = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "Follower"] = 0;
            values[valuesById[1] = "Leader"] = 1;
            return values;
        })();

        return MixAndMatch;
    })();

    dss.SoloJazz = (function() {

        /**
         * Properties of a SoloJazz.
         * @memberof dss
         * @interface ISoloJazz
         * @property {boolean|null} [paid] SoloJazz paid
         */

        /**
         * Constructs a new SoloJazz.
         * @memberof dss
         * @classdesc Represents a SoloJazz.
         * @implements ISoloJazz
         * @constructor
         * @param {dss.ISoloJazz=} [properties] Properties to set
         */
        function SoloJazz(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * SoloJazz paid.
         * @member {boolean} paid
         * @memberof dss.SoloJazz
         * @instance
         */
        SoloJazz.prototype.paid = false;

        /**
         * Creates a new SoloJazz instance using the specified properties.
         * @function create
         * @memberof dss.SoloJazz
         * @static
         * @param {dss.ISoloJazz=} [properties] Properties to set
         * @returns {dss.SoloJazz} SoloJazz instance
         */
        SoloJazz.create = function create(properties) {
            return new SoloJazz(properties);
        };

        /**
         * Encodes the specified SoloJazz message. Does not implicitly {@link dss.SoloJazz.verify|verify} messages.
         * @function encode
         * @memberof dss.SoloJazz
         * @static
         * @param {dss.ISoloJazz} message SoloJazz message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SoloJazz.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.paid != null && Object.hasOwnProperty.call(message, "paid"))
                writer.uint32(/* id 1, wireType 0 =*/8).bool(message.paid);
            return writer;
        };

        /**
         * Encodes the specified SoloJazz message, length delimited. Does not implicitly {@link dss.SoloJazz.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.SoloJazz
         * @static
         * @param {dss.ISoloJazz} message SoloJazz message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        SoloJazz.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a SoloJazz message from the specified reader or buffer.
         * @function decode
         * @memberof dss.SoloJazz
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.SoloJazz} SoloJazz
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SoloJazz.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.SoloJazz();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.paid = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a SoloJazz message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.SoloJazz
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.SoloJazz} SoloJazz
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        SoloJazz.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a SoloJazz message.
         * @function verify
         * @memberof dss.SoloJazz
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        SoloJazz.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.paid != null && message.hasOwnProperty("paid"))
                if (typeof message.paid !== "boolean")
                    return "paid: boolean expected";
            return null;
        };

        /**
         * Creates a SoloJazz message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.SoloJazz
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.SoloJazz} SoloJazz
         */
        SoloJazz.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.SoloJazz)
                return object;
            var message = new $root.dss.SoloJazz();
            if (object.paid != null)
                message.paid = Boolean(object.paid);
            return message;
        };

        /**
         * Creates a plain object from a SoloJazz message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.SoloJazz
         * @static
         * @param {dss.SoloJazz} message SoloJazz
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        SoloJazz.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.paid = false;
            if (message.paid != null && message.hasOwnProperty("paid"))
                object.paid = message.paid;
            return object;
        };

        /**
         * Converts this SoloJazz to JSON.
         * @function toJSON
         * @memberof dss.SoloJazz
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        SoloJazz.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return SoloJazz;
    })();

    dss.TeamCompetition = (function() {

        /**
         * Properties of a TeamCompetition.
         * @memberof dss
         * @interface ITeamCompetition
         * @property {string|null} [name] TeamCompetition name
         * @property {boolean|null} [paid] TeamCompetition paid
         */

        /**
         * Constructs a new TeamCompetition.
         * @memberof dss
         * @classdesc Represents a TeamCompetition.
         * @implements ITeamCompetition
         * @constructor
         * @param {dss.ITeamCompetition=} [properties] Properties to set
         */
        function TeamCompetition(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * TeamCompetition name.
         * @member {string} name
         * @memberof dss.TeamCompetition
         * @instance
         */
        TeamCompetition.prototype.name = "";

        /**
         * TeamCompetition paid.
         * @member {boolean} paid
         * @memberof dss.TeamCompetition
         * @instance
         */
        TeamCompetition.prototype.paid = false;

        /**
         * Creates a new TeamCompetition instance using the specified properties.
         * @function create
         * @memberof dss.TeamCompetition
         * @static
         * @param {dss.ITeamCompetition=} [properties] Properties to set
         * @returns {dss.TeamCompetition} TeamCompetition instance
         */
        TeamCompetition.create = function create(properties) {
            return new TeamCompetition(properties);
        };

        /**
         * Encodes the specified TeamCompetition message. Does not implicitly {@link dss.TeamCompetition.verify|verify} messages.
         * @function encode
         * @memberof dss.TeamCompetition
         * @static
         * @param {dss.ITeamCompetition} message TeamCompetition message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TeamCompetition.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
            if (message.paid != null && Object.hasOwnProperty.call(message, "paid"))
                writer.uint32(/* id 2, wireType 0 =*/16).bool(message.paid);
            return writer;
        };

        /**
         * Encodes the specified TeamCompetition message, length delimited. Does not implicitly {@link dss.TeamCompetition.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.TeamCompetition
         * @static
         * @param {dss.ITeamCompetition} message TeamCompetition message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TeamCompetition.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a TeamCompetition message from the specified reader or buffer.
         * @function decode
         * @memberof dss.TeamCompetition
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.TeamCompetition} TeamCompetition
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TeamCompetition.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.TeamCompetition();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.paid = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a TeamCompetition message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.TeamCompetition
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.TeamCompetition} TeamCompetition
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TeamCompetition.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a TeamCompetition message.
         * @function verify
         * @memberof dss.TeamCompetition
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        TeamCompetition.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.paid != null && message.hasOwnProperty("paid"))
                if (typeof message.paid !== "boolean")
                    return "paid: boolean expected";
            return null;
        };

        /**
         * Creates a TeamCompetition message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.TeamCompetition
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.TeamCompetition} TeamCompetition
         */
        TeamCompetition.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.TeamCompetition)
                return object;
            var message = new $root.dss.TeamCompetition();
            if (object.name != null)
                message.name = String(object.name);
            if (object.paid != null)
                message.paid = Boolean(object.paid);
            return message;
        };

        /**
         * Creates a plain object from a TeamCompetition message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.TeamCompetition
         * @static
         * @param {dss.TeamCompetition} message TeamCompetition
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        TeamCompetition.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.name = "";
                object.paid = false;
            }
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.paid != null && message.hasOwnProperty("paid"))
                object.paid = message.paid;
            return object;
        };

        /**
         * Converts this TeamCompetition to JSON.
         * @function toJSON
         * @memberof dss.TeamCompetition
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        TeamCompetition.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return TeamCompetition;
    })();

    dss.TShirt = (function() {

        /**
         * Properties of a TShirt.
         * @memberof dss
         * @interface ITShirt
         * @property {dss.TShirt.Style|null} [style] TShirt style
         * @property {boolean|null} [paid] TShirt paid
         */

        /**
         * Constructs a new TShirt.
         * @memberof dss
         * @classdesc Represents a TShirt.
         * @implements ITShirt
         * @constructor
         * @param {dss.ITShirt=} [properties] Properties to set
         */
        function TShirt(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * TShirt style.
         * @member {dss.TShirt.Style} style
         * @memberof dss.TShirt
         * @instance
         */
        TShirt.prototype.style = 0;

        /**
         * TShirt paid.
         * @member {boolean} paid
         * @memberof dss.TShirt
         * @instance
         */
        TShirt.prototype.paid = false;

        /**
         * Creates a new TShirt instance using the specified properties.
         * @function create
         * @memberof dss.TShirt
         * @static
         * @param {dss.ITShirt=} [properties] Properties to set
         * @returns {dss.TShirt} TShirt instance
         */
        TShirt.create = function create(properties) {
            return new TShirt(properties);
        };

        /**
         * Encodes the specified TShirt message. Does not implicitly {@link dss.TShirt.verify|verify} messages.
         * @function encode
         * @memberof dss.TShirt
         * @static
         * @param {dss.ITShirt} message TShirt message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TShirt.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.style != null && Object.hasOwnProperty.call(message, "style"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.style);
            if (message.paid != null && Object.hasOwnProperty.call(message, "paid"))
                writer.uint32(/* id 2, wireType 0 =*/16).bool(message.paid);
            return writer;
        };

        /**
         * Encodes the specified TShirt message, length delimited. Does not implicitly {@link dss.TShirt.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.TShirt
         * @static
         * @param {dss.ITShirt} message TShirt message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        TShirt.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a TShirt message from the specified reader or buffer.
         * @function decode
         * @memberof dss.TShirt
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.TShirt} TShirt
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TShirt.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.TShirt();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.style = reader.int32();
                    break;
                case 2:
                    message.paid = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a TShirt message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.TShirt
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.TShirt} TShirt
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        TShirt.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a TShirt message.
         * @function verify
         * @memberof dss.TShirt
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        TShirt.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.style != null && message.hasOwnProperty("style"))
                switch (message.style) {
                default:
                    return "style: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                case 5:
                case 6:
                case 7:
                case 8:
                case 9:
                case 10:
                    break;
                }
            if (message.paid != null && message.hasOwnProperty("paid"))
                if (typeof message.paid !== "boolean")
                    return "paid: boolean expected";
            return null;
        };

        /**
         * Creates a TShirt message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.TShirt
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.TShirt} TShirt
         */
        TShirt.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.TShirt)
                return object;
            var message = new $root.dss.TShirt();
            switch (object.style) {
            case "UnisexS":
            case 0:
                message.style = 0;
                break;
            case "UnisexM":
            case 1:
                message.style = 1;
                break;
            case "UnisexL":
            case 2:
                message.style = 2;
                break;
            case "UnisexXL":
            case 3:
                message.style = 3;
                break;
            case "Unisex2XL":
            case 4:
                message.style = 4;
                break;
            case "Unisex3XL":
            case 5:
                message.style = 5;
                break;
            case "BellaS":
            case 6:
                message.style = 6;
                break;
            case "BellaM":
            case 7:
                message.style = 7;
                break;
            case "BellaL":
            case 8:
                message.style = 8;
                break;
            case "BellaXL":
            case 9:
                message.style = 9;
                break;
            case "Bella2XL":
            case 10:
                message.style = 10;
                break;
            }
            if (object.paid != null)
                message.paid = Boolean(object.paid);
            return message;
        };

        /**
         * Creates a plain object from a TShirt message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.TShirt
         * @static
         * @param {dss.TShirt} message TShirt
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        TShirt.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.style = options.enums === String ? "UnisexS" : 0;
                object.paid = false;
            }
            if (message.style != null && message.hasOwnProperty("style"))
                object.style = options.enums === String ? $root.dss.TShirt.Style[message.style] : message.style;
            if (message.paid != null && message.hasOwnProperty("paid"))
                object.paid = message.paid;
            return object;
        };

        /**
         * Converts this TShirt to JSON.
         * @function toJSON
         * @memberof dss.TShirt
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        TShirt.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Style enum.
         * @name dss.TShirt.Style
         * @enum {number}
         * @property {number} UnisexS=0 UnisexS value
         * @property {number} UnisexM=1 UnisexM value
         * @property {number} UnisexL=2 UnisexL value
         * @property {number} UnisexXL=3 UnisexXL value
         * @property {number} Unisex2XL=4 Unisex2XL value
         * @property {number} Unisex3XL=5 Unisex3XL value
         * @property {number} BellaS=6 BellaS value
         * @property {number} BellaM=7 BellaM value
         * @property {number} BellaL=8 BellaL value
         * @property {number} BellaXL=9 BellaXL value
         * @property {number} Bella2XL=10 Bella2XL value
         */
        TShirt.Style = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "UnisexS"] = 0;
            values[valuesById[1] = "UnisexM"] = 1;
            values[valuesById[2] = "UnisexL"] = 2;
            values[valuesById[3] = "UnisexXL"] = 3;
            values[valuesById[4] = "Unisex2XL"] = 4;
            values[valuesById[5] = "Unisex3XL"] = 5;
            values[valuesById[6] = "BellaS"] = 6;
            values[valuesById[7] = "BellaM"] = 7;
            values[valuesById[8] = "BellaL"] = 8;
            values[valuesById[9] = "BellaXL"] = 9;
            values[valuesById[10] = "Bella2XL"] = 10;
            return values;
        })();

        return TShirt;
    })();

    dss.ProvideHousing = (function() {

        /**
         * Properties of a ProvideHousing.
         * @memberof dss
         * @interface IProvideHousing
         * @property {string|null} [pets] ProvideHousing pets
         * @property {number|Long|null} [quantity] ProvideHousing quantity
         * @property {string|null} [details] ProvideHousing details
         */

        /**
         * Constructs a new ProvideHousing.
         * @memberof dss
         * @classdesc Represents a ProvideHousing.
         * @implements IProvideHousing
         * @constructor
         * @param {dss.IProvideHousing=} [properties] Properties to set
         */
        function ProvideHousing(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ProvideHousing pets.
         * @member {string} pets
         * @memberof dss.ProvideHousing
         * @instance
         */
        ProvideHousing.prototype.pets = "";

        /**
         * ProvideHousing quantity.
         * @member {number|Long} quantity
         * @memberof dss.ProvideHousing
         * @instance
         */
        ProvideHousing.prototype.quantity = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * ProvideHousing details.
         * @member {string} details
         * @memberof dss.ProvideHousing
         * @instance
         */
        ProvideHousing.prototype.details = "";

        /**
         * Creates a new ProvideHousing instance using the specified properties.
         * @function create
         * @memberof dss.ProvideHousing
         * @static
         * @param {dss.IProvideHousing=} [properties] Properties to set
         * @returns {dss.ProvideHousing} ProvideHousing instance
         */
        ProvideHousing.create = function create(properties) {
            return new ProvideHousing(properties);
        };

        /**
         * Encodes the specified ProvideHousing message. Does not implicitly {@link dss.ProvideHousing.verify|verify} messages.
         * @function encode
         * @memberof dss.ProvideHousing
         * @static
         * @param {dss.IProvideHousing} message ProvideHousing message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ProvideHousing.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.pets != null && Object.hasOwnProperty.call(message, "pets"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.pets);
            if (message.quantity != null && Object.hasOwnProperty.call(message, "quantity"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.quantity);
            if (message.details != null && Object.hasOwnProperty.call(message, "details"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.details);
            return writer;
        };

        /**
         * Encodes the specified ProvideHousing message, length delimited. Does not implicitly {@link dss.ProvideHousing.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.ProvideHousing
         * @static
         * @param {dss.IProvideHousing} message ProvideHousing message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ProvideHousing.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ProvideHousing message from the specified reader or buffer.
         * @function decode
         * @memberof dss.ProvideHousing
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.ProvideHousing} ProvideHousing
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ProvideHousing.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.ProvideHousing();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.pets = reader.string();
                    break;
                case 2:
                    message.quantity = reader.int64();
                    break;
                case 3:
                    message.details = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ProvideHousing message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.ProvideHousing
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.ProvideHousing} ProvideHousing
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ProvideHousing.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ProvideHousing message.
         * @function verify
         * @memberof dss.ProvideHousing
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ProvideHousing.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.pets != null && message.hasOwnProperty("pets"))
                if (!$util.isString(message.pets))
                    return "pets: string expected";
            if (message.quantity != null && message.hasOwnProperty("quantity"))
                if (!$util.isInteger(message.quantity) && !(message.quantity && $util.isInteger(message.quantity.low) && $util.isInteger(message.quantity.high)))
                    return "quantity: integer|Long expected";
            if (message.details != null && message.hasOwnProperty("details"))
                if (!$util.isString(message.details))
                    return "details: string expected";
            return null;
        };

        /**
         * Creates a ProvideHousing message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.ProvideHousing
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.ProvideHousing} ProvideHousing
         */
        ProvideHousing.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.ProvideHousing)
                return object;
            var message = new $root.dss.ProvideHousing();
            if (object.pets != null)
                message.pets = String(object.pets);
            if (object.quantity != null)
                if ($util.Long)
                    (message.quantity = $util.Long.fromValue(object.quantity)).unsigned = false;
                else if (typeof object.quantity === "string")
                    message.quantity = parseInt(object.quantity, 10);
                else if (typeof object.quantity === "number")
                    message.quantity = object.quantity;
                else if (typeof object.quantity === "object")
                    message.quantity = new $util.LongBits(object.quantity.low >>> 0, object.quantity.high >>> 0).toNumber();
            if (object.details != null)
                message.details = String(object.details);
            return message;
        };

        /**
         * Creates a plain object from a ProvideHousing message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.ProvideHousing
         * @static
         * @param {dss.ProvideHousing} message ProvideHousing
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ProvideHousing.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.pets = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.quantity = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.quantity = options.longs === String ? "0" : 0;
                object.details = "";
            }
            if (message.pets != null && message.hasOwnProperty("pets"))
                object.pets = message.pets;
            if (message.quantity != null && message.hasOwnProperty("quantity"))
                if (typeof message.quantity === "number")
                    object.quantity = options.longs === String ? String(message.quantity) : message.quantity;
                else
                    object.quantity = options.longs === String ? $util.Long.prototype.toString.call(message.quantity) : options.longs === Number ? new $util.LongBits(message.quantity.low >>> 0, message.quantity.high >>> 0).toNumber() : message.quantity;
            if (message.details != null && message.hasOwnProperty("details"))
                object.details = message.details;
            return object;
        };

        /**
         * Converts this ProvideHousing to JSON.
         * @function toJSON
         * @memberof dss.ProvideHousing
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ProvideHousing.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ProvideHousing;
    })();

    dss.RequireHousing = (function() {

        /**
         * Properties of a RequireHousing.
         * @memberof dss
         * @interface IRequireHousing
         * @property {string|null} [petAllergies] RequireHousing petAllergies
         * @property {string|null} [details] RequireHousing details
         */

        /**
         * Constructs a new RequireHousing.
         * @memberof dss
         * @classdesc Represents a RequireHousing.
         * @implements IRequireHousing
         * @constructor
         * @param {dss.IRequireHousing=} [properties] Properties to set
         */
        function RequireHousing(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RequireHousing petAllergies.
         * @member {string} petAllergies
         * @memberof dss.RequireHousing
         * @instance
         */
        RequireHousing.prototype.petAllergies = "";

        /**
         * RequireHousing details.
         * @member {string} details
         * @memberof dss.RequireHousing
         * @instance
         */
        RequireHousing.prototype.details = "";

        /**
         * Creates a new RequireHousing instance using the specified properties.
         * @function create
         * @memberof dss.RequireHousing
         * @static
         * @param {dss.IRequireHousing=} [properties] Properties to set
         * @returns {dss.RequireHousing} RequireHousing instance
         */
        RequireHousing.create = function create(properties) {
            return new RequireHousing(properties);
        };

        /**
         * Encodes the specified RequireHousing message. Does not implicitly {@link dss.RequireHousing.verify|verify} messages.
         * @function encode
         * @memberof dss.RequireHousing
         * @static
         * @param {dss.IRequireHousing} message RequireHousing message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequireHousing.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.petAllergies != null && Object.hasOwnProperty.call(message, "petAllergies"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.petAllergies);
            if (message.details != null && Object.hasOwnProperty.call(message, "details"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.details);
            return writer;
        };

        /**
         * Encodes the specified RequireHousing message, length delimited. Does not implicitly {@link dss.RequireHousing.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RequireHousing
         * @static
         * @param {dss.IRequireHousing} message RequireHousing message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequireHousing.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RequireHousing message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RequireHousing
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RequireHousing} RequireHousing
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequireHousing.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RequireHousing();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.petAllergies = reader.string();
                    break;
                case 2:
                    message.details = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RequireHousing message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RequireHousing
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RequireHousing} RequireHousing
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequireHousing.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RequireHousing message.
         * @function verify
         * @memberof dss.RequireHousing
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RequireHousing.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.petAllergies != null && message.hasOwnProperty("petAllergies"))
                if (!$util.isString(message.petAllergies))
                    return "petAllergies: string expected";
            if (message.details != null && message.hasOwnProperty("details"))
                if (!$util.isString(message.details))
                    return "details: string expected";
            return null;
        };

        /**
         * Creates a RequireHousing message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RequireHousing
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RequireHousing} RequireHousing
         */
        RequireHousing.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RequireHousing)
                return object;
            var message = new $root.dss.RequireHousing();
            if (object.petAllergies != null)
                message.petAllergies = String(object.petAllergies);
            if (object.details != null)
                message.details = String(object.details);
            return message;
        };

        /**
         * Creates a plain object from a RequireHousing message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RequireHousing
         * @static
         * @param {dss.RequireHousing} message RequireHousing
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RequireHousing.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.petAllergies = "";
                object.details = "";
            }
            if (message.petAllergies != null && message.hasOwnProperty("petAllergies"))
                object.petAllergies = message.petAllergies;
            if (message.details != null && message.hasOwnProperty("details"))
                object.details = message.details;
            return object;
        };

        /**
         * Converts this RequireHousing to JSON.
         * @function toJSON
         * @memberof dss.RequireHousing
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RequireHousing.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RequireHousing;
    })();

    dss.NoHousing = (function() {

        /**
         * Properties of a NoHousing.
         * @memberof dss
         * @interface INoHousing
         */

        /**
         * Constructs a new NoHousing.
         * @memberof dss
         * @classdesc Represents a NoHousing.
         * @implements INoHousing
         * @constructor
         * @param {dss.INoHousing=} [properties] Properties to set
         */
        function NoHousing(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new NoHousing instance using the specified properties.
         * @function create
         * @memberof dss.NoHousing
         * @static
         * @param {dss.INoHousing=} [properties] Properties to set
         * @returns {dss.NoHousing} NoHousing instance
         */
        NoHousing.create = function create(properties) {
            return new NoHousing(properties);
        };

        /**
         * Encodes the specified NoHousing message. Does not implicitly {@link dss.NoHousing.verify|verify} messages.
         * @function encode
         * @memberof dss.NoHousing
         * @static
         * @param {dss.INoHousing} message NoHousing message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        NoHousing.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified NoHousing message, length delimited. Does not implicitly {@link dss.NoHousing.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.NoHousing
         * @static
         * @param {dss.INoHousing} message NoHousing message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        NoHousing.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a NoHousing message from the specified reader or buffer.
         * @function decode
         * @memberof dss.NoHousing
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.NoHousing} NoHousing
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        NoHousing.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.NoHousing();
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
         * Decodes a NoHousing message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.NoHousing
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.NoHousing} NoHousing
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        NoHousing.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a NoHousing message.
         * @function verify
         * @memberof dss.NoHousing
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        NoHousing.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a NoHousing message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.NoHousing
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.NoHousing} NoHousing
         */
        NoHousing.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.NoHousing)
                return object;
            return new $root.dss.NoHousing();
        };

        /**
         * Creates a plain object from a NoHousing message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.NoHousing
         * @static
         * @param {dss.NoHousing} message NoHousing
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        NoHousing.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this NoHousing to JSON.
         * @function toJSON
         * @memberof dss.NoHousing
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        NoHousing.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return NoHousing;
    })();

    dss.RegistrationAddReq = (function() {

        /**
         * Properties of a RegistrationAddReq.
         * @memberof dss
         * @interface IRegistrationAddReq
         * @property {string|null} [idempotencyKey] RegistrationAddReq idempotencyKey
         * @property {dss.IRegistrationInfo|null} [registration] RegistrationAddReq registration
         * @property {string|null} [redirectUrl] RegistrationAddReq redirectUrl
         */

        /**
         * Constructs a new RegistrationAddReq.
         * @memberof dss
         * @classdesc Represents a RegistrationAddReq.
         * @implements IRegistrationAddReq
         * @constructor
         * @param {dss.IRegistrationAddReq=} [properties] Properties to set
         */
        function RegistrationAddReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationAddReq idempotencyKey.
         * @member {string} idempotencyKey
         * @memberof dss.RegistrationAddReq
         * @instance
         */
        RegistrationAddReq.prototype.idempotencyKey = "";

        /**
         * RegistrationAddReq registration.
         * @member {dss.IRegistrationInfo|null|undefined} registration
         * @memberof dss.RegistrationAddReq
         * @instance
         */
        RegistrationAddReq.prototype.registration = null;

        /**
         * RegistrationAddReq redirectUrl.
         * @member {string} redirectUrl
         * @memberof dss.RegistrationAddReq
         * @instance
         */
        RegistrationAddReq.prototype.redirectUrl = "";

        /**
         * Creates a new RegistrationAddReq instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {dss.IRegistrationAddReq=} [properties] Properties to set
         * @returns {dss.RegistrationAddReq} RegistrationAddReq instance
         */
        RegistrationAddReq.create = function create(properties) {
            return new RegistrationAddReq(properties);
        };

        /**
         * Encodes the specified RegistrationAddReq message. Does not implicitly {@link dss.RegistrationAddReq.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {dss.IRegistrationAddReq} message RegistrationAddReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationAddReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.idempotencyKey != null && Object.hasOwnProperty.call(message, "idempotencyKey"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.idempotencyKey);
            if (message.registration != null && Object.hasOwnProperty.call(message, "registration"))
                $root.dss.RegistrationInfo.encode(message.registration, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.redirectUrl != null && Object.hasOwnProperty.call(message, "redirectUrl"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.redirectUrl);
            return writer;
        };

        /**
         * Encodes the specified RegistrationAddReq message, length delimited. Does not implicitly {@link dss.RegistrationAddReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {dss.IRegistrationAddReq} message RegistrationAddReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationAddReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationAddReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationAddReq} RegistrationAddReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationAddReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationAddReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.idempotencyKey = reader.string();
                    break;
                case 2:
                    message.registration = $root.dss.RegistrationInfo.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.redirectUrl = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationAddReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationAddReq} RegistrationAddReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationAddReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationAddReq message.
         * @function verify
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationAddReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.idempotencyKey != null && message.hasOwnProperty("idempotencyKey"))
                if (!$util.isString(message.idempotencyKey))
                    return "idempotencyKey: string expected";
            if (message.registration != null && message.hasOwnProperty("registration")) {
                var error = $root.dss.RegistrationInfo.verify(message.registration);
                if (error)
                    return "registration." + error;
            }
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                if (!$util.isString(message.redirectUrl))
                    return "redirectUrl: string expected";
            return null;
        };

        /**
         * Creates a RegistrationAddReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationAddReq} RegistrationAddReq
         */
        RegistrationAddReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationAddReq)
                return object;
            var message = new $root.dss.RegistrationAddReq();
            if (object.idempotencyKey != null)
                message.idempotencyKey = String(object.idempotencyKey);
            if (object.registration != null) {
                if (typeof object.registration !== "object")
                    throw TypeError(".dss.RegistrationAddReq.registration: object expected");
                message.registration = $root.dss.RegistrationInfo.fromObject(object.registration);
            }
            if (object.redirectUrl != null)
                message.redirectUrl = String(object.redirectUrl);
            return message;
        };

        /**
         * Creates a plain object from a RegistrationAddReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationAddReq
         * @static
         * @param {dss.RegistrationAddReq} message RegistrationAddReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationAddReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.idempotencyKey = "";
                object.registration = null;
                object.redirectUrl = "";
            }
            if (message.idempotencyKey != null && message.hasOwnProperty("idempotencyKey"))
                object.idempotencyKey = message.idempotencyKey;
            if (message.registration != null && message.hasOwnProperty("registration"))
                object.registration = $root.dss.RegistrationInfo.toObject(message.registration, options);
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                object.redirectUrl = message.redirectUrl;
            return object;
        };

        /**
         * Converts this RegistrationAddReq to JSON.
         * @function toJSON
         * @memberof dss.RegistrationAddReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationAddReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationAddReq;
    })();

    dss.RegistrationAddRes = (function() {

        /**
         * Properties of a RegistrationAddRes.
         * @memberof dss
         * @interface IRegistrationAddRes
         * @property {string|null} [redirectUrl] RegistrationAddRes redirectUrl
         */

        /**
         * Constructs a new RegistrationAddRes.
         * @memberof dss
         * @classdesc Represents a RegistrationAddRes.
         * @implements IRegistrationAddRes
         * @constructor
         * @param {dss.IRegistrationAddRes=} [properties] Properties to set
         */
        function RegistrationAddRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationAddRes redirectUrl.
         * @member {string} redirectUrl
         * @memberof dss.RegistrationAddRes
         * @instance
         */
        RegistrationAddRes.prototype.redirectUrl = "";

        /**
         * Creates a new RegistrationAddRes instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {dss.IRegistrationAddRes=} [properties] Properties to set
         * @returns {dss.RegistrationAddRes} RegistrationAddRes instance
         */
        RegistrationAddRes.create = function create(properties) {
            return new RegistrationAddRes(properties);
        };

        /**
         * Encodes the specified RegistrationAddRes message. Does not implicitly {@link dss.RegistrationAddRes.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {dss.IRegistrationAddRes} message RegistrationAddRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationAddRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.redirectUrl != null && Object.hasOwnProperty.call(message, "redirectUrl"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.redirectUrl);
            return writer;
        };

        /**
         * Encodes the specified RegistrationAddRes message, length delimited. Does not implicitly {@link dss.RegistrationAddRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {dss.IRegistrationAddRes} message RegistrationAddRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationAddRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationAddRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationAddRes} RegistrationAddRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationAddRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationAddRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.redirectUrl = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationAddRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationAddRes} RegistrationAddRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationAddRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationAddRes message.
         * @function verify
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationAddRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                if (!$util.isString(message.redirectUrl))
                    return "redirectUrl: string expected";
            return null;
        };

        /**
         * Creates a RegistrationAddRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationAddRes} RegistrationAddRes
         */
        RegistrationAddRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationAddRes)
                return object;
            var message = new $root.dss.RegistrationAddRes();
            if (object.redirectUrl != null)
                message.redirectUrl = String(object.redirectUrl);
            return message;
        };

        /**
         * Creates a plain object from a RegistrationAddRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationAddRes
         * @static
         * @param {dss.RegistrationAddRes} message RegistrationAddRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationAddRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.redirectUrl = "";
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                object.redirectUrl = message.redirectUrl;
            return object;
        };

        /**
         * Converts this RegistrationAddRes to JSON.
         * @function toJSON
         * @memberof dss.RegistrationAddRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationAddRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationAddRes;
    })();

    dss.RegistrationGetReq = (function() {

        /**
         * Properties of a RegistrationGetReq.
         * @memberof dss
         * @interface IRegistrationGetReq
         * @property {string|null} [id] RegistrationGetReq id
         */

        /**
         * Constructs a new RegistrationGetReq.
         * @memberof dss
         * @classdesc Represents a RegistrationGetReq.
         * @implements IRegistrationGetReq
         * @constructor
         * @param {dss.IRegistrationGetReq=} [properties] Properties to set
         */
        function RegistrationGetReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationGetReq id.
         * @member {string} id
         * @memberof dss.RegistrationGetReq
         * @instance
         */
        RegistrationGetReq.prototype.id = "";

        /**
         * Creates a new RegistrationGetReq instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {dss.IRegistrationGetReq=} [properties] Properties to set
         * @returns {dss.RegistrationGetReq} RegistrationGetReq instance
         */
        RegistrationGetReq.create = function create(properties) {
            return new RegistrationGetReq(properties);
        };

        /**
         * Encodes the specified RegistrationGetReq message. Does not implicitly {@link dss.RegistrationGetReq.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {dss.IRegistrationGetReq} message RegistrationGetReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            return writer;
        };

        /**
         * Encodes the specified RegistrationGetReq message, length delimited. Does not implicitly {@link dss.RegistrationGetReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {dss.IRegistrationGetReq} message RegistrationGetReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationGetReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationGetReq} RegistrationGetReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationGetReq();
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
         * Decodes a RegistrationGetReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationGetReq} RegistrationGetReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationGetReq message.
         * @function verify
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationGetReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            return null;
        };

        /**
         * Creates a RegistrationGetReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationGetReq} RegistrationGetReq
         */
        RegistrationGetReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationGetReq)
                return object;
            var message = new $root.dss.RegistrationGetReq();
            if (object.id != null)
                message.id = String(object.id);
            return message;
        };

        /**
         * Creates a plain object from a RegistrationGetReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationGetReq
         * @static
         * @param {dss.RegistrationGetReq} message RegistrationGetReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationGetReq.toObject = function toObject(message, options) {
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
         * Converts this RegistrationGetReq to JSON.
         * @function toJSON
         * @memberof dss.RegistrationGetReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationGetReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationGetReq;
    })();

    dss.RegistrationGetRes = (function() {

        /**
         * Properties of a RegistrationGetRes.
         * @memberof dss
         * @interface IRegistrationGetRes
         * @property {dss.IRegistrationInfo|null} [registration] RegistrationGetRes registration
         */

        /**
         * Constructs a new RegistrationGetRes.
         * @memberof dss
         * @classdesc Represents a RegistrationGetRes.
         * @implements IRegistrationGetRes
         * @constructor
         * @param {dss.IRegistrationGetRes=} [properties] Properties to set
         */
        function RegistrationGetRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationGetRes registration.
         * @member {dss.IRegistrationInfo|null|undefined} registration
         * @memberof dss.RegistrationGetRes
         * @instance
         */
        RegistrationGetRes.prototype.registration = null;

        /**
         * Creates a new RegistrationGetRes instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {dss.IRegistrationGetRes=} [properties] Properties to set
         * @returns {dss.RegistrationGetRes} RegistrationGetRes instance
         */
        RegistrationGetRes.create = function create(properties) {
            return new RegistrationGetRes(properties);
        };

        /**
         * Encodes the specified RegistrationGetRes message. Does not implicitly {@link dss.RegistrationGetRes.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {dss.IRegistrationGetRes} message RegistrationGetRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.registration != null && Object.hasOwnProperty.call(message, "registration"))
                $root.dss.RegistrationInfo.encode(message.registration, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified RegistrationGetRes message, length delimited. Does not implicitly {@link dss.RegistrationGetRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {dss.IRegistrationGetRes} message RegistrationGetRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationGetRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationGetRes} RegistrationGetRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationGetRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.registration = $root.dss.RegistrationInfo.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationGetRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationGetRes} RegistrationGetRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationGetRes message.
         * @function verify
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationGetRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.registration != null && message.hasOwnProperty("registration")) {
                var error = $root.dss.RegistrationInfo.verify(message.registration);
                if (error)
                    return "registration." + error;
            }
            return null;
        };

        /**
         * Creates a RegistrationGetRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationGetRes} RegistrationGetRes
         */
        RegistrationGetRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationGetRes)
                return object;
            var message = new $root.dss.RegistrationGetRes();
            if (object.registration != null) {
                if (typeof object.registration !== "object")
                    throw TypeError(".dss.RegistrationGetRes.registration: object expected");
                message.registration = $root.dss.RegistrationInfo.fromObject(object.registration);
            }
            return message;
        };

        /**
         * Creates a plain object from a RegistrationGetRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationGetRes
         * @static
         * @param {dss.RegistrationGetRes} message RegistrationGetRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationGetRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.registration = null;
            if (message.registration != null && message.hasOwnProperty("registration"))
                object.registration = $root.dss.RegistrationInfo.toObject(message.registration, options);
            return object;
        };

        /**
         * Converts this RegistrationGetRes to JSON.
         * @function toJSON
         * @memberof dss.RegistrationGetRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationGetRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationGetRes;
    })();

    dss.RegistrationPricesReq = (function() {

        /**
         * Properties of a RegistrationPricesReq.
         * @memberof dss
         * @interface IRegistrationPricesReq
         */

        /**
         * Constructs a new RegistrationPricesReq.
         * @memberof dss
         * @classdesc Represents a RegistrationPricesReq.
         * @implements IRegistrationPricesReq
         * @constructor
         * @param {dss.IRegistrationPricesReq=} [properties] Properties to set
         */
        function RegistrationPricesReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new RegistrationPricesReq instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {dss.IRegistrationPricesReq=} [properties] Properties to set
         * @returns {dss.RegistrationPricesReq} RegistrationPricesReq instance
         */
        RegistrationPricesReq.create = function create(properties) {
            return new RegistrationPricesReq(properties);
        };

        /**
         * Encodes the specified RegistrationPricesReq message. Does not implicitly {@link dss.RegistrationPricesReq.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {dss.IRegistrationPricesReq} message RegistrationPricesReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationPricesReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified RegistrationPricesReq message, length delimited. Does not implicitly {@link dss.RegistrationPricesReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {dss.IRegistrationPricesReq} message RegistrationPricesReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationPricesReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationPricesReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationPricesReq} RegistrationPricesReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationPricesReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationPricesReq();
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
         * Decodes a RegistrationPricesReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationPricesReq} RegistrationPricesReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationPricesReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationPricesReq message.
         * @function verify
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationPricesReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a RegistrationPricesReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationPricesReq} RegistrationPricesReq
         */
        RegistrationPricesReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationPricesReq)
                return object;
            return new $root.dss.RegistrationPricesReq();
        };

        /**
         * Creates a plain object from a RegistrationPricesReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationPricesReq
         * @static
         * @param {dss.RegistrationPricesReq} message RegistrationPricesReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationPricesReq.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this RegistrationPricesReq to JSON.
         * @function toJSON
         * @memberof dss.RegistrationPricesReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationPricesReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationPricesReq;
    })();

    dss.RegistrationPricesRes = (function() {

        /**
         * Properties of a RegistrationPricesRes.
         * @memberof dss
         * @interface IRegistrationPricesRes
         * @property {number|Long|null} [weekendPassCost] RegistrationPricesRes weekendPassCost
         * @property {dss.FullWeekendPassTier|null} [weekendPassTier] RegistrationPricesRes weekendPassTier
         * @property {number|Long|null} [dancePassCost] RegistrationPricesRes dancePassCost
         * @property {number|Long|null} [mixAndMatchCost] RegistrationPricesRes mixAndMatchCost
         * @property {number|Long|null} [soloJazzCost] RegistrationPricesRes soloJazzCost
         * @property {number|Long|null} [teamCompetitionCost] RegistrationPricesRes teamCompetitionCost
         * @property {number|Long|null} [tshirtCost] RegistrationPricesRes tshirtCost
         * @property {dss.IDiscountAmount|null} [studentDiscount] RegistrationPricesRes studentDiscount
         */

        /**
         * Constructs a new RegistrationPricesRes.
         * @memberof dss
         * @classdesc Represents a RegistrationPricesRes.
         * @implements IRegistrationPricesRes
         * @constructor
         * @param {dss.IRegistrationPricesRes=} [properties] Properties to set
         */
        function RegistrationPricesRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationPricesRes weekendPassCost.
         * @member {number|Long} weekendPassCost
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.weekendPassCost = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RegistrationPricesRes weekendPassTier.
         * @member {dss.FullWeekendPassTier} weekendPassTier
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.weekendPassTier = 0;

        /**
         * RegistrationPricesRes dancePassCost.
         * @member {number|Long} dancePassCost
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.dancePassCost = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RegistrationPricesRes mixAndMatchCost.
         * @member {number|Long} mixAndMatchCost
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.mixAndMatchCost = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RegistrationPricesRes soloJazzCost.
         * @member {number|Long} soloJazzCost
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.soloJazzCost = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RegistrationPricesRes teamCompetitionCost.
         * @member {number|Long} teamCompetitionCost
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.teamCompetitionCost = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RegistrationPricesRes tshirtCost.
         * @member {number|Long} tshirtCost
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.tshirtCost = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * RegistrationPricesRes studentDiscount.
         * @member {dss.IDiscountAmount|null|undefined} studentDiscount
         * @memberof dss.RegistrationPricesRes
         * @instance
         */
        RegistrationPricesRes.prototype.studentDiscount = null;

        /**
         * Creates a new RegistrationPricesRes instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {dss.IRegistrationPricesRes=} [properties] Properties to set
         * @returns {dss.RegistrationPricesRes} RegistrationPricesRes instance
         */
        RegistrationPricesRes.create = function create(properties) {
            return new RegistrationPricesRes(properties);
        };

        /**
         * Encodes the specified RegistrationPricesRes message. Does not implicitly {@link dss.RegistrationPricesRes.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {dss.IRegistrationPricesRes} message RegistrationPricesRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationPricesRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.weekendPassCost != null && Object.hasOwnProperty.call(message, "weekendPassCost"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.weekendPassCost);
            if (message.weekendPassTier != null && Object.hasOwnProperty.call(message, "weekendPassTier"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.weekendPassTier);
            if (message.dancePassCost != null && Object.hasOwnProperty.call(message, "dancePassCost"))
                writer.uint32(/* id 3, wireType 0 =*/24).int64(message.dancePassCost);
            if (message.mixAndMatchCost != null && Object.hasOwnProperty.call(message, "mixAndMatchCost"))
                writer.uint32(/* id 4, wireType 0 =*/32).int64(message.mixAndMatchCost);
            if (message.soloJazzCost != null && Object.hasOwnProperty.call(message, "soloJazzCost"))
                writer.uint32(/* id 5, wireType 0 =*/40).int64(message.soloJazzCost);
            if (message.teamCompetitionCost != null && Object.hasOwnProperty.call(message, "teamCompetitionCost"))
                writer.uint32(/* id 6, wireType 0 =*/48).int64(message.teamCompetitionCost);
            if (message.tshirtCost != null && Object.hasOwnProperty.call(message, "tshirtCost"))
                writer.uint32(/* id 7, wireType 0 =*/56).int64(message.tshirtCost);
            if (message.studentDiscount != null && Object.hasOwnProperty.call(message, "studentDiscount"))
                $root.dss.DiscountAmount.encode(message.studentDiscount, writer.uint32(/* id 8, wireType 2 =*/66).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified RegistrationPricesRes message, length delimited. Does not implicitly {@link dss.RegistrationPricesRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {dss.IRegistrationPricesRes} message RegistrationPricesRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationPricesRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationPricesRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationPricesRes} RegistrationPricesRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationPricesRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationPricesRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.weekendPassCost = reader.int64();
                    break;
                case 2:
                    message.weekendPassTier = reader.int32();
                    break;
                case 3:
                    message.dancePassCost = reader.int64();
                    break;
                case 4:
                    message.mixAndMatchCost = reader.int64();
                    break;
                case 5:
                    message.soloJazzCost = reader.int64();
                    break;
                case 6:
                    message.teamCompetitionCost = reader.int64();
                    break;
                case 7:
                    message.tshirtCost = reader.int64();
                    break;
                case 8:
                    message.studentDiscount = $root.dss.DiscountAmount.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationPricesRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationPricesRes} RegistrationPricesRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationPricesRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationPricesRes message.
         * @function verify
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationPricesRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.weekendPassCost != null && message.hasOwnProperty("weekendPassCost"))
                if (!$util.isInteger(message.weekendPassCost) && !(message.weekendPassCost && $util.isInteger(message.weekendPassCost.low) && $util.isInteger(message.weekendPassCost.high)))
                    return "weekendPassCost: integer|Long expected";
            if (message.weekendPassTier != null && message.hasOwnProperty("weekendPassTier"))
                switch (message.weekendPassTier) {
                default:
                    return "weekendPassTier: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                    break;
                }
            if (message.dancePassCost != null && message.hasOwnProperty("dancePassCost"))
                if (!$util.isInteger(message.dancePassCost) && !(message.dancePassCost && $util.isInteger(message.dancePassCost.low) && $util.isInteger(message.dancePassCost.high)))
                    return "dancePassCost: integer|Long expected";
            if (message.mixAndMatchCost != null && message.hasOwnProperty("mixAndMatchCost"))
                if (!$util.isInteger(message.mixAndMatchCost) && !(message.mixAndMatchCost && $util.isInteger(message.mixAndMatchCost.low) && $util.isInteger(message.mixAndMatchCost.high)))
                    return "mixAndMatchCost: integer|Long expected";
            if (message.soloJazzCost != null && message.hasOwnProperty("soloJazzCost"))
                if (!$util.isInteger(message.soloJazzCost) && !(message.soloJazzCost && $util.isInteger(message.soloJazzCost.low) && $util.isInteger(message.soloJazzCost.high)))
                    return "soloJazzCost: integer|Long expected";
            if (message.teamCompetitionCost != null && message.hasOwnProperty("teamCompetitionCost"))
                if (!$util.isInteger(message.teamCompetitionCost) && !(message.teamCompetitionCost && $util.isInteger(message.teamCompetitionCost.low) && $util.isInteger(message.teamCompetitionCost.high)))
                    return "teamCompetitionCost: integer|Long expected";
            if (message.tshirtCost != null && message.hasOwnProperty("tshirtCost"))
                if (!$util.isInteger(message.tshirtCost) && !(message.tshirtCost && $util.isInteger(message.tshirtCost.low) && $util.isInteger(message.tshirtCost.high)))
                    return "tshirtCost: integer|Long expected";
            if (message.studentDiscount != null && message.hasOwnProperty("studentDiscount")) {
                var error = $root.dss.DiscountAmount.verify(message.studentDiscount);
                if (error)
                    return "studentDiscount." + error;
            }
            return null;
        };

        /**
         * Creates a RegistrationPricesRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationPricesRes} RegistrationPricesRes
         */
        RegistrationPricesRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationPricesRes)
                return object;
            var message = new $root.dss.RegistrationPricesRes();
            if (object.weekendPassCost != null)
                if ($util.Long)
                    (message.weekendPassCost = $util.Long.fromValue(object.weekendPassCost)).unsigned = false;
                else if (typeof object.weekendPassCost === "string")
                    message.weekendPassCost = parseInt(object.weekendPassCost, 10);
                else if (typeof object.weekendPassCost === "number")
                    message.weekendPassCost = object.weekendPassCost;
                else if (typeof object.weekendPassCost === "object")
                    message.weekendPassCost = new $util.LongBits(object.weekendPassCost.low >>> 0, object.weekendPassCost.high >>> 0).toNumber();
            switch (object.weekendPassTier) {
            case "Tier1":
            case 0:
                message.weekendPassTier = 0;
                break;
            case "Tier2":
            case 1:
                message.weekendPassTier = 1;
                break;
            case "Tier3":
            case 2:
                message.weekendPassTier = 2;
                break;
            case "Tier4":
            case 3:
                message.weekendPassTier = 3;
                break;
            case "Tier5":
            case 4:
                message.weekendPassTier = 4;
                break;
            }
            if (object.dancePassCost != null)
                if ($util.Long)
                    (message.dancePassCost = $util.Long.fromValue(object.dancePassCost)).unsigned = false;
                else if (typeof object.dancePassCost === "string")
                    message.dancePassCost = parseInt(object.dancePassCost, 10);
                else if (typeof object.dancePassCost === "number")
                    message.dancePassCost = object.dancePassCost;
                else if (typeof object.dancePassCost === "object")
                    message.dancePassCost = new $util.LongBits(object.dancePassCost.low >>> 0, object.dancePassCost.high >>> 0).toNumber();
            if (object.mixAndMatchCost != null)
                if ($util.Long)
                    (message.mixAndMatchCost = $util.Long.fromValue(object.mixAndMatchCost)).unsigned = false;
                else if (typeof object.mixAndMatchCost === "string")
                    message.mixAndMatchCost = parseInt(object.mixAndMatchCost, 10);
                else if (typeof object.mixAndMatchCost === "number")
                    message.mixAndMatchCost = object.mixAndMatchCost;
                else if (typeof object.mixAndMatchCost === "object")
                    message.mixAndMatchCost = new $util.LongBits(object.mixAndMatchCost.low >>> 0, object.mixAndMatchCost.high >>> 0).toNumber();
            if (object.soloJazzCost != null)
                if ($util.Long)
                    (message.soloJazzCost = $util.Long.fromValue(object.soloJazzCost)).unsigned = false;
                else if (typeof object.soloJazzCost === "string")
                    message.soloJazzCost = parseInt(object.soloJazzCost, 10);
                else if (typeof object.soloJazzCost === "number")
                    message.soloJazzCost = object.soloJazzCost;
                else if (typeof object.soloJazzCost === "object")
                    message.soloJazzCost = new $util.LongBits(object.soloJazzCost.low >>> 0, object.soloJazzCost.high >>> 0).toNumber();
            if (object.teamCompetitionCost != null)
                if ($util.Long)
                    (message.teamCompetitionCost = $util.Long.fromValue(object.teamCompetitionCost)).unsigned = false;
                else if (typeof object.teamCompetitionCost === "string")
                    message.teamCompetitionCost = parseInt(object.teamCompetitionCost, 10);
                else if (typeof object.teamCompetitionCost === "number")
                    message.teamCompetitionCost = object.teamCompetitionCost;
                else if (typeof object.teamCompetitionCost === "object")
                    message.teamCompetitionCost = new $util.LongBits(object.teamCompetitionCost.low >>> 0, object.teamCompetitionCost.high >>> 0).toNumber();
            if (object.tshirtCost != null)
                if ($util.Long)
                    (message.tshirtCost = $util.Long.fromValue(object.tshirtCost)).unsigned = false;
                else if (typeof object.tshirtCost === "string")
                    message.tshirtCost = parseInt(object.tshirtCost, 10);
                else if (typeof object.tshirtCost === "number")
                    message.tshirtCost = object.tshirtCost;
                else if (typeof object.tshirtCost === "object")
                    message.tshirtCost = new $util.LongBits(object.tshirtCost.low >>> 0, object.tshirtCost.high >>> 0).toNumber();
            if (object.studentDiscount != null) {
                if (typeof object.studentDiscount !== "object")
                    throw TypeError(".dss.RegistrationPricesRes.studentDiscount: object expected");
                message.studentDiscount = $root.dss.DiscountAmount.fromObject(object.studentDiscount);
            }
            return message;
        };

        /**
         * Creates a plain object from a RegistrationPricesRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationPricesRes
         * @static
         * @param {dss.RegistrationPricesRes} message RegistrationPricesRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationPricesRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.weekendPassCost = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.weekendPassCost = options.longs === String ? "0" : 0;
                object.weekendPassTier = options.enums === String ? "Tier1" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.dancePassCost = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.dancePassCost = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.mixAndMatchCost = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.mixAndMatchCost = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.soloJazzCost = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.soloJazzCost = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.teamCompetitionCost = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.teamCompetitionCost = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.tshirtCost = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.tshirtCost = options.longs === String ? "0" : 0;
                object.studentDiscount = null;
            }
            if (message.weekendPassCost != null && message.hasOwnProperty("weekendPassCost"))
                if (typeof message.weekendPassCost === "number")
                    object.weekendPassCost = options.longs === String ? String(message.weekendPassCost) : message.weekendPassCost;
                else
                    object.weekendPassCost = options.longs === String ? $util.Long.prototype.toString.call(message.weekendPassCost) : options.longs === Number ? new $util.LongBits(message.weekendPassCost.low >>> 0, message.weekendPassCost.high >>> 0).toNumber() : message.weekendPassCost;
            if (message.weekendPassTier != null && message.hasOwnProperty("weekendPassTier"))
                object.weekendPassTier = options.enums === String ? $root.dss.FullWeekendPassTier[message.weekendPassTier] : message.weekendPassTier;
            if (message.dancePassCost != null && message.hasOwnProperty("dancePassCost"))
                if (typeof message.dancePassCost === "number")
                    object.dancePassCost = options.longs === String ? String(message.dancePassCost) : message.dancePassCost;
                else
                    object.dancePassCost = options.longs === String ? $util.Long.prototype.toString.call(message.dancePassCost) : options.longs === Number ? new $util.LongBits(message.dancePassCost.low >>> 0, message.dancePassCost.high >>> 0).toNumber() : message.dancePassCost;
            if (message.mixAndMatchCost != null && message.hasOwnProperty("mixAndMatchCost"))
                if (typeof message.mixAndMatchCost === "number")
                    object.mixAndMatchCost = options.longs === String ? String(message.mixAndMatchCost) : message.mixAndMatchCost;
                else
                    object.mixAndMatchCost = options.longs === String ? $util.Long.prototype.toString.call(message.mixAndMatchCost) : options.longs === Number ? new $util.LongBits(message.mixAndMatchCost.low >>> 0, message.mixAndMatchCost.high >>> 0).toNumber() : message.mixAndMatchCost;
            if (message.soloJazzCost != null && message.hasOwnProperty("soloJazzCost"))
                if (typeof message.soloJazzCost === "number")
                    object.soloJazzCost = options.longs === String ? String(message.soloJazzCost) : message.soloJazzCost;
                else
                    object.soloJazzCost = options.longs === String ? $util.Long.prototype.toString.call(message.soloJazzCost) : options.longs === Number ? new $util.LongBits(message.soloJazzCost.low >>> 0, message.soloJazzCost.high >>> 0).toNumber() : message.soloJazzCost;
            if (message.teamCompetitionCost != null && message.hasOwnProperty("teamCompetitionCost"))
                if (typeof message.teamCompetitionCost === "number")
                    object.teamCompetitionCost = options.longs === String ? String(message.teamCompetitionCost) : message.teamCompetitionCost;
                else
                    object.teamCompetitionCost = options.longs === String ? $util.Long.prototype.toString.call(message.teamCompetitionCost) : options.longs === Number ? new $util.LongBits(message.teamCompetitionCost.low >>> 0, message.teamCompetitionCost.high >>> 0).toNumber() : message.teamCompetitionCost;
            if (message.tshirtCost != null && message.hasOwnProperty("tshirtCost"))
                if (typeof message.tshirtCost === "number")
                    object.tshirtCost = options.longs === String ? String(message.tshirtCost) : message.tshirtCost;
                else
                    object.tshirtCost = options.longs === String ? $util.Long.prototype.toString.call(message.tshirtCost) : options.longs === Number ? new $util.LongBits(message.tshirtCost.low >>> 0, message.tshirtCost.high >>> 0).toNumber() : message.tshirtCost;
            if (message.studentDiscount != null && message.hasOwnProperty("studentDiscount"))
                object.studentDiscount = $root.dss.DiscountAmount.toObject(message.studentDiscount, options);
            return object;
        };

        /**
         * Converts this RegistrationPricesRes to JSON.
         * @function toJSON
         * @memberof dss.RegistrationPricesRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationPricesRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationPricesRes;
    })();

    dss.RegistrationUpdateReq = (function() {

        /**
         * Properties of a RegistrationUpdateReq.
         * @memberof dss
         * @interface IRegistrationUpdateReq
         * @property {string|null} [idempotencyKey] RegistrationUpdateReq idempotencyKey
         * @property {dss.IRegistrationInfo|null} [registration] RegistrationUpdateReq registration
         * @property {string|null} [redirectUrl] RegistrationUpdateReq redirectUrl
         */

        /**
         * Constructs a new RegistrationUpdateReq.
         * @memberof dss
         * @classdesc Represents a RegistrationUpdateReq.
         * @implements IRegistrationUpdateReq
         * @constructor
         * @param {dss.IRegistrationUpdateReq=} [properties] Properties to set
         */
        function RegistrationUpdateReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationUpdateReq idempotencyKey.
         * @member {string} idempotencyKey
         * @memberof dss.RegistrationUpdateReq
         * @instance
         */
        RegistrationUpdateReq.prototype.idempotencyKey = "";

        /**
         * RegistrationUpdateReq registration.
         * @member {dss.IRegistrationInfo|null|undefined} registration
         * @memberof dss.RegistrationUpdateReq
         * @instance
         */
        RegistrationUpdateReq.prototype.registration = null;

        /**
         * RegistrationUpdateReq redirectUrl.
         * @member {string} redirectUrl
         * @memberof dss.RegistrationUpdateReq
         * @instance
         */
        RegistrationUpdateReq.prototype.redirectUrl = "";

        /**
         * Creates a new RegistrationUpdateReq instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {dss.IRegistrationUpdateReq=} [properties] Properties to set
         * @returns {dss.RegistrationUpdateReq} RegistrationUpdateReq instance
         */
        RegistrationUpdateReq.create = function create(properties) {
            return new RegistrationUpdateReq(properties);
        };

        /**
         * Encodes the specified RegistrationUpdateReq message. Does not implicitly {@link dss.RegistrationUpdateReq.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {dss.IRegistrationUpdateReq} message RegistrationUpdateReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationUpdateReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.idempotencyKey != null && Object.hasOwnProperty.call(message, "idempotencyKey"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.idempotencyKey);
            if (message.registration != null && Object.hasOwnProperty.call(message, "registration"))
                $root.dss.RegistrationInfo.encode(message.registration, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.redirectUrl != null && Object.hasOwnProperty.call(message, "redirectUrl"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.redirectUrl);
            return writer;
        };

        /**
         * Encodes the specified RegistrationUpdateReq message, length delimited. Does not implicitly {@link dss.RegistrationUpdateReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {dss.IRegistrationUpdateReq} message RegistrationUpdateReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationUpdateReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationUpdateReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationUpdateReq} RegistrationUpdateReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationUpdateReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationUpdateReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.idempotencyKey = reader.string();
                    break;
                case 2:
                    message.registration = $root.dss.RegistrationInfo.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.redirectUrl = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationUpdateReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationUpdateReq} RegistrationUpdateReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationUpdateReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationUpdateReq message.
         * @function verify
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationUpdateReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.idempotencyKey != null && message.hasOwnProperty("idempotencyKey"))
                if (!$util.isString(message.idempotencyKey))
                    return "idempotencyKey: string expected";
            if (message.registration != null && message.hasOwnProperty("registration")) {
                var error = $root.dss.RegistrationInfo.verify(message.registration);
                if (error)
                    return "registration." + error;
            }
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                if (!$util.isString(message.redirectUrl))
                    return "redirectUrl: string expected";
            return null;
        };

        /**
         * Creates a RegistrationUpdateReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationUpdateReq} RegistrationUpdateReq
         */
        RegistrationUpdateReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationUpdateReq)
                return object;
            var message = new $root.dss.RegistrationUpdateReq();
            if (object.idempotencyKey != null)
                message.idempotencyKey = String(object.idempotencyKey);
            if (object.registration != null) {
                if (typeof object.registration !== "object")
                    throw TypeError(".dss.RegistrationUpdateReq.registration: object expected");
                message.registration = $root.dss.RegistrationInfo.fromObject(object.registration);
            }
            if (object.redirectUrl != null)
                message.redirectUrl = String(object.redirectUrl);
            return message;
        };

        /**
         * Creates a plain object from a RegistrationUpdateReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationUpdateReq
         * @static
         * @param {dss.RegistrationUpdateReq} message RegistrationUpdateReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationUpdateReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.idempotencyKey = "";
                object.registration = null;
                object.redirectUrl = "";
            }
            if (message.idempotencyKey != null && message.hasOwnProperty("idempotencyKey"))
                object.idempotencyKey = message.idempotencyKey;
            if (message.registration != null && message.hasOwnProperty("registration"))
                object.registration = $root.dss.RegistrationInfo.toObject(message.registration, options);
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                object.redirectUrl = message.redirectUrl;
            return object;
        };

        /**
         * Converts this RegistrationUpdateReq to JSON.
         * @function toJSON
         * @memberof dss.RegistrationUpdateReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationUpdateReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationUpdateReq;
    })();

    dss.RegistrationUpdateRes = (function() {

        /**
         * Properties of a RegistrationUpdateRes.
         * @memberof dss
         * @interface IRegistrationUpdateRes
         * @property {string|null} [redirectUrl] RegistrationUpdateRes redirectUrl
         */

        /**
         * Constructs a new RegistrationUpdateRes.
         * @memberof dss
         * @classdesc Represents a RegistrationUpdateRes.
         * @implements IRegistrationUpdateRes
         * @constructor
         * @param {dss.IRegistrationUpdateRes=} [properties] Properties to set
         */
        function RegistrationUpdateRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationUpdateRes redirectUrl.
         * @member {string} redirectUrl
         * @memberof dss.RegistrationUpdateRes
         * @instance
         */
        RegistrationUpdateRes.prototype.redirectUrl = "";

        /**
         * Creates a new RegistrationUpdateRes instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {dss.IRegistrationUpdateRes=} [properties] Properties to set
         * @returns {dss.RegistrationUpdateRes} RegistrationUpdateRes instance
         */
        RegistrationUpdateRes.create = function create(properties) {
            return new RegistrationUpdateRes(properties);
        };

        /**
         * Encodes the specified RegistrationUpdateRes message. Does not implicitly {@link dss.RegistrationUpdateRes.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {dss.IRegistrationUpdateRes} message RegistrationUpdateRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationUpdateRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.redirectUrl != null && Object.hasOwnProperty.call(message, "redirectUrl"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.redirectUrl);
            return writer;
        };

        /**
         * Encodes the specified RegistrationUpdateRes message, length delimited. Does not implicitly {@link dss.RegistrationUpdateRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {dss.IRegistrationUpdateRes} message RegistrationUpdateRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationUpdateRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationUpdateRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationUpdateRes} RegistrationUpdateRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationUpdateRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationUpdateRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.redirectUrl = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationUpdateRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationUpdateRes} RegistrationUpdateRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationUpdateRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationUpdateRes message.
         * @function verify
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationUpdateRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                if (!$util.isString(message.redirectUrl))
                    return "redirectUrl: string expected";
            return null;
        };

        /**
         * Creates a RegistrationUpdateRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationUpdateRes} RegistrationUpdateRes
         */
        RegistrationUpdateRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationUpdateRes)
                return object;
            var message = new $root.dss.RegistrationUpdateRes();
            if (object.redirectUrl != null)
                message.redirectUrl = String(object.redirectUrl);
            return message;
        };

        /**
         * Creates a plain object from a RegistrationUpdateRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationUpdateRes
         * @static
         * @param {dss.RegistrationUpdateRes} message RegistrationUpdateRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationUpdateRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.redirectUrl = "";
            if (message.redirectUrl != null && message.hasOwnProperty("redirectUrl"))
                object.redirectUrl = message.redirectUrl;
            return object;
        };

        /**
         * Converts this RegistrationUpdateRes to JSON.
         * @function toJSON
         * @memberof dss.RegistrationUpdateRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationUpdateRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationUpdateRes;
    })();

    dss.RegistrationGetSummaryReq = (function() {

        /**
         * Properties of a RegistrationGetSummaryReq.
         * @memberof dss
         * @interface IRegistrationGetSummaryReq
         */

        /**
         * Constructs a new RegistrationGetSummaryReq.
         * @memberof dss
         * @classdesc Represents a RegistrationGetSummaryReq.
         * @implements IRegistrationGetSummaryReq
         * @constructor
         * @param {dss.IRegistrationGetSummaryReq=} [properties] Properties to set
         */
        function RegistrationGetSummaryReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new RegistrationGetSummaryReq instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {dss.IRegistrationGetSummaryReq=} [properties] Properties to set
         * @returns {dss.RegistrationGetSummaryReq} RegistrationGetSummaryReq instance
         */
        RegistrationGetSummaryReq.create = function create(properties) {
            return new RegistrationGetSummaryReq(properties);
        };

        /**
         * Encodes the specified RegistrationGetSummaryReq message. Does not implicitly {@link dss.RegistrationGetSummaryReq.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {dss.IRegistrationGetSummaryReq} message RegistrationGetSummaryReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetSummaryReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified RegistrationGetSummaryReq message, length delimited. Does not implicitly {@link dss.RegistrationGetSummaryReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {dss.IRegistrationGetSummaryReq} message RegistrationGetSummaryReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetSummaryReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationGetSummaryReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationGetSummaryReq} RegistrationGetSummaryReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetSummaryReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationGetSummaryReq();
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
         * Decodes a RegistrationGetSummaryReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationGetSummaryReq} RegistrationGetSummaryReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetSummaryReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationGetSummaryReq message.
         * @function verify
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationGetSummaryReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a RegistrationGetSummaryReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationGetSummaryReq} RegistrationGetSummaryReq
         */
        RegistrationGetSummaryReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationGetSummaryReq)
                return object;
            return new $root.dss.RegistrationGetSummaryReq();
        };

        /**
         * Creates a plain object from a RegistrationGetSummaryReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationGetSummaryReq
         * @static
         * @param {dss.RegistrationGetSummaryReq} message RegistrationGetSummaryReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationGetSummaryReq.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this RegistrationGetSummaryReq to JSON.
         * @function toJSON
         * @memberof dss.RegistrationGetSummaryReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationGetSummaryReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationGetSummaryReq;
    })();

    dss.RegistrationSummary = (function() {

        /**
         * Properties of a RegistrationSummary.
         * @memberof dss
         * @interface IRegistrationSummary
         * @property {string|null} [id] RegistrationSummary id
         * @property {string|null} [firstName] RegistrationSummary firstName
         * @property {string|null} [lastName] RegistrationSummary lastName
         * @property {string|null} [email] RegistrationSummary email
         * @property {string|null} [createdAt] RegistrationSummary createdAt
         * @property {boolean|null} [paid] RegistrationSummary paid
         */

        /**
         * Constructs a new RegistrationSummary.
         * @memberof dss
         * @classdesc Represents a RegistrationSummary.
         * @implements IRegistrationSummary
         * @constructor
         * @param {dss.IRegistrationSummary=} [properties] Properties to set
         */
        function RegistrationSummary(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationSummary id.
         * @member {string} id
         * @memberof dss.RegistrationSummary
         * @instance
         */
        RegistrationSummary.prototype.id = "";

        /**
         * RegistrationSummary firstName.
         * @member {string} firstName
         * @memberof dss.RegistrationSummary
         * @instance
         */
        RegistrationSummary.prototype.firstName = "";

        /**
         * RegistrationSummary lastName.
         * @member {string} lastName
         * @memberof dss.RegistrationSummary
         * @instance
         */
        RegistrationSummary.prototype.lastName = "";

        /**
         * RegistrationSummary email.
         * @member {string} email
         * @memberof dss.RegistrationSummary
         * @instance
         */
        RegistrationSummary.prototype.email = "";

        /**
         * RegistrationSummary createdAt.
         * @member {string} createdAt
         * @memberof dss.RegistrationSummary
         * @instance
         */
        RegistrationSummary.prototype.createdAt = "";

        /**
         * RegistrationSummary paid.
         * @member {boolean} paid
         * @memberof dss.RegistrationSummary
         * @instance
         */
        RegistrationSummary.prototype.paid = false;

        /**
         * Creates a new RegistrationSummary instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationSummary
         * @static
         * @param {dss.IRegistrationSummary=} [properties] Properties to set
         * @returns {dss.RegistrationSummary} RegistrationSummary instance
         */
        RegistrationSummary.create = function create(properties) {
            return new RegistrationSummary(properties);
        };

        /**
         * Encodes the specified RegistrationSummary message. Does not implicitly {@link dss.RegistrationSummary.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationSummary
         * @static
         * @param {dss.IRegistrationSummary} message RegistrationSummary message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationSummary.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
            if (message.firstName != null && Object.hasOwnProperty.call(message, "firstName"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.firstName);
            if (message.lastName != null && Object.hasOwnProperty.call(message, "lastName"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.lastName);
            if (message.email != null && Object.hasOwnProperty.call(message, "email"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.email);
            if (message.createdAt != null && Object.hasOwnProperty.call(message, "createdAt"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.createdAt);
            if (message.paid != null && Object.hasOwnProperty.call(message, "paid"))
                writer.uint32(/* id 6, wireType 0 =*/48).bool(message.paid);
            return writer;
        };

        /**
         * Encodes the specified RegistrationSummary message, length delimited. Does not implicitly {@link dss.RegistrationSummary.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationSummary
         * @static
         * @param {dss.IRegistrationSummary} message RegistrationSummary message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationSummary.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationSummary message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationSummary
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationSummary} RegistrationSummary
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationSummary.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationSummary();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.firstName = reader.string();
                    break;
                case 3:
                    message.lastName = reader.string();
                    break;
                case 4:
                    message.email = reader.string();
                    break;
                case 5:
                    message.createdAt = reader.string();
                    break;
                case 6:
                    message.paid = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationSummary message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationSummary
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationSummary} RegistrationSummary
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationSummary.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationSummary message.
         * @function verify
         * @memberof dss.RegistrationSummary
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationSummary.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isString(message.id))
                    return "id: string expected";
            if (message.firstName != null && message.hasOwnProperty("firstName"))
                if (!$util.isString(message.firstName))
                    return "firstName: string expected";
            if (message.lastName != null && message.hasOwnProperty("lastName"))
                if (!$util.isString(message.lastName))
                    return "lastName: string expected";
            if (message.email != null && message.hasOwnProperty("email"))
                if (!$util.isString(message.email))
                    return "email: string expected";
            if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                if (!$util.isString(message.createdAt))
                    return "createdAt: string expected";
            if (message.paid != null && message.hasOwnProperty("paid"))
                if (typeof message.paid !== "boolean")
                    return "paid: boolean expected";
            return null;
        };

        /**
         * Creates a RegistrationSummary message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationSummary
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationSummary} RegistrationSummary
         */
        RegistrationSummary.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationSummary)
                return object;
            var message = new $root.dss.RegistrationSummary();
            if (object.id != null)
                message.id = String(object.id);
            if (object.firstName != null)
                message.firstName = String(object.firstName);
            if (object.lastName != null)
                message.lastName = String(object.lastName);
            if (object.email != null)
                message.email = String(object.email);
            if (object.createdAt != null)
                message.createdAt = String(object.createdAt);
            if (object.paid != null)
                message.paid = Boolean(object.paid);
            return message;
        };

        /**
         * Creates a plain object from a RegistrationSummary message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationSummary
         * @static
         * @param {dss.RegistrationSummary} message RegistrationSummary
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationSummary.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.id = "";
                object.firstName = "";
                object.lastName = "";
                object.email = "";
                object.createdAt = "";
                object.paid = false;
            }
            if (message.id != null && message.hasOwnProperty("id"))
                object.id = message.id;
            if (message.firstName != null && message.hasOwnProperty("firstName"))
                object.firstName = message.firstName;
            if (message.lastName != null && message.hasOwnProperty("lastName"))
                object.lastName = message.lastName;
            if (message.email != null && message.hasOwnProperty("email"))
                object.email = message.email;
            if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                object.createdAt = message.createdAt;
            if (message.paid != null && message.hasOwnProperty("paid"))
                object.paid = message.paid;
            return object;
        };

        /**
         * Converts this RegistrationSummary to JSON.
         * @function toJSON
         * @memberof dss.RegistrationSummary
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationSummary.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationSummary;
    })();

    dss.RegistrationGetSummaryRes = (function() {

        /**
         * Properties of a RegistrationGetSummaryRes.
         * @memberof dss
         * @interface IRegistrationGetSummaryRes
         * @property {Array.<dss.IRegistrationSummary>|null} [summaries] RegistrationGetSummaryRes summaries
         */

        /**
         * Constructs a new RegistrationGetSummaryRes.
         * @memberof dss
         * @classdesc Represents a RegistrationGetSummaryRes.
         * @implements IRegistrationGetSummaryRes
         * @constructor
         * @param {dss.IRegistrationGetSummaryRes=} [properties] Properties to set
         */
        function RegistrationGetSummaryRes(properties) {
            this.summaries = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RegistrationGetSummaryRes summaries.
         * @member {Array.<dss.IRegistrationSummary>} summaries
         * @memberof dss.RegistrationGetSummaryRes
         * @instance
         */
        RegistrationGetSummaryRes.prototype.summaries = $util.emptyArray;

        /**
         * Creates a new RegistrationGetSummaryRes instance using the specified properties.
         * @function create
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {dss.IRegistrationGetSummaryRes=} [properties] Properties to set
         * @returns {dss.RegistrationGetSummaryRes} RegistrationGetSummaryRes instance
         */
        RegistrationGetSummaryRes.create = function create(properties) {
            return new RegistrationGetSummaryRes(properties);
        };

        /**
         * Encodes the specified RegistrationGetSummaryRes message. Does not implicitly {@link dss.RegistrationGetSummaryRes.verify|verify} messages.
         * @function encode
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {dss.IRegistrationGetSummaryRes} message RegistrationGetSummaryRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetSummaryRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.summaries != null && message.summaries.length)
                for (var i = 0; i < message.summaries.length; ++i)
                    $root.dss.RegistrationSummary.encode(message.summaries[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified RegistrationGetSummaryRes message, length delimited. Does not implicitly {@link dss.RegistrationGetSummaryRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {dss.IRegistrationGetSummaryRes} message RegistrationGetSummaryRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RegistrationGetSummaryRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RegistrationGetSummaryRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.RegistrationGetSummaryRes} RegistrationGetSummaryRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetSummaryRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.RegistrationGetSummaryRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.summaries && message.summaries.length))
                        message.summaries = [];
                    message.summaries.push($root.dss.RegistrationSummary.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RegistrationGetSummaryRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.RegistrationGetSummaryRes} RegistrationGetSummaryRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RegistrationGetSummaryRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RegistrationGetSummaryRes message.
         * @function verify
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RegistrationGetSummaryRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.summaries != null && message.hasOwnProperty("summaries")) {
                if (!Array.isArray(message.summaries))
                    return "summaries: array expected";
                for (var i = 0; i < message.summaries.length; ++i) {
                    var error = $root.dss.RegistrationSummary.verify(message.summaries[i]);
                    if (error)
                        return "summaries." + error;
                }
            }
            return null;
        };

        /**
         * Creates a RegistrationGetSummaryRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.RegistrationGetSummaryRes} RegistrationGetSummaryRes
         */
        RegistrationGetSummaryRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.RegistrationGetSummaryRes)
                return object;
            var message = new $root.dss.RegistrationGetSummaryRes();
            if (object.summaries) {
                if (!Array.isArray(object.summaries))
                    throw TypeError(".dss.RegistrationGetSummaryRes.summaries: array expected");
                message.summaries = [];
                for (var i = 0; i < object.summaries.length; ++i) {
                    if (typeof object.summaries[i] !== "object")
                        throw TypeError(".dss.RegistrationGetSummaryRes.summaries: object expected");
                    message.summaries[i] = $root.dss.RegistrationSummary.fromObject(object.summaries[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a RegistrationGetSummaryRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.RegistrationGetSummaryRes
         * @static
         * @param {dss.RegistrationGetSummaryRes} message RegistrationGetSummaryRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RegistrationGetSummaryRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.summaries = [];
            if (message.summaries && message.summaries.length) {
                object.summaries = [];
                for (var j = 0; j < message.summaries.length; ++j)
                    object.summaries[j] = $root.dss.RegistrationSummary.toObject(message.summaries[j], options);
            }
            return object;
        };

        /**
         * Converts this RegistrationGetSummaryRes to JSON.
         * @function toJSON
         * @memberof dss.RegistrationGetSummaryRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RegistrationGetSummaryRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RegistrationGetSummaryRes;
    })();

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
         * Callback as used by {@link dss.Discount#add}.
         * @memberof dss.Discount
         * @typedef AddCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.DiscountAddRes} [response] DiscountAddRes
         */

        /**
         * Calls Add.
         * @function add
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountAddReq} request DiscountAddReq message or plain object
         * @param {dss.Discount.AddCallback} callback Node-style callback called with the error, if any, and DiscountAddRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Discount.prototype.add = function add(request, callback) {
            return this.rpcCall(add, $root.dss.DiscountAddReq, $root.dss.DiscountAddRes, request, callback);
        }, "name", { value: "Add" });

        /**
         * Calls Add.
         * @function add
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountAddReq} request DiscountAddReq message or plain object
         * @returns {Promise<dss.DiscountAddRes>} Promise
         * @variation 2
         */

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

        /**
         * Callback as used by {@link dss.Discount#list}.
         * @memberof dss.Discount
         * @typedef ListCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.DiscountListRes} [response] DiscountListRes
         */

        /**
         * Calls List.
         * @function list
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountListReq} request DiscountListReq message or plain object
         * @param {dss.Discount.ListCallback} callback Node-style callback called with the error, if any, and DiscountListRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Discount.prototype.list = function list(request, callback) {
            return this.rpcCall(list, $root.dss.DiscountListReq, $root.dss.DiscountListRes, request, callback);
        }, "name", { value: "List" });

        /**
         * Calls List.
         * @function list
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountListReq} request DiscountListReq message or plain object
         * @returns {Promise<dss.DiscountListRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Discount#update}.
         * @memberof dss.Discount
         * @typedef UpdateCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.DiscountUpdateRes} [response] DiscountUpdateRes
         */

        /**
         * Calls Update.
         * @function update
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountUpdateReq} request DiscountUpdateReq message or plain object
         * @param {dss.Discount.UpdateCallback} callback Node-style callback called with the error, if any, and DiscountUpdateRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Discount.prototype.update = function update(request, callback) {
            return this.rpcCall(update, $root.dss.DiscountUpdateReq, $root.dss.DiscountUpdateRes, request, callback);
        }, "name", { value: "Update" });

        /**
         * Calls Update.
         * @function update
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountUpdateReq} request DiscountUpdateReq message or plain object
         * @returns {Promise<dss.DiscountUpdateRes>} Promise
         * @variation 2
         */

        /**
         * Callback as used by {@link dss.Discount#delete_}.
         * @memberof dss.Discount
         * @typedef DeleteCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {dss.DiscountDeleteRes} [response] DiscountDeleteRes
         */

        /**
         * Calls Delete.
         * @function delete
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountDeleteReq} request DiscountDeleteReq message or plain object
         * @param {dss.Discount.DeleteCallback} callback Node-style callback called with the error, if any, and DiscountDeleteRes
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(Discount.prototype["delete"] = function delete_(request, callback) {
            return this.rpcCall(delete_, $root.dss.DiscountDeleteReq, $root.dss.DiscountDeleteRes, request, callback);
        }, "name", { value: "Delete" });

        /**
         * Calls Delete.
         * @function delete
         * @memberof dss.Discount
         * @instance
         * @param {dss.IDiscountDeleteReq} request DiscountDeleteReq message or plain object
         * @returns {Promise<dss.DiscountDeleteRes>} Promise
         * @variation 2
         */

        return Discount;
    })();

    /**
     * PurchaseItem enum.
     * @name dss.PurchaseItem
     * @enum {number}
     * @property {number} FullWeekendPassPurchaseItem=0 FullWeekendPassPurchaseItem value
     * @property {number} DanceOnlyPassPurchaseItem=1 DanceOnlyPassPurchaseItem value
     * @property {number} MixAndMatchPurchaseItem=2 MixAndMatchPurchaseItem value
     * @property {number} SoloJazzPurchaseItem=3 SoloJazzPurchaseItem value
     * @property {number} TeamCompetitionPurchaseItem=4 TeamCompetitionPurchaseItem value
     * @property {number} TShirtPurchaseItem=5 TShirtPurchaseItem value
     */
    dss.PurchaseItem = (function() {
        var valuesById = {}, values = Object.create(valuesById);
        values[valuesById[0] = "FullWeekendPassPurchaseItem"] = 0;
        values[valuesById[1] = "DanceOnlyPassPurchaseItem"] = 1;
        values[valuesById[2] = "MixAndMatchPurchaseItem"] = 2;
        values[valuesById[3] = "SoloJazzPurchaseItem"] = 3;
        values[valuesById[4] = "TeamCompetitionPurchaseItem"] = 4;
        values[valuesById[5] = "TShirtPurchaseItem"] = 5;
        return values;
    })();

    dss.DiscountAmount = (function() {

        /**
         * Properties of a DiscountAmount.
         * @memberof dss
         * @interface IDiscountAmount
         * @property {number|Long|null} [dollar] DiscountAmount dollar
         * @property {string|null} [percent] DiscountAmount percent
         * @property {google.protobuf.IEmpty|null} [squareNotFound] DiscountAmount squareNotFound
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

        /**
         * DiscountAmount squareNotFound.
         * @member {google.protobuf.IEmpty|null|undefined} squareNotFound
         * @memberof dss.DiscountAmount
         * @instance
         */
        DiscountAmount.prototype.squareNotFound = null;

        // OneOf field names bound to virtual getters and setters
        var $oneOfFields;

        /**
         * DiscountAmount amount.
         * @member {"dollar"|"percent"|"squareNotFound"|undefined} amount
         * @memberof dss.DiscountAmount
         * @instance
         */
        Object.defineProperty(DiscountAmount.prototype, "amount", {
            get: $util.oneOfGetter($oneOfFields = ["dollar", "percent", "squareNotFound"]),
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
            if (message.squareNotFound != null && Object.hasOwnProperty.call(message, "squareNotFound"))
                $root.google.protobuf.Empty.encode(message.squareNotFound, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
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
                case 3:
                    message.squareNotFound = $root.google.protobuf.Empty.decode(reader, reader.uint32());
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
            if (message.squareNotFound != null && message.hasOwnProperty("squareNotFound")) {
                if (properties.amount === 1)
                    return "amount: multiple values";
                properties.amount = 1;
                {
                    var error = $root.google.protobuf.Empty.verify(message.squareNotFound);
                    if (error)
                        return "squareNotFound." + error;
                }
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
            if (object.squareNotFound != null) {
                if (typeof object.squareNotFound !== "object")
                    throw TypeError(".dss.DiscountAmount.squareNotFound: object expected");
                message.squareNotFound = $root.google.protobuf.Empty.fromObject(object.squareNotFound);
            }
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
            if (message.squareNotFound != null && message.hasOwnProperty("squareNotFound")) {
                object.squareNotFound = $root.google.protobuf.Empty.toObject(message.squareNotFound, options);
                if (options.oneofs)
                    object.amount = "squareNotFound";
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
         * @property {string|null} [name] SingleDiscount name
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
         * SingleDiscount name.
         * @member {string} name
         * @memberof dss.SingleDiscount
         * @instance
         */
        SingleDiscount.prototype.name = "";

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
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
            if (message.amount != null && Object.hasOwnProperty.call(message, "amount"))
                $root.dss.DiscountAmount.encode(message.amount, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            if (message.appliedTo != null && Object.hasOwnProperty.call(message, "appliedTo"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.appliedTo);
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
                    message.name = reader.string();
                    break;
                case 2:
                    message.amount = $root.dss.DiscountAmount.decode(reader, reader.uint32());
                    break;
                case 3:
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
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
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
            if (object.name != null)
                message.name = String(object.name);
            if (object.amount != null) {
                if (typeof object.amount !== "object")
                    throw TypeError(".dss.SingleDiscount.amount: object expected");
                message.amount = $root.dss.DiscountAmount.fromObject(object.amount);
            }
            switch (object.appliedTo) {
            case "FullWeekendPassPurchaseItem":
            case 0:
                message.appliedTo = 0;
                break;
            case "DanceOnlyPassPurchaseItem":
            case 1:
                message.appliedTo = 1;
                break;
            case "MixAndMatchPurchaseItem":
            case 2:
                message.appliedTo = 2;
                break;
            case "SoloJazzPurchaseItem":
            case 3:
                message.appliedTo = 3;
                break;
            case "TeamCompetitionPurchaseItem":
            case 4:
                message.appliedTo = 4;
                break;
            case "TShirtPurchaseItem":
            case 5:
                message.appliedTo = 5;
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
                object.name = "";
                object.amount = null;
                object.appliedTo = options.enums === String ? "FullWeekendPassPurchaseItem" : 0;
            }
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
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

    dss.DiscountAddReq = (function() {

        /**
         * Properties of a DiscountAddReq.
         * @memberof dss
         * @interface IDiscountAddReq
         * @property {dss.IDiscountBundle|null} [bundle] DiscountAddReq bundle
         */

        /**
         * Constructs a new DiscountAddReq.
         * @memberof dss
         * @classdesc Represents a DiscountAddReq.
         * @implements IDiscountAddReq
         * @constructor
         * @param {dss.IDiscountAddReq=} [properties] Properties to set
         */
        function DiscountAddReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountAddReq bundle.
         * @member {dss.IDiscountBundle|null|undefined} bundle
         * @memberof dss.DiscountAddReq
         * @instance
         */
        DiscountAddReq.prototype.bundle = null;

        /**
         * Creates a new DiscountAddReq instance using the specified properties.
         * @function create
         * @memberof dss.DiscountAddReq
         * @static
         * @param {dss.IDiscountAddReq=} [properties] Properties to set
         * @returns {dss.DiscountAddReq} DiscountAddReq instance
         */
        DiscountAddReq.create = function create(properties) {
            return new DiscountAddReq(properties);
        };

        /**
         * Encodes the specified DiscountAddReq message. Does not implicitly {@link dss.DiscountAddReq.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountAddReq
         * @static
         * @param {dss.IDiscountAddReq} message DiscountAddReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountAddReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.bundle != null && Object.hasOwnProperty.call(message, "bundle"))
                $root.dss.DiscountBundle.encode(message.bundle, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified DiscountAddReq message, length delimited. Does not implicitly {@link dss.DiscountAddReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountAddReq
         * @static
         * @param {dss.IDiscountAddReq} message DiscountAddReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountAddReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountAddReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountAddReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountAddReq} DiscountAddReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountAddReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountAddReq();
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
         * Decodes a DiscountAddReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountAddReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountAddReq} DiscountAddReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountAddReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountAddReq message.
         * @function verify
         * @memberof dss.DiscountAddReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountAddReq.verify = function verify(message) {
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
         * Creates a DiscountAddReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountAddReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountAddReq} DiscountAddReq
         */
        DiscountAddReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountAddReq)
                return object;
            var message = new $root.dss.DiscountAddReq();
            if (object.bundle != null) {
                if (typeof object.bundle !== "object")
                    throw TypeError(".dss.DiscountAddReq.bundle: object expected");
                message.bundle = $root.dss.DiscountBundle.fromObject(object.bundle);
            }
            return message;
        };

        /**
         * Creates a plain object from a DiscountAddReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountAddReq
         * @static
         * @param {dss.DiscountAddReq} message DiscountAddReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountAddReq.toObject = function toObject(message, options) {
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
         * Converts this DiscountAddReq to JSON.
         * @function toJSON
         * @memberof dss.DiscountAddReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountAddReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountAddReq;
    })();

    dss.DiscountAddRes = (function() {

        /**
         * Properties of a DiscountAddRes.
         * @memberof dss
         * @interface IDiscountAddRes
         */

        /**
         * Constructs a new DiscountAddRes.
         * @memberof dss
         * @classdesc Represents a DiscountAddRes.
         * @implements IDiscountAddRes
         * @constructor
         * @param {dss.IDiscountAddRes=} [properties] Properties to set
         */
        function DiscountAddRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new DiscountAddRes instance using the specified properties.
         * @function create
         * @memberof dss.DiscountAddRes
         * @static
         * @param {dss.IDiscountAddRes=} [properties] Properties to set
         * @returns {dss.DiscountAddRes} DiscountAddRes instance
         */
        DiscountAddRes.create = function create(properties) {
            return new DiscountAddRes(properties);
        };

        /**
         * Encodes the specified DiscountAddRes message. Does not implicitly {@link dss.DiscountAddRes.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountAddRes
         * @static
         * @param {dss.IDiscountAddRes} message DiscountAddRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountAddRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified DiscountAddRes message, length delimited. Does not implicitly {@link dss.DiscountAddRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountAddRes
         * @static
         * @param {dss.IDiscountAddRes} message DiscountAddRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountAddRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountAddRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountAddRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountAddRes} DiscountAddRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountAddRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountAddRes();
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
         * Decodes a DiscountAddRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountAddRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountAddRes} DiscountAddRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountAddRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountAddRes message.
         * @function verify
         * @memberof dss.DiscountAddRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountAddRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a DiscountAddRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountAddRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountAddRes} DiscountAddRes
         */
        DiscountAddRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountAddRes)
                return object;
            return new $root.dss.DiscountAddRes();
        };

        /**
         * Creates a plain object from a DiscountAddRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountAddRes
         * @static
         * @param {dss.DiscountAddRes} message DiscountAddRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountAddRes.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this DiscountAddRes to JSON.
         * @function toJSON
         * @memberof dss.DiscountAddRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountAddRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountAddRes;
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

    dss.DiscountListReq = (function() {

        /**
         * Properties of a DiscountListReq.
         * @memberof dss
         * @interface IDiscountListReq
         */

        /**
         * Constructs a new DiscountListReq.
         * @memberof dss
         * @classdesc Represents a DiscountListReq.
         * @implements IDiscountListReq
         * @constructor
         * @param {dss.IDiscountListReq=} [properties] Properties to set
         */
        function DiscountListReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new DiscountListReq instance using the specified properties.
         * @function create
         * @memberof dss.DiscountListReq
         * @static
         * @param {dss.IDiscountListReq=} [properties] Properties to set
         * @returns {dss.DiscountListReq} DiscountListReq instance
         */
        DiscountListReq.create = function create(properties) {
            return new DiscountListReq(properties);
        };

        /**
         * Encodes the specified DiscountListReq message. Does not implicitly {@link dss.DiscountListReq.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountListReq
         * @static
         * @param {dss.IDiscountListReq} message DiscountListReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountListReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified DiscountListReq message, length delimited. Does not implicitly {@link dss.DiscountListReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountListReq
         * @static
         * @param {dss.IDiscountListReq} message DiscountListReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountListReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountListReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountListReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountListReq} DiscountListReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountListReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountListReq();
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
         * Decodes a DiscountListReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountListReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountListReq} DiscountListReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountListReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountListReq message.
         * @function verify
         * @memberof dss.DiscountListReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountListReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a DiscountListReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountListReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountListReq} DiscountListReq
         */
        DiscountListReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountListReq)
                return object;
            return new $root.dss.DiscountListReq();
        };

        /**
         * Creates a plain object from a DiscountListReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountListReq
         * @static
         * @param {dss.DiscountListReq} message DiscountListReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountListReq.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this DiscountListReq to JSON.
         * @function toJSON
         * @memberof dss.DiscountListReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountListReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountListReq;
    })();

    dss.DiscountListRes = (function() {

        /**
         * Properties of a DiscountListRes.
         * @memberof dss
         * @interface IDiscountListRes
         * @property {Array.<dss.IDiscountBundle>|null} [bundles] DiscountListRes bundles
         */

        /**
         * Constructs a new DiscountListRes.
         * @memberof dss
         * @classdesc Represents a DiscountListRes.
         * @implements IDiscountListRes
         * @constructor
         * @param {dss.IDiscountListRes=} [properties] Properties to set
         */
        function DiscountListRes(properties) {
            this.bundles = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountListRes bundles.
         * @member {Array.<dss.IDiscountBundle>} bundles
         * @memberof dss.DiscountListRes
         * @instance
         */
        DiscountListRes.prototype.bundles = $util.emptyArray;

        /**
         * Creates a new DiscountListRes instance using the specified properties.
         * @function create
         * @memberof dss.DiscountListRes
         * @static
         * @param {dss.IDiscountListRes=} [properties] Properties to set
         * @returns {dss.DiscountListRes} DiscountListRes instance
         */
        DiscountListRes.create = function create(properties) {
            return new DiscountListRes(properties);
        };

        /**
         * Encodes the specified DiscountListRes message. Does not implicitly {@link dss.DiscountListRes.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountListRes
         * @static
         * @param {dss.IDiscountListRes} message DiscountListRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountListRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.bundles != null && message.bundles.length)
                for (var i = 0; i < message.bundles.length; ++i)
                    $root.dss.DiscountBundle.encode(message.bundles[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified DiscountListRes message, length delimited. Does not implicitly {@link dss.DiscountListRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountListRes
         * @static
         * @param {dss.IDiscountListRes} message DiscountListRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountListRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountListRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountListRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountListRes} DiscountListRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountListRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountListRes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.bundles && message.bundles.length))
                        message.bundles = [];
                    message.bundles.push($root.dss.DiscountBundle.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a DiscountListRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountListRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountListRes} DiscountListRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountListRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountListRes message.
         * @function verify
         * @memberof dss.DiscountListRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountListRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.bundles != null && message.hasOwnProperty("bundles")) {
                if (!Array.isArray(message.bundles))
                    return "bundles: array expected";
                for (var i = 0; i < message.bundles.length; ++i) {
                    var error = $root.dss.DiscountBundle.verify(message.bundles[i]);
                    if (error)
                        return "bundles." + error;
                }
            }
            return null;
        };

        /**
         * Creates a DiscountListRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountListRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountListRes} DiscountListRes
         */
        DiscountListRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountListRes)
                return object;
            var message = new $root.dss.DiscountListRes();
            if (object.bundles) {
                if (!Array.isArray(object.bundles))
                    throw TypeError(".dss.DiscountListRes.bundles: array expected");
                message.bundles = [];
                for (var i = 0; i < object.bundles.length; ++i) {
                    if (typeof object.bundles[i] !== "object")
                        throw TypeError(".dss.DiscountListRes.bundles: object expected");
                    message.bundles[i] = $root.dss.DiscountBundle.fromObject(object.bundles[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a DiscountListRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountListRes
         * @static
         * @param {dss.DiscountListRes} message DiscountListRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountListRes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.bundles = [];
            if (message.bundles && message.bundles.length) {
                object.bundles = [];
                for (var j = 0; j < message.bundles.length; ++j)
                    object.bundles[j] = $root.dss.DiscountBundle.toObject(message.bundles[j], options);
            }
            return object;
        };

        /**
         * Converts this DiscountListRes to JSON.
         * @function toJSON
         * @memberof dss.DiscountListRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountListRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountListRes;
    })();

    dss.DiscountUpdateReq = (function() {

        /**
         * Properties of a DiscountUpdateReq.
         * @memberof dss
         * @interface IDiscountUpdateReq
         * @property {string|null} [oldCode] DiscountUpdateReq oldCode
         * @property {dss.IDiscountBundle|null} [bundle] DiscountUpdateReq bundle
         */

        /**
         * Constructs a new DiscountUpdateReq.
         * @memberof dss
         * @classdesc Represents a DiscountUpdateReq.
         * @implements IDiscountUpdateReq
         * @constructor
         * @param {dss.IDiscountUpdateReq=} [properties] Properties to set
         */
        function DiscountUpdateReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountUpdateReq oldCode.
         * @member {string} oldCode
         * @memberof dss.DiscountUpdateReq
         * @instance
         */
        DiscountUpdateReq.prototype.oldCode = "";

        /**
         * DiscountUpdateReq bundle.
         * @member {dss.IDiscountBundle|null|undefined} bundle
         * @memberof dss.DiscountUpdateReq
         * @instance
         */
        DiscountUpdateReq.prototype.bundle = null;

        /**
         * Creates a new DiscountUpdateReq instance using the specified properties.
         * @function create
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {dss.IDiscountUpdateReq=} [properties] Properties to set
         * @returns {dss.DiscountUpdateReq} DiscountUpdateReq instance
         */
        DiscountUpdateReq.create = function create(properties) {
            return new DiscountUpdateReq(properties);
        };

        /**
         * Encodes the specified DiscountUpdateReq message. Does not implicitly {@link dss.DiscountUpdateReq.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {dss.IDiscountUpdateReq} message DiscountUpdateReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountUpdateReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.oldCode != null && Object.hasOwnProperty.call(message, "oldCode"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.oldCode);
            if (message.bundle != null && Object.hasOwnProperty.call(message, "bundle"))
                $root.dss.DiscountBundle.encode(message.bundle, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified DiscountUpdateReq message, length delimited. Does not implicitly {@link dss.DiscountUpdateReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {dss.IDiscountUpdateReq} message DiscountUpdateReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountUpdateReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountUpdateReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountUpdateReq} DiscountUpdateReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountUpdateReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountUpdateReq();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.oldCode = reader.string();
                    break;
                case 2:
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
         * Decodes a DiscountUpdateReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountUpdateReq} DiscountUpdateReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountUpdateReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountUpdateReq message.
         * @function verify
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountUpdateReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.oldCode != null && message.hasOwnProperty("oldCode"))
                if (!$util.isString(message.oldCode))
                    return "oldCode: string expected";
            if (message.bundle != null && message.hasOwnProperty("bundle")) {
                var error = $root.dss.DiscountBundle.verify(message.bundle);
                if (error)
                    return "bundle." + error;
            }
            return null;
        };

        /**
         * Creates a DiscountUpdateReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountUpdateReq} DiscountUpdateReq
         */
        DiscountUpdateReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountUpdateReq)
                return object;
            var message = new $root.dss.DiscountUpdateReq();
            if (object.oldCode != null)
                message.oldCode = String(object.oldCode);
            if (object.bundle != null) {
                if (typeof object.bundle !== "object")
                    throw TypeError(".dss.DiscountUpdateReq.bundle: object expected");
                message.bundle = $root.dss.DiscountBundle.fromObject(object.bundle);
            }
            return message;
        };

        /**
         * Creates a plain object from a DiscountUpdateReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountUpdateReq
         * @static
         * @param {dss.DiscountUpdateReq} message DiscountUpdateReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountUpdateReq.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.oldCode = "";
                object.bundle = null;
            }
            if (message.oldCode != null && message.hasOwnProperty("oldCode"))
                object.oldCode = message.oldCode;
            if (message.bundle != null && message.hasOwnProperty("bundle"))
                object.bundle = $root.dss.DiscountBundle.toObject(message.bundle, options);
            return object;
        };

        /**
         * Converts this DiscountUpdateReq to JSON.
         * @function toJSON
         * @memberof dss.DiscountUpdateReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountUpdateReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountUpdateReq;
    })();

    dss.DiscountUpdateRes = (function() {

        /**
         * Properties of a DiscountUpdateRes.
         * @memberof dss
         * @interface IDiscountUpdateRes
         */

        /**
         * Constructs a new DiscountUpdateRes.
         * @memberof dss
         * @classdesc Represents a DiscountUpdateRes.
         * @implements IDiscountUpdateRes
         * @constructor
         * @param {dss.IDiscountUpdateRes=} [properties] Properties to set
         */
        function DiscountUpdateRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new DiscountUpdateRes instance using the specified properties.
         * @function create
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {dss.IDiscountUpdateRes=} [properties] Properties to set
         * @returns {dss.DiscountUpdateRes} DiscountUpdateRes instance
         */
        DiscountUpdateRes.create = function create(properties) {
            return new DiscountUpdateRes(properties);
        };

        /**
         * Encodes the specified DiscountUpdateRes message. Does not implicitly {@link dss.DiscountUpdateRes.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {dss.IDiscountUpdateRes} message DiscountUpdateRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountUpdateRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified DiscountUpdateRes message, length delimited. Does not implicitly {@link dss.DiscountUpdateRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {dss.IDiscountUpdateRes} message DiscountUpdateRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountUpdateRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountUpdateRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountUpdateRes} DiscountUpdateRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountUpdateRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountUpdateRes();
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
         * Decodes a DiscountUpdateRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountUpdateRes} DiscountUpdateRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountUpdateRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountUpdateRes message.
         * @function verify
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountUpdateRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a DiscountUpdateRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountUpdateRes} DiscountUpdateRes
         */
        DiscountUpdateRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountUpdateRes)
                return object;
            return new $root.dss.DiscountUpdateRes();
        };

        /**
         * Creates a plain object from a DiscountUpdateRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountUpdateRes
         * @static
         * @param {dss.DiscountUpdateRes} message DiscountUpdateRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountUpdateRes.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this DiscountUpdateRes to JSON.
         * @function toJSON
         * @memberof dss.DiscountUpdateRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountUpdateRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountUpdateRes;
    })();

    dss.DiscountDeleteReq = (function() {

        /**
         * Properties of a DiscountDeleteReq.
         * @memberof dss
         * @interface IDiscountDeleteReq
         * @property {string|null} [code] DiscountDeleteReq code
         */

        /**
         * Constructs a new DiscountDeleteReq.
         * @memberof dss
         * @classdesc Represents a DiscountDeleteReq.
         * @implements IDiscountDeleteReq
         * @constructor
         * @param {dss.IDiscountDeleteReq=} [properties] Properties to set
         */
        function DiscountDeleteReq(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * DiscountDeleteReq code.
         * @member {string} code
         * @memberof dss.DiscountDeleteReq
         * @instance
         */
        DiscountDeleteReq.prototype.code = "";

        /**
         * Creates a new DiscountDeleteReq instance using the specified properties.
         * @function create
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {dss.IDiscountDeleteReq=} [properties] Properties to set
         * @returns {dss.DiscountDeleteReq} DiscountDeleteReq instance
         */
        DiscountDeleteReq.create = function create(properties) {
            return new DiscountDeleteReq(properties);
        };

        /**
         * Encodes the specified DiscountDeleteReq message. Does not implicitly {@link dss.DiscountDeleteReq.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {dss.IDiscountDeleteReq} message DiscountDeleteReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountDeleteReq.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.code != null && Object.hasOwnProperty.call(message, "code"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.code);
            return writer;
        };

        /**
         * Encodes the specified DiscountDeleteReq message, length delimited. Does not implicitly {@link dss.DiscountDeleteReq.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {dss.IDiscountDeleteReq} message DiscountDeleteReq message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountDeleteReq.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountDeleteReq message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountDeleteReq} DiscountDeleteReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountDeleteReq.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountDeleteReq();
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
         * Decodes a DiscountDeleteReq message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountDeleteReq} DiscountDeleteReq
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountDeleteReq.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountDeleteReq message.
         * @function verify
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountDeleteReq.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.code != null && message.hasOwnProperty("code"))
                if (!$util.isString(message.code))
                    return "code: string expected";
            return null;
        };

        /**
         * Creates a DiscountDeleteReq message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountDeleteReq} DiscountDeleteReq
         */
        DiscountDeleteReq.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountDeleteReq)
                return object;
            var message = new $root.dss.DiscountDeleteReq();
            if (object.code != null)
                message.code = String(object.code);
            return message;
        };

        /**
         * Creates a plain object from a DiscountDeleteReq message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountDeleteReq
         * @static
         * @param {dss.DiscountDeleteReq} message DiscountDeleteReq
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountDeleteReq.toObject = function toObject(message, options) {
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
         * Converts this DiscountDeleteReq to JSON.
         * @function toJSON
         * @memberof dss.DiscountDeleteReq
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountDeleteReq.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountDeleteReq;
    })();

    dss.DiscountDeleteRes = (function() {

        /**
         * Properties of a DiscountDeleteRes.
         * @memberof dss
         * @interface IDiscountDeleteRes
         */

        /**
         * Constructs a new DiscountDeleteRes.
         * @memberof dss
         * @classdesc Represents a DiscountDeleteRes.
         * @implements IDiscountDeleteRes
         * @constructor
         * @param {dss.IDiscountDeleteRes=} [properties] Properties to set
         */
        function DiscountDeleteRes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new DiscountDeleteRes instance using the specified properties.
         * @function create
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {dss.IDiscountDeleteRes=} [properties] Properties to set
         * @returns {dss.DiscountDeleteRes} DiscountDeleteRes instance
         */
        DiscountDeleteRes.create = function create(properties) {
            return new DiscountDeleteRes(properties);
        };

        /**
         * Encodes the specified DiscountDeleteRes message. Does not implicitly {@link dss.DiscountDeleteRes.verify|verify} messages.
         * @function encode
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {dss.IDiscountDeleteRes} message DiscountDeleteRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountDeleteRes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            return writer;
        };

        /**
         * Encodes the specified DiscountDeleteRes message, length delimited. Does not implicitly {@link dss.DiscountDeleteRes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {dss.IDiscountDeleteRes} message DiscountDeleteRes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        DiscountDeleteRes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a DiscountDeleteRes message from the specified reader or buffer.
         * @function decode
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {dss.DiscountDeleteRes} DiscountDeleteRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountDeleteRes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.dss.DiscountDeleteRes();
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
         * Decodes a DiscountDeleteRes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {dss.DiscountDeleteRes} DiscountDeleteRes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        DiscountDeleteRes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a DiscountDeleteRes message.
         * @function verify
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        DiscountDeleteRes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            return null;
        };

        /**
         * Creates a DiscountDeleteRes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {dss.DiscountDeleteRes} DiscountDeleteRes
         */
        DiscountDeleteRes.fromObject = function fromObject(object) {
            if (object instanceof $root.dss.DiscountDeleteRes)
                return object;
            return new $root.dss.DiscountDeleteRes();
        };

        /**
         * Creates a plain object from a DiscountDeleteRes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof dss.DiscountDeleteRes
         * @static
         * @param {dss.DiscountDeleteRes} message DiscountDeleteRes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        DiscountDeleteRes.toObject = function toObject() {
            return {};
        };

        /**
         * Converts this DiscountDeleteRes to JSON.
         * @function toJSON
         * @memberof dss.DiscountDeleteRes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        DiscountDeleteRes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return DiscountDeleteRes;
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

        protobuf.Empty = (function() {

            /**
             * Properties of an Empty.
             * @memberof google.protobuf
             * @interface IEmpty
             */

            /**
             * Constructs a new Empty.
             * @memberof google.protobuf
             * @classdesc Represents an Empty.
             * @implements IEmpty
             * @constructor
             * @param {google.protobuf.IEmpty=} [properties] Properties to set
             */
            function Empty(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Creates a new Empty instance using the specified properties.
             * @function create
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.IEmpty=} [properties] Properties to set
             * @returns {google.protobuf.Empty} Empty instance
             */
            Empty.create = function create(properties) {
                return new Empty(properties);
            };

            /**
             * Encodes the specified Empty message. Does not implicitly {@link google.protobuf.Empty.verify|verify} messages.
             * @function encode
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.IEmpty} message Empty message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Empty.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                return writer;
            };

            /**
             * Encodes the specified Empty message, length delimited. Does not implicitly {@link google.protobuf.Empty.verify|verify} messages.
             * @function encodeDelimited
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.IEmpty} message Empty message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Empty.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes an Empty message from the specified reader or buffer.
             * @function decode
             * @memberof google.protobuf.Empty
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {google.protobuf.Empty} Empty
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Empty.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.google.protobuf.Empty();
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
             * Decodes an Empty message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof google.protobuf.Empty
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {google.protobuf.Empty} Empty
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Empty.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies an Empty message.
             * @function verify
             * @memberof google.protobuf.Empty
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Empty.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                return null;
            };

            /**
             * Creates an Empty message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof google.protobuf.Empty
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {google.protobuf.Empty} Empty
             */
            Empty.fromObject = function fromObject(object) {
                if (object instanceof $root.google.protobuf.Empty)
                    return object;
                return new $root.google.protobuf.Empty();
            };

            /**
             * Creates a plain object from an Empty message. Also converts values to other types if specified.
             * @function toObject
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.Empty} message Empty
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Empty.toObject = function toObject() {
                return {};
            };

            /**
             * Converts this Empty to JSON.
             * @function toJSON
             * @memberof google.protobuf.Empty
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Empty.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Empty;
        })();

        return protobuf;
    })();

    return google;
})();

module.exports = $root;
