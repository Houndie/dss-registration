import React from 'react'
import Page from '../../components/Page'
import CompResults, { team } from '../../components/CompResults'

export default () => (
	<Page title="2024 Competition Results">
		<CompResults
			team={{
				teams: [
					team("Windy City Lindy Hoppers", "https://www.youtube.com/watch?v=8xno08oMjCs&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=4"),
					team("Traffic Jam Swing", "https://www.youtube.com/watch?v=fZB3BBoS3Ew&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=11"),
					team("STL Live Wires", "https://www.youtube.com/watch?v=aUAlQ5ljbFk&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=9"),
					team("Rapid Rhythms", "https://www.youtube.com/watch?v=P26QG3Lv5XI&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=6"),
					team("Gem City Swing", "https://www.youtube.com/watch?v=tflCdHqgGDU&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=2"),
					team("Swing Columbus", "https://www.youtube.com/watch?v=DdqFQsiPlD4&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=1"),
					team("Naptown Stomp", "https://www.youtube.com/watch?v=nAnNeLXA0jE&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=10"),
					team("Boiling Hot!", "https://www.youtube.com/watch?v=SaFLbMr16rE&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=5", true),
					team("Swing On Sloopy", "https://www.youtube.com/watch?v=R35b50dSuco&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=8"),
					team("Motor City Swing", "https://www.youtube.com/watch?v=qgpQYTz7j4I&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=3"),
					team("Cleveland", "https://www.youtube.com/watch?v=I8zPGSeddJw&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=7"),
				],
			}}
			mixAndMatches={[{
				competitors: [
					"Rachel Kreher & Kemper Talley",
					"Scott Herdegen & Zoe Willingham",
					"Marina Clements & Isaac Swets",
					"Doug Bae & Melissa Rutherford",
					"Ismene Potakis & Ian Shepard",
					"Stanley Steers & Arleigh Wiler",
					"Alina Lusebrink & Desiree Roffers",
				],
				links: {
					finalsVideo: "https://www.youtube.com/watch?v=yPGRlIXojN8&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=12",
					prelimVideo: "https://www.youtube.com/watch?v=Hjr6xGGDTiI&list=PL79I7kj_rd05ywDgg9D8D24pQ_VbNeASM&index=13"
				}
			}]}
			solo={{
				competitors: [
					"Marina Clements",
					"Ben Todd",
					"Sarah Beth"
				],
				links: {}
			}}
		/>
	</Page>
)
