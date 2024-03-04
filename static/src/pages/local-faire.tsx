import React from 'react'
import Page from '../components/Page'
import GoogleMap from '../components/GoogleMap'

export default () => (
	<Page title="Local Fare">
		<p>See some of our favorite lunch and dinner places on google maps!  Please note that the lunch break is not very long...there's a number of fast food places a short drive away that may be worth travelling to if you aren't sure about making it to your lesson in time!</p>
		<p>We also have a food truck outside the Opera House on Saturday Lunch and outside Elegance in Dance for Saturday Late Night...both great convienent options!</p>
		<GoogleMap title="Local Fare" link="https://www.google.com/maps/d/embed?mid=1Ib2DqXGX-PGvGdTy1Ot_QrPI91k&ehbc=2E312F" />
		<p><i>Be sure to expand the left bar or the map to see a list of favorite places!</i></p>
	</Page>
)
