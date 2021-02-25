// Generated by protoc-gen-twirp_typescript. DO NOT EDIT
import {dss} from './forms.pb';
import {createTwirpAdapter} from 'pbjs-twirp';
import Axios from 'axios';

const getServiceMethodName = (fn: any): string => {
	if (fn === dss.Forms.prototype.contactUs) {
		return 'ContactUs';
    }
	if (fn === dss.Forms.prototype.safetyReport) {
		return 'SafetyReport';
    }

    throw new Error('Unknown Method');
};


export const FormsPathPrefix = '/twirp/dss.Forms/';

export const createForms = (baseURL: string, options = {}): dss.Forms => {
    const defaultOpts = {
        baseURL: baseURL + FormsPathPrefix,
        headers: {
          Accept: 'application/protobuf'
        }
    };
    const axiosOpts = { ...defaultOpts, ...options };
    
    const axios = Axios.create(axiosOpts);

    return dss.Forms.create(createTwirpAdapter(axios, getServiceMethodName));
};
