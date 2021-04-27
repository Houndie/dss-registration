import React from 'react'

interface CompItemProps {
	title: string
	results: React.ReactNode[]
	links: React.ReactNode[]
}

const CompItem = ({title, results, links}: CompItemProps) => {
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

interface Team {
	name: string
	url: string
	collegiateChamp: boolean
}

interface TeamLinks {
	score?: string
}

interface TeamCompProps {
	teams: Team[]
	isOldStyle?: boolean
	links?: TeamLinks
}

const TeamComp = ({teams, isOldStyle, links}: TeamCompProps) => {
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

export const team = (name: string, url: string, collegiateChamp: boolean = false) => {
	return {
		name: name,
		url: url,
		collegiateChamp: collegiateChamp
	}
}

interface MixAndMatchLinks {
	fPrelimScore?: string
	lPrelimScore?: string
	prelimVideo?: string
	finalsScore?: string
	finalsVideo?: string
}

interface MixAndMatchProps {
	type?: string
	competitors: string[]
	links: MixAndMatchLinks
}

const MixAndMatch = ({type, competitors, links}: MixAndMatchProps) => {
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
			results={competitors.map(c => <>{c}</>)}
			links={linkList}
		/>
	)
}

export const MixAndMatchOld = "Jack and Jill"
export const MixAndMatchOpen = "Open Jack And Jill"
export const MixAndMatchAmateur = "Amateur Jack And Jill"

interface SoloJazzLinks {
	prelimScore?: string
	prelimVideo?: string
	finalsScore?: string
	finalsVideo?: string
}

interface SoloJazzProps {
	competitors: string[]
	links: SoloJazzLinks
}

const SoloJazz = ({competitors, links}: SoloJazzProps) => {
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
			results={competitors.map(c => <>{c}</>)}
			links={linkList}
		/>
	)
}

interface CompResultsProps {
	team?: TeamCompProps
	mixAndMatches?: MixAndMatchProps[]
	solo?: SoloJazzProps
}

const CompResults = ({team, mixAndMatches, solo}: CompResultsProps) => (
	<>
		{team && (<TeamComp isOldStyle={team.isOldStyle} teams={team.teams} links={team.links}/>)}
		{mixAndMatches && mixAndMatches.map((mixAndMatch) => (<MixAndMatch type={mixAndMatch.type} competitors={mixAndMatch.competitors} links={mixAndMatch.links} />))}
		{solo && (<SoloJazz competitors={solo.competitors} links={solo.links}/>)}
	</>
)

export default CompResults
