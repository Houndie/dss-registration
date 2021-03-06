// Generated by protoc-gen-twirp_typescript. DO NOT EDIT
import {dss} from './registration.pb';
import {createTwirpAdapter} from 'pbjs-twirp';
import Axios from 'axios';

const getServiceMethodName = (fn: any): string => {
	if (fn === dss.Registration.prototype.add) {
		return 'Add';
    }
	if (fn === dss.Registration.prototype.get) {
		return 'Get';
    }
	if (fn === dss.Registration.prototype.getSummary) {
		return 'GetSummary';
    }
	if (fn === dss.Registration.prototype.prices) {
		return 'Prices';
    }
	if (fn === dss.Registration.prototype.update) {
		return 'Update';
    }

    throw new Error('Unknown Method');
};


export const RegistrationPathPrefix = '/twirp/dss.Registration/';

export const createRegistration = (baseURL: string, options = {}): dss.Registration => {
    const defaultOpts = {
        baseURL: baseURL + RegistrationPathPrefix,
        headers: {
          Accept: 'application/protobuf'
        }
    };
    const axiosOpts = { ...defaultOpts, ...options };
    
    const axios = Axios.create(axiosOpts);

    return dss.Registration.create(createTwirpAdapter(axios, getServiceMethodName));
};
