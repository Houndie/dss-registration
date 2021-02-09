import React from 'react'
import Page from '../../components/Page.js'
import CompResults, {team, MixAndMatchAmateur, MixAndMatchOpen} from '../../components/CompResults.js'

const C2016 = () => (
	<Page title="2016 Competition Results">
		<CompResults 
			team={{
				teams: [
					team("Naptown", "https://www.youtube.com/watch?v=eRfF_oX8yOs&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS&index=1"),
					team("CBUS", "https://www.youtube.com/watch?v=_Pj21YhLL-s&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS&index=2"),
					team("CTA Blueline", "https://www.youtube.com/watch?v=GMGA0q0jH8w&index=3&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS"),
					team("OSU Jitterbucks", "https://www.youtube.com/watch?v=g0cjYLFl95Q&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS&index=4", true),
					team("Rapid Rhythms", "https://www.youtube.com/watch?v=TY-CPfevhDQ&index=12&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS"),
					team("Daytonnati", "https://www.youtube.com/watch?v=dXadkqpdvhk&index=8&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS"),
					team("Tenacious Swing", "https://www.youtube.com/watch?v=tCgvq2SqCY8&index=11&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS"),
					team("Lindy 500", "https://www.youtube.com/watch?v=ma_BTzSkQ90&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS&index=7"),
					team("Knoxville Lindy Hoppers", "https://www.youtube.com/watch?v=wCMt2MQVMU4&index=10&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS"),
					team("Swing Out Loud", "https://www.youtube.com/watch?v=OcTH-AsA_1w&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS&index=9")
				],
				links: {
					score: "/images/competitions/2016_team.png"
				}
			}}
			mixAndMatches={[{
				type: MixAndMatchOpen,
				competitors: [
					"James Pack & Kerry Kapaku",
					"David Barry & Ali Lodico",
					"Yulai Liu & Cansu Bozkus",
					"Kenneth Cebrian & Hilary-Lynn McCabe",
					"David Deenik & Celia Mooradian",
					"Daniel Hoy & Cassie Stoa",
					"Kemper Talley & Amanda Guieb",
					"Emory Thompson & Binaebi Calkins"
				],
				links: {
					fPrelimScore: "/images/competitions/2014_jack_and_jill_follower.png",
					lPrelimScore: "/images/competitions/2014_jack_and_jill_leader.png",
					prelimVideo: "http://www.youtube.com/watch?v=oMGIqs3rgrA&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&feature=share&index=6",
					finalsScore: "/images/competitions/2014_jack_and_jill.png",
					finalsVideo: "http://www.youtube.com/watch?v=bt1BGzXZpK8&list=PLw2dfcFL5AM7PdzmUM1Yry8xRVnVqH8WA&feature=share&index=4"
				}
			},{
				type: MixAndMatchAmateur,
				competitors: [
					"Stanley Steers & Ismene Potakis",
					"Viktor Lillard & Malory Bertisky",
					"Tony Goldsmith & Rebecca Combs",
					"Matthew Keller & Mellissa Rutherford",
					"Chelsea Oswald & Jacqueline George"
				],
				links: {
					lPrelimScore: "/images/competitions/2016_jack_and_jill_amateur_leader.png",
					fPrelimScore: "/images/competitions/2016_jack_and_jill_amateur_follower.png",
					finalsScore: "/images/competitions/2016_jack_and_jill_amateur.png",
					finalsVideo: "https://www.youtube.com/watch?v=2nMWOxp3_wY&index=13&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS"
				}
			}]}
			solo={{
				competitors: [
					"James Pack",
					"Emory Thompson",
					"PJ Ryan",
					"Yulai Liu",
					"Cassie Stoa"
				],
				links: {
					prelimVideo: "https://www.youtube.com/watch?v=3rlHLlJULis&index=15&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS",
					finalsScore: "/images/competitions/2016_solo.png",
					finalsVideo: "https://www.youtube.com/watch?v=ZXYxV9sA0Vk&list=PLw2dfcFL5AM6ooifDSDifw0zV7Z0p1eMS&index=6"
				}
			}}
		/>
	</Page>
)

export default C2016
