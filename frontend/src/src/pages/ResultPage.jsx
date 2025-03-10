import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import { Typography, Link } from '@mui/material';
import Search from '../components/Search'

const ResultPage = () => {
	const [data, setData] = useState(null);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);
	const [responseTime, setResponseTime] = useState(null);

	const location = useLocation();
	const query = new URLSearchParams(location.search).get('query')

	useEffect(() => {
		const startTime = new Date().getTime();
		const api_host = import.meta.env.VITE_API_HOST
		const api_port = import.meta.env.VITE_API_PORT
		axios.get('https://' + api_host+ ':' + api_port + '/search?query=' + query).then((response) => {
			const endTime = new Date().getTime();
			const timeTaken = endTime - startTime;
			setResponseTime(timeTaken);
			setData(response.data);
			setLoading(false);
		}).catch((error) => {
			setError(error.message);
			setLoading(false);
		});
	}, [location.search]);

	var i = 0;
	return (
		<>
			<Search query={query}></Search>
			{loading && <Typography>Loading...</Typography>}
			{error && <Typography color="error">Error: {error}</Typography>}
			{!loading && !error && data &&
				<>
					<br />
					<Typography>{data.length} results in {responseTime}ms</Typography>
					<ul>
						<br />
						{data.map((page) => (
							<li key={i++}>
								<Typography>
									<Link href={page.Url} sx={{ fontFamily: 'inherit' }}>{page.Title}</Link>
								</Typography>
								<Typography>{page.Description}</Typography>
								{page.Description == "" && <Typography sx={{ fontStyle: 'italic' }}>No description</Typography>}
								<br />
							</li>
						))}
					</ul>
				</>
			}
			{!loading && !error && !data &&
				<>
					<br />
					<Typography>0 results in {responseTime}ms</Typography>
				</>
			}
		</>
	)
}

export default ResultPage