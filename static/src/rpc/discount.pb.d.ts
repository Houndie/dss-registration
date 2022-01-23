import * as $protobuf from "protobufjs";
export namespace dss {

    class Discount extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): Discount;
        public get(request: dss.IDiscountGetReq, callback: dss.Discount.GetCallback): void;
        public get(request: dss.IDiscountGetReq): Promise<dss.DiscountGetRes>;
    }

    namespace Discount {

        type GetCallback = (error: (Error|null), response?: dss.DiscountGetRes) => void;
    }

    enum PurchaseItem {
        Unknown = 0,
        FullWeekendPassPurchaseItem = 1,
        DanceOnlyPassPurchaseItem = 2,
        MixAndMatchPurchaseItem = 3,
        SoloJazzPurchaseItem = 4,
        TeamCompetitionPurchaseItem = 5,
        TShirtPurchaseItem = 6
    }

    interface IDiscountAmount {
        dollar?: (number|Long|null);
        percent?: (string|null);
    }

    class DiscountAmount implements IDiscountAmount {
        constructor(properties?: dss.IDiscountAmount);
        public dollar: (number|Long);
        public percent: string;
        public amount?: ("dollar"|"percent");
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
        amount?: (dss.IDiscountAmount|null);
        appliedTo?: (dss.PurchaseItem|null);
    }

    class SingleDiscount implements ISingleDiscount {
        constructor(properties?: dss.ISingleDiscount);
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
}
