import React from 'react'
import Page from '../../components/Page.js'
import CompResults, {team, MixAndMatchOld} from '../../components/CompResults.js'

const C2015 = () => (
	<Page title="2015 Competition Results">
		{() => 
			<CompResults 
				team={{
					teams: [
						team("SwingColumbus", "https://www.youtube.com/watch?v=mzhJ14Lmmnk&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s&index=1"),
						team("CTA Blueline", "https://www.youtube.com/watch?v=52uw22HKnlw&index=2&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s"),
						team("Daytonnati", "https://www.youtube.com/watch?v=_uRaSVSIR0M&index=3&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s"),
						team("Tenacious Swing", "https://www.youtube.com/watch?v=AO7qT03zF4U&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s&index=4", true),
						team("Dirty Scrubbable Princesses", "https://www.youtube.com/watch?v=xKQ6fE1D0gs&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s&index=6"),
						team("Denver Swing Project", "https://www.youtube.com/watch?v=fodB65e_No4&index=8&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s"),
						team("Demon Shauna Posse", "https://www.youtube.com/watch?v=e7MIkoWL0ds&index=7&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s"),
						team("CASE of Swing Fever", "https://www.youtube.com/watch?v=IpgFeAzQzJA&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s&index=9"),
						team("Knoxville Lindy Hoppers", "https://youtu.be/l0rDXCccgk8?list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s")
					]
				}}
				mixAndMatches={[{
					type: MixAndMatchOld,
					competitors: [
						"Emory Thompson & Madeline Ford",
						"John Holmstrom & Jenni Bevell",
						"Bradley Smith & Kerry Kapaku",
						"David Berry & Brittany Lewton",
						"Kyle Hankins & Ali Lodico",
						"Kemper Talley & Briana Ayers",
						"Ryan Anderson & Heather Mirletz"
					],
					links: {
						finalsScore: "/images/competitions/2015_jack_and_jill.png",
						finalsVideo: "https://www.youtube.com/watch?v=0Lu4DT0eAC4&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s&index=5"
					}
				}]}
				solo={{
					competitors: [
						"Erin Morris",
						"Bradley Smith",
						"Jenni Bevell",
						"Kyle Hankins"
					],
					links: {
						finalsScore: "/images/competitions/2015_solo.png",
						finalsVideo: "https://www.youtube.com/watch?v=CzAJI8mj2Y0&list=PLw2dfcFL5AM56Dy14GztcALInVCb6e83s&index=10"
					}
				}}
			/>
		}
	</Page>
)

export default C2015
