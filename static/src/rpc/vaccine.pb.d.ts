import * as $protobuf from "protobufjs";
export namespace dss {

    class Vaccine extends $protobuf.rpc.Service {
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): Vaccine;
        public upload(request: dss.IVaccineUploadReq, callback: dss.Vaccine.UploadCallback): void;
        public upload(request: dss.IVaccineUploadReq): Promise<dss.VaccineUploadRes>;
        public get(request: dss.IVaccineGetReq, callback: dss.Vaccine.GetCallback): void;
        public get(request: dss.IVaccineGetReq): Promise<dss.VaccineGetRes>;
        public approve(request: dss.IVaccineApproveReq, callback: dss.Vaccine.ApproveCallback): void;
        public approve(request: dss.IVaccineApproveReq): Promise<dss.VaccineApproveRes>;
        public reject(request: dss.IVaccineRejectReq, callback: dss.Vaccine.RejectCallback): void;
        public reject(request: dss.IVaccineRejectReq): Promise<dss.VaccineRejectRes>;
    }

    namespace Vaccine {

        type UploadCallback = (error: (Error|null), response?: dss.VaccineUploadRes) => void;

        type GetCallback = (error: (Error|null), response?: dss.VaccineGetRes) => void;

        type ApproveCallback = (error: (Error|null), response?: dss.VaccineApproveRes) => void;

        type RejectCallback = (error: (Error|null), response?: dss.VaccineRejectRes) => void;
    }

    interface IVaxApproved {
    }

    class VaxApproved implements IVaxApproved {
        constructor(properties?: dss.IVaxApproved);
        public static create(properties?: dss.IVaxApproved): dss.VaxApproved;
        public static encode(message: dss.IVaxApproved, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaxApproved, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaxApproved;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaxApproved;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaxApproved;
        public static toObject(message: dss.VaxApproved, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaxApprovalPending {
        url?: (string|null);
    }

    class VaxApprovalPending implements IVaxApprovalPending {
        constructor(properties?: dss.IVaxApprovalPending);
        public url: string;
        public static create(properties?: dss.IVaxApprovalPending): dss.VaxApprovalPending;
        public static encode(message: dss.IVaxApprovalPending, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaxApprovalPending, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaxApprovalPending;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaxApprovalPending;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaxApprovalPending;
        public static toObject(message: dss.VaxApprovalPending, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface INoVaxProofSupplied {
    }

    class NoVaxProofSupplied implements INoVaxProofSupplied {
        constructor(properties?: dss.INoVaxProofSupplied);
        public static create(properties?: dss.INoVaxProofSupplied): dss.NoVaxProofSupplied;
        public static encode(message: dss.INoVaxProofSupplied, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.INoVaxProofSupplied, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.NoVaxProofSupplied;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.NoVaxProofSupplied;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.NoVaxProofSupplied;
        public static toObject(message: dss.NoVaxProofSupplied, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineUploadReq {
        id?: (string|null);
        filesize?: (number|Long|null);
    }

    class VaccineUploadReq implements IVaccineUploadReq {
        constructor(properties?: dss.IVaccineUploadReq);
        public id: string;
        public filesize: (number|Long);
        public static create(properties?: dss.IVaccineUploadReq): dss.VaccineUploadReq;
        public static encode(message: dss.IVaccineUploadReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineUploadReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineUploadReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineUploadReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineUploadReq;
        public static toObject(message: dss.VaccineUploadReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineUploadRes {
        url?: (string|null);
    }

    class VaccineUploadRes implements IVaccineUploadRes {
        constructor(properties?: dss.IVaccineUploadRes);
        public url: string;
        public static create(properties?: dss.IVaccineUploadRes): dss.VaccineUploadRes;
        public static encode(message: dss.IVaccineUploadRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineUploadRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineUploadRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineUploadRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineUploadRes;
        public static toObject(message: dss.VaccineUploadRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineGetReq {
        id?: (string|null);
    }

    class VaccineGetReq implements IVaccineGetReq {
        constructor(properties?: dss.IVaccineGetReq);
        public id: string;
        public static create(properties?: dss.IVaccineGetReq): dss.VaccineGetReq;
        public static encode(message: dss.IVaccineGetReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineGetReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineGetReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineGetReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineGetReq;
        public static toObject(message: dss.VaccineGetReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineGetRes {
        vaxApproved?: (dss.IVaxApproved|null);
        vaxApprovalPending?: (dss.IVaxApprovalPending|null);
        noVaxProofSupplied?: (dss.INoVaxProofSupplied|null);
    }

    class VaccineGetRes implements IVaccineGetRes {
        constructor(properties?: dss.IVaccineGetRes);
        public vaxApproved?: (dss.IVaxApproved|null);
        public vaxApprovalPending?: (dss.IVaxApprovalPending|null);
        public noVaxProofSupplied?: (dss.INoVaxProofSupplied|null);
        public info?: ("vaxApproved"|"vaxApprovalPending"|"noVaxProofSupplied");
        public static create(properties?: dss.IVaccineGetRes): dss.VaccineGetRes;
        public static encode(message: dss.IVaccineGetRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineGetRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineGetRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineGetRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineGetRes;
        public static toObject(message: dss.VaccineGetRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineApproveReq {
        id?: (string|null);
    }

    class VaccineApproveReq implements IVaccineApproveReq {
        constructor(properties?: dss.IVaccineApproveReq);
        public id: string;
        public static create(properties?: dss.IVaccineApproveReq): dss.VaccineApproveReq;
        public static encode(message: dss.IVaccineApproveReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineApproveReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineApproveReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineApproveReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineApproveReq;
        public static toObject(message: dss.VaccineApproveReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineApproveRes {
    }

    class VaccineApproveRes implements IVaccineApproveRes {
        constructor(properties?: dss.IVaccineApproveRes);
        public static create(properties?: dss.IVaccineApproveRes): dss.VaccineApproveRes;
        public static encode(message: dss.IVaccineApproveRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineApproveRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineApproveRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineApproveRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineApproveRes;
        public static toObject(message: dss.VaccineApproveRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineRejectReq {
        id?: (string|null);
        reason?: (string|null);
    }

    class VaccineRejectReq implements IVaccineRejectReq {
        constructor(properties?: dss.IVaccineRejectReq);
        public id: string;
        public reason: string;
        public static create(properties?: dss.IVaccineRejectReq): dss.VaccineRejectReq;
        public static encode(message: dss.IVaccineRejectReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineRejectReq, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineRejectReq;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineRejectReq;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineRejectReq;
        public static toObject(message: dss.VaccineRejectReq, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }

    interface IVaccineRejectRes {
    }

    class VaccineRejectRes implements IVaccineRejectRes {
        constructor(properties?: dss.IVaccineRejectRes);
        public static create(properties?: dss.IVaccineRejectRes): dss.VaccineRejectRes;
        public static encode(message: dss.IVaccineRejectRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static encodeDelimited(message: dss.IVaccineRejectRes, writer?: $protobuf.Writer): $protobuf.Writer;
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): dss.VaccineRejectRes;
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): dss.VaccineRejectRes;
        public static verify(message: { [k: string]: any }): (string|null);
        public static fromObject(object: { [k: string]: any }): dss.VaccineRejectRes;
        public static toObject(message: dss.VaccineRejectRes, options?: $protobuf.IConversionOptions): { [k: string]: any };
        public toJSON(): { [k: string]: any };
    }
}
