import React from 'react'

const CompItem = ({title, results, links}) => {
	return (
	<>
		<h2>{title}</h2>
		<ol>
			{results.map((result, idx) => (
				<li key={idx}>{result}</li>
			))}
		</ol>
		{(links && links.length) ? (
			<>
				<h3>Links:</h3>
				<ul>
					{links.map((link, idx) => (
						<li key={idx}>{link}</li>
					))}
				</ul>
			</>
		) : null}
	</>
)}

const TeamComp = ({teams, isOldStyle, links}) => {
	let linkList = []
	if(links) {
		if(links.score){
			linkList.push((<a href={links.score}>Scoresheet</a>))
		}
	}
	return (
		<CompItem
			title={"Team Competition " + (isOldStyle ? "(Choreography + Strictly)" : "")}
			results={teams.map((item) => (
				<>{item.name} {item.collegiateChamp && (<i>Collegiate Cup Champions</i>)} (<a href={item.url}>video</a>)</>
				))}
			links={linkList}
		/>
	)
}

export const team = (name, url, collegiateChamp) => {
	return {
		name: name,
		url: url,
		collegiateChamp: collegiateChamp
	}
}

const MixAndMatch = ({type, competitors, links}) => {
	let linkList = []
	if(links){
		if(links.fPrelimScore){
			linkList.push((<a href={links.fPrelimScore}>Follower Prelim Scoresheet</a>))
		}
		if(links.lPrelimScore){
			linkList.push((<a href={links.lPrelimScore}>Leader Prelim Scoresheet</a>))
		}
		if(links.prelimVideo){
			linkList.push((<a href={links.prelimVideo}>Prelim Video</a>))
		}
		if(links.finalsScore){
			linkList.push((<a href={links.finalsScore}>Finals Scoresheet</a>))
		}
		if(links.finalsVideo){
			linkList.push((<a href={links.finalsVideo}>Finals Video</a>))
		}
	}
	return (
		<CompItem
			title={type ? type : "Mix and Match"}
			results={competitors}
			links={linkList}
		/>
	)
}

export const MixAndMatchOld = "Jack and Jill"
export const MixAndMatchOpen = "Open Jack And Jill"
export const MixAndMatchAmateur = "Amateur Jack And Jill"

const SoloJazz = ({competitors, links}) => {
	let linkList = []
	if(links) {
		if(links.prelimScore){
			linkList.push((<a href={links.prelimScore}>Prelim Scoresheet</a>))
		}
		if(links.prelimVideo){
			linkList.push((<a href={links.prelimVideo}>Prelim Video</a>))
		}
		if(links.finalsScore){
			linkList.push((<a href={links.finalsScore}>Finals Scoresheet</a>))
		}
		if(links.finalsVideo){
			linkList.push((<a href={links.finalsVideo}>Finals Video</a>))
		}
	}
	return (
		<CompItem
			title="Solo Jazz"
			results={competitors}
			links={linkList}
		/>
	)
}

const CompResults = ({team, mixAndMatches, solo}) => (
	<>
		{team && (<TeamComp isOldStyle={team.isOldStyle} teams={team.teams} links={team.links}/>)}
		{mixAndMatches && mixAndMatches.map((mixAndMatch) => (<MixAndMatch type={mixAndMatch.type} competitors={mixAndMatch.competitors} links={mixAndMatch.links} />))}
		{solo && (<SoloJazz competitors={solo.competitors} links={solo.links}/>)}
	</>
)

export default CompResults
