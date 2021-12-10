import * as $protobuf from "protobufjs";
export namespace dss {

    class Registration extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): Registration;
        public add(request: dss.IRegistrationAddReq, callback: dss.Registration.AddCallback): void;
        public add(request: dss.IRegistrationAddReq): Promise<dss.RegistrationAddRes>;
        public pay(request: dss.IRegistrationPayReq, callback: dss.Registration.PayCallback): void;
        public pay(request: dss.IRegistrationPayReq): Promise<dss.RegistrationPayRes>;
        public get(request: dss.IRegistrationGetReq, callback: dss.Registration.GetCallback): void;
        public get(request: dss.IRegistrationGetReq): Promise<dss.RegistrationGetRes>;
        public listByUser(request: dss.IRegistrationListByUserReq, callback: dss.Registration.ListByUserCallback): void;
        public listByUser(request: dss.IRegistrationListByUserReq): Promise<dss.RegistrationListByUserRes>;
        public list(request: dss.IRegistrationListReq, callback: dss.Registration.ListCallback): void;
        public list(request: dss.IRegistrationListReq): Promise<dss.RegistrationListRes>;
        public prices(request: dss.IRegistrationPricesReq, callback: dss.Registration.PricesCallback): void;
        public prices(request: dss.IRegistrationPricesReq): Promise<dss.RegistrationPricesRes>;
        public update(request: dss.IRegistrationUpdateReq, callback: dss.Registration.UpdateCallback): void;
        public update(request: dss.IRegistrationUpdateReq): Promise<dss.RegistrationUpdateRes>;
    }

    namespace Registration {

        type AddCallback = (error: (Error|null), response?: dss.RegistrationAddRes) => void;

        type PayCallback = (error: (Error|null), response?: dss.RegistrationPayRes) => void;

        type GetCallback = (error: (Error|null), response?: dss.RegistrationGetRes) => void;

        type ListByUserCallback = (error: (Error|null), response?: dss.RegistrationListByUserRes) => void;

        type ListCallback = (error: (Error|null), response?: dss.RegistrationListRes) => void;

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
        enabled?: (boolean|null);
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
        public enabled: boolean;
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
        squarePaid?: (boolean|null);
        adminPaymentOverride?: (boolean|null);
    }

    class FullWeekendPass implements IFullWeekendPass {
        constructor(properties?: dss.IFullWeekendPass);
        public tier: dss.FullWeekendPassTier;
        public level: dss.FullWeekendPassLevel;
        public squarePaid: boolean;
        public adminPaymentOverride: boolean;
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
        squarePaid?: (boolean|null);
        adminPaymentOverride?: (boolean|null);
    }

    class DanceOnlyPass implements IDanceOnlyPass {
        constructor(properties?: dss.IDanceOnlyPass);
        public squarePaid: boolean;
        public adminPaymentOverride: boolean;
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
        squarePaid?: (boolean|null);
        adminPaymentOverride?: (boolean|null);
    }

    class MixAndMatch implements IMixAndMatch {
        constructor(properties?: dss.IMixAndMatch);
        public role: dss.MixAndMatch.Role;
        public squarePaid: boolean;
        public adminPaymentOverride: boolean;
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
        squarePaid?: (boolean|null);
        adminPaymentOverride?: (boolean|null);
    }

    class SoloJazz implements ISoloJazz {
        constructor(properties?: dss.ISoloJazz);
        public squarePaid: boolean;
        public adminPaymentOverride: boolean;
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
        squarePaid?: (boolean|null);
        adminPaymentOverride?: (boolean|null);
    }

    class TeamCompetition implements ITeamCompetition {
        constructor(properties?: dss.ITeamCompetition);
        public name: string;
        public squarePaid: boolean;
        public adminPaymentOverride: boolean;
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
        squarePaid?: (boolean|null);
        adminPaymentOverride?: (boolean|null);
    }

    class TShirt implements ITShirt {
        constructor(properties?: dss.ITShirt);
        public style: dss.TShirt.Style;
        public squarePaid: boolean;
        public adminPaymentOverride: boolean;
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
        registration?: (dss.IRegistrationInfo|null);
    }

    class RegistrationAddReq implements IRegistrationAddReq {
        constructor(properties?: dss.IRegistrationAddReq);
        public registration?: (dss.IRegistrationInfo|null);
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
        registration?: (dss.IRegistrationInfo|null);
    }

    class RegistrationAddRes implements IRegistrationAddRes {
        constructor(properties?: dss.IRegistrationAddRes);
        public registration?: (dss.IRegistrationInfo|null);
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

    interface IRegistrationPayReq {
        id?: (string|null);
        idempotencyKey?: (string|null);
        redirectUrl?: (string|null);
    }

    class RegistrationPayReq implements IRegistrationPayReq {
        constructor(properties?: dss.IRegistrationPayReq);
        public id: string;
        public idempotencyKey: string;
        public redirectUrl: string;
        public static create(properties?: dss.IRegistrationPayReq): dss.RegistrationPayReq;
        public static encode(message: dss.IRegistrationPayReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationPayReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationPayReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationPayReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationPayReq;
        public static toObject(message: dss.RegistrationPayReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationPayRes {
        checkoutUrl?: (string|null);
    }

    class RegistrationPayRes implements IRegistrationPayRes {
        constructor(properties?: dss.IRegistrationPayRes);
        public checkoutUrl: string;
        public static create(properties?: dss.IRegistrationPayRes): dss.RegistrationPayRes;
        public static encode(message: dss.IRegistrationPayRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationPayRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationPayRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationPayRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationPayRes;
        public static toObject(message: dss.RegistrationPayRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
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
        weekendPassTier?: (dss.FullWeekendPassTier|null);
    }

    class RegistrationPricesRes implements IRegistrationPricesRes {
        constructor(properties?: dss.IRegistrationPricesRes);
        public weekendPassTier: dss.FullWeekendPassTier;
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
        registration?: (dss.IRegistrationInfo|null);
    }

    class RegistrationUpdateReq implements IRegistrationUpdateReq {
        constructor(properties?: dss.IRegistrationUpdateReq);
        public registration?: (dss.IRegistrationInfo|null);
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
        registration?: (dss.IRegistrationInfo|null);
    }

    class RegistrationUpdateRes implements IRegistrationUpdateRes {
        constructor(properties?: dss.IRegistrationUpdateRes);
        public registration?: (dss.IRegistrationInfo|null);
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

    interface IRegistrationListByUserReq {
    }

    class RegistrationListByUserReq implements IRegistrationListByUserReq {
        constructor(properties?: dss.IRegistrationListByUserReq);
        public static create(properties?: dss.IRegistrationListByUserReq): dss.RegistrationListByUserReq;
        public static encode(message: dss.IRegistrationListByUserReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationListByUserReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationListByUserReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationListByUserReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationListByUserReq;
        public static toObject(message: dss.RegistrationListByUserReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationListByUserRes {
        registrations?: (dss.IRegistrationInfo[]|null);
    }

    class RegistrationListByUserRes implements IRegistrationListByUserRes {
        constructor(properties?: dss.IRegistrationListByUserRes);
        public registrations: dss.IRegistrationInfo[];
        public static create(properties?: dss.IRegistrationListByUserRes): dss.RegistrationListByUserRes;
        public static encode(message: dss.IRegistrationListByUserRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationListByUserRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationListByUserRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationListByUserRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationListByUserRes;
        public static toObject(message: dss.RegistrationListByUserRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationListReq {
    }

    class RegistrationListReq implements IRegistrationListReq {
        constructor(properties?: dss.IRegistrationListReq);
        public static create(properties?: dss.IRegistrationListReq): dss.RegistrationListReq;
        public static encode(message: dss.IRegistrationListReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationListReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationListReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationListReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationListReq;
        public static toObject(message: dss.RegistrationListReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IRegistrationListRes {
        registrations?: (dss.IRegistrationInfo[]|null);
    }

    class RegistrationListRes implements IRegistrationListRes {
        constructor(properties?: dss.IRegistrationListRes);
        public registrations: dss.IRegistrationInfo[];
        public static create(properties?: dss.IRegistrationListRes): dss.RegistrationListRes;
        public static encode(message: dss.IRegistrationListRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IRegistrationListRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.RegistrationListRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.RegistrationListRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.RegistrationListRes;
        public static toObject(message: dss.RegistrationListRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }
}
