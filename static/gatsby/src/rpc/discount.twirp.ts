// Generated by protoc-gen-twirp_typescript. DO NOT EDIT
import {dss} from './discount.pb';
import {createTwirpAdapter} from 'pbjs-twirp';
import Axios from 'axios';

const getServiceMethodName = (fn: any): string => {
	if (fn === dss.Discount.prototype.add) {
		return 'Add';
    }
	if (fn === dss.Discount.prototype.get) {
		return 'Get';
    }
	if (fn === dss.Discount.prototype.list) {
		return 'List';
    }

    throw new Error('Unknown Method');
};


export const DiscountPathPrefix = '/twirp/dss.Discount/';

export const createDiscount = (baseURL: string, options = {}): dss.Discount => {
    const defaultOpts = {
        baseURL: baseURL + DiscountPathPrefix,
        headers: {
          Accept: 'application/protobuf'
        }
    };
    const axiosOpts = { ...defaultOpts, ...options };
    
    const axios = Axios.create(axiosOpts);

    return dss.Discount.create(createTwirpAdapter(axios, getServiceMethodName));
};
