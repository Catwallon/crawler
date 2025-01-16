import { useNavigate } from 'react-router-dom';
import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { IconButton, Box, TextField, Typography, Link } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';

function Search() {
	const navigate = useNavigate();
	const [searchInput, setSearchInput] = useState('');

	const handleChange = (event) => {
		setSearchInput(event.target.value);
	};

	const search = (e) => {
		e.preventDefault();
		navigate('/result?query=' + searchInput);
	};

	return (
		<>
			<Link component={RouterLink} to="/" sx={{textDecoration: 'none' }}>
				<Typography variant="h5" color="black" sx={{ textAlign: 'center', mt: 5, mb: 3 }}>Search engine</Typography>
			</Link>
			<Box
				sx={{
					display: 'flex',
					justifyContent: 'center',
					gap: 2,
					mb: 5,
					mx: 'auto'
				}}
			>
				<form onSubmit={search}>
					<TextField
						value={searchInput}
						onChange={handleChange}
						placeholder="Tap your search"
						size="small"
						slotProps={{
							input: {
								endAdornment: (
									<IconButton variant="contained" type="submit" aria-label="search">
										<SearchIcon />
									</IconButton>
								)
							}
					}}
					/>
				</form>
			</Box>
		</>
	);
}

export default Search;
