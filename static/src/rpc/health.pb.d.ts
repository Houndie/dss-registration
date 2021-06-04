import * as $protobuf from "protobufjs";
export namespace dss {

    class Info extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): Info;
        public health(request: dss.IInfoHealthReq, callback: dss.Info.HealthCallback): void;
        public health(request: dss.IInfoHealthReq): Promise<dss.InfoHealthRes>;
        public version(request: dss.IInfoVersionReq, callback: dss.Info.VersionCallback): void;
        public version(request: dss.IInfoVersionReq): Promise<dss.InfoVersionRes>;
    }

    namespace Info {

        type HealthCallback = (error: (Error|null), response?: dss.InfoHealthRes) => void;

        type VersionCallback = (error: (Error|null), response?: dss.InfoVersionRes) => void;
    }

    interface IInfoHealthReq {
    }

    class InfoHealthReq implements IInfoHealthReq {
        constructor(properties?: dss.IInfoHealthReq);
        public static create(properties?: dss.IInfoHealthReq): dss.InfoHealthReq;
        public static encode(message: dss.IInfoHealthReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IInfoHealthReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.InfoHealthReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.InfoHealthReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.InfoHealthReq;
        public static toObject(message: dss.InfoHealthReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IInfoHealthRes {
        healthiness?: (dss.InfoHealthRes.Healthiness|null);
    }

    class InfoHealthRes implements IInfoHealthRes {
        constructor(properties?: dss.IInfoHealthRes);
        public healthiness: dss.InfoHealthRes.Healthiness;
        public static create(properties?: dss.IInfoHealthRes): dss.InfoHealthRes;
        public static encode(message: dss.IInfoHealthRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IInfoHealthRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.InfoHealthRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.InfoHealthRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.InfoHealthRes;
        public static toObject(message: dss.InfoHealthRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    namespace InfoHealthRes {

        enum Healthiness {
            Unknown = 0,
            Healthy = 1,
            Unhealthy = 2
        }
    }

    interface IInfoVersionReq {
    }

    class InfoVersionReq implements IInfoVersionReq {
        constructor(properties?: dss.IInfoVersionReq);
        public static create(properties?: dss.IInfoVersionReq): dss.InfoVersionReq;
        public static encode(message: dss.IInfoVersionReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IInfoVersionReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.InfoVersionReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.InfoVersionReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.InfoVersionReq;
        public static toObject(message: dss.InfoVersionReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IInfoVersionRes {
        version?: (string|null);
    }

    class InfoVersionRes implements IInfoVersionRes {
        constructor(properties?: dss.IInfoVersionRes);
        public version: string;
        public static create(properties?: dss.IInfoVersionRes): dss.InfoVersionRes;
        public static encode(message: dss.IInfoVersionRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IInfoVersionRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.InfoVersionRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.InfoVersionRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.InfoVersionRes;
        public static toObject(message: dss.InfoVersionRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }
}
