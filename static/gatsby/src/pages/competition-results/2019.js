import React from 'react'
import Page from '../../components/Page.js'
import CompResults, {team} from '../../components/CompResults.js'

const C2019 = () => (
	<Page title="2019 Competition Results">
		{() => 
			<CompResults 
				team={{
					teams: [
						team("The STL Live Wires", "https://www.youtube.com/watch?v=g_Lwzy6jCSI&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=3"),
						team("SwingColumbus", "https://www.youtube.com/watch?v=WptTFugF3tc&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=4"),
						team("513 Special", "https://www.youtube.com/watch?v=BgJNV66VH-E&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=5"),
						team("Rapid Rhythms", "https://www.youtube.com/watch?v=ywET0n2FKnc&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=8"),
						team("Gem City Swing", "https://www.youtube.com/watch?v=i2Qsec5y8F0&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=7"),
						team("Swing Cats", "https://www.youtube.com/watch?v=DrrYBJ2Temg&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=6", true)
					],
				}}
				mixAndMatches={[{
					competitors: [
						"Katie Leatherberry & Ike Swets",
						"Katie Dillon & Kemper Talley",
						"Christina Cacciatore & Scott Herdegen",
						"Elizabeth Hokanson & Douglas Bae",
						"Lindsay Kelly & Aleksandr Daskalov (Tied with Below)",
						"Aliceann Talley & Tyedric Hill (Tied with Above)"
					],
					links: {
						finalsVideo: "https://www.youtube.com/watch?v=jjVCKiBdUe8&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=1"
					}
				}]}
				solo={{
					competitors: [
						"Christina Cacciatore",
						"Katie Leatherberry",
						"Elizabeth Hokanson"
					],
					links: {
						finalsVideo: "https://www.youtube.com/watch?v=jr4FFwGZr5k&list=PLw2dfcFL5AM7BKuyq5a2SIASCQVbbvWH5&index=2"
					}
				}}
			/>
		}
	</Page>
)

export default C2019
