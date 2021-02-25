import * as $protobuf from "protobufjs";
export namespace dss {

    class Forms extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): Forms;
        public contactUs(request: dss.IContactUsReq, callback: dss.Forms.ContactUsCallback): void;
        public contactUs(request: dss.IContactUsReq): Promise<dss.ContactUsRes>;
        public safetyReport(request: dss.ISafetyReportReq, callback: dss.Forms.SafetyReportCallback): void;
        public safetyReport(request: dss.ISafetyReportReq): Promise<dss.SafetyReportRes>;
    }

    namespace Forms {

        type ContactUsCallback = (error: (Error|null), response?: dss.ContactUsRes) => void;

        type SafetyReportCallback = (error: (Error|null), response?: dss.SafetyReportRes) => void;
    }

    interface IContactUsReq {
        name?: (string|null);
        email?: (string|null);
        msg?: (string|null);
        recaptchaResponse?: (string|null);
    }

    class ContactUsReq implements IContactUsReq {
        constructor(properties?: dss.IContactUsReq);
        public name: string;
        public email: string;
        public msg: string;
        public recaptchaResponse: string;
        public static create(properties?: dss.IContactUsReq): dss.ContactUsReq;
        public static encode(message: dss.IContactUsReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IContactUsReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.ContactUsReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.ContactUsReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.ContactUsReq;
        public static toObject(message: dss.ContactUsReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IContactUsRes {
    }

    class ContactUsRes implements IContactUsRes {
        constructor(properties?: dss.IContactUsRes);
        public static create(properties?: dss.IContactUsRes): dss.ContactUsRes;
        public static encode(message: dss.IContactUsRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IContactUsRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.ContactUsRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.ContactUsRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.ContactUsRes;
        public static toObject(message: dss.ContactUsRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface ISafetyReportReq {
        occurredOn?: (google.protobuf.ITimestamp|null);
        description?: (string|null);
        severity?: (number|null);
        issuesBefore?: (boolean|null);
        resolution?: (string|null);
        name?: (string|null);
        email?: (string|null);
        phoneNumber?: (string|null);
        recaptchaResponse?: (string|null);
    }

    class SafetyReportReq implements ISafetyReportReq {
        constructor(properties?: dss.ISafetyReportReq);
        public occurredOn?: (google.protobuf.ITimestamp|null);
        public description: string;
        public severity: number;
        public issuesBefore: boolean;
        public resolution: string;
        public name: string;
        public email: string;
        public phoneNumber: string;
        public recaptchaResponse: string;
        public static create(properties?: dss.ISafetyReportReq): dss.SafetyReportReq;
        public static encode(message: dss.ISafetyReportReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.ISafetyReportReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.SafetyReportReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.SafetyReportReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.SafetyReportReq;
        public static toObject(message: dss.SafetyReportReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface ISafetyReportRes {
    }

    class SafetyReportRes implements ISafetyReportRes {
        constructor(properties?: dss.ISafetyReportRes);
        public static create(properties?: dss.ISafetyReportRes): dss.SafetyReportRes;
        public static encode(message: dss.ISafetyReportRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.ISafetyReportRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.SafetyReportRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.SafetyReportRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.SafetyReportRes;
        public static toObject(message: dss.SafetyReportRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }
}

export namespace google {

    namespace protobuf {

        interface ITimestamp {
            seconds?: (number|Long|null);
            nanos?: (number|null);
        }

        class Timestamp implements ITimestamp {
            constructor(properties?: google.protobuf.ITimestamp);
            public seconds: (number|Long);
            public nanos: number;
            public static create(properties?: google.protobuf.ITimestamp): google.protobuf.Timestamp;
            public static encode(message: google.protobuf.ITimestamp, writer?: $protobuf.Writer): $protobuf.Writer;
            public static encodeDelimited(message: google.protobuf.ITimestamp, writer?: $protobuf.Writer): $protobuf.Writer;
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.Timestamp;
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): google.protobuf.Timestamp;
            public static verify(message: { [k: string]: any }): (string|null);
            public static fromObject(object: { [k: string]: any }): google.protobuf.Timestamp;
            public static toObject(message: google.protobuf.Timestamp, options?: $protobuf.IConversionOptions): { [k: string]: any };
            public toJSON(): { [k: string]: any };
        }
    }
}
