import * as $protobuf from "protobufjs";
export namespace dss {

    class Registration extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): Registration;
        public add(request: dss.IRegistrationAddReq, callback: dss.Registration.AddCallback): void;
        public add(request: dss.IRegistrationAddReq): Promise<dss.RegistrationAddRes>;
        public get(request: dss.IRegistrationGetReq, callback: dss.Registration.GetCallback): void;
        public get(request: dss.IRegistrationGetReq): Promise<dss.RegistrationGetRes>;
        public getSummary(request: dss.IRegistrationGetSummaryReq, callback: dss.Registration.GetSummaryCallback): void;
        public getSummary(request: dss.IRegistrationGetSummaryReq): Promise<dss.RegistrationGetSummaryRes>;
        public prices(request: dss.IRegistrationPricesReq, callback: dss.Registration.PricesCallback): void;
        public prices(request: dss.IRegistrationPricesReq): Promise<dss.RegistrationPricesRes>;
        public update(request: dss.IRegistrationUpdateReq, callback: dss.Registration.UpdateCallback): void;
        public update(request: dss.IRegistrationUpdateReq): Promise<dss.RegistrationUpdateRes>;
    }

    namespace Registration {

        type AddCallback = (error: (Error|null), response?: dss.RegistrationAddRes) => void;

        type GetCallback = (error: (Error|null), response?: dss.RegistrationGetRes) => void;

        type GetSummaryCallback = (error: (Error|null), response?: dss.RegistrationGetSummaryRes) => void;

        type PricesCallback = (error: (Error|null), response?: dss.RegistrationPricesRes) => void;

        type UpdateCallback = (error: (Error|null), response?: dss.RegistrationUpdateRes) => void;
    }

    interface IRegistrationInfo {
        id?: (string|null);
        firstName?: (string|null);
        lastName?: (string|null);
        streetAddress?: (string|null);
        city?: (string|null);
        state?: (string|null);
        zipCode?: (string|null);
        email?: (string|null);
        homeScene?: (string|null);
        isStudent?: (boolean|null);
        fullWeekendPass?: (dss.IFullWeekendPass|null);
        danceOnlyPass?: (dss.IDanceOnlyPass|null);
        noPass?: (dss.INoPass|null);
        mixAndMatch?: (dss.IMixAndMatch|null);
        soloJazz?: (dss.ISoloJazz|null);
        teamCompetition?: (dss.ITeamCompetition|null);
        tshirt?: (dss.ITShirt|null);
        provideHousing?: (dss.IProvideHousing|null);
        requireHousing?: (dss.IRequireHousing|null);
        noHousing?: (dss.INoHousing|null);
        discountCodes?: (string[]|null);
        createdAt?: (string|null);
    }

    class RegistrationInfo implements IRegistrationInfo {
        constructor(properties?: dss.IRegistrationInfo);
        public id: string;
        public firstName: string;
        public lastName: string;
        public streetAddress: string;
        public city: string;
        public state: string;
        public zipCode: string;
        public email: string;
        public homeScene: string;
        public isStudent: boolean;
        public fullWeekendPass?: (dss.IFullWeekendPass|null);
        public danceOnlyPass?: (dss.IDanceOnlyPass|null);
        public noPass?: (dss.INoPass|null);
        public mixAndMatch?: (dss.IMixAndMatch|null);
        public soloJazz?: (dss.ISoloJazz|null);
        public teamCompetition?: (dss.ITeamCompetition|null);
        public tshirt?: (dss.ITShirt|null);
        public provideHousing?: (dss.IProvideHousing|null);
        public requireHousing?: (dss.IRequireHousing|null);
        public noHousing?: (dss.INoHousing|null);
        public discountCodes: string[];
        public createdAt: string;
        public passType?: ("fullWeekendPass"|"danceOnlyPass"|"noPass");
        public housing?: ("provideHousing"|"requireHousing"|"noHousing");
        public static create(properties?: dss.IRegistrationInfo): dss.RegistrationInfo;
        public static encode(message: dss.IRegistrationInfo, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationInfo, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationInfo;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationInfo;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationInfo;
        public static toObject(message: dss.RegistrationInfo, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    enum FullWeekendPassTier {
        Tier1 = 0,
        Tier2 = 1,
        Tier3 = 2,
        Tier4 = 3,
        Tier5 = 4
    }

    enum FullWeekendPassLevel {
        Level1 = 0,
        Level2 = 1,
        Level3 = 2
    }

    interface IFullWeekendPass {
        tier?: (dss.FullWeekendPassTier|null);
        level?: (dss.FullWeekendPassLevel|null);
        paid?: (boolean|null);
    }

    class FullWeekendPass implements IFullWeekendPass {
        constructor(properties?: dss.IFullWeekendPass);
        public tier: dss.FullWeekendPassTier;
        public level: dss.FullWeekendPassLevel;
        public paid: boolean;
        public static create(properties?: dss.IFullWeekendPass): dss.FullWeekendPass;
        public static encode(message: dss.IFullWeekendPass, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IFullWeekendPass, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.FullWeekendPass;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.FullWeekendPass;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.FullWeekendPass;
        public static toObject(message: dss.FullWeekendPass, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDanceOnlyPass {
        paid?: (boolean|null);
    }

    class DanceOnlyPass implements IDanceOnlyPass {
        constructor(properties?: dss.IDanceOnlyPass);
        public paid: boolean;
        public static create(properties?: dss.IDanceOnlyPass): dss.DanceOnlyPass;
        public static encode(message: dss.IDanceOnlyPass, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDanceOnlyPass, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DanceOnlyPass;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DanceOnlyPass;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DanceOnlyPass;
        public static toObject(message: dss.DanceOnlyPass, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface INoPass {
    }

    class NoPass implements INoPass {
        constructor(properties?: dss.INoPass);
        public static create(properties?: dss.INoPass): dss.NoPass;
        public static encode(message: dss.INoPass, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.INoPass, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.NoPass;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.NoPass;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.NoPass;
        public static toObject(message: dss.NoPass, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IMixAndMatch {
        role?: (dss.MixAndMatch.Role|null);
        paid?: (boolean|null);
    }

    class MixAndMatch implements IMixAndMatch {
        constructor(properties?: dss.IMixAndMatch);
        public role: dss.MixAndMatch.Role;
        public paid: boolean;
        public static create(properties?: dss.IMixAndMatch): dss.MixAndMatch;
        public static encode(message: dss.IMixAndMatch, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IMixAndMatch, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.MixAndMatch;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.MixAndMatch;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.MixAndMatch;
        public static toObject(message: dss.MixAndMatch, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    namespace MixAndMatch {

        enum Role {
            Follower = 0,
            Leader = 1
        }
    }

    interface ISoloJazz {
        paid?: (boolean|null);
    }

    class SoloJazz implements ISoloJazz {
        constructor(properties?: dss.ISoloJazz);
        public paid: boolean;
        public static create(properties?: dss.ISoloJazz): dss.SoloJazz;
        public static encode(message: dss.ISoloJazz, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.ISoloJazz, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.SoloJazz;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.SoloJazz;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.SoloJazz;
        public static toObject(message: dss.SoloJazz, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface ITeamCompetition {
        name?: (string|null);
        paid?: (boolean|null);
    }

    class TeamCompetition implements ITeamCompetition {
        constructor(properties?: dss.ITeamCompetition);
        public name: string;
        public paid: boolean;
        public static create(properties?: dss.ITeamCompetition): dss.TeamCompetition;
        public static encode(message: dss.ITeamCompetition, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.ITeamCompetition, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.TeamCompetition;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.TeamCompetition;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.TeamCompetition;
        public static toObject(message: dss.TeamCompetition, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface ITShirt {
        style?: (dss.TShirt.Style|null);
        paid?: (boolean|null);
    }

    class TShirt implements ITShirt {
        constructor(properties?: dss.ITShirt);
        public style: dss.TShirt.Style;
        public paid: boolean;
        public static create(properties?: dss.ITShirt): dss.TShirt;
        public static encode(message: dss.ITShirt, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.ITShirt, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.TShirt;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.TShirt;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.TShirt;
        public static toObject(message: dss.TShirt, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    namespace TShirt {

        enum Style {
            UnisexS = 0,
            UnisexM = 1,
            UnisexL = 2,
            UnisexXL = 3,
            Unisex2XL = 4,
            Unisex3XL = 5,
            BellaS = 6,
            BellaM = 7,
            BellaL = 8,
            BellaXL = 9,
            Bella2XL = 10
        }
    }

    interface IProvideHousing {
        pets?: (string|null);
        quantity?: (number|Long|null);
        details?: (string|null);
    }

    class ProvideHousing implements IProvideHousing {
        constructor(properties?: dss.IProvideHousing);
        public pets: string;
        public quantity: (number|Long);
        public details: string;
        public static create(properties?: dss.IProvideHousing): dss.ProvideHousing;
        public static encode(message: dss.IProvideHousing, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IProvideHousing, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.ProvideHousing;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.ProvideHousing;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.ProvideHousing;
        public static toObject(message: dss.ProvideHousing, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRequireHousing {
        petAllergies?: (string|null);
        details?: (string|null);
    }

    class RequireHousing implements IRequireHousing {
        constructor(properties?: dss.IRequireHousing);
        public petAllergies: string;
        public details: string;
        public static create(properties?: dss.IRequireHousing): dss.RequireHousing;
        public static encode(message: dss.IRequireHousing, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRequireHousing, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RequireHousing;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RequireHousing;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RequireHousing;
        public static toObject(message: dss.RequireHousing, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface INoHousing {
    }

    class NoHousing implements INoHousing {
        constructor(properties?: dss.INoHousing);
        public static create(properties?: dss.INoHousing): dss.NoHousing;
        public static encode(message: dss.INoHousing, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.INoHousing, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.NoHousing;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.NoHousing;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.NoHousing;
        public static toObject(message: dss.NoHousing, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationAddReq {
        idempotencyKey?: (string|null);
        registration?: (dss.IRegistrationInfo|null);
        redirectUrl?: (string|null);
    }

    class RegistrationAddReq implements IRegistrationAddReq {
        constructor(properties?: dss.IRegistrationAddReq);
        public idempotencyKey: string;
        public registration?: (dss.IRegistrationInfo|null);
        public redirectUrl: string;
        public static create(properties?: dss.IRegistrationAddReq): dss.RegistrationAddReq;
        public static encode(message: dss.IRegistrationAddReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationAddReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationAddReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationAddReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationAddReq;
        public static toObject(message: dss.RegistrationAddReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationAddRes {
        redirectUrl?: (string|null);
    }

    class RegistrationAddRes implements IRegistrationAddRes {
        constructor(properties?: dss.IRegistrationAddRes);
        public redirectUrl: string;
        public static create(properties?: dss.IRegistrationAddRes): dss.RegistrationAddRes;
        public static encode(message: dss.IRegistrationAddRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationAddRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationAddRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationAddRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationAddRes;
        public static toObject(message: dss.RegistrationAddRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationGetReq {
        id?: (string|null);
    }

    class RegistrationGetReq implements IRegistrationGetReq {
        constructor(properties?: dss.IRegistrationGetReq);
        public id: string;
        public static create(properties?: dss.IRegistrationGetReq): dss.RegistrationGetReq;
        public static encode(message: dss.IRegistrationGetReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationGetReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationGetReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationGetReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationGetReq;
        public static toObject(message: dss.RegistrationGetReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationGetRes {
        registration?: (dss.IRegistrationInfo|null);
    }

    class RegistrationGetRes implements IRegistrationGetRes {
        constructor(properties?: dss.IRegistrationGetRes);
        public registration?: (dss.IRegistrationInfo|null);
        public static create(properties?: dss.IRegistrationGetRes): dss.RegistrationGetRes;
        public static encode(message: dss.IRegistrationGetRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationGetRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationGetRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationGetRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationGetRes;
        public static toObject(message: dss.RegistrationGetRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationPricesReq {
    }

    class RegistrationPricesReq implements IRegistrationPricesReq {
        constructor(properties?: dss.IRegistrationPricesReq);
        public static create(properties?: dss.IRegistrationPricesReq): dss.RegistrationPricesReq;
        public static encode(message: dss.IRegistrationPricesReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationPricesReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationPricesReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationPricesReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationPricesReq;
        public static toObject(message: dss.RegistrationPricesReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationPricesRes {
        weekendPassCost?: (number|Long|null);
        weekendPassTier?: (dss.FullWeekendPassTier|null);
        dancePassCost?: (number|Long|null);
        mixAndMatchCost?: (number|Long|null);
        soloJazzCost?: (number|Long|null);
        teamCompetitionCost?: (number|Long|null);
        tshirtCost?: (number|Long|null);
        studentDiscount?: (dss.IDiscountAmount|null);
    }

    class RegistrationPricesRes implements IRegistrationPricesRes {
        constructor(properties?: dss.IRegistrationPricesRes);
        public weekendPassCost: (number|Long);
        public weekendPassTier: dss.FullWeekendPassTier;
        public dancePassCost: (number|Long);
        public mixAndMatchCost: (number|Long);
        public soloJazzCost: (number|Long);
        public teamCompetitionCost: (number|Long);
        public tshirtCost: (number|Long);
        public studentDiscount?: (dss.IDiscountAmount|null);
        public static create(properties?: dss.IRegistrationPricesRes): dss.RegistrationPricesRes;
        public static encode(message: dss.IRegistrationPricesRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationPricesRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationPricesRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationPricesRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationPricesRes;
        public static toObject(message: dss.RegistrationPricesRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationUpdateReq {
        idempotencyKey?: (string|null);
        registration?: (dss.IRegistrationInfo|null);
        redirectUrl?: (string|null);
    }

    class RegistrationUpdateReq implements IRegistrationUpdateReq {
        constructor(properties?: dss.IRegistrationUpdateReq);
        public idempotencyKey: string;
        public registration?: (dss.IRegistrationInfo|null);
        public redirectUrl: string;
        public static create(properties?: dss.IRegistrationUpdateReq): dss.RegistrationUpdateReq;
        public static encode(message: dss.IRegistrationUpdateReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationUpdateReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationUpdateReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationUpdateReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationUpdateReq;
        public static toObject(message: dss.RegistrationUpdateReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationUpdateRes {
        redirectUrl?: (string|null);
    }

    class RegistrationUpdateRes implements IRegistrationUpdateRes {
        constructor(properties?: dss.IRegistrationUpdateRes);
        public redirectUrl: string;
        public static create(properties?: dss.IRegistrationUpdateRes): dss.RegistrationUpdateRes;
        public static encode(message: dss.IRegistrationUpdateRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationUpdateRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationUpdateRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationUpdateRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationUpdateRes;
        public static toObject(message: dss.RegistrationUpdateRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationGetSummaryReq {
    }

    class RegistrationGetSummaryReq implements IRegistrationGetSummaryReq {
        constructor(properties?: dss.IRegistrationGetSummaryReq);
        public static create(properties?: dss.IRegistrationGetSummaryReq): dss.RegistrationGetSummaryReq;
        public static encode(message: dss.IRegistrationGetSummaryReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationGetSummaryReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationGetSummaryReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationGetSummaryReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationGetSummaryReq;
        public static toObject(message: dss.RegistrationGetSummaryReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationSummary {
        id?: (string|null);
        firstName?: (string|null);
        lastName?: (string|null);
        email?: (string|null);
        createdAt?: (string|null);
        paid?: (boolean|null);
    }

    class RegistrationSummary implements IRegistrationSummary {
        constructor(properties?: dss.IRegistrationSummary);
        public id: string;
        public firstName: string;
        public lastName: string;
        public email: string;
        public createdAt: string;
        public paid: boolean;
        public static create(properties?: dss.IRegistrationSummary): dss.RegistrationSummary;
        public static encode(message: dss.IRegistrationSummary, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationSummary, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationSummary;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationSummary;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationSummary;
        public static toObject(message: dss.RegistrationSummary, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationGetSummaryRes {
        summaries?: (dss.IRegistrationSummary[]|null);
    }

    class RegistrationGetSummaryRes implements IRegistrationGetSummaryRes {
        constructor(properties?: dss.IRegistrationGetSummaryRes);
        public summaries: dss.IRegistrationSummary[];
        public static create(properties?: dss.IRegistrationGetSummaryRes): dss.RegistrationGetSummaryRes;
        public static encode(message: dss.IRegistrationGetSummaryRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationGetSummaryRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationGetSummaryRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationGetSummaryRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationGetSummaryRes;
        public static toObject(message: dss.RegistrationGetSummaryRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    class Discount extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): Discount;
        public add(request: dss.IDiscountAddReq, callback: dss.Discount.AddCallback): void;
        public add(request: dss.IDiscountAddReq): Promise<dss.DiscountAddRes>;
        public get(request: dss.IDiscountGetReq, callback: dss.Discount.GetCallback): void;
        public get(request: dss.IDiscountGetReq): Promise<dss.DiscountGetRes>;
        public list(request: dss.IDiscountListReq, callback: dss.Discount.ListCallback): void;
        public list(request: dss.IDiscountListReq): Promise<dss.DiscountListRes>;
        public update(request: dss.IDiscountUpdateReq, callback: dss.Discount.UpdateCallback): void;
        public update(request: dss.IDiscountUpdateReq): Promise<dss.DiscountUpdateRes>;
        public delete(request: dss.IDiscountDeleteReq, callback: dss.Discount.DeleteCallback): void;
        public delete(request: dss.IDiscountDeleteReq): Promise<dss.DiscountDeleteRes>;
    }

    namespace Discount {

        type AddCallback = (error: (Error|null), response?: dss.DiscountAddRes) => void;

        type GetCallback = (error: (Error|null), response?: dss.DiscountGetRes) => void;

        type ListCallback = (error: (Error|null), response?: dss.DiscountListRes) => void;

        type UpdateCallback = (error: (Error|null), response?: dss.DiscountUpdateRes) => void;

        type DeleteCallback = (error: (Error|null), response?: dss.DiscountDeleteRes) => void;
    }

    enum PurchaseItem {
        FullWeekendPassPurchaseItem = 0,
        DanceOnlyPassPurchaseItem = 1,
        MixAndMatchPurchaseItem = 2,
        SoloJazzPurchaseItem = 3,
        TeamCompetitionPurchaseItem = 4,
        TShirtPurchaseItem = 5
    }

    interface IDiscountAmount {
        dollar?: (number|Long|null);
        percent?: (string|null);
        squareNotFound?: (google.protobuf.IEmpty|null);
    }

    class DiscountAmount implements IDiscountAmount {
        constructor(properties?: dss.IDiscountAmount);
        public dollar: (number|Long);
        public percent: string;
        public squareNotFound?: (google.protobuf.IEmpty|null);
        public amount?: ("dollar"|"percent"|"squareNotFound");
        public static create(properties?: dss.IDiscountAmount): dss.DiscountAmount;
        public static encode(message: dss.IDiscountAmount, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountAmount, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountAmount;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountAmount;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountAmount;
        public static toObject(message: dss.DiscountAmount, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface ISingleDiscount {
        name?: (string|null);
        amount?: (dss.IDiscountAmount|null);
        appliedTo?: (dss.PurchaseItem|null);
    }

    class SingleDiscount implements ISingleDiscount {
        constructor(properties?: dss.ISingleDiscount);
        public name: string;
        public amount?: (dss.IDiscountAmount|null);
        public appliedTo: dss.PurchaseItem;
        public static create(properties?: dss.ISingleDiscount): dss.SingleDiscount;
        public static encode(message: dss.ISingleDiscount, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.ISingleDiscount, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.SingleDiscount;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.SingleDiscount;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.SingleDiscount;
        public static toObject(message: dss.SingleDiscount, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountBundle {
        code?: (string|null);
        discounts?: (dss.ISingleDiscount[]|null);
    }

    class DiscountBundle implements IDiscountBundle {
        constructor(properties?: dss.IDiscountBundle);
        public code: string;
        public discounts: dss.ISingleDiscount[];
        public static create(properties?: dss.IDiscountBundle): dss.DiscountBundle;
        public static encode(message: dss.IDiscountBundle, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountBundle, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountBundle;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountBundle;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountBundle;
        public static toObject(message: dss.DiscountBundle, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountAddReq {
        bundle?: (dss.IDiscountBundle|null);
    }

    class DiscountAddReq implements IDiscountAddReq {
        constructor(properties?: dss.IDiscountAddReq);
        public bundle?: (dss.IDiscountBundle|null);
        public static create(properties?: dss.IDiscountAddReq): dss.DiscountAddReq;
        public static encode(message: dss.IDiscountAddReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountAddReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountAddReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountAddReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountAddReq;
        public static toObject(message: dss.DiscountAddReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountAddRes {
    }

    class DiscountAddRes implements IDiscountAddRes {
        constructor(properties?: dss.IDiscountAddRes);
        public static create(properties?: dss.IDiscountAddRes): dss.DiscountAddRes;
        public static encode(message: dss.IDiscountAddRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountAddRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountAddRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountAddRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountAddRes;
        public static toObject(message: dss.DiscountAddRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountGetReq {
        code?: (string|null);
    }

    class DiscountGetReq implements IDiscountGetReq {
        constructor(properties?: dss.IDiscountGetReq);
        public code: string;
        public static create(properties?: dss.IDiscountGetReq): dss.DiscountGetReq;
        public static encode(message: dss.IDiscountGetReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountGetReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountGetReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountGetReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountGetReq;
        public static toObject(message: dss.DiscountGetReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountGetRes {
        bundle?: (dss.IDiscountBundle|null);
    }

    class DiscountGetRes implements IDiscountGetRes {
        constructor(properties?: dss.IDiscountGetRes);
        public bundle?: (dss.IDiscountBundle|null);
        public static create(properties?: dss.IDiscountGetRes): dss.DiscountGetRes;
        public static encode(message: dss.IDiscountGetRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountGetRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountGetRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountGetRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountGetRes;
        public static toObject(message: dss.DiscountGetRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountListReq {
    }

    class DiscountListReq implements IDiscountListReq {
        constructor(properties?: dss.IDiscountListReq);
        public static create(properties?: dss.IDiscountListReq): dss.DiscountListReq;
        public static encode(message: dss.IDiscountListReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountListReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountListReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountListReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountListReq;
        public static toObject(message: dss.DiscountListReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountListRes {
        bundles?: (dss.IDiscountBundle[]|null);
    }

    class DiscountListRes implements IDiscountListRes {
        constructor(properties?: dss.IDiscountListRes);
        public bundles: dss.IDiscountBundle[];
        public static create(properties?: dss.IDiscountListRes): dss.DiscountListRes;
        public static encode(message: dss.IDiscountListRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountListRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountListRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountListRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountListRes;
        public static toObject(message: dss.DiscountListRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountUpdateReq {
        oldCode?: (string|null);
        bundle?: (dss.IDiscountBundle|null);
    }

    class DiscountUpdateReq implements IDiscountUpdateReq {
        constructor(properties?: dss.IDiscountUpdateReq);
        public oldCode: string;
        public bundle?: (dss.IDiscountBundle|null);
        public static create(properties?: dss.IDiscountUpdateReq): dss.DiscountUpdateReq;
        public static encode(message: dss.IDiscountUpdateReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountUpdateReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountUpdateReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountUpdateReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountUpdateReq;
        public static toObject(message: dss.DiscountUpdateReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountUpdateRes {
    }

    class DiscountUpdateRes implements IDiscountUpdateRes {
        constructor(properties?: dss.IDiscountUpdateRes);
        public static create(properties?: dss.IDiscountUpdateRes): dss.DiscountUpdateRes;
        public static encode(message: dss.IDiscountUpdateRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountUpdateRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountUpdateRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountUpdateRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountUpdateRes;
        public static toObject(message: dss.DiscountUpdateRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountDeleteReq {
        code?: (string|null);
    }

    class DiscountDeleteReq implements IDiscountDeleteReq {
        constructor(properties?: dss.IDiscountDeleteReq);
        public code: string;
        public static create(properties?: dss.IDiscountDeleteReq): dss.DiscountDeleteReq;
        public static encode(message: dss.IDiscountDeleteReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountDeleteReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountDeleteReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountDeleteReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountDeleteReq;
        public static toObject(message: dss.DiscountDeleteReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IDiscountDeleteRes {
    }

    class DiscountDeleteRes implements IDiscountDeleteRes {
        constructor(properties?: dss.IDiscountDeleteRes);
        public static create(properties?: dss.IDiscountDeleteRes): dss.DiscountDeleteRes;
        public static encode(message: dss.IDiscountDeleteRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IDiscountDeleteRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.DiscountDeleteRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.DiscountDeleteRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.DiscountDeleteRes;
        public static toObject(message: dss.DiscountDeleteRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }
}

export namespace google {

    namespace protobuf {

        interface IEmpty {
        }

        class Empty implements IEmpty {
            constructor(properties?: google.protobuf.IEmpty);
            public static create(properties?: google.protobuf.IEmpty): google.protobuf.Empty;
            public static encode(message: google.protobuf.IEmpty, writer?: $protobuf.Writer): $protobuf.Writer;
            public static encodeDelimited(message: google.protobuf.IEmpty, writer?: $protobuf.Writer): $protobuf.Writer;
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.Empty;
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): google.protobuf.Empty;
            public static verify(message: { [k: string]: any }): (string|null);
            public static fromObject(object: { [k: string]: any }): google.protobuf.Empty;
            public static toObject(message: google.protobuf.Empty, options?: $protobuf.IConversionOptions): { [k: string]: any };
            public toJSON(): { [k: string]: any };
        }
    }
}
