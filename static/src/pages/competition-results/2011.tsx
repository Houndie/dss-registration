import React from 'react'
import Page from '../../components/Page'
import CompResults, { team, MixAndMatchOld } from '../../components/CompResults'

const C2011 = () => (
	<Page title="2011 Competition Results">
		<CompResults
			team={{
				isOldStyle: true,
				teams: [
					team("SwingColumbus", "http://youtu.be/8R21voAJZ_4"),
					team("UD Swing Club", "http://youtu.be/H1oUE3lNgrY", true),
					team("Ann Arbor", "http://youtu.be/VN3eSMvKTjQ"),
					team("Rhythm Cats", "http://youtu.be/3euiIrkp8E8"),
					team("OSU Jitterbucks", "http://youtu.be/xW2e04AFBlE"),
					team("Miami Swingers", "http://youtu.be/IryAc9iTAJY"),
					team("Red Hawks", "http://youtu.be/iUCysvqpm3Q"),
					team("Ohio University Swing Club", "http://youtu.be/NRGe2YNTbLA")
				],
			}}
			mixAndMatches={[{
				type: MixAndMatchOld,
				competitors: [
					"Chris Schoenfelder & Amanda Guieb",
					"Dan Rosenthal & Ellen Huffman",
					"Brian Tietz & Ali Lodico",
					"Jeff Johnston & Lisa Huneke",
					"Aleks Daskalov & Heather Lemire",
					"Cal Lin & Dani Dowler"
				],
				links: {
					finalsVideo: "http://youtu.be/yAHPOkDGhy4"
				}
			}]}
		/>
	</Page>
)

export default C2011
