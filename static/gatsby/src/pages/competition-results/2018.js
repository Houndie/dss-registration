import React from 'react'
import Page from '../../components/Page.js'
import CompResults, {team} from '../../components/CompResults.js'

const C2018 = () => (
	<Page title="2018 Competition Results">
		{() => 
			<CompResults 
				team={{
					teams: [
						team("Southeast Scramble", "https://www.youtube.com/watch?v=q1rR_iPhoXg&index=1&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&t=0s"),
						team("The STL Live Wires", "https://www.youtube.com/watch?v=wOWKFQHx1lE&index=2&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X"),
						team("Tacky Annie and the Opus 5", "https://www.youtube.com/watch?v=N9xOhdidl6U&index=3&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X"),
						team("Team SwingColumbus", "https://www.youtube.com/watch?v=3J827VKRpoE&index=7&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X"),
						team("Mitten Magic", "https://www.youtube.com/watch?v=vM0RoXrbwRI&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=8"),
						team("The OSU Rhythm Hoppers", "https://www.youtube.com/watch?v=UAbJEczTGpI&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=4", true),
						team("Gem City Swing", "https://www.youtube.com/watch?v=3zg0hZ15-ik&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=9"),
						team("Rapid Rhythms", "https://www.youtube.com/watch?v=fh8S05Tl0o8&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=10"),
						team("Swing Out LOUD!", "https://www.youtube.com/watch?v=3hRbglzqyPU&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=11"),
						team("All Out of Bubblegum", "https://www.youtube.com/watch?v=4JH4TvOyTSQ&index=12&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X"),
						team("Tenacious Swing", "https://www.youtube.com/watch?v=Y0-AEKPYT6U&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=13")
					],
					teamLinks: {
						score: "/images/competitions/2018_team.png"
					}
				}}
				mixAndMatches={[{
					competitors: [
						"Brittany Morton & Viktor Lillard",
						"Mimi Liu & Christopher Glasow",
						"Kerry Kapaku & Cyle Dixon",
						"Cindiasaurus Rex & Farooq Khan",
						"Aliceann Talley & Jony Navarro",
						"Sarah Jones & Bryan Soto"
					],
					links: {
						fPrelimScore: "/images/competitions/2018_mix_and_match_follower.pdf",
						lPrelimScore: "/images/competitions/2018_mix_and_match_leader.pdf",
						finalsScore: "/images/competitions/2018_mix_and_match.pdf",
						finalsVideo: "https://www.youtube.com/watch?v=ZZoYB42lrow&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=5"
					}
				}]}
				solo={{
					competitors: [
						"Viktor Lillard",
						"Cindiasaurus Rex",
						"Katie Leatherberry",
						"Tyedric Hill",
						"Jony Navarro",
						"Christopher Glasow"
					],
					links: {
						finalsScore: "/images/competitions/2018_solo.pdf",
						finalsVideo: "https://www.youtube.com/watch?v=Icgj4cDhgng&list=PLw2dfcFL5AM7cVFmOgKbbKTzZec9pGI-X&index=6"
					}
				}}
			/>
		}
	</Page>
)

export default C2018
