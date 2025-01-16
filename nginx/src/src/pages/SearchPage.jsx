import Search from '../components/Search'

import { Box } from '@mui/material';

function SearchPage() {
	return (
		<>
			<Box minHeight="100vh" display="flex" flexDirection="column" justifyContent="center">
				<Search></Search>
			</Box>
		</>
	)
}

export default SearchPage