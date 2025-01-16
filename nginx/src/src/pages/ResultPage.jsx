import Search from '../components/Search'
import { Typography, Link } from '@mui/material';
import { useLocation } from 'react-router-dom';
import React, { useState, useEffect } from 'react';
import axios from 'axios';

function ResultPage() {
	const [data, setData] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);
	const [responseTime, setResponseTime] = useState(null);

	const location = useLocation();

	useEffect(() => {
		const query = new URLSearchParams(location.search).get('query')
		const startTime = new Date().getTime();
		const api_host = import.meta.env.VITE_API_HOST
		const api_port = import.meta.env.VITE_API_PORT
		console.log('http://' + api_host+ ':' + api_port + '/search?query=' + query)
		axios.get('http://' + api_host+ ':' + api_port + '/search?query=' + query).then((response) => {
			const endTime = new Date().getTime();
			const timeTaken = endTime - startTime;
			setResponseTime(timeTaken);
			setData(response.data);
			setLoading(false);
		}).catch((error) => {
			
			setError(error);
			setLoading(false);
		});
	}, [location.search]);

	var i = 0;
	return (
		<>
			<Search></Search>
			{loading && <Typography>Loading...</Typography>}
			{error && <Typography>Error: {error}</Typography>}
			{!loading && !error && data &&
				<>
					<br />
					<Typography>{data.length} results in {responseTime}ms</Typography>
					<ul>
						<br />
						{data.map((page) => (
							<li key={i++}>
								<Link href={page.Url}>{page.Title}</Link>
								<Typography>{page.Url}</Typography>
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