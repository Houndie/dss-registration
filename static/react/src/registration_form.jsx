import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';
import { Formik, Form, useField } from 'formik';
import { v4 as uuidv4 } from 'uuid';
//import createRegistrationClient from "../rpc/registration_pb_twirp.js";
const registrationTwirp = require("./rpc/registration_pb_twirp.js")

'use strict';

//const siteBase = document.getElementById("hugoVariables").getAttribute('data-sitebase');
//const backend = document.getElementById("hugoVariables").getAttribute('data-backend');

const errString = (err) => {
	return JSON.stringify({ code: err.code, msg: err.message, meta: err.meta }, null, 2)
}

const registrationClient = registrationTwirp.createRegistrationClient(backend);

function parseDollar(intCost) {
	let dollar = intCost.toString()
	while(dollar.length < 3) {
		dollar = "0" + dollar;
	}
	return "$" + dollar.slice(0, -2) + "." + dollar.slice(-2)
}

function tierToString(tier) {
	switch (tier) {
	case registrationTwirp.FullWeekendPassTier.TIER1:
		return "Tier 1";
	case registrationTwirp.FullWeekendPassTier.TIER2:
		return "Tier 2";
	case registrationTwirp.FullWeekendPassTier.TIER3:
		return "Tier 3";
	case registrationTwirp.FullWeekendPassTier.TIER4:
		return "Tier 4";
	case registrationTwirp.FullWeekendPassTier.TIER5:
		return "Tier 5";
	}
}

const DSSTextInput = ({ label, ...props }) => {
	const [field, meta] = useField(props);
	return (
		<>
			<div className="form-group">
				<label htmlFor={props.id || props.name}>{label}</label>
				<br />
				<input className="form-control" {...field} {...props} />
				{meta.touched && meta.error ? (
					<>
						<br />
						<div className="error">{meta.error}</div>
					</>
				) : null}
			</div>
		</>
	);
};

const DSSTextarea = ({ label, ...props }) => {
	const [field, meta] = useField(props);
	return (
		<>
			<div className="form-group">
				<label htmlFor={props.id || props.name}>{label}</label>
				<br />
				<textarea className="form-control" {...field} {...props} />
				{meta.touched && meta.error ? (
					<>
						<br />
						<div className="error">{meta.error}</div>
					</>
				) : null}
			</div>
		</>
	);
};

const DSSSelect = ({ label, ...props }) => {
	const [field, meta] = useField(props);
	return (
		<>
			<div className="form-group">
				<label htmlFor={props.id || props.name}>{label}</label>
				<br />
				<select className="form-control" {...field} {...props} />
				{meta.touched && meta.error ? (
					<>
						<br />
						<div className="error">{meta.error}</div>
					</>
				) : null}
			</div>
		</>
	);
};

const DSSCheckbox = ({ children, ...props }) => {
	const [field, meta] = useField({ ...props, type: 'checkbox' });
	return (
		<>

			<div className="form-group form-check">
				<label className="form-check-label">
					<input type="checkbox" className="form-check-input" {...field} {...props} />
					{children}
				</label>
				{meta.touched && meta.error ? (
					<>
						<br />
						<div className="error">{meta.error}</div>
					</>
				) : null}
			</div>
		</>
	);
};

const fullWeekendPassOption = "Full";
const danceOnlyPassOption = "Dance";
const noPassOption = "None";

const provideOption = "Provide";
const requireOption = "Require";
const noHousingOption = "None";

const RegistrationForm = () => {
	const [prices, setPrices] = useState(null)
	useEffect(() => {
		registrationClient.prices(new registrationTwirp.RegistrationPricesReq()).then(res => {
			setPrices(res);
		}, err => {
			console.error(errString(err));
			window.location.href = siteBase+'/error/';	
		});
	}, [])

	return (
		<Formik
			initialValues={{
				firstName: '', 
				lastName: '',
				streetAddress: '',
				city: '',
				state: '',
				zipCode: '',
				email: '',
				homeScene: '',
				isStudent: false,
				passType: noPassOption,
				level: '',
				mixAndMatch: false,
				role: '',
				soloJazz: false,
				teamCompetition: false,
				teamName: '',
				tshirt: false,
				style: '',
				housing: noHousingOption,
				pets: '',
				quantity: 0,
				provideDetails: '',
				petAllergies: '',
				requireDetails: '',
			}}
			onSubmit={(values, { setSubmitting }) => {

				const r = new registrationTwirp.RegistrationInfo();
				r.setFirstName(values.firstName);
				r.setLastName(values.lastName);
				r.setStreetAddress(values.streetAddress);
				r.setCity(values.city);
				r.setState(values.state);
				r.setZipCode(values.zipCode);
				r.setEmail(values.email);
				r.setHomeScene(values.homeScene);
				r.setIsStudent(values.isStudent);

				switch (values.passType) {
					case fullWeekendPassOption: {
						const f = new registrationTwirp.FullWeekendPass();
						f.setTier(prices.weekendPassTier);
						f.setLevel(parseInt(values.level));
						r.setFullWeekendPass(f);
					}
					case danceOnlyPassOption:
						r.setDanceOnlyPass(new registrationTwirp.DanceOnlyPass());
					default:
						r.setNoPass(new registrationTwirp.NoPass());
				}

				if (values.mixAndMatch) {
					const m = new registrationTwirp.MixAndMatch();
					m.setRole(parseInt(values.role));
					r.setMixAndMatch(m);
				}

				if (values.soloJazz) {
					r.setSoloJazz(new registrationTwirp.SoloJazz());
				}

				if (values.teamCompetition) {
					const t = new registrationTwirp.TeamCompetition();
					t.setName(values.teamName);
					r.setTeamCompetition(t);
				}

				if (values.tshirt) {
					const t = new registrationTwirp.TShirt();
					t.setStyle(parseInt(values.style));
					r.setTshirt(t);
				}

				switch (values.housing) {
					case provideOption: {
						const h = new registrationTwirp.ProvideHousing();
						h.setPets(values.pets);
						h.setQuantity(values.quantity);
						h.setDetails(values.provideDetails);
						r.setProvideHousing(h);
					}
					case requireOption: {
						const h = new registrationTwirp.RequireHousing();
						h.setPetAllergies(values.petAllergies);
						h.setDetails(values.requireDetails);
						r.setRequireHousing(h);
					}
					default:
						r.setNoHousing(new registrationTwirp.NoHousing());
				}

				const req = new registrationTwirp.RegistrationAddReq();
				req.setIdempotencyKey(uuidv4());
				req.setRedirectUrl(siteBase+"/registration-complete");
				req.setRegistration(r);

				return registrationClient.add(req).then(
					res => {
						window.location.href = res.redirectUrl;	
					}, err => {
						console.error(errString(err));
						window.location.href = siteBase+'/error/';	
					}
				);
			}}
		>
		{props => prices != null ? (
			<Form id="formElement">
				<fieldset>
					<h2>Personal Information</h2>
					<div className="form-row">
						<div className="col">
							<DSSTextInput label="First Name" name="firstName" type="text" />
						</div>
						<div className="col">
							<DSSTextInput label="Last Name" name="lastName" type="text" />
						</div>
					</div>
					<DSSTextInput label="Street Address" name="streetAddress" type="text" />
					<div className="form-row">
						<div className="col-8"> 
							<DSSTextInput label="City" name="city" type="text" />
						</div>
						<div className="col-1"> 
							<DSSSelect label="State" name="state">
								<option value=""></option>
								<option value="AK">AK</option>
								<option value="AL">AL</option>
								<option value="AR">AR</option>
								<option value="AZ">AZ</option>
								<option value="CA">CA</option>
								<option value="CO">CO</option>
								<option value="CT">CT</option>
								<option value="DE">DE</option>
								<option value="FL">FL</option>
								<option value="GA">GA</option>
								<option value="HI">HI</option>
								<option value="IA">IA</option>
								<option value="ID">ID</option>
								<option value="IL">IL</option>
								<option value="IN">IN</option>
								<option value="KS">KS</option>
								<option value="KY">KY</option>
								<option value="LA">LA</option>
								<option value="MA">MA</option>
								<option value="MD">MD</option>
								<option value="ME">ME</option>
								<option value="MI">MI</option>
								<option value="MN">MN</option>
								<option value="MO">MO</option>
								<option value="MS">MS</option>
								<option value="MT">MT</option>
								<option value="NC">NC</option>
								<option value="ND">ND</option>
								<option value="NE">NE</option>
								<option value="NH">NH</option>
								<option value="NJ">NJ</option>
								<option value="NM">NM</option>
								<option value="NV">NV</option>
								<option value="NY">NY</option>
								<option value="OH">OH</option>
								<option value="OK">OK</option>
								<option value="OR">OR</option>
								<option value="PA">PA</option>
								<option value="RI">RI</option>
								<option value="SC">SC</option>
								<option value="SD">SD</option>
								<option value="TN">TN</option>
								<option value="TX">TX</option>
								<option value="UT">UT</option>
								<option value="VA">VA</option>
								<option value="VT">VT</option>
								<option value="WA">WA</option>
								<option value="WI">WI</option>
								<option value="WV">WV</option>
								<option value="WY">WY</option>
							</DSSSelect>
						</div>
						<div className="col-3"> 
							<DSSTextInput label="Zip Code" name="zipCode" type="text" />
						</div>
					</div>
					<DSSTextInput label="Email" name="email" type="email" />
					<DSSTextInput label="Home Scene" name="homeScene" type="text" />
					<DSSCheckbox name="isStudent">I am a student</DSSCheckbox>
					<hr />
				</fieldset>
				<fieldset>
					<h2>Purchase</h2>
					<div className="col-6">
						<DSSSelect label="Weekend Pass Type" name="passType">
							<option value={noPassOption} />
							<option value={fullWeekendPassOption}>{"Full Weekend Pass - "+tierToString(prices.weekendPassTier)+" ("+parseDollar(prices.weekendPassCost)+")"}</option>
							<option value={danceOnlyPassOption}>{"Dance Only Pass ("+parseDollar(prices.dancePassCost)+")"}</option>
						</DSSSelect>
					</div>
					{props.values.passType === fullWeekendPassOption ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSSelect label="Level" name="level">
									<option value="" />
									<option value={registrationTwirp.FullWeekendPassLevel.LEVEL1}>Level 1</option>
									<option value={registrationTwirp.FullWeekendPassLevel.LEVEL2}>Level 2</option>
									<option value={registrationTwirp.FullWeekendPassLevel.LEVEL3}>Level 3</option>
								</DSSSelect>
							</div>
						</div>
					) : null}
					<DSSCheckbox name="mixAndMatch">{"Mix & Match Competition ("+parseDollar(prices.mixAndMatchCost)+")"}</DSSCheckbox>
					{props.values.mixAndMatch ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSSelect label="Role" name="role">
									<option value="" />
									<option value={registrationTwirp.MixAndMatch.Role.FOLLOWER}>Follower</option>
									<option value={registrationTwirp.MixAndMatch.Role.LEADER}>Leader</option>
								</DSSSelect>
							</div>
						</div>
					) : null}
					<DSSCheckbox name="soloJazz">{"Solo Jazz Competition ("+parseDollar(prices.soloJazzCost)+")"}</DSSCheckbox>
					<DSSCheckbox name="teamCompetition">{"Team Competition ("+parseDollar(prices.teamCompetitionCost)+")"}</DSSCheckbox>
					{props.values.teamCompetition ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSTextInput label="Team Name" name="teamName" type="text" />
							</div>
						</div>
					) : null}
					<DSSCheckbox name="tshirt">{"T-Shirt ("+parseDollar(prices.tshirtCost)+")"}</DSSCheckbox>
					{props.values.tshirt ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSSelect label="T-Shirt Size/Style" name="style">
									<option value=''></option>
									<option value={registrationTwirp.TShirt.Style.UNISEXS}>Unisex S</option>
									<option value={registrationTwirp.TShirt.Style.UNISEXM}>Unisex M</option>
									<option value={registrationTwirp.TShirt.Style.UNISEXL}>Unisex L</option>
									<option value={registrationTwirp.TShirt.Style.UNISEXXL}>Unisex XL</option>
									<option value={registrationTwirp.TShirt.Style.UNISEX2XL}>Unisex 2XL</option>
									<option value={registrationTwirp.TShirt.Style.UNISEX3XL}>Unisex 3XL</option>
									<option value={registrationTwirp.TShirt.Style.BELLAS}>Bella S</option>
									<option value={registrationTwirp.TShirt.Style.BELLAM}>Bella M</option>
									<option value={registrationTwirp.TShirt.Style.BELLAL}>Bella L</option>
									<option value={registrationTwirp.TShirt.Style.BELLAXL}>Bella XL</option>
									<option value={registrationTwirp.TShirt.Style.BELLA2XL}>Bella 2XL</option>
								</DSSSelect>
							</div>
					</div>
				) : null}
				<hr />
				</fieldset>
				<fieldset>
					<h2>Housing</h2>
					<DSSSelect label="Housing Status" name="housing">
						<option value={noHousingOption}>I neither require nor can provide housing</option>
						<option value={provideOption}>I can provide housing</option>
						<option value={requireOption}>I require housing</option>
					</DSSSelect>
					{props.values.housing === "Provide" ? (
						<>
						<DSSTextInput label="I have the following pets (cats, dogs, etc)" name="pets" type="text" />
						<DSSTextInput label="I can house this many people" name="quantity" type="number" />
						<DSSTextarea label="Any other details the organizers should know about my house" name="provideDetails" />
						</>
					) : (props.values.housing === "Require" ? (
						<>
						<DSSTextInput label="I am allergic to the following pets" name="petAllergies" type="text" />
						<DSSTextarea label="Anything else I would like to say about my housing request" name="requireDetails" />
						</>
					) : null)}
				</fieldset>
				<button type="submit" className="btn btn-info">Submit</button>
			</Form>
		) : null}
		</Formik>
	);
};

ReactDOM.render(<RegistrationForm />, document.getElementById('registration_form'));
