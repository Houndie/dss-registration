import React from 'react'
import Page from '../components/Page'

export default () => (
	<Page title="Covid Policy">
		{() => 
			<>
				<p>Our covid policy is as follows this year:</p> 
				<h2>Vaccines/Covid Tests</h2>
				<p>When attending Dayton Swing Smackdown 2023, you must do one of the following:</p>
				<ul>
					<li><b>Provide evidence of vaccination:</b> You must povide evidence of covid vaccination at the door.  The last shot/booster can be no earlier than February 25, 2022.</li>
					<li><b>Take a covid test when arriving at the event for the first time</b>:  You must take the test while being observed by one of our registration employees.  You can provide a test yourself, we will also have tests for sale at the door. You must then wait for the test to complete before you can enter the dance (usually 15 minutes).  You must remain masked during this waiting period.</li>
				</ul>
				<h2>Masks</h2>
				<p>Masks are optional this year at Smackdown.  That said, we highly encourage their use...swing dancing involves much closer contact than day to day life and n95 and kn95 masks have been shown to reduce covid spread even in close contacts</p>
				<h2>Other Testing</h2>
				<p>While not required for admittance, we strongly encourage all attendees to take a covid test prior to attending the event. They are inexpensive and quick, and you can easily help prevent others from getting sick.</p>
				<h2>Refunds</h2>
				<p>If requested, we will refund your money fully up to the start of the event, and proportionally throughout the course of the event, if requested, for any reason.  If you are sick, stay home, we will refund you.  If you think you might be exposed, stay home, we will refund you.  If you are getting bad vibes, you may stay home and we will refund you.  Refunds are not conditional on having to provide any explanation. Please don't feel obligated to come just because you paid; we want the event to be as safe as possible.</p>
			</>}
	</Page>
)
