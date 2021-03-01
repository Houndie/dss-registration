import * as $protobuf from "protobufjs";
export namespace dss {

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
    }

    namespace Discount {

        type AddCallback = (error: (Error|null), response?: dss.DiscountAddRes) => void;

        type GetCallback = (error: (Error|null), response?: dss.DiscountGetRes) => void;

        type ListCallback = (error: (Error|null), response?: dss.DiscountListRes) => void;

        type UpdateCallback = (error: (Error|null), response?: dss.DiscountUpdateRes) => void;
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
