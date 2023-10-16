import React from 'react'
import Page from '../../components/Page'
import CompResults, { team, MixAndMatchOld } from '../../components/CompResults'

const C2014 = () => (
	<Page title="2014 Competition Results">
		<CompResults
			team={{
				teams: [
					team("SwingColumbus", "http://www.youtube.com/watch?v=XncSyzSMIuw&feature=share&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA"),
					team("Tenacious Swing", "http://www.youtube.com/watch?v=omfVkBaAD2c&feature=share&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&index=1", true),
					team("Swing Dings", "http://www.youtube.com/watch?v=C8yUA9nr8SY&feature=share&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&index=2"),
					team("OSU Jitterbucks", "http://www.youtube.com/watch?v=RD7hab7exWo&feature=share&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&index=3")
				],
				links: {
					score: "/images/competitions/2014_team.png"
				}
			}}
			mixAndMatches={[{
				type: MixAndMatchOld,
				competitors: [
					"Brent Watson & Jenni Bevell",
					"Daniel Hoy & Heather Lemire",
					"Bradley Smith & Allison Lodico",
					"Emmory Thompson & Emily Stienecker",
					"Yulai Liu & Brittany Lewton",
					"Jonathan Fisk & Celia Mooradian"
				],
				links: {
					fPrelimScore: "/images/competitions/2014_jack_and_jill_follower.png",
					lPrelimScore: "/images/competitions/2014_jack_and_jill_leader.png",
					prelimVideo: "http://www.youtube.com/watch?v=oMGIqs3rgrA&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&feature=share&index=6",
					finalsScore: "/images/competitions/2014_jack_and_jill.png",
					finalsVideo: "http://www.youtube.com/watch?v=bt1BGzXZpK8&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&feature=share&index=4"
				}
			}]}
			solo={{
				competitors: [
					"Bradley Smith",
					"Heather Lemire",
					"Andry Rakotomalala",
					"Dillon Grandinette"
				],
				links: {
					prelimVideo: "http://www.youtube.com/watch?v=IcMdu5PKf3Y&amp;feature=share&amp;list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&amp;index=7",
					prelimScore: "/images/competitions/2014_solo.png",
					finalsVideo: "http://www.youtube.com/watch?v=PGKMojh_kEM&amp;list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&amp;feature=share&amp;index=5"
				}
			}}
		/>
	</Page>
)

export default C2014
