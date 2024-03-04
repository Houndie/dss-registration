import React from 'react'

const GoogleMap: React.FC<{ title: string, link: string }> = ({ title, link }) => (
	<iframe title={title} className="embed-responsive-item" src={link} width="800" height="450" frameBorder="0" style={{ border: 0 }} allowFullScreen />
)

export default GoogleMap
