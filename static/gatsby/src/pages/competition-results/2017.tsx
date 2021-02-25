import React from 'react'
import Page from '../../components/Page'
import CompResults, {team, MixAndMatchOld} from '../../components/CompResults'

const C2017 = () => (
	<Page title="2017 Competition Results">
		{() =>
			<CompResults 
				team={{
					teams: [
						team("SwingColumbus", "https://youtu.be/SdctsYtOBSI?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy"),
						team("Tri City Six", "https://youtu.be/CM_CH0GujI8?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy"),
						team("STL Live Wires", "https://youtu.be/mDFIKfjUccw?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy"),
						team("Naptown Stomp", "https://youtu.be/k46WPn8Ovt0?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy"),
						team("Swing Out Loud", "https://youtu.be/yIN0lPzMlNo?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy", true),
						team("Tenacious Swing", "https://youtu.be/H0xgvK0i91o?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy"),
						team("Rapid Rhythms", "https://youtu.be/qUx0jf54b2o?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy"),
						team("Knoxville Rhythm Steppers", "https://youtu.be/aU8o7L7lLL0?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy")
					],
					links: {
						score: "/images/competitions/2017_team.png"
					}
				}}
				mixAndMatches={[{
					type: MixAndMatchOld,
					competitors: [
						"Cindiasaurus Rex & David Deenik",
						"Jenni Bevell & Kemper Talley",
						"Sarah Campbell & Viktor Lillard",
						"Kerry Kapaku & Brent Watson",
						"Anna Young & XiaoXing Shi",
					],
					links: {
						fPrelimScore: "/images/competitions/2017_jack_and_jill_follower.png",
						lPrelimScore: "/images/competitions/2017_jack_and_jill_leader.png",
						prelimVideo: "https://youtu.be/gLFlUMvAIqk?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy",
						finalsScore: "/images/competitions/2017_jack_and_jill.png",
						finalsVideo: "https://youtu.be/BZF55m9bol0?list=PLw2dfcFL5AM5IsGCcfS9N13W5OW0lrbpy"
					}
				}]}
			/>
		}
	</Page>
)

export default C2017
