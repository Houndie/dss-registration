import React from 'react'
import Page from '../components/Page.js'

export default () => (
	<Page title="Safety">
		{() => 
			<>
				<p>Dayton Swing Smackdown is committed to providing a safe place for all attendees to dance and enjoy themselves.  To that end, we rely on all guests to obey a code of conduct to make the experience the best it can be for everyone.  Guests who disobey the code of conduct may be given warnings, or barred entry to portions or the entirety of the event.</p>
				<h2>Code of Conduct</h2>
				<ul>
					<li>Dayton Swing Smackdown is open to all people, regardless of orientation, identity, race, appearance or religion.  No attendees should use language deemed discriminatory towards any person or group of people</li>
					<li>Attendees at Dayton Smackdown are allowed to say “no” to dance for any reason</li>
					<li>Do not attempt to teach or give advice to any other attendee, unless they ask for it.</li>
					<li>Harassment of any kind is strictly prohibited.</li>
				</ul>
				<h2>Report a Violation</h2>
				<p>If any attendee feels there is a violation of the code of conduct, feels unsafe or harassed, or has any other safety concerns, they are encouraged to report it!  Please bring any concerns to a Dayton Swing Smackdown staff member.  If you don’t feel comfortable talking to a staff member, <a href="/safety-report">you can submit a report here</a>.</p>
			</>
		}
	</Page>
)
