import React, { useState } from 'react';
import { useNavigate, Link as RouterLink } from 'react-router-dom';
import { IconButton, Box, TextField, Typography, Link } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import Logo from "../assets/logo.svg?react";

const Search = ({ query }) => {
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
			<Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', mt: 5, mb: 3 }}>
				<Logo style={{ marginRight: 8 }} />
				<Typography variant="h5" color="black">Search engine</Typography>
  			</Box>
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
						value={query}
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
