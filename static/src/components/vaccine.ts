import {dss} from "../rpc/vaccine.pb"

export enum VaccineInfoEnum {
	VaxApprovedEnum = "vax_approved",
	VaxApprovalPendingEnum = "vax_approval_pending",
	NoVaxProofSuppliedEnum = "no_vax_proof_supplied"
}

type VaxApproved = {
	type: VaccineInfoEnum.VaxApprovedEnum
}

type VaxApprovalPending = {
	type: VaccineInfoEnum.VaxApprovalPendingEnum,
	URL: string
}

type NoVaxProofSupplied = {
	type: VaccineInfoEnum.NoVaxProofSuppliedEnum,
}

export type VaccineInfo = VaxApproved | VaxApprovalPending | NoVaxProofSupplied

export const fromProtoVaccine = (res: dss.IVaccineGetRes): VaccineInfo => {
	if(res.vaxApproved){
		return {
			type: VaccineInfoEnum.VaxApprovedEnum
		}
	}

	if(res.vaxApprovalPending){
		return {
			type: VaccineInfoEnum.VaxApprovalPendingEnum,	
			URL: (res.vaxApprovalPending.url ? res.vaxApprovalPending.url : "")
		}
	}

	if(res.noVaxProofSupplied){
		return {
			type: VaccineInfoEnum.NoVaxProofSuppliedEnum
		}
	}

	throw "no vaccine result?"
}
