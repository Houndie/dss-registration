import React from 'react';
import ReactDOM from 'react-dom';
import { Formik, Form, useField } from 'formik';

'use strict';

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

const RegistrationForm = () => {
	return (
		<Formik
			initialValues={{firstName: '', lastName: ''}}
			onSubmit={(values, { setSubmitting }) => {
				setTimeout(() => {
				alert(JSON.stringify(values, null, 2));
				setSubmitting(false);
				}, 400);
			}}
		>
		{props => (
			<Form>
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
							<DSSTextInput label="Zip Code" name="zip code" type="text" />
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
							<option value="None" />
							<option value="Full">Full Weekend Pass</option>
							<option value="Dance">Dance Only Pass</option>
						</DSSSelect>
					</div>
					{props.values.passType === "Full" ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSSelect label="Level" name="level">
									<option value="None" />
									<option value="Level1">Level 1</option>
									<option value="Level2">Level 2</option>
									<option value="Level3">Level 3</option>
								</DSSSelect>
							</div>
						</div>
					) : null}
					<DSSCheckbox name="mixAndMatch">Mix & Match Competition</DSSCheckbox>
					{props.values.mixAndMatch ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSSelect label="Role" name="level">
									<option value="None" />
									<option value="Follower">Follower</option>
									<option value="Leader">Leader</option>
								</DSSSelect>
							</div>
						</div>
					) : null}
					<DSSCheckbox name="soloJazz">Solo Jazz Competition</DSSCheckbox>
					<DSSCheckbox name="teamCompetition">Team Competition</DSSCheckbox>
					{props.values.teamCompetition ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSTextInput label="Team Name" name="teamName" type="text" />
							</div>
						</div>
					) : null}
					<DSSCheckbox name="tshirt">T-Shirt</DSSCheckbox>
					{props.values.tshirt ? (
						<div className="form-row">
							<div className="col-1"></div>
							<div className="col-6">
								<DSSSelect label="T-Shirt Size/Style" name="style">
									<option value=""></option>
									<option value="Unisex S">Unisex S</option>
									<option value="Unisex M">Unisex M</option>
									<option value="Unisex L">Unisex L</option>
									<option value="Unisex XL">Unisex XL</option>
									<option value="Unisex 2XL">Unisex 2XL</option>
									<option value="Unisex 3XL">Unisex 3XL</option>
									<option value="Bella S">Bella S</option>
									<option value="Bella M">Bella M</option>
									<option value="Bella L">Bella L</option>
									<option value="Bella XL">Bella XL</option>
									<option value="Bella 2XL">Bella 2XL</option>
								</DSSSelect>
							</div>
					</div>
				) : null}
				<hr />
				</fieldset>
				<fieldset>
					<h2>Housing</h2>
					<DSSSelect label="Housing Status" name="housing">
						<option value="None">I neither require nor can provide housing</option>
						<option value="Provide">I can provide housing</option>
						<option value="Require">I require housing</option>
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
		)}
		</Formik>
	);
};

ReactDOM.render(<RegistrationForm />, document.getElementById('registration_form'));
