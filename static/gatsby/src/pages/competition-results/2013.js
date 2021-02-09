import React from 'react'
import Page from '../../components/Page.js'
import CompResults, {team, MixAndMatchOld} from '../../components/CompResults.js'

const C2013 = () => (
	<Page title="2013 Competition Results">
		<CompResults 
			team={{
				teams: [
					team("SwingColumbus", "http://youtu.be/g1VFLyfAVEA"),
					team("Rhythm Cats", "http://youtu.be/yHfIDWPt5l4"),
					team("Miami Redhawks", "http://youtu.be/UxaRppRLiss", true),
					team("Tenacious Swing", "http://youtu.be/5zBeUNrrGJA"),
					team("Jitterbucks", "http://youtu.be/tjEJDtdZIlI")
				],
				links: {
					score: "/images/competitions/2013_team.jpg"
				}
			}}
			mixAndMatches={[{
				type: MixAndMatchOld,
				competitors: [
					"Chris Schoenfelder & Ada Milby",
					"Jay Benze & Emily Schuhmann",
					"Warren Erath & Anna Young",
					"Dan Hoy & Gail Clendenin",
					"Brian Tietz & Ali Lodico",
					"Dan Young & Brittany Lewton",
					"Brent Watson & Kat Bloom"
				],
				links: {
					fPrelimScore: "/images/competitions/2013_jack_and_jill_follower.jpg",
					lPrelimScore: "/images/competitions/2013_jack_and_jill_leader.jpg",
					finalsScore: "/images/competitions/2013_jack_and_jill.jpg",
					finalsVideo: "http://youtu.be/jptJB735P2g"
				}
			}]}
			solo={{
				competitors: [
					"Emily Schuhmann",
					"Heather Lemire",
					"Jay Benze"
				],
				links: {
					finalsScore: "/images/competitions/2013_solo.jpg"
				}
			}}
		/>
	</Page>
)

export default C2013
