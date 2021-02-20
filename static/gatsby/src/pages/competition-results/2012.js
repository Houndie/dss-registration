import React from 'react'
import Page from '../../components/Page.js'
import CompResults, {team, MixAndMatchOld} from '../../components/CompResults.js'

const C2012 = () => (
	<Page title="2012 Competition Results">
		{() => 
			<CompResults 
				team={{
					teams: [
						team("SwingColumbus", "http://youtu.be/09ba3duU9h8"),
						team("Rhythm Cats", "http://youtu.be/yJZ9cpRygP8"),
						team("Jitterbucks", "http://youtu.be/nknd5FOxZVQ", true),
						team("UD Swing Dance Club", "http://youtu.be/a32OZ1c2bc4"),
						team("Miami Swing Syndicate", "http://youtu.be/eDvmshIzUlY"),
						team("The Razz Ma Tazz", "http://youtu.be/dBv2z1GifCg")
					]
				}}
				mixAndMatches={[{
					type: MixAndMatchOld,
					competitors: [
						"Danny Beyrer & Amanda Guieb",
						"Chris Schoenfelder & Brittany Radke",
						"T.J. Sweda & Ali Lodico",
						"Brent Watson & Ada Milby",
						"Jonathan Fisk & Chelsea Dvorchak",
						"Cody Ker & Heather Lemire",
						"Dan Young & Emily Schuhmann"
					],
					links: {
						finalsVideo: "http://youtu.be/OG37nMFjIUA"
					}
				}]}
				solo={{
					competitors: [
						"Chris Schoenfelder",
						"Emily Schuhmann",
						"Ellen McIntire"
					],
					links: {
						finalsVideo: "http://youtu.be/4i2yH04cN4s"
					}
				}}
			/>
		}
	</Page>
)

export default C2012
